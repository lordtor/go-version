package go_version

import (
	"fmt"
)

type (
	ApplicationVersion struct {
		Version        string `json:"version"`
		BuildTimeStamp string `json:"buildTimeStamp"`
		GitBranch      string `json:"gitBranch"`
		GitHash        string `json:"gitHash"`
	}
)

var (
	AppVersion = ApplicationVersion{}
)

//InitVersion is ...
func InitVersion(binVersion, aBuildNumber, aBuildTimeStamp, aGitBranch, aGitHash string) {
	AppVersion.Version = fmt.Sprintf("%v.%v", binVersion, aBuildNumber)
	AppVersion.GitBranch = aGitBranch
	AppVersion.GitHash = aGitHash
	AppVersion.BuildTimeStamp = aBuildTimeStamp
}

func GetVersion() ApplicationVersion {
	return AppVersion
}
