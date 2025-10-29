package asciitize_test

import (
	"errors"
	"os"
	"strings"
	"syscall"
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
	}{
		{
			name:     "error returned when failure to open",
			filepath: "testdata/nonexistent.png",
			wantErr: &os.PathError{
				Op:   "open",
				Path: "testdata/nonexistent.png",
				Err:  syscall.Errno(2),
			},
		},
		{
			name:     "error returned when unsupported file type",
			filepath: "testdata/unsupported.txt",
			wantErr:  errors.New("image: unknown format"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := strings.Builder{}
			asciitizer := asciitize.NewAsciitizer(
				&output,
			)
			err := asciitizer.Asciitize(tc.filepath)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantContents, output.String())
		})
	}
}
