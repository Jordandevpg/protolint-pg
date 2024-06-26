package config

import "github.com/Jordandevpg/protolint-pg/internal/stringsutil"

// Files represents the target files.
type Files struct {
	Exclude []string `yaml:"exclude" json:"exclude" toml:"exclude"`
}

func (d Files) shouldSkipRule(
	displayPath string,
) bool {
	return stringsutil.ContainsCrossPlatformPathInSlice(displayPath, d.Exclude)
}
