package query

type FavoriteStatisticsRequest struct {
	All      bool `json:"all"`
	Gender   bool `json:"gender"`
	AgeStart int  `json:"age_start"`
	AgeEnd   int  `json:"age_end"`
}
