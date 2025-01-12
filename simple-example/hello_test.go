package hello

import (
	"testing"
)

// always gets a t variable of type *testing.T as the first argument
func TestSayHello(t *testing.T) {

	// dynamic variable declaration
	want := "Hello test1, test2!"
	got := SayHello([]string{"test1", "test2"})

	if got != want { 
		t.Errorf("wanted %s, but got %s", want, got)
	}
}