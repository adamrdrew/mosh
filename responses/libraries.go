package responses

import "encoding/xml"

type LibraryDirectory struct {
	XMLName xml.Name `xml:"Directory"`
	Type    string   `xml:"type,attr"`
	UUID    string   `xml:"uuid,attr"`
	Title   string   `xml:"title,attr"`
}

type LibraryMediaContainer struct {
	XMLName     xml.Name           `xml:"MediaContainer"`
	Directories []LibraryDirectory `xml:"Directory"`
}
