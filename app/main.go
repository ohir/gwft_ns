package main

import (
	"fmt"
	"runtime/debug"

	c "example.com/mC"
	p "example.com/mP"
	s "example.com/mS"
)

func main() {
	fmt.Printf("\n%s\n---\n", Info())
	println(p.Say())
	println(s.Say())
	println(c.Say())
}

func Info() string {
	bi, _ := debug.ReadBuildInfo()

	r := fmt.Sprintf("Main module %s [%s/%s]\nUses:",
		bi.Main.Path, Vstamp, bi.Main.Version)
	for _, mo := range bi.Deps {
		r += "\n\t" + mo.Path
		if mo.Replace.Path != "" {
			r += "; from " + mo.Replace.Path
			r += " (instead of " + mo.Version + ")"
		} else {
			r += " @" + mo.Version
		}

		if mo.Replace.Version != "" {
			r += "; instead of " + mo.Replace.Version
		}
	}
	return r
}
