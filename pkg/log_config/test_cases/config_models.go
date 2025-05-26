package test_cases

type (
	CfgSuccessfulCompare struct {
		CfgWithCorrectTags CfgWithCorrectTags
		CfgWithCorrectTag2 CfgWithCorrectTags
	}

	CfgCompareWithPrivateFields struct {
		CfgWithPrivateFields  CfgWithPrivateFields
		cfgWithPrivateFields2 CfgWithPrivateFields
	}
)

var (
	SuccessfulCfg = CfgSuccessfulCompare{
		CfgWithCorrectTags: CorrectCfg,
		CfgWithCorrectTag2: CorrectCfg,
	}

	CompareWithPrivateFields = CfgCompareWithPrivateFields{
		CfgWithPrivateFields:  PrivateFieldCfg,
		cfgWithPrivateFields2: PrivateFieldCfg,
	}
)
