package data

import (
	grpcclt "miopkg/grpc/client"
	commentv1 "redbook/api/domain/comment/v1"
	userv1 "redbook/api/domain/user/v1"
)

func NewCommentRpcClient() commentv1.CommentRpcClient {
	conn := grpcclt.StdConfig("directserver.comment").Build()
	return commentv1.NewCommentRpcClient(conn)
}
func NewUserRpcClient() userv1.UserRpcClient {
	conn := grpcclt.StdConfig("directserver.user").Build()
	return userv1.NewUserRpcClient(conn)
}
