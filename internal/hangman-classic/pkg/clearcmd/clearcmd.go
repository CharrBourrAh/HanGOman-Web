package clearcmd

import (
	"fmt"
	"os"
	"os/exec"
)

func ClearCMD() {
	// clears the command prompt for better readability
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	// handling errors
	if err != nil {
		fmt.Println("This type of terminal is not supported by this game. Please use Windows' newer or classic Terminal app")
		return
	}
}
