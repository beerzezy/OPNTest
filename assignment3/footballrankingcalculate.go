package assignment3

import (
	"fmt"
	"sort"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

const (
	winPoint  = 3
	drawPoint = 1
)

type inputTeamStats struct {
	TeamName    string
	Win         int
	Lose        int
	Draw        int
	Score       int
	GoalFor     int
	GoalAgainst int
	GoalDiff    int
}

type resultTeamStat struct {
	TeamName    string
	Win         int
	Lose        int
	Draw        int
	Score       int
	GoalFor     int
	GoalAgainst int
	GoalDiff    int
}

func FootballRankingCalculate() {

	var teamAmount int

	fmt.Println("\n------------ Program Football Rankings Calculation. ------------")
	fmt.Print("Enter Football Team Amount: ")
	fmt.Scanln(&teamAmount)

	teamStatList := inputFootballTeam(teamAmount)
	calculatedStats := calculateStats(teamStatList)
	calculatedRanking := calculateRanking(calculatedStats)

	inputRender(teamStatList)
	resultRender(calculatedRanking)

	fmt.Println("\n---------------------------------------------------------------")
}

func inputFootballTeam(teamAmount int) []inputTeamStats {
	teamStatList := []inputTeamStats{}
	countTeamAmount := 0

	for countTeamAmount < teamAmount {

		countTeamAmount++
		fmt.Printf("\n----------------- No. %d -----------------", countTeamAmount)

		teamStat := inputTeamStats{}

		fmt.Print("\nTeam: ")
		fmt.Scanln(&teamStat.TeamName)

		for {
			fmt.Print("Win: ")
			_, err := fmt.Scanln(&teamStat.Win)
			if err != nil {
				fmt.Println("Invalid Input Type")
			} else {
				break
			}
		}

		for {
			fmt.Print("Lose: ")
			_, err := fmt.Scanln(&teamStat.Lose)
			if err != nil {
				fmt.Println("Invalid Input Type")
			} else {
				break
			}
		}

		for {
			fmt.Print("Draw: ")
			_, err := fmt.Scanln(&teamStat.Draw)
			if err != nil {
				fmt.Println("Invalid Input Type")
			} else {
				break
			}
		}

		for {
			fmt.Print("Goal For: ")
			_, err := fmt.Scanln(&teamStat.GoalFor)
			if err != nil {
				fmt.Println("Invalid Input Type")
			} else {
				break
			}
		}

		for {
			fmt.Print("Goal Against: ")
			_, err := fmt.Scanln(&teamStat.GoalAgainst)
			if err != nil {
				fmt.Println("Invalid Input Type")
			} else {
				break
			}
		}

		teamStatList = append(teamStatList, teamStat)

	}

	return teamStatList
}

func calculateStats(teamStatList []inputTeamStats) []inputTeamStats {

	for i, teamStat := range teamStatList {
		teamStatList[i].GoalDiff = teamStat.GoalFor - teamStat.GoalAgainst
		teamStatList[i].Score = teamStat.Win*winPoint + teamStat.Draw
	}

	return teamStatList
}

func calculateRanking(teamStatList []inputTeamStats) []resultTeamStat {

	resultTeamStatList := []resultTeamStat{}

	for _, teamStat := range teamStatList {
		resultTeamStatList = append(resultTeamStatList, resultTeamStat{
			TeamName:    teamStat.TeamName,
			Win:         teamStat.Win,
			Lose:        teamStat.Lose,
			Draw:        teamStat.Draw,
			Score:       teamStat.Score,
			GoalDiff:    teamStat.GoalDiff,
			GoalFor:     teamStat.GoalFor,
			GoalAgainst: teamStat.GoalAgainst,
		})
	}

	sort.Slice(resultTeamStatList, func(i, j int) bool {
		if resultTeamStatList[i].Score != resultTeamStatList[j].Score {
			return resultTeamStatList[i].Score > resultTeamStatList[j].Score
		}
		if resultTeamStatList[i].GoalDiff != resultTeamStatList[j].GoalDiff {
			return resultTeamStatList[i].GoalDiff > resultTeamStatList[j].GoalDiff
		}
		return resultTeamStatList[i].GoalFor > resultTeamStatList[j].GoalFor
	})

	return resultTeamStatList
}

func inputRender(teamStatList []inputTeamStats) {

	fmt.Println("\nInput:")
	fmt.Println("=========================================")

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Team", "Win", "Lose", "Draw", "Goal For", "Goal Against")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, teamStats := range teamStatList {
		tbl.AddRow(teamStats.TeamName, teamStats.Win, teamStats.Lose, teamStats.Draw, teamStats.GoalFor, teamStats.GoalAgainst)
	}

	tbl.Print()
}

func resultRender(rankingList []resultTeamStat) {

	fmt.Println("\nResult:")
	fmt.Println("=========================================")

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Team", "Goal For", "Goal Diff", "Score")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, ranking := range rankingList {
		tbl.AddRow(ranking.TeamName, ranking.GoalFor, ranking.GoalDiff, ranking.Score)
	}

	tbl.Print()
}
