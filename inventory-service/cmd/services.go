package main

type Services struct{}

func initServices() *Services {
	return &Services{}
}

var repo = initRepository()

// !
// NEW PRODUCT STOCK
func (x *Services) newProductStock(product Product) error {
	err := repo.newProductStock(product)
	if err != nil {
		return err
	}

	return nil

}

// !
// GET PRODUCT STOCK BY ID
func (x *Services) getStock(productID string) (*Product, error) {
	data, err := repo.getStock(productID)
	if err != nil {
		return &Product{}, err
	}
	return data, nil
}

// !
// GET PRODUCT STOCK BY ID
func (x *Services) productReStock(product Product) error {

	err := repo.productReStock(product)
	if err != nil {
		return err
	}
	return nil
}

// !
// CANCEL RESERVATION
func (x *Services) cancelReservation(product ProductData) error {
	err := repo.cancelReservation(product)
	if err != nil {
		return err
	}
	return nil
}

// !
// UPDATE STOCK VIA RESERVED
func (x *Services) updateStockViaReserved(product ProductData) error {
	err := repo.updateStockViaReserved(product)
	if err != nil {
		return err
	}
	return nil
}

// !
// UPDATE STOCK VIA RESERVED
func (x *Services) updateStockViaSold(product ProductData) error {
	err := repo.updateStockViaSold(product)
	if err != nil {
		return err
	}
	return nil
}
