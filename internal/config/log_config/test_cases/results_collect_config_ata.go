package test_cases

var (
	SuccessfulResult = map[string]map[string]string{
		"cfg_with_correct_tags": {
			"host":            "127.0.0.1",
			"port":            "80",
			"user_exists":     "false",
			"password_exists": "true",
		},
	}

	WithoutPrivateFields = map[string]map[string]string{
		"cfg_with_private_fields": {
			"user_exists":     "true",
			"password_exists": "true",
		},
	}
)
