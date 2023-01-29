package models

type District struct {
	DIS_ID       int64  `json:"DIS_ID" uri:"id" binding:"required"`
	PRV_ID       int64  `json:"PRV_ID"`
	DIS_NAME     string `json:"DIS_NAME"`
	DIS_UBIGEOUS string `json:"DIS_UBIGEOUS"`
	DIS_STATUS   bool   `json:"-"`
}
