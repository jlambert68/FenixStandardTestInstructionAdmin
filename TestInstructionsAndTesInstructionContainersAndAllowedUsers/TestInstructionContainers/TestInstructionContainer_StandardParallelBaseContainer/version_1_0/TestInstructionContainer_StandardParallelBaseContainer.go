package version_1_0

import (
	"github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/DomainData"
	"github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructionContainers"
	"github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructionContainers/TestInstructionContainer_StandardParallelBaseContainer"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/Bonds"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/Domains"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestCaseModelElementTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
)

const (
	// *** TestInstructionContainer 'StandardParallelBaseContainer'
	TestInstructionContainerUUID_FenixStandard_StandardParallelBaseContainer                        TypeAndStructs.OriginalElementUUIDType          = TestInstructionContainer_StandardParallelBaseContainer.TestInstructionContainerUUID_FenixStandard_StandardParallelBaseContainer
	TestInstructionContainerName_FenixStandard_StandardParallelBaseContainer                        TypeAndStructs.TestInstructionContainerNameType = "Fenix Standard Parallel TestInstructionsContainer"
	TestInstructionContainerTypeUUID_FenixStandard_StandardParallelBaseContainer                                                                    = TestInstructionContainers.TestInstructionContainerTypeUUID_FenixStandard_BaseContainers
	TestInstructionContainerTypeName_FenixStandard_StandardParallelBaseContainer                                                                    = TestInstructionContainers.TestInstructionContainerTypeNameType_FenixStandard_BaseContainers
	TestInstructionContainerDescription_FenixStandard_StandardParallelBaseContainer                 string                                          = "Children of this special container is processed in parallel"
	TestInstructionContainerMouseOverText_FenixStandard_StandardParallelBaseContainer               string                                          = "Children of this special container is processed in parallel"
	TestInstructionContainerDeprecated_FenixStandard_StandardParallelBaseContainer                  bool                                            = false
	TestInstructionContainerEnabled_FenixStandard_StandardParallelBaseContainer                     bool                                            = true
	TestInstructionContainerMajorVersionNumber_FenixStandard_StandardParallelBaseContainer          int                                             = 1
	TestInstructionContainerMinorVersionNumber_FenixStandard_StandardParallelBaseContainer          int                                             = 0
	TestInstructionContainerChildrenIsParallelProcessed_FenixStandard_StandardParallelBaseContainer bool                                            = true
	TestInstructionContainerColor_FenixStandard_StandardParallelBaseContainer                       TypeAndStructs.ColorType                        = "#AAAAAA"
	TCRuleDeletion_FenixStandard_StandardParallelBaseContainer                                      TypeAndStructs.TCRuleDeletionType               = "TCRuleDeletion011"
	TCRuleSwap_FenixStandard_StandardParallelBaseContainer                                          TypeAndStructs.TCRuleSwapType                   = "TCRuleSwap012"
	TestInstructionContainerCreatingTimeStamp                                                       TypeAndStructs.UpdatedTimeStampType             = "2024-01-29 13:00:00"

	// *** DropZone *** 'xxxxxxx'

	// Attribute - 'xxxxx'
)

// TestInstructionContainer_FenixStandard_StandardParallelBase
// Variable that holds the data for the TestInstructionContainer
var TestInstructionContainer_FenixStandard_StandardParallelBase *TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainerStruct

// Initiate_TestInstructionContainer_FenixStandard_Parallel
// Function that creates all data for the TestInstructionContainer
func Initiate_TestInstructionContainer_FenixStandard_Parallel(testInstructionsAndTestInstructionContainersMessage *TestInstructionAndTestInstuctionContainerTypes.TestInstructionsAndTestInstructionsContainersStruct) {

	// Initiate the full structure
	TestInstructionContainer_FenixStandard_StandardParallelBase = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainerStruct{}

	// TestInstructionContainer - 'StandardParallelBaseContainer'
	TestInstructionContainer_FenixStandard_StandardParallelBase.TestInstructionContainer = &TypeAndStructs.TestInstructionContainerStruct{
		DomainUUID:                            DomainData.DomainUUID_FenixStandard,
		DomainName:                            DomainData.DomainName_FenixStandard,
		TestInstructionContainerUUID:          TestInstructionContainerUUID_FenixStandard_StandardParallelBaseContainer,
		TestInstructionContainerName:          TestInstructionContainerName_FenixStandard_StandardParallelBaseContainer,
		TestInstructionContainerTypeUUID:      TestInstructionContainerTypeUUID_FenixStandard_StandardParallelBaseContainer,
		TestInstructionContainerTypeName:      TestInstructionContainerTypeName_FenixStandard_StandardParallelBaseContainer,
		TestInstructionContainerDescription:   TestInstructionContainerDescription_FenixStandard_StandardParallelBaseContainer,
		TestInstructionContainerMouseOverText: TestInstructionContainerMouseOverText_FenixStandard_StandardParallelBaseContainer,
		Deprecated:                            TestInstructionContainerDeprecated_FenixStandard_StandardParallelBaseContainer,
		Enabled:                               TestInstructionContainerEnabled_FenixStandard_StandardParallelBaseContainer,
		MajorVersionNumber:                    TestInstructionContainerMajorVersionNumber_FenixStandard_StandardParallelBaseContainer,
		MinorVersionNumber:                    TestInstructionContainerMinorVersionNumber_FenixStandard_StandardParallelBaseContainer,
		UpdatedTimeStamp:                      TestInstructionContainerCreatingTimeStamp,
		ChildrenIsParallelProcessed:           TestInstructionContainerChildrenIsParallelProcessed_FenixStandard_StandardParallelBaseContainer,
	}

	// BasicTestInstructionContainerInformation - 'StandardParallelBaseContainer'
	TestInstructionContainer_FenixStandard_StandardParallelBase.BasicTestInstructionContainerInformation = &TypeAndStructs.BasicTestInstructionContainerInformationStruct{
		DomainUUID:                            DomainData.DomainUUID_FenixStandard,
		DomainName:                            DomainData.DomainName_FenixStandard,
		TestInstructionContainerUUID:          TestInstructionContainerUUID_FenixStandard_StandardParallelBaseContainer,
		TestInstructionContainerName:          TestInstructionContainerName_FenixStandard_StandardParallelBaseContainer,
		TestInstructionContainerTypeUUID:      TestInstructionContainerTypeUUID_FenixStandard_StandardParallelBaseContainer,
		TestInstructionContainerTypeName:      TestInstructionContainerTypeName_FenixStandard_StandardParallelBaseContainer,
		Deprecated:                            TestInstructionContainerDeprecated_FenixStandard_StandardParallelBaseContainer,
		MajorVersionNumber:                    TestInstructionContainerMajorVersionNumber_FenixStandard_StandardParallelBaseContainer,
		MinorVersionNumber:                    TestInstructionContainerMinorVersionNumber_FenixStandard_StandardParallelBaseContainer,
		UpdatedTimeStamp:                      TestInstructionContainerCreatingTimeStamp,
		TestInstructionContainerColor:         TestInstructionContainerColor_FenixStandard_StandardParallelBaseContainer,
		TCRuleDeletion:                        TCRuleDeletion_FenixStandard_StandardParallelBaseContainer,
		TCRuleSwap:                            TCRuleSwap_FenixStandard_StandardParallelBaseContainer,
		TestInstructionContainerDescription:   TestInstructionContainerDescription_FenixStandard_StandardParallelBaseContainer,
		TestInstructionContainerMouseOverText: TestInstructionContainerMouseOverText_FenixStandard_StandardParallelBaseContainer,
		Enabled:                               TestInstructionContainerEnabled_FenixStandard_StandardParallelBaseContainer,
		TestInstructionContainerExecutionType: Domains.TestInstructionContainerExecutionType_PARALLELLED_PROCESSED,
	}

	// ImmatureTestInstructionContainerMessage - 'StandardParallelBaseContainer'
	// No DropZone for 'EmptySerialContainer'

	// ImmatureElementModelMessage - 'StandardParallelBaseContainer' - 'TIC' in 'TIC(B10)'
	var ImmatureElementModel_TIC *TypeAndStructs.ImmatureElementModelMessageStruct
	ImmatureElementModel_TIC = &TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               DomainData.DomainUUID_FenixStandard,
		DomainName:               DomainData.DomainName_FenixStandard,
		ImmatureElementUUID:      TestInstructionContainerUUID_FenixStandard_StandardParallelBaseContainer,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(TestInstructionContainerName_FenixStandard_StandardParallelBaseContainer),
		PreviousElementUUID:      TestInstructionContainerUUID_FenixStandard_StandardParallelBaseContainer,
		NextElementUUID:          TestInstructionContainerUUID_FenixStandard_StandardParallelBaseContainer,
		FirstChildElementUUID:    Bonds.Bond_B10_BondUuid,
		ParentElementUUID:        TestInstructionContainerUUID_FenixStandard_StandardParallelBaseContainer,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TIC,
		OriginalElementUUID:      TestInstructionContainerUUID_FenixStandard_StandardParallelBaseContainer,
		TopImmatureElementUUID:   TestInstructionContainerUUID_FenixStandard_StandardParallelBaseContainer,
		IsTopElement:             true,
	}
	TestInstructionContainer_FenixStandard_StandardParallelBase.ImmatureElementModel = append(TestInstructionContainer_FenixStandard_StandardParallelBase.ImmatureElementModel, ImmatureElementModel_TIC)

	// ImmatureElementModelMessage - 'StandardParallelBaseContainer' - 'B10' in 'TIC(B10)'
	var ImmatureElementModel_B10 *TypeAndStructs.ImmatureElementModelMessageStruct
	ImmatureElementModel_B10 = &TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               DomainData.DomainUUID_FenixStandard,
		DomainName:               DomainData.DomainName_FenixStandard,
		ImmatureElementUUID:      Bonds.Bond_B10_BondUuid,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(Bonds.Bond_B10_BondName),
		PreviousElementUUID:      Bonds.Bond_B10_BondUuid,
		NextElementUUID:          Bonds.Bond_B10_BondUuid,
		FirstChildElementUUID:    Bonds.Bond_B10_BondUuid,
		ParentElementUUID:        TestInstructionContainerUUID_FenixStandard_StandardParallelBaseContainer,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_B10,
		OriginalElementUUID:      Bonds.Bond_B10_BondUuid,
		TopImmatureElementUUID:   TestInstructionContainerUUID_FenixStandard_StandardParallelBaseContainer,
		IsTopElement:             false,
	}

	TestInstructionContainer_FenixStandard_StandardParallelBase.ImmatureElementModel = append(TestInstructionContainer_FenixStandard_StandardParallelBase.ImmatureElementModel, ImmatureElementModel_B10)

}
