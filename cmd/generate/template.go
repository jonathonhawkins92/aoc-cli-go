/*
Copyright © 2024 Jonathon Hawkins <garbage@thedumpster.com>
*/
package generate

import (
	_ "embed"
)

type Template struct{}

//go:embed templates/main.go.tmpl
var templateMainRoot []byte

func (c Template) MainRoot() []byte {
	return templateMainRoot
}

//go:embed templates/go.mod.tmpl
var templateMod []byte

func (c Template) Mod() []byte {
	return templateMod
}

//go:embed templates/day/main.go.tmpl
var templateDayRoot []byte

func (c Template) DayRoot() []byte {
	return templateDayRoot
}

//go:embed templates/day/part1/main.go.tmpl
var templateDayPart1 []byte

func (c Template) DayPart1() []byte {
	return templateDayPart1
}

//go:embed templates/day/part1/main_test.go.tmpl
var templateDayPart1Test []byte

func (c Template) DayPart1Test() []byte {
	return templateDayPart1Test
}

//go:embed templates/day/part2/main.go.tmpl
var templateDayPart2 []byte

func (c Template) DayPart2() []byte {
	return templateDayPart2
}

//go:embed templates/day/part2/main_test.go.tmpl
var templateDayPart2Test []byte

func (c Template) DayPart2Test() []byte {
	return templateDayPart2Test
}
