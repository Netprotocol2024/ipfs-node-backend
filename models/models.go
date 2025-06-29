package models

type StorageInfo struct {
	TotalGB          float64 `json:"totalGB"`
	UsedGB           float64 `json:"usedGB"`
	FreeGB           float64 `json:"freeGB"`
	UsedPercent      float64 `json:"usedPercent"`
	TotalUsedStorage float64 `json:"totalUsedStorage"`
}
