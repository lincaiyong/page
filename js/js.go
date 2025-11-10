package js

import (
	"regexp"
)

var data = map[string]map[string]string{}

func Set(clsName, code string) {
	if data[clsName] == nil {
		data[clsName] = map[string]string{}
	}
	re := regexp.MustCompile(`(?ms)^function\s+([a-zA-Z_][a-zA-Z0-9_]*)\([^)]*\)\s*\{(?:.|\n)*?^}`)
	matches := re.FindAllString(code, -1)
	for _, match := range matches {
		nameRe := regexp.MustCompile(`function\s+([a-zA-Z_][a-zA-Z0-9_]*)`)
		nameMatch := nameRe.FindStringSubmatch(match)
		if len(nameMatch) > 1 {
			functionName := nameMatch[1]
			data[clsName][functionName] = match
		}
	}
}

func Get(clsName, funName string) string {
	if data[clsName] == nil {
		return ""
	}
	return data[clsName][funName]
}
