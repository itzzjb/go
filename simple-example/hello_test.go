package hello

import (
	"testing"
)

// always gets a t variable of type *testing.T as the first argument
func TestSayHello(t *testing.T) {

	// created a slice of struct to test multiple scenarios - one set for each test case
	// struct is like a new data type that can hold multiple values
	// here we haven't given it a name, so it's an anonymous struct
	subtests := []struct { 
		items []string
		result string
	} {
		// if we don't provide items, it should return "Hello world!"
		{
			result : "Hello world!", 
		}, 
		// if we provide 1 item, it should return "Hello test!"
		{	
			items : []string{"test"}, 
			result: "Hello test!", 
		}, 
		// if we provide 2 items, it should return "Hello test1, test2!"
		{	
			items : []string{"test1", "test2"}, 
			result: "Hello test1, test2!", 
		}, 
	}

	// _ is a blank identifier - it's used to ignore the index of the slice 
	// _, st is a tuple assignment - it assigns the first value to _ and the second value to st
	for _, st := range subtests {
		if s := SayHello(st.items); s != st.result {
			t.Errorf("wanted %s [%v], but got %s", st.result, st.items, s)
		}
	}

	// dynamic variable declaration
	// want := "Hello test1, test2!"
	// got := SayHello([]string{"test1", "test2"})

	// if got != want { 
	// 	t.Errorf("wanted %s, but got %s", want, got)
	// }
}