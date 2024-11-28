package save

import (
	"encoding/json"
	"log"
	"main/internal/hangman-classic/pkg/structs"
	"os"
)

func StopAndSaveGame(data *structs.HangManData) {
	// create an array with the elements you want to save in save.txt
	mapReturn := map[string]interface{}{
		"ToFind":       data.ToFind,
		"Word":         data.Word,
		"Attempts":     data.Attempts,
		"AlreadyTried": data.AlreadyTried,
	}
	// creating a save file with json format in the save/save.txt file
	DataGameJson, _ := json.MarshalIndent(mapReturn, "", " ")
	err := os.WriteFile("save/save.txt", DataGameJson, 0777)
	if err != nil {
		log.Fatal(err)
	}
}

func StartWithFlag(Start string) structs.HangManData {
	var data structs.HangManData
	// read the save file
	JsonData, _ := os.ReadFile(Start)
	// extracting data from the file
	err := json.Unmarshal(JsonData, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
