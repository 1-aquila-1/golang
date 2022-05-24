package basico

import (
	"context"
	"fmt"

	"github.com/1-aquila-1/golang/dynamodb/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func Update(svc *dynamodb.Client, user models.Usuario) {
	out, err := svc.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String("ESTUDO_TB_USUARIO"),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: user.Id},
		},
		UpdateExpression: aws.String("set USUARIO = :USUARIO"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":USUARIO": &types.AttributeValueMemberS{Value: user.Usuario},
		},
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(out.Attributes)
}
