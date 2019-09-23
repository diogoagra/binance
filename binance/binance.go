package binance

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// Binance -
type Binance struct {
	APIKey    string
	APISecret string
	Debug     bool
	Lock      *sync.RWMutex
}

const binanceURL = "https://api.binance.com"

// New -
func New(apikey, apisecret string, debug bool) *Binance {
	return &Binance{
		APIKey:    apikey,
		APISecret: apisecret,
		Debug:     debug,
		Lock:      &sync.RWMutex{},
	}
}

// Sign -
func (b *Binance) Sign(message string) (response string) {
	apisignhmac256 := hmac.New(sha256.New, []byte(b.APISecret))
	apisignhmac256.Write([]byte(message))
	response = hex.EncodeToString(apisignhmac256.Sum(nil))
	return
}

// GetURL -
func (b *Binance) GetURL(endpoint, method string, private bool) (response []byte, err error) {
	b.Lock.Lock()
	defer b.Lock.Unlock()

	uri := fmt.Sprintf("%s%s", binanceURL, endpoint)

	if b.Debug {
		fmt.Println("Request:", uri)
	}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	req, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return
	}

	req.Header.Set("User-Agent", "BexsClient/Golang")

	if private {
		req.Header.Set("X-MBX-APIKEY", b.APIKey)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		if len(response) > 100 {
			response = response[0:100]
		}
		err = fmt.Errorf("Status code %v - %s", resp.StatusCode, response)
		return
	}

	if b.Debug {
		fmt.Println("Response:", string(response))
	}

	return
}

// StringToFloat -
func (b *Binance) StringToFloat(f string) (total float64) {
	total, _ = strconv.ParseFloat(f, 64)
	return
}

func unixMillis(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}
