package utils_test

import (
	"testing"

	"github.com/Xacor/fishing_company/pkg/utils"
)

func TestEmptyUserPass(t *testing.T) {
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Empty username", args{"", "qwerty"}, true},
		{"Empty password", args{"username", ""}, true},
		{"Empty username and password", args{"", ""}, true},
		{"Not empty username and password", args{"user", "qwerty"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.EmptyUserPass(tt.args.username, tt.args.password); got != tt.want {
				t.Errorf("EmptyUserPass() = %v, want %v", got, tt.want)
			}
		})
	}
}
