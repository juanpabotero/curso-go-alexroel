package greeting

import "testing"

func TestHelloName(t *testing.T) {
	name := "Juan"
	want := "Hello, " + name
	got, _ := Hello(name)
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	want := "empty name"
	got, err := Hello("")
	if got != "" || err.Error() != want {
		t.Errorf("got %q want %q", got, want)
	}
}
