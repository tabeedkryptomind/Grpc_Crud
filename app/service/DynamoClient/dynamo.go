package dynamo_client

import (
	"Grpc-crud/app/pb"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type DynamoDB struct {
	tableName string
}
func DynamoDbTable() DynamoDB {
	return DynamoDB{
		tableName: "blocks_table",
	}
}
func CreatDynamoClient() *dynamodb.DynamoDB {
	session := session.Must(
		session.NewSessionWithOptions(
			session.Options{
				SharedConfigState: session.SharedConfigEnable,
			}))
	return dynamodb.New(session)
}

func (db * DynamoDB)Save(blocks pb.Binder){
	dynamoDbClient :=  CreatDynamoClient()

	attributeValue, err := dynamodbattribute.MarshalMap(blocks)
	fmt.Println(attributeValue)
	if err != nil {
		log.Fatal(err)
	}
	item :=  &dynamodb.PutItemInput{
		Item: attributeValue,
		TableName: aws.String("blocks_table"),
	}
	_, err = dynamoDbClient.PutItem(item)
	if err != nil {
		log.Fatal(err)
	}
}