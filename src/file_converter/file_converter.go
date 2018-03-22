package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path"
	"path/filepath"
	"strings"
)

func main() {

	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	searchDir := user.HomeDir + "/Music"

	fmt.Println(searchDir)

	//ffmpeg -i audio.ogg audio.wav
	filepath.Walk(searchDir, func(inFile string, f os.FileInfo, err error) error {

		if path.Ext(inFile) == ".ogg" {

			outFile := strings.TrimSuffix(inFile, ".ogg") + ".wav"

			fmt.Println(inFile + " >>> " + outFile)

			cmd := exec.Command("ffmpeg", "-i", inFile, outFile)

			err := cmd.Run()
			if err != nil {
				log.Println(err)
			}
		}

		return nil
	})
}
