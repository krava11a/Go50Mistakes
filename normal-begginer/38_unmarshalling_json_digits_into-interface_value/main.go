package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {

	//var data = []byte(`{"status":200}`)

	// 0-1
	//var result map[string]interface{}
	//if err := json.Unmarshal(data, &result); err != nil {
	//	fmt.Println("error: ",err)
	//	return
	//}
	//
	//// 0 - MISTAKE
	////var status = result["status"].(int)
	//
	//// 1 - CAST
	//var status = uint64(result["status"].(float64))

	// 2 - DECODER, It's good

	//var result map[string]interface{}
	//
	//var decoder = json.NewDecoder(bytes.NewReader(data))
	//decoder.UseNumber()
	//
	//
	//if err := decoder.Decode(&result);err!=nil{
	//	fmt.Println("error: ",err)
	//	return
	//}
	//
	//var status,_ = result["status"].(json.Number).Int64()

	// 3 - Unmarshal
	//var result map[string]interface{}
	//var decoder = json.NewDecoder(bytes.NewReader(data))
	//decoder.UseNumber()
	//
	//if err := decoder.Decode(&result); err != nil {
	//	fmt.Println("error:", err)
	//	return
	//}
	//
	//var status uint64
	//if err := json.Unmarshal([]byte(result["status"].(json.Number).String()), &status); err != nil {
	//	fmt.Println("error:", err)
	//	return
	//}

	// 4 - STRUCT
	//var result struct {
	//	Status uint64 `json:"status"`
	//}
	//
	//if err := json.NewDecoder(bytes.NewReader(data)).Decode(&result); err != nil {
	//	fmt.Println("error:", err)
	//	return
	//}
	//fmt.Printf("result => %+v",result)

	// 5 - deferred processing
	records := [][]byte{
		[]byte(`{"status":200, "tag":"one"}`),
		[]byte(`{"status":"ok", "tag":"two"}`),
	}

	for idx, record := range records {
		var result struct {
			StatusCode uint64
			StatusName string
			Status     json.RawMessage `json:"status"`
			Tag        string          `json:"tag"`
		}

		if err := json.NewDecoder(bytes.NewReader(record)).Decode(&result); err != nil {
			fmt.Println("error:", err)
			return
		}
		var sstatus string
			if err := json.Unmarshal(result.Status, &sstatus); err == nil {
			result.StatusName = sstatus
		}

		var nstatus uint64
		if err := json.Unmarshal(result.Status, &nstatus); err == nil {
			result.StatusCode = nstatus
		}

		fmt.Printf("[%v] result => %+v\n",idx,result)

	}

	//fmt.Println(status)

}
