package service

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	excelize "github.com/xuri/excelize/v2"
	"kcers-survey/biz/infras/service/common"
)

func Export(tale []interface{}, resp []map[int]interface{}, name string) (string, error) {
	if name == "" {
		name = "列表导出"
	}
	exportFilePath, domain := common.ExportFilePath(name)

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			hlog.Errorf("关闭文件失败: %v", err)
		}
	}()
	sw, err := f.NewStreamWriter("Sheet1")
	if err != nil {
		fmt.Println(err)

	}
	err = sw.SetColWidth(1, 20, 16)
	if err != nil {
		hlog.Errorf("设置列宽失败: %v", err)

	}
	cell, err := excelize.CoordinatesToCellName(1, 1)
	if err != nil {
		hlog.Errorf("转换坐标失败: %v", err)
		return "", err
	}

	styleID, err := f.NewStyle(
		&excelize.Style{
			Font:      &excelize.Font{Color: "E83723", Size: 10, Family: "微软雅黑", Bold: true},
			Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center", WrapText: true},
		},
	)
	if err != nil {
		hlog.Errorf("设置行数据失败: %v", err)

	}
	styleID2, err := f.NewStyle(
		&excelize.Style{
			Font:      &excelize.Font{Color: "000000", Size: 9, Family: "Arial", Bold: false},
			Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center", WrapText: true},
		},
	)
	if err != nil {
		hlog.Errorf("设置行数据失败: %v", err)

	}
	var tales []interface{}
	for _, v := range tale {
		tales = append(tales, excelize.Cell{StyleID: styleID, Value: v})
	}

	if err := sw.SetRow("A1", tales,
		excelize.RowOpts{Height: 45, Hidden: false, OutlineLevel: 1}); err != nil {
		hlog.Errorf("设置行数据失败: %v", err)
	}

	for idx, row := range resp {
		cell, err = excelize.CoordinatesToCellName(1, idx+2)
		if err != nil {
			hlog.Errorf("转换坐标失败2: %v", err)
			return "", err
		}

		var r []interface{}

		for i := 1; i <= len(row); i++ {
			r = append(r, excelize.Cell{StyleID: styleID2, Value: row[i]})
		}
		//value := reflect.ValueOf(row)
		//for i := 0; i < value.NumField(); i++ {
		//	r = append(r, value.Field(i))
		//}

		err = sw.SetRow(cell, r)
		if err != nil {
			hlog.Errorf("设置行数据失败: %v", err)
			return "", err
		}
	}
	if err := sw.Flush(); err != nil {
		hlog.Errorf("Flush失败: %v", err)
	}

	//Save spreadsheet by the given path.
	if err := f.SaveAs(exportFilePath); err != nil {
		hlog.Errorf("保存文件失败: %v", err)
	}

	return domain, nil

}
