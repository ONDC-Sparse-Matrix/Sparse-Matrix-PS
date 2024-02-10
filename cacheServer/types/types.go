package types

import "time"

type ServerInfo struct {
	Host string
	Port string
}

type CacheInfo struct {
	LastModified time.Time
	MerchantList []string
}

type PincodeInfo struct {
	Pincode      string   `json:"pincode"`
	MerchantList []string `json:"merchantList"`
}

var PincodeInfoList []PincodeInfo

// var data struct {
// 	InfoList []PincodeInfo `json:"PincodeInfo"`
// }

var Server ServerInfo
