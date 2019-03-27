package kafka100

//协议版本定义
const (
	KAFKA_PROTO_VERSION_001 = 0x01000001 //v1.0.0版本协议
)

//卡夫卡数据类型
const (
	KAFKA_DT_DEVICE_ONLINE          = 1  //设备状态上线
	KAFKA_DT_DEVICE_OFFLINE         = 2  //设备状态下线
	KAFKA_DT_ALARM                  = 3  //报警
	KAFKA_DT_DISALARM               = 4  //消警
	KAFKA_DT_LIFT_MV_STATUS         = 5  //直梯触发数据
	KAFKA_DT_LIFT_RT_STATUS         = 6  //直梯实时数据
	KAFKA_DT_LIFT_SHAKE_STATUS      = 7  //直梯震动数据
	KAFKA_DT_GATEWAY_RT_STATUS      = 8  //局端网关实时数据
	KAFKA_DT_PHONE_MAC              = 9  //用户手机MAC地址上报
	KAFKA_DT_SCREEN_ADVERT_SHOT     = 10 //智能屏素材截屏
	KAFKA_DT_SCREEN_ADVERT_PLAYSTAT = 11 //智能屏广播播放统计
	KAFKA_DT_ZHEDA_ALARM_RECORD     = 12 //浙大告警录像完成上报
	KAFKA_DT_EAG_DEVICE_ONLINE      = 13 //Eag设备上线
	KAFKA_DT_EAG_DEVICE_OFFLINE     = 14 //Eag设备下线
)

//卡夫卡数据级别
const (
	KAFKA_LEVEL_NON_DISCARDING = 1 //不可丢弃的
	KAFKA_LEVEL_DISCARDING     = 2 //可丢弃的
)