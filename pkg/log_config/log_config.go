package log_config

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"unicode"

	"go.uber.org/zap"
)

const (
	logTag = "logKey"
	nested = "nested"

	falseString = "false"
	trueString  = "true"

	logSecret = "secret"
)

func LogConfig(cfg any) {
	configMap := collectConfigData(cfg)
	logConfigData(configMap)
}

func collectConfigData(config any) map[string]map[string]string {
	configMap := make(map[string]map[string]string)

	v, t := dereferenceValueAndType(config)

	for i := 0; i < v.NumField(); i++ {
		subCfg := v.Field(i)
		subCfgType := t.Field(i)

		if !subCfgType.IsExported() {
			continue
		}

		if isStruct(subCfg) {
			if hasNestedStructs(subCfg) && subCfgType.Tag.Get(nested) != trueString {
				subCfgMap := collectConfigData(subCfg.Interface())
				for k, v := range subCfgMap {
					configMap[k] = v
				}
			} else {
				nameFromTitle, info := structConfigToMap(subCfg.Interface())

				if name := subCfgType.Tag.Get(logTag); name != "" {
					nameFromTitle = name
				}

				configMap[nameFromTitle] = info
			}
		}
	}

	return configMap
}

func hasNestedStructs(v reflect.Value) bool {
	if v.Kind() == reflect.Ptr {
		v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		if isStruct(field) {
			return true
		}
	}

	return false
}

func structConfigToMap(config any) (string, map[string]string) {
	result := make(map[string]string)

	v, t := dereferenceValueAndType(config)

	configName := toSnakeCase(t.Name())

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		if !field.IsExported() {
			continue
		}

		if field.Type.Kind() == reflect.Struct {
			_, nestedResult := structConfigToMap(value.Interface())
			for k, v := range nestedResult {
				result[k] = v
			}
			continue
		}

		key, secret := parseLogKey(field.Tag.Get(logTag))
		if key == "" {
			continue
		}

		formattedValue := formatValue(value, secret)
		result[buildKey(key, secret)] = formattedValue
	}

	return configName, result
}

func logConfigData(configMap map[string]map[string]string) {
	log := newLogger()
	//goland:noinspection GoUnhandledErrorResult
	defer log.Sync()

	jsonData, err := json.MarshalIndent(configMap, "", "  ")
	if err != nil {
		log.Error("Failed to marshal config data", zap.Error(err))
	}

	log.Info("application configuration data", zap.String("config", fmt.Sprintf("\n%s\n", string(jsonData))))

}

func isStruct(subCfg reflect.Value) bool {
	return subCfg.Kind() == reflect.Struct || (subCfg.Kind() == reflect.Ptr && subCfg.Elem().Kind() == reflect.Struct)
}

func dereferenceValueAndType(config any) (reflect.Value, reflect.Type) {
	v := reflect.ValueOf(config)
	t := reflect.TypeOf(config)

	if t.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}

	return v, t
}

func formatValue(value reflect.Value, secretKey bool) string {
	if secretKey {
		if !value.IsZero() {
			return trueString
		} else {
			return falseString
		}
	}

	return fmt.Sprintf("%v", value.Interface())
}

func buildKey(baseKey string, secretKey bool) string {
	if secretKey {
		return fmt.Sprintf("%s_exists", baseKey)
	}
	return baseKey
}

func toSnakeCase(input string) string {
	var result strings.Builder
	for i, r := range input {
		if unicode.IsUpper(r) {
			if i > 0 && (unicode.IsLower(rune(input[i-1])) || (i+1 < len(input) && unicode.IsLower(rune(input[i+1])))) {
				result.WriteByte('_')
			}
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

func parseLogKey(logKey string) (string, bool) {
	if logKey == "" {
		return "", false
	}

	var field string
	var hasSecret bool
	keys := strings.Split(logKey, ",")

	for i, key := range keys {
		trimmedKey := strings.TrimSpace(key)

		if i == 0 {
			field = trimmedKey
		}

		if strings.EqualFold(trimmedKey, logSecret) {
			hasSecret = true
		}
	}

	return field, hasSecret
}
