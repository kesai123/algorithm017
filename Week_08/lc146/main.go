package main

import "fmt"

func main(){
	c := Constructor(2)
	c.Put(1,1)
	c.Put(2,2)
	fmt.Println("get1:", c.Get(1))
	c.Put(3,3)
	fmt.Println("get2:", c.Get(2))
	c.Put(4,4)
	fmt.Println("get1:", c.Get(1))
	fmt.Println("get3:", c.Get(3))
	fmt.Println("get4:", c.Get(4))
}