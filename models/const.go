package models

// Space标签
const (
	SpaceNFT        string = "NFT"
	SpaceWEB3       string = "WEB3"
	SpaceGame       string = "Game"
	SpaceMetaVerse  string = "MetaVerse"
	SpaceDeFi       string = "DeFi"
	SpaceDAO        string = "DAO"
	SpaceLayer2     string = "Layer2"
	SpaceInvestment string = "Investment"
	SpaceInvesting  string = "Investing"
	SpaceTechnology string = "Technology"
)

// SpaceStatus Space 状态
type SpaceStatus int

const (
	SpaceStatusLive SpaceStatus = iota
	SpaceStatusScheduled
	SpaceStatusEnded
	SpaceStatusCanceled
)

// 数据状态
const (
	DataStatusDisable int = 0
	DataStatusEnable  int = 1
)
