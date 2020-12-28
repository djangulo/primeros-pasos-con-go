package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	for _, tt := range []struct {
		in                  string
		lines, words, bytes int
	}{
		{"Hello, is it me you're looking for?", 1, 7, 35},
		{"Hi\nthere\nhow\nare\nyou\ndoing today?", 6, 7, 28},
		{"", 0, 0, 0},
		{"\n", 1, 0, 0},
	} {
		t.Run(tt.in, func(t *testing.T) {
			path := createTestFile(t, tt.in)
			cmd := exec.Command("go", "run", "./main.go", path)
			out, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatal(err)
			}
			got := strings.Split(string(out), " ")
			lines, _ := strconv.Atoi(got[0])
			words, _ := strconv.Atoi(got[1])
			bytes, _ := strconv.Atoi(got[2])

			if lines != tt.lines {
				t.Errorf("lines: expected %d got %d", tt.lines, lines)
			}
			if words != tt.words {
				t.Errorf("words: expected %d got %d", tt.words, words)
			}
			if bytes != tt.bytes {
				t.Errorf("bytes: expected %d got %d", tt.bytes, bytes)
			}
		})
	}
}

func createTestFile(t *testing.T, contents string) (path string) {
	testdir := t.TempDir()
	path = filepath.Join(testdir, "testfile.txt")
	fh, err := os.Create(path)
	if err != nil {
		t.Fatalf("error creating test file %q: %v", path, err)
	}
	defer fh.Close()
	fh.WriteString(contents)

	return
}
