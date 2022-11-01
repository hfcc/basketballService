package backendtool

import "time"
type NbaScheduleStruct struct {
	Standard []Standard `json:"standard"`
}
type Period struct {
	Current int `json:"current"`
	Type int `json:"type"`
	MaxRegular int `json:"maxRegular"`
}
type Nugget struct {
	Text string `json:"text"`
}
type HTeam struct {
	TeamID string `json:"teamId"`
	Score string `json:"score"`
	Win string `json:"win"`
	Loss string `json:"loss"`
}
type VTeam struct {
	TeamID string `json:"teamId"`
	Score string `json:"score"`
	Win string `json:"win"`
	Loss string `json:"loss"`
}
type Broadcasters struct {
	ShortName string `json:"shortName"`
	LongName string `json:"longName"`
}
type National struct {
	Broadcasters []Broadcasters `json:"broadcasters"`
}
type Video struct {
	RegionalBlackoutCodes string `json:"regionalBlackoutCodes"`
	IsLeaguePass bool `json:"isLeaguePass"`
	IsNationalBlackout bool `json:"isNationalBlackout"`
	IsTNTOT bool `json:"isTNTOT"`
	CanPurchase bool `json:"canPurchase"`
	IsVR bool `json:"isVR"`
	IsNextVR bool `json:"isNextVR"`
	IsNBAOnTNTVR bool `json:"isNBAOnTNTVR"`
	IsMagicLeap bool `json:"isMagicLeap"`
	IsOculusVenues bool `json:"isOculusVenues"`
	National National `json:"national"`
	Canadian []interface{} `json:"canadian"`
	SpanishNational []interface{} `json:"spanish_national"`
}
type Broadcast struct {
	Video Video `json:"video"`
}
type Watch struct {
	Broadcast Broadcast `json:"broadcast"`
}
type Standard struct {
	GameID string `json:"gameId"`
	SeasonStageID int `json:"seasonStageId"`
	GameURLCode string `json:"gameUrlCode"`
	StatusNum int `json:"statusNum"`
	ExtendedStatusNum int `json:"extendedStatusNum"`
	IsStartTimeTBD bool `json:"isStartTimeTBD"`
	StartTimeUTC time.Time `json:"startTimeUTC"`
	StartDateEastern string `json:"startDateEastern"`
	IsNeutralVenue bool `json:"isNeutralVenue"`
	StartTimeEastern string `json:"startTimeEastern"`
	IsBuzzerBeater bool `json:"isBuzzerBeater"`
	Period Period `json:"period"`
	Nugget Nugget `json:"nugget"`
	HTeam HTeam `json:"hTeam"`
	VTeam VTeam `json:"vTeam"`
	Watch Watch `json:"watch"`
}