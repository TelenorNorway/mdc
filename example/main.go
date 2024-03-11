package main

import (
	"fmt"
	"github.com/telenornorway/mdc"
	"github.com/telenornorway/mdc/revertable"
	"sync"
)

func printCurrentMdc(message string) {
	copy := mdc.Copy()
	if len(copy) == 0 {
		fmt.Printf("%s: empty\n", message)
		return
	}
	first := true
	str := ""
	for k, v := range copy {
		if first {
			first = false
		} else {
			str += ", "
		}
		str += k + "=" + v
	}
	fmt.Printf("%s: %s\n", message, str)
}

func main() {
	wg := &sync.WaitGroup{}

	mdc.Put("hello", "world")
	printCurrentMdc("main")

	wg.Add(2)

	go func() {
		defer wg.Done()
		printCurrentMdc("go1")
		mdc.Put("hello", "world2")
		printCurrentMdc("go1")
	}()

	go func() {
		defer wg.Done()
		printCurrentMdc("go2")
		mdc.Put("hello", "world3")
		printCurrentMdc("go2, should have hello")

		revert := revertable.New().
			Put("hello", "no-one").
			Apply()

		printCurrentMdc("go2, no-one expected")

		revert.Revert()

		printCurrentMdc("go2, world3 expected")
	}()

	wg.Wait()
}
