package ds

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	es8 "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/filipebuba/restaurante/internal/core/domain"
)

type elasticSearch struct {
	client *es8.Client
	index  string
}

func NewEsRepository(client *es8.Client) *elasticSearch {
	return &elasticSearch{
		client: client,
	}
}

func (r *elasticSearch) GetAllClientes(ctx context.Context, limit int, searchAfter []interface{}) ([]domain.Cliente, interface{}, error) {
	query := map[string]interface{}{
		"size": limit,
		"sort": []map[string]string{
			{"created_at": "asc"},
		},
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}

	if searchAfter != nil {
		query["search_after"] = searchAfter
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, nil, fmt.Errorf("error encoding query: %s", err)
	}

	req := esapi.SearchRequest{
		Index: []string{r.index},
		Body:  &buf,
	}

	res, err := req.Do(ctx, r.client)
	if err != nil {
		return nil, nil, fmt.Errorf("search request: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, nil, fmt.Errorf("error response: %s", res.String())
	}

	var result struct {
		Hits struct {
			Hits []struct {
				ID     string         `json:"id"`
				Source domain.Cliente `json:"_source"`
				Sort   []interface{}  `json:"sort"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, nil, fmt.Errorf("error decoding response: %s", err)
	}

	items := make([]domain.Cliente, len(result.Hits.Hits))
	for i, hit := range result.Hits.Hits {
		hit.Source.ID = hit.ID
		items[i] = hit.Source
	}

	if len(result.Hits.Hits) == 0 {
		return items, nil, nil
	}

	return items, result.Hits.Hits[len(result.Hits.Hits)-1].Sort, nil
}

func (r *elasticSearch) CreateCliente(ctx context.Context, client domain.Cliente) (*domain.Cliente, error) {
	client.ID = ""

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(client); err != nil {
		return nil, fmt.Errorf("error encoding document: %s", err)
	}

	req := esapi.IndexRequest{
		Index: r.index,
		Body:  &buf,
	}

	res, err := req.Do(ctx, r.client)
	if err != nil {
		return nil, fmt.Errorf("insert: request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode == 409 {
		return nil, errors.New("conflict")
	}

	if res.IsError() {
		return nil, fmt.Errorf("insert: response: %s", res.String())
	}

	var resBody map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
		return nil, fmt.Errorf("error parsing the response body: %s", err)
	}

	if id, ok := resBody["id"].(string); ok {
		client.ID = id
	} else {
		return nil, fmt.Errorf("error: no _id returned in response")
	}

	return &client, nil
}

func (r *elasticSearch) UpdateCliente(ctx context.Context, editCliente domain.Cliente) (*domain.Cliente, error) {

	bdy, err := json.Marshal(editCliente)
	if err != nil {
		return nil, fmt.Errorf("update: marshall: %w", err)
	}

	req := esapi.UpdateRequest{
		Index:      r.index,
		DocumentID: editCliente.ID,
		Body:       bytes.NewReader([]byte(fmt.Sprintf(`{"doc":%s}`, bdy))),
	}

	res, err := req.Do(ctx, r.client)
	if err != nil {
		return nil, fmt.Errorf("update: request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return nil, fmt.Errorf("not found")
	}

	if res.IsError() {
		return nil, fmt.Errorf("update: response: %s", res.String())
	}

	return &editCliente, nil
}

func (r *elasticSearch) DeleteCliente(ctx context.Context, id string) error {
	req := esapi.DeleteRequest{
		Index:      r.index,
		DocumentID: id,
	}

	res, err := req.Do(ctx, r.client)
	if err != nil {
		return fmt.Errorf("delete: request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return fmt.Errorf("not found")
	}

	if res.IsError() {
		return fmt.Errorf("delete: response: %s", res.String())
	}

	return nil
}

// document represents a single document in Get API response body.
type document struct {
	Source interface{} `json:"_source"`
}
