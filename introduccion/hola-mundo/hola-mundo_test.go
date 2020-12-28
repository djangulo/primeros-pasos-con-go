package main

import (
	"os/exec"
	"testing"
)

func TestMain(t *testing.T) {
	cmd := exec.Command("go", "run", "./hola-mundo.go")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}
	got := string(out)
	want := "¡Hola mundo!\n"

	if got != want {
		t.Errorf("expected %q got %q", want, got)
	}
}
