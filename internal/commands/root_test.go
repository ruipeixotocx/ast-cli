package commands

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"gotest.tools/assert"
)

func TestMain(m *testing.M) {
	log.Println("Commands tests started")
	// Run all tests
	exitVal := m.Run()
	log.Println("Commands tests done")
	os.Exit(exitVal)
}

func TestRootHelp(t *testing.T) {
	cmd := createASTCommand()
	args := fmt.Sprintf("--help")
	err := executeTestCommand(cmd, args)
	assert.NilError(t, err)
}

func executeTestCommand(cmd *cobra.Command, args ...string) error {
	cmd.SetArgs(args)
	cmd.SilenceUsage = true
	return cmd.Execute()
}