package test_cases

var (
	SuccessfulResultWithCorrectTags = map[string]string{
		"host":            "127.0.0.1",
		"port":            "80",
		"user_exists":     "false",
		"password_exists": "true",
	}

	SuccessfulWithNestedTag = map[string]string{
		"host":            "127.0.0.1",
		"port":            "80",
		"user_exists":     "true",
		"password_exists": "true",
	}

	ResultWithIncorrectTag = map[string]string{
		"host":     "127.0.0.1",
		"port":     "80",
		"user":     "root",
		"password": "123456",
	}

	ResultWithPrivateFields = map[string]string{
		"user_exists":     "true",
		"password_exists": "true",
	}

	ResulWithoutTags = map[string]string{}

	ResultWithEmptyValues = map[string]string{
		"host":            "",
		"port":            "0",
		"user_exists":     "true",
		"password_exists": "true",
		"db_name":         "",
	}

	SuccessfulResultWithNestedStruct = map[string]string{
		"host":            "127.0.0.1",
		"port":            "80",
		"user_exists":     "false",
		"password_exists": "true",
		"db_name":         "test",
	}
)
