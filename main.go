package main

import (
	_ "embed"
	"fmt"
	"github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers"
)

//go:embed TestInstructionsAndTesInstructionContainersAndAllowedUsers/allowedUsers/allowedUsers.json
var allowedUsers []byte

func main() {

	TestInstructionsAndTesInstructionContainersAndAllowedUsers.GenerateTestInstructionsAndTestInstructionContainersAndAllowedUsers_FenixStandard(allowedUsers)
	TestInstructionsAndTesInstructionContainersAndAllowedUsers.GenerateAndVerifyRPCMessages()

	fmt.Println(TestInstructionsAndTesInstructionContainersAndAllowedUsers.TestInstructionsAndTestInstructionContainersAndAllowedUsers_FenixStandard)
	fmt.Println("Success when generating and verifying all messages for TestInstructions, TestInstructionContainers and AllowedUsers")
}
