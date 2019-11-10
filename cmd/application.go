package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	name        string
	showToken   string
	deleteToken string
)

// applicationCmd represents the application command
var applicationCmd = &cobra.Command{
	Use:   "application",
	Short: "Manage applications",
	Long: `manage your applications. For example:

client application
client application --new app1
client application --show f47c3e1f9c9960fe5c3d2749505e32004aced4b6
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if name != "" {
			return newApplication()
		} else if showToken != "" {
			return showApplication()
		}
		return listApplications()
	},
}

func init() {
	rootCmd.AddCommand(applicationCmd)
	applicationCmd.PersistentFlags().StringVarP(&name, "new", "n", "", "create new application")
	applicationCmd.PersistentFlags().StringVarP(&showToken, "show", "s", "", "show application")
}

func newApplication() error {
	c := &http.Client{}
	params := fmt.Sprintf(`{"name":"%s"}`, name)
	var jsonStr = []byte(params)
	resp, err := c.Post("http://0.0.0.0:3000/api/applications", "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		return errors.New("something went wrong, try again later")
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return nil
}

func showApplication() error {
	c := &http.Client{}
	url := fmt.Sprintf(`http://0.0.0.0:3000/api/applications/%s`, showToken)
	resp, err := c.Get(url)
	if err != nil {
		return errors.New("something went wrong, try again later")
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return nil
}

func listApplications() error {
	c := &http.Client{}
	url := fmt.Sprintf(`http://0.0.0.0:3000/api/applications`)
	resp, err := c.Get(url)
	if err != nil {
		return errors.New("something went wrong, try again later")
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return nil
}
