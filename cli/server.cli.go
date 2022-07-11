package cli

import (
	"errors"

	"github.com/samuelsih/fakeapi/cmd"
	"github.com/samuelsih/fakeapi/utils"
	"github.com/spf13/cobra"
)

var serverCommand = &cobra.Command{
	Use: "server",
	Short: "Use to generate a fake data api & run it on localhost",
	Long: `
	Generate fake data api from file and run it on localhost. 
	Default port is 7000.
	Default filename is fakeapi.yaml (filename only accepts .yaml file).
	You can set your own custom port and yaml file with flag -p and -s.
	`,
	Args: func(_ *cobra.Command, _ []string) error {
		if !utils.IsYAMLFile(sourceFilename) {
			return errors.New("file must be .yaml extension")
		}

		return nil
	},
	Run: func(_ *cobra.Command, _ []string) {
		cmd.RunDefault(sourceFilename, port)
	},
}

func init() {
	serverCommand.Flags().StringVarP(&sourceFilename, "source", "s", "fakeapi.yaml", "Source file to read from")
	serverCommand.Flags().StringVarP(&port, "port", "p", "7000", "Port to run the localhost server")
}