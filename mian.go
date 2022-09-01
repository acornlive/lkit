package main

import (
	"fmt"
	"github.com/acornlive/lkit/logkit"
	"github.com/acornlive/lkit/strkit"
	"time"
)

func main() {

	fmt.Println(strkit.FormatTime(time.Now()))
	logkit.Log().SetPrefix("ACORN").SetLevel(logkit.ERROR)

	ma := map[string]string{
		"name": "wty",
		"age":  "aa",
	}
	logkit.Info("sss")
	logkit.Error("sss id {},name {} ，map {}", 1, "wty", ma)
}
