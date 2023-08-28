package controllers

import (
	"context"
	"strings"

	"go/protos"

	"go-grpc-practice/libs/tool"
	"go-grpc-practice/models"
)

type Server struct {
	protos.UnimplementedTestRpcServiceServer
}

func (s *Server) AddUserData(ctx context.Context, req *protos.AddUserDataRequest) (*protos.AddUserDataResponse, error) {
	result := &protos.AddUserDataResponse{Status: 0}

	if strings.TrimSpace(req.Param) == "" {
		return result, nil
	}

	// to do something
	models.AddUserForMySQL(req.Param)

	result.Status = 1
	result.Uid = "100008868"

	return result, nil
}

var userRestrict = []string{"uid", "other"}

func (s *Server) GetUserData(ctx context.Context, req *protos.GetUserDataRequest) (*protos.GetUserDataResponse, error) {
	result := &protos.GetUserDataResponse{Status: 0}

	if strings.TrimSpace(req.Uid) == "" {
		return result, nil
	}

	query := []string{}
	param := []string{}

	if strings.TrimSpace(req.Param) != "" {
		query = strings.Split(req.Param, ",")
	}

	// 清除不在列表的字段
	for _, v := range query {
		if tool.Contain(userRestrict, v) == true {
			param = append(param, v)
		}
	}

	if len(param) > 0 {
		// to do something
		res := models.GetUserForMySQL(req.Uid, param)
		data := string(tool.MarshalJson(res))

		result.Status = 1
		result.Data = data
	}

	return result, nil
}
