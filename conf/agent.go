package conf

// Agent 程序主体配置
type Agent struct {
	// 主目录
	Home string `default:"${ZIYUNIX_HOME=/config}" yaml:"home" json:"home" xml:"home" toml:"home"`
}
