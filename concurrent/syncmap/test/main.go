package main

import (
	"fmt"

	"github.com/xyersh/xuyacs/concurrent/syncmap"
)

func main() {
	sm := syncmap.NewSyncMap[string, int](4)

	sm.Set("Vasya", 44)
	sm.Set("Alex", 41)
	sm.Set("Natasha", 38)
	sm.Set("Beta", 0)
	sm.Set("Misha", 38)

	val, ok := sm.Get("Vasya")
	fmt.Printf("k: %s \tv: %d \t%t\n", "Vasya", val, ok)

	val, ok = sm.Get("Alex")
	fmt.Printf("k: %s \tv: %d \t%t\n", "Alex", val, ok)

	val, ok = sm.Get("Natasha")
	fmt.Printf("k: %s \tv: %d \t%t\n", "Natasha", val, ok)

	val, ok = sm.Get("Beta")
	fmt.Printf("k: %s \tv: %d \t%t\n", "Beta", val, ok)

	val, ok = sm.Get("Misha")
	fmt.Printf("k: %s \tv: %d \t%t\n", "Misha", val, ok)

	sm.Delete("Vasya")
	val, ok = sm.Get("Vasya")
	fmt.Printf("after deleting k: %s \tv: %d \t%t\n", "Vasya", val, ok)

	fmt.Println("Iteration over syncMap:")
	for key, val := range sm.All() {
		fmt.Printf("key: %v   val: %v\n", key, val)
	}
}
