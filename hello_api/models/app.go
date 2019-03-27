package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"github.com/astaxie/beego/logs"
)

type App struct {
	Id           int       `orm:"column(id);auto;pk"`
	AppKey       string    `orm:"column(app_key);size(32)"`
	AppSecret    string    `orm:"column(app_secret);size(32)"`
	Priority     int       `orm:"column(priority)"`
	CreatedTime  time.Time `orm:"column(created_time);type(datetime)"`
	UpdatedTime  time.Time `orm:"column(updated_time);type(datetime)"`
}

func (t *App) TableName() string{
	return "app"
}

func init()  {
	//orm.RegisterModel(new(App))
	orm.RegisterModelWithPrefix("xinzailing_",new(App))
}

func AddApp(app string, secret string, priority int) (int64,error){
	o := orm.NewOrm()
	v := &App{AppKey:app, AppSecret:secret, Priority:priority,CreatedTime:time.Now(),UpdatedTime:time.Now()}
	id, err := o.Insert(v)
	return id, err
}

func DeleteApp(app string)  {
	o := orm.NewOrm()
	res, err := o.Raw("delete from xinzailing_app where app_key = ?",app).Exec()
	if err != nil {
		logs.Error("DeleteAPP falied:",res)
	}
}

func UpdateApp(app string, secret string)  {
	o := orm.NewOrm()
	res, err := o.Raw("update xinzailing_app set app_secret =? where app_key = ?",secret,app).Exec()
	if err != nil {
		logs.Error("DeleteAPP falied:",res)
	}
}