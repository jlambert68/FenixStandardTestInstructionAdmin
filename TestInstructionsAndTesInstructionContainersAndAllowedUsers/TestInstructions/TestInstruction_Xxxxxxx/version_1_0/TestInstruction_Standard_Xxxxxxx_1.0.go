package version_1_0

import (
	"FenixStandardTestInstructionAdmin/LocalExecutionMethods"
	"FenixStandardTestInstructionAdmin/LocalExecutionMethods/FangEngineClassesAndMethods"
	"FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/DomainData"
	"FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions"
	fixedValuesOverVersions "FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_Xxxxxxx"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/Domains"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestCaseModelElementTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
)

const (

	// *************************************
	// *** TestInstruction *** 'Xxxxxxx'
	TestInstructionUUID_FenixStandard_Xxxxxxx               TypeAndStructs.OriginalElementUUIDType = fixedValuesOverVersions.TestInstructionUUID_FenixStandard_Xxxxxxx
	TestInstructionName_FenixStandard_Xxxxxxx               TypeAndStructs.TestInstructionNameType = "Xxxxxxx"
	TestInstructionTypeUUID_FenixStandard_Xxxxxxx                                                  = TestInstructions.TestInstructionTypeUUID_FenixStandard_Standard
	TestInstructionTypeName_FenixStandard_Xxxxxxx                                                  = TestInstructions.TestInstructionTypeName_FenixStandard_Standard
	TestInstructionDescription_FenixStandard_Xxxxxxx        string                                 = "This TestInstruction Checks if server 'X' is alive and responding "
	TestInstructionMouseOverText_FenixStandard_Xxxxxxx      string                                 = "This TestInstruction Checks if server 'X' is alive and responding"
	TestInstructionDeprecated_FenixStandard_Xxxxxxx         bool                                   = false
	TestInstructionEnabled_FenixStandard_Xxxxxxx            bool                                   = true
	TestInstructionMajorVersionNumber_FenixStandard_Xxxxxxx int                                    = 1
	TestInstructionMinorVersionNumber_FenixStandard_Xxxxxxx int                                    = 0
	TestInstructionColor_FenixStandard_Xxxxxxx              TypeAndStructs.ColorType               = "#fff000AA"
	TCRuleDeletion_FenixStandard_Xxxxxxx                    TypeAndStructs.TCRuleDeletionType      = "TCRuleDeletion020"
	TCRuleSwap_FenixStandard_Xxxxxxx                        TypeAndStructs.TCRuleSwapType          = "TCRuleSwap020"
	TestInstructionCreatingTimeStamp                        TypeAndStructs.UpdatedTimeStampType    = "2024-01-15 20:00:00"

	// *** DropZone *** 'Xxxxxxx_ExpectsToSucceed'
	TestInstructionDropZoneUUID_FenixStandard_Xxxxxxx_ExpectsToSucceed        TypeAndStructs.DropZoneUUIDType = "6344f79a-15a1-476e-8599-067088f6b73a"
	TestInstructionDropZoneName_FenixStandard_Xxxxxxx_ExpectsToSucceed        TypeAndStructs.DropZoneNameType = "Xxxxxxx_ExpectsToSucceed"
	TestInstructionDropZoneDescription_FenixStandard_Xxxxxxx_ExpectsToSucceed string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneMouseOver_FenixStandard_Xxxxxxx_ExpectsToSucceed   string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneColor_FenixStandard_Xxxxxxx_ExpectsToSucceed       TypeAndStructs.ColorType        = "#ffff00AA"

	// Attribute - 'ExpectedToBePassed'
	TestInstructionAttributeUUID_FenixStandard_Xxxxxxx_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeUUIDType = TestInstructions.TestInstructionAttributeUUID_FenixStandard_ExpectedToBePassed // TODO fix so they use the same UUID, Can't bu done now because UUID is key in Attrubutes-table in DB .TestInstructionAttributeUUID_FenixStandard_ExpectedToBePassed
	TestInstructionAttributeName_FenixStandard_Xxxxxxx_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeNameType = TestInstructions.TestInstructionAttributeName_FenixStandard_ExpectedToBePassed
	TestInstructionAttributeType_FenixStandard_Xxxxxxx_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeTypeType = TestInstructions.TestInstructionAttributeType_FenixStandard_ExpectedToBePassed
	TestInstructionAttributeActionCommand_FenixStandard_Xxxxxxx_ExpectedToBePassed      TypeAndStructs.AttributeActionCommandType       = Domains.AttributeActionCommand_USE_DROPZONE_VALUE_FOR_ATTRIBUTE
	TestInstructionAttributeValueAsStringValue_FenixStandard_Xxxxxxx_ExpectedToBePassed TypeAndStructs.AttributeValueAsStringType       = Domains.TestInstructionAttributeValueAsStringValue_TRUE
	TestInstructionAttributeValueUUID_FenixStandard_Xxxxxxx_ExpectedToBePassed          TypeAndStructs.AttributeValueUUIDType           = Domains.TestInstructionAttributeValueUUID_TRUE
	TestInstructionAttributeDescription_FenixStandard_Xxxxxxx_ExpectedToBePassed        string                                          = "Should the TestInstruction execution to be expected to succeed or not"
	TestInstructionAttributeMouseOverText_FenixStandard_Xxxxxxx_ExpectedToBePassed      string                                          = "Should the TestInstruction execution to be expected to succeed or not"

	// Attribute - 'ServerServerIpAddress'
	TestInstructionAttributeUUID_FenixStandard_Xxxxxxx_ServerIpAddress          TypeAndStructs.TestInstructionAttributeUUIDType = "7cd1932e-1842-48ee-b4f7-c5782d6e7453" // TODO fix so they use the same UUID, Can't bu done now because UUID is key in Attrubutes-table in DB .TestInstructionAttributeUUID_FenixStandard_ServerIpAddress
	TestInstructionAttributeName_FenixStandard_Xxxxxxx_ServerIpAddress          TypeAndStructs.TestInstructionAttributeNameType = "The IP-address of the server"
	TestInstructionAttributeType_FenixStandard_Xxxxxxx_ServerIpAddress          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeDescription_FenixStandard_Xxxxxxx_ServerIpAddress   string                                          = "The IP-address to the server to be checked"
	TestInstructionAttributeMouseOverText_FenixStandard_Xxxxxxx_ServerIpAddress string                                          = "The IP-address to the server to be checked"
)

// TestInstruction_FenixStandard_Xxxxxxx
// Variable that holds the data for the TestInstruction
var TestInstruction_FenixStandard_Xxxxxxx *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct

// Initate_TestInstruction_FenixStandard_Xxxxxxx
// Function that creates all data for the TestInstruction
func Initate_TestInstruction_FenixStandard_Xxxxxxx() *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct {

	// Initiate variable to be able to store all TestInstruction data
	TestInstruction_FenixStandard_Xxxxxxx = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct{
		TestInstruction:                    &TypeAndStructs.TestInstructionStruct{},
		BasicTestInstructionInformation:    &TypeAndStructs.BasicTestInstructionInformationStruct{},
		ImmatureTestInstructionInformation: nil,
		TestInstructionAttribute:           nil,
		ImmatureElementModel:               nil,

		// Local Execution Methods are specified here
		LocalExecutionMethods: TestInstructionAndTestInstuctionContainerTypes.AnyType{
			&LocalExecutionMethods.MethodsForLocalExecutionsStruct{
				LocalParametersUsedInRunTime: &LocalExecutionMethods.LocalParametersUsedInRunTimeStruct{
					ExpectedTestInstructionExecutionDurationInSeconds: 360,
				},
				FangEngineClassesMethodsAttributes: &FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct{},
			},
		},
	}

	// TestInstruction - Xxxxxxx
	TestInstruction_FenixStandard_Xxxxxxx.TestInstruction = &TypeAndStructs.TestInstructionStruct{
		DomainUUID:                   DomainData.DomainUUID_FenixStandard,
		DomainName:                   DomainData.DomainName_FenixStandard,
		ExecutionDomainUUID:          DomainData.ExecutionDomainUUID_FenixStandard,
		ExecutionDomainName:          DomainData.ExecutionDomainName_FenixStandard,
		TestInstructionUUID:          TestInstructionUUID_FenixStandard_Xxxxxxx,
		TestInstructionName:          TestInstructionName_FenixStandard_Xxxxxxx,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_FenixStandard_Xxxxxxx,
		TestInstructionTypeName:      TestInstructionTypeName_FenixStandard_Xxxxxxx,
		TestInstructionDescription:   TestInstructionDescription_FenixStandard_Xxxxxxx,
		TestInstructionMouseOverText: TestInstructionMouseOverText_FenixStandard_Xxxxxxx,
		Deprecated:                   TestInstructionDeprecated_FenixStandard_Xxxxxxx,
		Enabled:                      TestInstructionEnabled_FenixStandard_Xxxxxxx,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_FenixStandard_Xxxxxxx,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_FenixStandard_Xxxxxxx,
		UpdatedTimeStamp:             TestInstructionCreatingTimeStamp,
	}

	// BasicTestInstructionInformation - Xxxxxxx
	TestInstruction_FenixStandard_Xxxxxxx.BasicTestInstructionInformation = &TypeAndStructs.BasicTestInstructionInformationStruct{
		DomainUUID:                   DomainData.DomainUUID_FenixStandard,
		DomainName:                   DomainData.DomainName_FenixStandard,
		ExecutionDomainUUID:          DomainData.ExecutionDomainUUID_FenixStandard,
		ExecutionDomainName:          DomainData.ExecutionDomainName_FenixStandard,
		TestInstructionUUID:          TestInstructionUUID_FenixStandard_Xxxxxxx,
		TestInstructionName:          TestInstructionName_FenixStandard_Xxxxxxx,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_FenixStandard_Xxxxxxx,
		TestInstructionTypeName:      TestInstructionTypeName_FenixStandard_Xxxxxxx,
		Deprecated:                   TestInstructionDeprecated_FenixStandard_Xxxxxxx,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_FenixStandard_Xxxxxxx,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_FenixStandard_Xxxxxxx,
		UpdatedTimeStamp:             TestInstructionCreatingTimeStamp,
		TestInstructionColor:         TestInstructionColor_FenixStandard_Xxxxxxx,
		TCRuleDeletion:               TCRuleDeletion_FenixStandard_Xxxxxxx,
		TCRuleSwap:                   TCRuleSwap_FenixStandard_Xxxxxxx,
		TestInstructionDescription:   TestInstructionDescription_FenixStandard_Xxxxxxx,
		TestInstructionMouseOverText: TestInstructionMouseOverText_FenixStandard_Xxxxxxx,
		Enabled:                      TestInstructionEnabled_FenixStandard_Xxxxxxx,
	}

	// DropZone 'Xxxxxxx_ExpectsToSucceed'
	// ImmatureTestInstructionInformation  - DropZone: Xxxxxxx_ExpectsToSucceed, Attr: ExpectedToBePassed
	var TestInstruction_FenixStandard_Xxxxxxx_ExpectedToBePassed *TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstruction_FenixStandard_Xxxxxxx_ExpectedToBePassed = &TypeAndStructs.ImmatureTestInstructionInformationStruct{
		DomainUUID:                   DomainData.DomainUUID_FenixStandard,
		DomainName:                   DomainData.DomainName_FenixStandard,
		TestInstructionUUID:          TestInstructionUUID_FenixStandard_Xxxxxxx,
		TestInstructionName:          TestInstructionName_FenixStandard_Xxxxxxx,
		DropZoneUUID:                 TestInstructionDropZoneUUID_FenixStandard_Xxxxxxx_ExpectsToSucceed,
		DropZoneName:                 TestInstructionDropZoneName_FenixStandard_Xxxxxxx_ExpectsToSucceed,
		DropZoneDescription:          TestInstructionDropZoneDescription_FenixStandard_Xxxxxxx_ExpectsToSucceed,
		DropZoneMouseOver:            TestInstructionDropZoneMouseOver_FenixStandard_Xxxxxxx_ExpectsToSucceed,
		DropZoneColor:                TestInstructionDropZoneColor_FenixStandard_Xxxxxxx_ExpectsToSucceed,
		TestInstructionAttributeType: TestInstructionAttributeType_FenixStandard_Xxxxxxx_ExpectedToBePassed,
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_FenixStandard_Xxxxxxx_ExpectedToBePassed,
		TestInstructionAttributeName: TestInstructionAttributeName_FenixStandard_Xxxxxxx_ExpectedToBePassed,
		AttributeValueAsString:       TestInstructionAttributeValueAsStringValue_FenixStandard_Xxxxxxx_ExpectedToBePassed,
		AttributeValueUUID:           TestInstructionAttributeValueUUID_FenixStandard_Xxxxxxx_ExpectedToBePassed,
		FirstImmatureElementUUID:     TestInstructionUUID_FenixStandard_Xxxxxxx,
		AttributeActionCommand:       TestInstructionAttributeActionCommand_FenixStandard_Xxxxxxx_ExpectedToBePassed,
	}
	TestInstruction_FenixStandard_Xxxxxxx.ImmatureTestInstructionInformation = append(
		TestInstruction_FenixStandard_Xxxxxxx.ImmatureTestInstructionInformation,
		TestInstruction_FenixStandard_Xxxxxxx_ExpectedToBePassed)

	// TestInstruction Attribute - 'ExpectedToBePassed'
	var TestInstructionAttribute_FenixStandard_Xxxxxxx_ExpectedToBePassed *TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_FenixStandard_Xxxxxxx_ExpectedToBePassed = &TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    DomainData.DomainUUID_FenixStandard,
		DomainName:                                    DomainData.DomainName_FenixStandard,
		TestInstructionUUID:                           TestInstructionUUID_FenixStandard_Xxxxxxx,
		TestInstructionName:                           TestInstructionName_FenixStandard_Xxxxxxx,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_FenixStandard_Xxxxxxx_ExpectedToBePassed,
		TestInstructionAttributeName:                  TestInstructionAttributeName_FenixStandard_Xxxxxxx_ExpectedToBePassed,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_FenixStandard_Xxxxxxx_ExpectedToBePassed,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_FenixStandard_Xxxxxxx_ExpectedToBePassed,
		TestInstructionAttributeTypeUUID:              TestInstructions.TestInstructionAttributeTypeUUID_FenixStandard_ExpectedToPass,
		TestInstructionAttributeTypeName:              TestInstructions.TestInstructionAttributeTypeName_FenixStandard_ExpectedToPass,
		TestInstructionAttributeValueAsString:         Domains.TestInstructionAttributeValueAsStringValue_NO_VALUE,
		TestInstructionAttributeValueUUID:             Domains.TestInstructionAttributeValueUUID_NO_VALUE,
		TestInstructionAttributeVisible:               true,
		TestInstructionAttributeEnabled:               true,
		TestInstructionAttributeMandatory:             true,
		TestInstructionAttributeVisibleInTestCaseArea: false,
		TestInstructionAttributeIsDeprecated:          false,
		TestInstructionAttributeInputMask:             ".",
		TestInstructionAttributeType:                  TestInstructionAttributeType_FenixStandard_Xxxxxxx_ExpectedToBePassed,
	}
	TestInstruction_FenixStandard_Xxxxxxx.TestInstructionAttribute = append(
		TestInstruction_FenixStandard_Xxxxxxx.TestInstructionAttribute,
		TestInstructionAttribute_FenixStandard_Xxxxxxx_ExpectedToBePassed)

	// Add FangEngine relation for Attribute - 'ExpectedToBePassed'
	// Nothing here

	// TestInstruction Attribute - 'ServerIpAddress'
	var TestInstructionAttribute_FenixStandard_Xxxxxxx_ServerIpAddress *TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_FenixStandard_Xxxxxxx_ServerIpAddress = &TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    DomainData.DomainUUID_FenixStandard,
		DomainName:                                    DomainData.DomainName_FenixStandard,
		TestInstructionUUID:                           TestInstructionUUID_FenixStandard_Xxxxxxx,
		TestInstructionName:                           TestInstructionName_FenixStandard_Xxxxxxx,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_FenixStandard_Xxxxxxx_ServerIpAddress,
		TestInstructionAttributeName:                  TestInstructionAttributeName_FenixStandard_Xxxxxxx_ServerIpAddress,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_FenixStandard_Xxxxxxx_ServerIpAddress,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_FenixStandard_Xxxxxxx_ServerIpAddress,
		TestInstructionAttributeTypeUUID:              TestInstructions.TestInstructionAttributeTypeUUID_FenixStandard_Standard,
		TestInstructionAttributeTypeName:              TestInstructions.TestInstructionAttributeTypeName_FenixStandard_Standard,
		TestInstructionAttributeValueAsString:         Domains.TestInstructionAttributeValueAsStringValue_NO_VALUE,
		TestInstructionAttributeValueUUID:             Domains.TestInstructionAttributeValueUUID_NO_VALUE,
		TestInstructionAttributeVisible:               true,
		TestInstructionAttributeEnabled:               true,
		TestInstructionAttributeMandatory:             true,
		TestInstructionAttributeVisibleInTestCaseArea: false,
		TestInstructionAttributeIsDeprecated:          false,
		TestInstructionAttributeInputMask:             ".",
		TestInstructionAttributeType:                  TestInstructionAttributeType_FenixStandard_Xxxxxxx_ServerIpAddress,
	}
	TestInstruction_FenixStandard_Xxxxxxx.TestInstructionAttribute = append(
		TestInstruction_FenixStandard_Xxxxxxx.TestInstructionAttribute,
		TestInstructionAttribute_FenixStandard_Xxxxxxx_ServerIpAddress)

	// Add FangEngine relation for Attribute - 'ExpectedToBePassed'
	// Nothing here

	// ImmatureElementModel - Xxxxxxx
	var TestInstructionImmatureElementModel_FenixStandard_Xxxxxxx *TypeAndStructs.ImmatureElementModelMessageStruct
	TestInstructionImmatureElementModel_FenixStandard_Xxxxxxx = &TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               DomainData.DomainUUID_FenixStandard,
		DomainName:               DomainData.DomainName_FenixStandard,
		ImmatureElementUUID:      TestInstructionUUID_FenixStandard_Xxxxxxx,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(TestInstructionName_FenixStandard_Xxxxxxx),
		PreviousElementUUID:      TestInstructionUUID_FenixStandard_Xxxxxxx,
		NextElementUUID:          TestInstructionUUID_FenixStandard_Xxxxxxx,
		FirstChildElementUUID:    TestInstructionUUID_FenixStandard_Xxxxxxx,
		ParentElementUUID:        TestInstructionUUID_FenixStandard_Xxxxxxx,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TI,
		OriginalElementUUID:      TestInstructionUUID_FenixStandard_Xxxxxxx,
		TopImmatureElementUUID:   TestInstructionUUID_FenixStandard_Xxxxxxx,
		IsTopElement:             true,
	}
	TestInstruction_FenixStandard_Xxxxxxx.ImmatureElementModel = append(
		TestInstruction_FenixStandard_Xxxxxxx.ImmatureElementModel,
		TestInstructionImmatureElementModel_FenixStandard_Xxxxxxx)

	return TestInstruction_FenixStandard_Xxxxxxx
}
