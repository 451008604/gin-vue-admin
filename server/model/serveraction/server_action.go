package serveraction

type PlanJson struct {
	WHITE_IP   string // 白名单
	CLOSE_TIME string // 停服时间
	OPEN_TIME  string // 对外开放时间
	TITLE      string // 标题
	BODY       string // 正文
	BUTTON     string // 按钮
}
