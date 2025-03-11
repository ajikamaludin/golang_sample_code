package main

import (
	"fmt"
	"strings"
)

func ReplaceWords(rootwords []string, sentence string) string {
	words := strings.Split(sentence, " ")
	for i, word := range words {
		for _, root := range rootwords {
			if strings.HasPrefix(word, root) {
				words[i] = root
				break
			}
		}
	}
	return strings.Join(words, " ")
}

func main() {
	rootwords := []string{"cat", "bat", "rat"}
	sentence := "the cattle was rattled by the battery"
	fmt.Println(ReplaceWords(rootwords, sentence))
	rootwords = []string{"dog", "car", "bike"}
	sentence = "the dog were barking near the cars and bikers"
	fmt.Println(ReplaceWords(rootwords, sentence))
}
