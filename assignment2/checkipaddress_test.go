package assignment2

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CheckIPAddress(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
		wantErr  bool
	}{
		{
			name:     "CheckIPAddress_1_Success",
			input:    "172.254.254.1",
			expected: true,
			wantErr:  false,
		},
		{
			name:     "CheckIPAddress_2_Success",
			input:    "0.1.1.256",
			expected: false,
			wantErr:  false,
		},
		{
			name:     "CheckIPAddress_3_Success",
			input:    "1.1.1.a",
			expected: false,
			wantErr:  false,
		},
		{
			name:     "CheckIPAddress_4_Success",
			input:    "1",
			expected: false,
			wantErr:  false,
		},
		{
			name:     "CheckIPAddress_5_Success",
			input:    "64.233.16.00",
			expected: false,
			wantErr:  false,
		},
		{
			name:     "CheckIPAddress_6_Success",
			input:    "7283728",
			expected: false,
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
			result, err := CheckIPAddress()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func Test_validateIPAddress(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "CheckIPAddress_1_Success",
			input:    "172.254.254.1",
			expected: true,
		},
		{
			name:     "CheckIPAddress_2_Success",
			input:    "0.1.1.256",
			expected: false,
		},
		{
			name:     "CheckIPAddress_3_Success",
			input:    "1.1.1.a",
			expected: false,
		},
		{
			name:     "CheckIPAddress_4_Success",
			input:    "1",
			expected: false,
		},
		{
			name:     "CheckIPAddress_5_Success",
			input:    "64.233.16.00",
			expected: false,
		},
		{
			name:     "CheckIPAddress_6_Success",
			input:    "7283728",
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validateIPAddress(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
