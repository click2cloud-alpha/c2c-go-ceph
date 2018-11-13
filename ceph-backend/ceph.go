package main

import (
	"fmt"
	"github.com/click2cloud-alpha/c2c-go-ceph/pkg/api"
)

func main() {
	var access_key = "VX6NBR04LD2YAPQXWN9I"
	var secretKey = "QGg0xdIlq2e6HiWAjpJrgQuXhjy20hOA8FxCm7xN"
	var url = "http://192.168.1.173:7480"
	var prefix = "admin"

	api, error := radosAPI.New(url, access_key, secretKey, prefix)
	if error != nil {
		fmt.Println(error)
	}
	if error == nil {
		fmt.Println(api)
	}
	//new_user, err := api.CreateUser(radosAPI.UserConfig{
	//	UID:"user1",
	//	AccessKey:"TestUserAccessKey",
	//	SecretKey:"TestUserSecretKey",
	//	DisplayName:"User 1",
	//})
	//if err!=nil{
	//	fmt.Println(err)
	//}
	//if new_user!= nil{
	//	fmt.Println(new_user)
	//}
	user, err := api.GetUser("user1")
	if err != nil {
		fmt.Println(err)
	}
	//var user_ak = user.Keys[0].AccessKey
	//var user_sk = user.Keys[0].SecretKey
	var uid = user.Keys[0].User
	//fmt.Printf("Access Key: %v Secret Key: %v\n UID: %T\n ",user_ak, user_sk,uid)

	//api1, error1 := radosAPI.New(url, user_ak, user_sk, uid)
	//if error != nil {
	//	fmt.Println(error1)
	//}
	//if error == nil {
	//	fmt.Println(api1)
	//}
	bucketConfig := radosAPI.BucketConfig{
		Bucket: "opensds-test",
		UID:    uid,
	}
	//api, err := radosAPI.New(url, )
	api.LinkBucket(bucketConfig)

}
