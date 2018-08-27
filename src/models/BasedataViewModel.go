package models

//基础信息
type BaseDataViewModel struct {
	Department   []DepartmentViewModel
	TiJianXianMu []TiJiangXianMuViewModel
}

//部门
type DepartmentViewModel struct {
	Id       int    `json:"id"`
	MinCheng string `json:"minCheng"`
}

//体检项目
type TiJiangXianMuViewModel struct {
	Id            int    `json:"id"`
	MinCheng      string `json:"minCheng"`
	TiJianZhiBiao []TiJianZhiBiaoViewModel
}

//体检指标
type TiJianZhiBiaoViewModel struct {
	Id       int    `json:"id"`
	MinCheng string `json:"minCheng"`
}
