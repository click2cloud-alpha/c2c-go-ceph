package main

import (
	"fmt"
	"github.com/QuentinPerez/go-radosgw/pkg/api"
)

func main() {
	var access_key = "VX6NBR04LD2YAPQXWN9I"
	var secretKey = "QGg0xdIlq2e6HiWAjpJrgQuXhjy20hOA8FxCm7xN"
	var url = "http://192.168.1.173:7480"
	var prefix = "admin"

	api, error := radosAPI.New(url, access_key, secretKey, prefix)
	if(error!=nil){
		fmt.Println(error)
	}
	if(error == nil) {
		fmt.Println(api)
	}
	//user, err := api.CreateUser(radosAPI.UserConfig{
	//	UID:         "Aditya",
	//	DisplayName: "Aditya Bammrotwar",
	//})
	//
	//fmt.Println(err)
	//fmt.Println(user)

	bucket, error := api.GetBucket(radosAPI.BucketConfig{

	})
	fmt.Println(bucket)
	fmt.Println(error)
}
