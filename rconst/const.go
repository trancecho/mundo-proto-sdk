package rconst

var (
	StreamComment      = "mundo_comment_event" // 评论事件 (发布评论后触发
	NavyGroup          = "navy_group"          // 评论事件消费者组
	NavyConsumerPrefix = "navy_consumer_"      // 评论事件消费者前缀
)

type CommentEvent struct {
	PostId  uint `json:"post_id"`  // 被评论的帖子ID
	RobotId uint `json:"robot_id"` // 被评论的帖子的机器人ID
}
