package leaderboard

import (
	"sort"
)

type LeaderBoardEntry struct {
	Nickname   string
	Attempts   int
	Difficulty string
}

var LeaderBoardData []LeaderBoardEntry

func AddPlayerToLeaderBoard(nickname string, attempts int, difficulty string) {
	entry := LeaderBoardEntry{
		Nickname:   nickname,
		Attempts:   attempts,
		Difficulty: difficulty,
	}
	LeaderBoardData = append(LeaderBoardData, entry)
}

func GetLeaderBoard() []LeaderBoardEntry {
	sortedLeaderBoard := make([]LeaderBoardEntry, len(LeaderBoardData))
	copy(sortedLeaderBoard, LeaderBoardData)

	sort.Slice(sortedLeaderBoard, func(i, j int) bool {
		return sortedLeaderBoard[i].Attempts > sortedLeaderBoard[j].Attempts
	})
	return sortedLeaderBoard
}

func getTopPlayers(limit int) []LeaderBoardEntry {
	leaderBoard := GetLeaderBoard()
	if limit > len(leaderBoard) {
		return leaderBoard
	}
	return leaderBoard[:limit]
}
