package encapsulation

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

/*
	创建员工对象方式
*/
func TestCreateEmployeeObj(t *testing.T) {

	e := Employee{"0", "Job", 20}
	e1 := Employee{Name: "Mike", Age: 20}
	e2 := new(Employee) // 注意：这里返回的引用指针，相当于 e := &Employee{}
	e2.Id = "2"
	e2.Age = 20
	e2.Name = "Rose"

	t.Log(e)
	t.Log(e1)
	t.Log(e1.Name) // 注意： 通过实例的指针访问成员变量不需要使用 ->
	t.Log(e2)

	t.Logf("e is %T", e)
	t.Logf("e is %T", &e) // 注意：观察两条语句打印日志区别
	t.Logf("e2 is %T ", e2)

}

func TestStructOperations(t *testing.T) {
	e := Employee{"0", "Job", 20}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	t.Logf(e.String())
	t.Logf(e.String1())
}

/*
	在实例对应方法调用时，实例的成员进行值复制
*/
func (e Employee) String() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

/*
	通常情况下为了避免内存拷贝推荐使用下面定义方式
*/
func (e *Employee) String1() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}
