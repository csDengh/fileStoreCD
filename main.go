package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"

	appsv1 "k8s.io/api/apps/v1"
	"sigs.k8s.io/yaml"
)

func main() {
	var deapp appsv1.Deployment
	yamlFile, err := ioutil.ReadFile("manifest/fileStore.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}
	if err = yaml.Unmarshal(yamlFile, &deapp); err != nil {
		fmt.Println(err)
		return
	}
	name := os.Getenv("IMAGE")

	if name != "" {
		deapp.Spec.Template.Spec.Containers[0].Image = name
	}

	if out, err := yaml.Marshal(deapp); err == nil {
		if err := ioutil.WriteFile("manifest/fileStore.yaml", out, fs.ModeAppend); err != nil {
			log.Fatalln(err)
		}
	} else {
		log.Fatalln(err)
	}

}
