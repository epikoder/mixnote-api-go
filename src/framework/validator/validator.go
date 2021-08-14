package validator

import (
	"fmt"
	"net/http"
	"strings"
)

type (
	validator struct {
		R        *http.Request
		Rules    map[string][]string
		Messages map[string][]string
	}
)

var (
	applicationJSON      string = "application/json"
	applicationURLEncode string = "application/x-www-form-urlencoded"
)


func SetDefaultRequired() {
	defaultRequired = true
}

func NewValidator(r *http.Request, rules map[string][]string, messages map[string][]string) *validator {
	return &validator{
		R:        r,
		Rules:    rules,
		Messages: messages,
	}
}

func (v *validator) Validate() (i FormVar, errBag ErrorBag, err error) {
	ct := v.R.Header.Get("Content-Type")
	if ct == "" {
		ct = applicationJSON
	}
	switch ct {
	case applicationJSON:
		if err = getFromJson(&i, v.R); err != nil {
			return nil, nil, err
		}
	case applicationURLEncode:
		if err = getFromForm(&i, v.R); err != nil {
			return nil, nil, err
		}
	}

	errBag = make(ErrorBag)
	for field, rules := range v.Rules {
		s := fmt.Sprintf("%v", rules)
		hasRequired = strings.Contains(s, "required")
		hasNullable = strings.Contains(s, "nullable")

		for _, rule := range rules {
			value, ok := i[field]
			if !ok {
				continue
			}

			if defaultRequired && isEmpty(value) || hasRequired && isEmpty(value) {
				if len(errBag[field]) < 1 {
					errBag[field] = append(errBag[field], "field is required")
				}
				continue
			}

			if hasNullable && isEmpty(value) {
				continue
			}

			rf, err := getRuleFunc(rule)
			if err != nil {
				return nil, nil, err
			}

			if err := rf(rule, nil, &value); err != nil {
				errBag[field] = append(errBag[field], err.Error())
			}

			if value != nil || value != "" {
				i[field] = value
			}
		}
	}
	return
}

