package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHola(t *testing.T) {
	for _, tt := range []struct {
		in, want string
	}{
		{"", "¡Hola mundo!\n"},
		{"Juan", "¡Hola Juan!\n"},
		{"Pedro", "¡Hola Pedro!\n"},
		{"-", "¡Hola -!\n"},
	} {
		t.Run(fmt.Sprintf("hola(w, %q)", tt.in), func(t *testing.T) {
			var b strings.Builder
			hola(&b, tt.in)
			got := b.String()
			if got != tt.want {
				t.Errorf("expected %q got %q", tt.want, got)
			}
		})
	}
}

func TestHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	for _, tt := range []struct {
		in, want string
	}{
		{"", "¡Hola mundo!\n"},
		{"Juan", "¡Hola Juan!\n"},
		{"Pedro", "¡Hola Pedro!\n"},
		{"-", "¡Hola -!\n"},
	} {
		t.Run(fmt.Sprintf("hola(w, %q)", tt.in), func(t *testing.T) {
			url := ts.URL
			if tt.in != "" {
				url += "/?nombre=" + tt.in
			}
			res, err := http.Get(url)
			if err != nil {
				t.Fatal(err)
			}

			got, err := ioutil.ReadAll(res.Body)
			res.Body.Close()
			if err != nil {
				t.Fatal(err)
			}

			if string(got) != tt.want {
				t.Errorf("expected %q got %q", tt.want, string(got))
			}
		})
	}

}
