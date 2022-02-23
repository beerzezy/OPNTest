package assignment1

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckFirstNumeric(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{
		{
			name:     "CheckFirstNumeric_Success",
			input:    "Var-----___1_int1",
			expected: "1",
			wantErr:  false,
		},
		{
			name:     "CheckFirstNumeric_Success",
			input:    "Q2q-q",
			expected: "2",
			wantErr:  false,
		},
		{
			name:     "CheckFirstNumeric_Success",
			input:    "eef3243**s",
			expected: "3",
			wantErr:  false,
		},
		{
			name:     "CheckFirstNumeric_Not_Found",
			input:    "abcdefg",
			expected: "Numeric not found.",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		content := []byte(tt.input)
		tmpfile, _ := ioutil.TempFile("", "")
		defer os.Remove(tmpfile.Name()) // clean up
		tmpfile.Write(content)
		tmpfile.Seek(0, 0)
		oldStdin := os.Stdin
		defer func() { os.Stdin = oldStdin }() // Restore original Stdin
		os.Stdin = tmpfile

		t.Run(tt.name, func(t *testing.T) {
			result, err := CheckFirstNumeric()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func Test_findFirstNumeric(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "CheckFirstNumeric_Success",
			input:    "Var-----___1_int1",
			expected: "1",
		},
		{
			name:     "CheckFirstNumeric_Success",
			input:    "Q2q-q",
			expected: "2",
		},
		{
			name:     "CheckFirstNumeric_Success",
			input:    "eef3243**s",
			expected: "3",
		},
		{
			name:     "CheckFirstNumeric_Not_Found",
			input:    "abcdefg",
			expected: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findFirstNumeric(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
