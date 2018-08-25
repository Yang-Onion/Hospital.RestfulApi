package models

import (
	"time"
)

//体检详情
type PhysicalDetailViewModel struct {
	Id                       int       `json:"id"`
	RenWuId                  int       `josn:"renWuId"`
	BingRenId                int       `json:"bingRenId"`
	XinMin                   string    `json:"xinMin"`
	XinBie                   string    `json:"xinBie"`
	Nianlin                  string    `json:"nianlin"`
	ChuShengRiQi             time.Time `josn:"chuShengRiQi"`
	ShenFengZhengHao         string    `json:"shenFengZhengHao"`
	TiJianBianHao            string    `json:"tiJianBianHao"`
	KaiShiShiJian            time.Time `json:"kaiShiShiJian"`
	JieShuShiJian            time.Time `json:"jieShuShiJian"`
	IsDisabled               int       `json:"isDisabled"`
	CanKaoJianYi             []PhysicalConclusion
	PhysicalCatalogViewModel []PhysicalCatalogViewModel
}

//体检大项
type PhysicalCatalogViewModel struct {
	QingDanId              int    `json:"qingDanId"`
	TiJianXianMuId         int    `json:"tiJianXianMuId"`
	MinCheng               string `json:"minCheng"`
	ZhiXinKeShiId          int    `json:"zhiXinKeShiId"`
	ZhiXinKeShiMinCheng    string `json:"zhiXinKeShiMinCheng"`
	PhysicalItemsViewModel []PhysicalItemsViewModel
}

//体检小项及结果
type PhysicalItemsViewModel struct {
	QingDanId       int    `json:"qingDanId"`
	TiJianZhiBiaoId int    `json:"tiJianZhiBiaoId"`
	MinCheng        string `json:"minCheng"`
	JieGuo          string `json:"jieGuo"`
	CanKao          string `json:"canKao"`
	DanWei          string `json:"danWei"`
	BaoJing         string `json:"baoJing"`
	TiShi           string `json:"tiShi"`
}

//总检结论
type PhysicalConclusion struct {
	Id           int    `json:"id"`
	RenWuId      int    `json:"renWuId"`
	BingRenId    int    `json:"bingRenId"`
	JiLuXuHao    int    `json:"jiLuXuHao"`
	CanKaoJianYi string `json:"canKaoJianYi"`
}
