package version

import (
	"reflect"
	"testing"
)

func TestInitVersion(t *testing.T) {
	type args struct {
		binVersion      string
		aBuildNumber    string
		aBuildTimeStamp string
		aGitBranch      string
		aGitHash        string
	}
	tests := []struct {
		name string
		args args
	}{
		{"INIT", args{"0.0.0", "123", "01.01.2021", "master", "some"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitVersion(tt.args.binVersion, tt.args.aBuildNumber, tt.args.aBuildTimeStamp, tt.args.aGitBranch, tt.args.aGitHash)
		})
	}
}

func TestGetVersion(t *testing.T) {
	tests := []struct {
		name string
		want ApplicationVersion
	}{
		{"Version", ApplicationVersion{"0.0.0.123", "01.01.2021", "master", "some"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetVersion(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
