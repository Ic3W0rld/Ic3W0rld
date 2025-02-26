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

func removeOccurrences(s string, part string) string {
	for {
		// Find the index of the leftmost occurrence of `part`
		index := strings.Index(s, part)
		if index == -1 {
			// If no more occurrences are found, break the loop
			break
		}
		// Remove the leftmost occurrence of `part`
		s = s[:index] + s[index+len(part):]
	}
	return s
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

	s1 := "daabcbaabcbc"
	part1 := "abc"
	fmt.Println(removeOccurrences(s1, part1))

	s2 := "axxxxyyyyb"
	part2 := "xy"
	fmt.Println(removeOccurrences(s2, part2))

	time.Sleep(5 * time.Second)
}
