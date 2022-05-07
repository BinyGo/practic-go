package main

import (
	"errors"
	"fmt"
)

// 一个经典的场景就是我们在使用事务时，发生错误需要回滚，这时我们就可以用使用defer来保证程序退出时保证事务回滚，
// 代码摘自之前写的 Leaf-segment数据库获取ID方案：https://github.com/asong2020/go-algorithm/blob/master/leaf/dao/leaf_dao.go
// 示例代码如下
/*
func (l *LeafDao) NextSegment(ctx context.Context, bizTag string) (*model.Leaf, error) {
	// 开启事务
	tx, err := l.sql.Begin()
	defer func() {
		if err != nil {
			l.rollback(tx)
		}
	}()
	if err = l.checkError(err); err != nil {
		return nil, err
	}
	err = l.db.UpdateMaxID(ctx, bizTag, tx)
	if err = l.checkError(err); err != nil {
		return nil, err
	}
	leaf, err := l.db.Get(ctx, bizTag, tx)
	if err = l.checkError(err); err != nil {
		return nil, err
	}
	// 提交事务
	err = tx.Commit()
	if err = l.checkError(err); err != nil {
		return nil, err
	}
	return leaf, nil
}
*/

func main() {
	fmt.Println("Anonymous():", Anonymous()) //Anonymous(): 0
	fmt.Println("HasName():", HasName())     //HasName(): 2
	fmt.Println("Demo3():", Demo3())         //HasName(): 1
	fmt.Println("Demo4():", Demo4())         //HasName(): 2
	fmt.Println("Demo5():", Demo5())         //HasName(): 4
	e1()                                     //<nil>
	e2()                                     //e2 defer err
	e3()                                     //<nil>
}

func Demo1() {
	fmt.Println("reciprocal")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	// reciprocal 9 8 7 6 5 4 3 2 1
	// 跟栈一样，即"先进后出"特性，越后面的defer表达式越先被调用。
}

func Demo2() {
	fmt.Println(Sum(1, 2))
}

func Sum(num1, num2 int) int {
	defer fmt.Println("num1:", num1) //1
	defer fmt.Println("num2:", num2) //2
	num1++
	num2++
	return num1 + num2
	//defer将语句放入到栈中时，也会将相关的值拷贝同时入栈。这两个变量并不受num1++、num2++的影响
}

// 匿名函数
func Anonymous() int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2 value is ", i) //defer2 value is  2
	}()

	defer func() {
		i++
		fmt.Println("defer1 in value is ", i) //defer1 in value is  1
	}()

	return i
	//Anonymous()的返回值为0
}

//命名返回值的函数
func HasName() (j int) {
	defer func() {
		j++
		fmt.Println("defer2 in value", j) //defer2 in value 2
	}()

	defer func() {
		j++
		fmt.Println("defer1 in value", j) //defer1 in value 1
	}()

	return j
	//HasName()的返回值为2
}

// go官方文档，文档指出，defer的执行顺序有以下三个规则：
// A deferred function’s arguments are evaluated when the defer statement is evaluated.
// Deferred function calls are executed in Last In First Out order after the surrounding function returns.
// Deferred functions may read and assign to the returning function’s named return values.

func Demo3() (r int) {
	i := 1
	defer func() {
		i = i + 1
	}()
	return i
	//1
}
func Demo4() (r int) {
	defer func(r int) {
		r = r + 2
	}(r)
	return 2
	//2
}
func Demo5() (r int) {
	defer func(r *int) {
		*r = *r + 2
	}(&r)
	return 2
	//4
}

func e1() {
	var err error
	defer fmt.Println(err) //<nil>
	err = errors.New("e1 defer err")
}

func e2() {
	var err error
	defer func() {
		fmt.Println(err) //e2 defer err
	}()
	err = errors.New("e2 defer err")
}

func e3() {
	var err error
	defer func(err error) {
		fmt.Println(err) //<nil>
	}(err)
	err = errors.New("e3 defer err")
}

//原文:https://mp.weixin.qq.com/s?__biz=MzkyNzI1NzM5NQ==&mid=2247484778&idx=1&sn=7ceb16f634b3d479a8d5b0b8c4d50b27&chksm=c22b8336f55c0a20e1099d062a69c16436d3b7cbc9c3da6c43bab34fc2f029e0c48e0cb0f08c&cur_album_id=1932319304830517254&scene=189#wechat_redirect
