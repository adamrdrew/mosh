package responses

import "encoding/xml"

type ResponseArtistDirectory struct {
	XMLName xml.Name `xml:"Directory"`
	//Will always be artist
	Type string `xml:"type,attr"`
	//This is the endpoint we use to interact with this artist
	//example /library/metadata/100098/children
	Key string `xml:"key,attr"`
	//Artist title
	Title string `xml:"title,attr"`
	//This is like the artist's ID. We can use it to reverse engineer
	//the key. This makes running commands easier
	RatingKey string `xml:"ratingKey,attr"`
}

type ResponseArtistMediaContainer struct {
	XMLName     xml.Name                  `xml:"MediaContainer"`
	Directories []ResponseArtistDirectory `xml:"Directory"`
}
