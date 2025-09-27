package request

type CreateDictTypeRequest struct {
	SystemFlag bool   `form:"systemFlag" json:"systemFlag"`
	DictCode   string `form:"dictCode" json:"dictCode"`
	DictName   string `form:"dictName" json:"dictName"`
	Remark     string `form:"remark" json:"remark"`
	Status     bool   `form:"status" json:"status"`
}
