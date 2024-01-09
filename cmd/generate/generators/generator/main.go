/*
Copyright © 2024 Jonathon Hawkins <garbage@thedumpster.com>
*/
package generator

import (
	"os"
	"strings"
)

type Config struct {
	Repository      string
	TargetDirectory string
	PackageName     string
}

type Generator = func(config Config)

func RightRemoveSegment(path string) string {
	parts := strings.Split(path, string(os.PathSeparator))
	allByLast := parts[:len(parts)-1]
	return strings.Join(allByLast, string(os.PathSeparator))
}
