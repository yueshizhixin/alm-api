package conf

/**
	http相关配置
 */

var (
	CheckOrigin bool
	CheckHost   bool
	CheckRefer  bool
	AllowOrigin *[]string
	AllowHost   *[]string
	AllowRefer  *[]string
)

func init() {
	CheckOrigin = true
	CheckHost = true
	CheckRefer = true
	AllowOrigin = &[]string{"https://yueshizhixin.github.io"}
	AllowOrigin = &[]string{"https://yueshizhixin.github.io"}
	AllowRefer = &[]string{"https://yueshizhixin.github.io"}
}
