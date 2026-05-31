package sitemap_generator

import "encoding/xml"

type XMLUrl struct {
	Loc string `xml:"loc"`
}

type XMLUrlSet struct {
	XMLName xml.Name `xml:"urlset"`
	XMLNS   string   `xml:"xmlns,attr"`
	URLs    []XMLUrl `xml:"url"`
}

func (x *XMLUrlSet) ToXML(links []string) (string, error) {
	for _, link := range links {
		x.URLs = append(x.URLs, XMLUrl{Loc: link})
	}

	output, err := xml.MarshalIndent(x, "", " ")
	return xml.Header + string(output), err
}
