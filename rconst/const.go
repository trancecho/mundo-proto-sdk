package rconst

var (
	StreamComment      = "mundo_comment_event" // 评论事件 (发布评论后触发
	NavyGroup          = "navy_group"          // 评论事件消费者组
	NavyConsumerPrefix = "navy_consumer_"      // 评论事件消费者前缀
)

type CommentEvent struct {
	PostId uint `json:"post_id"` // 被评论的帖子ID
}

// StatEventType 统计事件类型常量
const (
	// 用户行为事件
	EventLogin      = "login"       // 登录
	EventLogout     = "logout"      // 登出
	EventFirstLogin = "first_login" // 首次登录
	EventOnline     = "online"      // 上线
	EventOffline    = "offline"     // 下线

	// 页面浏览事件
	EventPageView = "page_view" // 页面浏览
	EventPageStay = "page_stay" // 页面停留

	// 通用交互事件
	EventClick = "click" // 点击

	// 题库相关事件
	EventQuestionSolve = "question_solve" // 做题

	// 聊天相关事件
	EventChatSend = "chat_sent" // 发送聊天消息
)
