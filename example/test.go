package main

import (
	"bufio"
	"bytes"
	"container/list"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

//获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。
func WalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix)                                                     //忽略后缀匹配的大小写
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		//if err != nil { //忽略错误
		// return err
		//}
		if fi.IsDir() { // 忽略目录
			return nil
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}
		return nil
	})
	return files, err
}

// 计算字符串中有多少个显示用的字符
func CalcUtf8StringShowNum(s string) (n int) {
	num, chn := CalcUtf8StringNum(s)
	return (chn*2 + num - chn)
}

// 计算一个字符串中有多少个字符，多少个中文
// n: 字符总数
// chn: 中文字符总数
func CalcUtf8StringNum(s string) (n int, chn int) {
	str := s
	for len(str) > 0 {
		_, size := utf8.DecodeRuneInString(str)
		if size > 1 {
			chn++
		}
		str = str[size:]
		n++
	}
	return n, chn
}

// 读取一个文件中所有的行,需要注意的是必须是\n作为换行符
// fn: 文件路径
// lines: 存储所有的行
func ReadAllLines(fn string, lines *list.List) error {
	file, err := os.Open(fn)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(str) > 0 {
				lines.PushBack(str)
			}
			break
		}
		lines.PushBack(str)
	}
	return nil
}

type FormatInterface interface {
	Before(line string, tagName string)
	ToString(line string, tagName string) string
}

func format_some(fn string, out string, tagName string, obj FormatInterface) error {
	lines := list.New()
	rerr := ReadAllLines(fn, lines)
	if rerr != nil {
		return rerr
	}
	for e := lines.Front(); e != nil; e = e.Next() {
		var str string = e.Value.(string)
		if strings.Contains(str, tagName) {
			obj.Before(str, tagName)
		}
	}

	if fn == out {
		os.Remove(fn)
	}

	file, werr := os.OpenFile(out, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if werr != nil {
		return werr
	}
	defer file.Close()
	for e := lines.Front(); e != nil; e = e.Next() {
		var str string = e.Value.(string)
		if strings.Contains(str, tagName) {
			str = obj.ToString(str, tagName)
		}
		file.Write([]byte(str))
	}

	return nil
}

func FormatListAllTag(fn string, dt map[string]string) error {
	bs, err := ioutil.ReadFile(fn)
	if err != nil {
		return err
	}

	decoder := xml.NewDecoder(bytes.NewBuffer(bs))

	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		switch token := t.(type) {
		// 处理元素开始（标签）
		case xml.StartElement:
			name := token.Name.Local
			dt[name] = name
			//fmt.Printf("tag:%s\n", name)

		// 处理元素结束（标签）
		case xml.EndElement:
		case xml.CharData:
		default:
			// ...
		}
	}
	return nil
}

func FormatOneXmlFile(fn string, out string) error {
	dt := make(map[string]string)
	FormatListAllTag(fn, dt)
	bs, err := ioutil.ReadFile(fn)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(out, bs, os.ModePerm)
	if err != nil {
		return err
	}

	for k, _ := range dt {
		obj := NewFormatStruct()
		format_some(out, out, k, obj)
	}

	return nil
}

func FormatXmlPath(inpath string, outpath string) error {
	files, err := WalkDir(inpath, "xml")
	if err != nil {
		return err
	}

	num := len(inpath) + 1
	tmpmp := make(map[string]string)

	// 创建指定的目录
	oup := strings.Replace(outpath, "/", "\\", 100)
	for _, v := range files {
		fn := (oup + "\\" + v[num:])
		idx := strings.LastIndex(fn, "\\")
		fp := fn[0:idx]
		//fmt.Printf("outpath:%s\n", fp)
		if _, ok := tmpmp[fp]; ok == false {
			tmpmp[fp] = fp
			os.MkdirAll(fp, os.ModePerm)
		}
		FormatOneXmlFile(v, fn)
	}
	return nil
}

// 这里是转换成空格的数量
func GetTabNumFormSpaceNum(max int, has int, key string, value string) string {
	num := max - has
	if num == 0 {
		return ""
	}

	ss := ""
	for i := 0; i < num; i++ {
		ss += " "
	}
	return ss
}

// 获取一个xml形式的字符串中的所有属性生成map,list两种形式,list存储的是key而已,并且是有顺序的
// line 输入字符串
// tagName 需要查找的标签
// data 存储所有属性的key,value
// tags 所有属性的key并且保留顺序
func GetAttr(line string, tagName string, data map[string]string, tags *list.List) error {
	decoder := xml.NewDecoder(bytes.NewBuffer([]byte(line)))

	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		switch token := t.(type) {
		// 处理元素开始（标签）
		case xml.StartElement:
			name := token.Name.Local
			if name == tagName {
				for _, attr := range token.Attr {
					attrName := attr.Name.Local
					attrValue := attr.Value
					data[attrName] = attrValue
					//fmt.Printf("An attribute is:_%s_%s\n", attrName, attrValue)
					if tags != nil {
						tags.PushBack(attrName)
					}
				}
			}

		// 处理元素结束（标签）
		case xml.EndElement:
		case xml.CharData:
		default:
			// ...
		}
	}
	return nil
}

//========================================================
type FormatStruct struct {
	FormatInterface
	Tags      *list.List     // 存储所有的属性并且按照顺序来
	Id2MaxNum map[string]int // 存储每个属性和属性值的最大显示字符长度
	Start     string         // 开始字符比如,每行前面有缩进的处理
	End       string         // 这里关键是针对> 或者 />这样的两种情况进行必要的处理的
	IsHas     bool           // 在第一次计算的时候进行标记使用
}

func NewFormatStruct() *FormatStruct {
	clzz := &FormatStruct{
		End:       ">",
		Id2MaxNum: make(map[string]int),
		IsHas:     false}
	return clzz
}

func (self *FormatStruct) Before(line string, tagName string) {
	fidx := strings.Index(line, "<"+tagName)
	if fidx == -1 {
		return
	}
	if self.IsHas == false {
		idx := strings.Index(line, "<"+tagName)
		if idx > -1 {
			if idx == 0 {
				self.Start = ""
			} else {
				self.Start = line[:idx]
			}

			self.IsHas = true
		}
		nidx := strings.Index(line, "/>")
		if nidx > -1 {
			self.End = "/>"
		}
	}

	tags := list.New()

	data := make(map[string]string)
	GetAttr(line, tagName, data, tags)

	if self.Tags == nil {
		self.Tags = tags
	} else {
		if self.Tags.Len() < tags.Len() {
			self.Tags = tags
		}
	}

	for k, v := range data {
		num := CalcUtf8StringShowNum(v)
		nn, ok := self.Id2MaxNum[k]
		if ok {
			if num > nn {
				self.Id2MaxNum[k] = num
			}
		} else {
			self.Id2MaxNum[k] = num
		}
	}
}

func (self *FormatStruct) ToString(line string, tagName string) string {
	fidx := strings.Index(line, "<"+tagName)
	if fidx == -1 {
		return line
	}
	if self.Tags.Len() == 0 {
		return self.Start + "<" + tagName + self.End + "\r\n"
	}

	data := make(map[string]string)
	GetAttr(line, tagName, data, nil)

	var str string = self.Start + "<" + tagName + " "
	for e := self.Tags.Front(); e != nil; e = e.Next() {
		k := e.Value.(string)
		max, _ := self.Id2MaxNum[k]
		val, _ := data[k]
		has := CalcUtf8StringShowNum(val)
		str += k + "=\"" + val + "\" " + GetTabNumFormSpaceNum(max, has, k, val)
	}

	str += self.End + "\r\n"
	return str
}

func Usage() {
	str := os.Args[0]
	fmt.Printf(`
    usage:
        1. ` + str + ` -t [in] [out] [tag]
        2. ` + str + ` -f [in] [out]
        3. ` + str + ` -p [inpath] [output]
        [-t]: format xml <tag line
        [-f]: format one xml file
        [-p]: format one path all xml 

        [in]: input file like g:\\xx\ch.xml
        [out]: output file like g:\\xx\ch.xml_out
        [inpath]: a path for xml input
        [output]: a path for xml output
        [tag]: <root>
                <node id="10" name="haha" go="1,1"/>
            </root>
            tag is node
    `)
}

func main() {
	num := len(os.Args)
	if num < 4 {
		Usage()
		return
	}

	tt := os.Args[1]
	in := os.Args[2]
	out := os.Args[3]

	fmt.Print("==========start format xml=======\n")
	fmt.Printf("input file: %s\nout file:%s\n\n", in, out)
	for {
		if tt == "-t" {
			if num < 5 {
				Usage()
				break
			}
			tag := os.Args[4]
			clzz := NewFormatStruct()
			format_some(in, out, tag, clzz)
		} else if tt == "-f" {
			FormatOneXmlFile(in, out)
		} else if tt == "-p" {
			FormatXmlPath(in, out)
		} else {
			Usage()
		}
		break
	}

	fmt.Print("=================================\n")
}
