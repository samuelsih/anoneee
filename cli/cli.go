package cli

import (
	"errors"
	"log"

	"github.com/samuelsih/fakeapi/cmd"
	"github.com/samuelsih/fakeapi/utils"
	"github.com/spf13/cobra"
)

var (
	source, port string
)

var rootCmd = &cobra.Command{
	Use:   "run",
	Short: "Use to generate a fake data api & run it on localhost",
	Long: `
	Generate fake data api from file and run it on localhost. 
	Default port is 7000.
	Default filename is fakeapi.yaml (filename only accepts .yaml file).
	You can set your own custom port and yaml file with flag -p and -s.
	Also you can write the result to json file with persistent flag -w (default fakeapi_result.json)
	`,
	Args: func(_ *cobra.Command, _ []string) error {
		if source != "" {
			if !utils.IsYAMLFile(source) {
				return errors.New("file must be .yaml extension")
			}
		}

		return nil
	},
	Run: func(_ *cobra.Command, _ []string) {
		cmd.Run(source, port)
	},
}

func Do() {
	rootCmd.LocalFlags().StringVarP(&source, "source", "s", "fakeapi.yaml", "Source directory to read from")
	rootCmd.LocalFlags().StringVarP(&port, "port", "p", "7000", "Port to serve the api")
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}