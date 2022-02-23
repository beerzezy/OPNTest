package assignment3

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

const (
	winPoint  = 3
	drawPoint = 1
)

type inputTeamStats struct {
	TeamName    string `field:"Team"`
	Win         int    `field:"Win"`
	Lose        int    `field:"Lose"`
	Draw        int    `field:"Draw"`
	Score       int    `field:"Score"`
	GoalFor     int    `field:"Goal For"`
	GoalAgainst int    `field:"Goal Against"`
	GoalDiff    int    `field:"Goal Diff"`
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
		fmt.Printf("\n----------------- No. %d -----------------\n", countTeamAmount)

		teamStat := inputTeamStats{}

		teamStat = inputStats("TeamName", teamStat)
		teamStat = inputStats("Win", teamStat)
		teamStat = inputStats("Lose", teamStat)
		teamStat = inputStats("Draw", teamStat)
		teamStat = inputStats("GoalFor", teamStat)
		teamStat = inputStats("GoalAgainst", teamStat)

		teamStatList = append(teamStatList, teamStat)

	}

	return teamStatList
}

func inputStats(fieldName string, teamStat inputTeamStats) inputTeamStats {
	scanner := bufio.NewScanner(os.Stdin)
	stats := teamStat

	t := reflect.ValueOf(stats).Type()
	n, _ := t.FieldByName(fieldName)
	tagName := n.Tag.Get("field")

	v := reflect.ValueOf(&stats).Elem()
	for {
		fmt.Printf("%s: ", tagName)
		if scanner.Scan() {
			input := scanner.Text()
			itemType := v.FieldByName(fieldName).Type()

			if itemType.Kind() == reflect.String {
				v.FieldByName(fieldName).SetString(input)
				break
			}
			if itemType.Kind() == reflect.Int {
				inputInt, err := strconv.Atoi(input)
				if err != nil {
					fmt.Println("Invalid Input Type")
				} else {
					v.FieldByName(fieldName).SetInt(int64(inputInt))
					break
				}
			}

		}
	}

	return stats
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
