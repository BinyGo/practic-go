package basic

import (
	"fmt"
	"sync"
	"time"
) /*

令牌桶算法

算法思想
	想象有一个木桶，以固定的速度往木桶里加入令牌，木桶满了则不再加入令牌。服务收到请求时尝试从木桶中取出一个令牌，
	如果能够得到令牌则继续执行后续的业务逻辑；如果没有得到令牌，直接返回反问频率超限的错误码或页面等，不继续执行后续的业务逻辑

特点：由于木桶内只要有令牌，请求就可以被处理，所以令牌桶算法可以支持突发流量。同时由于往木桶添加令牌的速度是固定的，
	且木桶的容量有上限，所以单位时间内处理的请求书也能够得到控制，起到限流的目的。假设加入令牌的速度为 1token/10ms，桶的容量为500，
	在请求比较的少的时候（小于每10毫秒1个请求）时，木桶可以先"攒"一些令牌（最多500个）。当有突发流量时，一下把木桶内的令牌取空，
	也就是有500个在并发执行的业务逻辑，之后要等每10ms补充一个新的令牌才能接收一个新的请求。

参数设置：木桶的容量 - 考虑业务逻辑的资源消耗和机器能承载并发处理多少业务逻辑。生成令牌的速度 - 太慢的话起不到“攒”令牌应对突发流量的效果。

适用场景：
	适合电商抢购或者微博出现热点事件这种场景，因为在限流的同时可以应对一定的突发流量。
	如果采用均匀速度处理请求的算法，在发生热点时间的时候，会造成大量的用户无法访问，对用户体验的损害比较大。

go语言实现：
	假设每100ms生产一个令牌，按user_id/IP记录访问最近一次访问的时间戳 t_last 和令牌数，
	每次请求时如果 now - last > 100ms, 增加 (now - last) / 100ms个令牌。
	然后，如果令牌数 > 0，令牌数 -1 继续执行后续的业务逻辑，否则返回请求频率超限的错误码或页面。
*/

// 并发访问同一个user_id/ip的记录需要上锁
var recordMu map[string]*sync.RWMutex

func init() {
	recordMu = make(map[string]*sync.RWMutex)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type TokenBucket struct {
	BucketSize int                // 木桶内的容量：最多可以存放多少个令牌
	TokenRate  time.Duration      // 多长时间生成一个令牌
	records    map[string]*record // 报错user_id/ip的访问记录
}

// 上次访问时的时间戳和令牌数
type record struct {
	last  time.Time
	token int
}

func NewTokenBucket(bucketSize int, tokenRate time.Duration) *TokenBucket {
	return &TokenBucket{
		BucketSize: bucketSize,
		TokenRate:  tokenRate,
		records:    make(map[string]*record),
	}
}

func (t *TokenBucket) getUidOrIp() string {
	// 获取请求用户的user_id或者ip地址
	return "127.0.0.1"
}

// 获取这个user_id/ip上次访问时的时间戳和令牌数
func (t *TokenBucket) getRecord(uidOrIp string) *record {
	if r, ok := t.records[uidOrIp]; ok {
		return r
	}
	return &record{}
}

// 保存user_id/ip最近一次请求时的时间戳和令牌数量
func (t *TokenBucket) storeRecord(uidOrIp string, r *record) {
	t.records[uidOrIp] = r
}

// 验证是否能获取一个令牌
func (t *TokenBucket) validate(uidOrIp string) bool {
	// 并发修改同一个用户的记录上写锁
	rl, ok := recordMu[uidOrIp]
	if !ok {
		var mu sync.RWMutex
		rl = &mu
		recordMu[uidOrIp] = rl
	}
	rl.Lock()
	defer rl.Unlock()

	r := t.getRecord(uidOrIp)
	now := time.Now()
	if r.last.IsZero() {
		// 第一次访问初始化为最大令牌数
		r.last, r.token = now, t.BucketSize
	} else {
		if r.last.Add(t.TokenRate).Before(now) {
			// 如果与上次请求的间隔超过了token rate
			// 则增加令牌，更新last
			r.token += max(int(now.Sub(r.last)/t.TokenRate), t.BucketSize)
			fmt.Println("增加令牌,token:", r.token)
			r.last = now
		}
	}
	var result bool
	if r.token > 0 {
		// 如果令牌数大于1，取走一个令牌，validate结果为true
		r.token--
		fmt.Println("取走令牌,token", r.token)
		result = true
	}

	// 保存最新的record
	t.storeRecord(uidOrIp, r)
	return result
}

// 返回是否被限流
func (t *TokenBucket) IsLimited() bool {
	return !t.validate(t.getUidOrIp())
}

func TokenBucketExec() {
	tokenBucket := NewTokenBucket(5, 100*time.Millisecond)
	for i := 0; i < 30; i++ {
		fmt.Println("限流", tokenBucket.IsLimited())
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println(tokenBucket.IsLimited())
}
