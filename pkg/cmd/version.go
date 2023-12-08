package cmd

import "fmt"

type Version struct {
	Version   string
	GoVersion string
	BuildTime string
	Platform  string
}

var (
	cliVersion         Version
	version   string
	goVersion string
	buildTime string
	platform  string
)

func init() {
	cliVersion = Version{
		Version:   version,
		GoVersion: goVersion,
		BuildTime: buildTime,
		Platform:  platform,
	}
}

func (ver Version) String() string {
	return fmt.Sprintf(
		"\ngit hash: %s\nGo version: %s\nBuild time: %s\nPlatform: %s",
		ver.Version,
		ver.GoVersion,
		ver.BuildTime,
		ver.Platform,
	)
}