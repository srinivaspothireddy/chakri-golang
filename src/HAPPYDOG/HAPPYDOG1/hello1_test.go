package hello1

import "testing"

func TestHello(t *testing.T) {
	want := "Hello1, world."
	if got := Hello1(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}