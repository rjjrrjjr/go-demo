package export

import (
	"encoding/csv"
	"os"
	"testing"
)

func TestCreateCVS(t *testing.T) {
	// 创建 CSV 文件并写入数据
	file, err := os.Create("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// 写入表头
	header := []string{"姓名", "年龄"}
	if err := writer.Write(header); err != nil {
		panic(err)
	}

	// 写入数据行
	row1 := []string{"张三", "20"}
	if err := writer.Write(row1); err != nil {
		panic(err)
	}

	row2 := []string{"李四", "25"}
	if err := writer.Write(row2); err != nil {
		panic(err)
	}

	row3 := []string{"王二麻子", "25"}
	if err := writer.Write(row3); err != nil {
		panic(err)
	}
}
