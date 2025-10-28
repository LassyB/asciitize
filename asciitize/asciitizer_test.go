package asciitize_test

import (
	"strings"
	"testing"

	"github.com/LassyB/asciitize/asciitize"
	"github.com/stretchr/testify/assert"
)

func TestAsciitizer_Asciitize(t *testing.T) {
	testCases := []struct {
		name         string
		filepath     string
		wantErr      error
		wantContents string
	}{}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := strings.Builder{}
			asciitizer := asciitize.NewAsciitizer(&output)
			err := asciitizer.Asciitize(tc.filepath)
			assert.Equal(t, err, tc.wantErr)
			assert.Equal(t, output.String(), tc.wantContents)
		})
	}
}
