package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jmoiron/jsonq"
)

func main() {
	j := `{
  "result": {
    "token": "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJuYW1laWQiOiIwIiwidW5pcXVlX25hbWUiOiJzaG9wLTcwMDAwNTMzODciLCJodHRwOi8vc2NoZW1hcy5taWNyb3NvZnQuY29tL2FjY2Vzc2NvbnRyb2xzZXJ2aWNlLzIwMTAvMDcvY2xhaW1zL2lkZW50aXR5cHJvdmlkZXIiOiJBU1AuTkVUIElkZW50aXR5IiwiaHR0cDovL3d3dy5hc3BuZXRib2lsZXJwbGF0ZS5jb20vaWRlbnRpdHkvY2xhaW1zL3RlbmFudElkIjoiMCIsInJvbGUiOiJTcG90IiwiaXNzIjoiRUUiLCJleHAiOjE0OTA3NTE1NzUsIm5iZiI6MTQ5MDQ5MjM3NX0.xYfr5n_1CLFtgUF6dkRmD7raK1Ff0T90focHD3xvuTo",
    "userId": 0,
    "userName": "赵金红",
    "isDeveloper": false,
    "roles": [
      {
        "name": "Spot",
        "isOwner": false
      }
    ],
    "tenantId": 0,
    "tenantCode": "EE",
    "tenantName": "EE",
    "appId": null,
    "spots": [
      {
        "id": 0,
        "code": "",
        "displayName": "",
        "info": {
          "cslInfo": [
            {
              "shopCode": "C30R",
              "brandCode": "EE",
              "tenantId": 0
            }
          ]
        }
      }
    ],
    "spotsGroup": null,
    "brands": [
      {
        "brandCode": "EE",
        "isChief": false,
        "chiefBrandCode": null
      }
    ],
    "userInformation": {
      "cslInfo": {
        "empId": "7000053387"
      }
    }
  },
  "targetUrl": null,
  "success": true,
  "error": null,
  "unAuthorizedRequest": false,
  "__abp": true
}`

	data := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(j))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)
	// data["foo"] -> 1
	// fmt.Println(jq.Int("foo"))

	// // data["subobj"]["subarray"][1] -> 2
	// fmt.Println(jq.Int("subobj", "subarray", "1"))

	// // data["subobj"]["subarray"]["array"][0] -> "hello"
	// fmt.Println(jq.String("subobj", "subsubobj", "array", "0"))

	// // data["subobj"] -> map[string]interface{}{"subobj": ...}
	// fmt.Println(jq.Object("subobj"))
	fmt.Println(jq.ArrayOfObjects("result", "spots", "0", "info", "cslInfo"))
}
