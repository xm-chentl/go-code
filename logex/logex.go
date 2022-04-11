package logex

type ILogger interface {
	// Title 标题固定
	Title(format string, args ...interface{}) ILogger
	// Label 标签
	Label(name string, format string, args ...interface{}) ILogger
	// GetLog 获取日志对象
	GetLog() ILog
	// 记录日志, 然后panic;
	Panic(error)
	// Fatal 致命错误, 出现错误时程序无法正常运转. 输出日志后, 程序退出;
	Fatal(error)
	// Error 错误日志, 需要查看原因;
	Error(error)
	// Warn 警告信息, 提醒程序员注意;
	Warn(error)
	// Info 关键操作, 核心流程的日志;
	Info(error)
	// Debug 一般程序中输出的调试信息;
	Debug(error)
	// Trace很细粒度的信息, 一般用不到;
	Trace(error)
}

type ILog interface {
	Logger() ILogger
}
