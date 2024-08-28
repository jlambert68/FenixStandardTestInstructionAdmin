package version_1_0

import (
	"github.com/jlambert68/FenixStandardTestInstructionAdmin/LocalExecutionMethods"
	"github.com/jlambert68/FenixStandardTestInstructionAdmin/LocalExecutionMethods/TestApiEngineClassesAndMethods"
	"github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/DomainData"
	"github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions"
	fixedValuesOverVersions "github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendTestDataToThisDomain"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/Domains"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestCaseModelElementTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
)

const (

	// *************************************
	// *** TestInstruction *** 'SendTestDataToThisDomain'
	TestInstructionUUID_FenixSentToUsersDomain_SendTestDataToThisDomain               TypeAndStructs.OriginalElementUUIDType = fixedValuesOverVersions.TestInstructionUUID_FenixSentToUsersDomain_SendTestDataToThisDomain
	TestInstructionName_FenixSentToUsersDomain_SendTestDataToThisDomain               TypeAndStructs.TestInstructionNameType = "SendTestDataToThisDomain"
	TestInstructionTypeUUID_FenixSentToUsersDomain_SendTestDataToThisDomain                                                  = TestInstructions.TestInstructionTypeUUID_FenixSentToUsersDomain_SendTestDataUsedByExecution
	TestInstructionTypeName_FenixSentToUsersDomain_SendTestDataToThisDomain                                                  = TestInstructions.TestInstructionTypeName_FenixSentToUsersDomain_SendTestDataUsedByExecution
	TestInstructionDescription_FenixSentToUsersDomain_SendTestDataToThisDomain        string                                 = "This TestInstruction sends the used TestData for the execution to ExecutionDomain"
	TestInstructionMouseOverText_FenixSentToUsersDomain_SendTestDataToThisDomain      string                                 = "This TestInstruction sends the used TestData for the execution to ExecutionDomain"
	TestInstructionDeprecated_FenixSentToUsersDomain_SendTestDataToThisDomain         bool                                   = false
	TestInstructionEnabled_FenixSentToUsersDomain_SendTestDataToThisDomain            bool                                   = true
	TestInstructionMajorVersionNumber_FenixSentToUsersDomain_SendTestDataToThisDomain int                                    = 1
	TestInstructionMinorVersionNumber_FenixSentToUsersDomain_SendTestDataToThisDomain int                                    = 0
	TestInstructionColor_FenixSentToUsersDomain_SendTestDataToThisDomain              TypeAndStructs.ColorType               = "#fff000AA"
	TCRuleDeletion_FenixSentToUsersDomain_SendTestDataToThisDomain                    TypeAndStructs.TCRuleDeletionType      = "TCRuleDeletion020"
	TCRuleSwap_FenixSentToUsersDomain_SendTestDataToThisDomain                        TypeAndStructs.TCRuleSwapType          = "TCRuleSwap020"
	TestInstructionCreatingTimeStamp                                                  TypeAndStructs.UpdatedTimeStampType    = "2024-08-28 13:00:00"
	ExpectedMaxTestInstructionExecutionDurationInSeconds                              int64                                  = 300

	// Attribute - 'SendTestDataToThisExecutionDomain'
	TestInstructionAttributeUUID_FenixSentToUsersDomain_SendTestDataToThisDomain_SendTestDataToThisExecutionDomain          TypeAndStructs.TestInstructionAttributeUUIDType = "47bce83a-3b71-439f-8283-33d0e26d62d9" // TODO fix so they use the same UUID, Can't bu done now because UUID is key in Attrubutes-table in DB .TestInstructionAttributeUUID_FenixSentToUsersDomain_SendTestDataToThisExecutionDomain
	TestInstructionAttributeName_FenixSentToUsersDomain_SendTestDataToThisDomain_SendTestDataToThisExecutionDomain          TypeAndStructs.TestInstructionAttributeNameType = TestInstructions.TestInstructionAttributeName_FenixSentToUsersDomain_SendTestDataToThisDomain
	TestInstructionAttributeType_FenixSentToUsersDomain_SendTestDataToThisDomain_SendTestDataToThisExecutionDomain          TypeAndStructs.TestInstructionAttributeTypeType = TestInstructions.TestInstructionAttributeType_FenixSentToUsersDomain_SendTestDataToThisDomain
	TestInstructionAttributeDescription_FenixSentToUsersDomain_SendTestDataToThisDomain_SendTestDataToThisExecutionDomain   string                                          = "The Uuid of the ExecutionDomain that should receive the TestData"
	TestInstructionAttributeMouseOverText_FenixSentToUsersDomain_SendTestDataToThisDomain_SendTestDataToThisExecutionDomain string                                          = "The Uuid of the ExecutionDomain that should receive the TestData"
)

var (
	// Attribute - 'SendTestDataToThisExecutionDomain'
	TestInstructionAttributeComboBoxPredefinedValues_SendTestDataToThisDomain_SendTestDataToThisExecutionDomain = TypeAndStructs.TestInstructionAttributeComboBoxPredefinedValuesType{}
)

// TestInstruction_FenixSentToUsersDomain_SendTestDataToThisDomain
// Variable that holds the data for the TestInstruction
var TestInstruction_FenixSentToUsersDomain_SendTestDataToThisDomain *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct

// Initate_TestInstruction_FenixSentToUsersDomain_SendTestDataToThisDomain
// Function that creates all data for the TestInstruction
func Initate_TestInstruction_FenixSentToUsersDomain_SendTestDataToThisDomain() *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct {

	// Initiate variable to be able to store all TestInstruction data
	TestInstruction_FenixSentToUsersDomain_SendTestDataToThisDomain = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct{
		TestInstruction:                    &TypeAndStructs.TestInstructionStruct{},
		BasicTestInstructionInformation:    &TypeAndStructs.BasicTestInstructionInformationStruct{},
		ImmatureTestInstructionInformation: nil,
		TestInstructionAttribute:           nil,
		ImmatureElementModel:               nil,

		// Local Execution Methods are specified here
		LocalExecutionMethods: TestInstructionAndTestInstuctionContainerTypes.AnyType{
			&LocalExecutionMethods.MethodsForLocalExecutionsStruct{
				LocalParametersUsedInRunTime: &LocalExecutionMethods.LocalParametersUsedInRunTimeStruct{
					ExpectedTestInstructionExecutionDurationInSeconds: ExpectedMaxTestInstructionExecutionDurationInSeconds,
				},
				TestApiEngineClassesMethodsAttributes: &TestApiEngineClassesAndMethods.TestApiEngineClassesMethodsAttributesStruct{},
			},
		},
	}

	// TestInstruction - SendTestDataToThisDomain
	TestInstruction_FenixSentToUsersDomain_SendTestDataToThisDomain.TestInstruction = &TypeAndStructs.TestInstructionStruct{
		DomainUUID:                   DomainData.DomainUUID_FenixStandard,
		DomainName:                   DomainData.DomainName_FenixStandard,
		ExecutionDomainUUID:          DomainData.ExecutionDomainUUID_FenixStandard,
		ExecutionDomainName:          DomainData.ExecutionDomainName_FenixStandard,
		TestInstructionUUID:          TestInstructionUUID_FenixSentToUsersDomain_SendTestDataToThisDomain,
		TestInstructionName:          TestInstructionName_FenixSentToUsersDomain_SendTestDataToThisDomain,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_FenixSentToUsersDomain_SendTestDataToThisDomain,
		TestInstructionTypeName:      TestInstructionTypeName_FenixSentToUsersDomain_SendTestDataToThisDomain,
		TestInstructionDescription:   TestInstructionDescription_FenixSentToUsersDomain_SendTestDataToThisDomain,
		TestInstructionMouseOverText: TestInstructionMouseOverText_FenixSentToUsersDomain_SendTestDataToThisDomain,
		Deprecated:                   TestInstructionDeprecated_FenixSentToUsersDomain_SendTestDataToThisDomain,
		Enabled:                      TestInstructionEnabled_FenixSentToUsersDomain_SendTestDataToThisDomain,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_FenixSentToUsersDomain_SendTestDataToThisDomain,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_FenixSentToUsersDomain_SendTestDataToThisDomain,
		UpdatedTimeStamp:             TestInstructionCreatingTimeStamp,
	}

	// BasicTestInstructionInformation - SendTestDataToThisDomain
	TestInstruction_FenixSentToUsersDomain_SendTestDataToThisDomain.BasicTestInstructionInformation = &TypeAndStructs.BasicTestInstructionInformationStruct{
		DomainUUID:                   DomainData.DomainUUID_FenixStandard,
		DomainName:                   DomainData.DomainName_FenixStandard,
		ExecutionDomainUUID:          DomainData.ExecutionDomainUUID_FenixStandard,
		ExecutionDomainName:          DomainData.ExecutionDomainName_FenixStandard,
		TestInstructionUUID:          TestInstructionUUID_FenixSentToUsersDomain_SendTestDataToThisDomain,
		TestInstructionName:          TestInstructionName_FenixSentToUsersDomain_SendTestDataToThisDomain,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_FenixSentToUsersDomain_SendTestDataToThisDomain,
		TestInstructionTypeName:      TestInstructionTypeName_FenixSentToUsersDomain_SendTestDataToThisDomain,
		Deprecated:                   TestInstructionDeprecated_FenixSentToUsersDomain_SendTestDataToThisDomain,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_FenixSentToUsersDomain_SendTestDataToThisDomain,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_FenixSentToUsersDomain_SendTestDataToThisDomain,
		UpdatedTimeStamp:             TestInstructionCreatingTimeStamp,
		TestInstructionColor:         TestInstructionColor_FenixSentToUsersDomain_SendTestDataToThisDomain,
		TCRuleDeletion:               TCRuleDeletion_FenixSentToUsersDomain_SendTestDataToThisDomain,
		TCRuleSwap:                   TCRuleSwap_FenixSentToUsersDomain_SendTestDataToThisDomain,
		TestInstructionDescription:   TestInstructionDescription_FenixSentToUsersDomain_SendTestDataToThisDomain,
		TestInstructionMouseOverText: TestInstructionMouseOverText_FenixSentToUsersDomain_SendTestDataToThisDomain,
		Enabled:                      TestInstructionEnabled_FenixSentToUsersDomain_SendTestDataToThisDomain,
	}

	// TestInstruction Attribute - 'SendTestDataToThisExecutionDomain'
	var TestInstructionAttribute_FenixSentToUsersDomain_SendTestDataToThisDomain_SendTestDataToThisExecutionDomain *TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_FenixSentToUsersDomain_SendTestDataToThisDomain_SendTestDataToThisExecutionDomain = &TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                       DomainData.DomainUUID_FenixStandard,
		DomainName:                                       DomainData.DomainName_FenixStandard,
		TestInstructionUUID:                              TestInstructionUUID_FenixSentToUsersDomain_SendTestDataToThisDomain,
		TestInstructionName:                              TestInstructionName_FenixSentToUsersDomain_SendTestDataToThisDomain,
		TestInstructionAttributeUUID:                     TestInstructionAttributeUUID_FenixSentToUsersDomain_SendTestDataToThisDomain_SendTestDataToThisExecutionDomain,
		TestInstructionAttributeName:                     TestInstructionAttributeName_FenixSentToUsersDomain_SendTestDataToThisDomain_SendTestDataToThisExecutionDomain,
		TestInstructionAttributeDescription:              TestInstructionAttributeDescription_FenixSentToUsersDomain_SendTestDataToThisDomain_SendTestDataToThisExecutionDomain,
		TestInstructionAttributeMouseOver:                TestInstructionAttributeMouseOverText_FenixSentToUsersDomain_SendTestDataToThisDomain_SendTestDataToThisExecutionDomain,
		TestInstructionAttributeTypeUUID:                 TestInstructions.TestInstructionAttributeTypeUUID_FenixStandard_Standard,
		TestInstructionAttributeTypeName:                 TestInstructions.TestInstructionAttributeTypeName_FenixStandard_Standard,
		TestInstructionAttributeValueAsString:            Domains.TestInstructionAttributeValueAsStringValue_NO_VALUE,
		TestInstructionAttributeValueUUID:                Domains.TestInstructionAttributeValueUUID_NO_VALUE,
		TestInstructionAttributeVisible:                  true,
		TestInstructionAttributeEnabled:                  true,
		TestInstructionAttributeMandatory:                true,
		TestInstructionAttributeVisibleInTestCaseArea:    false,
		TestInstructionAttributeIsDeprecated:             false,
		TestInstructionAttributeInputMask:                "^.+$",
		TestInstructionAttributeType:                     TestInstructionAttributeType_FenixSentToUsersDomain_SendTestDataToThisDomain_SendTestDataToThisExecutionDomain,
		TestInstructionAttributeComboBoxPredefinedValues: TestInstructionAttributeComboBoxPredefinedValues_SendTestDataToThisDomain_SendTestDataToThisExecutionDomain,
	}
	TestInstruction_FenixSentToUsersDomain_SendTestDataToThisDomain.TestInstructionAttribute = append(
		TestInstruction_FenixSentToUsersDomain_SendTestDataToThisDomain.TestInstructionAttribute,
		TestInstructionAttribute_FenixSentToUsersDomain_SendTestDataToThisDomain_SendTestDataToThisExecutionDomain)

	// Add TestApiEngine relation for Attribute - 'SendTestDataToThisExecutionDomain'
	// Nothing here

	// ImmatureElementModel - SendTestDataToThisDomain
	var TestInstructionImmatureElementModel_FenixSentToUsersDomain_SendTestDataToThisDomain *TypeAndStructs.ImmatureElementModelMessageStruct
	TestInstructionImmatureElementModel_FenixSentToUsersDomain_SendTestDataToThisDomain = &TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               DomainData.DomainUUID_FenixStandard,
		DomainName:               DomainData.DomainName_FenixStandard,
		ImmatureElementUUID:      TestInstructionUUID_FenixSentToUsersDomain_SendTestDataToThisDomain,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(TestInstructionName_FenixSentToUsersDomain_SendTestDataToThisDomain),
		PreviousElementUUID:      TestInstructionUUID_FenixSentToUsersDomain_SendTestDataToThisDomain,
		NextElementUUID:          TestInstructionUUID_FenixSentToUsersDomain_SendTestDataToThisDomain,
		FirstChildElementUUID:    TestInstructionUUID_FenixSentToUsersDomain_SendTestDataToThisDomain,
		ParentElementUUID:        TestInstructionUUID_FenixSentToUsersDomain_SendTestDataToThisDomain,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TI,
		OriginalElementUUID:      TestInstructionUUID_FenixSentToUsersDomain_SendTestDataToThisDomain,
		TopImmatureElementUUID:   TestInstructionUUID_FenixSentToUsersDomain_SendTestDataToThisDomain,
		IsTopElement:             true,
	}
	TestInstruction_FenixSentToUsersDomain_SendTestDataToThisDomain.ImmatureElementModel = append(
		TestInstruction_FenixSentToUsersDomain_SendTestDataToThisDomain.ImmatureElementModel,
		TestInstructionImmatureElementModel_FenixSentToUsersDomain_SendTestDataToThisDomain)

	return TestInstruction_FenixSentToUsersDomain_SendTestDataToThisDomain
}
