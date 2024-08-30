package TestInstructions

import (
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
)

// *************************************
// *** Shared TestInstruction Attributes ***
const (
	// Attribute 'ExpectedToBePassed' is the parameter sent to CA-execution engine telling if the TestInstruction is expected to pass or fail
	TestInstructionAttributeUUID_FenixStandard_ExpectedToBePassed TypeAndStructs.TestInstructionAttributeUUIDType = "43fef553-2424-4a82-ac7e-01db0a9a8e17"
	TestInstructionAttributeName_FenixStandard_ExpectedToBePassed TypeAndStructs.TestInstructionAttributeNameType = "ExpectedToBePassed"
	TestInstructionAttributeType_FenixStandard_ExpectedToBePassed TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeTypeTypeName_FenixStandard_ExpectedToBePassed

	TestInstructionAttributeTypeUUID_StandardAttributes TypeAndStructs.TestInstructionAttributeTypeUUIDType = "a0161997-4323-45a8-93fa-8ff7d64d1b7a"
	TestInstructionAttributeTypeName_StandardAttributes TypeAndStructs.TestInstructionAttributeNameType     = "Standard Attributes"

	// Attribute 'SendTestDataToThisDomain' is the parameter that is populated with values during runtime by the TestCaseBuilder-server.
	// The values consists of a list of which Domains/ExecutionDomains that will receive the TestData
	TestInstructionAttributeUUID_FenixSentToUsersDomain_SendTestDataToThisDomain TypeAndStructs.TestInstructionAttributeUUIDType = "136d712d-a6bf-48e8-8e0f-a1d97971d2ce"
	TestInstructionAttributeName_FenixSentToUsersDomain_SendTestDataToThisDomain TypeAndStructs.TestInstructionAttributeNameType = "SendTestDataToThisDomain"
	TestInstructionAttributeType_FenixSentToUsersDomain_SendTestDataToThisDomain TypeAndStructs.TestInstructionAttributeTypeType = "TESTCASE_BUILDER_SERVER_INJECTED_COMBOBOX"
)
