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
	appID     string
	chatNo    string
	message   string
	messageNo string
	query     string
)

// messageCmd represents the message command
var messageCmd = &cobra.Command{
	Use:   "messages",
	Short: "Manage messages",
	Long: `manage your messages. For example:

client messages --app f47c3e1f9c9960fe5c3d2749505e32004aced4b6 --chat 1
client messages --app f47c3e1f9c9960fe5c3d2749505e32004aced4b6 --chat 1 --search hey
client messages --app f47c3e1f9c9960fe5c3d2749505e32004aced4b6 --chat 1 --new hello
client messages --app f47c3e1f9c9960fe5c3d2749505e32004aced4b6 --chat 1 --show 1
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if appID == "" {
			return errors.New("app param is required")
		}
		if chatNo == "" {
			return errors.New("chat number param is required")
		}
		if message != "" {
			return newMessage()
		} else if messageNo != "" {
			return showMessage()
		}
		return listMessages()
	},
}

func init() {
	rootCmd.AddCommand(messageCmd)
	messageCmd.PersistentFlags().StringVarP(&appID, "app", "a", "", "application token")
	messageCmd.PersistentFlags().StringVarP(&chatNo, "chat", "c", "", "chat number")
	messageCmd.PersistentFlags().StringVar(&query, "search", "", "search term")
	messageCmd.PersistentFlags().StringVarP(&message, "new", "n", "", "message")
	messageCmd.PersistentFlags().StringVarP(&messageNo, "show", "s", "", "message number")
}

func listMessages() error {
	c := &http.Client{}
	url := fmt.Sprintf(`%s/api/applications/%s/chats/%s/messages?query=%s`, BaseURL, appID, chatNo, query)
	resp, err := c.Get(url)
	if err != nil {
		return errors.New("something went wrong, try again later")
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return nil
}

func newMessage() error {
	c := &http.Client{}
	url := fmt.Sprintf(`%s/api/applications/%s/chats/%s/messages`, BaseURL, appID, chatNo)
	jsondata := fmt.Sprintf(`{"message": "%s"}`, message)
	resp, err := c.Post(url, "application/json", bytes.NewBuffer([]byte(jsondata)))
	if err != nil {
		return errors.New("something went wrong, try again later")
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return nil
}

func showMessage() error {
	c := &http.Client{}
	url := fmt.Sprintf(`%s/api/applications/%s/chats/%s/messages/%s`, BaseURL, appID, chatNo, messageNo)
	resp, err := c.Get(url)
	if err != nil {
		return errors.New("something went wrong, try again later")
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return nil
}
