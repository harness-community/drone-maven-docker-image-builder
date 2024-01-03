package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	pomPath := os.Getenv("PLUGIN_POM_PATH")

	if pomPath == "" {
		fmt.Println("POM Path is empty, exiting...")
		os.Exit(1)
	}

	fmt.Println("POM Path: ", pomPath)

	cmd := exec.Command("mvn", "-f", fmt.Sprintf("%s/pom.xml", pomPath), "help:evaluate", "-Dexpression=project.version", "-q", "-DforceStdout")
	output, err := cmd.Output()

	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}

	pomVersion := strings.TrimSpace(string(output))

	fmt.Println("POM Version: ", pomVersion)
	os.Setenv("POM_VERSION", pomVersion)
}
