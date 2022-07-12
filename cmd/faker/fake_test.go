package faker

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	testcases := []struct {
		name    string
		args    map[string]any
		wantErr bool
	}{
		{
			"test-true-1",
			map[string]any{
				"name":    "fullname",
				"age":     "age",
				"address": "street",
			},
			false,
		},
		{
			"test-false",
			map[string]any{
				"name":    "name",
				"age":     "24",
				"address": "address",
			},
			true,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := Generate(tt.args); err != nil {
				if !tt.wantErr {
					t.Errorf("Generate() error = %v, want %v", (err != nil), tt.wantErr)
					return
				}
			}
		})
	}
}

func Test_summonFakeData(t *testing.T) {
	type args struct {
		fakeType string
	}

	tests := []struct {
		name     string
		args     args
		wantType string
		wantErr  bool
	}{
		{"test-no-error-1", args{"title"}, "string", false},
		{"test-no-error-3", args{"year"}, "int", false},
		{"test-no-error-3", args{"bool"}, "bool", false},
		{"test-no-error-4", args{"age"}, "int", false},
	
		{"test-error-1", args{"brands"}, "nil", true},
		{"test-error-2", args{"years"}, "nil", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := summonFakeData(tt.args.fakeType)
			if (err != nil) != tt.wantErr {
				t.Errorf("summonFakeData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantType == "nil" {
				if reflect.ValueOf(got).IsValid() {
					t.Errorf("summonFakeData() type = %v, want %v", reflect.ValueOf(got).IsValid(), tt.wantType)
					return
				}
			} else {
				reflectValue := reflect.ValueOf(got)
				typeValue := fmt.Sprintf("%v", reflectValue.Type())
	
				if typeValue != tt.wantType {
					t.Errorf("summonFakeData() type = %v, want %v", typeValue, tt.wantType)
					return
				}
			}
		})
	}
}

