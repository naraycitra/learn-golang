package main

import "fmt"

const sizeOfAlphabet = 26

// Node represent node of tries
type Node struct {
	children [sizeOfAlphabet]*Node
	isEnd    bool
}

// Tries represent tries
type Tries struct {
	root *Node
}

// Init initialize root of tries
func Init() *Tries {
	result := &Tries{root: &Node{}}
	return result
}

// Insert is method to add a word into the tries
func (t *Tries) Insert(w string) {
	currentNode := t.root
	for i := 0; i < len(w); i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

// Search is method to check the word given is in our tries
func (t *Tries) Search(w string) bool {
	currentNode := t.root
	for i := 0; i < len(w); i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd == true {
		return true
	}
	return false
}

func main() {
	myTries := Init()
	toAdd := []string{
		"naray",
		"nurul",
		"nasa",
		"nusa",
		"namun",
		"nusantara",
	}
	for _, v := range toAdd {
		myTries.Insert(v)
	}
	fmt.Println(myTries.Search("nusantara"))
}
