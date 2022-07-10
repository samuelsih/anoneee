package builder

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/samuelsih/fakeapi/cmd/faker"
)

type Builder struct {
	Prefix       string
	IsPaginate   bool
	AmountOfData int
	Value        map[string]any
	SliceValue   []map[string]any
}

func NewBuilder() *Builder {
	builder := Builder{
		Prefix:       "api",
		IsPaginate:   false,
		AmountOfData: 10,
		Value:        make(map[string]any, 1),
	}

	return &builder
}

func (b *Builder) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(b.SliceValue)
}

func (b *Builder) handleID() {
	fmt.Println("[3] Handling the ID...")

	if _, ok := b.Value["id"]; !ok {
		b.Value["id"] = "default"
	}
}

func (b *Builder) Execute() error {
	b.handleID()

	fmt.Println("[4] Generating fake value...")

	if b.Value["id"] == "default" {
		for i := 0; i < b.AmountOfData; i++ {
			dataMap, err := faker.Generate(b.Value)
			if err != nil {
				return err
			}

			b.SliceValue = append(b.SliceValue, dataMap)
		}

		return nil
	}

	var wg sync.WaitGroup

	wg.Add(b.AmountOfData)
	
	for i := 0; i < b.AmountOfData; i++ {
		var genErr error = nil
		go func(v map[string]any) {
			defer wg.Done()
			dataMap, err := faker.Generate(v)
			if err != nil {
				genErr = err
				return
			}

			b.SliceValue = append(b.SliceValue, dataMap)
		}(b.Value)

		if genErr != nil {
			return genErr
		}
	}

	wg.Wait()

	return nil
}

func (b *Builder) WriteToJSONFile(filename string) error {
	if filename == "" {
		filename = "fakeapi_result.json"
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	jsonBytes, err := json.MarshalIndent(b.SliceValue, "", "\t")
	if err != nil {
		return err
	}

	_, err = file.Write(jsonBytes)
	return err
}
