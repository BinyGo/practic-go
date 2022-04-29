package proxy

import "fmt"

type IObject interface {
	ObjDo(action string)
}

type Object struct {
	action string
}

func (obj *Object) ObjDo() {
	fmt.Printf("I can, %s", obj.action)
}

type ProxyObject struct {
	object *Object
}

func (p *ProxyObject) ObjDo(action string) {
	if p.object == nil {
		p.object = &Object{action: "Object Run"}
	}
	if action == "Run" {
		p.object.ObjDo()
	}
}

func Demo1() {
	app := ProxyObject{}
	app.ObjDo("Run")
}
