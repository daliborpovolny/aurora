package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/mail"
	"reflect"
	"strconv"
	"strings"
)

// func Decode(r *http.Request, dst interface{}) error {
// 	ct := r.Header.Get("Content-type")

// 	// json
// 	if strings.HasPrefix(ct, "application/json") {
// 		return json.NewDecoder(r.Body).Decode(&dst)
// 	}

// 	// html form
// 	if err := r.ParseForm(); err != nil {
// 		return err
// 	}

// 	val := reflect.ValueOf(dst).Elem()
// 	typ := val.Type()

// 	for i := 0; i < val.NumField(); i++ {
// 		field := val.Field(i)
// 		structField := typ.Field(i)

// 		tag := structField.Tag.Get(structField.Name)
// 		if tag == "" {
// 			tag = strings.ToLower(structField.Name)
// 		} else {
// 			tag = strings.Split(tag, ",")[0]
// 		}

// 		formValue := r.FormValue(tag)
// 		if formValue == "" {
// 			continue
// 		}

// 		if field.CanSet() {
// 			switch field.Kind() {

// 			case reflect.String:
// 				field.SetString(formValue)

// 			case reflect.Int, reflect.Int64:
// 				n, err := strconv.ParseInt(formValue, 10, 64)
// 				if err != nil {
// 					return fmt.Errorf("invalid int for field %s", tag)
// 				}
// 				field.SetInt(n)

// 			case reflect.Float64:
// 				f, err := strconv.ParseFloat(formValue, 64)
// 				if err != nil {
// 					return fmt.Errorf("invalid float for field %s", tag)
// 				}
// 				field.SetFloat(f)

// 			case reflect.Slice:
// 				if field.Type().Elem().Kind() == reflect.String {
// 					field.Set(reflect.ValueOf(r.Form[tag]))
// 				} else if field.Type().Elem().Kind() == reflect.Int {
// 					var ints []int
// 					for _, s := range r.Form[tag] {
// 						n, err := strconv.Atoi(s)
// 						if err != nil {
// 							return fmt.Errorf("invalid int in list for field %s", tag)
// 						}
// 						ints = append(ints, n)
// 					}
// 					field.Set(reflect.ValueOf(ints))
// 				}
// 			}
// 		}
// 	}
// 	return nil
// }

func DecodeJson(r *http.Request, dst interface{}) error {
	ct := r.Header.Get("Content-Type")

	if !strings.HasPrefix(ct, "application/json") {
		return errors.New("expected json request")
	}

	return json.NewDecoder(r.Body).Decode(dst)
}

func DecodeForm(r *http.Request, dst interface{}) error {
	ct := r.Header.Get("Content-Type")

	if strings.HasPrefix(ct, "application/json") {
		return errors.New("expected non json request")
	}

	if err := r.ParseForm(); err != nil {
		return err
	}

	val := reflect.ValueOf(dst).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		structField := typ.Field(i)

		tag := structField.Tag.Get("json")
		if tag == "" {
			tag = strings.ToLower(structField.Name)
		} else {
			tag = strings.Split(tag, ",")[0]
		}

		if tag == "-" {
			continue
		}

		formValue := r.FormValue(tag)
		if formValue == "" && field.Kind() != reflect.Slice {
			continue
		}

		if field.CanSet() {
			switch field.Kind() {
			case reflect.String:
				field.SetString(formValue)

			case reflect.Int, reflect.Int64:
				n, err := strconv.ParseInt(formValue, 10, 64)
				if err != nil {
					return fmt.Errorf("invalid int for field %s", tag)
				}
				field.SetInt(n)

			case reflect.Float64:
				f, err := strconv.ParseFloat(formValue, 64)
				if err != nil {
					return fmt.Errorf("invalid float for field %s", tag)
				}
				field.SetFloat(f)

			case reflect.Slice:
				values := r.Form[tag]
				if field.Type().Elem().Kind() == reflect.String {
					field.Set(reflect.ValueOf(values))
				} else if field.Type().Elem().Kind() == reflect.Int {
					var ints []int
					for _, s := range values {
						n, err := strconv.Atoi(s)
						if err != nil {
							return fmt.Errorf("invalid int in list for field %s", tag)
						}
						ints = append(ints, n)
					}
					field.Set(reflect.ValueOf(ints))
				}
			}
		}
	}
	return nil
}

func IsJson(r *http.Request) bool {
	cnt := r.Header.Get("Content-Type")
	return strings.HasPrefix(cnt, "application/json")
}

type JSONResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func IsEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
