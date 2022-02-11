/*
 * @Auther: BinyGo
 * @Description:
 * @Date: 2022-02-11 19:19:29
 * @LastEditTime: 2022-02-11 22:04:30
 */
package singleton

import (
	"fmt"
	"sync"
	"unsafe"
)

//单例设计模式（Singleton Design Pattern）
//一个类只允许创建一个对象（或者实例），那这个类就是一个单例类
//这种设计模式就叫作单例设计模式，简称单例模式。

type singleton struct {
	name string
}

var (
	once sync.Once
	//懒汉模式
	//懒汉+once,once中使用atomic,对比饿汉,还是需要多消耗一些资源
	instance *singleton

	//如果不需要传参,不是启动频繁,且不介意初始化就启动的场景可以采用饿汉模式
	//饿汉模式 instance 在类加载的时候，就已经创建并初始化好了
	//所以，instance 实例的创建过程是线程安全的
	//不过，这样的实现方式不支持延迟加载（在真正用到的时候，再创建实例）
	instance2 *singleton = &singleton{"BinyGo"}
)

func GetInstance(name string) *singleton {
	once.Do(func() {
		instance = &singleton{name}
		fmt.Println("create instance singleton", instance)
	})
	fmt.Println(instance, "懒汉 is already created:", unsafe.Pointer(&instance))
	return instance
}

func GetInstance2() *singleton {
	fmt.Println(instance, "饿汉 is already created:", unsafe.Pointer(&instance))
	return instance2
}

/*
1. 单例存在哪些问题？
	单例对 OOP 特性的支持不友好
	单例会隐藏类之间的依赖关系
	单例对代码的扩展性不友好
	单例对代码的可测试性不友好
	单例不支持有参数的构造函数(类似上面懒汉代码例子,仅在第一次调用时的参数会被使用)
2. 单例有什么替代解决方案？
	通过工厂模式、IOC 容器来保证，由程序员自己来保证（自己在编写代码的时候自己保证不要创建两个类对象）。
	有人把单例当作反模式，主张杜绝在项目中使用。这有点极端。模式没有对错，关键看你怎么用。
	如果单例类并没有后续扩展的需求，并且不依赖外部系统，那设计成单例类就没有太大问题。
	对于一些全局的类，我们在其他地方 new 的话，还要在类之间传来传去，不如直接做成单例类，使用起来简洁方便。
3. 如何实现集群环境下的单例？
*/
