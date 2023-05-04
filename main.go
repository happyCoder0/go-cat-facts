package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

const url string = "https://catfact.ninja/fact"

type CatFact struct {
	Text   string `json:"fact"`
	Length int
}

func HandleError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func FetchFact(ch chan string) {
	var fact CatFact
	client := http.DefaultClient
	response, err := client.Get(url)

	HandleError(err)

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	HandleError(err)

	err = json.Unmarshal(body, &fact)
	HandleError(err)

	ch <- string(body)
}

func main() {
	ch := make(chan string)
	go FetchFact(ch)

	fmt.Println("Loading...")
	fact := <-ch
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println(fact)

}
