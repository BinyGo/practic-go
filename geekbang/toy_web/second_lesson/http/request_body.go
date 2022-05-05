package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello,this is home page")
}

//http://localhost:8998/body/once?name=biny&age=18
func readyBodyOnce(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "read body failed: %v \n", err)
		// 记住要返回，不然就还会执行后面的代码
		return
	}
	// 类型转换，将 []byte 转换为 string
	fmt.Fprintf(w, "read data: %s \n", string(body)) //read data:

	// 尝试再次读取，啥也读不到，但是也不会报错
	body, err = io.ReadAll(r.Body)
	if err != nil {
		// 不会进来这里
		fmt.Fprintf(w, "read the data one more time got error: %v \n", err)
		return
	}
	fmt.Fprintf(w, "read the data one more time: [%s] and read data length %d \n", string(body), len(body)) //read the data one more time: [] and read data length 0

}

//http://localhost:8998/body/multi?name=biny&age=18
func getBodyIsNil(w http.ResponseWriter, r *http.Request) {
	if r.GetBody == nil {
		fmt.Fprint(w, "GetBody is nil") //GetBody is nil
	} else {
		fmt.Fprint(w, "GetBody not nil")
	}
}

//http://localhost:8998/url/query?name=biny&age=18
func queryParams(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	fmt.Fprintf(w, "Query is %v \n", values) //Query is map[age:[18] name:[biny]]
}

//http://localhost:8998/wholeUrl?name=biny&age=18
func wholeUrl(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(r.URL)
	fmt.Fprint(w, string(data))
	// {
	// 	"Scheme": "",
	// 	"Opaque": "",
	// 	"User": null,
	// 	"Host": "",
	// 	"Path": "/wholeUrl",
	// 	"RawPath": "",
	// 	"ForceQuery": false,
	// 	"RawQuery": "name=biny&age=18",
	// 	"Fragment": "",
	// 	"RawFragment": ""
	// }
}

//http://localhost:8998/header?dd=dd
func header(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "header is %v \n", r.Header)
	//header is map[Accept:[text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9]
	//Accept-Encoding:[gzip, deflate, br] Accept-Language:[zh-CN,zh;q=0.9]
	//Connection:[keep-alive] Cookie:[vipPromorunningtmr=; isvipretainend=; discount_free_trigger=; i_like_gitea=3fea4c0d6ea53ffa; lang=zh-CN;
	//session=eyJ1c2VyX2lkIjo3fQ.Yg-u_g.7vFRKoojkh2BHlgyIxjvtWquCfA; grafana_session=5859d02d9d528039bbde8b40002066f8]
	//Sec-Ch-Ua:[" Not A;Brand";v="99", "Chromium";v="100", "Google Chrome";v="100"]
	//Sec-Ch-Ua-Mobile:[?0] Sec-Ch-Ua-Platform:["Windows"]
	//Sec-Fetch-Dest:[document] Sec-Fetch-Mode:[navigate] Sec-Fetch-Site:[none] Sec-Fetch-User:[?1] Upgrade-Insecure-Requests:[1]
	//User-Agent:[Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36]]
}

//http://localhost:8998/form?name=biny&age=18
func form(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "before parse form %v\n", r.Form) //before parse form map[]
	//要先r.ParseForm(),才能取到r.Form数据
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "parse form error: %v \n", r.Form)
	}
	fmt.Fprintf(w, "after parse form %v \n", r.Form) //after parse form map[age:[18] name:[biny]]
}

func main() {
	http.HandleFunc("/home", home)
	http.HandleFunc("/body/once", readyBodyOnce)
	http.HandleFunc("/body/multi", getBodyIsNil)
	http.HandleFunc("/url/query", queryParams)
	http.HandleFunc("/header", header)
	http.HandleFunc("/wholeUrl", wholeUrl)
	http.HandleFunc("/form", form)
	if err := http.ListenAndServe(":8998", nil); err != nil {
		panic(err)
	}
}
