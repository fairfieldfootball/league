package yahoo

import "encoding/xml"

const (
	PlayerFilterPosition   = "position"
	PlayerFilterStatus     = "status"
	PlayerFilterSearch     = "search"
	PlayerFilterSort       = "sort"
	PlayerFilterSortType   = "sort_type"
	PlayerFilterSortSeason = "sort_season"
	PlayerFilterSortDate   = "sort_date"
	PlayerFilterSortWeek   = "sort_week"
	PlayerFilterStart      = "start"
	PlayerFilterCount      = "count"

	PlayerStatusAvailable = "A"
	PlayerStatusFreeAgent = "FA"
	PlayerStatusWaivers   = "W"
	PlayerStatusTaken     = "T"
	PlayerStatusKeepers   = "K"
)

type Player struct {
	XMLName xml.Name `xml:"player" json:"-"`
	// Stats         PlayerStats          `xml:"stats>stat" json:"stats,omitempty"`
	// Ownership     PlayerOwnership      `xml:"ownership" json:"ownership,omitempty"`
	// PercentOwned  PlayerPercentOwneded `xml:"percent_owned" json:"percent_owned"`
	// DraftAnalysis PlayerDraftAnalysis  `xml:"draft_analysis" json:"draft_analysis"`
}
