package validator

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/mixnote/mixnote-api-go/src/framework/utilities"
)

type (
	validator struct {
		R        *http.Request
		Rules    map[string][]string
		Messages map[string][]string
		Form     FormVar
		ErrorBag ErrorBag
	}
)

var (
	applicationJSON      string = "application/json"
	applicationURLEncode string = "application/x-www-form-urlencoded"
	applicationFormData  string = "multipart/form-data"
)

func SetDefaultRequired() {
	defaultRequired = true
}

func SetFileMemory(i int) {
	defaultMemory = i
}

func NewValidator(r *http.Request, rules map[string][]string, messages map[string][]string) *validator {
	return &validator{
		R:        r,
		Rules:    rules,
		Messages: messages,
	}
}

func (v *validator) Validate() (FormVar, ErrorBag, error) {
	v.ErrorBag = make(ErrorBag)
	fn := func() string {
		s := v.R.Header.Get("Content-Type")
		if strings.Contains(s, ";") {
			return strings.Split(s, ";")[0]
		}
		return s
	}
	ct := fn()

	utilities.Console.Debug(ct)
	if ct == "" {
		ct = applicationJSON
	}

	v.Form = make(FormVar)
	switch ct {
	case applicationJSON:
		getFromJson(&v.Form, v.R)
	case applicationURLEncode:
		getFromForm(&v.Form, v.R)
	case applicationFormData:
		getFormData(&v.Form, v.R)
	default:
		return nil, nil, fmt.Errorf("Content-Type unknown")
	}

	callRF := func(field, rule string, value interface{}, rf ruleFunc) {
		if err := rf(rule, nil, value); err != nil {
			v.addError(field, "field is required")
		}
		if value != nil || value != "" {
			v.Form[field] = value
		}
	}

	for field, rules := range v.Rules {
		s := fmt.Sprintf("%v", rules)
		hasRequired = strings.Contains(s, "required")
		hasNullable = strings.Contains(s, "nullable")
		hasFile := strings.Contains(s, "file")
		value, ok := v.Form[field]

		if !ok && !hasFile {
			v.addError(field, "field is required")
			continue
		}

		if hasFile {
			v.validateFile(field, rules, (defaultRequired || hasRequired))
			continue
		}

		for _, rule := range rules {
			rf, err := getRuleFunc(rule)
			if err != nil {
				return nil, nil, err
			}

			if (defaultRequired || hasRequired) && isEmpty(value) {
				v.addError(field, "field is required")
				continue
			}

			if hasNullable && isEmpty(value) {
				continue
			}
			callRF(field, rule, &value, rf)
		}
	}
	return v.Form, v.ErrorBag, nil
}

func (v *validator) validateFile(field string, rules []string, isRequired bool) {
	_, _, err := v.R.FormFile(field)
	if err != nil && isRequired {
		v.addError(field, "field is required")
		return
	}
	for _, rule := range rules {
		fF, err := getFileRuleFunc(rule)
		if err != nil {
			v.addError(field, err.Error())
		}
		fF()
	}

}

func (v *validator) addError(field string, message string) {
	if len(v.ErrorBag[field]) < 1 {
		v.ErrorBag[field] = append(v.ErrorBag[field], message)
	}
}
