package responses

import "encoding/xml"

type ResponseAlbumDirectory struct {
	XMLName xml.Name `xml:"Directory"`
	//Will always be album
	Type string `xml:"type,attr"`
	//This is the endpoint we use to interact with this album
	//example /library/metadata/33342/children
	Key string `xml:"key,attr"`
	//Album title
	Title string `xml:"title,attr"`
	//This is like the artist's ID. We can use it to reverse engineer
	//the key. This makes running commands easier
	RatingKey string `xml:"ratingKey,attr"`
	//Artist Name
	ParentTitle string `xml:"parentTitle,attr"`
}

type ResponseAlbumMediaContainer struct {
	XMLName     xml.Name                 `xml:"MediaContainer"`
	Directories []ResponseAlbumDirectory `xml:"Directory"`
}
