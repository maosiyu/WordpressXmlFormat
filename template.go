/**
 * 将 wordpress 博客 导出的xml文件中，HTML语法 格式为 markdown语法
 *
 * @date 2018-09-20
 * @author mao_siyu
 */
package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

/** 声明一个空字典 */
var dictionary map[string]string

/**
 * 入口
 */
func main() {

	// 填充字典
	fillDictionary()

	var fileName string;
	fmt.Println("Please input your full file path: ")
	fmt.Scanln(&fileName)

	// 读取xml文件 并且进行格式化
	xml := readXmlFile(fileName)
	// 将html语法 替换成 markdown语法
	newFileName := strings.Replace(fileName, ".xml", "_markdown.xml", -1)
	// 将xml文件写入磁盘
	writeXmlFile(xml, newFileName)
	//
	fmt.Println("文件已生成：", newFileName)
	fmt.Println("按回车键退出!")
	fmt.Scanln()
}

/**
 * 读取XML文件
 *
 * return 字符串文本
 */
func readXmlFile(fileName string) string {

	// 读取文件
	b, err := ioutil.ReadFile(fileName)
	// 如果有错
	if err != nil {
		fmt.Println(err)
	}
	// 将流转为字符串
	xml := string(b)
	// strings.Replace 返回 xml 的副本，并将副本中的 old 字符串替换为 new 字符串
	// 替换次数为 n 次，如果 n 为 -1，则全部替换
	for key := range dictionary {
		xml = strings.Replace(xml, key, dictionary[key], -1)
	}
	return xml
}

/**
 * 向磁盘写入XML文件
 */
func writeXmlFile(srcFile string, fileName string) {

	d1 := []byte(srcFile)
	err := ioutil.WriteFile(fileName, d1, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

/**
 * 填充字典
 */
func fillDictionary() {

	// 初始化 map
	dictionary = make(map[string]string)
	dictionary["<pre class=\"prism-highlight line-numbers\" data-start=\"1\"><code class=\"language-null\">"] = "```\n"
	dictionary["<pre class=\"line-numbers prism-highlight\" data-start=\"1\"><code class=\"language-null\">"] = "```\n"
	dictionary["<pre><code class=\" line-numbers\">"] = "```\n"

	dictionary["<pre class=\"line-numbers prism-highlight\" data-start=\"1\"><code class=\"language-javascript\">"] = "``` javascript\n"
	dictionary["<pre class=\"line-numbers prism-highlight\" data-start=\"1\"><code class=\"language-js\">"] = "``` javascript\n"
	dictionary["<pre data-language=><code class=\"language-javascript  line-numbers\">"] = "``` javascript\n"
	dictionary["<pre><code class=\"language-javascript  line-numbers\">"] = "``` javascript\n"

	dictionary["<pre class=\"line-numbers prism-highlight\" data-start=\"1\"><code class=\"language-json\">"] = "``` json\n"
	dictionary["<pre><code class=\"language-javascripton  line-numbers\">"] = "``` json\n"
	dictionary["<pre><code class=\"language-json  line-numbers\">"] = "``` json\n"

	dictionary["<pre class=\"line-numbers prism-highlight\" data-start=\"1\"><code class=\"language-java\">"] = "``` java\n"
	dictionary["<pre><code class=\"language-java  line-numbers\">"] = "``` java\n"

	dictionary["<pre class=\"line-numbers prism-highlight\" data-start=\"1\"><code class=\"language-ruby\">"] = "``` ruby\n"
	dictionary["<pre data-language=><code class=\"language-ruby  line-numbers\">"] = "``` ruby\n"
	dictionary["<pre><code class=\"language-ruby  line-numbers\">"] = "``` ruby\n"

	dictionary["<pre class=\"line-numbers prism-highlight\" data-start=\"1\"><code class=\"language-sql\">"] = "``` sql\n"
	dictionary["<pre class=\"line-numbers prism-highlight\" data-start=\"1\"><code class=\"language-powershell\">"] = "``` powershell\n"

	dictionary["<pre class=\"line-numbers prism-highlight\" data-start=\"1\"><code class=\"language-python\">"] = "``` python\n"
	dictionary["<pre><code class=\"language-python  line-numbers\">"] = "``` python\n"

	dictionary["<pre class=\"line-numbers prism-highlight\" data-start=\"1\"><code class=\"language-bash\">"] = "``` bash\n"
	dictionary["<pre class=\"line-numbers prism-highlight\" data-start=\"1\"><code class=\"language-cpp\">"] = "``` cpp\n"
	dictionary["<pre><code class=\"language-c  line-numbers\">"] = "``` c\n"
	dictionary["<pre><code class=\"language-tree  line-numbers\">"] = "``` tree\n"
	dictionary["<pre><code class=\"language-go  line-numbers\">"] = "``` go\n"

	dictionary["<pre class=\"line-numbers prism-highlight\" data-start=\"1\"><code class=\"language-html\">"] = "``` html\n"
	dictionary["<pre data-language=HTML><code class=\"language-markup  line-numbers\">"] = "``` html\n"

	dictionary["<pre><code class=\"language-css  line-numbers\"><style lang=\"scss\">"] = "``` css\n"
	dictionary["<pre><code class=\"language-css  line-numbers\">"] = "``` css\n"

	dictionary["<pre class=\"line-numbers prism-highlight\" data-start=\"1\"><code class=\"language-xml\">"] = "``` xml\n"
	dictionary["<pre data-language=XML><code class=\"language-markup  line-numbers\">"] = "``` xml\n"

	dictionary["<pre><code class=\"language-yml  line-numbers\">"] = "``` yml\n"

	dictionary["</code></pre>"] = "```\n"
	dictionary["<blockquote>"] = "```\n"
	dictionary["</blockquote>"] = "```"

	dictionary["<h1>"] = "\n# "
	dictionary["<h2>"] = "\n## "
	dictionary["<h3>"] = "\n### "
	dictionary["<h4>"] = "\n#### "
	dictionary["<h5>"] = "\n##### "
	dictionary["<h6>"] = "\n###### "

	dictionary["</h1>"] = ""
	dictionary["</h2>"] = ""
	dictionary["</h3>"] = ""
	dictionary["</h4>"] = ""
	dictionary["</h5>"] = ""
	dictionary["</h6>"] = ""

	dictionary["<br />"] = ""
	dictionary["<hr />"] = ""
	dictionary["<ul>"] = ""
	dictionary["<li>"] = "- "
	dictionary["</li>"] = ""
	dictionary["</ul>"] = ""

	dictionary["<strong>"] = "**"
	dictionary["</strong>"] = "**"

	dictionary["&lt;"] = "<"
	dictionary["&gt;"] = ">"
	dictionary["&#062;"] = ">"
	dictionary["&#124;"] = "|"
	dictionary["&#039;"] = "'"
	dictionary["&quot;"] = "\""
	dictionary["&#42;"] = "*"

	dictionary["<code>"] = "`"
	dictionary["</code>"] = "`"

	dictionary["<![CDATA[<![CDATA["] = "<![CDATA["

}
