package main

import "fmt"

type MyHashMap struct {
	data []int
}

func Constructor() MyHashMap {
	// Inisialisasi slice dengan panjang 1_000_001 dan isi -1
	size := 1000001
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = -1
	}
	return MyHashMap{data: data}
}

func (this *MyHashMap) Put(key int, value int) {
	this.data[key] = value
}

func (this *MyHashMap) Get(key int) int {
	return this.data[key]
}

func (this *MyHashMap) Remove(key int) {
	this.data[key] = -1
}

// Contoh penggunaan
func main() {
	obj := Constructor()
	obj.Put(1, 10)
	obj.Put(2, 20)
	fmt.Println(obj.Get(1)) // Output: 10
	fmt.Println(obj.Get(3)) // Output: -1
	obj.Put(2, 30)
	fmt.Println(obj.Get(2)) // Output: 30
	obj.Remove(2)
	fmt.Println(obj.Get(2)) // Output: -1
}
