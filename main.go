package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var (
		displaySpace bool
		displayOrg   bool
	)

	flag.BoolVar(&displaySpace, "space", false, "display space")
	flag.BoolVar(&displayOrg, "org", false, "display org")
	flag.Parse()

	var config struct {
		OrganizationFields struct {
			Name string
		}
		SpaceFields struct {
			Name string
		}
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

	if displaySpace {
		println(config.SpaceFields.Name)
		return
	}

	if displayOrg {
		println(config.OrganizationFields.Name)
		return
	}

	if displaySpace == displayOrg {
		fmt.Printf("%s | %s", config.OrganizationFields.Name, config.SpaceFields.Name)
	}
}
