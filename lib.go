package hcat

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/yosssi/gohtml"

	"github.com/kendfss/but"
)

type node struct {
	Attr     []xml.Attr
	XMLName  xml.Name
	Children []node `xml:",any"`
	Text     string `xml:",chardata"`
}

func Scrape(url string) io.ReadCloser {
	resp, err := http.Get(url)
	but.Must(err)
	return resp.Body
}

func Read(body io.ReadCloser) []byte {
	data, err := ioutil.ReadAll(body)
	but.Must(err)
	return data
}

func Prettify(body []byte) []byte {
	return []byte(gohtml.Format(string(body)))
}
