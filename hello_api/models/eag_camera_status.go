package models

import (
	"github.com/astaxie/beego/orm"
)

type CameraOnlineStat struct {
	Id           int       `orm:"column(id);auto;pk"`
	BeforeTime   string    `orm:"column(beforetime);size(32)"`
	AfterTime    string    `orm:"column(aftertime);size(32)"`
	OnlineNum    int    	`orm:"column(onlinenumber)"`
	OfflineNum   int 		`orm:"column(offlinenumber)"`
}

func (t *CameraOnlineStat) TableName() string{
	return "CameraOnlineStat"
}

func init()  {
	orm.RegisterModelWithPrefix("eag_",new(CameraOnlineStat))
}

func AddCamerOnlineStatData(beforetime string, aftertime string, onlinenumber int, offlinenumber int) (int64,error){
	o := orm.NewOrm()
	v := &CameraOnlineStat{BeforeTime:beforetime, AfterTime:aftertime, OnlineNum:onlinenumber,OfflineNum:offlinenumber}
	id, err := o.Insert(v)
	return id, err
}