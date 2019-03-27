package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
	"hello_api/kafka/pb"
	"github.com/golang/protobuf/proto"
	"hello_api/global"

	"time"
	"github.com/astaxie/beego"
	"hello_api/models"
)

type DeviceConsumer struct {
	KfkConsumer
}

func init(){
	global.G_OffLine = 0
	global.G_OnLine = 0
	global.G_Time = time.Now().Unix()
	period, err:= beego.AppConfig.Int("statPeriod")
	if nil != err{
		logs.Error("get statPeriod failed")
		global.G_Period = 600
	}
	global.G_Period = period
}
func newBasicConsumer(name string, broker []string, topics []string) (*DeviceConsumer){
	return &DeviceConsumer{
		KfkConsumer{name: name, brokers: broker, topics: topics},
	}
}

func (ths DeviceConsumer) ProcessData(msg *sarama.ConsumerMessage) error{
	data := msg.Value
	notify := &kafka100.PushData{}
	err := proto.Unmarshal(data, notify)
	if err != nil {
		logs.Error("parse data error:%s, data:%s", err.Error(), string(data))
		return nil
	}
	logs.Debug("basic notify message,data:%s", string(data))

	if notify.DataType == kafka100.KAFKA_DT_EAG_DEVICE_ONLINE{
		deviceOnline := &kafka100.EagDeviceOnline{}
		err = proto.Unmarshal(notify.Data, deviceOnline)
		if err != nil {
			logs.Error("parse data error:%s,data:%s", err.Error(), string(data))
			return nil
		}
		if uint32(kafka100.ONLINE_EXT_TYPE_EAGDEV) == deviceOnline.ExtType {
			logs.Notice("deviceID[online]:%s",deviceOnline.DeviceId)
			global.G_OnLine ++
			camExt := &kafka100.CamEagOnlineExtInfo{}
			if proto.Unmarshal(deviceOnline.ExtInfo, camExt) == nil{
				logs.Info("Device = %+v;",camExt)
			}else{
				logs.Error("online unmarshal failed")
			}
		}
	}

	if notify.DataType == kafka100.KAFKA_DT_EAG_DEVICE_OFFLINE{
		deviceOffline := &kafka100.EagDeviceOffline{}
		err = proto.Unmarshal(notify.Data, deviceOffline)
		if err != nil {
			logs.Error("parse data error:%s,data:%s", err.Error(), string(data))
			return nil
		}else{
			logs.Notice("deviceID[offline]:%s",deviceOffline.DeviceId)
			global.G_OffLine++
		}
	}
	nowtime := time.Now().Unix()
	if (nowtime - global.G_Time) > int64(global.G_Period){

		timeLayout := "2006-01-02 15:04:05"
		before_datetime := time.Unix(global.G_Time, 0).Format(timeLayout)
		after_datetime := time.Unix(time.Now().Unix(), 0).Format(timeLayout)

		logs.Warn("time(before):%s--(after)%s: online:%d; offline:%d",before_datetime,after_datetime,global.G_OnLine,global.G_OffLine)
		models.AddCamerOnlineStatData(before_datetime,after_datetime,global.G_OnLine,global.G_OffLine)
		global.G_Time = time.Now().Unix()
		global.G_OnLine = 0
		global.G_OffLine = 0
	}
	return nil
}
