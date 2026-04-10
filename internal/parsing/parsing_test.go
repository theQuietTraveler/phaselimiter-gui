package parsing

import "testing"

func TestParseDroppedFilePath(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		goos    string
		want    string
		wantErr bool
	}{
		{name: "decodes escaped spaces", input: "file:///tmp/my%20song.wav", goos: "linux", want: "/tmp/my song.wav"},
		{name: "decodes unicode characters", input: "file:///tmp/h%C3%B8st.wav", goos: "linux", want: "/tmp/høst.wav"},
		{name: "windows drive format", input: "file:///C:/audio/test.wav", goos: "windows", want: "C:/audio/test.wav"},
		{name: "unsupported scheme", input: "https://example.com/audio.wav", goos: "linux", wantErr: true},
		{name: "invalid uri", input: "file://%zz", goos: "linux", wantErr: true},
		{name: "empty", input: "", goos: "linux", wantErr: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ParseDroppedFilePath(tc.input, tc.goos)
			if (err != nil) != tc.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tc.wantErr)
			}
			if err == nil && got != tc.want {
				t.Fatalf("path = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestExtractProgression(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		want     float64
		wantOkay bool
	}{
		{name: "integer progression", line: "progression: 1", want: 1, wantOkay: true},
		{name: "decimal progression", line: "progression: 0.625", want: 0.625, wantOkay: true},
		{name: "signed progression", line: "progression: -0.2", want: -0.2, wantOkay: true},
		{name: "no progression", line: "processing file", wantOkay: false},
		{name: "invalid number", line: "progression: not-a-number", wantOkay: false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := ExtractProgression(tc.line)
			if ok != tc.wantOkay {
				t.Fatalf("ok = %v, want %v", ok, tc.wantOkay)
			}
			if ok && got != tc.want {
				t.Fatalf("progression = %v, want %v", got, tc.want)
			}
		})
	}
}
