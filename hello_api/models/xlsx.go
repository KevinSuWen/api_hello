package models

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"time"
	"github.com/astaxie/beego/logs"
	"fmt"
)

func main(){
	//time := time.Now().String()
	XlsCreat("Xin")

	XlsWrite("simline.xlsx",2,"Xin",21,23)
	XlsWrite("simline.xlsx",3,"Xin",26,23)
	XlsWrite("simline.xlsx",4,"Xin",14,79)
	XlsWrite("simline.xlsx",5,"Xin",78,23)
	XlsWrite("simline.xlsx",6,"Xin",22,23)
}

func XlsCreat(sheet string){
	xlsx := excelize.NewFile()
	//xlsx.NewSheet(sheet)
	xlsx.SetCellValue("Sheet1","A1","Time")
	xlsx.SetCellValue("Sheet1","B1","OnLine")
	xlsx.SetCellValue("Sheet1","C1","OffLine")
	xlsx.SetActiveSheet(1)
	//filename := fmt.Sprintf("%s.xlsx",time.Now().String())
	err := xlsx.SaveAs("simline.xlsx")
	if err != nil{
		logs.Error("create xlsx err:",err)
	}
	XlsWrite("simline.xlsx",3,"Xin",26,23)
}

func XlsWrite(xlsxname string,index int32, sheet string,onlineNum int32, offlineNum int32){
	xlsx, err := excelize.OpenFile(xlsxname)
	logs.Error("open xlsx xlsxname:",xlsxname)
	if err != nil {
		logs.Error("open xlsx err:",err)
		fmt.Println("444444444444444")
		return
	}
	xlsx.SetActiveSheet(1)
	logs.Error("open xlsx xlsxname-2:",sheet)
	xlsx.SetCellValue("Sheet1","A2",66)
	idx_a := fmt.Sprintf("A%s",index)
	idx_b := fmt.Sprintf("B%s",index)
	idx_c := fmt.Sprintf("C%s",index)
	xlsx.SetCellValue("Sheet1",idx_a,time.Now())
	xlsx.SetCellValue("Sheet1",idx_b,onlineNum)
	xlsx.SetCellValue("Sheet1",idx_c,offlineNum)
}