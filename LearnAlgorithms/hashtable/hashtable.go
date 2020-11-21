package main

import "fmt"

// ArraySize is the size that for store hash table
const ArraySize = 10

// HashTable represent structure of hash table
type HashTable struct {
	Array [ArraySize]*bucket
}

type bucket struct {
	head   *bucketNode
	length int
}

type bucketNode struct {
	key  string
	next *bucketNode
}

func hash(key string) int {
	var sum int
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
		b.length++
	} else {
		fmt.Println(k, "already exists!")
	}
}

func (b *bucket) search(k string) bool {
	currentNode := b.head
	for i := 0; i < b.length; i++ {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (b *bucket) delete(k string) {
	// handle if node to delete is in the head
	if b.head.key == k {
		b.head = b.head.next
		b.length--
		return
	}

	previousNode := b.head
	for i := 0; i < b.length; i++ {
		if previousNode.key == k {
			previousNode.next = previousNode.next.next
			b.length--
		}
		previousNode = previousNode.next
	}
}

// Init will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.Array {
		result.Array[i] = &bucket{}
	}
	return result
}

// Insert method to add key into hash table
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.Array[index].insert(key)
}

// Search method to search the given key and return true id key is exists
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.Array[index].search(key)
}

// Delete method to delete the given key from hashtable
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.Array[index].delete(key)
}

func main() {
	myHashTable := Init()
	list := []string{
		"naray",
		"nurul",
		"abdullah",
		"fatimah",
		"aisyah",
	}
	for _, v := range list {
		myHashTable.Insert(v)
	}
	myHashTable.Delete("aisyah")
	fmt.Println(myHashTable.Search("abdullah"))
	fmt.Println(myHashTable.Search("aisyah"))
}
