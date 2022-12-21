package buildpacks

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func BuildImage(Workdir string, Name string) error {

	fmt.Println("Building image...")

	//Check os type
	os_name := runtime.GOOS

	// //Get current directory
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	//Get dockerfile path
	dockerfile := dir + "/buildpacks/Dockerfile"

	//Execute command
	if os_name == "windows" {

		//Command for windows
		command := `docker build -t ` + Name + ` .`
		cmd := exec.Command("cmd", "/C", command)
		//Get html output
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}

	} else {
		command := "DOCKER_BUILDKIT=1 docker build -t " + Name + " -f " + dockerfile + " " + Workdir
		cmd := exec.Command("sh", "-c", command)
		//Get html output
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	}

	return nil
}
