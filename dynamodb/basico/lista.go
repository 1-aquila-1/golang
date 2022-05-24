package basico

import (
	"context"
	"fmt"

	"github.com/1-aquila-1/golang/dynamodb/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func Lista(db *dynamodb.Client) {
	out, err := db.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String("ESTUDO_TB_USUARIO"),
	})

	if err != nil {
		panic(err)
	}
	usuarios := []models.Usuario{}
	err = attributevalue.UnmarshalListOfMaps(out.Items, &usuarios)
	if err != nil {
		panic(err)
	}

	for _, item := range usuarios {
		fmt.Println(item)
	}

}
