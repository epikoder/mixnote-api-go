package validator

import "strings"

func isEmail(str string) bool {
	return regexEmail.MatchString(str)
}

func isNumeric(str string) bool {
	return regexNumeric.MatchString(str)
}

func isRuleExist(name string) bool {
	if strings.Contains(name, ":") {
		name = strings.Split(name, ":")[0]
	}
	_, ok := ruleMapFunc[name]
	return ok 
}

func isAlphaNumeric(s string) bool {
	return regexAlphaNumeric.MatchString(s)
}

func isEmpty(v interface{}) bool {
	return v == nil || v == ""
}
