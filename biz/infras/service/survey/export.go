package survey

import (
	"github.com/pkg/errors"
	service2 "kcers-survey/biz/infras/service"
	"kcers-survey/idl_gen/model/service"
)

func (s Survey) ListResponseExport(req *service.ResponseListReq) (string, error) {
	//TODO implement me

	req.Page = 1
	req.PageSize = 100000
	resp, total, _ := s.ListResponse(req)

	if total == 0 {
		return "", errors.New("暂无数据")
	}
	tale := []interface{}{
		"编号",
		"受访人",
		"受访人联系电话",
		"调研员",
		"调研员",
		"填写问卷时间",
		"完成度",
		"省",
		"市（州）",
		"县（区、旗）",
		"乡（镇）",
		"村",
	}
	var list []map[int]interface{}
	for _, row := range resp {
		list = append(list, map[int]interface{}{
			1:  row.Sn,
			2:  row.Respondent,
			3:  row.RespondentPhone,
			4:  row.Researcher,
			5:  row.ResearcherPhone,
			6:  row.CreatedAt,
			7:  row.AnswerCount,
			8:  row.Area,
			9:  row.City,
			10: row.District,
			11: row.Village,
			12: row.Address,
		})

	}
	domain, err := service2.Export(tale, list, "")
	if err != nil {
		return "", err
	}
	return domain, nil
}
