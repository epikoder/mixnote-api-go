package validator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

func getFromJson(i interface{}, r *http.Request) (err error) {
	if reflect.TypeOf(i).Kind() != reflect.Ptr {
		return fmt.Errorf("non pointer value given")
	}
	var b []byte
	if b, err = ioutil.ReadAll(r.Body); err != nil {
		return err
	}
	if err = json.Unmarshal(b, i); err != nil {
		return err
	}
	return
}

func getFromForm(i interface{}, r *http.Request) (err error) {
	if reflect.TypeOf(i).Kind() != reflect.Ptr {
		return fmt.Errorf("non pointer value given")
	}
	r.ParseForm()
	(*(i.(*FormVar))) = make(FormVar)
	for k, v := range r.Form {
		if strings.HasPrefix(k, "{") {
			return fmt.Errorf("wrong content-type header %s. Use %s instead", applicationURLEncode, applicationJSON)
		}
		(*(i.(*FormVar)))[k] = v[0]
	}
	return
}

func getRuleFunc(rule string) (ruleFunc, error) {
	var fn = func() string {
		if strings.Contains(rule, ":") {
			return strings.Split(rule, ":")[0]
		}
		return rule
	}

	rf, ok := ruleMapFunc[fn()]
	if !ok {
		return nil, fmt.Errorf("rule %s does not exist", rule)
	}
	return rf, nil
}


func readStringFromPtr(v interface{}) (string, bool) {
	s, ok := (*(v.(*interface{}))).(string)
	return s, ok
}
