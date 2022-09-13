package yahoo

import "encoding/xml"

type Transaction struct {
	XMLName xml.Name `xml:"transaction" json:"-"`
	// Players          []Player                     `xml:"players>player" json:"players,omitempty"`
	// Types            []TransactionType            `xml:"types>type" json:"types,omitempty"`
	// TeamKey          TeamKey                      `xml:"team_key" json:"team_key,omitempty"`
	// TypesWithTeamKey []TransactionTypeWithTeamKey `xml:"types_with_team_key" json:"types_with_team_key"`
	// Count            int                          `xml:"count" json:"count"`
}
