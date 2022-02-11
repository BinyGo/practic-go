/*
 * @Auther: BinyGo
 * @Description:
 * @Date: 2022-02-11 19:01:38
 * @LastEditTime: 2022-02-11 20:38:23
 */
package main

func main() {

}

type StorageType int

const (
	DiskStorage StorageType = 1 << iota
	TempStorage
	MemoryStorage
)
