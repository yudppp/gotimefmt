package timefmt

import (
	"strings"
)

var replaceMap []string

func init() {
	replaceMap = make([]string, 0, 16)
	AddReplaceSet("MM", "01")
	AddReplaceSet("DD", "02")
	AddReplaceSet("hh", "15")
	AddReplaceSet("h", "3")
	AddReplaceSet("mm", "04")
	AddReplaceSet("ss", "05")
	AddReplaceSet("YYYY", "2006")
}

func AddReplaceSet(k, v string) {
	replaceMap = append(replaceMap, k, v)
}

func Format(p string) string {
	r := strings.NewReplacer(replaceMap...)
	return r.Replace(p)
}
