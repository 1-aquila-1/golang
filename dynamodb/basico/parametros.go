package basico

import (
	"context"
	"fmt"

	"github.com/1-aquila-1/golang/dynamodb/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func BuscaId(svc *dynamodb.Client, userId string) {

	out, err := svc.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("ESTUDO_TB_USUARIO"),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: userId},
		},
	})

	if err != nil {
		panic(err)
	}

	usuario := models.Usuario{}

	err = attributevalue.UnmarshalMap(out.Item, &usuario)

	if err != nil {
		panic(err)
	}

	fmt.Println("ID :=: ", usuario.Id, "USUARIO :=: ", usuario.Usuario)
}
