### golang格式化/压缩xml


### 安装

	go get github.com/zouhuigang/xmlformat

### 使用

1.格式化

	xmlstring := `<?xml version="1.0" encoding="UTF-8"?><Book><Page><Stroke duration="1240" startSecond="1439307219" startMillisecond="517" lineWidth="1"><Packet x="171.816298" y="143.266575" pressure="0.000000" deltaTime="0.000000"></Packet><Packet x="183.791436" y="191.167127" pressure="120.000000" deltaTime="80.000000"></Packet></Stroke></Page></Book>`

	x1 := xmlformat.FormatXML(xmlstring, " ", "  ")
	fmt.Println(x1)

2.压缩

	x3 := `<?xml version="1.0" encoding="utf-8"?>


	<Book>


  	<Page>
    <Stroke duration="1240"   startSecond="1439307219" startMillisecond="517" lineWidth="1">   
      <Packet x="171.816298" y="143.266575" pressure="0.000000" deltaTime="0.000000"/>
      <Packet x="183.791436" y="191.167127" pressure="120.000000" deltaTime="80.000000"/>
      <Packet x="194.357735" y="225.683702" pressure="122.000000" deltaTime="189.333334"/>
	</Stroke>
	  </Page>
	</Book>`

	x1 := xmlformat.CompressXml(x3)
	fmt.Println(x1)

注：使用在网站上，尽量用js来写。避免请求服务器

### 测试

	go test -v 

 -v参数，如果不加这个参数的话，只会显示错误的测试用例，否则就显示所有的测试用例（成功 + 错误）

### 测试单个文件，一定要带上被测试的原文件

    Go test -v  wechat_test.go wechat.go 


### 测试单个方法

    go test -v -xmlformat_test.run example2




