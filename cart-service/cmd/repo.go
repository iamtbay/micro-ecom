package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository struct{}

func initRepo() *Repository {
	return &Repository{}
}

// !
// GET CART
func (x *Repository) getCart(userID uuid.UUID) (CartOrder, error) {
	//ctx
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	//find
	pattern := fmt.Sprintf("cart:%s:*", userID.String())
	cartData, err := rdb.Keys(ctx, pattern).Result()
	if err != nil {
		return CartOrder{}, err
	}

	var totalPrice float64
	var products []CartItem
	for _, key := range cartData {
		keyIDs := strings.Split(key, ":")
		//
		var product CartItem
		productRd, err := rdb.HGetAll(ctx, key).Result()
		fmt.Println(productRd)
		if err != nil {
			return CartOrder{}, err
		}

		product.Name = productRd["name"]
		product.ProductID, _ = primitive.ObjectIDFromHex(keyIDs[len(keyIDs)-1])
		product.Quantity, _ = strconv.Atoi(productRd["quantity"])
		product.Price, _ = strconv.ParseFloat(productRd["price"], 64)
		totalPrice += product.Price * float64(product.Quantity)
		products = append(products, product)
	}
	return CartOrder{
		UserID:     userID,
		Products:   products,
		TotalPrice: totalPrice,
	}, nil
}

// !
// ADD TO CART
func (x *Repository) addToCart(userID uuid.UUID, product CartItem) error {
	//ctx
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	//todo
	//check product already user's list or not?
	exists, err := rdb.Exists(ctx, fmt.Sprintf("cart:%s:%s", userID.String(), product.ProductID.Hex()), "").Result()
	if err != nil {
		return err
	}

	//check product is already in cart or not
	if exists == 0 {
		err = rdb.HSet(ctx,
			fmt.Sprintf("cart:%s:%s", userID.String(), product.ProductID.Hex()),
			"name", product.Name,
			"quantity", product.Quantity,
			"price", product.Price,
		).Err()
		if err != nil {
			return err
		}
	} else {
		//todo
		// can we change the amount dynamically? especially when coding front-end
		// maybe create router for increase and decraese
		err = rdb.HIncrBy(ctx, fmt.Sprintf("cart:%s:%s", userID.String(), product.ProductID.Hex()), "quantity", 1).Err()

		if err != nil {
			return err
		}
	}

	//save

	return nil
}

// !
// UPDATE QUANTITY OF PRODUCT
func (x *Repository) updateQuantityOfProduct(userID uuid.UUID, productID string, quantity int, isExact bool) (string, error) {
	//ctx
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	//
	key := fmt.Sprintf("cart:%s:%s", userID, productID)
	quantityStr, err := rdb.HGet(ctx, key, "quantity").Result()
	if err != nil {
		return "", errors.New("product isn't in your cart")
	}

	productQuantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		return "", err
	}
	if isExact {
		//todo
		// fix here isnt arrange the quantity
		err := rdb.HSet(ctx, key, "quantity", quantity).Err()
		if err != nil {
			return "", err
		}
		return "quantity arranged", nil
	}

	if quantity == -1 && productQuantity <= 1 {
		//del product from cart
		err := rdb.Del(ctx, key).Err()
		if err != nil {
			return "", err
		}
		return "item removed from your cart", nil
	} else if quantity == -1 {
		err := rdb.HIncrBy(ctx, key, "quantity", -1).Err()
		if err != nil {
			return "", err
		}
		return "quantity decreased", nil
	} else if quantity == 1 {
		err := rdb.HIncrBy(ctx, key, "quantity", 1).Err()
		if err != nil {
			return "", err
		}
		return "quantity increased", nil
	} else if quantity >= 1 && !isExact {
		err := rdb.HIncrBy(ctx, key, "quantity", int64(quantity)).Err()
		if err != nil {
			return "", err
		}
		return "quantity increased", nil
	}
	return "", nil
}

//!
//UPDATE PRODUCT INFO ON CART

func (x *Repository) updateProduct(product UpdateProductType) error {
	//ctx
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	//
	var cursor uint64
	for {
		keys, nextCursor, err := rdb.Scan(ctx, cursor, fmt.Sprintf("cart:*:%v", product.ID.Hex()), 0).Result()
		if err != nil {
			return err
		}

		for _, key := range keys {
			fmt.Println(key)
			err := rdb.HSet(ctx, key, "name", product.Name, "price", product.Price).Err()
			if err != nil {
				return err
			}
		}
		cursor = nextCursor

		if cursor == 0 {
			break
		}
	}

	return nil
}

// !
// DELETE PRODUCT ON CART
func (x *Repository) deleteProductOnCart(userID uuid.UUID, productID string) error {
	//ctx
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	key := fmt.Sprintf("cart:%s:%s", userID, productID)
	data, err := rdb.Keys(ctx, key).Result()
	if err != nil {
		return err
	}
	if len(data) > 0 {
		_, err := rdb.HDel(ctx, key, "name", "quantity", "price").Result()
		if err != nil {
			return err
		}
	} else {
		return errors.New("product couldn't find in your cart")
	}
	return nil
}

// !
// DELETE USER CART
func (x *Repository) deleteUserCart(userID uuid.UUID) error {
	//ctx
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	key := fmt.Sprintf("cart:%s:*", userID.String())

	var cursor uint64
	for {
		keys, newCursor, err := rdb.Scan(ctx, cursor, key, 0).Result()
		if err != nil {
			return err
		}

		if len(keys) > 0 {
			_, err := rdb.Del(ctx, keys...).Result()
			if err != nil {
				return err
			}
			fmt.Printf("deleted keys %v \n", keys)
		}
		cursor = newCursor
		if cursor == 0 {
			break
		}
	}
	return nil
}
