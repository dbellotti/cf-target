package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var (
		displaySpace  bool
		displayOrg    bool
		displayTarget bool
	)

	flag.BoolVar(&displaySpace, "space", false, "display space")
	flag.BoolVar(&displayOrg, "org", false, "display org")
	flag.BoolVar(&displayTarget, "api", false, "display api target")
	flag.Parse()

	var config struct {
		Target             string
		OrganizationFields struct{ Name string }
		SpaceFields        struct{ Name string }
	}

	cfHome := os.Getenv("CF_HOME")
	if cfHome == "" {
		cfHome = filepath.Join(os.Getenv("HOME"), ".cf")
	}

	f, err := os.Open(filepath.Join(cfHome, "config.json"))
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(f).Decode(&config)
	if err != nil {
		panic(err)
	}

	// not targeted
	if displayTarget && config.Target == "" {
		return
	}

	result := []string{}

	// 0 0 0
	if !displayTarget && !displayOrg && !displaySpace {
		result = append(result, config.Target)
		result = append(result, config.OrganizationFields.Name)
		result = append(result, config.SpaceFields.Name)

		fmt.Printf("%s", strings.Join(result, " > "))
		return
	}

	if displayTarget {
		result = append(result, config.Target)
	}

	if displayOrg {
		result = append(result, config.OrganizationFields.Name)
	}

	if displaySpace {
		result = append(result, config.SpaceFields.Name)
	}

	fmt.Printf("%s", strings.Join(result, " > "))
	return
}
