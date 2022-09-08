package dict

type DeviceKind string
type DeviceStatus string
type DevicePosition string
type PayType string
type PayStatus string
type UserRole int64
type PeriodType string
type StepType int64
type TaskStatus string
type TaskType int64
type Channel string
type EventType string
type EventStatus string
type EventLevel string
type BillCheckStatus string
type TransStatus string
type ShopStatus string

const (
	AppChannel     = Channel("0")
	MiniAppChannel = Channel("1")
	WindowsChannel = Channel("2")
	DemoChannel    = Channel("3")

	SubIPCDevice   = DeviceKind("1") //附属摄像头(ipc)
	NVRDevice      = DeviceKind("2") //录像机 nvr
	CashDevice     = DeviceKind("3") //收银机
	VoiceIPCDevice = DeviceKind("4") //语音摄像头

	// 设备位置 1-监控区1  2-监控区2  3-监控区3  4-监控区4  5-监控区5  6-监控区6  7-主监控   8-大门   9-贵重区
	PositionArea1     = DevicePosition("1") // 监控区1
	PositionArea2     = DevicePosition("2") // 监控区2
	PositionArea3     = DevicePosition("3") // 监控区3
	PositionArea4     = DevicePosition("4") // 监控区4
	PositionArea5     = DevicePosition("5") // 监控区5
	PositionArea6     = DevicePosition("6") // 监控区6
	PositionMain      = DevicePosition("7") // 主监控
	PositionDoor      = DevicePosition("8") // 大门
	PositionImportant = DevicePosition("9") // 贵重区

	DeviceStatusOk  = DeviceStatus("0") //正常
	DeviceStatusErr = DeviceStatus("1") //异常
	DeviceStatusOff = DeviceStatus("2") //离线

	WeChatPay = PayType("1") // 微信支付
	AliPay    = PayType("2") // 支付宝
	UnionPay  = PayType("3") // 云闪付

	ReadyToPay = PayStatus("0") // 待支付
	PaySuccess = PayStatus("1") // 支付成功
	PayFail    = PayStatus("2") // 支付失败

	RoleApproveUser = UserRole(1) // 1-审核员
	RoleDutyUser    = UserRole(2) // 2-云值守员
	RoleServingUser = UserRole(3) // 3-工客服人员
	RoleDeveloper   = UserRole(8) // 8-开发人员
	RoleSuperAdmin  = UserRole(9) // 9-超级管理员

	StatToday     = PeriodType("1") // 今天
	StatYesterday = PeriodType("2") // 昨天
	StatThisWeek  = PeriodType("3") // 本周
	StatThisMonth = PeriodType("4") // 本月

	StepHour  = StepType(0)
	StepDay   = StepType(1)
	StepWeek  = StepType(2)
	StepMonth = StepType(3)
	//值守状态（0-待审核 1-值守中 2-下线待确认 3-值守完成，8-异常处理中 9-完成）
	TaskCreateReadyToApprove     = TaskStatus("0")
	TaskApprovedAndOnDuty        = TaskStatus("1")
	TaskCancelDutyReadyToConfirm = TaskStatus("2")
	TaskDutyFinished             = TaskStatus("3")
	TaskHandlingEvents           = TaskStatus("8")
	TaskFinished                 = TaskStatus("9")
	//工单类型（0-事中正常 1-事中异常 2-事后）
	TaskInProcessNormal = TaskType(0)
	TaskInProcessError  = TaskType(1)
	TaskAfterwards      = TaskType(2)

	TaskEvent  = EventType("1") // 工单异常
	TransEvent = EventType("2") // 交易异常
	BillEvent  = EventType("3") // 订单异常

	EventCreated  = EventStatus("0") // 待处理
	EventHandling = EventStatus("1") // 处理中
	EventFinished = EventStatus("2") // 完成

	EventNormal    = EventLevel("1") // 普通
	EventImportant = EventLevel("2") // 重要
	EventEmergency = EventLevel("3") // 紧急

	BillReadyToCheck = BillCheckStatus("0") // 待审核
	BillCheckOk      = BillCheckStatus("1") // 审核通过
	BillCheckReject  = BillCheckStatus("2") // 审核拒绝

	TransInProgress          = TransStatus("0")  // 0-购物中
	TransFinishedWithoutBill = TransStatus("1")  // 1-完成（无购物离店）
	TransFinished            = TransStatus("2")  // 2-完成（购物离店）
	TransFinishedWithEvent   = TransStatus("-1") // -1-异常

	ShopStatusOk  = ShopStatus("0") // 0-正常
	ShopStatusErr = ShopStatus("1") // 1-异常

)

func GetDevicePosition(positionStr string) string {
	position := DevicePosition(positionStr)
	if position == PositionArea1 {
		return "监控区1"
	} else if position == PositionArea2 {
		return "监控区2"
	} else if position == PositionArea3 {
		return "监控区3"
	} else if position == PositionArea4 {
		return "监控区4"
	} else if position == PositionArea5 {
		return "监控区5"
	} else if position == PositionArea6 {
		return "监控区6"
	} else if position == PositionMain {
		return "主监控"
	} else if position == PositionDoor {
		return "大门"
	} else if position == PositionImportant {
		return "贵重区"
	}
	return ""
}
