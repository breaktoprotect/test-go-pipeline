package simplehttpclient

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetHttpBin() {
	response, err := http.Get(`https://httpbin.org/get`)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	bytes_content, err := ioutil.ReadAll(response.Body)
	content := string(bytes_content)

	log.Printf("Response (String): %s", content)

}
