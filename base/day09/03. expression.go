package main

import "fmt"

type User11 struct {
	id   int
	name string
}

func (self *User11) Test() {
	fmt.Printf("%p, %v\n", self, self)
}

func (self User11) Test2() {
	fmt.Printf("%p, %v\n", self, self)
}

func (self *User11) TestPointer() {
	fmt.Printf("TestPointer: %p, %v\n", self, self)
}

func (self User11) TestValue() {
	fmt.Printf("TestValue: %p, %v\n", &self, self)
}

type Data struct{}

func (Data) TestValue() {}

func (*Data) TestPointer() {}

func f1() {
	u := User11{1, "Tom"}
	u.Test()

	mValue := u.Test
	mValue() // 隐式传递 receiver

	mExpression := (*User11).Test
	mExpression(&u) // 显式传递 receiver
}

func f2() {
	u := User11{1, "Tom"}
	mValue := u.Test2 // 立即复制 receiver，因为不是指针类型，不受后续修改影响。
	u.id, u.name = 2, "Jack"
	u.Test2()
	mValue()
}

func f3() {
	u := User11{1, "Tom"}
	fmt.Printf("User: %p, %v\n", &u, u)

	mv := User11.TestValue
	mv(u)

	mp := (*User11).TestPointer
	mp(&u)

	mp2 := (*User11).TestValue // *User 方法集包含 TestValue。签名变为 func TestValue(self *User)。实际依然是 receiver value copy。
	mp2(&u)
}

func f4() {
	var p *Data = nil
	p.TestPointer()

	(*Data)(nil).TestPointer() // method value
	(*Data).TestPointer(nil)   // method expression

	// p.TestValue()            // invalid memory address or nil pointer dereference

	// (Data)(nil).TestValue()  // cannot convert nil to type Data
	// Data.TestValue(nil)      // cannot use nil as type Data in function argument
}
func main() {
	// f1()
	// f2()
	// f3()
	f4()
}
