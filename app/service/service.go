package service

import (
	"Grpc-crud/app/pb"
	er "Grpc-crud/app/utils/errors"
	res "Grpc-crud/app/utils/response"
	db "Grpc-crud/app/service/DynamoClient"
) 

var (
	Dy db.DynamoDB
)
func WriteRequest(b pb.Binder)(*res.ResRespo, *er.ResErr) {
	Dy.Save(b)
	return nil, nil
}