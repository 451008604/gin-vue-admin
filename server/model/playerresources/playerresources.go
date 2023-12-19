package playerresources

type SetFormData struct {
	Input100456   string `json:"input100456"`   // 金币数量
	Input108790   string `json:"input108790"`   // 钻石数量
	Input81625    string `json:"input81625"`    // 体力数量
	Input46774    string `json:"input46774"`    // 经验数量
	Input16999    string `json:"input16999"`    // 工具数量
	Input17994    string `json:"input17994"`    // 护具数量
	Input107097   string `json:"input107097"`   // 油漆数量
	Input89875    string `json:"input89875"`    // 线材数量
	Textarea47524 string `json:"textarea47524"` // 道具资源
	Textarea94373 string `json:"textarea94373"` // 玩家ID列表
}

type MailJson struct {
	Roleid        []uint32    // 角色ID
	Prize         []ItemPrize // 奖励
	DescribeID    int32       // 描述ID
	ExpireTime    uint32      // 过期时间, 时间戳
	StrExpireTime string      // 可读性较强的时间
	AliveTime     uint32      // 存活时间, 单位秒
}

// 道具奖励
type ItemPrize struct {
	ID  int
	Num int
}