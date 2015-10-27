package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"reflect"
	"strings"
)

////////////////////////////////////////////////////////////////

type Json struct {
	ErrorCode    int            `json:"errorCode"`
	ErrorMessage string         `json:"errorMessage"`
	TimeWait     string         `json:"timeWait"`
	TimeWork     string         `json:"timeWork"`
	Translation  []TranslationT `json:"translation"`
}

type TranslationT struct {
	ErrorCode    int           `json:"errorCode,omitempty"`
	ErrorMessage string        `json:"errorMessage,omitempty"`
	Src          string        `json:"src,omitempty"`
	SrcTokenized string        `json:"src-tokenized,omitempty"`
	Translated   []TranslatedT `json:"translated,omitempty"`
}

type TranslatedT struct {
	Text         string          `json:"text,omitempty"`
	Score        float64         `json:"score"`
	Rank         int             `json:"rank"`
	Tokenized    string          `json:"tokenized,omitempty"`
	AlignmentRaw []AlignmentRawT `json:"alignment-raw,omitempty"`
}

type AlignmentRawT struct {
	SrcStart int `json:"src-start"`
	SrcEnd   int `json:"src-end"`
	TgtStart int `json:"tgt-start"`
	TgtEnd   int `json:"tgt-end"`
}

////////////////////////////////////////////////////////////////

var JSON = `{
  "errorCode": 99,
  "errorMessage": "Failed to translate some sentence(s)",
  "timeWait": "309.124us",
  "timeWork": "2.882872363s",
  "translation": [
    {
      "src-tokenized": "this is a test .",
      "translated": [
        {
          "text": "Dit is een test.",
          "score": -1.21259033679962,
          "rank": 0,
          "tokenized": "dit is een test .",
          "alignment-raw": [
            {
              "src-start": 0,
              "src-end": 2,
              "tgt-start": 0,
              "tgt-end": 2
            },
            {
              "src-start": 3,
              "src-end": 4,
              "tgt-start": 3,
              "tgt-end": 4
            }
          ]
        },
        {
          "text": "Dit is een test..",
          "score": -2.04888892173767,
          "rank": 1,
          "tokenized": "dit is een test ..",
          "alignment-raw": [
            {
              "src-start": 0,
              "src-end": 2,
              "tgt-start": 0,
              "tgt-end": 2
            },
            {
              "src-start": 3,
              "src-end": 4,
              "tgt-start": 3,
              "tgt-end": 4
            }
          ]
        }
      ]
    },
    {
      "src-tokenized": "and this .",
      "translated": [
        {
          "text": "En dit.",
          "score": -1.571592569351196,
          "rank": 0,
          "tokenized": "en dit .",
          "alignment-raw": [
            {
              "src-start": 0,
              "src-end": 0,
              "tgt-start": 0,
              "tgt-end": 0
            },
            {
              "src-start": 1,
              "src-end": 1,
              "tgt-start": 1,
              "tgt-end": 1
            },
            {
              "src-start": 2,
              "src-end": 2,
              "tgt-start": 2,
              "tgt-end": 2
            }
          ]
        },
        {
          "text": "En.",
          "score": -1.681366324424743,
          "rank": 1,
          "tokenized": "en .",
          "alignment-raw": [
            {
              "src-start": 0,
              "src-end": 0,
              "tgt-start": 0,
              "tgt-end": 0
            },
            {
              "src-start": 1,
              "src-end": 2,
              "tgt-start": 1,
              "tgt-end": 1
            }
          ]
        }
      ]
    },
    {
      "errorCode": 5,
      "errorMessage": "Line has more than 100 words (after tokenisation)"
    },
    {
      "src-tokenized": "and a final test .",
      "translated": [
        {
          "text": "En een laatste test.",
          "score": -2.985783100128174,
          "rank": 0,
          "tokenized": "en een laatste test .",
          "alignment-raw": [
            {
              "src-start": 0,
              "src-end": 0,
              "tgt-start": 0,
              "tgt-end": 0
            },
            {
              "src-start": 1,
              "src-end": 2,
              "tgt-start": 1,
              "tgt-end": 2
            },
            {
              "src-start": 3,
              "src-end": 4,
              "tgt-start": 3,
              "tgt-end": 4
            }
          ]
        },
        {
          "text": "En een laatste.",
          "score": -3.696867704391479,
          "rank": 1,
          "tokenized": "en een laatste .",
          "alignment-raw": [
            {
              "src-start": 0,
              "src-end": 0,
              "tgt-start": 0,
              "tgt-end": 0
            },
            {
              "src-start": 1,
              "src-end": 2,
              "tgt-start": 1,
              "tgt-end": 2
            },
            {
              "src-start": 3,
              "src-end": 4,
              "tgt-start": 3,
              "tgt-end": 3
            }
          ]
        }
      ]
    }
  ]
}
`

func main() {

	j := Json{}

	err := json.Unmarshal([]byte(JSON), &j)
	if err != nil {
		log.Fatal(err)
	}

	/*
		b, err := json.MarshalIndent(j, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))
	*/

	var buf bytes.Buffer

	buf.WriteString("<?xml version='1.0' encoding='UTF-8'?>\n<methodResponse><params><param><value>\n")
	marshall(reflect.ValueOf(j), &buf)
	buf.WriteString("</value></param></params></methodResponse>\n")

	fmt.Print(buf.String())
}

func marshall(r reflect.Value, buf *bytes.Buffer) {
	switch k := r.Kind(); k {
	case reflect.Struct:
		fmt.Fprintln(buf, "<struct>")
		t := r.Type()
		for i := 0; i < r.NumField(); i++ {
			f := t.Field(i)
			tag := f.Tag.Get("json")
			s := strings.Split(tag, ",")
			var n string
			omitempty := false
			if len(s) > 0 {
				n = s[0]
				for _, opt := range s[1:] {
					if opt == "omitempty" {
						omitempty = true
					}
				}
			} else {
				n = f.Name
			}
			r2 := r.Field(i)
			if omitempty && isempty(r2) {
				continue
			}
			fmt.Fprintln(buf, "<member>")
			fmt.Fprintf(buf, "<name>%s</name>\n", html.EscapeString(n))
			fmt.Fprintln(buf, "<value>")
			marshall(r2, buf)
			fmt.Fprintln(buf, "</value>")
			fmt.Fprintln(buf, "</member>")
		}
		fmt.Fprintln(buf, "</struct>")
	case reflect.Int:
		fmt.Fprintf(buf, "<int>%d</int>\n", r.Int())
	case reflect.Float64:
		fmt.Fprintf(buf, "<double>%g</double>\n", r.Float())
	case reflect.String:
		fmt.Fprintf(buf, "<string>%s</string>\n", r.String())
	case reflect.Slice:
		fmt.Fprintln(buf, "<array><data>")
		for i := 0; i < r.Len(); i++ {
			fmt.Fprintln(buf, "<value>")
			marshall(r.Index(i), buf)
			fmt.Fprintln(buf, "</value>")
		}
		fmt.Fprintln(buf, "</data></array>")
	default:
		panic(fmt.Errorf("unknown type: %s", k))
	}
}

func isempty(r reflect.Value) bool {
	switch k := r.Kind(); k {
	case reflect.Struct:
		t := r.Type()
		for i := 0; i < r.NumField(); i++ {
			s := strings.Split(t.Field(i).Tag.Get("json"), ",")
			omitempty := false
			if len(s) > 0 {
				for _, opt := range s[1:] {
					if opt == "omitempty" {
						omitempty = true
					}
				}
			}
			if !omitempty {
				return false
			}
			if !isempty(r.Field(i)) {
				return false
			}
		}
		return true
	case reflect.Int:
		if r.Int() == 0 {
			return true
		}
	case reflect.Float64:
		if r.Float() == 0 {
			return true
		}
	case reflect.String:
		if r.String() == "" {
			return true
		}
	case reflect.Bool:
		return r.Bool()
	case reflect.Slice:
		for i := 0; i < r.Len(); i++ {
			if !isempty(r.Index(i)) {
				return false
			}
		}
		return true
	default:
		panic(fmt.Errorf("unknown type: %s", k))
	}
	return false
}
