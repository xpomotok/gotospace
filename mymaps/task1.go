package mymaps
package main

import "strings"

func wordFrequency(text string) map[string]int {
    words := strings.Fields()

dict := make(map[string]int)

for i, v := range words{
    dict[v] += 1
}
    return dict
}

// Пример: "hello world hello" -> map[string]int{"hello": 2, "world": 1}