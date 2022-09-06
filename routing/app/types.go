package app

type CommonHttpRep struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

type GetShopListReq struct {
	ShopName   string `json:"shopName"`  // 店名
	RegionCode string `json:"regionId"`  // 地区号 废弃
	IsDuty     string `json:"isDuty"`    // 是否云值守
	IsServing  string `json:"isServing"` // 是否有人
	IsLike     string `json:"isLike"`    // 是否收藏
}

type GetShopListRep struct {
	Data []ShopInfoOutputData `json:"data"`
	CommonHttpRep
}

type ShopInfoOutputData struct {
	ShopId           string `json:"shopId"`           // 店id
	ShopName         string `json:"shopName"`         // 店名
	ShortAddress     string `json:"shortAddress"`     // 短地址
	RegionName       string `json:"regionName"`       // 地区名
	DutyOpt          string `json:"isDuty"`           // 是否云值守
	OnLineDeviceNum  string `json:"onLineDeviceNum"`  // 在线设备数
	OffLineDeviceNum string `json:"offLineDeviceNum"` // 离线设备数
	ServingNum       string `json:"servingNum"`       // 在线人数
	LikeOpt          string `json:"isLike"`           // 是否收藏
}

type GetShopDetailReq struct {
	ShopId string `json:"shopId"` // 店id
}

type GetShopDetailRep struct {
	Data ShopDetailData `json:"data"`
	CommonHttpRep
}

//type ShopDetailData map[string]string

type ShopDetailData struct {
	ShopName          string `json:"shopName"`          // 门店名
	Address           string `json:"shopAddress"`       // 店铺详细地址
	ContactName       string `json:"contactName"`       // 联系人姓名
	ContactPhone      string `json:"contactPhone"`      // 联系方式
	Contract          string `json:"contractUrl"`       // 合同扫描图片地址
	ContractBeginTime string `json:"contractBeginTime"` // 合同开始时间（10位数字时间戳）
	ContractEndTime   string `json:"contractEndTime"`   //	合同结束时间（10位数字时间戳）
	Status            string `json:"status"`            //	店铺状态（0-正常 1-异常）
}

type ShopLikeReq struct {
	ShopId  string `json:"shopId"` // 店id
	LikeOpt string `json:"isLike"` // 是否收藏
}

type GetDeviceListReq struct {
	ShopId string `json:"shopId"` // 店id
}

type GetDeviceListRep struct {
	Token string                 `json:"token"` // 视频播放token
	Data  []DeviceInfoOutputData `json:"data"`
	CommonHttpRep
}

type DeviceInfoOutputData struct {
	DeviceName     string `json:"deviceName"`     // 设备名
	DeviceSerial   string `json:"deviceSerial"`   // 设备序列号
	DevicePosition string `json:"devicePosition"` // 设备位置
	DeviceStatus   string `json:"deviceStatus"`   // 设备状态
	UpdateTime     string `json:"updateTime"`     // 状态更新时间
	VideoUrl       string `json:"videoUrl"`       // 视频链接
	ImageUrl       string `json:"imageUrl"`       // 当前截图
}

type LoginReq struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
type RegisterReq struct {
	ClientName string `json:"name"`
	Phone      string `json:"phone"`
	Password   string `json:"password"`
}

type LoginRep struct {
	Token string `json:"token"`
	CommonHttpRep
	Data ClientInfoData `json:"data"`
}
type ClientInfoData struct {
	ClientId    int64  `json:"clientId"`
	ClientName  string `json:"clientName"`
	ClientPhone string `json:"clientPhone"`
}
type ChangePasswordReq struct {
	PasswordNew string `json:"passwordNew"`
	PasswordOld string `json:"passwordOld"`
}

type ApplyDutyReq struct {
	ShopId  string `json:"shopId"`  // 店铺id
	OptFlag string `json:"optFlag"` // 操作标志 (0退出，1申请）
}

type CheckDutyReq struct {
	RelSerial string `json:"relSerial"` // 关联流水号
	Content   string `json:"content"`   // 异常原因
}

type ConfirmSuggestionReq struct {
	Phone   string `json:"phone"`   // 手机号
	Content string `json:"content"` // 投诉内容
}

type GetTaskStatusListRep struct {
	Data []TaskStatusInfoOutputData `json:"data"`
	CommonHttpRep
}

type TaskStatusInfoOutputData struct {
	ShopId       string `json:"shopId"`       // 门店id
	ShopName     string `json:"shopName"`     // 门店名
	Address      string `json:"shopAddress"`  // 店铺详细地址
	TaskSerialNo string `json:"taskSerialNo"` // 工单流水号
	TaskType     string `json:"taskType"`     // 工单类型（0-事中正常 1-事中异常 2-事后）
	TaskStatus   string `json:"status"`       // 设备位置
	CreateTime   string `json:"createTime"`   // 创建时间
	UpdateTime   string `json:"updateTime"`   // 状态更新时间
}

type GetMessageListReq struct {
	CurrentIndex int64 `json:"currentIndex"` // 当前最大索引
}
type GetMessageListRep struct {
	Data []MessageData `json:"data"`
	CommonHttpRep
}
type MessageData struct {
	MessageIndex int64  `json:"messageIndex"` // 消息类型（0-申请通知，1-预警通知）
	MessageType  string `json:"messageType"`  // 消息类型（0-申请通知，1-预警通知）
	Topic        string `json:"topic"`        // 主题
	Content      string `json:"content"`      // 内容
	KeyString    string `json:"keyString"`    // 点击链接跳转
	CreateTime   string `json:"createTime"`   // 创建时间
	SendTime     string `json:"sendTime"`     // 发送时间
}

type ShopVisitRankRep struct {
	Data []ShopVisitRankOutputData `json:"data"`
	CommonHttpRep
}
type ShopVisitRankOutputData struct {
	ShopName string `json:"shopName"` // 门店名
	VisitNum string `json:"visitNum"` // 在店人数
}

type StatVisitSumReq struct {
	ShopId     string `json:"shopId"`     // 门店id 门店查询类型为1时，必输
	SearchType string `json:"searchType"` // 门店查询类型（0-全部门店，1-具体门店）
	PeriodType string `json:"periodType"` // 时间段类型（1-今天，2-昨天，3-本周，4-本月）
}

type StatVisitSumRep struct {
	Data VisitSumData `json:"data"`
	CommonHttpRep
}

type VisitSumData struct {
	TotalVisit     string `json:"totalVisit"`     // 总客流 整数
	PayVisit       string `json:"payVisit"`       // 付款客流 整数
	VisitFrequency string `json:"visitFrequency"` // 进入频率 整数
	DiffRate       string `json:"diffRate"`       // 较上期 单位%
}

type StatVisitChartReq struct {
	ShopId     string `json:"shopId"`     // 门店id 门店查询类型为1时，必输
	SearchType string `json:"searchType"` // 门店查询类型（0-全部门店，1-具体门店）
	StepType   int64  `json:"stepType"`   // 步幅类型（0-时，1-日，2-周，3-月）
}

type StatVisitChartRep struct {
	Data []ChartData `json:"data"`
	CommonHttpRep
}

type ChartData struct {
	Label string `json:"label"` //	统计标签，按日统计，为时分，按7d/30d统计为月日, 按年统计为月份
	Value string `json:"value"` //	数量
}

type StatIncomeSumReq struct {
	ShopId     string `json:"shopId"`     // 门店id 门店查询类型为1时，必输
	SearchType string `json:"searchType"` // 门店查询类型（0-全部门店，1-具体门店）
	PeriodType string `json:"periodType"` // 时间段类型（1-今天，2-昨天，3-本周，4-本月）
}

type StatIncomeSumRep struct {
	Data IncomeSumData `json:"data"`
	CommonHttpRep
}

type IncomeSumData struct {
	TotalIncome    string `json:"totalIncome"`    // 总营收 数字到分
	PayCount       string `json:"payCount"`       // 付款单数 付款单数
	PayPerCustomer string `json:"payPerCustomer"` // 客单价	客单价，数据到分
	DiffRate       string `json:"diffRate"`       // 较上期 单位%
}

type StatIncomeChartReq struct {
	ShopId     string `json:"shopId"`     // 门店id 门店查询类型为1时，必输
	SearchType string `json:"searchType"` // 门店查询类型（0-全部门店，1-具体门店）
	StepType   int64  `json:"stepType"`   // 步幅类型（0-时，1-日，2-周，3-月）
}

type StatIncomeChartRep struct {
	Data []ChartData `json:"data"`
	CommonHttpRep
}

type Stat24hTransReq struct {
	ShopId string `json:"shopId"`
}

type Stat24hTransRep struct {
	Data Stat24hTransData `json:"data"`
	CommonHttpRep
}

type Stat24hTransData struct {
	TransNum    string `json:"transNum"`
	TransAmount string `json:"transAmount"`
}
