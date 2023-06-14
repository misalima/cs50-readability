package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"math"
)

func main() {
	var text string
	var L, S float64

	fmt.Println("Enter the text to determine the grade:")
	in := bufio.NewReader(os.Stdin)
	text, err := in.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading string...")
	}
	
	lettersNum := countLetters(text)
	wordsNum := countWords(text)
	sentencesNum := countSentences(text)
	
	fmt.Println("\n", lettersNum, "letters,", wordsNum, "words and", sentencesNum, "sentences.")
	
	L = float64(lettersNum) / float64(wordsNum) * 100
	S = float64(sentencesNum) / float64(wordsNum) * 100

	index := 0.0588 * L - 0.296 * S - 15.8

	fmt.Println("Grade: ", int(math.Round(index)))

}

func countLetters(text string) int {
	var lettersNum int
	for _, letter := range text {
		if (letter >= 65 && letter <= 90) || (letter >= 97 && letter <= 122) {
			lettersNum++
		} 
	}
	return lettersNum
}

func countWords(text string) int {
	var wordsNum int
	for _, letter := range text {
		if letter == 32 || letter == 10 {
			wordsNum++
		}
	}
	return wordsNum
}

func countSentences(text string) int {
	var sentenceNum int
	for _, letter := range text {
		
		if	letter == 33 || letter == 46 || 
			letter == 63 {
			sentenceNum++
		}
	}
	return sentenceNum
}