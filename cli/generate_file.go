package cli

import (
	"errors"

	"github.com/samuelsih/fakeapi/cmd"
	"github.com/samuelsih/fakeapi/utils"
	"github.com/spf13/cobra"
)

var generateJSONFileCommand = &cobra.Command{
	Use: "gen",
	Short: "Generate a fake data api & store it to .json file",
	Long: `
	Use to generate a fake data api & store it to .json file. 
	Default filename result is fakeapi_result.json (only accepts json).
	Default filename source to read is fakeapi.yaml (filename only accepts .yaml file).
	You can set your own custom source and destination filename with flag -s and -d.
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
		cmd.OnlyGenerateJSON(sourceFilename, jsonFilename)
	},
}

func init() {
	generateJSONFileCommand.Flags().StringVarP(&sourceFilename, "source", "s", "fakeapi.yaml", "Source filename to read from")
	generateJSONFileCommand.Flags().StringVarP(&jsonFilename, "destination", "d", "fakeapi_result.json", "Destination filename to write")
}