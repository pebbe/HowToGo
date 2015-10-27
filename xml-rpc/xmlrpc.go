package main

import (
	"encoding/xml"
	"fmt"
)

////////////////////////////////////////////////////////////////

type MethodResponseT struct {
	XMLName xml.Name `xml:"methodResponse"`
	Params  *ParamsT `xml:"params,omitempty"`
	Fault   *FaultT  `xml:"fault,omitempty"`
}

type FaultT struct {
	Value ValueT `xml:"value"`
}

type ParamsT struct {
	Param []ParamT `xml:"param"`
}

type ParamT struct {
	Value ValueT `xml:"value"`
}

type ValueT struct {
	I4       *int     `xml:"i4,omitempty"`
	Int      *int     `xml:"int,omitempty"`
	Boolean  *int     `xml:"bool,omitempty"`
	String   *string  `xml:"string,omitempty"`
	Text     string   `xml:",chardata"`
	Double   *float64 `xml:"double,omitempty"`
	DateTime *string  `xml:"dateTime.iso8601,omitempty"`
	Base64   *string  `xml:"base64,omitempty"`
	Struct   *StructT `xml:"struct,omitempty"`
	Array    *ArrayT  `xml:"array,omitempty"`
}

type StructT struct {
	Member []MemberT `xml:"member"`
}

type MemberT struct {
	Name  string `xml:"name"`
	Value ValueT `xml:"value"`
}

type ArrayT struct {
	Data DataT `xml:"data"`
}

type DataT struct {
	Value []ValueT `xml:"value"`
}

////////////////////////////////////////////////////////////////

type ResponseT struct {
	Translation TranslationT `xmlrpc:"translation"`
}

type TranslationT struct {
	Translated []TranslatedT `xmlrpc:"translated"`
}

type TranslatedT struct {
	Text         string  `xmlrpc:"text"`
	Rank         int     `xmlrpc:"rank"`
	TgtTokenized string  `xmlrpc:"tgt-tokenized"`
	Score        float64 `xmlrpc:"score"`
}

////////////////////////////////////////////////////////////////

func main() {

	resp := []byte(`<?xml version="1.0"?>
<methodResponse>
   <params>
      <param>
         <value><string>Boo!</string></value>
         </param>
      </params>
   </methodResponse>
`)

	r := &MethodResponseT{}
	err := xml.Unmarshal(resp, r)
	fmt.Println(err)
	if r.Params.Param[0].Value.DateTime != nil {
		fmt.Println("DateTime:", *r.Params.Param[0].Value.DateTime)
	}
	if r.Params.Param[0].Value.String != nil {
		fmt.Println("String:", *r.Params.Param[0].Value.String)
	}
	if r.Params.Param[0].Value.Text != "" {
		fmt.Println("Text:", "x"+r.Params.Param[0].Value.Text+"x")
	}

	i := 4
	r.Params.Param[0].Value.I4 = &i
	r.Params.Param[0].Value.String = nil

	output, err := xml.MarshalIndent(r, "", "  ")
	fmt.Println(err)
	fmt.Println(string(output))

}
