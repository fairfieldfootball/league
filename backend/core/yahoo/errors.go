package yahoo

import (
	"encoding/xml"
)

type Err struct {
	XMLName     xml.Name `xml:"error"`
	Description string   `json:"description" xml:"description"`
}
