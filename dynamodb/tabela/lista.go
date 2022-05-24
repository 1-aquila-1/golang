package tabela

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func Lista(svc *dynamodb.Client) {
	p := dynamodb.NewListTablesPaginator(svc, nil, func(o *dynamodb.ListTablesPaginatorOptions) {
		o.StopOnDuplicateToken = true
	})

	for p.HasMorePages() {
		out, err := p.NextPage(context.TODO())
		if err != nil {
			panic(err)
		}

		for _, tn := range out.TableNames {
			fmt.Println(tn)
		}
	}
}
