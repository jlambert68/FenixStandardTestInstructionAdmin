package version_1_0

import (
	"github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/DomainData"
	"github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructionContainers"
	"github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructionContainers/TestInstructionContainer_StandardSerialBaseContainer"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/Bonds"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/Domains"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestCaseModelElementTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
)

const (
	// *** TestInstructionContainer 'StandardSerialBaseContainer'
	TestInstructionContainerUUID_FenixStandard_StandardSerialBaseContainer                        TypeAndStructs.OriginalElementUUIDType          = TestInstructionContainer_StandardSerialBaseContainer.TestInstructionContainerUUID_FenixStandard_StandardSerialBaseContainer
	TestInstructionContainerName_FenixStandard_StandardSerialBaseContainer                        TypeAndStructs.TestInstructionContainerNameType = "Fenix Standard Serial TestInstructionsContainer"
	TestInstructionContainerTypeUUID_FenixStandard_StandardSerialBaseContainer                                                                    = TestInstructionContainers.TestInstructionContainerTypeUUID_FenixStandard_BaseContainers
	TestInstructionContainerTypeName_FenixStandard_StandardSerialBaseContainer                                                                    = TestInstructionContainers.TestInstructionContainerTypeNameType_FenixStandard_BaseContainers
	TestInstructionContainerDescription_FenixStandard_StandardSerialBaseContainer                 string                                          = "Children of this special container is processed in serial"
	TestInstructionContainerMouseOverText_FenixStandard_StandardSerialBaseContainer               string                                          = "Children of this special container is processed in serial"
	TestInstructionContainerDeprecated_FenixStandard_StandardSerialBaseContainer                  bool                                            = false
	TestInstructionContainerEnabled_FenixStandard_StandardSerialBaseContainer                     bool                                            = true
	TestInstructionContainerMajorVersionNumber_FenixStandard_StandardSerialBaseContainer          int                                             = 1
	TestInstructionContainerMinorVersionNumber_FenixStandard_StandardSerialBaseContainer          int                                             = 0
	TestInstructionContainerChildrenIsParallelProcessed_FenixStandard_StandardSerialBaseContainer bool                                            = false
	TestInstructionContainerColor_FenixStandard_StandardSerialBaseContainer                       TypeAndStructs.ColorType                        = "#AAAAAA"
	TCRuleDeletion_FenixStandard_StandardSerialBaseContainer                                      TypeAndStructs.TCRuleDeletionType               = "TCRuleDeletion011"
	TCRuleSwap_FenixStandard_StandardSerialBaseContainer                                          TypeAndStructs.TCRuleSwapType                   = "TCRuleSwap012"
	TestInstructionContainerCreatingTimeStamp                                                     TypeAndStructs.UpdatedTimeStampType             = "2024-01-29 13:00:00"

	// *** DropZone *** 'xxxxxxx'

	// Attribute - 'xxxxx'
)

// TestInstructionContainer_FenixStandard_StandardSerialBase
// Variable that holds the data for the TestInstructionContainer
var TestInstructionContainer_FenixStandard_StandardSerialBase *TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainerStruct

// Initiate_TestInstructionContainer_FenixStandard_Serial
// Function that creates all data for the TestInstructionContainer
func Initiate_TestInstructionContainer_FenixStandard_Serial(testInstructionsAndTestInstructionContainersMessage *TestInstructionAndTestInstuctionContainerTypes.TestInstructionsAndTestInstructionsContainersStruct) {

	// Initiate the full structure
	TestInstructionContainer_FenixStandard_StandardSerialBase = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainerStruct{}

	// TestInstructionContainer - 'StandardSerialBaseContainer'
	TestInstructionContainer_FenixStandard_StandardSerialBase.TestInstructionContainer = &TypeAndStructs.TestInstructionContainerStruct{
		DomainUUID:                            DomainData.DomainUUID_FenixStandard,
		DomainName:                            DomainData.DomainName_FenixStandard,
		TestInstructionContainerUUID:          TestInstructionContainerUUID_FenixStandard_StandardSerialBaseContainer,
		TestInstructionContainerName:          TestInstructionContainerName_FenixStandard_StandardSerialBaseContainer,
		TestInstructionContainerTypeUUID:      TestInstructionContainerTypeUUID_FenixStandard_StandardSerialBaseContainer,
		TestInstructionContainerTypeName:      TestInstructionContainerTypeName_FenixStandard_StandardSerialBaseContainer,
		TestInstructionContainerDescription:   TestInstructionContainerDescription_FenixStandard_StandardSerialBaseContainer,
		TestInstructionContainerMouseOverText: TestInstructionContainerMouseOverText_FenixStandard_StandardSerialBaseContainer,
		Deprecated:                            TestInstructionContainerDeprecated_FenixStandard_StandardSerialBaseContainer,
		Enabled:                               TestInstructionContainerEnabled_FenixStandard_StandardSerialBaseContainer,
		MajorVersionNumber:                    TestInstructionContainerMajorVersionNumber_FenixStandard_StandardSerialBaseContainer,
		MinorVersionNumber:                    TestInstructionContainerMinorVersionNumber_FenixStandard_StandardSerialBaseContainer,
		UpdatedTimeStamp:                      TestInstructionContainerCreatingTimeStamp,
		ChildrenIsParallelProcessed:           TestInstructionContainerChildrenIsParallelProcessed_FenixStandard_StandardSerialBaseContainer,
	}

	// BasicTestInstructionContainerInformation - 'StandardSerialBaseContainer'
	TestInstructionContainer_FenixStandard_StandardSerialBase.BasicTestInstructionContainerInformation = &TypeAndStructs.BasicTestInstructionContainerInformationStruct{
		DomainUUID:                            DomainData.DomainUUID_FenixStandard,
		DomainName:                            DomainData.DomainName_FenixStandard,
		TestInstructionContainerUUID:          TestInstructionContainerUUID_FenixStandard_StandardSerialBaseContainer,
		TestInstructionContainerName:          TestInstructionContainerName_FenixStandard_StandardSerialBaseContainer,
		TestInstructionContainerTypeUUID:      TestInstructionContainerTypeUUID_FenixStandard_StandardSerialBaseContainer,
		TestInstructionContainerTypeName:      TestInstructionContainerTypeName_FenixStandard_StandardSerialBaseContainer,
		Deprecated:                            TestInstructionContainerDeprecated_FenixStandard_StandardSerialBaseContainer,
		MajorVersionNumber:                    TestInstructionContainerMajorVersionNumber_FenixStandard_StandardSerialBaseContainer,
		MinorVersionNumber:                    TestInstructionContainerMinorVersionNumber_FenixStandard_StandardSerialBaseContainer,
		UpdatedTimeStamp:                      TestInstructionContainerCreatingTimeStamp,
		TestInstructionContainerColor:         TestInstructionContainerColor_FenixStandard_StandardSerialBaseContainer,
		TCRuleDeletion:                        TCRuleDeletion_FenixStandard_StandardSerialBaseContainer,
		TCRuleSwap:                            TCRuleSwap_FenixStandard_StandardSerialBaseContainer,
		TestInstructionContainerDescription:   TestInstructionContainerDescription_FenixStandard_StandardSerialBaseContainer,
		TestInstructionContainerMouseOverText: TestInstructionContainerMouseOverText_FenixStandard_StandardSerialBaseContainer,
		Enabled:                               TestInstructionContainerEnabled_FenixStandard_StandardSerialBaseContainer,
		TestInstructionContainerExecutionType: Domains.TestInstructionContainerExecutionType_SERIAL_PROCESSED,
	}

	// ImmatureTestInstructionContainerMessage - 'StandardSerialBaseContainer'
	// No DropZone for 'EmptySerialContainer'

	// ImmatureElementModelMessage - 'StandardSerialBaseContainer' - 'TIC' in 'TIC(B10)'
	var ImmatureElementModel_TIC *TypeAndStructs.ImmatureElementModelMessageStruct
	ImmatureElementModel_TIC = &TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               DomainData.DomainUUID_FenixStandard,
		DomainName:               DomainData.DomainName_FenixStandard,
		ImmatureElementUUID:      TestInstructionContainerUUID_FenixStandard_StandardSerialBaseContainer,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(TestInstructionContainerName_FenixStandard_StandardSerialBaseContainer),
		PreviousElementUUID:      TestInstructionContainerUUID_FenixStandard_StandardSerialBaseContainer,
		NextElementUUID:          TestInstructionContainerUUID_FenixStandard_StandardSerialBaseContainer,
		FirstChildElementUUID:    Bonds.Bond_B10_BondUuid,
		ParentElementUUID:        TestInstructionContainerUUID_FenixStandard_StandardSerialBaseContainer,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TIC,
		OriginalElementUUID:      TestInstructionContainerUUID_FenixStandard_StandardSerialBaseContainer,
		TopImmatureElementUUID:   TestInstructionContainerUUID_FenixStandard_StandardSerialBaseContainer,
		IsTopElement:             true,
	}
	TestInstructionContainer_FenixStandard_StandardSerialBase.ImmatureElementModel = append(TestInstructionContainer_FenixStandard_StandardSerialBase.ImmatureElementModel, ImmatureElementModel_TIC)

	// ImmatureElementModelMessage - 'StandardSerialBaseContainer' - 'B10' in 'TIC(B10)'
	var ImmatureElementModel_B10 *TypeAndStructs.ImmatureElementModelMessageStruct
	ImmatureElementModel_B10 = &TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               DomainData.DomainUUID_FenixStandard,
		DomainName:               DomainData.DomainName_FenixStandard,
		ImmatureElementUUID:      Bonds.Bond_B10_BondUuid,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(Bonds.Bond_B10_BondName),
		PreviousElementUUID:      Bonds.Bond_B10_BondUuid,
		NextElementUUID:          Bonds.Bond_B10_BondUuid,
		FirstChildElementUUID:    Bonds.Bond_B10_BondUuid,
		ParentElementUUID:        TestInstructionContainerUUID_FenixStandard_StandardSerialBaseContainer,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_B10,
		OriginalElementUUID:      Bonds.Bond_B10_BondUuid,
		TopImmatureElementUUID:   TestInstructionContainerUUID_FenixStandard_StandardSerialBaseContainer,
		IsTopElement:             false,
	}

	TestInstructionContainer_FenixStandard_StandardSerialBase.ImmatureElementModel = append(TestInstructionContainer_FenixStandard_StandardSerialBase.ImmatureElementModel, ImmatureElementModel_B10)

}
