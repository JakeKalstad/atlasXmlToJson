package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type TextureAtlasXML struct {
	XMLName    xml.Name `xml:"TextureAtlas"`
	Text       string   `xml:",chardata"`
	ImagePath  string   `xml:"imagePath,attr"`
	SubTexture []struct {
		Text   string `xml:",chardata"`
		Name   string `xml:"name,attr"`
		X      string `xml:"x,attr"`
		Y      string `xml:"y,attr"`
		Width  string `xml:"width,attr"`
		Height string `xml:"height,attr"`
	} `xml:"SubTexture"`
}

type Frame struct {
	W int `json:"w"`
	H int `json:"h"`
	X int `json:"x"`
	Y int `json:"y"`
}

type Anchor struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Meta struct {
	Description string `json:"description"`
}

type FrameData struct {
	Filename string `json:"filename"`
	Frame    Frame  `json:"frame"`
	Anchor   Anchor `json:"anchor"`
}
type AtlasData struct {
	Frames []FrameData `json:"frames"`
	Meta   Meta        `json:"meta"`
}

func main() {
	inputFile := os.Args[1]
	outputFile := os.Args[2]
	desc := ""
	if len(os.Args) == 2 {
		desc = os.Args[3]
	}
	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	xmlData := &TextureAtlasXML{}
	xml.Unmarshal(content, &xmlData)
	jsonData := AtlasData{
		Frames: []FrameData{},
		Meta:   Meta{Description: desc},
	}
	for _, n := range xmlData.SubTexture {
		w, err := strconv.Atoi(n.Width)
		if err != nil {
			fmt.Println(err.Error())
		}
		h, err := strconv.Atoi(n.Height)
		if err != nil {
			fmt.Println(err.Error())
		}
		x, err := strconv.Atoi(n.X)
		if err != nil {
			fmt.Println(err.Error())
		}
		y, err := strconv.Atoi(n.Y)
		if err != nil {
			fmt.Println(err.Error())
		}
		nJsonData := FrameData{
			Filename: n.Name,
			Frame: Frame{
				W: w,
				H: h,
				X: x,
				Y: y,
			},
			Anchor: Anchor{X: 0.5, Y: 0.5},
		}
		jsonData.Frames = append(jsonData.Frames, nJsonData)
	}
	bytes, err := json.Marshal(jsonData)
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile(outputFile, bytes, 0644)
}
