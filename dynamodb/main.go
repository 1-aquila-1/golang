package main

import (
	"context"

	"github.com/1-aquila-1/golang/dynamodb/basico"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func awsConfig() (aws.Config, error) {
	return config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = "sa-east-1"
		return nil
	})
}

func awsConfigLocal() (aws.Config, error) {
	return config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{URL: "http://localhost:8000"}, nil
			})),
	)
}

func main() {

	cfg, err := awsConfigLocal()
	if err != nil {
		panic(err)
	}
	svc := dynamodb.NewFromConfig(cfg)

	// tabela.Lista(svc)
	// tabela.Criacao(svc)

	// user := models.Usuario{
	// 	Id:      "u4",
	// 	Usuario: "paulo123",
	// }
	// basico.Insert(svc, user)

	// basico.BuscaId(svc, "u4")

	// user := models.Usuario{
	// 	Id:      "12346",
	// 	Usuario: "joao1",
	// }
	// basico.Update(svc, user)

	// basico.BuscaId(svc, "12346")
	basico.Lista(svc)

}
