package buildpacks

import (
	"fmt"
	"os"
)

func CreateDockerfile(Workdir string, Image_Name string, Port string, language string) (err error) {
	fmt.Println("Creating dockerfile...")

	docker_nodejs := `
	FROM node:16
	WORKDIR ` + Workdir + `
	COPY package*.json ./
	COPY . ./app
	EXPOSE ` + Port + `
	RUN npm install
	CMD [ "node", "index.js" ]
	`

	docker_golang := `
	FROM golang:1.16-alpine
	WORKDIR ` + Workdir + `
	COPY go.mod ./
	COPY go.sum ./
	COPY *.go ./
	RUN go get
	RUN go build main.go
	EXPOSE ` + Port + `
	CMD [ "./main" ]
	`
	docker_python := `
	FROM python:3.9.5-alpine3.13
	WORKDIR ` + Workdir + `
	COPY . .
	RUN pip install -r requirements.txt
	EXPOSE ` + Port + `
	CMD [ "python", "main.py" ]
	`
	docker_php := `
	FROM php:8.0.3-apache-buster
	WORKDIR ` + Workdir + `
	COPY . .
	EXPOSE ` + Port + `
	CMD [ "apache2-foreground" ]
	`
	docker_rust := `
	FROM rust:1.51.0-alpine3.13
	WORKDIR ` + Workdir + `
	COPY . .
	RUN cargo build --release
	EXPOSE ` + Port + `
	CMD [ "./target/release/main" ]
	`
	docker_clojure := `
	FROM clojure:openjdk-11-lein-2.9.5-alpine
	WORKDIR ` + Workdir + `
	COPY . .
	RUN lein uberjar
	EXPOSE ` + Port + `
	CMD [ "java", "-jar", "target/uberjar/main-0.1.0-SNAPSHOT-standalone.jar" ]
	`
	docker_java := `
	FROM openjdk:11.0.10-jdk-buster
	WORKDIR ` + Workdir + `
	COPY . .
	RUN javac main.java
	EXPOSE ` + Port + `
	CMD [ "java", "main" ]
	`
	docker_ruby := `
	FROM ruby:2.7.2-alpine3.13
	WORKDIR ` + Workdir + `
	COPY . .
	RUN bundle install
	EXPOSE ` + Port + `
	CMD [ "ruby", "main.rb" ]
	`
	docker_c := `
	FROM gcc:10.2.0-alpine3.13
	WORKDIR ` + Workdir + `
	COPY . .
	RUN gcc main.c -o main
	EXPOSE ` + Port + `
	CMD [ "./main" ]
	`
	docker_csharp := `
	FROM mcr.microsoft.com/dotnet/sdk:5.0-alpine
	WORKDIR ` + Workdir + `
	COPY . .
	RUN dotnet publish -c Release -o out
	EXPOSE ` + Port + `
	CMD [ "dotnet", "out/main.dll" ]
	`
	docker_swift := `
	FROM swift:5.4.2-focal
	WORKDIR ` + Workdir + `
	COPY . .
	RUN swift build -c release
	EXPOSE ` + Port + `
	CMD [ ".build/release/main" ]
	`
	docker_elixir := `
	FROM elixir:1.11.3-alpine
	WORKDIR ` + Workdir + `
	COPY . .
	RUN mix local.hex --force
	RUN mix local.rebar --force
	RUN mix deps.get
	RUN mix compile
	EXPOSE ` + Port + `
	CMD [ "mix", "run", "--no-halt" ]
	`
	docker_haskell := `
	FROM haskell:8.10.4-alpine
	WORKDIR ` + Workdir + `
	COPY . .
	RUN stack setup
	RUN stack build
	EXPOSE ` + Port + `
	CMD [ "stack", "exec", "main" ]
	`
	docker_dart := `
	FROM google/dart:2.12.0
	WORKDIR ` + Workdir + `
	COPY . .
	RUN pub get
	EXPOSE ` + Port + `
	CMD [ "dart", "main.dart" ]
	`
	docker_kotlin := `
	FROM openjdk:11.0.10-jdk-buster
	WORKDIR ` + Workdir + `
	COPY . .
	RUN kotlinc main.kt -include-runtime -d main.jar
	EXPOSE ` + Port + `
	CMD [ "java", "-jar", "main.jar" ]
	`
	docker_perl := `
	FROM perl:5.32.0-alpine3.13
	WORKDIR ` + Workdir + `
	COPY . .
	RUN cpanm --installdeps .
	EXPOSE ` + Port + `
	CMD [ "perl", "main.pl" ]
	`
	docker_scala := `
	FROM hseeberger/scala-sbt:8u282_1.5.2_2.13.5
	WORKDIR ` + Workdir + `
	COPY . .
	RUN sbt assembly
	EXPOSE ` + Port + `
	CMD [ "java", "-jar", "target/scala-2.13/main.jar" ]
	`

	//Get current directory
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	//Create Docker file using os
	file, err := os.Create(dir + "/buildpacks/Dockerfile")
	if err != nil {
		return err
	}

	//Switch method based on language
	switch language {
	case "go":
		//Write Dockerfile
		_, err = file.WriteString(docker_golang)
		if err != nil {
			return err
		}

	case "nodejs":
		//Write Dockerfile
		_, err = file.WriteString(docker_nodejs)
		if err != nil {
			return err
		}

	case "python":
		//Write Dockerfile
		_, err = file.WriteString(docker_python)
		if err != nil {
			return err
		}

	case "php":
		//Write Dockerfile
		_, err = file.WriteString(docker_php)
		if err != nil {
			return err
		}

	case "rust":
		//Write Dockerfile
		_, err = file.WriteString(docker_rust)
		if err != nil {
			return err
		}

	case "clojure":
		//Write Dockerfile
		_, err = file.WriteString(docker_clojure)
		if err != nil {
			return err
		}

	case "java":
		//Write Dockerfile
		_, err = file.WriteString(docker_java)
		if err != nil {
			return err
		}

	case "ruby":
		//Write Dockerfile
		_, err = file.WriteString(docker_ruby)
		if err != nil {
			return err
		}

	case "c":
		//Write Dockerfile
		_, err = file.WriteString(docker_c)
		if err != nil {
			return err
		}

	case "csharp":
		//Write Dockerfile
		_, err = file.WriteString(docker_csharp)
		if err != nil {
			return err
		}

	case "swift":
		//Write Dockerfile
		_, err = file.WriteString(docker_swift)
		if err != nil {
			return err
		}

	case "elixir":
		//Write Dockerfile
		_, err = file.WriteString(docker_elixir)
		if err != nil {
			return err
		}

	case "haskell":
		//Write Dockerfile
		_, err = file.WriteString(docker_haskell)
		if err != nil {
			return err
		}

	case "dart":
		//Write Dockerfile
		_, err = file.WriteString(docker_dart)
		if err != nil {
			return err
		}

	case "kotlin":
		//Write Dockerfile
		_, err = file.WriteString(docker_kotlin)
		if err != nil {
			return err
		}

	case "perl":
		//Write Dockerfile
		_, err = file.WriteString(docker_perl)
		if err != nil {
			return err
		}

	case "scala":
		//Write Dockerfile
		_, err = file.WriteString(docker_scala)
		if err != nil {
			return err
		}
	}

	//Close file
	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}
