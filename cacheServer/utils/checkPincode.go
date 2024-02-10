package utils

import (
	"cache-server/types"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	// "net/url"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
)

const (
	DefaultExpiration = 48 * time.Hour
	PurgeInterval     = 48 * time.Hour
)

// JSON Format to recieve in body from database
// {
// {pincode : 221144, merchantList : ["merchant1", "merchant2", "merchant3"]},
// {pincode : 221144, merchantList : ["merchant1", "merchant2", "merchant3"]},
// {pincode : 221144, merchantList : ["merchant1", "merchant2", "merchant3"]},
// {pincode : 221144, merchantList : ["merchant1", "merchant2", "merchant3"]},
// {pincode : 221144, merchantList : ["merchant1", "merchant2", "merchant3"]},
// {pincode : 221144, merchantList : ["merchant1", "merchant2", "merchant3"]},
// {pincode : 221144, merchantList : ["merchant1", "merchant2", "merchant3"]},
// }

func cacheMiss(pincode string, c *cache.Cache) []types.PincodeInfo {
	pincodeInt, _ := strconv.Atoi(pincode)
	for i := pincodeInt - 5; i <= pincodeInt+5; i++ {
		fmt.Println("Cache miss")
		var pincodeInfo types.PincodeInfo
		baseUrl := fmt.Sprintf("http://%s:%s/pincode/%s", types.Server.Host, types.Server.Port, pincode)
		resp, err := http.Get(baseUrl)
		body, err := ioutil.ReadAll(resp.Body)

		err = json.Unmarshal(body, &(types.PincodeInfoList))
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return nil
		}

		// Access each element in the array
		for _, PincodeInfo := range types.PincodeInfoList {
			var cacheInfo types.CacheInfo

			if pincode == PincodeInfo.Pincode {
				pincodeInfo.MerchantList = PincodeInfo.MerchantList
				pincodeInfo.Pincode = PincodeInfo.Pincode
				fmt.Println("asdasd", pincodeInfo)
			}
			cacheInfo.LastModified = time.Now().UTC()
			cacheInfo.MerchantList = PincodeInfo.MerchantList
			c.Set(pincode, cacheInfo, cache.DefaultExpiration)
			fmt.Printf("Pincode: %s, Merchants: %v\n", PincodeInfo.Pincode, PincodeInfo.MerchantList)
		}
		fmt.Println(pincodeInfo)
		types.PincodeInfoList = append(types.PincodeInfoList, pincodeInfo)
	}
	return types.PincodeInfoList
}

func CheckPincode(pincode string, ctx *fiber.Ctx) []byte {
	types.PincodeInfoList = types.PincodeInfoList[:0]
	c := cache.New(48*time.Hour, 48*time.Hour)

	if x, found := c.Get(pincode); found {
		fmt.Println("Cache hit")
		var cacheInfo types.CacheInfo
		cacheInfo = x.(types.CacheInfo)

		// lastModified := cacheInfo.LastModified
		// baseUrl := fmt.Sprintf("http://%s:%s/checkLastModified?pincode=%s&date=%s", types.Server.Host, types.Server.Port, pincode, lastModified.String())

		// resp, _ := http.Get(baseUrl)
		// body, _ := ioutil.ReadAll(resp.Body) // {status: 200}
		// if string(body) != "200" {
		// 	types.PincodeInfoList = cacheMiss(pincode, c)
		// }

		types.PincodeInfoList[0].Pincode = pincode
		types.PincodeInfoList[0].MerchantList = cacheInfo.MerchantList
	} else {
		types.PincodeInfoList = cacheMiss(pincode, c)
	}

	jsonData, _ := json.Marshal(types.PincodeInfoList)
	return jsonData
}
