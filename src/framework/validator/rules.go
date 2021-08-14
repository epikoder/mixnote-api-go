package validator

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/mixnote/mixnote-api-go/src/framework/utilities"
)

type (
	ruleFunc func(rule string, message interface{}, value interface{}) (err error)
)

var (
	ruleMapFunc          = make(map[string]ruleFunc)
	defaultRequired bool = false
	hasRequired     bool = false
	hasNullable     bool = false
)

func AddRule(name string, rf ruleFunc) {
	if isRuleExist(name) {
		utilities.Console.Fatal("validator: rule %s already exist", name)
	}
	ruleMapFunc[name] = rf
}

func init() {
	AddRule("nullable", func(rule string, message, value interface{}) (err error) {
		return nil
	})

	AddRule("required", func(rule string, message, value interface{}) (err error) {
		err = fmt.Errorf("field is required")
		if s, ok := message.(string); ok && s != "" {
			err = fmt.Errorf(s)
		}

		v := *(value.(*interface{}))
		if v == nil || v == "" {
			return err
		}
		return nil
	})

	AddRule("email", func(rule string, message interface{}, value interface{}) (err error) {
		err = fmt.Errorf("invalid email value")
		if s, ok := message.(string); ok && s != "" {
			err = fmt.Errorf(s)
		}

		v, ok := readStringFromPtr(value)
		if !ok {
			return fmt.Errorf("should be of type string")
		}
		if !isEmail(v) {
			return err
		}
		return nil
	})

	AddRule("numeric", func(rule string, message interface{}, value interface{}) (err error) {
		if *(value.(*interface{})) != nil {
			err = fmt.Errorf("invalid type expecting numeric got %s", reflect.TypeOf(*(value.(*interface{}))).Kind().String())
		} else {
			err = fmt.Errorf("invalid value numeric expected")
		}

		if s, ok := message.(string); ok && s != "" {
			err = fmt.Errorf(s)
		}

		v, ok := readStringFromPtr(value)
		if !ok && !isEmpty(value) {
			v = strings.Split(fmt.Sprintf("%f", (*(value.(*interface{}))).(float64)), ".")[0]
		}

		if !isNumeric(v) {
			return err
		}

		if *(value.(*interface{})), err = strconv.ParseUint(v, 10, 64); err != nil {
			return fmt.Errorf("can not convert value to numeric")
		}

		return nil
	})

	AddRule("alpha_numeric", func(rule string, message, value interface{}) (err error) {
		v := *(value.(*interface{}))
		if v != nil || v == "" {
			err = fmt.Errorf("invalid type expecting alpha-numeric got %s", reflect.TypeOf(*(value.(*interface{}))).Kind().String())
		} 
		if s, ok := message.(string); ok && s != "" {
			err = fmt.Errorf(s)
		}

		s, _ := readStringFromPtr(value)
		if !isAlphaNumeric(s) {
			return err
		}
		return nil
	})

	AddRule("min", func(rule string, message, value interface{}) (err error) {
		var l int
		if l, err = strconv.Atoi(strings.Split(rule, ":")[1]); err != nil {
			return fmt.Errorf("invalid min value specified in rule")
		}
		if v := fmt.Sprintf("%v", *(value.(*interface{}))); len(v) < l {
			return fmt.Errorf("minimium value is %d", l)
		}
		return nil
	})

	AddRule("max", func(rule string, message, value interface{}) (err error) {
		var l int
		if l, err = strconv.Atoi(strings.Split(rule, ":")[1]); err != nil {
			return fmt.Errorf("invalid max value specified in rule")
		}
		if v := fmt.Sprintf("%v", *(value.(*interface{}))); len(v) < l {
			return fmt.Errorf("minimium value is %d", l)
		}
		return nil
	})
}
