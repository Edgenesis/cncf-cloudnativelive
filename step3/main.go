// an HTTP server that periodically gets deviceshifu-sensor/sensor
// the return data is a json looks like {"displacement":97.7842341380047,"voltage":0.3911369365520188}
// when the displacement is greater than 300, the server will send a get to deviceshifu-rtspcamera/capture and save the result as an image in the current directory
// the server will serve all images on /images API
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	sensorURL    = "http://deviceshifu-sensor/sensor"
	cameraURL    = "http://deviceshifu-rtspcamera/capture"
	pollInterval = 500 * time.Millisecond
)

type SensorData struct {
	Displacement float64 `json:"displacement"`
	Voltage      float64 `json:"voltage"`
}

func main() {
	go startPolling()

	http.HandleFunc("/images", serveImages)
	http.HandleFunc("/images/", serveIndividualImage)
	http.ListenAndServe(":8080", nil)
}

func startPolling() {
	for {
		data, err := getSensorData()
		if err != nil {
			log.Printf("Error getting sensor data: %v", err)
			time.Sleep(pollInterval)
			continue
		}

		if data.Displacement > 200 {
			fmt.Println(data)
			if err := captureImage(); err != nil {
				log.Printf("Error capturing image: %v", err)
			}
		}
		time.Sleep(pollInterval)
	}
}

func getSensorData() (*SensorData, error) {
	resp, err := http.Get(sensorURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data SensorData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func captureImage() error {
	resp, err := http.Get(cameraURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fileName := time.Now().Format("20060102_150405.jpg")
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}

func serveImages(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir(".")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")

	fmt.Fprintf(w, "<html><head><title>Images</title></head><body>")

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		ext := filepath.Ext(file.Name())
		if ext == ".jpg" {
			fmt.Fprintf(w, "<a href=\"/images/%s\" target=\"_blank\"><img src=\"/images/%s\" width=\"300\"></a><br>", file.Name(), file.Name())
		}
	}

	fmt.Fprintf(w, "</body></html>")
}

func serveIndividualImage(w http.ResponseWriter, r *http.Request) {
	imageName := r.URL.Path[len("/images/"):]
	http.ServeFile(w, r, imageName)
}
