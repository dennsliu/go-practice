package main

import (
	"fmt"
)

//声明一个角色的结构体，本例中，角色也是 OnSkill 事件的响应处理方
type Actor struct {
}

//为角色结构添加一个OnEvent()方法，这个方法拥有 param 参数，类型为 interface{}，
//与事件系统的函数(func(interface{}))签名一致。
func (a *Actor) OnEvent(param interface{}) {
	fmt.Println("actor event:", param)
}

//GolobalEvent()函数为全局响应事件。
//有时需要全局进行侦听或者处理一些事件，这里使用普通函数实现全局事件的处理。
func GlobalEvent(param interface{}) {
	fmt.Println("global event:", param)
}

//创建一个 map 实例，这个 map 通过事件名(string)关联回调列表([]func(interace{})),
//同一个事件名称可能存在多个事件回调，因此使用回调列表保存。
//回调的函数声明为 func(interface{})。
var eventByName = make(map[string][]func(interface{}))

//注册事件，提供事件名和回调函数。
//提供给外部的通过事件注册响应函数的入口。
func RegisterEvent(name string, callback func(interface{})) {

	//通过名字查找事件列表
	//eventByName 通过事件名（name）进行查询，返回回调列表（[]func(interface{})）
	list := eventByName[name]

	//在列表切片中添加函数
	//为同一个事件名称在已经注册的事件回调的列表中再添加一个回调函数。
	list = append(list, callback)

	//将修改后的函数列表设置到 map 的对应事件中。
	eventByName[name] = list
}

//调用事件的入口。提供事件名称name和参数param。
//事件的参数表示描述事件具体的细节，例如门打开的事件触发时，参数可以传入谁进来了。
func CallEvent(name string, param interface{}) {

	//通过注册事件回调的 eventByName 和事件名字查询处理函数列表 list。
	list := eventByName[name]

	//遍历这个事件列表，如果没有找到对应的事件，list将是一个空切片。
	for _, callback := range list {

		//将每个函数回调转入事件参数并调用，就会触发事件实现方的逻辑处理。
		callback(param)
	}
}

func main() {

	//实例化一个角色
	a := new(Actor)

	//注册一个 OnSkill 事件，实现代码由 a 的 OnEvent 进行处理。也就是 Actor 的 OnEvent()方法。
	RegisterEvent("Onskill", a.OnEvent)

	//再次注册一个 OnSkill 事件，实现代码由 GlobalEvent 进行处理。
	//虽然注册了2次，也注册的同一个名字，但前面注册的事件不会被覆盖，而是被添加到事件系统中，关联 OnSkill 事件的函数列表中。
	RegisterEvent("Onskill", GlobalEvent)

	//模拟处理事件，通过 CallEvent() 函数传入两个参数，第一个为事件名，第二个为处理函数的参数。
	CallEvent("Onskill", 100)
}
