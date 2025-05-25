package log_config

import (
	"testing"

	testCases "gnss_corr/pkg/log_config/test_cases"
)

func TestStructConfigToMap(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		config         any
		expectedName   string
		expectedResult map[string]string
	}{
		{
			name:           "Successful with all tags",
			config:         testCases.CorrectCfg,
			expectedName:   "cfg_with_correct_tags",
			expectedResult: testCases.SuccessfulResultWithCorrectTags,
		},
		{
			name:           "Successful with nested tag",
			config:         testCases.CfgWithNested,
			expectedName:   "cfg_with_nested_tags",
			expectedResult: testCases.SuccessfulWithNestedTag,
		},
		{
			name:           "With incorrect logSecret tag",
			config:         testCases.IncorrectCfg,
			expectedName:   "cfg_with_incorrect_tags",
			expectedResult: testCases.ResultWithIncorrectTag,
		},
		{
			name:           "Config without tags",
			config:         testCases.WithoutTagsCfg,
			expectedName:   "cfg_without_tags",
			expectedResult: testCases.ResulWithoutTags,
		},
		{
			name:           "Config with private fields",
			config:         testCases.PrivateFieldCfg,
			expectedName:   "cfg_with_private_fields",
			expectedResult: testCases.ResultWithPrivateFields,
		},
		{
			name:           "Config with empty string values",
			config:         testCases.EmptyValuesCfg,
			expectedName:   "cfg_with_empty_values",
			expectedResult: testCases.ResultWithEmptyValues,
		},
		{
			name:           "Config with pointer",
			config:         &testCases.CorrectCfg,
			expectedName:   "cfg_with_correct_tags",
			expectedResult: testCases.SuccessfulResultWithCorrectTags,
		},
		{
			name:           "Config with nested structure",
			config:         testCases.NestedStructCfg,
			expectedName:   "config_with_nested_struct",
			expectedResult: testCases.SuccessfulResultWithNestedStruct,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			name, info := structConfigToMap(tc.config)

			if name != tc.expectedName {
				t.Errorf("Expected config name '%s', but got '%s'", tc.expectedName, name)
			}

			if !mapsAreEqual(info, tc.expectedResult) {
				t.Errorf("Expected %v, but got %v", tc.expectedResult, info)
			}
		})
	}
}

func TestCollectConfigData(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		cfg            any
		expectedResult map[string]map[string]string
	}{
		{
			name:           "Successful",
			cfg:            testCases.SuccessfulCfg,
			expectedResult: testCases.SuccessfulResult,
		},
		{
			name:           "with private fields",
			cfg:            testCases.CompareWithPrivateFields,
			expectedResult: testCases.WithoutPrivateFields,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			data := collectConfigData(tc.cfg)

			if !mapsAreEqual2D(data, tc.expectedResult) {
				t.Errorf("Expected %v, but got %v", tc.expectedResult, data)
			}
		})
	}
}

func mapsAreEqual(obtainedMap, expectedMap map[string]string) bool {
	if len(obtainedMap) != len(expectedMap) {
		return false
	}

	for key, obtainedValue := range obtainedMap {
		if expectedValue, ok := expectedMap[key]; !ok || obtainedValue != expectedValue {
			return false
		}
	}
	return true
}

func mapsAreEqual2D(obtainedMap, expectedMap map[string]map[string]string) bool {
	if len(obtainedMap) != len(expectedMap) {
		return false
	}

	for key, obtainedValue := range obtainedMap {
		expectedValue, ok := expectedMap[key]
		if !ok || !mapsAreEqual(obtainedValue, expectedValue) {
			return false
		}
	}
	return true
}
