package main

import "fmt"

type DynamicArray struct {
	data     []int
	length   int
	capacity int
}

func NewDynamicArray(cap int) *DynamicArray {
	return &DynamicArray{
		data:     make([]int, cap),
		length:   0,
		capacity: cap,
	}
}

func (da *DynamicArray) Append(value int) {
	if da.length == da.capacity {
		da.resize()
	}

	da.data[da.length] = value
	da.length++
}

func (da *DynamicArray) resize() {
	newCapacity := da.capacity * 2
	newData := make([]int, newCapacity)

	for i := 0; i < da.length; i++ {
		newData[i] = da.data[i]
	}

	da.data = newData
	da.capacity = newCapacity
}

func (da *DynamicArray) Get(index int) int {

	if index < 0 || index >= da.length {
		panic("index out of bounds")
	}

	return da.data[index]
}

func main() {
	arr := NewDynamicArray(2)

	arr.Append(10)
	arr.Append(20)
	arr.Append(30)

	fmt.Println(arr.data)
	fmt.Println("Length", arr.length)
	fmt.Println("Capacity", arr.capacity)

}

