package logging

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type GetLoggerConfigTestCase struct {
	TestName      string
	LogLevel      string
	LogDevMode    string
	ExpectedError string
}

func TestGetLoggerConfig(t *testing.T) {
	tests := []GetLoggerConfigTestCase{
		{
			TestName:   "successful load config",
			LogLevel:   "1",
			LogDevMode: "true",
		},
		{
			TestName:      "env SEARCH_DATA_EXTRACTOR_LOG_LEVEL not enabled",
			LogLevel:      "",
			LogDevMode:    "true",
			ExpectedError: "parsing \"\": invalid syntax",
		},
		{
			TestName:      "env SEARCH_DATA_EXTRACTOR_LOG_LEVEL is not int",
			LogLevel:      "test",
			LogDevMode:    "true",
			ExpectedError: "parsing \"test\": invalid syntax",
		},
		{
			TestName:      "value of env SEARCH_DATA_EXTRACTOR_LOG_LEVEL is less than -1",
			LogLevel:      "-5",
			LogDevMode:    "true",
			ExpectedError: "invalid log level",
		},
		{
			TestName:      "value of env SEARCH_DATA_EXTRACTOR_LOG_LEVEL is greater than 2",
			LogLevel:      "5",
			LogDevMode:    "true",
			ExpectedError: "invalid log level",
		},
		{
			TestName:      "env SEARCH_DATA_EXTRACTOR_LOG_DEV_MODE not enabled",
			LogLevel:      "1",
			LogDevMode:    "",
			ExpectedError: "parsing \"\": invalid syntax",
		},
		{
			TestName:      "env SEARCH_DATA_EXTRACTOR_LOG_DEV_MODE is not bool",
			LogLevel:      "1",
			LogDevMode:    "test",
			ExpectedError: "parsing \"test\": invalid syntax",
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			os.Clearenv()

			t.Setenv("SEARCH_DATA_EXTRACTOR_LOG_LEVEL", test.LogLevel)
			t.Setenv("SEARCH_DATA_EXTRACTOR_LOG_DEV_MODE", test.LogDevMode)

			_, err := LoadLoggerConfig()
			if test.ExpectedError != "" {
				if err != nil {
					assert.ErrorContains(t, err, test.ExpectedError)
				} else {
					t.Fatalf("error expected: %s", test.ExpectedError)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
