package config

var locallhost string

// Initialize 配置信息初始化
func Initialize() {
	// 触发加载本目录下其他文件中的 init 方法

	// 请将ip地址赋值给locallhost
	locallhost = "0.0.0.0"
}
