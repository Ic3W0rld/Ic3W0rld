package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/trace"
	"strings"
	"time"
)

func getHappyString(n int, k int) string {
	chars := []rune{'a', 'b', 'c'}
	count := 0
	var result strings.Builder
	backtrack(n, &count, k, &result, chars, "")
	if count >= k {
		return result.String()
	}
	return ""
}

func backtrack(n int, count *int, k int, result *strings.Builder, chars []rune, current string) {
	if len(current) == n {
		*count++
		if *count == k {
			result.WriteString(current)
		}
		return
	}
	for _, ch := range chars {
		if len(current) == 0 || current[len(current)-1] != byte(ch) {
			backtrack(n, count, k, result, chars, current+string(ch))
			if *count >= k {
				return
			}
		}
	}
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

	n1 := 1
	k1 := 3
	fmt.Println(getHappyString(n1, k1))

	n2 := 1
	k2 := 4
	fmt.Println(getHappyString(n2, k2))

	n3 := 3
	k3 := 9
	fmt.Println(getHappyString(n3, k3))

	time.Sleep(5 * time.Second)
}
