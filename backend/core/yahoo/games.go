package yahoo

import "encoding/xml"

const (
	GameCodeMLB = "mlb"
	GameCodeNBA = "nba"
	GameCodeNFL = "nfl"
	GameCodeNHL = "nhl"

	GameFilterIsAvailable = "is_available"
	GameFilterTypes       = "game_types"
	GameFilterCodes       = "game_codes"
	GameFilterSeasons     = "seasons"
)

type Game struct {
	XMLName          xml.Name `xml:"game" json:"-"`
	Code             string   `xml:"code" json:"code"`
	ID               string   `xml:"game_id" json:"id"`
	Key              string   `xml:"game_key" json:"key"`
	Name             string   `xml:"name" json:"name"`
	Season           string   `xml:"season" json:"season"`
	Type             string   `xml:"type" json:"type"`
	URL              string   `xml:"url" json:"url"`
	RegistrationOver int      `xml:"is_registration_over" json:"registration_over"`
	GameOver         int      `xml:"is_game_over" json:"game_over"`
	Offseason        int      `xml:"is_offseason" json:"offseason"`
	Leagues          []League `xml:"leagues>league" json:"leagues,omitempty"`
	Teams            []Team   `xml:"teams>team" json:"teams,omitempty"`
	// Players          []Player         `xml:"players>player" json:"players,omitempty"`
	// GameWeeks        []GameWeeks      `xml:"game_weeks>game_week" json:"game_weeks,omitempty"`
	// StatCategories   []StatCategories `xml:"stat_categories>stat_category" json:"stat_categories,omitempty"`
	// PositionTypes    []PositionType   `xml:"position_types>position_type" json:"position_types,omitempty"`
	// RosterPositions  []RosterPosition `xml:"roster_positions>roster_position" json:"roster_positions,omitempty"`
}
