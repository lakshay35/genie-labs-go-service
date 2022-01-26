package models

type Filter struct {
	VerifiedOnly bool     `json:"verifiedOnly"`
	Sort         string   `json:"sort"`
	Collections  []string `json:"collections"`
	Traits       []string `json:"traits"`
	NSFW         bool     `json:"nsfw"`
}

type GetNFTAssetsPayload struct {
	Filter Filter `json:"filter"`
	Size   int    `json:"size"`
}

type Asset struct {
	Id string `json:"id"`
}
