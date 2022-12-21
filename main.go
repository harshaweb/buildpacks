package main

import (
	bp "dauqu.com/buildpacks/buildpacks"
	"fmt"
)

func main() {

	project_dir := "/Users/harshaweb/Documents/go/vhost"
	project_name := "my_name"
	project_port := "8000"

	//Dectect language
	language, err := bp.DetectLanguage(project_dir)
	if err != nil {
		fmt.Println(err)
	}

	//Create dockerfile
	err = bp.CreateDockerfile(project_dir, project_name, project_port, language)
	if err != nil {
		fmt.Println(err)
	}

	// //Build image
	err = bp.BuildImage(project_dir, project_name)
	if err != nil {
		fmt.Println(err)
	}
}
