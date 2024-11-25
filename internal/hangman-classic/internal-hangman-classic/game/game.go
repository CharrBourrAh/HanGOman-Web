package game

import (
	"fmt"
	"log"
	"main/internal/hangman-classic/internal-hangman-classic/input"
	"main/internal/hangman-classic/pkg/clearcmd"
	"main/internal/hangman-classic/pkg/structs"
	"math/rand"
	"net/http"
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

/*
func showHangman(hangman [][]string, attempts int) {
	for i := 8 * attempts; i < 8+8*attempts-1; i++ {
		for j := 0; j < len(hangman[i]); j++ {
			fmt.Print(hangman[i][j])
		}
		fmt.Print("\n")
	}
}
*/

func Init(data *structs.HangManData, r *http.Request) {
	// variables initialisation
	data.Attempts = 10
	data.AlreadyTried = []string{}
	randomWord(ReadFile(data.WordFile), data)
	randLetter(data)
	//data.HangmanPositions = ReadFile("data/hangman.txt")
	Game(data, r)
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

/*
func Resume(fileData structs.HangManData) {
	//fileData.HangmanPositions = ReadFile("data/hangman.txt")
	Game(&fileData)
}
*/

func Game(data *structs.HangManData, r *http.Request) {
	copyWord := strings.Split(data.Word, "")
	for data.Word != data.ToFind && data.Attempts > 0 {
		/*showHangman(data.HangmanPositions, 10-data.Attempts)
		fmt.Println(data.Word)
		fmt.Print("Already guessed letters / words : ")
		fmt.Println(data.AlreadyTried)
		fmt.Println()*/
		userInput := strings.ToLower(input.Input(r))
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
							clearcmd.ClearCMD()
							fmt.Println("This word is incorrect. You have", data.Attempts, "attempts remaining")
							if isInList(data.AlreadyTried, userInput) == false {
								data.AlreadyTried = append(data.AlreadyTried, userInput)
							}
							break
						} else {
							// If the given word is correct
							data.Word = data.ToFind
							break
						}
					}
				}
			}
			//clearcmd.ClearCMD()
		}
		if IsLetter(userInput) == true && len(userInput) == 1 {
			if strings.Join(copyWord, "") == data.Word && slices.Contains(data.AlreadyTried, userInput) == false {
				// if the previous letter list is similar with the new list, one attempt is remove
				data.Attempts -= 1
				fmt.Println("Not present in the word,", data.Attempts, "attempts remaining")
			} else if slices.Contains(data.AlreadyTried, userInput) == true {
				// if the letter has already been tried
				clearcmd.ClearCMD()
				fmt.Println("You've already used this letter before")
			} else {
				// if the guessed letter is correct
				data.Word = strings.Join(copyWord, "")
			}
			if isInList(data.AlreadyTried, userInput) == false {
				data.AlreadyTried = append(data.AlreadyTried, userInput)
			}
		}
	}
	//clearcmd.ClearCMD()
	if data.Attempts == 0 {
		// loose scenario
		fmt.Println(data.Word + "\nYou loose :( \nYou've needed to find " + data.ToFind)
	} else {
		// win scenario
		fmt.Println("You've won horray :D\nYou've successfully found " + data.ToFind)
	}
	//Menu()
}

/*
	func Menu() {
	for {
		mainMenuAscii := "  __  __       _                                    \n |  \\/  |     (_)                                   \n | \\  / | __ _ _ _ __    _ __ ___   ___ _ __  _   _ \n | |\\/| |/ _` | | '_ \\  | '_ ` _ \\ / _ \\ '_ \\| | | |\n | |  | | (_| | | | | | | | | | | |  __/ | | | |_| |\n |_|  |_|\\__,_|_|_| |_| |_| |_| |_|\\___|_| |_|\\__,_|\n                                                    \n                                                    "
		fmt.Print(mainMenuAscii)
		fmt.Println("\n" + "\033[32m" + "s" + "\033[0m" + " : launch a new game")
		fmt.Println("\x1b[33m" + "o" + "\033[0m" + " : opens the game's settings (change the words files)")
		fmt.Println("\033[31m" + "q" + "\033[0m" + " : exit the game")
		choice := input.Input()
		if choice == "s" {
			// launches the game
			clearcmd.ClearCMD()
			Init("")
		}
		if choice == "o" {
			// open the settings
			for {
				clearcmd.ClearCMD()
				fmt.Println("\n   _____      _   _   _                 \n  / ____|    | | | | (_)                \n | (___   ___| |_| |_ _ _ __   __ _ ___ \n  \\___ \\ / _ \\ __| __| | '_ \\ / _` / __|\n  ____) |  __/ |_| |_| | | | | (_| \\__ \\\n |_____/ \\___|\\__|\\__|_|_| |_|\\__, |___/\n                               __/ |    \n                              |___/     \n")
				fmt.Println("Choose which word file you want to use to play\n")
				fmt.Println("wi : Launch a new game while using a word in words.txt")
				fmt.Println("wii : Launch a new game while using a word in words2.txt")
				fmt.Println("wiii : Launch a new game while using a word in words3.txt\n")
				fmt.Println("e : Go back to the main menu")
				choice = input.Input()
				if choice == "e" {
					Menu()
				} else if choice == "wi" {
					// launches the game with words.txt
					clearcmd.ClearCMD()
					Init("data/words.txt")
				} else if choice == "wii" {
					// launches the game with words2.txt
					clearcmd.ClearCMD()
					Init("data/words2.txt")
				} else if choice == "wiii" {
					// launches the game with words3.txt
					clearcmd.ClearCMD()
					Init("data/words3.txt")
				}
				choice = input.Input()
			}
		}
		if choice == "q" {
			// Exit the program
			clearcmd.ClearCMD()
			os.Exit(3)
		}
	}
}
*/
