package model

// Product holds "product" info from an exchange
type Product struct {
	Asset          string `json:"asset"`
	Currency       string `json:"currency"`
	MinSize        string `json:"min_size"`
	MaxSize        string `json:"max_size"`
	Increment      string `json:"increment"`
	AssetIncrement string `json:"asset_increment"`
	Label          string `json:"label"`
}
