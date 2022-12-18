package telegram

import (
	"fmt"
	"io"
	"net/http"
)

func SendMessage(message string) {
	formatter := fmt.Sprintf("**%s**", message)
	url := fmt.Sprintf(`https://api.telegram.org/bot{BOT_ID}/sendMessage?chat_id={CHANNEL_ID}&text=%s&parse_mode=markdown`, formatter)

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
