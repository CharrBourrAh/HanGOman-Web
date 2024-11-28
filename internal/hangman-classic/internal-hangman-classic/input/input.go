package input

import "fmt"

func Input() string {
	var userInput string
	// all the forbidden characters are written in the following string
	forbiddenCharacters := "1234567890°+&é\"(-è_çà)'=^$*ù!:;,?.§%¨£µ€<>~#{[|`\\]}@"
	fmt.Println("Enter an input :")
	data, err := fmt.Scanf("%s \n", &userInput)
	// handling errors
	if err != nil || data != 1 {
		fmt.Println("Error")
	}
	for i := 0; i < len(userInput); i++ {
		for j := 0; j < len(forbiddenCharacters); j++ {
			if userInput[i] == forbiddenCharacters[j] {
				fmt.Println("Error")
			}
		}
	}
	fmt.Println("You've chosen : " + userInput)
	return userInput
}
