package models

type Province struct {
	PRV_ID       int64  `json:"PRV_ID" uri:"id" binding:"required"`
	DEP_ID       int64  `json:"DEP_ID"`
	PRV_NAME     string `json:"PRV_NAME"`
	PRV_UBIGEOUS string `json:"PRV_UBIGEOUS"`
	PRV_STATUS   bool   `json:"-"`
}
