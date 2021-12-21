package rater

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/text/encoding/charmap"
)

type Rater struct {
	//
}

type Currency struct {
	ID       string `xml:"ID,attr"`
	NumCode  uint   `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Nom      uint   `xml:"Nominal"`
	Name     string `xml:"Name"`
	Value    string `xml:"Value"`
}

type Rates struct {
	XMLName    xml.Name   `xml:"ValCurs"`
	Date       string     `xml:"Date,attr"`
	Currencies []Currency `xml:"Valute"`
}

const cbr string = "https://www.cbr.ru/scripts/XML_daily.asp"

func NewRater() *Rater {
	return &Rater{}
}

func (r *Rater) Load(date string, rates *Rates) error {
	url := fmt.Sprintf("%s?date_req=%s", cbr, date)
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	decoder := xml.NewDecoder(res.Body)
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}
		return nil, fmt.Errorf("unknown charset: %s", charset)
	}

	err = decoder.Decode(&rates)
	if err != nil {
		return err
	}

	return nil
}
