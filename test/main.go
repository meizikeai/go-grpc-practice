package main

import (
	"fmt"

	"go-grpc-practice/test/unit"
)

func main() {
	// testAddUserData()
	// testGetUserData()
}

func testAddUserData() {
	data := "name=love,sex=1"
	res, err := unit.Client.AddUserData(data)

	if err != nil {
		fmt.Println("test AddUserData", err)
		return
	}

	state := int(res.GetStatus().Number())

	fmt.Println("test AddUserData", state, res.GetStatus().String(), res.Uid)
}

func testGetUserData() {
	data := "user_id,area,phone,state,is_robot,robot_id"
	res, err := unit.Client.GetUserData("1674295877562273792", data)

	if err != nil {
		fmt.Println("test GetUserData", err)
		return
	}

	state := int(res.GetStatus().Number())

	fmt.Println("test GetUserData", state, res.GetStatus().String(), res.Data)
}
