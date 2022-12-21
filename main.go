package main

import (
	"fmt"

	bp "dauqu.com/buildpacks/buildpacks"
)

func main() {

	//Get current directory
	// dir, err := os.Getwd()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	project_dir := "/Users/harshaweb/Documents/go/vhost"

	//Dectect language
	language, err := bp.DetectLanguage(project_dir)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(language)

	//Create dockerfile
	err = bp.CreateDockerfile(project_dir, "golang", "8000", language)
	if err != nil {
		fmt.Println(err)
	}

	// //Build image
	err = bp.BuildImage(project_dir, "golang")
	if err != nil {
		fmt.Println(err)
	}

	// command := "DOCKER_BUILDKIT=1 docker build ."

	// //Execute command
	// cmd := exec.Command("sh", "-c", command)
	// //Get html output
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	// err := cmd.Run()
	// if err != nil {
	// 	fmt.Println(err)
	// }

}

