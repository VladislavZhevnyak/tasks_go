package hw01

import "testing"

func TestSayHello(t *testing.T) {
	want := "привет, мир"
	if got := SayHello(); got != want {
		t.Errorf("SayHello() == %q, want %q", got, want)
	}
}

func BenchmarkSayHello(b *testing.B) {
	if testing.Short() {
		b.Skip("пропуск")
	}
	for i := 0; i < b.N; i++ {
		SayHello()
	}
}
