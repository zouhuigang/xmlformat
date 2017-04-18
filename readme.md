### 测试

	go test -v 

 -v参数，如果不加这个参数的话，只会显示错误的测试用例，否则就显示所有的测试用例（成功 + 错误）

### 测试单个文件，一定要带上被测试的原文件

    Go test -v  wechat_test.go wechat.go 


### 测试单个方法

    go test -v -xmlformat_test.run example2


### 安装

	go get github.com/zouhuigang/xmlformat