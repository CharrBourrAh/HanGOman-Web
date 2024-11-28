package leaderboard

import (
	"sort"
)

type Entry struct {
	Nickname   string
	Attempts   int
	Difficulty string
}

var Data []Entry

func AddPlayerToLeaderBoard(nickname string, attempts int, difficulty string) {
	entry := Entry{
		Nickname: nickname,
		Attempts: attempts,
	}
	switch difficulty {
	case "./internal/hangman-classic/data/words.txt":
		entry.Difficulty = "Facile"
	case "./internal/hangman-classic/data/words2.txt":
		entry.Difficulty = "Moyen"
	case "./internal/hangman-classic/data/words3.txt":
		entry.Difficulty = "Force à toi"

	}
	Data = append(Data, entry)
}

func GetLeaderBoard() []Entry {
	sortedLeaderBoard := make([]Entry, len(Data))
	copy(sortedLeaderBoard, Data)

	sort.Slice(sortedLeaderBoard, func(i, j int) bool {
		return sortedLeaderBoard[i].Attempts > sortedLeaderBoard[j].Attempts
	})
	return sortedLeaderBoard
}

func getTopPlayers(limit int) []Entry {
	leaderBoard := GetLeaderBoard()
	if limit > len(leaderBoard) {
		return leaderBoard
	}
	return leaderBoard[:limit]
}
