package goiex

type AMC struct {
	Earnings
}

type BTO struct {
	Earnings
}

type Earnings struct {
	Symbol   string `json:"symbol"`
	Earnings []earnings
}

type EarningsToday struct {
	Bto BTO
	Amc AMC
}

type Quote struct {
	quote
}

type earnings struct {
	ActualEPS              float32 `json:"actualEPS"`
	ConcensusEPS           float32 `json:"consensusEPS"`
	EstimatedEPS           float32 `json:"estimatedEPS"`
	AnnounceTime           string  `json:"announceTime"`
	NumberOfEstimates      int32   `json:"numberOfEstimates"`
	EPSSurpriseDollar      float32 `json:"EPSSurpriseDollar"`
	EPSReportDate          string  `json:"EPSReportDate"` // make date type?
	FiscalPeriod           string  `json:"fiscalPeriod"`
	FiscalEndDate          string  `json:"fiscalEndDate"`
	YearAgo                float32 `json:"yearAgo"`
	YearAgoChangePercent   float32 `json:"yearAgoChangePercent"`
	EstimatedChangePercent float32 `json:"estimatedChangePercent"`
	SymbolId               int32   `json:"symbolId"`
}

type quote struct {
	Symbol                string
	CompanyName           string
	PrimaryExchange       string
	Sector                string
	CalculationPrice      string
	Open                  float32
	OpenTime              int64
	Close                 float32
	CloseTime             int64
	High                  float32
	Low                   float32
	LatestPrice           float32
	LatestSource          string
	LatestTime            string
	LatestUpdate          int64
	LatestVolume          int32
	IexRealtimePrice      float32
	IexRealtimeSize       int32
	IexLastUpdated        int64
	DelayedPrice          float32
	DelayedPriceTime      int64
	ExtendedPrice         float32
	ExtendedChange        float32
	ExtendedChangePercent float32
	ExtendedPriceTime     int64
	PreviousClose         float32
	Change                float32
	ChangePercent         float32
	IexMarketPercent      float32
	IexVolume             int32
	AvgTotalVolume        int32
	IexBidPrice           float32
	IexBidSize            int32
	IexAskPrice           float32
	IexAskSize            int32
	MarketCap             int64
	PeRatio               float32
	Week52High            float32
	Week52Low             float32
	YtdChange             float32
}
