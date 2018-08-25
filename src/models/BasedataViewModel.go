package models

type BaseDataViewModel struct {
	Department   []DepartmentViewModel
	TiJianXianMu []TiJiangXianMuViewModel
}

type DepartmentViewModel struct {
	Id       int    `json:"id"`
	MinCheng string `json:"minCheng"`
}

type TiJiangXianMuViewModel struct {
	Id            int    `json:"id"`
	MinCheng      string `json:"minCheng"`
	TiJianZhiBiao []TiJianZhiBiaoViewModel
}

type TiJianZhiBiaoViewModel struct {
	Id       int    `json:"id"`
	MinCheng string `json:"minCheng"`
}
