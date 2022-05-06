package main

func leak() {
	ch := make(chan struct{})
	go func() {
		ch <- struct{}{}
	}()
	//<-ch //注释,未取ch,将造成内存泄漏,在单元测试中查看详情
}
