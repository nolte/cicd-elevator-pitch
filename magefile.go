//+build mage

package main

import (
	"os"

	"github.com/magefile/mage/sh"

	// mage:import
	_ "github.com/nolte/plumbing/cmd/kind"
)

// StartUnitTests will be execute the Terratest Bases Unit Tests
func StartUnitTests() error {
	os.Chdir("./tests")
	defer os.Chdir("./..")
	return sh.Run("go", "test", "-v", "-run", "TestDockerElevatorPitch")
}
func StartHelmUnitTests() error {
	os.Chdir("./tests")
	defer os.Chdir("./..")
	return sh.Run("go", "test", "-v", "-run", "TestHelmCiCdPitchDeployment")
}
func StartMarkdownLintTests() error {
	return sh.Run("markdownlint", "-i", "{**/pitch/themes/**,**/node_modules/**}", ".")
}
