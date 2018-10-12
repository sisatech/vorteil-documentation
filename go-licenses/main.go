package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	err := cli.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

var cli = cobra.Command{
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var projectPrefix = args[0]
		if strs := strings.Split(args[0], "/"); len(strs) >= 3 {
			projectPrefix = strings.Join(strs[:3], "/")
		}
		c := exec.Command("licenses", args[0])
		data, err := c.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		s := string(data)
		lines := strings.Split(s, "\n")
		for i := 0; i < len(lines); i++ {
			if strings.Contains(lines[i], "/vendor/") {
				x := strings.Split(lines[i], "/vendor/")
				lines[i] = x[len(x)-1]
			} else if strings.HasPrefix(lines[i], projectPrefix) {
				lines = append(lines[:i], lines[i+1:]...)
				i--
			}
		}

		fmt.Println("# 3rd Party Libraries")
		for _, line := range lines {
			var name, url string
			index := strings.IndexAny(line, " \t")
			if index == 0 || index < 0 {
				continue
			}
			name = line[:index]
			url = "https://" + name
			fmt.Printf("* [%s](%s)\n", name, url)
		}
	},
}
