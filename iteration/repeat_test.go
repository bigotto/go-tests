package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	assertCorrectMessage(t, repeated, expected)
}

func TestRepeatString(t *testing.T) {
	repeated := RepeatString("a", 4)
	expected := "aaaa"

	assertCorrectMessage(t, repeated, expected)
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func BenchmarkRepeatString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatString("a", 5)
	}
}