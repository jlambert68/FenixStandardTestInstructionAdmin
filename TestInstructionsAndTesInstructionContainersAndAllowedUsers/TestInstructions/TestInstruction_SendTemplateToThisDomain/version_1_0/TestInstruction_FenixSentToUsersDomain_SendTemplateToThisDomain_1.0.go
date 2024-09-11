package version_1_0

import (
	"github.com/jlambert68/FenixStandardTestInstructionAdmin/LocalExecutionMethods"
	"github.com/jlambert68/FenixStandardTestInstructionAdmin/LocalExecutionMethods/TestApiEngineClassesAndMethods"
	"github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/DomainData"
	"github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions"
	fixedValuesOverVersions "github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendTemplateToThisDomain"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/Domains"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestCaseModelElementTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
)

const (

	// *************************************
	// *** TestInstruction *** 'FenixOwnedSendTemplateToThisDomain'
	TestInstructionUUID_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain               TypeAndStructs.OriginalElementUUIDType = fixedValuesOverVersions.TestInstructionUUID_FenixSentToUsersDomain_SendTemplateToThisDomain
	TestInstructionName_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain               TypeAndStructs.TestInstructionNameType = "FenixOwnedSendTemplateToThisDomain"
	TestInstructionTypeUUID_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain                                                  = TestInstructions.TestInstructionTypeUUID_FenixSentToUsersDomain_SendTemplate
	TestInstructionTypeName_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain                                                  = TestInstructions.TestInstructionTypeName_FenixSentToUsersDomain_SendTemplate
	TestInstructionDescription_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain        string                                 = "This TestInstruction sends the used Template for the execution to ExecutionDomain"
	TestInstructionMouseOverText_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain      string                                 = "This TestInstruction sends the used Template for the execution to ExecutionDomain"
	TestInstructionDeprecated_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain         bool                                   = false
	TestInstructionEnabled_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain            bool                                   = true
	TestInstructionMajorVersionNumber_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain int                                    = 1
	TestInstructionMinorVersionNumber_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain int                                    = 0
	TestInstructionColor_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain              TypeAndStructs.ColorType               = "#fff000AA"
	TCRuleDeletion_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain                    TypeAndStructs.TCRuleDeletionType      = "TCRuleDeletion020"
	TCRuleSwap_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain                        TypeAndStructs.TCRuleSwapType          = "TCRuleSwap020"
	TestInstructionCreatingTimeStamp                                                            TypeAndStructs.UpdatedTimeStampType    = "2024-08-28 13:00:00"
	ExpectedMaxTestInstructionExecutionDurationInSeconds                                        int64                                  = 30

	// Attribute - 'SendTemplateToThisExecutionDomain'
	TestInstructionAttributeUUID_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain_SendTemplateToThisExecutionDomain          TypeAndStructs.TestInstructionAttributeUUIDType = "5de20bcf-701c-4f6f-bd96-ba6f4d721a69" // TODO fix so they use the same UUID, Can't bu done now because UUID is key in Attrubutes-table in DB .TestInstructionAttributeUUID_FenixSentToUsersDomain_SendTemplateToThisExecutionDomain
	TestInstructionAttributeName_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain_SendTemplateToThisExecutionDomain          TypeAndStructs.TestInstructionAttributeNameType = TestInstructions.TestInstructionAttributeName_FenixSentToUsersDomain_SendTemplateToThisDomain
	TestInstructionAttributeType_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain_SendTemplateToThisExecutionDomain          TypeAndStructs.TestInstructionAttributeTypeType = TestInstructions.TestInstructionAttributeType_FenixSentToUsersDomain_SendTemplateToThisDomain
	TestInstructionAttributeDescription_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain_SendTemplateToThisExecutionDomain   string                                          = "The Uuid of the ExecutionDomain that should receive the Template"
	TestInstructionAttributeMouseOverText_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain_SendTemplateToThisExecutionDomain string                                          = "The Uuid of the ExecutionDomain that should receive the Template"

	// Attribute - 'FenixOwnedSendTemplateToThisDomainTextBox'
	TestInstructionAttributeUUID_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateToThisDomainTextBox          TypeAndStructs.TestInstructionAttributeUUIDType = "39789a8a-7e5a-4c6a-a589-8a23327c5239" //TestInstructionAttributeUUID_SubCustody_FenixOwnedSendTemplateToThisDomainTextBox
	TestInstructionAttributeName_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateToThisDomainTextBox          TypeAndStructs.TestInstructionAttributeNameType = "FenixOwnedSendTemplateToThisDomainTextBox"
	TestInstructionAttributeType_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateToThisDomainTextBox          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeDescription_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateToThisDomainTextBox   string                                          = "The Domain the will receive the Template"
	TestInstructionAttributeMouseOverText_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateToThisDomainTextBox string                                          = "The Domain the will receive the Template"

	// Attribute - 'SendTemplateToThisExecutionDomainTextBox'
	TestInstructionAttributeUUID_FenixOwnedSendTemplateToThisDomain_SendTemplateToThisExecutionDomainTextBox          TypeAndStructs.TestInstructionAttributeUUIDType = "881b1094-8742-43e1-a3aa-7352061a9a61" //TestInstructionAttributeUUID_SubCustody_SendTemplateToThisExecutionDomainTextBox
	TestInstructionAttributeName_FenixOwnedSendTemplateToThisDomain_SendTemplateToThisExecutionDomainTextBox          TypeAndStructs.TestInstructionAttributeNameType = "SendTemplateToThisExecutionDomainTextBox"
	TestInstructionAttributeType_FenixOwnedSendTemplateToThisDomain_SendTemplateToThisExecutionDomainTextBox          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeDescription_FenixOwnedSendTemplateToThisDomain_SendTemplateToThisExecutionDomainTextBox   string                                          = "The ExecutionDomain the will receive the Template"
	TestInstructionAttributeMouseOverText_FenixOwnedSendTemplateToThisDomain_SendTemplateToThisExecutionDomainTextBox string                                          = "The ExecutionDomain the will receive the Template"

	// Attribute - 'TemplateAsString'
	TestInstructionAttributeUUID_FenixOwnedSendTemplateToThisDomain_TemplateAsString          TypeAndStructs.TestInstructionAttributeUUIDType = "53c5960b-e0d4-4535-b333-7a9f6681db15" //TestInstructionAttributeUUID_SubCustody_TemplateAsString
	TestInstructionAttributeName_FenixOwnedSendTemplateToThisDomain_TemplateAsString          TypeAndStructs.TestInstructionAttributeNameType = "TemplateAsString"
	TestInstructionAttributeType_FenixOwnedSendTemplateToThisDomain_TemplateAsString          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeDescription_FenixOwnedSendTemplateToThisDomain_TemplateAsString   string                                          = "Template as a json string"
	TestInstructionAttributeMouseOverText_FenixOwnedSendTemplateToThisDomain_TemplateAsString string                                          = "Template as a json string"
)

var (
	// Attribute - 'SendTemplateToThisExecutionDomain'
	TestInstructionAttributeComboBoxPredefinedValues_FenixOwnedSendTemplateToThisDomain_SendTemplateToThisExecutionDomain TypeAndStructs.TestInstructionAttributeComboBoxPredefinedValuesType
)

// TestInstruction_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain
// Variable that holds the data for the TestInstruction
var TestInstruction_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct

// Initate_TestInstruction_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain
// Function that creates all data for the TestInstruction
func Initate_TestInstruction_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain() *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct {

	// Initiate variable to be able to store all TestInstruction data
	TestInstruction_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct{
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

	// TestInstruction - FenixOwnedSendTemplateToThisDomain
	TestInstruction_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain.TestInstruction = &TypeAndStructs.TestInstructionStruct{
		DomainUUID:                   DomainData.DomainUUID_FenixStandard,
		DomainName:                   DomainData.DomainName_FenixStandard,
		ExecutionDomainUUID:          DomainData.ExecutionDomainUUID_FenixSentToUsersDomain,
		ExecutionDomainName:          DomainData.ExecutionDomainName_FenixSentToUsersDomain,
		TestInstructionUUID:          TestInstructionUUID_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
		TestInstructionName:          TestInstructionName_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
		TestInstructionTypeName:      TestInstructionTypeName_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
		TestInstructionDescription:   TestInstructionDescription_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
		TestInstructionMouseOverText: TestInstructionMouseOverText_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
		Deprecated:                   TestInstructionDeprecated_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
		Enabled:                      TestInstructionEnabled_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
		UpdatedTimeStamp:             TestInstructionCreatingTimeStamp,
	}

	// BasicTestInstructionInformation - FenixOwnedSendTemplateToThisDomain
	TestInstruction_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain.BasicTestInstructionInformation = &TypeAndStructs.BasicTestInstructionInformationStruct{
		DomainUUID:                   DomainData.DomainUUID_FenixStandard,
		DomainName:                   DomainData.DomainName_FenixStandard,
		ExecutionDomainUUID:          DomainData.ExecutionDomainUUID_FenixSentToUsersDomain,
		ExecutionDomainName:          DomainData.ExecutionDomainName_FenixSentToUsersDomain,
		TestInstructionUUID:          TestInstructionUUID_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
		TestInstructionName:          TestInstructionName_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
		TestInstructionTypeName:      TestInstructionTypeName_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
		Deprecated:                   TestInstructionDeprecated_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
		UpdatedTimeStamp:             TestInstructionCreatingTimeStamp,
		TestInstructionColor:         TestInstructionColor_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
		TCRuleDeletion:               TCRuleDeletion_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
		TCRuleSwap:                   TCRuleSwap_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
		TestInstructionDescription:   TestInstructionDescription_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
		TestInstructionMouseOverText: TestInstructionMouseOverText_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
		Enabled:                      TestInstructionEnabled_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
	}

	// TestInstruction Attribute - 'SendTemplateToThisExecutionDomain'
	var TestInstructionAttribute_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain_SendTemplateToThisExecutionDomain *TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain_SendTemplateToThisExecutionDomain = &TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                       DomainData.DomainUUID_FenixStandard,
		DomainName:                                       DomainData.DomainName_FenixStandard,
		TestInstructionUUID:                              TestInstructionUUID_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
		TestInstructionName:                              TestInstructionName_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
		TestInstructionAttributeUUID:                     TestInstructionAttributeUUID_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain_SendTemplateToThisExecutionDomain,
		TestInstructionAttributeName:                     TestInstructionAttributeName_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain_SendTemplateToThisExecutionDomain,
		TestInstructionAttributeDescription:              TestInstructionAttributeDescription_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain_SendTemplateToThisExecutionDomain,
		TestInstructionAttributeMouseOver:                TestInstructionAttributeMouseOverText_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain_SendTemplateToThisExecutionDomain,
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
		TestInstructionAttributeType:                     TestInstructionAttributeType_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain_SendTemplateToThisExecutionDomain,
		TestInstructionAttributeComboBoxPredefinedValues: TestInstructionAttributeComboBoxPredefinedValues_FenixOwnedSendTemplateToThisDomain_SendTemplateToThisExecutionDomain,
	}
	TestInstruction_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain.TestInstructionAttribute = append(
		TestInstruction_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain.TestInstructionAttribute,
		TestInstructionAttribute_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain_SendTemplateToThisExecutionDomain)

	// Add TestApiEngine relation for Attribute - 'SendTemplateToThisExecutionDomain'
	// Nothing here

	// TestInstruction Attribute - 'FenixOwnedSendTemplateToThisDomainTextBox'
	var TestInstructionAttribute_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomainTextBox *TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomainTextBox = &TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                       DomainData.DomainUUID_FenixStandard,
		DomainName:                                       DomainData.DomainName_FenixStandard,
		TestInstructionUUID:                              TestInstructionUUID_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
		TestInstructionName:                              TestInstructionName_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
		TestInstructionAttributeUUID:                     TestInstructionAttributeUUID_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateToThisDomainTextBox,
		TestInstructionAttributeName:                     TestInstructionAttributeName_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateToThisDomainTextBox,
		TestInstructionAttributeDescription:              TestInstructionAttributeDescription_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateToThisDomainTextBox,
		TestInstructionAttributeMouseOver:                TestInstructionAttributeMouseOverText_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateToThisDomainTextBox,
		TestInstructionAttributeTypeUUID:                 TestInstructions.TestInstructionAttributeTypeUUID_FenixStandard_Standard_Hidden,
		TestInstructionAttributeTypeName:                 TestInstructions.TestInstructionAttributeTypeName_FenixStandard_Standard_Hidden,
		TestInstructionAttributeValueAsString:            Domains.TestInstructionAttributeValueAsStringValue_NO_VALUE,
		TestInstructionAttributeValueUUID:                Domains.TestInstructionAttributeValueUUID_NO_VALUE,
		TestInstructionAttributeVisible:                  true,
		TestInstructionAttributeEnabled:                  false,
		TestInstructionAttributeMandatory:                false,
		TestInstructionAttributeVisibleInTestCaseArea:    false,
		TestInstructionAttributeIsDeprecated:             false,
		TestInstructionAttributeInputMask:                ".*",
		TestInstructionAttributeType:                     TestInstructionAttributeType_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateToThisDomainTextBox,
		TestInstructionAttributeComboBoxPredefinedValues: nil,
	}
	TestInstruction_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain.TestInstructionAttribute = append(
		TestInstruction_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain.TestInstructionAttribute,
		TestInstructionAttribute_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomainTextBox)

	// Add TestApiEngine relation for Attribute - 'FenixOwnedSendTemplateToThisDomainTextBox'
	// Nothing here

	// TestInstruction Attribute - 'SendTemplateToThisExecutionDomainTextBox'
	var TestInstructionAttribute_FenixSentToUsersDomain_SendTemplateToThisExecutionDomainTextBox *TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_FenixSentToUsersDomain_SendTemplateToThisExecutionDomainTextBox = &TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                       DomainData.DomainUUID_FenixStandard,
		DomainName:                                       DomainData.DomainName_FenixStandard,
		TestInstructionUUID:                              TestInstructionUUID_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
		TestInstructionName:                              TestInstructionName_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
		TestInstructionAttributeUUID:                     TestInstructionAttributeUUID_FenixOwnedSendTemplateToThisDomain_SendTemplateToThisExecutionDomainTextBox,
		TestInstructionAttributeName:                     TestInstructionAttributeName_FenixOwnedSendTemplateToThisDomain_SendTemplateToThisExecutionDomainTextBox,
		TestInstructionAttributeDescription:              TestInstructionAttributeDescription_FenixOwnedSendTemplateToThisDomain_SendTemplateToThisExecutionDomainTextBox,
		TestInstructionAttributeMouseOver:                TestInstructionAttributeMouseOverText_FenixOwnedSendTemplateToThisDomain_SendTemplateToThisExecutionDomainTextBox,
		TestInstructionAttributeTypeUUID:                 TestInstructions.TestInstructionAttributeTypeUUID_FenixStandard_Standard_Hidden,
		TestInstructionAttributeTypeName:                 TestInstructions.TestInstructionAttributeTypeName_FenixStandard_Standard_Hidden,
		TestInstructionAttributeValueAsString:            Domains.TestInstructionAttributeValueAsStringValue_NO_VALUE,
		TestInstructionAttributeValueUUID:                Domains.TestInstructionAttributeValueUUID_NO_VALUE,
		TestInstructionAttributeVisible:                  true,
		TestInstructionAttributeEnabled:                  false,
		TestInstructionAttributeMandatory:                false,
		TestInstructionAttributeVisibleInTestCaseArea:    false,
		TestInstructionAttributeIsDeprecated:             false,
		TestInstructionAttributeInputMask:                ".*",
		TestInstructionAttributeType:                     TestInstructionAttributeType_FenixOwnedSendTemplateToThisDomain_SendTemplateToThisExecutionDomainTextBox,
		TestInstructionAttributeComboBoxPredefinedValues: nil,
	}
	TestInstruction_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain.TestInstructionAttribute = append(
		TestInstruction_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain.TestInstructionAttribute,
		TestInstructionAttribute_FenixSentToUsersDomain_SendTemplateToThisExecutionDomainTextBox)

	// Add TestApiEngine relation for Attribute - 'SendTemplateToThisExecutionDomainTextBox'
	// Nothing here

	// TestInstruction Attribute - 'TemplateAsString'
	var TestInstructionAttribute_FenixSentToUsersDomain_TemplateAsString *TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_FenixSentToUsersDomain_TemplateAsString = &TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                       DomainData.DomainUUID_FenixStandard,
		DomainName:                                       DomainData.DomainName_FenixStandard,
		TestInstructionUUID:                              TestInstructionUUID_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
		TestInstructionName:                              TestInstructionName_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
		TestInstructionAttributeUUID:                     TestInstructionAttributeUUID_FenixOwnedSendTemplateToThisDomain_TemplateAsString,
		TestInstructionAttributeName:                     TestInstructionAttributeName_FenixOwnedSendTemplateToThisDomain_TemplateAsString,
		TestInstructionAttributeDescription:              TestInstructionAttributeDescription_FenixOwnedSendTemplateToThisDomain_TemplateAsString,
		TestInstructionAttributeMouseOver:                TestInstructionAttributeMouseOverText_FenixOwnedSendTemplateToThisDomain_TemplateAsString,
		TestInstructionAttributeTypeUUID:                 TestInstructions.TestInstructionAttributeTypeUUID_FenixStandard_Standard_Hidden,
		TestInstructionAttributeTypeName:                 TestInstructions.TestInstructionAttributeTypeName_FenixStandard_Standard_Hidden,
		TestInstructionAttributeValueAsString:            Domains.TestInstructionAttributeValueAsStringValue_NO_VALUE,
		TestInstructionAttributeValueUUID:                Domains.TestInstructionAttributeValueUUID_NO_VALUE,
		TestInstructionAttributeVisible:                  true,
		TestInstructionAttributeEnabled:                  false,
		TestInstructionAttributeMandatory:                false,
		TestInstructionAttributeVisibleInTestCaseArea:    false,
		TestInstructionAttributeIsDeprecated:             false,
		TestInstructionAttributeInputMask:                ".*",
		TestInstructionAttributeType:                     TestInstructionAttributeType_FenixOwnedSendTemplateToThisDomain_TemplateAsString,
		TestInstructionAttributeComboBoxPredefinedValues: nil,
	}
	TestInstruction_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain.TestInstructionAttribute = append(
		TestInstruction_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain.TestInstructionAttribute,
		TestInstructionAttribute_FenixSentToUsersDomain_TemplateAsString)

	// Add TestApiEngine relation for Attribute - 'TemplateAsString'
	// Nothing here

	// ImmatureElementModel - FenixOwnedSendTemplateToThisDomain
	var TestInstructionImmatureElementModel_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain *TypeAndStructs.ImmatureElementModelMessageStruct
	TestInstructionImmatureElementModel_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain = &TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               DomainData.DomainUUID_FenixStandard,
		DomainName:               DomainData.DomainName_FenixStandard,
		ImmatureElementUUID:      TestInstructionUUID_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(TestInstructionName_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain),
		PreviousElementUUID:      TestInstructionUUID_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
		NextElementUUID:          TestInstructionUUID_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
		FirstChildElementUUID:    TestInstructionUUID_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
		ParentElementUUID:        TestInstructionUUID_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TI,
		OriginalElementUUID:      TestInstructionUUID_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
		TopImmatureElementUUID:   TestInstructionUUID_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain,
		IsTopElement:             true,
	}
	TestInstruction_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain.ImmatureElementModel = append(
		TestInstruction_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain.ImmatureElementModel,
		TestInstructionImmatureElementModel_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain)

	return TestInstruction_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain
}
