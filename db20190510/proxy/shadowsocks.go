package proxy

import (
	"fmt"
	ss "github.com/shadowsocks/shadowsocks-go/shadowsocks"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

/**
代理请求翻墙
*/

var config struct {
	server   string
	port     int
	method   string
	password string
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func HTTPClientBySocks5(uri string) *http.Client {
	parseUri, err := url.Parse(uri)
	handleError(err)

	host, _, err := net.SplitHostPort(parseUri.Host)
	if err != nil {
		if parseUri.Scheme == "https" {
			host = net.JoinHostPort(parseUri.Host, "443")
		} else {
			host = net.JoinHostPort(parseUri.Host, "80")
		}
	} else {
		host = parseUri.Host
	}

	rawAddr, err := ss.RawAddr(host)
	handleError(err)

	serverAddr := net.JoinHostPort(config.server, strconv.Itoa(config.port))
	cipher, err := ss.NewCipher(config.method, config.password)
	handleError(err)
	dailFunc := func(network, addr string) (net.Conn, error) {
		return ss.DialWithRawAddr(rawAddr, serverAddr, cipher.Copy())
	}

	tr := &http.Transport{
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	tr.Dial = dailFunc
	return &http.Client{Transport: tr}
}

func Get(uri string) (resp *http.Response, err error) {
	client := HTTPClientBySocks5(uri)
	return client.Get(uri)
}
func Post(uri string, contentType string, body io.Reader) (resp *http.Response, err error) {
	client := HTTPClientBySocks5(uri)
	return client.Post(uri, contentType, body)
}
func PostForm(uri string, data url.Values) (resp *http.Response, err error) {
	client := HTTPClientBySocks5(uri)
	return client.PostForm(uri, data)
}
func Head(uri string) (resp *http.Response, err error) {
	client := HTTPClientBySocks5(uri)
	return client.Head(uri)
}

func Test() {
	config.method = "aes-256-cfb"
	config.password = "missssi"
	config.port = 10011
	config.server = "35.220.100.200"
	uri := "https://clients5.google.com/pagead/drt/dn/dn.js"
	resp, err := Get(uri)
	handleError(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	handleError(err)
	fmt.Println(string(body))
}
