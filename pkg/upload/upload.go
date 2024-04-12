package upload

import (
    "log"
    "net/http"
	"bytes"
	"time"
)

func UploadToObs(resJson []byte, obsUrl string, fileName string)(error){
    
    currentTime := time.Now()
    timeString := currentTime.Format("2006-01-02 15:04:05")

	targetUrl := obsUrl+"/"+timeString+fileName
	// fmt.Println("[+] target url: %v", targetUrl)

	client := &http.Client{}
	req, err := http.NewRequest("PUT", targetUrl, bytes.NewBuffer(resJson))
	if err != nil {
		log.Println("[!] Error creating request:", err)
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("[!] Error sending request:", err)
		return err
	}
	defer resp.Body.Close()

	log.Println("POST Response Status:", resp.Status)
	return nil
}

func UploadToHostObs(resJson []byte, URL string, obsUrl string, fileName string)(error){
    currentTime := time.Now()
    timeString := currentTime.Format("2006-01-02 15:04:05")

    url := URL+"/"+timeString+fileName

	client := &http.Client{}
	request, err := http.NewRequest("PUT", url, bytes.NewBuffer(resJson))
	if err != nil {
		log.Println("Error creating PUT request:", err)
		return err
	}

	request.Host = obsUrl
	// 发送请求
	postResponse, err := client.Do(request)
	if err != nil {
		log.Println("Error making PUT request:", err)
		return err
	}
	defer postResponse.Body.Close()

	// 输出响应
	log.Println("POST Response Status:", postResponse.Status)
    return nil
}