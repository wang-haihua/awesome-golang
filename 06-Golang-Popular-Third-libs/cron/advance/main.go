package main

/*
cron对象创建使用了选项模式，支持如下选项：
	WithLocation：指定时区；
	WithParser：使用自定义的解析器；
	WithSeconds：让时间格式支持秒，实际上内部调用了WithParser。
	WithLogger：自定义Logger；
	WithChain：Job 包装器。
*/
