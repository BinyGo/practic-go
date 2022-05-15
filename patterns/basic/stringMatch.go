package basic

import "fmt"

// 1. 字符串匹配
// 1.1.1. 字符串匹配
// BF(Brute force)算法
// 实现：每次向后移动一位进行匹配

// RK(Rabin-Karp)算法
// 实现：将每组要匹配长度的字符串进行hash，再hash后的元素里找

// BM(Boyer-Moore)算法
// 有两部分组成：并且是由大到小，倒着匹配
// 坏前缀：普通匹配只一位一位移动，移动规则为 si(坏字符的位置) xi(坏字符在匹配字符最后出现的位置) 都没有xi=-1 移动距离等于si-xi
// 好后缀：坏前缀有可能产生负数，所以还要利用好后缀来进行匹配，好后缀类似坏前缀如果匹配串中有和好后缀相同的子串 ，移动到最靠后的子串的位置，如果没有相同的子串，就需要在匹配的子串中，查找和前缀子串匹配最长的子串进行移动。

// KMP（Knuth Morris Pratt）算法
// 实现：关键部分next数组，失效函数。next数据就是匹配串字符串最长匹配前缀和最长匹配后缀的关系。

func KMPStringMatch() {
	a := "ababaeabacaaaaaddfdfdfdfdf"
	b := "aca"
	pos := Kmp(b, a)
	fmt.Println(pos)
}
func Kmp(needle string, str string) int {
	next := getNext(needle)
	fmt.Println("k", next)
	j := 0
	for i := 0; i < len(str); i++ {
		for j > 0 && str[i] != needle[j] {
			j = next[j-1] + 1
		}
		if str[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i - j + 1
		}
	}
	return -1
}

func getNext(needle string) []int {
	var next = make([]int, len(needle))
	fmt.Println("n", next)
	next[0] = -1
	k := -1
	for i := 1; i < len(needle); i++ {
		for k != -1 && needle[k+1] != needle[i] {
			k = next[k]
		}
		if needle[k+1] == needle[i] {
			k++
		}
		next[i] = k
	}
	return next
}
