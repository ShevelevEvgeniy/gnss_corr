package test_cases

type (
	CfgWithCorrectTags struct {
		Host       string `logKey:"host"`
		Port       int    `logKey:"port"`
		DBUser     string `logKey:"user,secret"`
		DBPassword string `logKey:"password,secret"`
	}

	CfgWithNestedTags struct {
		Host string     `logKey:"host"`
		Port int        `logKey:"port"`
		User UserStruct `nested:"true"`
	}

	UserStruct struct {
		DBUser     string `logKey:"user,secret"`
		DBPassword string `logKey:"password,secret"`
	}

	CfgWithIncorrectTags struct {
		Host       string `logKey:"host"`
		Port       int    `logKey:"port"`
		DBUser     string `logKey:"user"`
		DBPassword string `logKey:"password,ecret"`
	}

	CfgWithoutTags struct {
		Host       string
		Port       int
		DBUser     string
		DBPassword string
		DBName     string
	}

	CfgWithPrivateFields struct {
		host       string `logKey:"host"`
		port       int    `logKey:"port"`
		DBUser     string `logKey:"user,secret"`
		DBPassword string `logKey:"password,secret"`
		dbName     string `logKey:"db_name"`
	}

	CfgWithEmptyValues struct {
		Host       string `logKey:"host"`
		Port       int    `logKey:"port"`
		DBUser     string `logKey:"user,secret"`
		DBPassword string `logKey:"password,secret"`
		DBName     string `logKey:"db_name"`
	}

	ConfigWithNestedStruct struct {
		CfgWithCorrectTags
		DBName string `logKey:"db_name"`
	}
)

var (
	CorrectCfg = CfgWithCorrectTags{
		Host:       "127.0.0.1",
		Port:       80,
		DBUser:     "",
		DBPassword: "123456",
	}

	CfgWithNested = CfgWithNestedTags{
		Host: "127.0.0.1",
		Port: 80,
		User: UserStruct{
			DBUser:     "root",
			DBPassword: "secret",
		},
	}

	IncorrectCfg = CfgWithIncorrectTags{
		Host:       "127.0.0.1",
		Port:       80,
		DBUser:     "root",
		DBPassword: "123456",
	}

	WithoutTagsCfg = CfgWithoutTags{
		Host:       "127.0.0.1",
		Port:       80,
		DBUser:     "root",
		DBPassword: "123456",
		DBName:     "test",
	}

	PrivateFieldCfg = CfgWithPrivateFields{
		host:       "127.0.0.1",
		port:       80,
		DBUser:     "root",
		DBPassword: "123456",
		dbName:     "test",
	}

	EmptyValuesCfg = CfgWithEmptyValues{
		Host:       "",
		Port:       0,
		DBUser:     "root",
		DBPassword: "123456",
		DBName:     "",
	}

	NestedStructCfg = ConfigWithNestedStruct{
		CfgWithCorrectTags: CorrectCfg,
		DBName:             "test",
	}
)
