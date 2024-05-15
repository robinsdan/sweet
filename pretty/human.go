package pretty

import (
	"fmt"
	"github.com/robinsdan/sweet"
	"strings"
)

var aliases = make(map[string]string)

func Alias(origin fmt.Stringer, alias string) {
	aliases[origin.String()] = alias
}

func Human(origin string) string {
	for o, r := range aliases {
		origin = strings.ReplaceAll(origin, o, r)
	}

	return origin
}

func Humanf(fmter string, args ...interface{}) string {
	origin := fmt.Sprintf(fmter, args...)
	for o, r := range aliases {
		origin = strings.ReplaceAll(origin, o, r)
	}

	return origin
}

func JsonStr(v interface{}) string {
	return Human(sweet.JsonStr(v))
}

func JsonString(v interface{}) string {
	return Human(sweet.JsonString(v))
}
