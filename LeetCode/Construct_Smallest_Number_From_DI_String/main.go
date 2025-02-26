package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/trace"
	"time"
)

func smallestNumber(pattern string) string {
	n := len(pattern)
	result := make([]byte, 0, n+1)
	stack := make([]byte, 0)
	currentDigit := byte('1')

	for i := 0; i <= n; i++ {
		stack = append(stack, currentDigit)
		currentDigit++
		if i == n || pattern[i] == 'I' {
			for len(stack) > 0 {
				result = append(result, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
		}
	}

	return string(result)
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// Start tracing
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("Failed to create trace file: %v", err)
	}
	defer f.Close()

	if err := trace.Start(f); err != nil {
		log.Fatalf("Failed to start trace: %v", err)
	}
	defer trace.Stop()

	s1 := "IIIDIDDD"
	fmt.Println(smallestNumber(s1))

	s2 := "DDD"
	fmt.Println(smallestNumber(s2))

	time.Sleep(5 * time.Second)
}
