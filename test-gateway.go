package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type handle struct {
	host string
	port string
}

func (h *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uri := r.RequestURI
	fmt.Println(uri)
	fmt.Println(strings.Index(uri[1:len(uri)], "/"))
	//if strings.Index(uri[1:len(uri)], "/") <= 1 {
	//	return
	//}

	//	serviceName := uri[1 : strings.Index(uri[1:len(uri)], "/")+1]
	//	 获取user_id
	query := r.URL.Query()
	if len(query["access_token"]) > 0 {
		// 验证access_token的正确性
		//resp, err := http.Get("http://oauth-center/oauth/user/getUser?access_token=" + query["access_token"][0])
		resp, err := http.Get("http://airport.zhiweicloud.com/oauth/user/getUser?access_token=" + query["access_token"][0])
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		var dat map[string]string
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &dat)
		fmt.Println(dat)

		if dat["user_id"] == "" {
			return
		}

		// 检查用户是否有权限
		if strings.Index(uri, "?") > 0 {
			dat["url"] = uri[0:strings.Index(uri, "?")]
		} else {
			dat["url"] = uri
		}
		if query["queryOrderType"] != nil {
			fmt.Println(query["queryOrderType"][0])
			dat["queryOrderType"] = query["queryOrderType"][0]
		}
		paramString, _ := json.Marshal(dat)

		//resp, err = http.Post("https://api.iairportcloud.com/guest-permission/get-user-permission", "application/json", bytes.NewReader(paramString))
		resp, err = http.Post("http://localhost:8080/get-user-permission", "application/json", bytes.NewReader(paramString))
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		var permission map[string]string
		body, _ = ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &permission)
		fmt.Println(permission)
		if permission[dat["url"]] == "false" {
			println(dat["url"] + ":没有权限")
			return
		}

		// 将user_id添加到header中
		r.Header.Set("User-Id", dat["user_id"])
		r.Header.Set("Client-Id", dat["client_id"])
		r.Header.Set("Role-Ids", permission["roleId"])
		fmt.Println(r.Header)
	} else if uri != "/flight-info/updateFlight" {
		println("找不到access_token")
		return
	}

	//	remote, err := url.Parse("http://" + serviceName)
	remote, err := url.Parse("http://192.168.0.143:8080")
	if err != nil {
		// panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(w, r)
}

func startServer() {
	//被代理的服务器host和port
	h := &handle{host: "airport.zhiweicloud.com"}
err := http.ListenAndServe(":8887", h)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}

func main() {
	startServer()
}
