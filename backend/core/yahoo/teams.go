package yahoo

import "encoding/xml"

type Team struct {
	XMLName             xml.Name       `xml:"team" json:"-"`
	ID                  string         `xml:"team_id" json:"id"`
	Key                 string         `xml:"team_key" json:"key"`
	Name                string         `xml:"name" json:"name"`
	OwnedByCurrentLogin int            `xml:"is_owned_by_current_login" json:"owned_by_current_login"`
	URL                 string         `xml:"url" json:"url"`
	TeamLogos           []TeamLogo     `xml:"team_logos>team_logo" json:"team_logos"`
	DivisionID          int            `xml:"division_id" json:"division_id"`
	WaiverPriority      int            `xml:"waiver_priority" json:"waiver_priority"`
	FAABBalance         int            `xml:"faab_balance" json:"faab_balance"`
	NumberOfMoves       int            `xml:"number_of_moves" json:"number_of_moves"`
	NumberOfTrades      int            `xml:"number_of_trades" json:"number_of_trades"`
	RosterAdds          TeamRosterAdds `xml:"roster_adds" json:"roster_adds"`
	ClinchedPlayoffs    int            `xml:"clinched_playoffs" json:"clinched_playoffs"`
	LeagueScoringType   string         `xml:"league_scoring_type" json:"league_scoring_type"`
	DraftPosition       int            `xml:"draft_position" json:"draft_position"`
	AuctionBudgetTotal  int            `xml:"auction_budget_total" json:"auction_budget_total"`
	HasDraftGrade       int            `xml:"has_draft_grade" json:"has_draft_grade"`
	DraftGrade          string         `xml:"draft_grade" json:"draft_grade"`
	DraftRecapURL       string         `xml:"draft_recap_url" json:"draft_recap_url"`
	Managers            []TeamManager  `xml:"managers>manager" json:"managers"`
	// Stats               []TeamStat       `xml:"stats>stat" json:"stats,omitempty"`
	// Standings           TeamStandings    `xml:"standings" json:"standings,omitempty"`
	// Roster              TeamRoster       `xml:"roster" json:"roster,omitempty"`
	// DraftResults        TeamDraftResults `xml:"draft_results" json:"draft_results"`
	// Matchups            []TeamMatchup    `xml:"matchups>matchup" json:"matchups,omitempty"`
}

type TeamLogo struct {
	XMLName xml.Name `xml:"team_logo" json:"-"`
	Size    string   `xml:"size" json:"size"`
	URL     string   `xml:"url" json:"url"`
}

type TeamRosterAdds struct {
	XMLName       xml.Name `xml:"roster_adds" json:"-"`
	CoverageType  string   `xml:"coverage_type" json:"coverage_type"`
	CoverageValue string   `xml:"coverage_value" json:"coverage_value"`
	Value         int      `xml:"value" json:"value"`
}

type TeamManager struct {
	XMLName      xml.Name `xml:"manager" json:"-"`
	ID           int      `xml:"manager_id" json:"id"`
	Nickname     string   `xml:"nickname" json:"nickname"`
	UserID       string   `xml:"guid" json:"guid"`
	CurrentLogin int      `xml:"is_current_login" json:"current_login"`
	Commissioner int      `xml:"is_commissioner" json:"commissioner"`
	ImageURL     string   `xml:"image_url" json:"image_url"`
	FeloScore    int      `xml:"felo_score" json:"felo_score"`
	FeloTier     string   `xml:"felo_tier" json:"felo_tier"`
}
