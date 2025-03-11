package main

import (
	"fmt"
	"strings"
)

// Fungsi untuk mencari akar kata terpendek yang cocok dengan kata
func findRoot(word string, rootwords []string) string {
	shortestRoot := word // Defaultnya adalah kata itu sendiri jika tidak ada yang cocok
	for _, root := range rootwords {
		if strings.HasPrefix(word, root) { // Cek apakah kata diawali dengan rootword
			if len(root) < len(shortestRoot) { // Pilih yang paling pendek
				shortestRoot = root
			}
		}
	}
	return shortestRoot
}

// Fungsi untuk mengganti kata dalam kalimat
func replaceWords(rootwords []string, sentence string) string {
	words := strings.Fields(sentence) // Pisahkan kalimat menjadi kata-kata
	for i, word := range words {
		words[i] = findRoot(word, rootwords) // Ganti kata dengan root yang sesuai
	}
	return strings.Join(words, " ") // Gabungkan kembali menjadi kalimat
}

func main() {
	rootwords := []string{"cat", "bat", "rat"}
	sentence := "the cattle was rattled by the battery"
	fmt.Println(replaceWords(rootwords, sentence))
	rootwords = []string{"dog", "car", "bike"}
	sentence = "the dog were barking near the cars and bikers"
	fmt.Println(replaceWords(rootwords, sentence))
}
