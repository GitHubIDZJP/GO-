package ginDemo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type auth struct {
	Username string `json:"username"`
	Pwd      string `json:"password"`
}

func MainLogin() {
	Get()
	PostWithJson()
	PostWithUrlencoded()
}

func Get() {
	//get请求
	//http.Get的参数必须是带http://协议头的完整url,不然请求结果为空
	resp, _ := http.Get("http://localhost:8080/login2?username=admin&password=123456")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	fmt.Printf("Get request result: %s\n", string(body))
}

func PostWithJson() {
	//post请求提交json数据
	auths := auth{"admin", "123456"}
	ba, _ := json.Marshal(auths)
	resp, _ := http.Post("http://localhost:8080/login1", "application/json", bytes.NewBuffer([]byte(ba)))
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Post request with json result: %s\n", string(body))
}

func PostWithUrlencoded() {
	//post请求提交application/x-www-form-urlencoded数据
	form := make(url.Values)
	form.Set("username", "admin")
	form.Add("password", "123456")
	resp, _ := http.PostForm("http://localhost:8080/login2", form)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("发布带有 application/x-www-form-urlencoded结果的请求: %s\n", string(body))
}
