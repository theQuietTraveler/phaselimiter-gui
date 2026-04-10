package parsing

import "testing"

func BenchmarkParseDroppedFilePath(b *testing.B) {
	input := "file:///tmp/Artist%20Name%20-%20Track%20%231.wav"
	for i := 0; i < b.N; i++ {
		_, _ = ParseDroppedFilePath(input, "linux")
	}
}

func BenchmarkExtractProgression(b *testing.B) {
	line := "progression: 0.8342"
	for i := 0; i < b.N; i++ {
		_, _ = ExtractProgression(line)
	}
}
