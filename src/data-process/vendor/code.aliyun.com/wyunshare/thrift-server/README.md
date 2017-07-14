安装
```bash
go get code.aliyun.com/wyunshare/thrift-server
govendor add code.aliyun.com/wyunshare/thrift-server
```
服务端使用

router.go
```go
import (
	"code.aliyun.com/wyunshare/thrift-server/gen-go/server"
	"encoding/json"
)

type BusinessServiceImpl struct {
}

// 通过 BusinessServiceImpl 实现 IBusinessService 接口的 Send 方法，从而实现 IBusinessService 接口
func (msi *BusinessServiceImpl) Handle(operation string, paramJSON []byte) (r *server.Response, err error) {

	paramInput := make(map[string]interface{})
	paramInput["d"] = "a"
	json.Unmarshal(paramJSON, &paramInput)

	dd := make([]byte, 1000)
	dd, err  = json.Marshal(paramInput)

	switch operation {
	case "test":
		a := &server.Response{
			ResponeCode:  server.RESCODE__200,
			ResponseJSON: dd,
		}
		return a, nil
	default:
		return nil, nil
	}

}
```
main.go(单个service)
```go
import "code.aliyun.com/wyunshare/thrift-server"

type BusinessServiceImpl struct {
}

StartSingleServer("0.0.0.0", "9093", "businessService", &BusinessServiceImpl{})
```
main.go(多个service)
```go

import (
	"code.aliyun.com/wyunshare/thrift-server/business"
	"code.aliyun.com/wyunshare/thrift-server/processor"
	"code.aliyun.com/wyunshare/thrift-server/server"
)

type BusinessServiceImpl struct {
}

func main() {

	const thrift_server_address = "0.0.0.0" // 0.0.0.0 表示监听所有端口
	const thrift_server_port = "8888"

	businessServiceMap := &business.BusinessServiceMap{
		ServiceMap: make(map[string]business.IBusinessService),
	}
	businessServiceMap.RegisterService("businessService", &BusinessServiceImpl{})

	wyunServiceImpl := processor.WyunServiceImpl{
		BusinessServiceMap: businessServiceMap,
	}

	server.StartServer(thrift_server_address, thrift_server_port, &wyunServiceImpl)

}
```
客户端使用
```go

import (
	"code.aliyun.com/wyunshare/thrift-server/gen-go/server"
	"code.aliyun.com/wyunshare/thrift-server"
	"log"
	"encoding/json"
)

func main() {

		po := GetPool(net.JoinHostPort("localhost", "9092"))
		pooledClient, err := po.Get() //从连接池获取thrift client
		if err != nil {
			log.Println("Thrift pool get client error", err)
			return
		}
	
		defer po.Put(pooledClient, false)
	
		rawClient, ok := pooledClient.(*server.MyServiceClient)
		if !ok {
			log.Println("convert to raw client failed")
			return
		}
	
		req := server.NewRequest() //创建request
		req.ServiceName = "businessService"
		req.Operation = "test"
	
		paramInput := make(map[string]interface{})
	
		req.ParamJSON, _ = json.Marshal(paramInput)
	
		r, err := rawClient.Send(req) //发送reque	
```