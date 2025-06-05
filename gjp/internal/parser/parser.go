package parser

import "errors"

func Parse(input string) error {
	trimmed := trimWhitespace(input)
	if trimmed == "{}" {
		return nil
	}
	return errors.New("invalid JSON")
}

func trimWhitespace(s string) string {
	result := ""
	for _, r := range s {
		if r != ' ' && r != '\n' && r != '\t' && r != '\r' {
			result += string(r)
		}
	}
	return result
}
