package excel

import (
	"github.com/xuri/excelize/v2"
	"guilt-type-service/internal/model"
	"guilt-type-service/internal/repository"
	
)

func ImportFromExcel(path string, repo repository.GuiltTypeRepository) error {
	f, err := excelize.OpenFile(path)
	if err != nil {
		return err
	}
	defer f.Close()

	sheet := f.GetSheetName(0)
	rows, err := f.GetRows(sheet)
	if err != nil {
		return err
	}

	var guiltTypes []model.GuiltType
	for i, row := range rows {
		if i == 0 { // skip header
			continue
		}
		if len(row) < 1 {
			continue
		}
		gt := model.GuiltType{
			Name: row[0],
		}
		if len(row) > 1 {
			gt.OtherInfo = row[1]
		}
		guiltTypes = append(guiltTypes, gt)
	}

	return repo.BulkInsert(guiltTypes)
}