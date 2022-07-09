package builder

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/samuelsih/anoneee/cmd/faker"
)

type Builder struct {
	Prefix       string
	IsPaginate   bool
	AmountOfData int
	Value        map[string]any
	SliceValue   []map[string]any
}

func NewBuilder() *Builder {
	builder := Builder {
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

	var hasPrintGenerated bool

	if b.Value["id"] == "default" {
		fmt.Println("Generating fake value...")

		for i := 0; i < b.AmountOfData; i++ {
			dataMap, err := faker.Generate(b.Value)
			if err != nil {
				return err
			}
	
			b.SliceValue = append(b.SliceValue, dataMap)
		}
	
		return nil
	}

	errChan := make(chan error, b.AmountOfData)
	var mu sync.Mutex

	for i := 0; i < b.AmountOfData; i++ {
		go func () {
			dataMap, err := faker.Generate(b.Value)
			if err != nil {
				errChan <- err
				return
			}
	
			b.SliceValue = append(b.SliceValue, dataMap)
		}()

		select {
		case err := <-errChan:
			if err != nil {
				close(errChan)
				return err
			}
		default:
			printGeneratedFromFile(&mu, &hasPrintGenerated)
		}
	}

	close(errChan)
	return nil
}

func (b *Builder) WriteToJSONFile() error {
	file, err := os.Create("result.json")
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


func printGeneratedFromFile(mu *sync.Mutex, checker *bool) {
	if !(*checker) {
		mu.Lock()
		fmt.Println("[4] Generating fake value...")
		*checker = true
		mu.Unlock()
	}

}