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
	app  string
	chat string
	new  bool
)

// chatCmd represents the chat command
var chatCmd = &cobra.Command{
	Use:   "chats",
	Short: "Manage Chats",
	Long: `manage your chats. For example:

client chats --app f47c3e1f9c9960fe5c3d2749505e32004aced4b6
client chats --app f47c3e1f9c9960fe5c3d2749505e32004aced4b6 --new
client chats --app f47c3e1f9c9960fe5c3d2749505e32004aced4b6 --show 1
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if app == "" {
			return errors.New("app param is required")
		}
		if new {
			return newChat()
		} else if chat != "" {
			return showChat()
		}
		return listChats()
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)
	chatCmd.PersistentFlags().StringVarP(&app, "app", "a", "", "application token")
	chatCmd.PersistentFlags().BoolVarP(&new, "new", "n", false, "create new chat")
	chatCmd.PersistentFlags().StringVarP(&chat, "show", "s", "", "chat number")
}

func listChats() error {
	c := &http.Client{}
	url := fmt.Sprintf(`%s/api/applications/%s/chats`, BaseURL, app)
	resp, err := c.Get(url)
	if err != nil {
		return errors.New("something went wrong, try again later")
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return nil
}

func newChat() error {
	c := &http.Client{}
	url := fmt.Sprintf(`%s/api/applications/%s/chats`, BaseURL, app)
	resp, err := c.Post(url, "application/json", bytes.NewBuffer([]byte(`{}`)))
	if err != nil {
		return errors.New("something went wrong, try again later")
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return nil
}

func showChat() error {
	c := &http.Client{}
	url := fmt.Sprintf(`%s/api/applications/%s/chats/%s`, BaseURL, app, chat)
	resp, err := c.Get(url)
	if err != nil {
		return errors.New("something went wrong, try again later")
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return nil
}
