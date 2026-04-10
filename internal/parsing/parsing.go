package parsing

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

var (
	windowsDrivePathPattern = regexp.MustCompile(`^/([a-zA-Z]:/)`)
	progressionPattern      = regexp.MustCompile(`progression: ([-+]?[0-9]*\.?[0-9]+)`)
)

func ParseDroppedFilePath(line string, goos string) (string, error) {
	trimmed := strings.TrimSpace(line)
	if trimmed == "" {
		return "", fmt.Errorf("empty URI")
	}

	fileURL, err := url.Parse(trimmed)
	if err != nil {
		return "", fmt.Errorf("invalid URI: %w", err)
	}
	if fileURL.Scheme != "" && fileURL.Scheme != "file" {
		return "", fmt.Errorf("unsupported URI scheme: %s", fileURL.Scheme)
	}

	decodedPath, err := url.PathUnescape(fileURL.Path)
	if err != nil {
		return "", fmt.Errorf("failed to decode URI path: %w", err)
	}
	if decodedPath == "" {
		return "", fmt.Errorf("empty file path")
	}

	if goos == "windows" {
		decodedPath = windowsDrivePathPattern.ReplaceAllString(decodedPath, "$1")
	}

	return decodedPath, nil
}

func ExtractProgression(line string) (float64, bool) {
	matches := progressionPattern.FindStringSubmatch(line)
	if len(matches) < 2 {
		return 0, false
	}

	value, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return 0, false
	}

	return value, true
}
