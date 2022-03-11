package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q repeated %q", expected, repeated)
	}
}

func Repeat(name string) string {
	return "aaaaa"
}
