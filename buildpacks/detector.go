package buildpacks

import (
	"fmt"
	"os"
)

// Program detects the programming language of the project
func DetectLanguage(Source string) (res string, err error) {

	fmt.Println("Detecting language...")

	//Get all files and folders in current directory
	files, err := os.ReadDir(Source)
	if err != nil {
		return "", err
	}

	//Check if there is already a buildpacks folder
	for _, file := range files {

		//If there is a main.go and go.mod file
		if file.Name() == "main.go" {
			return "go", nil
		}

		//If there is a package.json file
		if file.Name() == "package.json" || file.Name() == "yarn.lock" || file.Name() == "package-lock.json" {
			return "nodejs", nil
		}

		//If there is a index.php file
		if file.Name() == "index.php" {
			return "php", nil
		}

		//Find python project by checking if there is a requirements.txt file or not
		if file.Name() == "main.py" || file.Name() == "requirements.txt" {
			return "python", nil
		}

		//Find Ruby project
		if file.Name() == "Cargo.toml" {
			return "rust", nil
		}

		//Find Clojure project
		if file.Name() == "project.clj" {
			return "clojure", nil
		}

		//Find a Crystal project by checking if there is a shard.yml file or not
		if file.Name() == "shard.yml" {
			return "crystal", nil
		}

		//Find dart project by checking if there is a pubspec.yaml and a main.dart both files or not
		if file.Name() == "pubspec.yaml" || file.Name() == "main.dart" {
			return "dart", nil
		}

		//Find Elixir project by checking if there is a mix.exs file or not
		if file.Name() == "mix.exs" {
			return "elixir", nil
		}

		//Find package.yml and any .hs file to detect Haskell project
		if file.Name() == "package.yml" || file.Name() == ".hs" {
			return "haskell", nil
		}

		//Find Java project by checking if there is a pom.xml file or not
		if file.Name() == "pom.xml" {
			return "java", nil
		}

		//Find Kotlin project by checking if there is a build.gradle.kts file or not
		if file.Name() == "build.gradle.kts" {
			return "kotlin", nil
		}

		//Find Lua project by checking if there is a main.lua file or not
		if file.Name() == "main.lua" {
			return "lua", nil
		}

		//Find Perl project by checking if there is a cpanfile file or not
		if file.Name() == "cpanfile" {
			return "perl", nil
		}

		//Find Ruby project by checking if there is a Gemfile file or not
		if file.Name() == "Gemfile" {
			return "ruby", nil
		}

		//Find Scala project by checking if there is a build.sbt file or not
		if file.Name() == "build.sbt" {
			return "scala", nil
		}

		//Find Swift project by checking if there is a Package.swift file or not
		if file.Name() == "Package.swift" {
			return "swift", nil
		}

		//Find TypeScript project by checking if there is a tsconfig.json file or not
		if file.Name() == "tsconfig.json" {
			return "typescript", nil
		}

		//Find C project by checking if there is a Makefile file or not
		if file.Name() == "Makefile" {
			return "c", nil
		}

		//Find C++ project by checking if there is a Makefile file or not
		if file.Name() == "Makefile" {
			return "cpp", nil
		}

		//Find C# project by checking if there is a .csproj file or not
		if file.Name() == ".csproj" {
			return "c#", nil
		}

		//Find Objective-C project by checking if there is a .m file or not
		if file.Name() == ".m" {
			return "objective-c", nil
		}

		//Find Objective-C++ project by checking if there is a .mm file or not
		if file.Name() == ".mm" {
			return "objective-c++", nil
		}

		//Dectect dino project by checking if there is a deno.json file or not
		if file.Name() == "deno.json" {
			return "dino", nil
		}

		//Detect perl6 project by checking if there is a META6.json file or not
		if file.Name() == "META6.json" {
			return "perl6", nil
		}

		//Detect scela project by checking if there is a build.sbt file or not
		if file.Name() == "build.sbt" {
			return "scala", nil
		}
	}

	//Return no project detected
	return "", nil
}
