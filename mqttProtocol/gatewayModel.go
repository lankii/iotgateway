package mqttProtocol

// 网关批量上报结构体
type (
	//设备上报请求
	GatewayBatchReq struct {
		Id      string       `json:"id"`
		Version string       `json:"version"`
		Sys     SysInfo      `json:"sys"`
		Params  PropertyInfo `json:"params"`
		Method  string       `json:"method"`
	}

	PropertyInfo struct {
		Properties map[string]interface{} `json:"properties"`
		Events     map[string]EventNode   `json:"events"`
		SubDevices []Sub                  `json:"subDevices"`
	}
	Sub struct {
		Identity   Identity               `json:"identity"`
		Properties map[string]interface{} `json:"properties"`
		Events     map[string]EventNode   `json:"events"`
	}
	Identity struct {
		ProductKey string `json:"productKey"`
		DeviceKey  string `json:"deviceKey"`
	}
	EventNode struct {
		Value      map[string]interface{} `json:"value"`
		CreateTime int64                  `json:"time"`
	}

	// GatewayBatchReply 设备上报回复
	GatewayBatchReply struct {
		Code int `json:"code"`
		Data struct {
		} `json:"data"`
		Id      string `json:"id"`
		Message string `json:"message"`
		Method  string `json:"method"`
		Version string `json:"version"`
	}
)
