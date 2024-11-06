package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Handlers struct{}

func initHandlers() *Handlers {
	return &Handlers{}
}

var services = initServices()

func (x *Handlers) SearchProduct(c *gin.Context) {
	name := c.Query("name")
	brand := c.Query("brand")
	minPrice := c.Query("minPrice")
	maxPrice := c.Query("maxPrice")
	added_by := c.Query("added_by")
	filters := make(map[string]string)

	if name != "" {
		filters["name"] = name
	}
	if brand != "" {
		filters["brand"] = brand
	}
	if minPrice != "" {
		filters["minPrice]"] = minPrice
	}
	if maxPrice != "" {
		filters["maxPrice]"] = maxPrice
	}
	if added_by != "" {
		filters["added_by"] = added_by
	}

	minPriceInt, _ := strconv.Atoi(minPrice)
	maxPriceInt, _ := strconv.Atoi(maxPrice)

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{
						"bool": map[string]interface{}{
							"should": []map[string]interface{}{
								{
									"prefix": map[string]interface{}{
										"product_name": filters["name"],
									},
								},
								{
									"prefix": map[string]interface{}{
										"brand": filters["brand"],
									},
								},
							},
						},
					},
					{
						"range": map[string]interface{}{
							"price": map[string]interface{}{
								"gte": minPriceInt,
								"lte": maxPriceInt,
							},
						},
					},
				},
			},
		},
	}
	queryBytes, err := json.Marshal(query)
	if err != nil {
		log.Fatalf("Error marshaling query: %s", err)
	}

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("products"),
		es.Search.WithBody(strings.NewReader(string(queryBytes))),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	defer res.Body.Close()

	var resBody map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Missing search query"})
		return
	}

	c.JSON(200, gin.H{"message": "Product found", "data": resBody})
}
