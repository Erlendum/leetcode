package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	data map[int]struct{}
}

func Constructor() RandomizedSet {
	data := make(map[int]struct{})
	var set RandomizedSet
	set.data = data
	return set
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.data[val]; ok {
		return false
	}
	this.data[val] = struct{}{}
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.data[val]; !ok {
		return false
	}
	delete(this.data, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	l := len(this.data)
	randNumber := rand.Intn(l)
	counter := 0

	for k := range this.data {
		if counter == randNumber {
			return k
		}
		counter++
	}
	return -1
}

func main() {
	set := Constructor()
	set.Insert(1)

	fmt.Println(set.GetRandom())
}
