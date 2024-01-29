package TestInstructionsAndTesInstructionContainersAndAllowedUsers

import (
	"fmt"
	"github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/DomainData"
	testInstructionContainer_StandardParallelBaseContainer "github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructionContainers/TestInstructionContainer_StandardParallelBaseContainer"
	testInstructionContainer_StandardParallelBaseContainer_v_1_0 "github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructionContainers/TestInstructionContainer_StandardParallelBaseContainer/version_1_0"
	testInstructionContainer_StandardSerialBaseContainer "github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructionContainers/TestInstructionContainer_StandardSerialBaseContainer"
	testInstructionContainer_StandardSerialBaseContainer_1_0 "github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructionContainers/TestInstructionContainer_StandardSerialBaseContainer/version_1_0"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/shared_code"
	"os"
	"time"
)

var TestInstructionsAndTestInstructionContainersAndAllowedUsers_FenixStandard *TestInstructionAndTestInstuctionContainerTypes.TestInstructionsAndTestInstructionsContainersStruct

func GenerateTestInstructionsAndTestInstructionContainersAndAllowedUsers_FenixStandard() {

	var err error

	// Load Allowed users from json-file
	err = shared_code.ImportAllowedUsersFromFile()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Generate TestInstructions
	// Standard::Xxxxxxx
	//Xxxxxxx_v_1_0.Initate_TestInstruction_FenixStandard_Xxxxxxx()

	// Build structure for all TestInstructions & TestInstructionContainers to be sent over gRPC to Fenix Backend
	TestInstructionsAndTestInstructionContainersAndAllowedUsers_FenixStandard = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionsAndTestInstructionsContainersStruct{

		// TestInstructions
		TestInstructions: &TestInstructionAndTestInstuctionContainerTypes.TestInstructionsStruct{
			TestInstructionsMap:  map[TypeAndStructs.OriginalElementUUIDType]*TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionsStruct{},
			TestInstructionsHash: shared_code.InitialValueBeforeHashed,
		},

		/*
					// TestInstruction 'Xxxxxxx'
					Xxxxxxx.TestInstructionUUID_FenixStandard_Xxxxxxx: {
						TestInstructionVersions: []*TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{

							// Version 'Xxxxxxx_1_0'
							{
								TestInstructionInstance:             Xxxxxxx_v_1_0.TestInstruction_FenixStandard_Xxxxxxx,
								TestInstructionInstanceMajorVersion: Xxxxxxx_v_1_0.TestInstruction_FenixStandard_Xxxxxxx.TestInstruction.MajorVersionNumber,
								TestInstructionInstanceMinorVersion: Xxxxxxx_v_1_0.TestInstruction_FenixStandard_Xxxxxxx.TestInstruction.MinorVersionNumber,
								Deprecated:                          Xxxxxxx_v_1_0.TestInstruction_FenixStandard_Xxxxxxx.TestInstruction.Deprecated,
								Enabled:                             Xxxxxxx_v_1_0.TestInstruction_FenixStandard_Xxxxxxx.TestInstruction.Enabled,
								TestInstructionInstanceVersionHash:  shared_code.InitialValueBeforeHashed,
							},
						},
						TestInstructionVersionsHash: shared_code.InitialValueBeforeHashed,
					},
				},
				TestInstructionsHash: shared_code.InitialValueBeforeHashed,
			},

		*/

		// TestInstructionContainers
		TestInstructionContainers: &TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainersStruct{},
		AllowedUsers:              shared_code.AllowedUsersLoadFromJsonFile,
		TestInstructionsAndTestInstructionsContainersAndUsersMessageHash: shared_code.InitialValueBeforeHashed,
		MessageCreationTimeStamp: time.Now(),
		ForceNewBaseLineForTestInstructionsAndTestInstructionContainers: false,
		ConnectorsDomain: TestInstructionAndTestInstuctionContainerTypes.ConnectorsDomainStruct{
			ConnectorsDomainUUID: DomainData.DomainUUID_FenixStandard,
			ConnectorsDomainName: DomainData.DomainName_FenixStandard,
			ConnectorsDomainHash: shared_code.InitialValueBeforeHashed,
		},
	}

	// Generate TestInstructionContainers
	// testInstructionContainer_StandardSerialBaseContainer
	testInstructionContainer_StandardSerialBaseContainer_1_0.Initiate_TestInstructionContainer_FenixStandard_Serial(TestInstructionsAndTestInstructionContainersAndAllowedUsers_FenixStandard)

	// testInstructionContainer_StandardParallelBaseContainer
	testInstructionContainer_StandardParallelBaseContainer_v_1_0.Initiate_TestInstructionContainer_FenixStandard_Parallel(TestInstructionsAndTestInstructionContainersAndAllowedUsers_FenixStandard)

	// TestInstructionContainers
	var testInstructionContainers *TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainersStruct
	testInstructionContainers = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainersStruct{
		TestInstructionContainersMap: map[TypeAndStructs.OriginalElementUUIDType]*TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainerInstanceVersionsStruct{

			// 'StandardSerialBaseContainer'
			testInstructionContainer_StandardSerialBaseContainer.TestInstructionContainerUUID_FenixStandard_StandardSerialBaseContainer: {
				TestInstructionContainerVersions: []*TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainerInstanceVersionStruct{

					//Version 'TestInstructionContainer_SpecialSerialBaseContainer_1.0'
					{
						TestInstructionContainerInstance:             testInstructionContainer_StandardSerialBaseContainer_1_0.TestInstructionContainer_FenixStandard_StandardSerialBase,
						TestInstructionContainerInstanceMajorVersion: testInstructionContainer_StandardSerialBaseContainer_1_0.TestInstructionContainer_FenixStandard_StandardSerialBase.TestInstructionContainer.MajorVersionNumber,
						TestInstructionContainerInstanceMinorVersion: testInstructionContainer_StandardSerialBaseContainer_1_0.TestInstructionContainer_FenixStandard_StandardSerialBase.TestInstructionContainer.MinorVersionNumber,
						Deprecated:                           testInstructionContainer_StandardSerialBaseContainer_1_0.TestInstructionContainer_FenixStandard_StandardSerialBase.TestInstructionContainer.Deprecated,
						Enabled:                              testInstructionContainer_StandardSerialBaseContainer_1_0.TestInstructionContainer_FenixStandard_StandardSerialBase.TestInstructionContainer.Enabled,
						TestInstructionContainerInstanceHash: shared_code.InitialValueBeforeHashed,
					},
				},
				TestInstructionContainerVersionsHash: shared_code.InitialValueBeforeHashed,
			},

			// 'StandardParallelBaseContainer'
			testInstructionContainer_StandardParallelBaseContainer.TestInstructionContainerUUID_FenixStandard_StandardParallelBaseContainer: {
				TestInstructionContainerVersions: []*TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainerInstanceVersionStruct{

					//Version 'TestInstructionContainer_SpecialSerialBaseContainer_1.0'
					{
						TestInstructionContainerInstance:             testInstructionContainer_StandardParallelBaseContainer_v_1_0.TestInstructionContainer_FenixStandard_StandardParallelBase,
						TestInstructionContainerInstanceMajorVersion: testInstructionContainer_StandardParallelBaseContainer_v_1_0.TestInstructionContainer_FenixStandard_StandardParallelBase.TestInstructionContainer.MajorVersionNumber,
						TestInstructionContainerInstanceMinorVersion: testInstructionContainer_StandardParallelBaseContainer_v_1_0.TestInstructionContainer_FenixStandard_StandardParallelBase.TestInstructionContainer.MinorVersionNumber,
						Deprecated:                           testInstructionContainer_StandardParallelBaseContainer_v_1_0.TestInstructionContainer_FenixStandard_StandardParallelBase.TestInstructionContainer.Deprecated,
						Enabled:                              testInstructionContainer_StandardParallelBaseContainer_v_1_0.TestInstructionContainer_FenixStandard_StandardParallelBase.TestInstructionContainer.Enabled,
						TestInstructionContainerInstanceHash: shared_code.InitialValueBeforeHashed,
					},
				},
				TestInstructionContainerVersionsHash: shared_code.InitialValueBeforeHashed,
			},
		},

		TestInstructionContainersHash: shared_code.InitialValueBeforeHashed,
	}
	TestInstructionsAndTestInstructionContainersAndAllowedUsers_FenixStandard.TestInstructionContainers = testInstructionContainers

	// Define temporary storage for 'LocalExecutionMethods.MethodsForLocalExecutionsStruct'
	/*
		var tempStorageLocalExecutionMethodsObject *TestInstructionAndTestInstuctionContainerTypes.AnyType


		var PushToTempStoreFunction = func(incomingLocalExecutionMethodsObject *TestInstructionAndTestInstuctionContainerTypes.AnyType) {
			tempStorageLocalExecutionMethodsObject = incomingLocalExecutionMethodsObject

		}

		var PullFromTempStoreFunction = func() (outgoingLocalExecutionMethodsObject *TestInstructionAndTestInstuctionContainerTypes.AnyType) {
			return tempStorageLocalExecutionMethodsObject
		}

	*/
	//type PushToTempStoreFunctionType[ T, any] func(T)
	//var PushToTempStore = shared_code.PushToTempStoreFunctionType[*TestInstructionAndTestInstuctionContainerTypes.AnyType](PushToTempStoreFunction)
	//var PullFromTempStore = shared_code.PullFromTempStoreFunctionType[*TestInstructionAndTestInstuctionContainerTypes.AnyType](PullFromTempStoreFunction)

	// Calculate hashes that is included in the Supported TestInstructions, TestInstructionContainers and Allowed Users message
	err = shared_code.CalculateTestInstructionAndTestInstructionContainerAndUsersMessageHashes(
		TestInstructionsAndTestInstructionContainersAndAllowedUsers_FenixStandard)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
