package models

type NFTAssetMetadataItems struct {
	Blacklisted bool       `json:"blacklisted"`
	Blockchain  string     `json:"blockchain"`
	Ownership   *Ownership `json:"ownership,omitempty"`
	Offer       *Offer     `json:"offer,omitempty"`
}

type Ownership struct {
	Status   string  `json:"status"`
	PriceEth float64 `json:"priceEth"`
	Token    string  `json:"token"`
	TokenId  string  `json:"tokenId"`
}

type Properties struct {
	Name string `json:"name"`
}

type Offer struct {
	Active    bool    `json:"active"`
	BuyPrice  float64 `json:"buyPrice"`
	Cancelled bool    `json:"cancelled"`
	Completed bool    `json:"completed"`
	Id        string  `json:"id"`
}

type NFTAssetMetadata struct {
	Id         string                `json:"id"`
	Item       NFTAssetMetadataItems `json:"item"`
	Properties Properties            `json:"properties"`
}

// {
// 		"blacklisted": false,
// 		"blockchain": "ETHEREUM",
// 		"categories": [],
// 		"creator": "0x9b1a1d93cbc4c3a5584b7e58656e6dac00b4fcee",
// 		"deleted": false,
// 		"id": "0x6184f10302cebeea0211f9310225f051cc549626:788",
// 		"lazySupply": 0,
// 		"likes": 0,
// 		"owners": [],
// 		"ownership": {
// 				"blacklisted": false,
// 				"buyToken": "0x0000000000000000000000000000000000000000",
// 				"buyTokenId": "",
// 				"categories": [],
// 				"date": "2021-10-10T08:15:22.905+00:00",
// 				"hide": false,
// 				"id": "0x6184f10302cebeea0211f9310225f051cc549626:788:0x9b1a1d93cbc4c3a5584b7e58656e6dac00b4fcee",
// 				"lazyValue": 0,
// 				"likes": 0,
// 				"owner": "0x9b1a1d93cbc4c3a5584b7e58656e6dac00b4fcee",
// 				"pending": [],
// 				"price": 0.03,
// 				"priceEth": 0.03,
// 				"selling": 0,
// 				"sold": 0,
// 				"status": "FIXED_PRICE",
// 				"stock": 0,
// 				"token": "0x6184f10302cebeea0211f9310225f051cc549626",
// 				"tokenId": "788",
// 				"value": 1,
// 				"verified": false
// 		},
// 		"royalties": [],
// 		"sellers": 1,
// 		"supply": 1,
// 		"token": "0x6184f10302cebeea0211f9310225f051cc549626",
// 		"tokenId": "788",
// 		"totalStock": 1,
// 		"unlockable": false,
// 		"verified": false,
// 		"version": 18,
// 		"statuses": [
// 				"FIXED_PRICE"
// 		],
// 		"meta": {
// 				"imageMeta": {
// 						"height": 1200,
// 						"type": "image/png",
// 						"width": 1200
// 				}
// 		},
// 		"properties": {
// 				"attributes": [
// 						{
// 								"key": "Patterns",
// 								"value": "Leopard"
// 						},
// 						{
// 								"key": "Pupils",
// 								"value": "Goat"
// 						},
// 						{
// 								"key": "Body Color",
// 								"value": "Red"
// 						},
// 						{
// 								"key": "Eye Color",
// 								"value": "Blue"
// 						},
// 						{
// 								"key": "Pattern Color",
// 								"value": "Purple"
// 						},
// 						{
// 								"key": "Eye",
// 								"value": "Bored"
// 						},
// 						{
// 								"key": "Backlit",
// 								"value": "Orange"
// 						},
// 						{
// 								"key": "Antennae",
// 								"value": "Moth"
// 						},
// 						{
// 								"key": "Agents",
// 								"value": "Eta"
// 						},
// 						{
// 								"key": "Headwear",
// 								"value": "Bandana"
// 						},
// 						{
// 								"key": "Ear",
// 								"value": "Antenna"
// 						},
// 						{
// 								"key": "Clothes",
// 								"value": "Caveman"
// 						}
// 				],
// 				"description": "Galactic Secret Agency is a collection of 10.000 unique Alien Agents. They're lurking, watching and influencing humanity for centuries. Each Alien is unique, generated from over 170 hand drawn assets. Holding one of the agent grant you access to future members-only benefits. Visit our website to learn more about GSA. https://www.GalacticSecretAgency.com",
// 				"name": "Agent #788",
// 				"mediaEntries": [
// 						{
// 								"url": "https://galacticsecretagency.com/assets/assets/gsa/avatar/788.png",
// 								"dimension": {
// 										"width": 1200,
// 										"height": 1200
// 								},
// 								"sizeType": "ORIGINAL",
// 								"contentType": "IMAGE"
// 						},
// 						{
// 								"url": "https://img.rarible.com/prod/image/upload/t_image_preview/prod-itemImages/0x6184f10302cebeea0211f9310225f051cc549626:788/ad3e31ad",
// 								"dimension": {
// 										"width": 400,
// 										"height": 400
// 								},
// 								"sizeType": "PREVIEW",
// 								"contentType": "IMAGE",
// 								"cropMode": "FIT"
// 						},
// 						{
// 								"url": "https://img.rarible.com/prod/image/upload/t_image_big/prod-itemImages/0x6184f10302cebeea0211f9310225f051cc549626:788/ad3e31ad",
// 								"dimension": {
// 										"width": 1000,
// 										"height": 1000
// 								},
// 								"sizeType": "BIG",
// 								"contentType": "IMAGE",
// 								"cropMode": "FIT"
// 						},
// 						{
// 								"url": "https://img.rarible.com/prod/image/upload/t_image_mobile_low/prod-itemImages/0x6184f10302cebeea0211f9310225f051cc549626:788/ad3e31ad",
// 								"dimension": {
// 										"width": 375,
// 										"height": 0
// 								},
// 								"sizeType": "MOBILE_LOW",
// 								"contentType": "IMAGE",
// 								"cropMode": "FIT"
// 						},
// 						{
// 								"url": "https://img.rarible.com/prod/image/upload/t_image_mobile_medium/prod-itemImages/0x6184f10302cebeea0211f9310225f051cc549626:788/ad3e31ad",
// 								"dimension": {
// 										"width": 750,
// 										"height": 0
// 								},
// 								"sizeType": "MOBILE_MEDIUM",
// 								"contentType": "IMAGE",
// 								"cropMode": "FIT"
// 						},
// 						{
// 								"url": "https://img.rarible.com/prod/image/upload/t_image_mobile_high/prod-itemImages/0x6184f10302cebeea0211f9310225f051cc549626:788/ad3e31ad",
// 								"dimension": {
// 										"width": 1125,
// 										"height": 0
// 								},
// 								"sizeType": "MOBILE_HIGH",
// 								"contentType": "IMAGE",
// 								"cropMode": "FIT"
// 						},
// 						{
// 								"url": "https://galacticsecretagency.com/assets/assets/gsa/avatar/788.png",
// 								"dimension": {
// 										"width": 1200,
// 										"height": 1200
// 								},
// 								"sizeType": "SOURCE",
// 								"contentType": "IMAGE"
// 						}
// 				]
// 		},
// 		"carbonNegativeStatus": "NON_CARBON_NEGATIVE"
// },
