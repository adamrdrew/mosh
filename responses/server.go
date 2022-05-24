package responses

type ServerMediaContainer struct {
	Server Server `xml:"Server"`
}

type Server struct {
	Name    string `xml:"name,attr"`
	Address string `xml:"address,attr"`
	Port    string `xml:"port,attr"`
}
