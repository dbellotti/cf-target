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

	// 0 0 0
	if !displayTarget && !displayOrg && !displaySpace {
		fmt.Printf("%s > %s > %s", config.Target, config.OrganizationFields.Name, config.SpaceFields.Name)
		return
	}

	// 0 0 1
	if !displayTarget && !displayOrg && displaySpace {
		fmt.Printf("%s", config.SpaceFields.Name)
		return
	}

	// 0 1 0
	if !displayTarget && displayOrg && !displaySpace {
		fmt.Printf("%s", config.OrganizationFields.Name)
		return
	}

	// 0 1 1
	if !displayTarget && displayOrg && displaySpace {
		fmt.Printf("%s > %s", config.OrganizationFields.Name, config.SpaceFields.Name)
		return
	}

	// 1 0 0
	if displayTarget && !displayOrg && !displaySpace {
		fmt.Printf("%s", config.Target)
		return
	}

	// 1 0 1
	if displayTarget && !displayOrg && displaySpace {
		fmt.Printf("%s > %s > %s", config.Target, config.OrganizationFields.Name, config.SpaceFields.Name)
		return
	}

	// 1 1 0
	if displayTarget && displayOrg && !displaySpace {
		fmt.Printf("%s > %s", config.Target, config.OrganizationFields.Name)
		return
	}

	// 1 1 1
	if displayTarget && displayOrg && displaySpace {
		fmt.Printf("%s > %s > %s", config.Target, config.OrganizationFields.Name, config.SpaceFields.Name)
		return
	}
}
