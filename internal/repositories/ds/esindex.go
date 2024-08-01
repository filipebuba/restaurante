package ds

import (
	"context"
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func (e *elasticSearch) CreateIndex(index string) error {
	reqExists := esapi.IndicesExistsRequest{
		Index: []string{index},
	}
	resExists, err := reqExists.Do(context.Background(), e.client)
	if err != nil {
		return fmt.Errorf("error checking if the index exists: %w", err)
	}
	defer resExists.Body.Close()

	if resExists.StatusCode == 404 {
		// Index does not exist, create it
		return e.createIndexWithMapping(index)
	}

	e.index = index

	return nil
}

func (e *elasticSearch) createIndexWithMapping(index string) error {
	// 1. Create an index with a mapping
	mapping := `{
		"mappings": {
			"properties": {
				"nome": {
					"type": "text"
				},
				"email": {
					"type": "text"
				},
				"telefone": {
					"type": "text"
				},
				"feedbacks": {
					"type": "text",
					"properties": {
					"email": {
						"type": "keyword"
					},
					"orders": {
						"type": "text"
					}
					}
				}
			}
		}
	}`

	req := esapi.IndicesCreateRequest{
		Index: index,
		Body:  strings.NewReader(mapping),
	}

	res, err := req.Do(context.Background(), e.client)
	if err != nil {
		return fmt.Errorf("error creating the index: %w", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		return fmt.Errorf("error response from Elasticsearch: %s", res.String())
	}

	e.index = index

	fmt.Println("Index created successfully")

	return nil
}
