package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/inconshreveable/log15"
	"github.com/spf13/cobra"
)

var log = log15.Root()

var sync bool

func main() {

	cli.Flags().BoolVar(&sync, "sync", sync, "Sync documentation with Kayako")

	err := cli.Execute()
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	log.Info("Documentation synchronized")
}

var cli = &cobra.Command{
	Use:   "kayako <documentation> [<backup>]",
	Short: "Synchronize documentation with the Kayako support portal",
	Long: `Synchronize documentation with the Kayako support portal. This tool requires 
very particular documentation structure to work, but it performs validation 
locally before trying to connect. Without any provided flags only the validation
will be performed, allowing documentation writers to quickly validate their 
contributions to the local project.

When using the sync flag to synchronize documentation with Kayako the tool needs
access to account credentials with relevant privileges to do so. The tool uses 
basic authentication and fetches a username and password from the KAYAKO_USER 
and KAYAKO_PASS environment variables.`,
	Args: cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {

		src := args[0]
		backup := ""
		if len(args) > 1 {
			backup = args[1]
		}

		// scan documentation directory
		err := scan(src, backup)
		if err != nil {
			log.Error(err.Error())
			os.Exit(1)
		}

	},
}

func do(r *http.Request) (*response, error) {
	var username = os.Getenv("KAYAKO_USER")
	var password = os.Getenv("KAYAKO_PASS")

	auth := fmt.Sprintf("%s:%s", username, password)
	b64Encoded := base64.StdEncoding.EncodeToString([]byte(auth))

	r.Header.Add("Authorization", fmt.Sprintf("Basic %s", b64Encoded))
	r.Header.Add("Content-Type", "application/json")

	client := http.DefaultClient

	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := new(response)
	err = json.Unmarshal(data, response)
	if err != nil {
		return nil, err
	}

	response.body = string(data)

	return response, nil
}
