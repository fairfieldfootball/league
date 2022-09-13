package yahoo

import "encoding/xml"

type League struct {
	XMLName               xml.Name       `xml:"league" json:"-"`
	ID                    string         `xml:"league_id" json:"id"`
	Key                   string         `xml:"league_key" json:"key"`
	Name                  string         `xml:"name" json:"name"`
	URL                   string         `xml:"url" json:"url"`
	Password              string         `xml:"password" json:"password"`
	LogoURL               string         `xml:"logo_url" json:"logo_url"`
	DraftStatus           string         `xml:"draft_status" json:"draft_status"`
	NumTeams              int            `xml:"num_teams" json:"num_teams"`
	EditKey               int            `xml:"edit_key" json:"edit_key"`
	WeeklyDeadline        string         `xml:"weekly_deadline" json:"weekly_deadline"`
	LeagueUpdateTimestamp string         `xml:"league_update_timestamp" json:"league_update_timestamp"`
	ScoringType           string         `xml:"scoring_type" json:"scoring_type"`
	LeagueType            string         `xml:"league_type" json:"league_type"`
	Renew                 string         `xml:"renew" json:"renew"`
	Renewed               string         `xml:"renewed" json:"renewed"`
	FeloTier              string         `xml:"felo_tier" json:"felo_tier"`
	IrisGroupChatID       string         `xml:"iris_group_chat_id" json:"iris_group_chat_id"`
	ShortInvitationURL    string         `xml:"short_invitation_url" json:"short_invitation_url"`
	AllowAddToDLExtraPos  int            `xml:"allow_add_to_dl_extra_pos" json:"allow_add_to_dl_extra_pos"`
	ProLeague             int            `xml:"is_pro_league" json:"pro_league"`
	CashLeague            int            `xml:"is_cash_league" json:"cash_league"`
	CurrentWeek           int            `xml:"current_week" json:"current_week"`
	StartWeek             int            `xml:"start_week" json:"start_week"`
	StartDate             string         `xml:"start_date" json:"start_date"`
	EndWeek               int            `xml:"end_week" json:"end_week"`
	EndDate               string         `xml:"end_date" json:"end_date"`
	Finished              int            `xml:"is_finished" json:"finished"`
	GameCode              string         `xml:"game_code" json:"game_code"`
	Season                int            `xml:"season" json:"season"`
	Settings              LeagueSettings `xml:"settings" json:"settings,omitempty"`
	// Standings             LeagueStandings    `xml:"standings" json:"standings,omitempty"`
	// ScoreBoard            LeagueScoreboard   `xml:"scoreboard" json:"scoreboard,omitempty"`
	Teams []Team `xml:"teams>team" json:"teams,omitempty"`
	// Players               []Player           `xml:"players>player" json:"players,omitempty"`
	// DraftResults          LeagueDraftResults `xml:"draft_results" json:"draft_results,omitempty"`
	// Transactions          []Transaction      `xml:"transactions>transaction" json:"transactions,omitempty"`
}

type LeagueSettings struct {
	XMLName                    xml.Name               `xml:"settings" json:"-"`
	DraftType                  string                 `xml:"draft_type" json:"draft_type"`
	AuctionDraft               int                    `xml:"is_auction_draft" json:"auction_draft"`
	ScoringType                string                 `xml:"scoring_type" json:"scoring_type"`
	PersistentURL              string                 `xml:"persistent_url" json:"persistent_url"`
	UsesPlayoff                int                    `xml:"uses_playoff" json:"uses_playoff"`
	HasPlayoffConsolationGames int                    `xml:"has_playoff_consolation_games" json:"has_playoff_consolation_games"`
	PlayoffStartWeek           int                    `xml:"playoff_start_week" json:"playoff_start_week"`
	PlayoffReseeding           int                    `xml:"uses_playoff_reseeding" json:"playoff_reseeding"`
	LockEliminatedTeams        int                    `xml:"uses_lock_eliminated_teams" json:"lock_eliminated_teams"`
	NumPlayoffTeams            int                    `xml:"num_playoff_teams" json:"num_playoff_teams"`
	NumPlayoffConsolationTeams int                    `xml:"num_playoff_consolation_teams" json:"num_playoff_consolation_teams"`
	MultiweekChampionship      int                    `xml:"has_multiweek_championship" json:"has_multiweek_championship"`
	WaiverType                 string                 `xml:"waiver_type" json:"waiver_type"`
	WaiverRule                 string                 `xml:"waiver_rule" json:"waiver_rule"`
	UsesFAAB                   int                    `xml:"uses_faab" json:"uses_faab"`
	DraftTime                  string                 `xml:"draft_time" json:"draft_time"`
	PostDraftPlayers           string                 `xml:"post_draft_players" json:"post_draft_players"`
	MaxTeams                   int                    `xml:"max_teams" json:"max_teams"`
	WaiverTime                 int                    `xml:"waiver_time" json:"waiver_time"`
	TradeEndDate               string                 `xml:"trade_end_date" json:"trade_end_date"`
	TradeRatifyType            string                 `xml:"trade_ratify_type" json:"trade_ratify_type"`
	TradeRejectTime            string                 `xml:"trade_reject_time" json:"trade_reject_time"`
	PlayerPool                 string                 `xml:"player_pool" json:"player_pool"`
	CantCutList                string                 `xml:"cant_cut_list" json:"cant_cut_list"`
	DraftTogether              int                    `xml:"draft_together" json:"draft_together"`
	SendbirdChannelURL         string                 `xml:"sendbird_channel_url" json:"sendbird_channel_url"`
	RosterPositions            []LeagueRosterPosition `xml:"roster_positions>roster_position" json:"roster_positions,omitempty"`
	StatCategories             LeagueStatCategories   `xml:"stat_categories" json:"stat_categories"`
	StatModifiers              LeagueStatModifiers    `xml:"stat_modifiers" json:"stat_modifiers"`
	MaxWeeklyAdds              int                    `xml:"max_weekly_adds" json:"max_weekly_adds"`
	Divisions                  []LeagueDivision       `xml:"divisions>division" json:"divisions,omitempty"`
	PickemEnabled              int                    `xml:"pickem_enabled" json:"pickem_enabled"`
	FractionalPoints           int                    `xml:"uses_fractional_points" json:"fractional_points"`
	NegativePoints             int                    `xml:"uses_negative_points" json:"negative_points"`
}

type LeagueRosterPosition struct {
	XMLName          xml.Name `xml:"roster_position" json:"-"`
	Position         string   `xml:"position" json:"position"`
	PositionType     string   `xml:"position_type" json:"position_type"`
	Count            int      `xml:"count" json:"count"`
	StartingPosition int      `xml:"is_starting_position" json:"starting_position"`
}

type LeagueStatCategories struct {
	XMLName xml.Name             `xml:"stat_categories" json:"-"`
	Stats   []LeagueStatCategory `xml:"stats>stat" json:"stats,omitempty"`
	Groups  []LeagueStatGroup    `xml:"groups>group" json:"groups,omitempty"`
}

type LeagueStatCategory struct {
	XMLName             xml.Name                         `xml:"stat" json:"-"`
	ID                  int                              `xml:"stat_id" json:"id"`
	Enabled             int                              `xml:"enabled" json:"enabled"`
	Name                string                           `xml:"name" json:"name"`
	DisplayName         string                           `xml:"display_name" json:"display_name"`
	Group               string                           `xml:"group" json:"group"`
	Abbr                string                           `xml:"abbr" json:"abbr"`
	SortOrder           int                              `xml:"sort_order" json:"sort_order"`
	PositionType        string                           `xml:"position_type" json:"position_type"`
	PositionTypes       []LeagueStatCategoryPositionType `xml:"stat_position_types>stat_position_type" json:"position_types,omitempty"`
	ExcludedFromDisplay int                              `xml:"is_excluded_from_display" json:"_excluded_from_display,omitempty"`
	OnlyDisplayStat     int                              `xml:"is_only_display_stat" json:"only_display_stat"`
}

type LeagueStatCategoryPositionType struct {
	XMLName         xml.Name `xml:"stat_position_type" json:"-"`
	PositionType    string   `xml:"position_type" json:"position_type"`
	OnlyDisplayStat int      `xml:"is_only_display_stat" json:"only_display_stat"`
}

type LeagueStatGroup struct {
	XMLName     xml.Name `xml:"group" json:"-"`
	Name        string   `xml:"group_name" json:"name"`
	DisplayName string   `xml:"group_display_name" json:"display_name"`
	Abbr        string   `xml:"group_abbr" json:"group_abbr"`
}

type LeagueStatModifiers struct {
	XMLName xml.Name             `xml:"stat_modifiers" json:"-"`
	Stats   []LeagueStatModifier `xml:"stats>stat" json:"stats,omitempty"`
}

type LeagueStatModifier struct {
	XMLName xml.Name                  `xml:"stat" json:"-"`
	ID      int                       `xml:"stat_id" json:"id"`
	Value   float64                   `xml:"value" json:"value"`
	Bonuses []LeagueStatModifierBonus `xml:"bonuses>bonus" json:"bonuses,omitempty"`
}

type LeagueStatModifierBonus struct {
	XMLName xml.Name `xml:"bonus" json:"-"`
	Target  int      `xml:"target" json:"target"`
	Points  float64  `xml:"points" json:"points"`
}

type LeagueDivision struct {
	XMLName xml.Name `xml:"division" json:"division"`
	ID      int      `xml:"division_id" json:"id"`
	Name    string   `xml:"name" json:"name"`
}

type LeagueStandings struct {
}
