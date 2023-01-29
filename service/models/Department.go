package models

type Department struct {
	DEP_ID       int64  `json:"DEP_ID" uri:"id" binding:"required"`
	CNT_ID       int64  `json:"CNT_ID"`
	DEP_NAME     string `json:"DEP_NAME"`
	DEP_UBIGEOUS string `json:"DEP_UBIGEOUS"`
	DEP_STATUS   bool   `json:"-"`
}
