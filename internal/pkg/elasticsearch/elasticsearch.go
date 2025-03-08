package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/muharik19/boiler-plate-grpc/pkg/logger"
	global "github.com/muharik19/boiler-plate-grpc/pkg/utils"
)

var (
	appName = *global.Getenv("APP_NAME")
	Client  *elasticsearch.Client
)

// Init initializes Elasticsearch connection
func Init() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			*global.Getenv("ELASTIC_URL_1"),
		},
		Transport: &http.Transport{
			// MaxIdleConnsPerHost:   10,
			// ResponseHeaderTimeout: time.Millisecond,
			// DialContext:           (&net.Dialer{Timeout: time.Nanosecond}).DialContext,
			// TLSClientConfig: &tls.Config{
			// MinVersion: tls.VersionTLS12,
			// ...
			// },
		},
	}

	var err error
	Client, err = elasticsearch.NewClient(cfg)
	if err != nil {
		logger.Errf("Error creating Elasticsearch client: %s", err.Error())
	}

	logger.Info("Elasticsearch successfully connected")
}

// Insert adds a document to Elasticsearch
func Insert(ctx context.Context, index string, doc any) error {
	jsonBody, err := json.Marshal(doc)
	if err != nil {
		return fmt.Errorf("failed to encode document: %w", err)
	}

	req := esapi.IndexRequest{
		Index:   fmt.Sprintf("%s.%s", appName, index),
		Body:    bytes.NewReader(jsonBody),
		Refresh: "true",
	}

	res, err := req.Do(ctx, Client)
	if err != nil {
		logger.Errf("Failed to insert document %v: Error: %s", req, err.Error())
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("insert failed, status: %s", res.Status())
	}

	logger.Info("Elasticsearch document inserted successfully")

	return nil
}

// Update modifies an existing document
func Update(ctx context.Context, index, ID string, update map[string]any) error {
	jsonBody, err := json.Marshal(map[string]any{"doc": update})
	if err != nil {
		return fmt.Errorf("failed to encode update document: %w", err)
	}

	req := esapi.UpdateRequest{
		Index:      index,
		DocumentID: ID,
		Body:       bytes.NewReader(jsonBody),
		Refresh:    "true",
	}

	res, err := req.Do(ctx, Client)
	if err != nil {
		logger.Errf("Failed to update document %v: Error: %s", req, err.Error())
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("update failed, status: %s", res.Status())
	}

	logger.Info("Elasticsearch document updated successfully")

	return nil
}

// Search queries Elasticsearch
func Search(ctx context.Context, index string, query map[string]any) (map[string]any, error) {
	jsonBody, err := json.Marshal(query)
	if err != nil {
		return nil, fmt.Errorf("failed to encode search query: %w", err)
	}

	req := esapi.SearchRequest{
		Index: []string{index},
		Body:  bytes.NewReader(jsonBody),
	}

	res, err := req.Do(ctx, Client)
	if err != nil {
		logger.Errf("Failed to execute search %v: Error: %s", req, err.Error())
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("search failed, status: %s", res.Status())
	}

	var result map[string]any
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode search response: %w", err)
	}

	logger.Info("Elasticsearch search executed successfully")

	return result, nil
}
