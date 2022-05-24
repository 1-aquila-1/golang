package basico

import (
	"context"
	"fmt"

	"github.com/1-aquila-1/golang/dynamodb/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func Insert(svc *dynamodb.Client, user models.Usuario) {
	out, err := svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("ESTUDO_TB_USUARIO"),
		Item: map[string]types.AttributeValue{
			"ID":      &types.AttributeValueMemberS{Value: user.Id},
			"USUARIO": &types.AttributeValueMemberS{Value: user.Usuario},
			"EMAIL":   &types.AttributeValueMemberS{Value: user.Email},
		},
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(out.Attributes)
}
