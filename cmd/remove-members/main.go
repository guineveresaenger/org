package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	members := []string{"guineveresaenger", "ameukam"}

	// parse the config directory given the flag configPath
	configPath := flag.String("configPath", "../../config/", "the path of the config files")

	flag.Parse()

	// TODO : Remove when ready for public usage
	fmt.Println("This is guin's awesome ability to focus. Yay!")

	// open all the yaml files in org/config
	// walk through the filenames and see if they end in `.yaml`
	yf := getYamlFiles(*configPath)
	fmt.Println(members)
	fmt.Printf(strings.Join(yf, "\n"))

	// process this data somehow

	// match an entry to a string
	for _, filename := range yf {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("sorry; error")
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()
			// compare line to each name in the remove list
			for _, member := range members {
				if line == "- "+member {
					fmt.Println(line, filename)
				}
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("sorry; scanner error")
			return
		}
	}
}

func checkForYaml(filename string) bool {
	return strings.HasSuffix(filename, ".yaml")
}

func getYamlFiles(folder string) []string {
	var files []string
	// error handling in case folder doesn't exist

	_, e := os.Stat(folder)
	if os.IsNotExist(e) {
		log.Fatal("Folder does not exist.")
	}

	// list all the files, accounting for subfolders
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".yaml") && !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("sorry; couldn't handle this file")
		return files
	}
	return files
}
