package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

const (
	SPI_SETDESKWALLPAPER = 0x0014
	SPIF_UPDATEINIFILE   = 0x01
)

func main() {
	imageURLs, err := getBingWallpaperURL(0, 1, "zh-CN")
	if err != nil {
		fmt.Printf("get bing wallpaper url error: %v\n", err)
		return
	}
	if len(imageURLs) == 0 {
		fmt.Println("bing wallpaper url array is empty")
		return
	}
	filename := time.Now().Format("20160102") + ".jpg"
	err = downloadImage(imageURLs[0], filename)
	if err != nil {
		fmt.Printf("download image error: %v\n", err)
		return
	}
	switch runtime.GOOS {
	case "darwin", "windows":
		err = setWallpaper(filename)
	case "linux":
		fmt.Println("not support linux")
	default:
		fmt.Println("unknown system")
	}
	if err != nil {
		fmt.Printf("set wallpaper error: %v\n", err)
		return
	}
}

type BingImageInfo struct {
	StartDate     string   `json:"startdate"`
	FullStartDate string   `json:"fullstartdate"`
	EndDate       string   `json:"enddate"`
	URL           string   `json:"url"`
	URLBase       string   `json:"urlbase"`
	Copyright     string   `json:"copyright"`
	CopyrightLink string   `json:"copyrightlink"`
	Title         string   `json:"title"`
	Quiz          string   `json:"quiz"`
	WP            bool     `json:"wp"`
	Hash          string   `json:"hsh"`
	Drk           int      `json:"drk"`
	Top           int      `json:"top"`
	Bot           int      `json:"bot"`
	HS            []string `json:"hs"`
}

type BingHPImageArchiveResponse struct {
	Images   []BingImageInfo `json:"images"`
	ToolTips interface{}     `json:"tooltips"` // 不使用
}

func getBingWallpaperURL(idx, num int, area string) ([]string, error) {
	downloadBaseURL := "https://cn.bing.com"
	reqURLTmp := "https://cn.bing.com/HPImageArchive.aspx?format=js&idx=%d&n=%d&mkt=%s"
	reqURL := fmt.Sprintf(reqURLTmp, idx, num, area)
	fmt.Printf("GET %s\n", reqURL)
	client := &http.Client{}
	resp, err := client.Get(reqURL)
	if err != nil {
		return nil, fmt.Errorf("http get [%s] error: %v", reqURL, err)
	}
	rawData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body error: %v", err)
	}

	data := &BingHPImageArchiveResponse{}
	err = json.Unmarshal(rawData, data)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal error: %v, data: %s", err, rawData)
	}

	imageURLs := make([]string, 0, num)
	for _, image := range data.Images {
		url := downloadBaseURL + image.URL
		imageURLs = append(imageURLs, url)
	}
	return imageURLs, nil
}

func downloadImage(url, filename string) error {
	fmt.Printf("download url: %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("http get [%s] error: %v", url, err)
	}
	filename, err = filepath.Abs(filename)
	if err != nil {
		return fmt.Errorf("get abs path [%s] error: %v", filename, err)
	}
	imageFile, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("create file [%s] error: %v", filename, err)
	}
	defer imageFile.Close()

	n, err := io.Copy(imageFile, resp.Body)
	if err != nil {
		return fmt.Errorf("write to file error: %v", err)
	}
	fmt.Printf("download file [%s] %d bytes", filename, n)
	return nil
}
