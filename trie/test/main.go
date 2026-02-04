package main

import (
	"fmt"

	"github.com/xyersh/xuyacs/trie"
)

func main() {
	trie := trie.NewTrie()
	trie.Insert("hello")
	trie.Insert("world")
	trie.Insert("hell")
	trie.Insert("he")
	trie.Insert("h")
	trie.Insert("helo")
	trie.Insert("helloworld")
	trie.Insert("helloworld1")
	trie.Insert("helloworld2")
	trie.Insert("helloworld3")
	trie.Insert("helloworld4")
	trie.Insert("helloworld5")
	trie.Insert("helloworld6")
	trie.Insert("helloworld7")
	trie.Insert("helloworld8")
	trie.Insert("helloworld9")
	trie.Insert("helloworld10")
	trie.Insert("helloworld11")
	trie.Insert("helloworld12")
	trie.Insert("helloworld13")
	trie.Insert("helloworld14")
	trie.Insert("helloworld15")
	trie.Insert("helloworld16")
	trie.Insert("helloworld17")
	trie.Insert("helloworld18")

	fmt.Printf("Search helloworld18 returns %t\n", trie.Search("helloworld18"))

	fmt.Printf("words stast with helloworld: %t\n", trie.StartsWith("helloworld"))

	trie.Delete("helloworld18")

	words := trie.GetAllWordsWithPrefix("hellow")
	fmt.Println(words)
}
