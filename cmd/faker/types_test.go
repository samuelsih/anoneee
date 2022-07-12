package faker

import (
	"testing"
)

func TestIsFakerType(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{"test-true-default", "default", true},
		{"test-true-uuid", "uuid", true},
		{"test-true-fullname", "fullname", true},
		{"test-true-gender", "gender", true},
		{"test-true-paragraphs", "paragraphs", true},

		{"test-false-defaults", "defaults", false},
		{"test-false-id", "id", false},
		{"test-false-name", "name", false},
		{"test-false-female", "female", false},
		{"test-false-word", "word", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFakerType(tt.args); got != tt.want {
				t.Errorf("IsFakerType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotBrokenID(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test-broken-id", args{"id", "defaults"}, false},
		{"test-broken-products", args{"products", "uuid"}, false},
		{"test-broken-name", args{"name", "default"}, false},

		{"test-not-broken-id-default", args{"id", "default"}, true},
		{"test-not-broken-id-uuid", args{"id", "uuid"}, true},
		{"test-not-broken-random-type", args{"name", "fullname"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NotBrokenID(tt.args.key, tt.args.value); got != tt.want {
				t.Errorf("NotBrokenID() = %v, want %v", got, tt.want)
			}
		})
	}
}
