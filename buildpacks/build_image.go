package buildpacks

import (
	// "archive/tar"
	// "bytes"
	// "context"
	"fmt"
	"os"
	"runtime"
	//Exec
	"os/exec"
	// "github.com/docker/docker/api/types"
	// "github.com/docker/docker/client"
	// "io"
	// "io/ioutil"
	// "os"
	// "time"
)

func BuildImage(Workdir string, Name string) error {

	fmt.Println("Building image...")

	//Check os type
	os_name := runtime.GOOS

	//Create docker client
	// cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	// if err != nil {
	// 	panic(err)
	// }

	//Create context
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	// defer cancel()

	//CHeck if docker is running
	// _, err = cli.Ping(ctx)
	// if err != nil {
	// 	return err
	// }

	//Current directory
	// dir, err := os.Getwd()
	// if err != nil {
	// 	return err
	// }

	// dockerfile := dir + "/buildpacks/Dockerfile"

	// Create a buffer
	// buf := new(bytes.Buffer)
	// tw := tar.NewWriter(buf)
	// defer tw.Close()

	// Create a filereader
	// dockerFileReader, err := os.Open(dockerfile)
	// if err != nil {
	// 	return err
	// }

	// Read the actual Dockerfile
	// readDockerFile, err := ioutil.ReadAll(dockerFileReader)
	// if err != nil {
	// 	return err
	// }

	// Make a TAR header for the file
	// tarHeader := &tar.Header{
	// 	Name: dockerfile,
	// 	Size: int64(len(readDockerFile)),
	// }

	// Writes the header described for the TAR file
	// err = tw.WriteHeader(tarHeader)
	// if err != nil {
	// 	return err
	// }

	// Writes the dockerfile data to the TAR file
	// _, err = tw.Write(readDockerFile)
	// if err != nil {
	// 	return err
	// }

	// dockerFileTarReader := bytes.NewReader(buf.Bytes())

	// Define the build options to use for the file
	// https://godoc.org/github.com/docker/docker/api/types#ImageBuildOptions
	// buildOptions := types.ImageBuildOptions{
	// 	Context:    dockerFileTarReader,
	// 	Dockerfile: dockerfile,
	// 	Remove:     true,
	// 	Tags:       []string{Name},
	// }

	// Build the actual image
	// imageBuildResponse, err := cli.ImageBuild(
	// 	ctx,
	// 	dockerFileTarReader,
	// 	buildOptions,
	// )

	// if err != nil {
	// 	return err
	// }

	// // Read the STDOUT from the build process
	// defer imageBuildResponse.Body.Close()
	// _, err = io.Copy(os.Stdout, imageBuildResponse.Body)
	// if err != nil {
	// 	return err
	// }

	// //Get current directory
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	//Get dockerfile path
	dockerfile := dir + "/buildpacks/Dockerfile"

	//Execute command
	if os_name == "windows" {

		command := "DOCKER_BUILDKIT=1 docker build -t " + Name + " -f " + dockerfile + " " + Workdir
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
