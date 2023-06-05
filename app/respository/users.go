package respository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
	"web-diet-bot/model"
)

type API interface {
	UserIsActive(user string) (bool, error)
	FindUser(user string) (*model.User, error)
	UpdateUserPhoto(user string, base64photo string) error
}

type Repository struct {
	Db *dynamodb.DynamoDB
}

func (r Repository) UserIsActive(user string) (bool, error) {
	findUser, err := r.FindUser(user)
	if err != nil {
		return false, err
	}
	return findUser != nil, nil
}

func (r Repository) FindUser(user string) (*model.User, error) {
	log.Println("Finding user: ", user)
	input := &dynamodb.GetItemInput{
		TableName: aws.String("Users"),
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String(user),
			},
		},
	}
	result, err := r.Db.GetItem(input)
	if err != nil {
		log.Println("Query fail")
		return nil, err
	}

	u := &model.User{}
	dynamodbattribute.UnmarshalMap(result.Item, u)
	return u, nil
}

func (r Repository) UpdateUserPhoto(user string, base64photo string) error {
	putItemInput := dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":r": {
				S: aws.String(base64photo),
			},
		},
		TableName: aws.String("Users"),
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String(user),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set LastPhoto = :r"),
	}
	_, err := r.Db.UpdateItem(&putItemInput)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("user photo saved")
	return nil
}
