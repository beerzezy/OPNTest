package assignment3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_calculateStats(t *testing.T) {

	tests := []struct {
		name     string
		input    []inputTeamStats
		expected []inputTeamStats
	}{
		{
			input: []inputTeamStats{
				{
					TeamName:    "D",
					Win:         6,
					Lose:        1,
					Draw:        1,
					GoalFor:     15,
					GoalAgainst: 11,
				},
				{
					TeamName:    "B",
					Win:         6,
					Lose:        1,
					Draw:        1,
					GoalFor:     24,
					GoalAgainst: 20,
				},
			},
			expected: []inputTeamStats{
				{
					TeamName:    "D",
					Win:         6,
					Lose:        1,
					Draw:        1,
					Score:       19,
					GoalFor:     15,
					GoalAgainst: 11,
					GoalDiff:    4,
				},
				{
					TeamName:    "B",
					Win:         6,
					Lose:        1,
					Draw:        1,
					Score:       19,
					GoalFor:     24,
					GoalAgainst: 20,
					GoalDiff:    4,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateStats(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func Test_calculateRanking(t *testing.T) {

	tests := []struct {
		name     string
		input    []inputTeamStats
		expected []resultTeamStat
	}{
		{
			input: []inputTeamStats{
				{
					TeamName:    "D",
					Win:         6,
					Lose:        1,
					Draw:        1,
					Score:       19,
					GoalFor:     15,
					GoalAgainst: 11,
					GoalDiff:    4,
				},
				{
					TeamName:    "B",
					Win:         6,
					Lose:        1,
					Draw:        1,
					Score:       19,
					GoalFor:     24,
					GoalAgainst: 20,
					GoalDiff:    4,
				},
			},
			expected: []resultTeamStat{
				{
					TeamName:    "B",
					Win:         6,
					Lose:        1,
					Draw:        1,
					Score:       19,
					GoalFor:     24,
					GoalAgainst: 20,
					GoalDiff:    4,
				},
				{
					TeamName:    "D",
					Win:         6,
					Lose:        1,
					Draw:        1,
					Score:       19,
					GoalFor:     15,
					GoalAgainst: 11,
					GoalDiff:    4,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateRanking(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
