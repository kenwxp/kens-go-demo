package types

type Client struct {
	ClientId    int64  // 客户id
	ClientName  string // 客户名
	ClientPhone string // 手机
	Password    string // 登录密码
	Salt        string // 盐
	Token       string // token
	CreateTime  string // 创建时间（10位数字时间戳）
	UpdateTime  string // 创建时间（10位数字时间戳）
	LogonTime   string // 登录时间（10位数字时间戳）
	IsValid     int64  // 启用标志(0-启用 1-禁用
}

type Customer struct {
	CustomerId  int64  // 自增	顾客id	pk
	Nickname    string // 用户昵称
	Avatar      string // 用户头像
	Phone       string // 手机号
	Gender      string // 性别 0-男 1-女
	Birth       string // 生日 yyyy-mm-dd
	Token       string // token
	OpenId      string // 微信openId
	SessionKey  string // 微信sessionKey
	IsMember    int64  // 是否会员 0-否 1-是
	JoinTime    string // 加入会员时间   （10位数字时间戳）
	JoinEndTime string // 退出会员时间   （10位数字时间戳）
	CreateTime  string // 创建时间   （10位数字时间戳）
	UpdateTime  string // 更新时间   （10位数字时间戳）
	IsBlack     int64  // 是否黑名单 0-否 1-是
}
type User struct {
	AccNo      string // 工号、登录账号	pk
	AccName    string // 工号、登录账号	pk
	Password   string // 密码
	Salt       string // 盐
	Token      string // token
	Phone      string // 手机号
	Email      string // 邮箱
	RoleId     int64  // 角色权限（1-管理员 2-云值守员 3-工客服人员）
	OnWork     string // 在岗状态（0-离岗 1-在岗）
	CreateTime string // 创建时间（10位数字时间戳）
	UpdateTime string // 更新时间（10位数字时间戳）
	LogonTime  string // 更新时间（10位数字时间戳）
	IsValid    int64  // 启用标志(0-启用 1-禁用)
	Role
}
type Role struct {
	RoleId     int64  // 角色ID
	RoleName   string // 角色名
	MenuItems  string // 菜单项
	CreateTime string // 创建时间（10位数字时间戳）
	UpdateTime string // 更新时间（10位数字时间戳）
	IsValid    string // 启用标志(0-启用 1-禁用)
}
type Region struct {
	RegionId        string // 地区主键编号	pk
	RegionName      string // 地区名称
	RegionShortName string // 地区缩写
	RegionCode      string // 行政地区编号
	RegionParentId  string // 地区父id
	RegionLevel     string // 地区级别 1-省、自治区、直辖市 2-地级市、地区、自治州、盟 3-市辖区、县级市、县
}

type ShippingAddress struct {
	AddressId  string // 自增 id pk
	CustomerId int64  // 顾客id
	AcceptName string // 收货人
	Mobile     string // 手机
	AreaCode   string // 地区号
	Street     string // 详细街道地址
	DoorNumber string // 门牌号
	IsDefault  int64  // 是否默认 0-否 1-是
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	IsValid    int64  // 是否启用 0-启用，1-禁用
	Customer
}

type Goods struct {
	GoodsId         int64  //  商品id
	GoodsName       string //  商品名
	GoodsType       string //  商品类别 1-普通 2-积分
	Category        string //  商品类别
	Content         string //  商品描述
	Images          string //  商品图片 数组
	Price           string //  售价 单位分
	MembershipPrice string //  会员价 单位分
	Unit            string //  单位
	Stock           int64  //  库存
	MinBuyNum       int64  //  最小起售数
	UseProperty     int64  //  是否使用规格
	PropertyGroups  string //  规格数组
	SpecGroups      string //  参数数组
	IsSell          string //  上架标志 0-下架 1-上架
	Sort            int64  //  排序
	CreateTime      string //  创建时间
	UpdateTime      string //  更新时间
	IsValid         int64  //  启用标志 0-启用 1-禁用
	GoodsCategory
}

type GoodsCategory struct {
	CategoryId   int64  //	类别id
	CategoryName string //	类别名
	Icon         string //	图表
	Sort         int64  //	排序
	CreateTime   string //	创建时间
	UpdateTime   string //	更新时间
	IsValid      int64  //	是否启用 0-启用 1-禁用
}

type GoodsProperty struct {
	Id        int64  // 自增id 主键
	GoodsId   string // 商品id
	SubId     int64  // 规格id
	SubParams string // 规格参数
	Stock     int64  // 库存
	Price     string // 价格
	PicUrl    string // 图片
	IsValid   int64  // 启用标志 0-启用 1-禁用
	Goods
}

type Order struct {
	OrderId           int64  // 订单id 主键
	OrderSerial       string // 订单序列号
	PaymentId         int64  // 支付id
	CommentId         int64  // 评论id
	CustomerId        int64  // 顾客id
	ShopId            int64  // 店铺id
	GoodsNum          string // 商品总件数
	GoodsAmount       string // 商品金额
	SendFee           string // 配送金额
	Discount          string // 折扣金额
	TotalAmount       string // 总金额
	GoodsDetail       int64  // 商品详情
	OrderStatus       string // 订单状态
	SendType          string // 寄送类型 1-自取 2-外卖
	ShippingAddressId int64  // 寄送地址id
	CreateTime        string // 创建时间
	ProduceTime       string // 制作时间
	SendTime          string // 寄送时间
	ReceiveTime       string // 送达时间
	FinishTime        string // 完成时间
	IsValid           int64  // 启用标志 0-启用 1-禁用
	Customer
	OrderPayment
}

type OrderPayment struct {
	PaymentId   int64  // 支付id
	PayOrderNum string // 微信支付的订单号
	PrepayId    string // 微信支付预付款单号
	MchId       string // 微信支付商户号
	OpenId      string // 支付的用户的微信openid
	CustomerId  int64  // 支付的用户id
	ShopId      int64  // 门店id
	PayAmount   string // 支付金额 单位分
	PayFee      string // 支付手续费 单位分
	PayStatus   string // 支付状态 0-待支付 1-已支付 2-支付失败
	CreateTime  string // 创建时间
	UpdateTime  string // 更新时间
	IsValid     int64  // 启用标志 0-启用 1-禁用
}
