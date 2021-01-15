package practice

import (
	"fmt"
	"reflect"
	"testing"
)

type HelloInterface interface {
	Hello()
}

type HelloImpl struct {
	Name string
}

func (h HelloImpl) Hello() {
	fmt.Println("hello", h.Name)
}

func TestReflect(t *testing.T) {
	var hello HelloInterface = HelloImpl{Name: "Go reflect"}
	v := reflect.ValueOf(hello)
	fmt.Printf("v=%#v\n", v)
	// 如果直接在一个struct类型上调用Elem()方法，会报如下的错误
	// panic: reflect: call of reflect.Value.Elem on struct Value [recovered]
	//	panic: reflect: call of reflect.Value.Elem on struct Value
	//e := v.Elem()
	//fmt.Printf("e=%#v\n", e)
	vt := v.Type()
	fmt.Printf("vt=%#v\n", vt)

	helloType := reflect.TypeOf(hello)
	fmt.Printf("helloType=%#v\n", helloType)

	// 在调用Elem()方法时报错如下
	// panic: reflect: Elem of invalid type main.HelloImpl [recovered]
	//	panic: reflect: Elem of invalid type main.HelloImpl
	//helloElem := helloType.Elem()
	//fmt.Printf("helloElem=%#v\n", helloElem)

	helloTypePtr := reflect.PtrTo(helloType)
	fmt.Printf("helloTypePtr=%#v\n", helloTypePtr)
	// 获取这个类型的方法的个数
	fmt.Printf("helloTypePtr.NumMethod()=%#v\n", helloTypePtr.NumMethod())
	// Field方法应该是获取struct的字段的，因为interface你所struct类型所以也没有字段，调用时就会报错如下：
	// panic: reflect: Field of non-struct type *practice.HelloImpl [recovered]
	//	panic: reflect: Field of non-struct type *practice.HelloImpl
	//fmt.Printf("helloTypePtr.Field(0)=%#v\n", helloTypePtr.Field(0))
	fmt.Printf("helloTypePtr.Name()=%#v\n", helloTypePtr.Name())
	name, ok := helloTypePtr.MethodByName("Hello")
	if ok {
		fmt.Printf("helloTypePtr.MethodByName()=%#v\n", name)
	}
}
