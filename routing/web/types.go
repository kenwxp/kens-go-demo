package web

type CommonHttpRep struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}
type LoginReq struct {
	AccNo    string `json:"accNo"`
	Password string `json:"password"`
}

type LoginRep struct {
	Token    string       `json:"token"`
	UserInfo UserInfoData `json:"userInfo"`
	CommonHttpRep
}

type UserInfoData struct {
	AccNo     string `json:"accNo"`     // 工号、登录账号	pk
	AccName   string `json:"accName"`   // 工号、登录账号	pk
	Phone     string `json:"phone"`     // 手机号
	Email     string `json:"email"`     // 邮箱
	RoleId    string `json:"roleId"`    // 角色权限（1-管理员 2-云值守员 3-工客服人员）
	OnWork    string `json:"onWork"`    // 在岗状态（0-离岗 1-在岗）
	LogonTime string `json:"logonTime"` // 登录时间（10位数字时间戳）
}
type ChangePasswordReq struct {
	PasswordNew string `json:"passwordNew"`
	PasswordOld string `json:"passwordOld"`
}

type GetClientListReq struct {
	ClientName string `json:"clientName"` // 客户名
	Phone      string `json:"phone"`      // 电话号码
}
type GetClientListRep struct {
	DataLen int64                  `json:"dataLen"`
	Data    []ClientListOutputData `json:"data"`
	CommonHttpRep
}
type ClientListOutputData struct {
	ClientId   int64  `json:"clientId"`   // 客户id
	ClientName string `json:"clientName"` // 客户名
	Phone      string `json:"phone"`      // 手机
	CreateTime string `json:"createTime"` // 创建时间（10位数字时间戳）
	UpdateTime string `json:"updateTime"` // 创建时间（10位数字时间戳）
	LogonTime  string `json:"logonTime"`  // 登录时间（10位数字时间戳）
	IsValid    int64  `json:"isValid"`    // 启用标志(0-启用 1-禁用)
}
type AddClientReq struct {
	ClientName string `json:"clientName"`
	Phone      string `json:"phone"`
}
type DeleteClientReq struct {
	ClientId int64 `json:"clientId"` // 客户id
}
type EditClientReq struct {
	ClientId int64 `json:"clientId"` // 客户id
	AddClientReq
}
type ResetClientPasswordReq struct {
	ClientId int64 `json:"clientId"` // 客户id
}

type ShopIdReq struct {
	ShopId string `json:"shopId"` // 客户id
}
type GetShopListWithConditionRep struct {
	DataLen int64                `json:"dataLen"`
	Data    []ShopListOutputData `json:"data"`
	CommonHttpRep
}
type ShopListReq struct {
	ShopId     int64  `json:"shopId"`
	ClientId   int64  `json:"clientId"`
	ClientName string `json:"clientName"`
	ShopName   string `json:"shopName"`
	RegionCode string `json:"regionCode"` // 区域
	Address    string `json:"address"`
	Status     string `json:"status"`
}
type ShopListOutputData struct {
	ShopId            int64  `json:"shopId"`            // 门店id
	ShopName          string `json:"shopName"`          // 门店名
	RegionCode        string `json:"regionCode"`        // 区域
	RegionName        string `json:"regionName"`        // 区域名
	ContactName       string `json:"contactName"`       // 联系人姓名
	ContactPhone      string `json:"contactPhone"`      // 联系方式
	Address           string `json:"address"`           // 地址
	ShortAddress      string `json:"shortAddress"`      // 短地址
	MaxReception      int64  `json:"maxReception"`      // 最大在店人数
	Area              string `json:"area"`              // 面积
	Contract          string `json:"contract"`          // 合同扫描图片地址
	ContractBeginTime string `json:"contractBeginTime"` // 合同开始时间（10位数字时间戳）
	ContractEndTime   string `json:"contractEndTime"`   // 合同结束时间（10位数字时间戳）
	CreateTime        string `json:"createTime"`        // 创建时间（10位数字时间戳）
	UpdateTime        string `json:"updateTime"`        // 更新时间（10位数字时间戳）
	IsValid           int64  `json:"isValid"`           // 启用标志(0-启用 1-禁用)
	ReceptionNum      int64  `json:"receptionNum"`      // 在店人数
	DutyOpt           string `json:"dutyOpt"`           // 云值守状态（0-否 1-是）
	DutyOnTime        string `json:"dutyOnTime"`        // 云值守时间 （10位数字时间戳）
	DutyOffTime       string `json:"dutyOffTime"`       // 取消云值守时间（10位数字时间戳）
	Status            string `json:"status"`            // 运行状态（0-正常 1-异常 2-离线 ）
	LikeOpt           string `json:"likeOpt"`           // 收藏操作 （0-否 1-是）
	LikeTime          string `json:"likeTime"`          // 收藏时间
	Remark            string `json:"remark"`            // 备注
	ClientId          int64  `json:"clientId"`          // 客户id
	ClientName        string `json:"clientName"`        // 客户名
	ClientPhone       string `json:"clientPhone"`       // 客户手机
	DeviceNum         string `json:"deviceNum"`         // 设备数
}
type AddShopReq struct {
	ShopName          string `json:"shopName"`          // 门店名
	RegionCode        string `json:"regionCode"`        // 地区号
	ClientId          int64  `json:"clientId"`          // 所属客户
	ContactName       string `json:"contactName"`       // 联系人姓名
	ContactPhone      string `json:"contactPhone"`      // 联系方式
	Address           string `json:"address"`           // 店铺详细地址
	ShortAddress      string `json:"shortAddress"`      // 店铺短地址
	MaxReception      int64  `json:"maxReception"`      // 最大在店人数
	Area              string `json:"area"`              // 店铺面积
	Contract          string `json:"contract"`          // 合同扫描图片地址
	ContractBeginTime string `json:"contractBeginTime"` // 合同开始时间（10位数字时间戳）
	ContractEndTime   string `json:"contractEndTime"`   //	合同结束时间（10位数字时间戳）
}
type EditShopReq struct {
	ShopId int64 `json:"shopId"` // 商店id
	AddShopReq
}

type DeviceIdReq struct {
	DeviceId int64 `json:"deviceId"` // 设备id
}

type GetDeviceListReq struct {
	ShopId     int64  `json:"shopId"`
	ShopName   string `json:"shopName"`
	ClientName string `json:"clientName"`
	DeviceName string `json:"deviceName"`
	DeviceKind string `json:"deviceKind"`
	Status     string `json:"status"`
}

type GetDeviceListRep struct {
	DataLen int64                  `json:"dataLen"`
	Data    []DeviceListOutputData `json:"data"`
	CommonHttpRep
}
type DeviceListOutputData struct {
	DeviceId       int64  `json:"deviceId"`       // 自增	设备id
	DeviceName     string `json:"deviceName"`     // 设备名
	Model          string `json:"model"`          // 设备型号
	DevicePosition string `json:"devicePosition"` // 设备位置
	DeviceKind     string `json:"deviceKind"`     // 种类（1-摄像头(ipc) 2-nvr 3-门禁卡 4-收银机）
	BridgeDevice   int64  `json:"bridgeDevice"`   // 桥接设备，ipc设置nvr的设备id
	ChannelNo      int64  `json:"channelNo"`      // ipc 桥接nvr的通道号
	Brand          string `json:"brand"`          // 品牌方
	DeviceSerial   string `json:"deviceSerial"`   // 设备序列号
	ValidateCode   string `json:"validateCode"`   // 设备验证码
	CreateTime     string `json:"CreateTime"`     // 创建时间（10位数字时间戳）
	UpdateTime     string `json:"updateTime"`     // 更新时间（10位数字时间戳）
	IsValid        int64  `json:"isValid"`        // 启用标志(0-启用 1-禁用)
	AddOpt         string `json:"addOpt"`         // 上云操作（0-否 1-是）
	AddOnTime      string `json:"addOnTime"`      // 设备上云时间（10位数字时间戳）
	AddOffTime     string `json:"addOffTime"`     // 设备取消上云时间（10位数字时间戳）
	Status         string `json:"status"`         // 运行状态（0-正常 1-异常 2-离线 ）
	Remark         string `json:"remark"`         // 备注
	ShopId         int64  `json:"shopId"`         // 门店id
	ShopName       string `json:"shopName"`       // 门店名
	ClientName     string `json:"clientName"`     // 客户名
}

type GetShopDeviceRelListReq struct {
	ShopId  int64 `json:"shopId"`
	RelFlag int64 `json:"relFlag"` //关联状态（0-待关联 2-已关联)
}

type GetShopDeviceRelListRep struct {
	DataLen int64                   `json:"dataLen"`
	Data    []ShopDeviceRelListItem `json:"data"`
	CommonHttpRep
}
type ShopDeviceRelListItem struct {
	DeviceId       int64  `json:"deviceId"`       // 自增	设备id
	DeviceName     string `json:"deviceName"`     // 设备名
	Model          string `json:"model"`          // 设备型号
	DevicePosition string `json:"devicePosition"` // 设备位置
	DeviceKind     string `json:"deviceKind"`     // 种类（1-摄像头(ipc) 2-nvr 3-收银机）
	BridgeDevice   int64  `json:"bridgeDevice"`   // 桥接设备，ipc设置nvr的设备id
	ChannelNo      int64  `json:"channelNo"`      // ipc 桥接nvr的通道号
	Brand          string `json:"brand"`          // 品牌方
}

type AddDeviceReq struct {
	DeviceName     string `json:"deviceName"`     // 设备名
	Model          string `json:"model"`          // 设备型号
	DevicePosition string `json:"devicePosition"` // 设备位置
	DeviceKind     string `json:"deviceKind"`     // 种类（1-摄像头(ipc) 2-nvr 3-门禁卡 4-收银机）
	BridgeDevice   int64  `json:"bridgeDevice"`   // 桥接设备，ipc设置nvr的设备id
	ChannelNo      int64  `json:"channelNo"`      // ipc 桥接nvr的通道号
	Brand          string `json:"brand"`          // 品牌方
	DeviceSerial   string `json:"deviceSerial"`   // 设备序列号
	ValidateCode   string `json:"validateCode"`   // 设备验证码
}
type EditDeviceReq struct {
	DeviceId int64 `json:"deviceId"` // 设备id
	AddDeviceReq
}
type DeviceAddOptReq struct {
	DeviceId int64 `json:"deviceId"` // 设备id
	AddOpt   int64 `json:"addOpt"`   // 上云操作（0-否 1-是）
}

type DeviceRelShopReq struct {
	ShopId   int64 `json:"shopId"`   // 门店id
	DeviceId int64 `json:"deviceId"` // 设备id
	RelFlag  int64 `json:"relFlag"`  // 关联操作标志（0-取消关联，1-关联)
}

type ShopDeviceTreeNode struct {
	Title    string               `json:"title"`
	Key      string               `json:"key"`
	IsLeaf   bool                 `json:"isLeaf"`
	Children []ShopDeviceTreeNode `json:"children"`
}

type RoleIdReq struct {
	RoleId string `json:"roleId"` // 角色ID
}
type GetRoleListRep struct {
	DataLen int64                `json:"dataLen"`
	Data    []RoleListOutputData `json:"data"`
	CommonHttpRep
}
type RoleListOutputData struct {
	RoleId     string `json:"roleId"`     // 角色ID
	RoleName   string `json:"roleName"`   // 角色名
	MenuItems  string `json:"menuItems"`  // 菜单项
	CreateTime string `json:"createTime"` // 创建时间（10位数字时间戳）
	UpdateTime string `json:"updateTime"` // 更新时间（10位数字时间戳）
	IsValid    string `json:"isValid"`    // 启用标志(0-启用 1-禁用)
}
type AddRoleReq struct {
	RoleName  string `json:"roleName"`  // 角色名
	MenuItems string `json:"menuItems"` // 菜单项
}
type EditRoleReq struct {
	RoleId string `json:"roleId"` // 角色ID
	AddRoleReq
}

type GetTaskListReq struct {
	ShopName     string `json:"shopName"`     // 门店名
	ClientName   string `json:"clientName"`   // 客户名
	DutyAccNo    string `json:"dutyAccNo"`    // 值守员工工号
	ApproveAccNo string `json:"approveAccNo"` // 审核员工工号
	TaskSerial   string `json:"taskSerial"`   // 工单流水号
	TaskType     string `json:"taskType"`     // 工单类型（0-事中正常 1-事中异常 2-事后）
	Status       string `json:"status"`       // 运行状态
	BeginTime    string `json:"beginTime"`    // 开始时间
	EndTime      string `json:"endTime"`      // 结束时间
}

type GetTaskListRep struct {
	DataLen int64            `json:"dataLen"`
	Data    []TaskOutPutData `json:"data"`
	CommonHttpRep
}

type TaskOutPutData struct {
	TaskId        int64  `json:"taskId"`        // 工单id
	TaskSerial    string `json:"taskSerial"`    // 工单流水号
	RelSerialNo   string `json:"relSerialNo"`   // 关联工单流水号（事后工单关联事中）
	ClientId      int64  `json:"clientId"`      // 申请客户id
	ClientName    string `json:"clientName"`    // 客户名
	ShopId        int64  `json:"shopId"`        // 门店id
	ShopName      string `json:"shopName"`      // 门店名
	TaskType      int64  `json:"taskType"`      // 工单类型（0-事中正常 1-事中异常 2-事后）
	ApproveAccNo  string `json:"approveAccNo"`  // 审批员工工号
	DutyAccNo     string `json:"dutyAccNo"`     // 值守员工工号
	Status        string `json:"status"`        // 值守状态（0-待审核 1-值守中 2-下线待确认 3-值守完成，8-异常处理中 9-完成）
	CreateTime    string `json:"createTime"`    // 创建时间（10位数字时间戳）
	ApproveTime   string `json:"approveTime"`   // 审核时间（10位数字时间戳）
	CancelTime    string `json:"cancelTime"`    // 取消云值守时间（10位数字时间戳）
	DutyEndTime   string `json:"dutyEndTime"`   // 云值守结束时间（10位数字时间戳）
	FinishTime    string `json:"finishTime"`    // 完成时间（10位数字时间戳）
	ApproveRemark string `json:"approveRemark"` // 审核批复
	CancelRemark  string `json:"cancelRemark"`  // 取消云值守批复
	DutyEndRemark string `json:"dutyEndRemark"` // 云值守结束批复
	FinishRemark  string `json:"finishRemark"`  // 完成工单结束批复
	TransNum      int64  `json:"transNum"`      // 交易数
	EventNum      int64  `json:"eventNum"`      // 事件数
}

type DutyApproveReq struct {
	TaskSerial    string `json:"taskSerial"`    // 工单流水号
	AutoFlag      string `json:"autoFlag"`      // 是否自动分配（0-否，1-是）	必填
	DutyAccNo     string `json:"dutyAccNo"`     // 值守员工工号
	ApproveRemark string `json:"approveRemark"` // 	审核批复
}

type OffDutyConfirmReq struct {
	TaskSerial string `json:"taskSerial"` // 工单流水号
	DutyRemark string `json:"dutyRemark"` // 云值守结束批复
}

type AccNoReq struct {
	AccNo string `json:"accNo"` // 工号、登录账号
}
type GetUserListReq struct {
	AccNo   string `json:"accNo"`   // 工号、登录账号
	AccName string `json:"accName"` // 工号、登录账号
	RoleId  int64  `json:"roleId"`  // 角色ID
	OnWork  string `json:"onWork"`  // 在岗状态（0-离岗 1-在岗）
}
type GetUserListRep struct {
	DataLen int64              `json:"dataLen"`
	Data    []UserListItemData `json:"data"`
	CommonHttpRep
}
type UserListItemData struct {
	AccNo      string `json:"accNo"`      // 工号、登录账号
	AccName    string `json:"accName"`    // 工号、登录账号
	Phone      string `json:"phone"`      // 手机号
	Email      string `json:"email"`      // 邮箱
	RoleId     int64  `json:"roleId"`     // 角色ID
	OnWork     string `json:"onWork"`     // 在岗状态（0-离岗 1-在岗）
	CreateTime string `json:"createTime"` // 创建时间（10位数字时间戳）
	UpdateTime string `json:"updateTime"` // 更新时间（10位数字时间戳）
	LogonTime  string `json:"logonTime"`  // 登录时间（10位数字时间戳）
	IsValid    int64  `json:"isValid"`    // 启用标志(0-启用 1-禁用)
}

type AddUserReq struct {
	AccName string `json:"AccName"` // 工号、登录账号
	RoleId  int64  `json:"roleId"`  // 角色ID
	Phone   string `json:"phone"`   // 手机号
	Email   string `json:"email"`   // 邮箱
}
type EditUserReq struct {
	AccNo string `json:"accNo"` // 工号、登录账号
	AddUserReq
}
type ResetUserPasswordReq struct {
	AccNo string `json:"accNo"` // 工号、登录账号
}

type DutyListRep struct {
	Token   string         `json:"token"` // 视频播放token
	DataLen int64          `json:"dataLen"`
	Data    []DutyListItem `json:"data"`
	CommonHttpRep
}
type DutyListItem struct {
	TaskId     string `json:"taskId"`     // 工单id
	TaskSerial string `json:"taskSerial"` // 工单流水号
	ClientId   string `json:"clientId"`   // 申请客户id
	ClientName string `json:"clientName"` // 客户名
	ShopId     string `json:"shopId"`     // 门店id
	ShopName   string `json:"shopName"`   // 门店名
	InTrans    string `json:"inTrans"`    // 是否有交易（0-否 1-是）
	VideoUrl   string `json:"videoUrl"`   // 主视频链接
}
type StationVideoListReq struct {
	TaskSerial string `json:"taskSerial"` // 工单流水号
}

type StationVideoListRep struct {
	DataLen      int64                  `json:"dataLen"`
	Data         []StationVideoListItem `json:"data"`
	ContactName  string                 `json:"contactName"`  // 店铺联系人姓名
	ContactPhone string                 `json:"contactPhone"` // 店铺联系人手机
	Token        string                 `json:"token"`        // 视频播放token
	CommonHttpRep
}

type StationVideoListItem struct {
	ShopId         string `json:"shopId"`         // 门店id
	ShopName       string `json:"shopName"`       // 门店名
	DeviceName     string `json:"deviceName"`     // 设备名
	DevicePosition string `json:"devicePosition"` // 位置
	DeviceKind     string `json:"deviceKind"`     // 设备种类
	VideoUrl       string `json:"videoUrl"`       // 视频链接
}

type StationGetBillReq struct {
	TaskSerial string `json:"taskSerial"` // 工单流水号
}

type StationGetBillRep struct {
	TaskSerial     string          `json:"taskSerial"`     // 工单流水号
	TransSerial    string          `json:"transSerial"`    // 交易流水号
	TransBeginTime string          `json:"transBeginTime"` // 交易开始时间
	CustomerPhone  string          `json:"customerPhone"`  // 顾客手机
	BillList       []*BillListItem `json:"billList"`
	CommonHttpRep
}
type BillListItem struct {
	BillSerial     string         `json:"billSerial"`     //	订单流水号
	BillStatus     string         `json:"billStatus"`     //	订单状态（0-进行中，1-支付成功，2-支付失败）
	BillCreateTime string         `json:"billCreateTime"` //	订单创建时间
	TotalNum       string         `json:"totalNum"`       //	商品总件数
	TotalAmount    string         `json:"totalAmount"`    //	商品总计金额 单位分
	PayAmount      string         `json:"payAmount"`      //	商品总计金额 单位分
	CheckStatus    string         `json:"checkStatus"`    //	订单审核状态（0-待审核，1-审核通过，2-拒绝）
	GoodList       []GoodListItem `json:"goodList"`
}

type SortedBills []*BillListItem

func (s SortedBills) Len() int {
	return len(s)
}
func (s SortedBills) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s SortedBills) Less(i, j int) bool {
	return s[i].BillCreateTime > s[j].BillCreateTime
}

type GoodListItem struct {
	GoodId    string `json:"goodId"`    //	商品id
	GoodName  string `json:"goodName"`  //	商品名
	GoodNum   string `json:"goodNum"`   //	商品数量
	GoodPrice string `json:"goodPrice"` //	商品单价
	GoodSum   string `json:"goodSum"`   //	商品小计
}

type StationGetStatusReq struct {
	TaskSerial string `json:"taskSerial"` // 工单流水号
}

type StationGetStatusRep struct {
	Status string `json:"status"` //工单状态（0-待审核 1-值守中 2-下线待确认 3-值守完成，8-异常处理中 9-完成）
	CommonHttpRep
}

type StationDutyAlarmReq struct {
	TaskSerial   string `json:"taskSerial"`   // 工单流水号
	TransSerial  string `json:"transSerial"`  // 交易流水号
	EventContent string `json:"eventContent"` // 事件说明
}

type StationOpenDoorReq struct {
	TaskSerial  string `json:"taskSerial"`  // 工单流水号
	TransSerial string `json:"transSerial"` // 交易流水号
}

type MessageCreationReq struct {
	MessageChannel string `json:"messageChannel"` // 消息发送渠道(0-app 1-miniapp 2-windows）
	UserId         string `json:"userId"`         // 消息发送用户
	MessageType    string `json:"messageType"`    // 消息类型：自定义字段，区分不同业务场景
	Topic          string `json:"topic"`          // 主题
	Content        string `json:"content"`        // 内容
	RelSerial      string `json:"relSerial"`      // 相关业务的主键或流水号
}

type GetMessageListReq struct {
	CurrentIndex int64  `json:"currentIndex"` // 当前最大索引
	BeginTime    string `json:"beginTime"`    // 开始时间
	EndTime      string `json:"endTime"`      // 结束时间
}
type GetMessageListRep struct {
	DataLen int64         `json:"dataLen"`
	Data    []MessageData `json:"data"`
	CommonHttpRep
}

type GetLatestMessageIndexRep struct {
	MaxIndex int64 `json:"maxIndex"` //当前用户最大消息索引
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

type GetEventListReq struct {
	TaskSerial  string `json:"taskSerial"`  // 工单流水号
	TransSerial string `json:"transSerial"` // 交易流水号
	BillSerial  string `json:"billSerial"`  // 订单流水号
	EventSerial string `json:"eventSerial"` // 事件流水号
	ClientName  string `json:"clientName"`  // 客户名
	ShopName    string `json:"shopName"`    // 门店名
	AccNo       string `json:"accNo"`       // 处理员工工号
	EventType   string `json:"eventType"`   // 事件类型 （1工单异常，2交易异常，3订单异常 ）
	EventLevel  string `json:"eventLevel"`  // 事件级别（1-普通，2-重要，3-紧急）
	Status      string `json:"status"`      // 处理状态（0待处理，1处理中，2完成）
	BeginTime   string `json:"beginTime"`   // 开始时间
	EndTime     string `json:"endTime"`     // 结束时间
}
type GetEventListRep struct {
	DataLen int64       `json:"dataLen"`
	Data    []EventData `json:"data"`
	CommonHttpRep
}
type EventData struct {
	EventId      int64  `json:"eventId"`      // 事件id
	EventSerial  string `json:"eventSerial"`  // 事件流水号
	TaskSerial   string `json:"taskSerial"`   // 工单流水号
	TransSerial  string `json:"transSerial"`  // 交易流水号
	BillSerial   string `json:"billSerial"`   // 订单流水号
	Status       string `json:"status"`       // 处理状态（0待处理，1处理中，2完成）
	EventType    string `json:"eventType"`    // 事件类型  （1工单异常，2交易异常，3订单异常 ）
	EventLevel   string `json:"eventLevel"`   // 事件级别（1-普通，2-重要，3-紧急）
	Content      string `json:"content"`      // 事件描述
	ReplyContent string `json:"replyContent"` // 处理批复
	AccNo        string `json:"accNo"`        // 处理员工工号
	CreateTime   string `json:"createTime"`   // 创建时间（10位数字时间戳）
	UpdateTime   string `json:"updateTime"`   // 更新时间（10位数字时间戳）
	ClientId     int64  `json:"clientId"`     // 申请客户id
	ClientName   string `json:"clientName"`   // 客户名
	ClientPhone  string `json:"clientPhone"`  // 客户手机
	ShopId       int64  `json:"shopId"`       // 门店id
	ShopName     string `json:"shopName"`     // 门店名
	ContactName  string `json:"contactName"`  // 门店联系人姓名
	ContactPhone string `json:"contactPhone"` // 门店联系方式
	Address      string `json:"address"`      // 门店详细地址
}

type EventConfirmReq struct {
	EventId      int64  `json:"eventId"`      // 工单id
	ReplyContent string `json:"replyContent"` // 处理批复
	FinishFlag   string `json:"finishFlag"`   // 是否完成 0-否 1-是
}

type GetTransListReq struct {
	TransSerial  string `json:"transSerial"`   // 交易流水号
	TaskSerial   string `json:"taskSerial"`    // 工单流水号
	ShopName     string `json:"shopName"`      // 门店名
	Phone        string `json:"customerPhone"` // 顾客手机号
	DutyAccNo    string `json:"dutyAccNo"`     // 值守员工工号
	ApproveAccNo string `json:"approveAccNo"`  // 审核员工工号
	Status       string `json:"status"`        // 交易状态
	BeginTime    string `json:"beginTime"`     // 开始时间
	EndTime      string `json:"endTime"`       // 结束时间
}
type GetTransListRep struct {
	DataLen int64       `json:"dataLen"`
	Data    []TransData `json:"data"`
	CommonHttpRep
}
type TransData struct {
	TransId       int64  `json:"transId"`       // 交易id
	TransSerial   string `json:"transSerial"`   // 流水号
	TaskSerial    string `json:"taskSerial"`    // 工单流水号
	DutyAccNo     string `json:"dutyAccNo"`     // 值守员工工号
	ApproveAccNo  string `json:"approveAccNo"`  // 审核员工工号
	BeginTime     string `json:"beginTime"`     // 交易开始时间（10位数字时间戳）
	EndTime       string `json:"endTime"`       // 交易结束时间（10位数字时间戳）
	Status        string `json:"status"`        // 交易状态（0-购物中 1-完成（无购物离店）2-完成（购物离店） -1-异常）
	ShopId        int64  `json:"shopId"`        // 门店id
	ShopName      string `json:"shopName"`      // 门店名
	CustomerId    int64  `json:"customerId"`    // 顾客id
	CustomerPhone string `json:"customerPhone"` // 顾客手机
}

type GetBillListReq struct {
	BillSerial    string `json:"billSerial"`    // 订单流水号
	TransSerial   string `json:"transSerial"`   // 交易流水号
	PayType       string `json:"payType"`       // 支付方式 1-微信 2-支付宝 3-云闪付
	PayStatus     string `json:"payStatus"`     // 支付状态（0-待支付，1-成功，2-失败）
	ShopName      string `json:"shopName"`      // 门店名
	CustomerPhone string `json:"customerPhone"` // 顾客名
	CheckStatus   string `json:"checkStatus"`   // 审核状态（0-待审核，1-审核通过，2-拒绝）
	BeginTime     string `json:"beginTime"`     // 开始时间
	EndTime       string `json:"endTime"`       // 结束时间
}
type GetBillListRep struct {
	DataLen int64      `json:"dataLen"`
	Data    []BillData `json:"data"`
	CommonHttpRep
}
type BillData struct {
	BillId        int64  `json:"billId"`        // 订单id
	BillSerial    string `json:"billSerial"`    // 流水号
	TransSerial   string `json:"transSerial"`   // 交易流水号
	TotalNum      string `json:"totalNum"`      // 总件数
	TotalAmount   string `json:"totalAmount"`   // 总金额
	PayAmount     string `json:"payAmount"`     // 支付金额 单位分
	PayType       string `json:"payType"`       // 支付方式 1-微信 2-支付宝 3-云闪付
	PayTime       string `json:"payTime"`       // 支付时间（10位数字时间戳）
	PaySerial     string `json:"paySerial"`     // 支付序列号
	PayStatus     string `json:"payStatus"`     // 支付状态（0-待支付，1-成功，2-失败）
	Detail        string `json:"detail"`        // 物品详细清单（由收银机推送具体数据待定）
	CreateTime    string `json:"createTime"`    // 创建时间（10位数字时间戳）
	ShopId        int64  `json:"shopId"`        // 门店id
	ShopName      string `json:"shopName"`      // 门店名
	CustomerId    int64  `json:"customerId"`    // 顾客id
	CustomerPhone string `json:"customerPhone"` // 顾客手机
	CheckStatus   string `json:"checkStatus"`   // 审核状态（0-待审核，1-审核通过，2-拒绝）
	CheckTime     string `json:"checkTime"`     // 审核时间（10位数字时间戳）
	EventSerial   string `json:"eventSerial"`   // 审核拒绝后追溯事件流水号
}

type BillApproveReq struct {
	BillSerial string `json:"billSerial"` // 订单流水号
	Flag       string `json:"flag"`       // 审核标志（1-通过 2-拒绝）
	Reason     string `json:"reason"`     // 拒绝原因
}

type GetLocalVideoReq struct {
	DeviceId  string `json:"deviceId"`  // 监控设备序列号
	BeginTime string `json:"beginTime"` // 开始时间
	EndTime   string `json:"endTime"`   // 结束时间
}

type GetLocalVideoRep struct {
	Data LocalVideoData `json:"data"`
	CommonHttpRep
}

type LocalVideoData struct {
	VideoUrl string `json:"videoUrl"` // 视频播放url
	Token    string `json:"token"`    // 视频播放token
}

type GetSuggestionListReq struct {
	Channel   string `json:"channel"`   // 投诉渠道
	UserName  string `json:"userName"`  // 投诉人名
	Status    string `json:"status"`    // 处理状态 （0-待处理 1-完成）
	BeginTime string `json:"beginTime"` // 开始时间
	EndTime   string `json:"endTime"`   // 结束时间
}

type GetSuggestionListRep struct {
	DataLen int64                   `json:"dataLen"`
	Data    []GetSuggestionListItem `json:"data"`
	CommonHttpRep
}
type GetSuggestionListItem struct {
	SuggestionId int64  `json:"suggestionId"` // id
	UserId       string `json:"userId"`       // 投诉人
	UserName     string `json:"userName"`     // 投诉人名
	Channel      string `json:"channel"`      // 投诉渠道
	Contact      string `json:"contact"`      // 联系方式（页面采集）
	Content      string `json:"content"`      // 投诉内容
	CreateTime   string `json:"createTime"`   // 投诉时间（10位数字时间戳）
	Status       string `json:"status"`       // 处理状态 （0-待处理 1-完成）
	ReplyAccNo   string `json:"replyAccNo"`   // 处理员工号
	ReplyAccName string `json:"replyAccName"` // 处理员工名
	ReplyTime    string `json:"replyTime"`    // 处理时间（10位数字时间戳）
	ReplyRemark  string `json:"replyRemark"`  // 处理批复
}
type SuggestionApproveReq struct {
	SuggestionId int64  `json:"suggestionId"` // id
	ReplyRemark  string `json:"replyRemark"`  // 处理批复
}
type GetShopRankListRep struct {
	Data []ShopRankLisItem `json:"data"`
	CommonHttpRep
}
type ShopRankLisItem struct {
	ShopName string `json:"shopName"` // 门店名
	VisitNum string `json:"visitNum"` // 在店人数
}

type GetShopCountRep struct {
	Data ShopCountData `json:"data"`
	CommonHttpRep
}

type ShopCountData struct {
	Total    int64 `json:"total"`    // 总设备数
	Online   int64 `json:"online"`   // 在线设备数
	Offline  int64 `json:"offline"`  // 离线设备数
	Abnormal int64 `json:"abnormal"` // 异常设备数
}

type StatTransSumRep struct {
	Data StatTransSumData `json:"data"`
	CommonHttpRep
}

type StatTransSumData struct {
	TotalVisit     string `json:"totalVisit"`     // 总客流
	TotalIncome    string `json:"totalIncome"`    // 总营收
	PayCount       string `json:"payCount"`       // 付款单数
	PayPerCustomer string `json:"payPerCustomer"` // 客单价
	DiffRate       string `json:"diffRate"`       // 较上期
}

type StatChartVisitReq struct {
	StepType int64 `json:"stepType"` //步幅类型  （0-时，1-日，2-周，3-月）
}

type StatChartVisitRep struct {
	Data []ChartData `json:"data"`
	CommonHttpRep
}

type StatChartIncomeReq struct {
	StepType int64 `json:"stepType"` //步幅类型  （0-时，1-日，2-周，3-月）
}

type StatChartIncomeRep struct {
	Data []ChartData `json:"data"`
	CommonHttpRep
}

type StatChartPayNumReq struct {
	StepType int64 `json:"stepType"` //步幅类型  （0-时，1-日，2-周，3-月）
}

type StatChartPayNumRep struct {
	Data []ChartData `json:"data"`
	CommonHttpRep
}

type ChartData struct {
	Label string `json:"label"` //	统计标签，按日统计，为时分，按7d/30d统计为月日, 按年统计为月份
	Value string `json:"value"` //	数量
}

type GetUserDutyListReq struct {
	AccNo   string `json:"accNo"`   // 工号、登录账号
	AccName string `json:"accName"` // 员工名
	OnWork  string `json:"onWork"`  // 在岗状态（0-离岗 1-在岗）
}

type GetUserDutyListRep struct {
	DataLen int64          `json:"dataLen"`
	Data    []UserDutyData `json:"data"`
	CommonHttpRep
}

type UserDutyData struct {
	AccNo          string `json:"accNo"`          // 工号、登录账号
	AccName        string `json:"accName"`        // 员工名
	Phone          string `json:"phone"`          // 手机号
	OnWork         string `json:"onWork"`         // 在岗状态（0-离岗 1-在岗）
	TotalDutyNum   int64  `json:"totalDutyNum"`   // 总任务数
	EndDutyNum     int64  `json:"endDutyNum"`     // 完成任务数
	CurrentDutyNum int64  `json:"currentDutyNum"` // 当前任务数
	TodayEndNum    int64  `json:"todayEndNum"`    // 今日完成任务数
}
