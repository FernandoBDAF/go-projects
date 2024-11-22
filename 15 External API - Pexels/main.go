package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"math/rand"

	"github.com/joho/godotenv"
)

const (
	PhotoApi = "https://api.pexels.com/v1"
	VideoApi = "https://api.pexels.com/videos"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	apiKey := os.Getenv("PEXELS_API_KEY")
	if apiKey == "" {
		log.Fatal("PEXELS_API_KEY is not set")
	}

	var c = NewClient(apiKey)

	result, err := c.SearchPhotos("nature", 1, 10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
	if result.Page == 0 {
		log.Fatal("Page is 0")
	}
}

type Client struct {
	apiKey         string
	httpClient     *http.Client
	RemainingTimes int32
}

type SearchPhotosResult struct {
	Page     int32   `json:"page"`
	PerPage  int32   `json:"per_page"`
	Total    int32   `json:"total"`
	NextPage string  `json:"next_page"`
	Photos   []Photo `json:"photos"`
}

type Photo struct {
	ID              int32  `json:"id"`
	Width           int32  `json:"width"`
	Height          int32  `json:"height"`
	URL             string `json:"url"`
	Src             Src    `json:"src"`
	Photographer    string `json:"photographer"`
	PhotographerURL string `json:"photographer_url"`
}

type CuratedPhotosResult struct {
	Page     int32   `json:"page"`
	PerPage  int32   `json:"per_page"`
	NextPage string  `json:"next_page"`
	Photos   []Photo `json:"photos"`
}

type Src struct {
	Original  string `json:"original"`
	Large2X   string `json:"large2x"`
	Large     string `json:"large"`
	Medium    string `json:"medium"`
	Small     string `json:"small"`
	Portrait  string `json:"portrait"`
	Landscape string `json:"landscape"`
	Square    string `json:"square"`
	Tiny      string `json:"tiny"`
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:         apiKey,
		httpClient:     &http.Client{Timeout: 10 * time.Second},
		RemainingTimes: 50,
	}
}

type SearchVideoResult struct {
	Page     int32   `json:"page"`
	PerPage  int32   `json:"per_page"`
	Total    int32   `json:"total"`
	NextPage string  `json:"next_page"`
	Videos   []Video `json:"videos"`
}

type Video struct {
	ID              int32  `json:"id"`
	Width           int32  `json:"width"`
	Height          int32  `json:"height"`
	URL             string `json:"url"`
	FullRes         interface{} `json:"full_res"`
	Duration        float64     `json:"duration"`
	Image           string `json:"image"`
	VideoFiles      []VideoFiles `json:"video_files"`
	VideoPictures   []VideoPictures `json:"video_pictures"`
}

type PopularVideosResult struct {
	Page     int32   `json:"page"`
	PerPage  int32   `json:"per_page"`
	Total    int32   `json:"total"`
	URL      string  `json:"url"`
	Videos   []Video `json:"videos"`
}

type VideoFiles struct {
	ID        int32  `json:"id"`
	Quality   string `json:"quality"`
	FileType  string `json:"file_type"`
	Width     int32  `json:"width"`
	Height    int32  `json:"height"`
	Link      string `json:"link"`
}

type VideoPictures struct {
	ID        int32  `json:"id"`
	Picture   string `json:"picture"`
	Number    int32  `json:"number"`
}

func (c *Client) SearchPhotos(query string, page int32, perPage int32) (*SearchPhotosResult, error) {
	requestUrl := fmt.Sprintf("%s/search?query=%s&page=%d&per_page=%d", PhotoApi, query, page, perPage)
	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	times, err := strconv.Atoi(resp.Header.Get("X-Ratelimit-Remaining"))
	if err != nil {
		return nil, err
	}
	c.RemainingTimes = int32(times)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search photos failed: %s", resp.Status)
	}

	var result SearchPhotosResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) CuratedPhotos(page int32, perPage int32) (*CuratedPhotosResult, error) {
	requestUrl := fmt.Sprintf("%s/curated?page=%d&per_page=%d", PhotoApi, page, perPage)
	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	times, err := strconv.Atoi(resp.Header.Get("X-Ratelimit-Remaining"))
	if err != nil {
		return nil, err
	}
	c.RemainingTimes = int32(times)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("curated photos failed: %s", resp.Status)
	}

	var result CuratedPhotosResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetPhoto(id int32) (*Photo, error) {
	requestUrl := fmt.Sprintf("%s/photos/%d", PhotoApi, id)
	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	times, err := strconv.Atoi(resp.Header.Get("X-Ratelimit-Remaining"))
	if err != nil {
		return nil, err
	}
	c.RemainingTimes = int32(times)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get photo failed: %s", resp.Status)
	}

	var result Photo
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetRandomPhoto() (*Photo, error) {
	localRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	randNum := localRand.Intn(1001)
	result, err := c.CuratedPhotos(int32(randNum), 1)
	if err != nil || len(result.Photos) == 0 {
		return nil, err
	}
	return &result.Photos[0], nil
}

// SearchVideo
func (c *Client) SearchVideo(query string, page int32, perPage int32) (*SearchVideoResult, error) {
	requestUrl := fmt.Sprintf("%s/videos/search?query=%s&page=%d&per_page=%d", VideoApi, query, page, perPage)
	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	times, err := strconv.Atoi(resp.Header.Get("X-Ratelimit-Remaining"))
	if err != nil {
		return nil, err
	}
	c.RemainingTimes = int32(times)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search video failed: %s", resp.Status)
	}

	var result SearchVideoResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Popular Videos
func (c *Client) PopularVideos(page int32, perPage int32) (*PopularVideosResult, error) {
	requestUrl := fmt.Sprintf("%s/videos/popular?page=%d&per_page=%d", VideoApi, page, perPage)
	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	times, err := strconv.Atoi(resp.Header.Get("X-Ratelimit-Remaining"))
	if err != nil {
		return nil, err
	}
	c.RemainingTimes = int32(times)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("popular videos failed: %s", resp.Status)
	}

	var result PopularVideosResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

// get random video
func (c *Client) GetRandomVideo() (*Video, error) {
	localRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	randNum := localRand.Intn(1001)
	result, err := c.PopularVideos(int32(randNum), 1)
	if err != nil || len(result.Videos) == 0 {
		return nil, err
	}
	return &result.Videos[0], nil
}

// get remaining requests in this month

func (c *Client) GetRemainingRequestsInThisMonth() int32 {
	return c.RemainingTimes
}
