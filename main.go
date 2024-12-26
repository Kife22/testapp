package main

import (
	"fmt"
	"sync"
)

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func worker(id int, texts <-chan string, wg *sync.WaitGroup, s string) {
	defer wg.Done()
	for text := range texts {
		reversed := reverseString(text)
		fmt.Printf("line %d, thread %d: %s\n", id, id%3+1, reversed)

	}
}

func main() {
	texts := []string{"Hello", "qwerty", "Golang", "platypus", "тест", "level", "generics"}
	textsChannel := make(chan string)

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i+1, textsChannel, &wg, "")
	}

	for _, text := range texts {
		textsChannel <- text
	}

	close(textsChannel)
	wg.Wait()

	fmt.Println("end")
}
