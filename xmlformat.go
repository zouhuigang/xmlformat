package xmlformat

import (
	"regexp"
	"strings"
)

var (
	reg = regexp.MustCompile(`<([/!]?)([^>]+?)(/?)>`)
	// NL is the newline string used in XML output, define for DOS-convenient.
	NL = "\r\n"
)

// FormatXML will (purly) reformat the XML string in a readable way, without any rewriting/altering the structure
func FormatXML(xmls, prefix, indent string) string {
	src := regexp.MustCompile(`>\s+<`).ReplaceAllString(xmls, "><")

	rf := replaceTag(prefix, indent)
	return (prefix + reg.ReplaceAllStringFunc(src, rf))
}

//压缩xml
/*
js压缩方法：
var str = this.preserveComments ? text
					: text.replace(/\<![ \r\n\t]*(--([^\-]|[\r\n]|-[^\-])*--[ \r\n\t]*)\>/g,"")
					.replace(/[ \r\n\t]{1,}xmlns/g, ' xmlns');
			return  str.replace(/>\s{0,}</g,"><");
*/
func CompressXml(xmls string) string {
	//re := regexp.MustCompile(`( <=>)\\s+( =<)`)
	//xmls = re.ReplaceAllString(xmls, "")
	//xmls = strings.Replace(xmls, "\n", "", -1) //去除换行符
	var re = regexp.MustCompile(`\<![ \r\n\t]*(--([^\-]|[\r\n]|-[^\-])*--[ \r\n\t]*)\>`)
	xmls = re.ReplaceAllString(xmls, "")

	re = regexp.MustCompile(`>\s{0,}<`)
	xmls = re.ReplaceAllString(xmls, "><")
	return xmls
}

// replaceTag returns a closure function to do 's/(?<=>)\s+(?=<)//g; s(<(/?)([^>]+?)(/?)>)($indent+=$3?0:$1?-1:1;"<$1$2$3>"."\n".("  "x$indent))ge' as in Perl
// and deal with comments as well
func replaceTag(prefix, indent string) func(string) string {
	indentLevel := 0
	return func(m string) string {
		parts := reg.FindStringSubmatch(m)
		// $3: A <foo/> tag. No alteration to indentation.
		// $1: A closing </foo> tag. Drop one indentation level
		// else: An opening <foo> tag. Increase one indentation level
		if len(parts[3]) == 0 {
			//print("] " + parts[1] + "-" + parts[3] + ".\n")
			if parts[1] == `/` {
				indentLevel--
			} else if parts[1] != `!` {
				indentLevel++
			}
		} //else {
		//print("] " + parts[1] + parts[2] + parts[3])
		//}
		return "<" + parts[1] + parts[2] + parts[3] + ">" +
			NL + prefix + strings.Repeat(indent, indentLevel)
	}
}
