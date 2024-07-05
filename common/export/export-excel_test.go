package export

import (
	"fmt"
	"github.com/tealeg/xlsx/v3"
	"github.com/xuri/excelize/v2"
	"testing"
)

/**
在Go语言中，标准库并没有直接提供导出Excel文件的功能。如果你不使用第三方库，将会更加复杂和繁琐。

要实现无第三方库导出Excel，你可以手动操作文件和格式，将数据以Excel支持的格式写入文件。但这需要对Excel文件格式和结构有较深入的了解，并编写大量的代码来生成正确的Excel文件。

考虑到这种情况，通常建议使用可靠的第三方库，如popular excelize、tealeg/xlsx或go-echarts/excel等，它们已经为我们封装了复杂的底层操作，使得导出Excel变得简单和高效

如果需要更多的功能和灵活性，popular excelize是一个不错的选择。如果只需要基本的读写能力且性能要求较高，tealeg/xlsx可能更适合。而如果你打算与Go-Echarts图表库集成来生成图表，并且数据源是Excel文件，则可以考虑使用go-echarts/excel库
*/

func TestCreateExcelTealeg(t *testing.T) {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf("Failed to add sheet: %s\n", err)
		return
	}
	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = "Hello"
	cell = row.AddCell()
	cell.Value = "World"

	err = file.Save("example.xlsx")
	if err != nil {
		fmt.Printf("Failed to save file: %s\n", err)
		return
	}
	fmt.Println("Excel file created and saved.")
}

func TestCreateExcelPopularExcelize(t *testing.T) {
	f := excelize.NewFile()
	data := [][]interface{}{
		{"ID", "Name", "Age"},
		{1, "John Doe", 30},
		{2, "Jane Smith", 25},
		{3, "Bob Johnson", 35},
	}
	sheetName := "Sheet1"
	for row, rowData := range data {
		for col, cellData := range rowData {
			cell, _ := excelize.CoordinatesToCellName(col+1, row+1)
			fmt.Println(cell)
			_ = f.SetCellValue(sheetName, cell, cellData)
		}
	}

	// 保存文件
	err := f.SaveAs("sample.xlsx")
	if err != nil {
		fmt.Println("Save failed:", err)
		return
	}

	fmt.Println("Excel file created and saved.")
}

func TestCreateExcel(t *testing.T) {
	f := excelize.NewFile()

	// 创建一个新的sheet

	// 设置单元格的值
	_ = f.SetCellValue("Sheet1", "A1", "姓名")
	_ = f.SetCellValue("Sheet1", "B1", "年龄")
	_ = f.SetCellValue("Sheet1", "A2", "张三")
	_ = f.SetCellValue("Sheet1", "B2", 28)
	_ = f.SetCellValue("Sheet1", "A3", "李四")
	_ = f.SetCellValue("Sheet1", "B3", 32)

	_ = f.SetCellValue("Sheet2", "A2", "张三")
	_ = f.SetCellValue("Sheet2", "B2", 28)
	_ = f.SetCellValue("Sheet2", "A3", "李四")
	_ = f.SetCellValue("Sheet2", "B3", 32)

	// 保存文件
	err := f.SaveAs("example.xlsx")
	if err != nil {
		fmt.Println("保存Excel时发生错误：", err)
	} else {
		fmt.Println("Excel文件保存成功！")
	}
}
