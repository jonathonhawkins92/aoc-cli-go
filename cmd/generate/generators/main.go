/*
Copyright © 2024 Jonathon Hawkins <garbage@thedumpster.com>
*/
package generators

import (
	"github.com/jonathonhawkins92/aoc-cli-go/cmd/generate/generators/generator"
	"github.com/jonathonhawkins92/aoc-cli-go/cmd/generate/generators/generator/days"
	"github.com/jonathonhawkins92/aoc-cli-go/cmd/generate/generators/generator/root"
)

func Root(config generator.Config) {
	root.Create(config)
}

func Days(config generator.Config) {
	days.Create(config)
}
