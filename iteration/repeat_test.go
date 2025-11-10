package iteration

import (
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 10)
	want := strings.Repeat("a", 10)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 0)
	}
}
