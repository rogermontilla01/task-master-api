package config

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

type Configuration interface {
	Validate() []error
}

func ValidateEnvConfig(config interface{}) []error {
	var errors []error
	v := reflect.ValueOf(config).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tag := t.Field(i).Tag.Get("env")
		tags := strings.Split(tag, ",")

		envVar := ""
		required := false
		defaultVal := ""

		// Parse the struct tags
		for _, t := range tags {
			if strings.HasPrefix(t, "value=") {
				envVar = strings.TrimPrefix(t, "value=")
			} else if t == "required" {
				required = true
			} else if strings.HasPrefix(t, "default=") {
				defaultVal = strings.TrimPrefix(t, "default=")
			}
		}

		envVal, exists := os.LookupEnv(envVar)
		if !exists {
			if defaultVal != "" {
				log.Warn().Msgf("Using default value for %s: %s", envVar, defaultVal)
				envVal = defaultVal
			} else if required {
				errors = append(errors, fmt.Errorf("%s is required but not set", envVar))
				continue
			}
		}

		switch field.Kind() {
		case reflect.Bool:
			value, err := strconv.ParseBool(envVal)
			if err != nil {
				errors = append(errors, fmt.Errorf("failed to cast %s to bool", envVar))
				continue
			}
			field.SetBool(value)
		case reflect.Int:
			value, err := strconv.Atoi(envVal)
			if err != nil {
				errors = append(errors, fmt.Errorf("failed to cast %s to int", envVar))
				continue
			}
			field.SetInt(int64(value))
		case reflect.Int64:
			value, err := strconv.ParseInt(envVal, 10, 64)
			if err != nil {
				errors = append(errors, fmt.Errorf("failed to cast %s to int", envVar))
				continue
			}
			field.SetInt(value)
		case reflect.String:
			field.SetString(envVal)
		case reflect.Slice:
			sliceValues := strings.Split(envVal, ",")
			slice := reflect.MakeSlice(field.Type(), len(sliceValues), len(sliceValues))
			for i, val := range sliceValues {
				elem := slice.Index(i)
				switch elem.Kind() {
				case reflect.Bool:
					value, err := strconv.ParseBool(val)
					if err != nil {
						errors = append(errors, fmt.Errorf("failed to cast %s to bool", envVar))
						continue
					}
					elem.SetBool(value)
				case reflect.Int:
					value, err := strconv.Atoi(val)
					if err != nil {
						errors = append(errors, fmt.Errorf("failed to cast %s to int", envVar))
						continue
					}
					elem.SetInt(int64(value))
				case reflect.Int64:
					value, err := strconv.ParseInt(val, 10, 64)
					if err != nil {
						errors = append(errors, fmt.Errorf("failed to cast %s to int", envVar))
						continue
					}
					elem.SetInt(value)
				case reflect.String:
					elem.SetString(val)
				default:
					errors = append(errors, fmt.Errorf("unsupported element type for %s", envVar))
				}
			}
			field.Set(slice)
		default:
			errors = append(errors, fmt.Errorf("unsupported field type for %s", envVar))
		}
	}

	switch c := config.(type) {
	case Configuration:
		if extraErrors := c.Validate(); len(extraErrors) > 0 {
			errors = append(errors, extraErrors...)
		}
	}

	return errors
}

func ValidateEnvConfigOrFail(config interface{}) {

	errors := ValidateEnvConfig(config)

	if len(errors) > 0 {
		for _, e := range errors {
			log.Error().Msg(e.Error())
		}
		// log.Fatal().Msg("Environment does not have a valid configuration")
		log.Panic().Msg("Environment does not have a valid configuration")
	}

}
