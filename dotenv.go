package dotenv

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

// Load env vars from a file and load them into environment for the current process.
// Doesn't overrides already set environment variables.
//
func Load(filename string) error {
	return loadFile(filename, false)
}

// ForceLoad env vars from a file and load them into environment for the current process.
// Does overrides already set environment variables.
//
func ForceLoad(filename string) error {
	return loadFile(filename, true)
}

func loadFile(filename string, overload bool) error {
	file, errOpen := os.Open(filename)
	if errOpen != nil {
		return errOpen
	}
	defer file.Close()

	envMap, errParse := parseReader(file)
	if errParse != nil {
		return errParse
	}

	osEnv := os.Environ()
	currEnv := make(map[string]struct{}, len(osEnv))

	for _, v := range osEnv {
		key := strings.Split(v, "=")[0]
		currEnv[key] = struct{}{}
	}

	for key, value := range envMap {
		_, ok := currEnv[key]
		if !ok || overload {
			os.Setenv(key, value)
		}
	}
	return nil
}

func parseReader(r io.Reader) (map[string]string, error) {
	lines, errScan := scanLines(r)
	if errScan != nil {
		return nil, errScan
	}

	envMap := make(map[string]string, len(lines))
	for _, line := range lines {
		if isIgnored(line) {
			continue
		}
		key, value, errParse := parseLine(line, envMap)
		if errParse != nil {
			return nil, errParse
		}
		envMap[key] = value
	}
	return envMap, nil
}

func scanLines(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func isIgnored(line string) bool {
	s := strings.TrimSpace(line)
	return len(s) == 0 || s[0] == '#'
}

func parseLine(line string, envMap map[string]string) (string, string, error) {
	idx := strings.Index(line, "=")
	if idx == -1 {
		return "", "", errors.New("Can't separate key from value")
	}
	splitString := strings.SplitN(line, "=", 2)
	return splitString[0], splitString[1], nil
}
