package yahoo

import "encoding/xml"

type Roster struct {
	XMLName xml.Name `xml:"roster" json:"-"`
	// Players []Player `xml:"players>player" json:"players,omitempty"`
}
