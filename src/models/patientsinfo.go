package models

import (
	"time"
)

type PatientsInfo struct {
	BingRenId     int    `josn:"bingRenId"`
	MenZhengNo    int64  `json:"menZhengNo"`
	ZhuYuanHao    int    `json:"zhuYuanHao"`
	JiuZhengCaHao string `json:"jiuZhengCaHao"`
}
