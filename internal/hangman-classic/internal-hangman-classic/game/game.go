package game

import (
	"fmt"
	"log"
	"main/internal/hangman-classic/pkg/clearcmd"
	"main/internal/hangman-classic/pkg/structs"
	"math/rand"
	//"net/http"
	"os"
	"slices"
	"strings"
	"unicode"
)

func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func isInList(list []string, s string) bool {
	for _, r := range list {
		if r == s {
			return true
		}
	}
	return false
}

func ReadFile(nameFile string) [][]string {
	content, err := os.ReadFile(nameFile)
	if err != nil {
		log.Fatal(err)
	}
	// Convert []byte to string
	wordFile := []rune(string(content))
	var wordTab [][]string
	counter := 0
	wordTab = append(wordTab, []string{})
	// filling wordTab
	for i := 0; i < len(wordFile); i++ {
		if wordFile[i] != 10 {
			// if the character is not a new line character
			wordTab[counter] = append(wordTab[counter], string(wordFile[i]))
		} else {
			counter++
			wordTab = append(wordTab, []string{})
		}
	}
	return wordTab
}

func randomWord(list [][]string, data *structs.HangManData) {
	randomWordPos := rand.Intn(len(list))
	for i := 0; i < len(list[randomWordPos])-1; i++ {
		data.Word += "_"
		data.ToFind += list[randomWordPos][i]
	}
}

func Init(data *structs.HangManData) {
	// variables initialisation
	data.Attempts = 10
	data.AlreadyTried = []string{}
	randomWord(ReadFile(data.WordFile), data)
	randLetter(data)
	//data.HangmanPositions = ReadFile("data/hangman.txt")
}

func randLetter(data *structs.HangManData) {
	n := len(data.ToFind)/2 - 1
	wordCopy := strings.Split(data.Word, "")
	for i := 0; i < n; i++ {
		randomIndex := rand.Intn(len(data.ToFind))
		wordCopy[randomIndex] = strings.Split(data.ToFind, "")[randomIndex]
	}
	data.Word = strings.Join(wordCopy, "")
}

func InsertChar(data *structs.HangManData) {
	copyWord := strings.Split(data.Word, "")
	userInput := ""
	fmt.Println(data.Word)
	fmt.Print("Already guessed letters / words : ")
	fmt.Println(data.AlreadyTried)
	fmt.Println()
	userInput = data.Input
	for i := 0; i < len(data.ToFind); i++ {
		if len(userInput) == 1 {
			for j := 0; j < len(data.Word); j++ {
				// Adding the given letter to the correct position(s)
				if userInput == string(data.ToFind[j]) {
					copyWord[j] = userInput
				}
			}
		}
		if len(userInput) >= 2 {
			if IsLetter(userInput) == true {
				if isInList(data.AlreadyTried, userInput) == true {
					clearcmd.ClearCMD()
					fmt.Println("You've already used this word before")
					break
				} else {
					if userInput != data.ToFind {
						// If the given word is incorrect
						data.Attempts -= 2
						fmt.Println("This word is incorrect. You have", data.Attempts, "attempts remaining")
						if isInList(data.AlreadyTried, userInput) == false {
							data.AlreadyTried = append(data.AlreadyTried, userInput)
						}
						return
					} else {
						// If the given word is correct
						data.Word = data.ToFind
						return
					}
				}
			}
		}
	}
	if IsLetter(userInput) == true && len(userInput) == 1 {
		// if the previous letter list is similar with the new list, one attempt is removed
		if strings.Join(copyWord, "") == data.Word && slices.Contains(data.AlreadyTried, userInput) == false {
			data.Attempts -= 1
			fmt.Println("Not present in the word,", data.Attempts, "attempts remaining")
			// if the letter has already been tried
		} else if slices.Contains(data.AlreadyTried, userInput) == true {
			fmt.Println("You've already used this letter before")
			// if the guessed letter is correct
			return
		} else {
			data.Word = strings.Join(copyWord, "")
			return
		}
		if isInList(data.AlreadyTried, userInput) == false {
			data.AlreadyTried = append(data.AlreadyTried, userInput)
			return
		}
	}

}

func StatusGame(data *structs.HangManData) string {
	if data.Attempts == 0 {
		// loose scenario
		fmt.Println(data.Word + "\nYou loose :( \nYou've needed to find " + data.ToFind)
		return "loose"
	} else if data.Word == data.ToFind {
		// win scenario
		fmt.Println("You've won horray :D\nYou've successfully found " + data.ToFind)
		return "win"
	} else {
		return "ingame"
	}
}
