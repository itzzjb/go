package hello

import (
	"testing"
)

// always get a t variable of type *testing.T as the first argument
func TestSayHello(t *testing.T) {

	// dynamic variable declaration
	want := "Hello, test!"
	got := SayHello("test")

	if got != want { 
		t.Errorf("wanted %s, but got %s", want, got)
	}
}