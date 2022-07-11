package cli

import (
	"log"
	"github.com/spf13/cobra"
)

var (
	sourceFilename, port, jsonFilename string
)

func Do() {
	rootCommand := &cobra.Command{
		Use: "fakeapi [command]",
	}

	rootCommand.AddCommand(serverCommand, generateJSONFileCommand, runAllCommand)
	
	if err := rootCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}