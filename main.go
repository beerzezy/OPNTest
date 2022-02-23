package main

import (
	"fmt"
	"os"

	"github.com/beerzezy/OPNTest/assignment1"
	"github.com/beerzezy/OPNTest/assignment2"
	"github.com/beerzezy/OPNTest/assignment3"
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

type programs struct {
	No   int
	Name string
}

func main() {

	for {
		var choice int

		menu()

		fmt.Print("Select choice: ")

		n, err := fmt.Scanln(&choice)
		if n < 1 || err != nil {
			fmt.Println("invalid input")
			continue
		}

		switch choice {
		case (1):
			assignment1.CheckFirstNumeric()
		case (2):
			assignment2.CheckIPAddress()
		case (3):
			assignment3.FootballRankingCalculate()
		case (4):
			os.Exit(0)
		default:
			fmt.Printf("No choice...\n\n")
		}
	}

}

func menu() {

	programsList := []programs{
		{
			No:   1,
			Name: "Check First Numeric of String.",
		},
		{
			No:   2,
			Name: "IP Address Checker.",
		},
		{
			No:   3,
			Name: "Football Rankings Calculation.",
		},
		{
			No:   4,
			Name: "Exit Program.",
		},
	}

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("No.", "Programs")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, program := range programsList {
		tbl.AddRow(program.No, program.Name)
	}
	fmt.Println()
	tbl.Print()
}
