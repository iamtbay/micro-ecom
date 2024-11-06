package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Services struct{}

func initServices() *Services {
	return &Services{}
}

func (x *Services) IndexProduct(product Product) error {

	productIndex := ProductIndex{
		Name:    product.Name,
		Brand:   product.Brand,
		Content: product.Content,
		Price:   product.Price,
		AddedBy: product.AddedBy,
	}
	data, err := json.Marshal(productIndex)
	if err != nil {
		return err
	}

	res, err := es.Index(
		"products",
		strings.NewReader(string(data)),
		es.Index.WithDocumentID(product.ID.Hex()),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("failed to index product: %v", res.String())
	}

	fmt.Printf("Succesfully indexed product %s \n", product.ID)
	return nil
}
