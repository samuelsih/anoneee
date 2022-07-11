package cli

import (
	"errors"

	"github.com/samuelsih/fakeapi/cmd"
	"github.com/samuelsih/fakeapi/utils"
	"github.com/spf13/cobra"
)

var runAllCommand = &cobra.Command{
	Use: "run",
	Short: "Use to generate a fake data api, run it on localhost and store the fake data to json file",
	Long: `
	Generate fake data api from file and run it on localhost. 
	Default port is 7000.
	Default source filename to read is fakeapi.yaml (only accepts .yaml file).
	Default filename result is fakeapi_result.json (only accepts json).
	You can set your own custom port and yaml file with flag -p and -s.
	You can set your own custom destination filename with flag -d.
	`,
	Args: func(_ *cobra.Command, _ []string) error {
		if !utils.IsYAMLFile(sourceFilename) {
			return errors.New("source filename must be .yaml extension")
		}

		if !utils.IsJSONFile(jsonFilename) {
			return errors.New("destination filename must be .json extension")
		}

		return nil
	},
	Run: func(_ *cobra.Command, _ []string) {
		cmd.RunAll(sourceFilename, port, jsonFilename)
	},
}

func init() {
	runAllCommand.Flags().StringVarP(&sourceFilename, "source", "s", "fakeapi.yaml", "Source file to read from")
	runAllCommand.Flags().StringVarP(&port, "port", "p", "7000", "Port to run the localhost server")
	runAllCommand.Flags().StringVarP(&jsonFilename, "destination", "d", "fakeapi_result.json", "Destination filename to write")
}