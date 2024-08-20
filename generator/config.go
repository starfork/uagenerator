package generator

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	FFVersions             []float32
	ChromeVersions         []string
	EdgeVersions           []string
	OperaVersions          []string
	Pixel7AndroidVersions  []string
	Pixel6AndroidVersions  []string
	Pixel5AndroidVersions  []string
	Pixel4AndroidVersions  []string
	Nexus10AndroidVersions []string
	Nexus10Builds          []string
	OsStrings              []string
}

func LoadConfig(f string) *Config {
	c := &Config{}
	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	if err = decoder.Decode(&c); err != nil {
		panic(err)
	}
	if len(c.FFVersions) > 0 {
		ffVersions = c.FFVersions
	}
	if len(c.ChromeVersions) > 0 {
		chromeVersions = c.ChromeVersions
	}
	if len(c.EdgeVersions) > 0 {
		edgeVersions = c.EdgeVersions
	}
	if len(c.OperaVersions) > 0 {
		operaVersions = c.OperaVersions
	}
	if len(c.Pixel7AndroidVersions) > 0 {
		pixel7AndroidVersions = c.Pixel7AndroidVersions
	}
	if len(c.Pixel6AndroidVersions) > 0 {
		pixel6AndroidVersions = c.Pixel6AndroidVersions
	}
	if len(c.Pixel5AndroidVersions) > 0 {
		pixel5AndroidVersions = c.Pixel5AndroidVersions
	}
	if len(c.Pixel4AndroidVersions) > 0 {
		pixel4AndroidVersions = c.Pixel4AndroidVersions
	}

	if len(c.Nexus10AndroidVersions) > 0 {
		nexus10AndroidVersions = c.Nexus10AndroidVersions
	}
	if len(c.Nexus10Builds) > 0 {
		nexus10Builds = c.Nexus10Builds
	}
	if len(c.OsStrings) > 0 {
		osStrings = c.OsStrings
	}

	return c
}
