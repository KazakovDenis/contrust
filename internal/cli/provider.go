package contrust

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/KazakovDenis/contrust/internal/common/log"
)

var providerCmd = &cobra.Command{
	Use:   "provider",
	Short: "Manage providers",
	Long:  "A set of commands to manage providers.",
}

var provider string

var addProvider = &cobra.Command{
	Use:   "add",
	Short: "Add a new provider",
	Long:  "Add a new provider to the system.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		provider := args[0]

		if provider == "" {
			fmt.Println("Provider is required")
			os.Exit(1)
		}

		payload := map[string]string{"name": provider}
		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			fmt.Printf("Error encoding JSON: %v\n", err)
			os.Exit(1)
		}

		url := fmt.Sprintf("%s/provider", config.ServerURL)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
		if err != nil {
			fmt.Printf("Error creating request: %v\n", err)
			os.Exit(1)
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error making request: %v\n", err)
			os.Exit(1)
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Error("Error while closing request body")
			}
		}(resp.Body)

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response body: %v\n", err)
			os.Exit(1)
		}

		if resp.StatusCode == http.StatusOK {
			green := color.New(color.FgGreen).SprintFunc()
			fmt.Println(green(string(respBody)))
		} else {
			red := color.New(color.FgRed).SprintFunc()
			fmt.Println(red(fmt.Sprintf("Error: %d %s\n%s", resp.StatusCode, resp.Status, string(respBody))))
		}
	},
}

func init() {
	rootCmd.AddCommand(providerCmd)
	providerCmd.AddCommand(addProvider)
}
