package utils

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// StructToURLValues converts a struct with `url` tags into url.Values.
// Handles basic types like string, int, bool pointers.
// Note: This is a simplified implementation. I tried not to use any other external packages.
// I also tried to make it as simple as possible.
// This was fu**ing hard to implement, because i had to learn about reflection and url encoding.
func StructToURLValues(s any) (url.Values, error) {
	values := url.Values{}
	if s == nil {
		return values, nil
	}

	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("tmdb:utils: expected a struct or pointer to struct, got %T", s)
	}

	t := v.Type()
	for i := range v.NumField() {
		field := t.Field(i)
		value := v.Field(i)

		// Skiping unexported fields
		if !field.IsExported() {
			continue
		}

		// Get the tag, skip if "-"
		tag := field.Tag.Get("url")
		if tag == "-" {
			continue
		}

		// Parse tag options (e.g., "name,omitempty")
		parts := strings.Split(tag, ",")
		urlKey := parts[0]
		omitempty := false
		if len(parts) > 1 && parts[1] == "omitempty" {
			omitempty = true
		}

		if urlKey == "" {
			urlKey = field.Name // Defaults to field name if tag name is empty
		}

		// Handle pointer types
		if value.Kind() == reflect.Ptr {
			if value.IsNil() {
				if !omitempty {
					// Since omitempty is not set, i decided to skip nil pointers, because TMDB API doesn't accept them.
				}
				continue
			}
			value = value.Elem() // Dereference the pointer
		}

		// Convert value to string
		// This section is jaunty, inspired from https://github.com/gorilla/schema
		// To implement something like this, i'd recommend just using the gorilla/schema package.
		var strValue string
		switch value.Kind() {
		case reflect.String:
			strValue = value.String()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			strValue = strconv.FormatInt(value.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			strValue = strconv.FormatUint(value.Uint(), 10)
		case reflect.Float32, reflect.Float64:
			strValue = strconv.FormatFloat(value.Float(), 'f', -1, 64)
		case reflect.Bool:
			strValue = strconv.FormatBool(value.Bool())
		case reflect.Slice:
			// Handle slices, typically comma-separated for TMDb
			if value.Type().Elem().Kind() == reflect.String {
				strSlice := make([]string, value.Len())
				for j := 0; j < value.Len(); j++ {
					strSlice[j] = value.Index(j).String()
				}
				strValue = strings.Join(strSlice, ",")
			} else if value.Type().Elem().Kind() == reflect.Int {
				intSlice := make([]string, value.Len())
				for j := 0; j < value.Len(); j++ {
					intSlice[j] = strconv.FormatInt(value.Index(j).Int(), 10)
				}
				strValue = strings.Join(intSlice, ",") // Joins int slice into a comma separated string, although TMDB API uses pipe (|) for some endpoints. I decided to use comma, because it's more common. (Will add pipe support later)
				// Also probably should add support for uint slices, but i don't need it for now.
			} else {
				// Skip unsupported slice types for now
				continue
			}
		default:
			// Skip unsupported types
			continue
		}

		// Add to values if not omitempty or if value is not zero/empty
		if !omitempty || (omitempty && strValue != "" && strValue != "0" && strValue != "false") {
			values.Set(urlKey, strValue)
		}
	}

	return values, nil
}
