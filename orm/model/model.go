package model

import "time"

//table 基础信息
const (
	SynastryHarvestDataTable           = "synastry_harvest_data"            // 收获文案
	ReportUserTable                    = "report_user"                      // 用户举报
	SynastryDataTable                  = "synastry_data"                    // 合盘星象数据，含相位和宫位
	TradeInfoTable                     = "trade_info"                       // 揪揪订单信息表
	V2SynastryRelationshipMatchTable   = "v2_synastry_relationship_match"   // v2 最适合的关系匹配表，根据最高和最低分的标签id来查找
	FeedCommentCounterTable            = "feed_comment_counter"             // 评论点赞计数
	FootholdConflictTable              = "foothold_conflict"                // 和谐与冲突
	FortuneContentTable                = "fortune_content"                  //
	FortuneDayLabelsTable              = "fortune_day_labels"               // 日运标签详情表
	LuckySymbolTable                   = "lucky_symbol"                     //
	MessageMeTable                     = "message_me"                       // 我点赞，评论，收藏信息表
	PushActivityTable                  = "push_activity"                    // 大春子活动页信息表
	TopicInfoTable                     = "topic_info"                       // 话题
	DirtyWordTable                     = "dirty_word"                       // 脏字系统
	FeedTable                          = "feed"                             // 用户个人详情页数据
	UserTable                          = "user"                             // 用户信息表
	V2AstrolabeMoonTable               = "v2_astrolabe_moon"                // 个人星盘月亮星座数据表
	CliConfTable                       = "cli_conf"                         // 客户端基础配置信息表
	CompensateInfoTable                = "compensate_info"                  // 补偿信息表
	FeedCommentLikeTable               = "feed_comment_like"                // 针对评论是点赞详情记录
	MessageTable                       = "message"                          // 互动消息表
	PhaseNotFoundPlaceholderTable      = "phase_not_found_placeholder"      // 指定相位找不到时的填充文案
	PlanetConstellationTable           = "planet_constellation"             // 行星，星座描述
	ArchiveStarTable                   = "archive_star"                     // 档案信息表(丢弃)
	BannerInfoTable                    = "banner_info"                      // banner基础信息
	WeekFortuneIndicesTable            = "week_fortune_indices"             // 周运基础数据索引
	PlanetHouseTable                   = "planet_house"                     // 行星宫位描述信息表
	V2AstrolabeSunTable                = "v2_astrolabe_sun"                 // 个人星盘太阳星座数据表
	FortuneLuckyBasicsTable            = "fortune_lucky_basics"             // 基础运势项目表
	FortuneMonthYearLabelsTable        = "fortune_month_year_labels"        // 月运年运星象标签表
	GoodsInfoTable                     = "goods_info"                       // 揪揪商品信息表
	SmsPhoneBlockTable                 = "sms_phone_block"                  // 短信号码黑名单
	ConstellationElementTable          = "constellation_element"            // 星座元素
	FortuneImportantMatchElementsTable = "fortune_important_match_elements" // 最大运势-幸运元素匹配表
	FortuneYearMatchContentTable       = "fortune_year_match_content"       // 年运匹配文案表
	LcTencentRegionsTable              = "lc_tencent_regions"               // 地理位置经纬度信息表
	V2AstrolabePlanetManTable          = "v2_astrolabe_planet_man"          // 个人星盘行星人数据表
	FortuneMajorDetailsTable           = "fortune_major_details"            //
	FortuneMonthYearPhasesTable        = "fortune_month_year_phases"        // 月运年运星象文案表
	FortuneHouseTable                  = "fortune_house"                    //
	MessageLikeTable                   = "message_like"                     // 互动点赞消息表
	PlanetDescTable                    = "planet_desc"                      // 星体描述
	PushTaskTable                      = "push_task"                        // 推送任务信息表
	ArticleUserTable                   = "article_user"                     // 星文作者
	FortuneDayPhasesTable              = "fortune_day_phases"               // 新版日运星象表
	SynastryRelationshipExtendTable    = "synastry_relationship_extend"     // 合盘最适合的关系，特殊关系扩充表
	V2AstrolabeAscTable                = "v2_astrolabe_asc"                 // 个人星盘上升星座数据表
	V2AstrolabeElementsTable           = "v2_astrolabe_elements"            // 个人星盘元素属性数据表
	V2SynastryDataTable                = "v2_synastry_data"                 // v2 优化版合盘星象数据，含相位和宫位
	ArticleTable                       = "article"                          // 发布的星文信息表
	FortunePhaseTable                  = "fortune_phase"                    //
	RouterInfoTable                    = "router_info"                      // 路由（接口）信息表
	SynastryRelationshipTable          = "synastry_relationship"            // 最适合的关系，通用文案
	V2AstrolabePlanetScoreTable        = "v2_astrolabe_planet_score"        // 个人星盘行星人得分数据表
	FortuneLuckyElementsTable          = "fortune_lucky_elements"           // 改版日运新增幸运元素表
	PlanetPhaseCombinedTable           = "planet_phase_combined"            // 相位合盘描述信息表
	FotruneShouldAvoidTable            = "fotrune_should_avoid"             //
	SplashInfoTable                    = "splash_info"                      // 启屏页基础信息
	ArchiveH5ShareTable                = "archive_h5_share"                 // 分享用户档案（合盘模块）
	FeedCollectionDetailTable          = "feed_collection_detail"           // 用户收藏详情表
	FateEnergyTable                    = "fate_energy"                      // 缘分能量区间描述
	FeedCounterTable                   = "feed_counter"                     // feed相关计数统计，用于feed后期推荐计算
	FortuneMajorTable                  = "fortune_major"                    //
	SynastryDataExtendTable            = "synastry_data_extend"             // 合盘特殊关系扩展文案表
	V25EncounterPlanetDataTable        = "v25_encounter_planet_data"        // 行星邂逅星体数据表
	V2SynastryHarvestDataTable         = "v2_synastry_harvest_data"         // v2 收获文案
	ApkUpdateInfoTable                 = "apk_update_info"                  // apk更新信息表
	ArticleJobTable                    = "article_job"                      // 星文抓取作业
	ReviveDaysTable                    = "revive_days"                      //
	UserPermissionTable                = "user_permission"                  // 用户权限树
	V25EncounterLevelDataTable         = "v25_encounter_level_data"         // 邂逅等级数据表
	CombinedHistoryTable               = "combined_history"                 // 合盘记录信息表
	LoginLogTable                      = "login_log"                        // 用户上次登录行为日志
	ArticleCounterTable                = "article_counter"                  // 星文点赞，阅读计数
	ArticleLikeDetailTable             = "article_like_detail"              // 星文点赞详情
	FeedCommentTable                   = "feed_comment"                     // 评论发布的星文|瞬间记录表
	FeedLikeDetailTable                = "feed_like_detail"                 // 用户点赞详情表
	FortuneEventHistoryTable           = "fortune_event_history"            // 揪揪事件记录
	FortuneJiujiuEventsTable           = "fortune_jiujiu_events"            // 揪揪事件和事件类型，按事件类型对应的不同星象区分每一条数据
	ArchiveTable                       = "archive"                          // 档案信息表
	ArchiveH5Table                     = "archive_h5"                       // h5合盘档案信息表（合盘模块）
	V2SynastryRelationshipTable        = "v2_synastry_relationship"         // v2 最适合的关系，通用文案
	WeekFortuneDataTable               = "week_fortune_data"                // 周运基础数据
	PurchaseGoodsInfoTable             = "purchase_goods_info"              // 用户购买商品信息表
	V2AstrolabeModeTable               = "v2_astrolabe_mode"                // 个人星盘星座属性数据表
	EventInfoTable                     = "event_info"                       // 客户端上报事件管理
	FeedTopicCountTable                = "feed_topic_count"                 // feed话题统计
	FollowTopicTable                   = "follow_topic"                     // 关注话题信息记录表
	FotruneAstroTipsTable              = "fotrune_astro_tips"               //
	ReportFeedTable                    = "report_feed"                      // feed举报信息表
	V2SynastryLabelsDescTable          = "v2_synastry_labels_desc"          // v2 合盘6个标签不同分段的描述词
	ArchivesRecommendTable             = "archives_recommend"               // 推荐类型档案（2.0）
	AsoInfoTable                       = "aso_info"                         // aso信息表
	QuadrantAnalysisTable              = "quadrant_analysis"                // 象限对应的文案
	SynastryNotFoundPlaceholderTable   = "synastry_not_found_placeholder"   // 指定相位找不到时的填充文案
	UserBlockTable                     = "user_block"                       // 用户拉黑列表
	V2AstrolabeQuadrantTable           = "v2_astrolabe_quadrant"            // 个人星盘四象限数据表
	FortuneAdviceTable                 = "fortune_advice"                   //
	PlanetPhaseSelfTable               = "planet_phase_self"                // 相位本命描述信息表
)

type SynastryHarvestData struct {
	ContentKey int32  `db:"content_key" json:"content_key,omitempty"` // 通过给定标签的有无，得到的二进制数值对应的编号
	Good       int32  `db:"good" json:"good,omitempty"`               // 收获文案的吉凶值，0 无效默认值，1 吉，2 凶
	Content    string `db:"content" json:"content,omitempty"`         // 收获文案详情
}

type ReportUser struct {
	ReportId      int64      `db:"report_id" json:"report_id,omitempty"` // 举报id
	Uid           int64      `db:"uid" json:"uid,omitempty"`             // 举报人
	Ruid          int64      `db:"ruid" json:"ruid,omitempty"`           // 被举报人id
	ReportType    byte       `db:"report_type" json:"report_type"`       // 举报类型： 1. 昵称； 2. 签名； 3. 背景图； 4. 聊天；
	ReportContent byte       `db:"report_content" json:"report_content"` // 1. 色情低俗； 2. 广告骚扰； 3. 辱骂、人身攻击； 4. 违法内容；
	EventStatus   byte       `db:"event_status" json:"event_status"`     // 事件处理状态:0 未处理 1.已经处理
	ReportTime    *time.Time `db:"report_time" json:"report_time"`       // 举报时间
}

type SynastryData struct {
	PhaseKey      string `db:"phase_key" json:"phase_key,omitempty"`           // 相位或者宫位key
	Type          int32  `db:"type" json:"type,omitempty"`                     // 0 默认无效值，1 相位，2 宫位
	Label         int32  `db:"label" json:"label,omitempty"`                   // 星象对应标签，6种之一
	Good          int32  `db:"good" json:"good,omitempty"`                     // 星象对应吉凶，0 默认无效值，1 吉，2 凶
	Score         int32  `db:"score" json:"score,omitempty"`                   // 相位对应的评分 0-100
	Content       string `db:"content" json:"content,omitempty"`               // 星象对应的通用文案
	ExtendContent string `db:"extend_content" json:"extend_content,omitempty"` // 准了来源的相位描述
	Id1           int32  `db:"id1" json:"id1,omitempty"`
	Id2           int32  `db:"id2" json:"id2,omitempty"`
	Phase         int32  `db:"phase" json:"phase,omitempty"`
}

type TradeInfo struct {
	Id               int64      `db:"id" json:"id,omitempty"`                                 // 流水id，可做为订单的系统自有id
	Uid              int64      `db:"uid" json:"uid,omitempty"`                               // 用户id
	NickName         string     `db:"nick_name" json:"nick_name,omitempty"`                   // 购买商品用户昵称
	GoodsId          int64      `db:"goods_id" json:"goods_id,omitempty"`                     // 商品id
	GoodsName        string     `db:"goods_name" json:"goods_name,omitempty"`                 // 商品名称
	KeepMonth        int32      `db:"keepMonth" json:"keepMonth,omitempty"`                   // 持续时间
	Detail           string     `db:"detail" json:"detail,omitempty"`                         // 交易详情
	OutTradeNo       string     `db:"out_trade_no" json:"out_trade_no,omitempty"`             // 商户订单编号
	Money            string     `db:"money" json:"money,omitempty"`                           // 交易金额 单位 元
	RecpetData       string     `db:"recpet_data" json:"recpet_data,omitempty"`               // applePay 支付凭证
	PayType          byte       `db:"pay_type" json:"pay_type"`                               // 支付方式 1.支付宝 2.微信 3.applePay
	OrderStatus      byte       `db:"order_status" json:"order_status"`                       // 订单状态：1.待支付 2.支付成功 3.交易完成；4.订单关闭；5.已退款
	OrderStatusDes   string     `db:"order_status_des" json:"order_status_des,omitempty"`     // 订单状态描述
	TradeCreateTime  *time.Time `db:"trade_create_time" json:"trade_create_time"`             // 订单创建时间
	TradeSuccessTime *time.Time `db:"trade_success_time" json:"trade_success_time"`           // 订单支付成功
	TradeCloseTime   *time.Time `db:"trade_close_time" json:"trade_close_time"`               // 订单关闭时间
	TradeDoneTime    *time.Time `db:"trade_done_time" json:"trade_done_time"`                 // 订单完成时间
	TradeRefundTime  *time.Time `db:"trade_refund_time" json:"trade_refund_time"`             // 订单退款时间
	TradeErrCodeDes  string     `db:"trade_err_code_des" json:"trade_err_code_des,omitempty"` // 交易错误代码描述
	BillCreateIp     string     `db:"bill_create_ip" json:"bill_create_ip,omitempty"`         // 支付API的机器IP
	ChannelName      string     `db:"channel_name" json:"channel_name,omitempty"`             // 渠道
	PlatForm         byte       `db:"plat_form" json:"plat_form"`                             // 1. 安卓  2 ios
}

type V2SynastryRelationshipMatch struct {
	UpperLabel      int32 `db:"upper_label" json:"upper_label,omitempty"`           // 得分最高的标签id
	LowerLabel      int32 `db:"lower_label" json:"lower_label,omitempty"`           // 得分最低的标签id
	RelationshipKey int32 `db:"relationship_key" json:"relationship_key,omitempty"` // 该标签组合，匹配到的关系id
}

type FeedCommentCounter struct {
	FeedCommentId int64 `db:"feed_comment_id" json:"feed_comment_id,omitempty"`
	LikeNum       int32 `db:"like_num" json:"like_num,omitempty"` // 评论点赞数
}

type FootholdConflict struct {
	DataType string `db:"data_type" json:"data_type,omitempty"`
	Num      int32  `db:"num" json:"num,omitempty"`
	Content  string `db:"content" json:"content,omitempty"`
}

type FortuneContent struct {
	Id                int32  `db:"id" json:"id,omitempty"`
	ContentKey        string `db:"content_key" json:"content_key,omitempty"`
	ContentFortune    string `db:"content_fortune" json:"content_fortune,omitempty"`
	AliasFortune      string `db:"alias_fortune" json:"alias_fortune,omitempty"`
	AdviceFortune     string `db:"advice_fortune" json:"advice_fortune,omitempty"`
	LabelList         string `db:"label_list" json:"label_list,omitempty"`
	ImportanceFortune int32  `db:"importance_fortune" json:"importance_fortune,omitempty"`
	GoodFortune       int32  `db:"good_fortune" json:"good_fortune,omitempty"`
}

type FortuneDayLabels struct {
	LabelKey    int32  `db:"label_key" json:"label_key,omitempty"`         // 从1开始，给这里的标签详情一套编号
	Label       int32  `db:"label" json:"label,omitempty"`                 // 日运的标签，枚举来自固定顺序，目前分别是1到7
	Score       int32  `db:"score" json:"score,omitempty"`                 // 指定的1-3，这3种标签才有分数，其它的为0
	Good        int32  `db:"good" json:"good,omitempty"`                   // 0 平，1 吉，2 凶
	Annotation  string `db:"annotation" json:"annotation,omitempty"`       // 注释，用于参考编写文案
	EventId     int32  `db:"event_id" json:"event_id,omitempty"`           // 对应的事件id
	EventTypeId int32  `db:"event_type_id" json:"event_type_id,omitempty"` // 对应事件类型id，和event_id属于同一张表的 一个字段
}

type LuckySymbol struct {
	Classify  string `db:"classify" json:"classify,omitempty"`     // 类别
	Num       int32  `db:"num" json:"num,omitempty"`               // 在类别内编号
	OriginUrl string `db:"origin_url" json:"origin_url,omitempty"` // 初始url
}

type MessageMe struct {
	Mid       int64      `db:"mid" json:"mid,omitempty"`
	Uid       int64      `db:"uid" json:"uid,omitempty"`               // 用户ID
	Fid       int64      `db:"fid" json:"fid,omitempty"`               // feed编号
	MsgType   int32      `db:"msg_type" json:"msg_type,omitempty"`     // 消息类型 1 点赞；2 评论；3 收藏
	MsgTime   *time.Time `db:"msg_time" json:"msg_time"`               // 点赞，评论，收藏时间
	MsgStatus int32      `db:"msg_status" json:"msg_status,omitempty"` // 0:无效 1.有效
}

type PushActivity struct {
	Id          int64      `db:"id" json:"id,omitempty"`                     // 流水id
	Title       string     `db:"title" json:"title,omitempty"`               // 活动标题
	Content     string     `db:"content" json:"content,omitempty"`           // 活动内容
	ImageUrl    string     `db:"image_url" json:"image_url,omitempty"`       // 活动页面图片地址
	ActivityUrl string     `db:"activity_url" json:"activity_url,omitempty"` // 活动跳转地址
	Status      int32      `db:"status" json:"status,omitempty"`             // 活动状态：0、下线 1.上线
	CreateTime  *time.Time `db:"create_time" json:"create_time"`             // 创建时间
}

type TopicInfo struct {
	Id          int64      `db:"id" json:"id,omitempty"`                     // 话题流水id
	Title       string     `db:"title" json:"title,omitempty"`               // 话题标题
	SubTitle    string     `db:"sub_title" json:"sub_title,omitempty"`       // 话题子标题
	Cover       string     `db:"cover" json:"cover,omitempty"`               // 话题封面
	Top         int32      `db:"top" json:"top,omitempty"`                   // 置顶
	TopicStatus int32      `db:"topic_status" json:"topic_status,omitempty"` // 1.上架 2.下架 3.删除
	TopicTime   *time.Time `db:"topic_time" json:"topic_time"`               // 话题更新时间
}

type DirtyWord struct {
	Id   int32  `db:"id" json:"id,omitempty"`
	Word string `db:"word" json:"word,omitempty"` // 脏字
}

type Feed struct {
	Fid             int64      `db:"fid" json:"fid,omitempty"`                             // 评论信息表id
	Uid             int64      `db:"uid" json:"uid,omitempty"`                             // 发布用户ID
	Topics          string     `db:"topics" json:"topics,omitempty"`                       // 标签
	Content         string     `db:"content" json:"content,omitempty"`                     // 主要内容
	Location        string     `db:"location" json:"location,omitempty"`                   // 发布内容地址
	Mood            string     `db:"mood" json:"mood,omitempty"`                           // 心情
	Images          string     `db:"images" json:"images,omitempty"`                       // 评论图片url
	ThumbImages     string     `db:"thumb_images" json:"thumb_images,omitempty"`           // feed 缩略图
	SchemeF         string     `db:"scheme_f" json:"scheme_f,omitempty"`                   // 跳转
	VideoUrl        string     `db:"video_url" json:"video_url,omitempty"`                 // 视频cdn地址
	VideoCover      string     `db:"video_cover" json:"video_cover,omitempty"`             // 视频conver
	VideoThumbCover string     `db:"video_thumb_cover" json:"video_thumb_cover,omitempty"` // 视频缩略图
	VideoWidth      int32      `db:"video_width" json:"video_width,omitempty"`             // 视频宽度
	VideoHeight     int32      `db:"video_height" json:"video_height,omitempty"`           // 视频高度
	AudioUrl        string     `db:"audio_url" json:"audio_url,omitempty"`                 // 音频地址
	AudioTime       float64    `db:"audio_time" json:"audio_time"`                         // 音频时间
	Permission      byte       `db:"permission" json:"permission"`                         // 权限，0代表广场可见，1表示主页可见，2表示自己可见，3表示陌生人可见
	FeedType        byte       `db:"feed_type" json:"feed_type"`                           // feed 内容分类 1. 文字 2.图片 3.音频 4.视频
	FeedTime        *time.Time `db:"feed_time" json:"feed_time"`                           // 发布时间
	FeedStatus      byte       `db:"feed_status" json:"feed_status"`                       // 0:正常 1.下架
	RndName         string     `db:"rnd_name" json:"rnd_name,omitempty"`
	PicNum          byte       `db:"pic_num" json:"pic_num"`
	VideoNum        byte       `db:"video_num" json:"video_num"`
	ContentNum      int32      `db:"content_num" json:"content_num,omitempty"`
	Score           float64    `db:"score" json:"score"` // feed内容得分
	MsgTime         *time.Time `db:"msg_time" json:"msg_time"`
	OriginFid       string     `db:"origin_fid" json:"origin_fid,omitempty"` // mock数据源对应的fid
}

type User struct {
	Uid                 int64      `db:"uid" json:"uid,omitempty"`                           // 用户唯一编号
	NickName            string     `db:"nick_name" json:"nick_name,omitempty"`               // 用户名
	Mobile              string     `db:"mobile" json:"mobile,omitempty"`                     // 手机号码
	Cover               string     `db:"cover" json:"cover,omitempty"`                       // 背景颜色
	AvatarIcon          string     `db:"avatar_icon" json:"avatar_icon,omitempty"`           // 头像icon
	Sex                 byte       `db:"sex" json:"sex"`                                     // 性别 0 女 1 男
	Instruction         string     `db:"instruction" json:"instruction,omitempty"`           // 个性签名，自我介绍
	FromW               string     `db:"from_w" json:"from_w,omitempty"`                     // 来自###
	ImPermission        byte       `db:"im_permission" json:"im_permission"`                 // Im聊天权限: 0 .正常,1禁用
	PublisherPermission byte       `db:"publisher_permission" json:"publisher_permission"`   // 信息流发布权限：0正常,1禁用
	SchemeU             string     `db:"scheme_u" json:"scheme_u,omitempty"`                 // 跳转到用户的url
	RoleType            int32      `db:"role_type" json:"role_type,omitempty"`               // 用户类型 0 普通用户； 1.后台用户；2.模拟用户 ;3.编辑用户 4.mock数据用户
	Pwd                 string     `db:"pwd" json:"pwd,omitempty"`                           // 后台管理员登入系统密码
	RegTime             *time.Time `db:"reg_time" json:"reg_time"`                           // 注册时间
	OriginUid           string     `db:"origin_uid" json:"origin_uid,omitempty"`             // mock用户uid，扩展字段
	UserStatus          byte       `db:"user_status" json:"user_status"`                     // 用户状态 0、正常 1、封禁
	EditProfile         int32      `db:"edit_profile" json:"edit_profile,omitempty"`         // 1表示未编辑，2表示已编辑，3表示未知
	ArchiveDemoDel      int32      `db:"archive_demo_del" json:"archive_demo_del,omitempty"` // 大春子档案是否删除 0.未删除 1.已经删除
	Birthday            *time.Time `db:"birthday" json:"birthday"`                           // 用户生日
	DeviceType          string     `db:"device_type" json:"device_type,omitempty"`           // 用户设备id
	WxId                string     `db:"wx_id" json:"wx_id,omitempty"`                       // 微信开放平台id
	QqId                string     `db:"qq_id" json:"qq_id,omitempty"`                       // qq开放平台id
	WbId                string     `db:"wb_id" json:"wb_id,omitempty"`                       // 微博开放平台id
	ShortId             int64      `db:"short_id" json:"short_id,omitempty"`
	ChartTag            int32      `db:"chart_tag" json:"chart_tag,omitempty"`
	ChannelName         string     `db:"channel_name" json:"channel_name,omitempty"`
	AppVersionName      string     `db:"app_version_name" json:"app_version_name,omitempty"`
	Platform            byte       `db:"platform" json:"platform"` // 1. android 2. appstore
	Vip                 byte       `db:"vip" json:"vip"`           // 1. 会员 2.非会员
	DeviceId            string     `db:"device_id" json:"device_id,omitempty"`
}

type V2AstrolabeMoon struct {
	PhaseKey string `db:"phase_key" json:"phase_key,omitempty"` // 星象key
	Title    string `db:"title" json:"title,omitempty"`         // 关键词
	Content  string `db:"content" json:"content,omitempty"`     // 文案
}

type CliConf struct {
	Id          int32  `db:"id" json:"id,omitempty"`                     // 流水ID
	ConfName    string `db:"conf_name" json:"conf_name,omitempty"`       // 配置名称
	ConfType    string `db:"conf_type" json:"conf_type,omitempty"`       // 配置类型 android，ios
	ConfContent string `db:"conf_content" json:"conf_content,omitempty"` // 配置内容 json格式
	Who         string `db:"who" json:"who,omitempty"`                   // 负责人
	UpdateTime  string `db:"update_time" json:"update_time,omitempty"`   // 修改时间
}

type CompensateInfo struct {
	Id               int64      `db:"id" json:"id,omitempty"`                               // 流水id
	Uid              int64      `db:"uid" json:"uid,omitempty"`                             // 补偿用户id
	NickName         string     `db:"nick_name" json:"nick_name,omitempty"`                 // 补偿用户昵称
	GoodsId          int64      `db:"goods_id" json:"goods_id,omitempty"`                   // 商品id
	GoodsName        string     `db:"goods_name" json:"goods_name,omitempty"`               // 商品名称
	KeepMonth        int32      `db:"keep_month" json:"keep_month,omitempty"`               // 商品持续时间
	CompensateStatus int32      `db:"compensate_status" json:"compensate_status,omitempty"` // 1. 补偿成功 2.补偿失败
	CompensateDes    string     `db:"compensate_des" json:"compensate_des,omitempty"`       // 补偿说明
	CompensateErr    string     `db:"compensate_err" json:"compensate_err,omitempty"`       // 补偿失败错误描述
	CompensateTime   *time.Time `db:"compensate_time" json:"compensate_time"`               // 补偿时间
}

type FeedCommentLike struct {
	FeedCommentId int64 `db:"feed_comment_id" json:"feed_comment_id,omitempty"` // 评论ID
	Uid           int64 `db:"uid" json:"uid,omitempty"`                         // 评论uid点赞详情记录
	Operate       byte  `db:"operate" json:"operate"`                           // 1.点赞 2.取消点赞
}

type Message struct {
	Mid        int64   `db:"mid" json:"mid,omitempty"`                 // 流水ID
	Uid        int64   `db:"uid" json:"uid,omitempty"`                 // 消息通知的用户ID
	Fid        int64   `db:"fid" json:"fid,omitempty"`                 // 瞬间ID
	BehaviorId int64   `db:"behavior_id" json:"behavior_id,omitempty"` // 行为ID（消息产生者），这里只记录评论相关的uid
	Content    string  `db:"content" json:"content,omitempty"`         // 消息详情
	Image      string  `db:"image" json:"image,omitempty"`             // 瞬间所对应的image(视频，对应的cover)
	AudioTime  float64 `db:"audio_time" json:"audio_time"`             // 音频长度
	Txt        string  `db:"txt" json:"txt,omitempty"`                 // 文字
	VideoImage string  `db:"video_image" json:"video_image,omitempty"` // 视频封面
	MsgScheme  string  `db:"msg_scheme" json:"msg_scheme,omitempty"`   // 消息跳转的scheme
	MsgType    int32   `db:"msg_type" json:"msg_type,omitempty"`       // 1.我的feed点赞 2.评论 3.被评论
	IsRead     byte    `db:"is_read" json:"is_read"`                   // 消息状态，默认未读取
}

type PhaseNotFoundPlaceholder struct {
	Id      int32  `db:"id" json:"id,omitempty"`           // 用于过滤相位的行星id
	Content string `db:"content" json:"content,omitempty"` // 填充文案
}

type PlanetConstellation struct {
	Id                     int32      `db:"id" json:"id,omitempty"`                                             // 流水ID
	PlanetConstellationKey string     `db:"planet_constellation_key" json:"planet_constellation_key,omitempty"` // 行星落在星座key
	Content                string     `db:"content" json:"content,omitempty"`                                   // 行星落在星座详细描述
	SourceFrom             string     `db:"source_from" json:"source_from,omitempty"`                           // 数据来源
	DateTime               *time.Time `db:"date_time" json:"date_time"`                                         // 修改时间
	HoroElement            int32      `db:"horo_element" json:"horo_element,omitempty"`                         // 星座元素属性，共4个枚举值，100到103分别为 火 土 风 水
	HoroMode               int32      `db:"horo_mode" json:"horo_mode,omitempty"`                               // 星座模式，共3个枚举值，100到102分别为 创始 固定 变动
	HoroGuard              int32      `db:"horo_guard" json:"horo_guard,omitempty"`                             // 守护星，共10个枚举值，100到109，分别为太阳到冥王星的10个星体
	EnName                 string     `db:"en_name" json:"en_name,omitempty"`                                   // 英文名，仅太阳星座的12条数据有
	Alias                  string     `db:"alias" json:"alias,omitempty"`                                       // 别名，仅太阳星座的12条数据有
	Label                  string     `db:"label" json:"label,omitempty"`                                       // 情绪标签，仅月亮星座的12条数据有
}

type ArchiveStar struct {
	Uid           int64      `db:"uid" json:"uid,omitempty"`                       // 用户ID
	Aid           int64      `db:"aid" json:"aid,omitempty"`                       // 档案id
	FileName      string     `db:"file_name" json:"file_name,omitempty"`           // 档案名称
	Icon          string     `db:"icon" json:"icon,omitempty"`                     // 用户icon
	RelationShip  byte       `db:"relation_ship" json:"relation_ship"`             // ta和你的关系:0 自己，1 恋人， 2 朋友，3 亲友，4 工作，5 客户，6 案例，7其他
	Sex           byte       `db:"sex" json:"sex"`                                 // 性别 ：1 男，2 女 ；
	BirthDay      *time.Time `db:"birth_day" json:"birth_day"`                     // 出生时间（公历）2006-01-02 15:04
	TimeZone      float64    `db:"time_zone" json:"time_zone"`                     // 时区，如东八区 8
	BirthAddress  string     `db:"birth_address" json:"birth_address,omitempty"`   // 出生地点(本命盘地点)
	BirthLng      float64    `db:"birth_lng" json:"birth_lng"`                     // 经度
	BirthLat      float64    `db:"birth_lat" json:"birth_lat"`                     // 纬度
	LivingPlacing string     `db:"living_placing" json:"living_placing,omitempty"` // 现居住地址
	LivingLng     float64    `db:"living_lng" json:"living_lng"`                   // 现居地经度
	Horoscope     byte       `db:"horoscope" json:"horoscope"`                     // 星座名称
	LivingLat     float64    `db:"living_lat" json:"living_lat"`                   // 现居地纬度
	UpdateTime    *time.Time `db:"update_time" json:"update_time"`                 // 更新时间
}

type BannerInfo struct {
	Id          int64      `db:"id" json:"id,omitempty"`         // 流水id
	Title       string     `db:"title" json:"title,omitempty"`   // banner 标题
	Cover       string     `db:"cover" json:"cover,omitempty"`   // 封面
	BannerType  byte       `db:"banner_type" json:"banner_type"` // banner类型，用于区分banner展示的位置 1.星文banner 2.消息通知banner 3.feed信息流banner
	LinkType    byte       `db:"link_type" json:"link_type"`     // banner跳转类型 ；1.星文地址 2.网页地址
	LinkUrl     string     `db:"link_url" json:"link_url,omitempty"`
	Status      byte       `db:"status" json:"status"`                       // 1. 上架 2.下架
	Platform    byte       `db:"platform" json:"platform"`                   // 0. 所有平台 1.android 2.ios;
	ChannelName string     `db:"channel_name" json:"channel_name,omitempty"` // 渠道，适配该渠道才显示
	Weight      byte       `db:"weight" json:"weight"`                       // 1~9: 值越小，权重越大
	CreateTime  *time.Time `db:"create_time" json:"create_time"`             // 创建时间
	StartTime   *time.Time `db:"start_time" json:"start_time"`               // 开始时间
	EndTime     *time.Time `db:"end_time" json:"end_time"`                   // 结束时间
}

type WeekFortuneIndices struct {
	Id       int64  `db:"id" json:"id,omitempty"`               // 流水id
	Year     int32  `db:"year" json:"year,omitempty"`           // 年
	Week     int32  `db:"week" json:"week,omitempty"`           // 周
	DateRage string `db:"date_rage" json:"date_rage,omitempty"` // 对应周日期
	EditNum  int32  `db:"edit_num" json:"edit_num,omitempty"`   // 当前数据编辑数
	Status   int32  `db:"status" json:"status,omitempty"`       // 0、未审核 1.已审核
}

type PlanetHouse struct {
	Id             int32      `db:"id" json:"id,omitempty"`                             // 流水ID
	PlanetHouseKey string     `db:"planet_house_key" json:"planet_house_key,omitempty"` // 行星宫位key
	Content        string     `db:"content" json:"content,omitempty"`                   // 行星宫位信息描述
	SourceFrom     string     `db:"source_from" json:"source_from,omitempty"`           // 数据来源
	DateTime       *time.Time `db:"date_time" json:"date_time"`                         // 修改时间
}

type V2AstrolabeSun struct {
	PhaseKey    string `db:"phase_key" json:"phase_key,omitempty"`       // 星象key
	Title       string `db:"title" json:"title,omitempty"`               // 关键词
	Content     string `db:"content" json:"content,omitempty"`           // 文案
	HoroElement int32  `db:"horo_element" json:"horo_element,omitempty"` // 星座元素属性
	HoroMode    int32  `db:"horo_mode" json:"horo_mode,omitempty"`       // 星座模式
	HoroGuard   int32  `db:"horo_guard" json:"horo_guard,omitempty"`     // 守护星
}

type FortuneLuckyBasics struct {
	BasicId int32  `db:"basic_id" json:"basic_id,omitempty"` // 基础运势项目id
	Type    int32  `db:"type" json:"type,omitempty"`         // 颜色1 数字2 星座3 方位4
	Title   string `db:"title" json:"title,omitempty"`       // 名称
	Icon    string `db:"icon" json:"icon,omitempty"`         // 图标地址的后缀
}

type FortuneMonthYearLabels struct {
	PhaseKey string `db:"phase_key" json:"phase_key,omitempty"` // 星象key
	Label    int32  `db:"label" json:"label,omitempty"`         // 标签id
	Good     int32  `db:"good" json:"good,omitempty"`           // 标签对应吉凶，0 平，1 吉，2 凶
	Score    int32  `db:"score" json:"score,omitempty"`         // 标签对应打分
	Id       int32  `db:"id" json:"id,omitempty"`               // 流水id
}

type GoodsInfo struct {
	Id         int64      `db:"id" json:"id,omitempty"`                 // 流水id
	SellId     string     `db:"sell_id" json:"sell_id,omitempty"`       // 商品售卖id,用于支持apple支付，其格式：包名.goods_type.id
	Name       string     `db:"name" json:"name,omitempty"`             // 商品名字
	GoodsType  byte       `db:"goods_type" json:"goods_type"`           // 1. 新人专享 2.揪揪商品信息
	Detail     string     `db:"detail" json:"detail,omitempty"`         // 商品描述
	Money      string     `db:"money" json:"money,omitempty"`           // 单位 元
	KeepMonth  int32      `db:"keep_month" json:"keep_month,omitempty"` // 产品有效时间 单位 月
	Status     byte       `db:"status" json:"status"`                   // 商品信息状态 1.正常 2.删除
	CreateTime *time.Time `db:"create_time" json:"create_time"`         // 创建时间
}

type SmsPhoneBlock struct {
	Id     int64      `db:"id" json:"id,omitempty"`
	Mobile string     `db:"mobile" json:"mobile,omitempty"`
	InDate *time.Time `db:"in_date" json:"in_date"`
}

type ConstellationElement struct {
	Id      int32  `db:"id" json:"id,omitempty"` // 元素枚举id
	Content string `db:"content" json:"content,omitempty"`
	Name    string `db:"name" json:"name,omitempty"`
}

type FortuneImportantMatchElements struct {
	PhaseKey     string `db:"phase_key" json:"phase_key,omitempty"`         // 星象key
	ElementsList string `db:"elements_list" json:"elements_list,omitempty"` // 对应的幸运元素id列表
	BasicsList   string `db:"basics_list" json:"basics_list,omitempty"`     // 对应的4个基础运势项目id列表
}

type FortuneYearMatchContent struct {
	ScoreFloor float64 `db:"score_floor" json:"score_floor"`     // 分数下限
	ScoreUpper float64 `db:"score_upper" json:"score_upper"`     // 分数上限
	Label      int32   `db:"label" json:"label,omitempty"`       // 对应标签id
	Good       int32   `db:"good" json:"good,omitempty"`         // 对应标签吉凶，0 平，1 吉，2 凶
	Content0   string  `db:"content0" json:"content0,omitempty"` // 年运短文案
	Content1   string  `db:"content1" json:"content1,omitempty"` // 年运固定文案

	Content2 string `db:"content2" json:"content2,omitempty"` // 年运长文案，好
	Content3 string `db:"content3" json:"content3,omitempty"` // 年运长文案，不好
}

type LcTencentRegions struct {
	RegionId   int32   `db:"region_id" json:"region_id,omitempty"`     // 主键 ID
	RegionName string  `db:"region_name" json:"region_name,omitempty"` // 地址全名
	Level      int32   `db:"level" json:"level,omitempty"`             // 行政层级
	Parent     int32   `db:"parent" json:"parent,omitempty"`
	RegionCode int32   `db:"region_code" json:"region_code,omitempty"` // 行政号码
	Pinyin     string  `db:"pinyin" json:"pinyin,omitempty"`           // 拼音
	Name       string  `db:"name" json:"name,omitempty"`               // 名称
	Lat        float64 `db:"lat" json:"lat"`                           // 纬度
	Lng        float64 `db:"lng" json:"lng"`                           // 经度
	Cidx       string  `db:"cidx" json:"cidx,omitempty"`
}

type V2AstrolabePlanetMan struct {
	PlanetId int32  `db:"planet_id" json:"planet_id,omitempty"` // 行星id
	Name     string `db:"name" json:"name,omitempty"`           // 行星人名称
	Title    string `db:"title" json:"title,omitempty"`         // 短文案
	Content  string `db:"content" json:"content,omitempty"`     // 长文案
}

type FortuneMajorDetails struct {
	Alias     string `db:"alias" json:"alias,omitempty"`           // 强平弱的中文解读，如 强强弱平
	DetailKey string `db:"detail_key" json:"detail_key,omitempty"` // 重点运势详细解读的key，格式为 19_0:20_2:17_1:18_1
	Content   string `db:"content" json:"content,omitempty"`       // 详细文案
}

type FortuneMonthYearPhases struct {
	PhaseKey  string `db:"phase_key" json:"phase_key,omitempty"` // 星象key
	Good      int32  `db:"good" json:"good,omitempty"`           // 运势吉凶，0 平，1 吉，2 凶
	Title     string `db:"title" json:"title,omitempty"`         // 运势标题
	Content   string `db:"content" json:"content,omitempty"`     // 运势文案
	Advice    string `db:"advice" json:"advice,omitempty"`       // 运势建议
	Important int32  `db:"important" json:"important,omitempty"` // 月运中的运势重要程度打分
}

type FortuneHouse struct {
	Id             int32  `db:"id" json:"id,omitempty"`                             // 流水id
	PlanetHouseKey string `db:"planet_house_key" json:"planet_house_key,omitempty"` // 星体-宫位 key
	ContentWeek    string `db:"content_week" json:"content_week,omitempty"`         // 周运对应文案
	AliasWeek      string `db:"alias_week" json:"alias_week,omitempty"`             // 周运对应别称
	ImportanceWeek int32  `db:"importance_week" json:"importance_week,omitempty"`   // 周运对应重要程度
	GoodWeek       int32  `db:"good_week" json:"good_week,omitempty"`               // 周运对应吉凶程度
	LabelList      string `db:"label_list" json:"label_list,omitempty"`
}

type MessageLike struct {
	Mid        int64   `db:"mid" json:"mid,omitempty"`                 // 流水ID
	Fid        int64   `db:"fid" json:"fid,omitempty"`                 // 瞬间ID
	BehaviorId int64   `db:"behavior_id" json:"behavior_id,omitempty"` // 行为ID（消息产生者）
	Content    string  `db:"content" json:"content,omitempty"`         // 消息详情
	Image      string  `db:"image" json:"image,omitempty"`             // 瞬间所对应的image(视频，对应的cover)
	AudioTime  float64 `db:"audio_time" json:"audio_time"`             // 音频长度
	Txt        string  `db:"txt" json:"txt,omitempty"`                 // 文字
	VideoImage string  `db:"video_image" json:"video_image,omitempty"`
	MsgScheme  string  `db:"msg_scheme" json:"msg_scheme,omitempty"` // 消息跳转的scheme
}

type PlanetDesc struct {
	Id      int32  `db:"id" json:"id,omitempty"`           // 行星id
	Content string `db:"content" json:"content,omitempty"` // 行星对应描述文案
}

type PushTask struct {
	Id           int64      `db:"id" json:"id,omitempty"`               // 流水id
	Platform     string     `db:"platform" json:"platform,omitempty"`   // 推送平台:android ,ios ,all
	Title        string     `db:"title" json:"title,omitempty"`         // 推送标题
	Content      string     `db:"content" json:"content,omitempty"`     // 推送内容
	ImageUrl     string     `db:"image_url" json:"image_url,omitempty"` // 推送携带的图片
	Status       int32      `db:"status" json:"status,omitempty"`       // 推送任务状态:0.待审核 1.待发送 2.已发送 3.已取消；
	Extras       string     `db:"extras" json:"extras,omitempty"`       // 跳转携带参数-json格式
	PlanPushTime *time.Time `db:"plan_push_time" json:"plan_push_time"` // 计划推送时间
	CreatTime    *time.Time `db:"creat_time" json:"creat_time"`         // 推送任务创建时间
}

type ArticleUser struct {
	Id         int64      `db:"id" json:"id,omitempty"`                   // 星文作者
	UserName   string     `db:"user_name" json:"user_name,omitempty"`     // 星文作者昵称
	Icon       string     `db:"icon" json:"icon,omitempty"`               // 星文作者icon
	UserStatus int32      `db:"user_status" json:"user_status,omitempty"` // 状态： 0.正常 1.删除
	AddTime    *time.Time `db:"add_time" json:"add_time"`                 // 星文作者添加时间
}

type FortuneDayPhases struct {
	PhaseKey    string `db:"phase_key" json:"phase_key,omitempty"`     // 星象key
	Good        int32  `db:"good" json:"good,omitempty"`               // 吉凶，0 平 1 吉 2 凶
	Score       int32  `db:"score" json:"score,omitempty"`             // 重要程度打分
	LabelKeys   string `db:"label_keys" json:"label_keys,omitempty"`   // 日运标签详情key的集合
	Title       string `db:"title" json:"title,omitempty"`             // 标题
	Content     string `db:"content" json:"content,omitempty"`         // 星象对应文案
	Advice      string `db:"advice" json:"advice,omitempty"`           // 星象对应运势建议
	Inspiration string `db:"inspiration" json:"inspiration,omitempty"` // 灵感
	BgId        int32  `db:"bg_id" json:"bg_id,omitempty"`             // 首页图id
	WeightType  int32  `db:"weight_type" json:"weight_type,omitempty"` // 3 相位，2宫位，1星座，大的权重高
}

type SynastryRelationshipExtend struct {
	ContentKey string `db:"content_key" json:"content_key,omitempty"` // 格式：最适合的关系_特殊关系编号，如(0-17)_07_12 自己-厌恶的人 男 女
	Title      string `db:"title" json:"title,omitempty"`             // 最适合的关系，特殊关系标题
	Content    string `db:"content" json:"content,omitempty"`         // 最适合的关系，特殊关系详情
}

type V2AstrolabeAsc struct {
	PhaseKey string `db:"phase_key" json:"phase_key,omitempty"` // 星象key
	Title    string `db:"title" json:"title,omitempty"`         // 关键词
	Content  string `db:"content" json:"content,omitempty"`     // 文案
}

type V2AstrolabeElements struct {
	ElementId int32  `db:"element_id" json:"element_id,omitempty"` // 元素id
	HoroList  string `db:"horo_list" json:"horo_list,omitempty"`   // 对应星座列表
	HouseList string `db:"house_list" json:"house_list,omitempty"` // 对应宫位列表
	Title     string `db:"title" json:"title,omitempty"`           // 短文案
	Content   string `db:"content" json:"content,omitempty"`       // 长文案
	HoroName  string `db:"horo_name" json:"horo_name,omitempty"`   // 星座元素名
}

type V2SynastryData struct {
	PhaseKey      string `db:"phase_key" json:"phase_key,omitempty"`       // 相位或者宫位key
	Type          int32  `db:"type" json:"type,omitempty"`                 // 0 默认无效值，1 相位，2 宫位
	Labels        string `db:"labels" json:"labels,omitempty"`             // 星象对应标签的集合，逗号分隔
	LabelsGood    string `db:"labels_good" json:"labels_good,omitempty"`   // 星象标签的吉凶集合，逗号分隔
	LabelsScore   string `db:"labels_score" json:"labels_score,omitempty"` // 星象标签对应的评分 0-100，逗号分隔
	Content       string `db:"content" json:"content,omitempty"`           // 星象对应的通用文案
	NewContent    string `db:"new_content" json:"new_content,omitempty"`   // v2.3优化文案
	Id1           int32  `db:"id1" json:"id1,omitempty"`
	Id2           int32  `db:"id2" json:"id2,omitempty"`
	Phase         int32  `db:"phase" json:"phase,omitempty"`
	Good          int32  `db:"good" json:"good,omitempty"`                     // 星象对应吉凶，0 默认无效值，1 吉，2 凶
	ExtendContent string `db:"extend_content" json:"extend_content,omitempty"` // 准了来源的相位描述
}

type Article struct {
	ArticleId     int64      `db:"article_id" json:"article_id,omitempty"` // 星文id
	Url           string     `db:"url" json:"url,omitempty"`               // 星文源地址
	Uid           int64      `db:"uid" json:"uid,omitempty"`               // 发布用户ID
	UserName      string     `db:"user_name" json:"user_name,omitempty"`
	Title         string     `db:"title" json:"title,omitempty"`                   // 星文标题
	ShortContent  string     `db:"short_content" json:"short_content,omitempty"`   // 星文简介
	Content       string     `db:"content" json:"content,omitempty"`               // 主要内容
	ArticleFrom   byte       `db:"article_from" json:"article_from"`               // 星文来源 1.用户发布的星文 2.编辑发布的星文
	Cover         string     `db:"cover" json:"cover,omitempty"`                   // 星文封面
	ArticleScheme string     `db:"article_scheme" json:"article_scheme,omitempty"` // 星文跳转
	ArticleStatus int32      `db:"article_status" json:"article_status,omitempty"` // 1.上架 2.下架
	ArticleTime   *time.Time `db:"article_time" json:"article_time"`               // 星文发布时间
	ArticleLog    string     `db:"article_log" json:"article_log,omitempty"`       // 星文操作记录
	ArticleType   byte       `db:"article_type" json:"article_type"`
	ArticleMood   string     `db:"article_mood" json:"article_mood,omitempty"`
}

type FortunePhase struct {
	Id             int32  `db:"id" json:"id,omitempty"`                             // 自增流水id
	PlanetPhaseKey string `db:"planet_phase_key" json:"planet_phase_key,omitempty"` // 相位key
	ContentWeek    string `db:"content_week" json:"content_week,omitempty"`         // 周运对应文案
	AliasWeek      string `db:"alias_week" json:"alias_week,omitempty"`             // 周运对应别称
	ImportanceWeek int32  `db:"importance_week" json:"importance_week,omitempty"`   // 周运对应重要程度
	GoodWeek       int32  `db:"good_week" json:"good_week,omitempty"`               // 周运对应吉凶程度
	LabelList      string `db:"label_list" json:"label_list,omitempty"`
}

type RouterInfo struct {
	Id           int64  `db:"id" json:"id,omitempty"`                       // 路由id
	Title        string `db:"title" json:"title,omitempty"`                 // 路由标题
	Icon         string `db:"icon" json:"icon,omitempty"`                   // 路由icon
	Path         string `db:"path" json:"path,omitempty"`                   // 路由
	ParentId     int64  `db:"parent_id" json:"parent_id,omitempty"`         // 父节点id,非0表示存在父节点；
	RouterStatus int32  `db:"router_status" json:"router_status,omitempty"` // 0. 正常 1.删除
}

type SynastryRelationship struct {
	RelationshipKey int32  `db:"relationship_key" json:"relationship_key,omitempty"` // 最适合的关系key
	Title           string `db:"title" json:"title,omitempty"`                       // 最适合的关系标题
	Content         string `db:"content" json:"content,omitempty"`                   // 最适合的关系描述
}

type V2AstrolabePlanetScore struct {
	PhaseKey   string `db:"phase_key" json:"phase_key,omitempty"`     // 星象key
	PlanetList string `db:"planet_list" json:"planet_list,omitempty"` // 对应星体列表
	ScoreList  string `db:"score_list" json:"score_list,omitempty"`   // 对应得分列表
}

type FortuneLuckyElements struct {
	EleId   int32  `db:"ele_id" json:"ele_id,omitempty"`   // 元素id，唯一
	Type    int32  `db:"type" json:"type,omitempty"`       // 元素类型，妆容1 化妆品2 搭配女3 活动女4 风格5 搭配男6 活动男7
	Title   string `db:"title" json:"title,omitempty"`     // 元素的标题
	Content string `db:"content" json:"content,omitempty"` // 元素的描述
	Part    int32  `db:"part" json:"part,omitempty"`       // 元素的部件id，妆容1 化妆品2 风格3 衣服4 裤子5 鞋子6 配饰7 活动8
}

type PlanetPhaseCombined struct {
	Id             int32      `db:"id" json:"id,omitempty"`                             // 流水ID
	PlanetPhaseKey string     `db:"planet_phase_key" json:"planet_phase_key,omitempty"` // 行星相位key
	Content        string     `db:"content" json:"content,omitempty"`                   // 行星相位信息描述
	SourceFrom     string     `db:"source_from" json:"source_from,omitempty"`           // 数据来源
	DateTime       *time.Time `db:"date_time" json:"date_time"`                         // 修改时间
	Placeholder    byte       `db:"placeholder" json:"placeholder"`                     // 是否存在占位符
	Good           string     `db:"good" json:"good,omitempty"`                         // 相位是否有利
	Love           string     `db:"love" json:"love,omitempty"`                         // 相位对感情影响
}

type FotruneShouldAvoid struct {
	DataType string `db:"data_type" json:"data_type,omitempty"`
	Num      int32  `db:"num" json:"num,omitempty"`
	Content  string `db:"content" json:"content,omitempty"`
}

type SplashInfo struct {
	SplashId      int64      `db:"splash_id" json:"splash_id,omitempty"`             // 启屏任务id
	Weight        byte       `db:"weight" json:"weight"`                             // 1~9排序权重
	Title         string     `db:"title" json:"title,omitempty"`                     // banner 标题
	Cover         string     `db:"cover" json:"cover,omitempty"`                     // 封面
	ShowNow       byte       `db:"show_now" json:"show_now"`                         // 是否即可展示；
	CutScreenTime int32      `db:"cut_screen_time" json:"cut_screen_time,omitempty"` // 切换屏幕时间超过设定的此值，展示广告（可用）
	LinkType      byte       `db:"link_type" json:"link_type"`                       // 1. http类型 2.自定义类型
	LinkUrl       string     `db:"link_url" json:"link_url,omitempty"`               // 跳转连接
	KeepSeconds   int32      `db:"keep_seconds" json:"keep_seconds,omitempty"`       // 停留时间，单位秒
	Status        byte       `db:"status" json:"status"`                             // 1. 上架 2.下架
	Platform      byte       `db:"platform" json:"platform"`                         // 0. 所有平台 1.android 2.ios;
	ChannelName   string     `db:"channel_name" json:"channel_name,omitempty"`       // 渠道，适配该渠道才显示
	CreateTime    *time.Time `db:"create_time" json:"create_time"`                   // 创建时间
	StartTime     *time.Time `db:"start_time" json:"start_time"`                     // 开始时间
	EndTime       *time.Time `db:"end_time" json:"end_time"`                         // 结束时间
}

type ArchiveH5Share struct {
	Uid    int64  `db:"uid" json:"uid,omitempty"`
	OpenId string `db:"open_id" json:"open_id,omitempty"` // 开放平台Id
}

type FeedCollectionDetail struct {
	Fid int64 `db:"fid" json:"fid,omitempty"` // feed id

	Uid     int64 `db:"uid" json:"uid,omitempty"` // 用户id
	Operate byte  `db:"operate" json:"operate"`   // 1.收藏 2.未收藏
}

type FateEnergy struct {
	Id      int32  `db:"id" json:"id,omitempty"`           // 区间id
	Content string `db:"content" json:"content,omitempty"` // 该区间描述文案
	Remark  string `db:"remark" json:"remark,omitempty"`
}

type FeedCounter struct {
	Fid           int64 `db:"fid" json:"fid,omitempty"`                       // feed id
	LikeNum       int32 `db:"like_num" json:"like_num,omitempty"`             // 点赞数
	CollectionNum int32 `db:"collection_num" json:"collection_num,omitempty"` // 收藏数
	ReadNum       int32 `db:"read_num" json:"read_num,omitempty"`             // 阅读数
	CommentNum    int32 `db:"comment_num" json:"comment_num,omitempty"`       // 评论数
}

type FortuneMajor struct {
	FortuneKey string `db:"fortune_key" json:"fortune_key,omitempty"` // 重点运势ID，格式为 四轴id_运势等级
	Cate       string `db:"cate" json:"cate,omitempty"`               // 重点运势分类
	Content    string `db:"content" json:"content,omitempty"`         // 重点运势详细文案
}

type SynastryDataExtend struct {
	ContentKey string `db:"content_key" json:"content_key,omitempty"` // 星象key_档案关系编号,如(6_7_180)_07_12 自己-厌恶的人 男 女
	Content    string `db:"content" json:"content,omitempty"`         // 星象key在特殊关系下的对应文案
}

type V25EncounterPlanetData struct {
	PlanetId        int32  `db:"planet_id" json:"planet_id,omitempty"`               // 星体id
	PlanetName      string `db:"planet_name" json:"planet_name,omitempty"`           // 星体名称
	PlanetCharacter string `db:"planet_character" json:"planet_character,omitempty"` // 对应性格
	EnergyName      string `db:"energy_name" json:"energy_name,omitempty"`           // 能量名称
	EnergyContent   string `db:"energy_content" json:"energy_content,omitempty"`     // 能量解释
	Content         string `db:"content" json:"content,omitempty"`                   // 行星文案
}

type V2SynastryHarvestData struct {
	GoodLabel int32  `db:"good_label" json:"good_label,omitempty"` // 包含吉最多的标签id
	BadLabel  int32  `db:"bad_label" json:"bad_label,omitempty"`   // 包含凶最多的标签id
	Content   string `db:"content" json:"content,omitempty"`       // 收获文案详情
}

type ApkUpdateInfo struct {
	Id               int32      `db:"id" json:"id,omitempty"`                     // apk流水id
	ChannelName      string     `db:"channel_name" json:"channel_name,omitempty"` // 渠道编号（名称）
	PkgName          string     `db:"pkg_name" json:"pkg_name,omitempty"`         // 包名
	ApkUrl           string     `db:"apk_url" json:"apk_url,omitempty"`
	IsForce          byte       `db:"is_force" json:"is_force"`                               // 是否强制更新
	ApkVersion       string     `db:"apk_version" json:"apk_version,omitempty"`               // 发布apk版本
	ApkVersionNum    string     `db:"apk_version_num" json:"apk_version_num,omitempty"`       // apk版本号
	ApkUpdateContent string     `db:"apk_update_content" json:"apk_update_content,omitempty"` // 更新内容
	Md5              string     `db:"md5" json:"md5,omitempty"`                               // 文件md5
	UpdateTime       *time.Time `db:"update_time" json:"update_time"`                         // 更新时间
}

type ArticleJob struct {
	Id        int64      `db:"id" json:"id,omitempty"`           // 流水id
	JobUrl    string     `db:"job_url" json:"job_url,omitempty"` // 抓取内容url
	JobStatus byte       `db:"job_status" json:"job_status"`     // 作业状态 0：未完成 1.已完成 2.重试多次失败
	JobStart  *time.Time `db:"job_start" json:"job_start"`       // 任务开始时间
	JobTry    int32      `db:"job_try" json:"job_try,omitempty"` // job尝试次数
	JobEnd    *time.Time `db:"job_end" json:"job_end"`           // 任务结束时间
	JobMood   string     `db:"job_mood" json:"job_mood,omitempty"`
	JobType   byte       `db:"job_type" json:"job_type"`
	JobErr    string     `db:"job_err" json:"job_err,omitempty"`
}

type ReviveDays struct {
	Title   string `db:"title" json:"title,omitempty"`     // 标题
	Id      int32  `db:"id" json:"id,omitempty"`           // 编号
	Content string `db:"content" json:"content,omitempty"` // xx之日信息描述
	Url     string `db:"url" json:"url,omitempty"`         // 对应图标url
}

type UserPermission struct {
	Uid        int64      `db:"uid" json:"uid,omitempty"`
	UserName   string     `db:"user_name" json:"user_name,omitempty"`
	UserType   int32      `db:"user_type" json:"user_type,omitempty"`     // 1. 管理员 2.运维员 3.编辑员
	UserStatus int32      `db:"user_status" json:"user_status,omitempty"` // 0.正常 1.删除
	Pwd        string     `db:"pwd" json:"pwd,omitempty"`
	RouterIds  string     `db:"router_ids" json:"router_ids,omitempty"` // 权限id数据,数据源来自router_info
	UpdateTime *time.Time `db:"update_time" json:"update_time"`         // 更新时间
}

type V25EncounterLevelData struct {
	LevelName     int32 `db:"level_name" json:"level_name,omitempty"`         // 邂逅等级
	LevelExp      int32 `db:"level_exp" json:"level_exp,omitempty"`           // 等级经验
	LevelAddition int32 `db:"level_addition" json:"level_addition,omitempty"` // 等级加成
}

type CombinedHistory struct {
	Id   int64 `db:"id" json:"id,omitempty"` // 流水id
	Uid  int64 `db:"uid" json:"uid,omitempty"`
	Uid2 int64 `db:"uid2" json:"uid2,omitempty"`
	Aid1 int64 `db:"aid1" json:"aid1,omitempty"` // 合盘档案id1
	Aid2 int64 `db:"aid2" json:"aid2,omitempty"` // 合盘档案id2

	CombinedType byte       `db:"combined_type" json:"combined_type"` // 合盘类型
	Score        int32      `db:"score" json:"score,omitempty"`       // 契合度得分
	CombinedTime *time.Time `db:"combined_time" json:"combined_time"` // 合盘时间
	Status       byte       `db:"status" json:"status"`               // 合盘记录状态 0.正常 1.删除
	CombinedFrom byte       `db:"combined_from" json:"combined_from"` // 1.微信好友 2.QQ好友
	OpenId       string     `db:"open_id" json:"open_id,omitempty"`
}

type LoginLog struct {
	Uid          int64   `db:"uid" json:"uid,omitempty"`                     // 用户ID
	Mobile       string  `db:"mobile" json:"mobile,omitempty"`               // 登入系统使用的手机号码
	DeviceId     string  `db:"device_id" json:"device_id,omitempty"`         // 手机设备ID
	Brand        string  `db:"brand" json:"brand,omitempty"`                 // 上次登入手机型号，brand+model
	LoginIp      string  `db:"login_ip" json:"login_ip,omitempty"`           // 上次登入的ip,首次登入此字段为空
	LoginTime    string  `db:"login_time" json:"login_time,omitempty"`       // 上次登入时间
	Lng          float64 `db:"lng" json:"lng"`                               // 经度
	Lat          float64 `db:"lat" json:"lat"`                               // 维度
	PublishCover string  `db:"publish_cover" json:"publish_cover,omitempty"` // 最近一次publish对应的url
	LoginAddress string  `db:"login_address" json:"login_address,omitempty"` // 登录地名
}

type ArticleCounter struct {
	Aid     int64 `db:"aid" json:"aid,omitempty"`           // 星文id
	LikeNum int32 `db:"like_num" json:"like_num,omitempty"` // 星文点赞数
	ReadNum int32 `db:"read_num" json:"read_num,omitempty"` // 星文阅读数
}

type ArticleLikeDetail struct {
	Aid     int64 `db:"aid" json:"aid,omitempty"` // 星文ID
	Uid     int64 `db:"uid" json:"uid,omitempty"`
	Operate byte  `db:"operate" json:"operate"` // 1.点赞 2.点赞取消
}

type FeedComment struct {
	Fid                 int64      `db:"fid" json:"fid,omitempty"`                         // 瞬间ID
	FeedCommentId       int64      `db:"feed_comment_id" json:"feed_comment_id,omitempty"` // 评论信息表id
	Uid                 int64      `db:"uid" json:"uid,omitempty"`                         // 发布评论用户ID
	FeedCommentStatus   byte       `db:"feed_comment_status" json:"feed_comment_status"`   // feed评论状态 0.正常 1.下架
	RndName             string     `db:"rnd_name" json:"rnd_name,omitempty"`               // 当前评论用户随机昵称
	Cuid                int64      `db:"cuid" json:"cuid,omitempty"`                       // 评论中被回复的用户ID
	Cname               string     `db:"cname" json:"cname,omitempty"`                     // 回复的随机名称
	Content             string     `db:"content" json:"content,omitempty"`                 // 评论内容 256个汉字
	ImageUrl            string     `db:"image_url" json:"image_url,omitempty"`             // 评论引用的图片url
	PublishTime         *time.Time `db:"publish_time" json:"publish_time"`                 // 发布时间
	OriginFeedCommentId string     `db:"origin_feed_comment_id" json:"origin_feed_comment_id,omitempty"`
}

type FeedLikeDetail struct {
	Fid     int64 `db:"fid" json:"fid,omitempty"` // feed id
	Uid     int64 `db:"uid" json:"uid,omitempty"` // 用户id
	Operate byte  `db:"operate" json:"operate"`   // 1.点赞 2.取消点赞
}

type FortuneEventHistory struct {
	Id       int64  `db:"id" json:"id,omitempty"`               // 流水id
	Uid      int64  `db:"uid" json:"uid,omitempty"`             // 用户id
	Aid      int64  `db:"aid" json:"aid,omitempty"`             // 揪揪档案id
	EventId  int32  `db:"event_id" json:"event_id,omitempty"`   // 揪揪事件ID
	PhaseKey string `db:"phase_key" json:"phase_key,omitempty"` // 揪揪事件对应的星象
	Start    int64  `db:"start" json:"start,omitempty"`         // 星象开始时间
	End      int64  `db:"end" json:"end,omitempty"`             // 星象结束时间
	Now      int64  `db:"now" json:"now,omitempty"`             // 揪揪发生时间
	Strong   byte   `db:"strong" json:"strong"`                 // 增强|减弱
	Good     int32  `db:"good" json:"good,omitempty"`           // 吉平凶
	Icon     string `db:"icon" json:"icon,omitempty"`
	Name     string `db:"name" json:"name,omitempty"` // 被揪人的昵称
}

type FortuneJiujiuEvents struct {
	EventId            int32  `db:"event_id" json:"event_id,omitempty"`                       // 事件id，唯一
	EventName          string `db:"event_name" json:"event_name,omitempty"`                   // 对应的事件名
	EventDesc          string `db:"event_desc" json:"event_desc,omitempty"`                   // 事件描述
	EventRelationships string `db:"event_relationships" json:"event_relationships,omitempty"` // 事件对应关系集合
	EventTypeName      string `db:"event_type_name" json:"event_type_name,omitempty"`         // 对应的事件类型名
	PhaseKey           string `db:"phase_key" json:"phase_key,omitempty"`                     // 本条事件类型对应的星象
	Title              string `db:"title" json:"title,omitempty"`                             // 揪揪事件的标题
	Content            string `db:"content" json:"content,omitempty"`                         // 揪揪事件的描述
	Advice             string `db:"advice" json:"advice,omitempty"`                           // 揪揪事件的建议
}

type Archive struct {
	Aid           int64      `db:"aid" json:"aid,omitempty"`                       // 档案id
	Uid           int64      `db:"uid" json:"uid,omitempty"`                       // 用户ID
	FileName      string     `db:"file_name" json:"file_name,omitempty"`           // 档案名称
	RelationShip  byte       `db:"relation_ship" json:"relation_ship"`             // ta和你的关系:0 自己，1 恋人， 2 朋友，3 亲友，4 工作，5 客户，6 案例，7其他
	Sex           byte       `db:"sex" json:"sex"`                                 // 性别 ：1 男，0 女 ；类别：2 事件
	Icon          string     `db:"icon" json:"icon,omitempty"`                     // 星座对应的图标
	BirthDay      *time.Time `db:"birth_day" json:"birth_day"`                     // 出生时间（公历）2006-01-02 15:04
	TimeZone      float64    `db:"time_zone" json:"time_zone"`                     // 时区，如东八区 8
	BirthAddress  string     `db:"birth_address" json:"birth_address,omitempty"`   // 出生地点(本命盘地点)
	BirthLng      float64    `db:"birth_lng" json:"birth_lng"`                     // 经度
	BirthLat      float64    `db:"birth_lat" json:"birth_lat"`                     // 纬度
	LivingPlacing string     `db:"living_placing" json:"living_placing,omitempty"` // 现居住地址
	LivingLng     float64    `db:"living_lng" json:"living_lng"`                   // 现居地经度
	Horoscope     byte       `db:"horoscope" json:"horoscope"`                     // 星座名称
	LivingLat     float64    `db:"living_lat" json:"living_lat"`                   // 现居地纬度
	UpdateTime    *time.Time `db:"update_time" json:"update_time"`                 // 更新时间
	CombinedFrom  byte       `db:"combined_from" json:"combined_from"`
	ArchiveStatus byte       `db:"archive_status" json:"archive_status"`
	Collection    byte       `db:"collection" json:"collection"`
	BgUrl         string     `db:"bg_url" json:"bg_url,omitempty"`
	OpenId        string     `db:"open_id" json:"open_id,omitempty"`
	Raid          int64      `db:"raid" json:"raid,omitempty"`
}

type ArchiveH5 struct {
	OpenId        string     `db:"open_id" json:"open_id,omitempty"`               // 微信|qq用户ID
	Aid           int64      `db:"aid" json:"aid,omitempty"`                       // 档案id
	FileName      string     `db:"file_name" json:"file_name,omitempty"`           // 档案名称
	Icon          string     `db:"icon" json:"icon,omitempty"`                     // 用户icon
	RelationShip  byte       `db:"relation_ship" json:"relation_ship"`             // ta和你的关系:0 自己，1 恋人， 2 朋友，3 亲友，4 工作，5 客户，6 案例，7其他
	Sex           byte       `db:"sex" json:"sex"`                                 // 性别 ：1 男，2 女 ；
	BirthDay      *time.Time `db:"birth_day" json:"birth_day"`                     // 出生时间（公历）2006-01-02 15:04
	TimeZone      float64    `db:"time_zone" json:"time_zone"`                     // 时区，如东八区 8
	BirthAddress  string     `db:"birth_address" json:"birth_address,omitempty"`   // 出生地点(本命盘地点)
	BirthLng      float64    `db:"birth_lng" json:"birth_lng"`                     // 经度
	BirthLat      float64    `db:"birth_lat" json:"birth_lat"`                     // 纬度
	LivingPlacing string     `db:"living_placing" json:"living_placing,omitempty"` // 现居住地址
	LivingLng     float64    `db:"living_lng" json:"living_lng"`                   // 现居地经度
	Horoscope     byte       `db:"horoscope" json:"horoscope"`                     // 星座名称
	LivingLat     float64    `db:"living_lat" json:"living_lat"`                   // 现居地纬度
	UpdateTime    *time.Time `db:"update_time" json:"update_time"`                 // 更新时间
	CombinedFrom  byte       `db:"combined_from" json:"combined_from"`             // 1.微信 2.QQ
}

type V2SynastryRelationship struct {
	RelationshipKey int32  `db:"relationship_key" json:"relationship_key,omitempty"` // 最适合的关系key
	Title           string `db:"title" json:"title,omitempty"`                       // 最适合的关系标题
	Content         string `db:"content" json:"content,omitempty"`                   // 最适合的关系描述
}

type WeekFortuneData struct {
	Id                  int64  `db:"id" json:"id,omitempty"`                                     // 周运数据索引id
	ConstellationNumber int32  `db:"constellation_number" json:"constellation_number,omitempty"` // 星座编号
	ConstellationName   string `db:"constellation_name" json:"constellation_name,omitempty"`     // 星座名称
	Title               string `db:"title" json:"title,omitempty"`                               // 标题
	Health              int32  `db:"health" json:"health,omitempty"`                             // 健康
	Work                int32  `db:"work" json:"work,omitempty"`                                 // 事业
	Love                int32  `db:"love" json:"love,omitempty"`                                 // 爱情
	Money               int32  `db:"money" json:"money,omitempty"`                               // 财富
	Content             string `db:"content" json:"content,omitempty"`                           // 星座周运详细描述
}

type PurchaseGoodsInfo struct {
	Uid          int64      `db:"uid" json:"uid,omitempty"`           // 用户id
	GoodsId      int64      `db:"goods_id" json:"goods_id,omitempty"` // 商品信息流水id
	PurchaseDate *time.Time `db:"purchase_date" json:"purchase_date"` // 商品购买时间
	ExpiredDate  *time.Time `db:"expired_date" json:"expired_date"`   // 商品失效时间
}

type V2AstrolabeMode struct {
	ModeId    int32  `db:"mode_id" json:"mode_id,omitempty"`       // 模式id
	HoroList  string `db:"horo_list" json:"horo_list,omitempty"`   // 对应星座列表
	HouseList string `db:"house_list" json:"house_list,omitempty"` // 对应宫位列表
	Title     string `db:"title" json:"title,omitempty"`           // 短文案
	Content   string `db:"content" json:"content,omitempty"`       // 长文案
	HoroName  string `db:"horo_name" json:"horo_name,omitempty"`   // 星座模式名
}

type EventInfo struct {
	EventId      int32      `db:"event_id" json:"event_id,omitempty"`         // 事件id
	EventName    string     `db:"event_name" json:"event_name,omitempty"`     // 事件名称
	EventDes     string     `db:"event_des" json:"event_des,omitempty"`       // 事件描述
	EventStatus  int32      `db:"event_status" json:"event_status,omitempty"` // 事件状态 1. 正常 2.删除
	EventAddTime *time.Time `db:"event_add_time" json:"event_add_time"`       // 事件注册事件
}

type FeedTopicCount struct {
	Id        int64      `db:"id" json:"id,omitempty"`
	Topic     string     `db:"topic" json:"topic,omitempty"`           // 话题标签
	FeedCount int32      `db:"feed_count" json:"feed_count,omitempty"` // 话题对应的feed数量
	CountTime *time.Time `db:"count_time" json:"count_time"`           // 统计时间
}

type FollowTopic struct {
	Uid       int64      `db:"uid" json:"uid,omitempty"`     // 用户id
	Topic     string     `db:"topic" json:"topic,omitempty"` // 话题
	Action    byte       `db:"action" json:"action"`         // 1.关注 2.取消
	TopicTime *time.Time `db:"topic_time" json:"topic_time"` // 关注时间
}

type FotruneAstroTips struct {
	Num     int32  `db:"num" json:"num,omitempty"`
	Content string `db:"content" json:"content,omitempty"`
}

type ReportFeed struct {
	ReportId      int64      `db:"report_id" json:"report_id,omitempty"`             // 举报id
	Uid           int64      `db:"uid" json:"uid,omitempty"`                         // 举报人id
	Fid           int64      `db:"fid" json:"fid,omitempty"`                         // feed id
	FeedCommentId int64      `db:"feed_comment_id" json:"feed_comment_id,omitempty"` // 评论id
	ReportType    byte       `db:"report_type" json:"report_type"`                   // 举报类型： 1. feed 举报 ; 2. 评论举报;
	ReportContent byte       `db:"report_content" json:"report_content"`             // 1. 色情低俗； 2. 广告骚扰； 3. 辱骂、人身攻击； 4. 违法内容；
	EventStatus   byte       `db:"event_status" json:"event_status"`                 // 举报事件状态：0、未处理 1、已经处理
	ReportTime    *time.Time `db:"report_time" json:"report_time"`                   // 举报时间
}

type V2SynastryLabelsDesc struct {
	LabelId    int32  `db:"label_id" json:"label_id,omitempty"`       // 标签id
	LabelDesc  string `db:"label_desc" json:"label_desc,omitempty"`   // 该分数段下的描述词
	ScoreFloor int32  `db:"score_floor" json:"score_floor,omitempty"` // 最低分
	ScoreUpper int32  `db:"score_upper" json:"score_upper,omitempty"` // 最高分
}

type ArchivesRecommend struct {
	Id            int64      `db:"id" json:"id,omitempty"`                         // 用户ID
	Aid           int64      `db:"aid" json:"aid,omitempty"`                       // 档案id
	UserName      string     `db:"user_name" json:"user_name,omitempty"`           // 档案名称
	Icon          string     `db:"icon" json:"icon,omitempty"`                     // 用户icon
	RelationShip  byte       `db:"relation_ship" json:"relation_ship"`             // ta和你的关系:0 自己，1 恋人， 2 朋友，3 亲友，4 工作，5 客户，6 案例，7其他
	Sex           byte       `db:"sex" json:"sex"`                                 // 性别 ：1 男，2 女 ；
	BirthDay      *time.Time `db:"birth_day" json:"birth_day"`                     // 出生时间（公历）2006-01-02 15:04
	TimeZone      float64    `db:"time_zone" json:"time_zone"`                     // 时区，如东八区 8
	BirthAddress  string     `db:"birth_address" json:"birth_address,omitempty"`   // 出生地点(本命盘地点)
	BirthLng      float64    `db:"birth_lng" json:"birth_lng"`                     // 经度
	BirthLat      float64    `db:"birth_lat" json:"birth_lat"`                     // 纬度
	LivingPlacing string     `db:"living_placing" json:"living_placing,omitempty"` // 现居住地址
	LivingLng     float64    `db:"living_lng" json:"living_lng"`                   // 现居地经度
	Horoscope     byte       `db:"horoscope" json:"horoscope"`                     // 星座名称
	LivingLat     float64    `db:"living_lat" json:"living_lat"`                   // 现居地纬度
	UpdateTime    *time.Time `db:"update_time" json:"update_time"`                 // 更新时间
	ArchiveType   byte       `db:"archive_type" json:"archive_type"`               // 1.热门 2.明星 3.动漫
	BgUrl         string     `db:"bg_url" json:"bg_url,omitempty"`
	ArchiveStatus int32      `db:"archive_status" json:"archive_status,omitempty"` // 1. 上架 2.下架 3.编辑 4.删除
}

type AsoInfo struct {
	Id         int64      `db:"id" json:"id,omitempty"`                   // 流水id
	AppId      string     `db:"app_id" json:"app_id,omitempty"`           // app store id
	Idfa       string     `db:"idfa" json:"idfa,omitempty"`               // iphone设备标识
	Keywords   string     `db:"keywords" json:"keywords,omitempty"`       // 任务关键字
	TaskSource string     `db:"task_source" json:"task_source,omitempty"` // 任务名称
	Active     int32      `db:"active" json:"active,omitempty"`           // 此设备是否激活 0:未激活 1.激活
	ActiveTime *time.Time `db:"active_time" json:"active_time"`
}

type QuadrantAnalysis struct {
	Id      int32  `db:"id" json:"id,omitempty"`           // 象限编号
	Content string `db:"content" json:"content,omitempty"` // 该象限对应文案描述
	Title   string `db:"title" json:"title,omitempty"`     // 该象限对应的称呼
}

type SynastryNotFoundPlaceholder struct {
	Id      int32  `db:"id" json:"id,omitempty"`           // 用于过滤相位的行星id
	Content string `db:"content" json:"content,omitempty"` // 填充文案
}

type UserBlock struct {
	Id          int64 `db:"id" json:"id,omitempty"`               // 流水id
	Uid         int64 `db:"uid" json:"uid,omitempty"`             // 用户id
	BlockUid    int64 `db:"block_uid" json:"block_uid,omitempty"` // 被拉黑人uid
	BlockStatus byte  `db:"block_status" json:"block_status"`     // 0.取消拉黑 1. 拉黑
}

type V2AstrolabeQuadrant struct {
	QuadrantId int32  `db:"quadrant_id" json:"quadrant_id,omitempty"` // 象限id
	HouseList  string `db:"house_list" json:"house_list,omitempty"`   // 对应宫位列表
	Alias      string `db:"alias" json:"alias,omitempty"`             // 别称
	Title      string `db:"title" json:"title,omitempty"`             // 短文案
	Content    string `db:"content" json:"content,omitempty"`         // 长文案
}

type FortuneAdvice struct {
	Classify string `db:"classify" json:"classify,omitempty"` // 运势类别
	Num      int32  `db:"num" json:"num,omitempty"`           // 在类别内编号
	Content  string `db:"content" json:"content,omitempty"`   // 运势建议详情描述
}

type PlanetPhaseSelf struct {
	Id             int32      `db:"id" json:"id,omitempty"`                             // 流水ID
	PlanetPhaseKey string     `db:"planet_phase_key" json:"planet_phase_key,omitempty"` // 行星相位key
	Content        string     `db:"content" json:"content,omitempty"`                   // 行星相位信息描述
	SourceFrom     string     `db:"source_from" json:"source_from,omitempty"`           // 数据来源
	DateTime       *time.Time `db:"date_time" json:"date_time"`                         // 修改时间
}
