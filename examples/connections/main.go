package main

import "github.com/goqml/goqml"

func assert(condition bool, msg string) {
	if !condition {
		panic("断言失败: " + msg)
	}
}

//go:generate go run ../../cmd/goqml gen -f ./contact.go
func main() {
	{ // connect
		c1 := NewContact()
		defer c1.Delete()
		c2 := NewContact()
		defer c2.Delete()
		conn := goqml.Connect(
			c1, goqml.MakeSignal("nameChanged", "QString"),
			c2, goqml.MakeSlot("setName", "QString"))
		defer conn.Delete()
		assert(c1.Name != "John" && c2.Name != "John", "b11")
		c1.SetName("John")
		assert(c1.Name == "John" && c2.Name == "John", "b12")
	}
	{ // connect func
		c1 := NewContact()
		defer c1.Delete()
		c2 := NewContact()
		defer c2.Delete()
		conn := goqml.ConnectFunc(c1, goqml.MakeSignal("nameChanged", "QString"), func(name string) {
			c2.Name = name
		})
		defer conn.Delete()
		assert(c1.Name != "John" && c2.Name != "John", "b21")
		c1.SetName("John")
		assert(c1.Name == "John" && c2.Name == "John", "b22")
	}
	{ // disconnect
		c1 := NewContact()
		defer c1.Delete()
		c2 := NewContact()
		defer c2.Delete()
		conn := goqml.Connect(
			c1, goqml.MakeSignal("nameChanged", "QString"),
			c2, goqml.MakeSlot("setName", "QString"))
		assert(c1.Name != "John" && c2.Name != "John", "b31")
		c1.SetName("John")
		assert(c1.Name == "John" && c2.Name == "John", "b32")
		conn.Delete()
		c1.SetName("Doo")
		assert(c1.Name == "Doo" && c2.Name == "John", "b33")
	}
}
