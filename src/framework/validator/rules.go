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
	fileFunc  func() (err error)
	mimes   struct {
		JPEG string
		MPEG string
		FLAC string
		PNG  string
		ARC  string
	}
)

var (
	ruleMapFunc          = make(map[string]ruleFunc)
	defaultRequired bool = false
	hasRequired     bool = false
	hasNullable     bool = false

	/// Defined rules
	Nullable         = "nullable"
	Required         = "required"
	File             = "file"
	Email            = "email"
	Numeric          = "numeric"
	AlphaSpaceNoDash = "alpha_space_no_dash"
	Alpha            = "alpha"
	AlphaNumeric     = "alpha_numeric"
	AlphaNumericCase = "alpha_numeric_case"
	Min              = "min"
	Max              = "max"
	UUID             = "uuid"
	Mime             = "mime"

	fileMapRules = make(map[string]fileFunc)
	Mimes = &mimes{
		JPEG: "JPEG",
	}
)

func AddRule(name string, rf ruleFunc) {
	if isRuleExist(name) {
		utilities.Console.Fatal("validator: rule %s already exist", name)
	}
	ruleMapFunc[name] = rf
}

func AddMimeRule(name string, mf fileFunc) {
	if isMimeRuleExist(name) {
		utilities.Console.Fatal("validator: rule %s already exist", name)
	}
	fileMapRules[name] = mf
}

func init() {
	AddRule(Nullable, func(rule string, message, value interface{}) (err error) {
		return nil
	})

	AddRule(Required, func(rule string, message, value interface{}) (err error) {
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

	AddRule(Email, func(rule string, message interface{}, value interface{}) (err error) {
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

	AddRule(Numeric, func(rule string, message interface{}, value interface{}) (err error) {
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

	AddRule(Alpha, func(rule string, message, value interface{}) (err error) {
		v := *(value.(*interface{}))
		if v != nil || v == "" {
			err = fmt.Errorf("invalid type expecting alpha got %s", reflect.TypeOf(*(value.(*interface{}))).Kind().String())
		}
		if s, ok := message.(string); ok && s != "" {
			err = fmt.Errorf(s)
		}

		s, _ := readStringFromPtr(value)
		if !isAlpha(s) {
			return err
		}
		return nil
	})

	AddRule(AlphaSpaceNoDash, func(rule string, message, value interface{}) (err error) {
		v := *(value.(*interface{}))
		if v != nil || v == "" {
			err = fmt.Errorf("string without dash expected got %s", reflect.TypeOf(*(value.(*interface{}))).Kind().String())
		}
		if s, ok := message.(string); ok && s != "" {
			err = fmt.Errorf(s)
		}

		s, _ := readStringFromPtr(value)
		if !isAlphaSpaceNoDash(s) {
			return err
		}
		return nil
	})

	AddRule(AlphaNumeric, func(rule string, message, value interface{}) (err error) {
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

	AddRule(AlphaNumericCase, func(rule string, message, value interface{}) (err error) {
		err = fmt.Errorf("invalid type expecting alpha-numeric case sensitive without space")
		if s, ok := message.(string); ok && s != "" {
			err = fmt.Errorf(s)
		}

		s, _ := readStringFromPtr(value)
		if !isAlphaNumericCase(s) {
			return err
		}
		return nil
	})

	AddRule(Min, func(rule string, message, value interface{}) (err error) {
		var l int
		if l, err = strconv.Atoi(strings.Split(rule, ":")[1]); err != nil {
			return fmt.Errorf("invalid min value specified in rule")
		}
		if v := fmt.Sprintf("%v", *(value.(*interface{}))); len(v) < l {
			return fmt.Errorf("minimium value is %d", l)
		}
		return nil
	})

	AddRule(Max, func(rule string, message, value interface{}) (err error) {
		var l int
		if l, err = strconv.Atoi(strings.Split(rule, ":")[1]); err != nil {
			return fmt.Errorf("invalid max value specified in rule")
		}
		if v := fmt.Sprintf("%v", *(value.(*interface{}))); len(v) < l {
			return fmt.Errorf("minimium value is %d", l)
		}
		return nil
	})

	AddRule(UUID, func(rule string, message, value interface{}) (err error) {
		err = fmt.Errorf("invalid type expecting uuid")
		if s, ok := message.(string); ok && s != "" {
			err = fmt.Errorf(s)
		}

		s, _ := readStringFromPtr(value)
		if !isUUID(s) {
			return err
		}
		return nil
	})

	AddRule(File, func(rule string, message, value interface{}) (err error) {
		fmt.Printf("file is : %v \n", value == nil)
		return (func ()  error {
			if value == nil {
				return fmt.Errorf("field is required")
			}
			return nil})()
	})
	
	AddRule(Mime, func(rule string, message, value interface{}) (err error) {
		rule_ := strings.Split(rule, ":")
		if len(rule_) != 2 {
			return fmt.Errorf("no mime type passed")
		}

		_rules := strings.Split(rule_[1], ",")
		for _, name := range _rules {
			mf, ok := fileMapRules[name]
			if !ok {
				return fmt.Errorf("mime type unkwown")
			}
			if err := mf(); err != nil {
				return err
			}
		}
		return nil
	})
}

func init() {
	AddMimeRule(Mimes.JPEG, func() (err error) {
		return nil
	})
}
