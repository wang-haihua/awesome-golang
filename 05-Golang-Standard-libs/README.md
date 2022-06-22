# Golang-Standard-libs

### 输入输出

- **fmt** fmt包实现了类似C语言printf和scanf的格式化I/O。主要分为向外输出内容和获取输入内容两大部分 
- **bufio** bufio包实现了有缓冲的I/O。它包装一个io.Reader或io.Writer接口对象，创建另一个也实现了该接口，且同时还提供了缓冲和一些文本I/O的帮助函数的对象。



### 自动化测试

- **testing** testing 提供对 Go 包的自动化测试的支持。通过 `go test` 命令，能够自动执行形如 `func TestXxx(*testing.T)` 形式的任何函数

