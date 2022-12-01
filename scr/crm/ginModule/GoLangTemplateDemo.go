package ginModule

import (
	"fmt"
	"html/template"
	"os"
)

func GoLangTemplateDemo() {
	f1()
	//f2()
	//f3()
	//f4()

}

func f4() {
	var textStr = `
		{{if eq "哈哈" .vv}} 
		  vv {{ssv .vv }}
		{{else if eq "test3" .vv }} 
		vvv
		{{else}}
		vvvv
		{{end}}`

	var tf = template.FuncMap{
		"ssv": func(v string) string {
			return v + "123"
		},
	}

	/*
			判断字符串
			如果eq 等于哈哈 则打印 哈哈123(拼接后)
		    不等于则 打印 vvvvv

	*/

	/*
			解析字符串模板
		t, err := template.New("foo").Parse(\{ {define "T"}}Hello, { {.}}!{ {end}}`)` 可以解析字符串模板，并设置它的名字。
	*/
	t := template.Must(template.New("fff").Funcs(tf).Parse(textStr)) //可以解析字符串模板，并设置它的名字
	t.Execute(os.Stdout, map[string]interface{}{
		"vv": "哈哈",
	})
	/*
			打印:
		 vv testvvvvvv


	*/
}

func f3() {
	var testStr = `
      {{range .lists}}
		 {{.}}
      {{end}}
      {{range .listsMap}}
		 {{.userName}}
      {{end}}
`
	t := template.Must(template.New("vv").Parse(testStr))
	param := make(map[string]interface{})
	param["lists"] = []string{"张三", "李四", "赵武"}
	param["listsMap"] = []map[string]interface{}{
		map[string]interface{}{"userName": "张三1"},
		map[string]interface{}{"userName": "张三2"},
		map[string]interface{}{"userName": "张三3"},
	}
	t.Execute(os.Stdout, param)
	/**

	  张三

	  李四

	  赵武


	  张三1

	  张三2

	  张三3




	*/
}

func f2() {
	var ttxt = `thi sis a test for tempalte:
	Name:{{.Name}}
	Age:{{.Age}}
	School:{{.School}}
	Married:{{.MarriedOK}}
`

	var ex = struct {
		Name      string
		Age       int
		School    string
		MarriedOK bool
	}{
		Name:      "小学",
		Age:       18,
		School:    "三大法撒旦法撒旦发送方",
		MarriedOK: true,
	}
	t := template.Must(template.New("index").Parse(ttxt))
	t.Execute(os.Stdout, ex)
	/*
		thi sis a test for tempalte:
		        Name:小学
		        Age:18
		        School:三大法撒旦法撒旦发送方
		        Married:true

	*/
}

func f1() {
	const letter = `
		Dear {{.Name}},
		{{if .Attended}}
		It was a pleasure to see you at the wedding.
		{{- else}}
		It is a shame you couldn't make it to the wedding.
		{{- end}}
		{{with .Gift -}}
		Thank you for the lovely {{.}}.
		{{end}}
		Best wishes,
		Josie
		`
	type Recep struct {
		Name, Gift string
		Attended   bool
	}

	//var receipents = []Recep{
	//	{"张三","李四",true},
	//	{"哈哈","微微",false},
	//}

	var reslice = []map[string]interface{}{
		{"Name": "张三", "Gift": "李四", "Attended": false},
		{"Name": "哈哈", "Gift": "微微", "Attended": true},
	}

	t := template.Must(template.New("letter").Parse(letter))
	for _, r := range reslice {
		//执行
		err := t.Execute(os.Stdout, r)
		if err != nil {
			fmt.Println("execute template :", err)
		}
	}
	/*
	 Dear 张三,

	                It is a shame you couldn't make it to the wedding.
	                Thank you for the lovely 李四.

	                Best wishes,
	                Josie

	                Dear 哈哈,

	                It was a pleasure to see you at the wedding.
	                Thank you for the lovely 微微.

	                Best wishes,
	                Josie

	*/
}
