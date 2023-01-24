package cli

import (
	"fmt"
	"github.com/tonnytg/kmtg/domain/pods"
	"log"
	"os"
	"text/template"
)

func Call(p pods.Pods) {
	template, err := template.ParseFiles("./pkg/cli/templates/pods.template")
	// Capture any error
	if err != nil {
		log.Fatalln(err)
	}

	// Print out the template to std
	template.Execute(os.Stdout, p)
	fmt.Println()
}
