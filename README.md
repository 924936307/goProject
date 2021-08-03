window Goland 2020.3 常用配置 +golang安装 : https://www.cnblogs.com/nickchou/p/14096944.html

go mod使用，包管理，提交git,更新,分支的处理：https://www.jianshu.com/p/760c97ff644c

go 游戏服务器架构：https://github.com/bobohume/gonet
                http://gonet2.github.io/

cellnet go通信架构 ：https://github.com/davyxu/cellnet

游戏架构发展及推荐的人、文章 （推荐） 

goland自动下载所有依赖 : https://blog.csdn.net/qq_17303159/article/details/110087790
go get -d -v ./...
go mod tidy

go mod init xx (初始化 xx是项目名)
go mod tidy （自动更新依赖）
go get github.com/davyxu/cellnet 拉取指定的第三方包

go env 查看配置信息

gonet项目的依赖环境 etcd：https://www.cnblogs.com/xigang8068/p/5786027.html


go做的简单的处理web请求的demo : https://blog.csdn.net/qq_36025814/article/details/106842775

编译运行及goos goarch对应情况：https://www.cnblogs.com/gbat/p/12809390.html

golang测试(功能测试、压力测试、代码覆盖率测试) ：https://www.cnblogs.com/gbat/p/12809319.html

tcp并发，吞吐量，socks5性能测试工具 ： https://github.com/huzhao37/goperf

goland 依赖的Go SDK中的方法报红：更新golang的版本就好了=

