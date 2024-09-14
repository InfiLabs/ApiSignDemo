/**
 * @author  zhaoliang.liang
 * @date  2024/1/19 0019 11:21
 */

package config

var Config *config

type config struct {
	AppId       string
	SignKey     string
	InfiWbsPath string
}

func init() {
	Config = &config{
		AppId:       "infi",
		SignKey:     "infi",
		InfiWbsPath: "https://api.infi.cn",
	}
}
