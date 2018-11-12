package main

import (
	"c2c-go-ceph/pkg/api"
	"fmt"
)

func main() {
	var access_key = "VX6NBR04LD2YAPQXWN9I"
	var secretKey = "QGg0xdIlq2e6HiWAjpJrgQuXhjy20hOA8FxCm7xN"
	var url = "http://192.168.1.173:8443"
	var prefix = "admin"

	api, error := radosAPI.New(url, access_key, secretKey, prefix)
	if error != nil {
		fmt.Println(error)
	}
	if error == nil {
		fmt.Println(api)
	}
	//user, err := api.CreateUser(radosAPI.UserConfig{
	//	UID:         "user1",
	//	DisplayName: "User1",
	//})
	//
	//fmt.Println(err)
	//fmt.Println(user)
	fmt.Println("API____", api)

	bucket, error := api.GetBucket(radosAPI.BucketConfig{})
	fmt.Println("bucket---", bucket)
	fmt.Println("Error-------", error)

}
