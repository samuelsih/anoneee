package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/samuelsih/fakeapi/cmd/builder"
	"github.com/samuelsih/fakeapi/cmd/faker"
	"github.com/samuelsih/fakeapi/cmd/server"
	"github.com/samuelsih/fakeapi/utils"

	"github.com/goccy/go-yaml"
)

func Run() {
	yamlData, err := readYAMLFile("fakeapi.yaml")
	utils.CheckError(err)

	structData := builder.NewBuilder()

	structData, err = extractYAMLData(yamlData, structData)
	if err != nil {
		log.Fatal(err)
	}

	if err := structData.Execute(); err != nil {
		log.Fatalf("Cant execute : %v", err)
	}

	fmt.Println("[5] DONE!")

	server.RunServer("7000", *structData)
}

func readYAMLFile(filename string) (map[string]any, error) {
	fmt.Println("[1] Reading file...")

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	yamlData := make(map[string]any)

	err = yaml.Unmarshal(data, &yamlData)
	return yamlData, err
}

func extractYAMLData(yamlData map[string]any, builderData *builder.Builder) (*builder.Builder, error) {
	fmt.Println("[2] Extracting data from file...")

	for yamlKey, yamlValue := range yamlData {
		switch yamlKey {
		case "PREFIX":
			if v, ok := yamlValue.(string); ok {
				builderData.Prefix = utils.FormatString(v)
			} else {
				return nil, errors.New("PREFIX value must be a string")
			}

		case "DATA":
			for _, val := range yamlValue.([]any) {

				for key, value := range val.(map[string]any) {

					if v, ok := value.(string); ok {
						if faker.IsFakerType(v) && faker.NotBrokenID(utils.FormatString(key), v) {
							builderData.Value[utils.FormatString(key)] = v
						} else {
							log.Fatalf("Unknown data type in [%v ==> %v]", key, value)
						}

					} else {
						return nil, errors.New("can't cast the type, either that is nested or array types")
					}
				}
			}

		case "AMOUNT":
			if v, ok := yamlValue.(uint64); ok {
				builderData.AmountOfData = int(v)
			} else {
				return nil, errors.New("AMOUNT must be integer")
			}

		default:
			return nil, errors.New("unknown " + yamlKey + " in yaml file")
		}
	}

	return builderData, nil
}
