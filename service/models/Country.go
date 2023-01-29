package models

type Country struct {
	CNT_ID     int64  `json:"CNT_ID" uri:"id" binding:"required"`
	CNT_NAME   string `json:"CNT_NAME"`
	CNT_CODE   string `json:"CNT_CODE"`
	CNT_STATUS bool   `json:"-"`
}
