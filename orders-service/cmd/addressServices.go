package main

import "github.com/google/uuid"

// !
// GET ADDRESSES
func (x *Services) getAddresses(userID uuid.UUID) ([]*GetAddresses, error) {
	addresses, err := repo.getAddresses(userID)
	if err != nil {
		return []*GetAddresses{}, err
	}
	return addresses, nil
}

// !
// GET SINGLE ADDRESS
func (x *Services) getSingleAddress(userID uuid.UUID, addressIDStr string) (*GetAddresses, error) {
	//parse address id
	addressID, err := uuid.Parse(addressIDStr)
	if err != nil {
		return &GetAddresses{}, nil
	}
	//repo
	addresses, err := repo.getSingleAddressByID(userID, addressID)
	if err != nil {
		return &GetAddresses{}, err
	}

	//
	return addresses, nil
}

// !
// GET ADDRESSES
func (x *Services) addNewAddress(address NewAddress) error {
	err := repo.addNewAddress(address)
	if err != nil {
		return err
	}
	return nil
}

// !
// EDIT ADDRESS BY ID
func (x *Services) editAddressByID(addressInfo NewAddress, addressIDStr string) error {
	//parse address id
	addressID, err := uuid.Parse(addressIDStr)
	if err != nil {
		return err
	}
	//repo
	err = repo.editAddressByID(addressInfo, addressID)
	if err != nil {
		return err
	}

	//
	return nil
}

// !
// DELETE ADDRESS BY ID
func (x *Services) deleteAddressByID(addressIDStr string, userID uuid.UUID) error {
	//parse address id
	addressID, err := uuid.Parse(addressIDStr)
	if err != nil {
		return err
	}
	//repo
	err = repo.deleteAddressByID(userID, addressID)
	if err != nil {
		return err
	}

	//
	return nil
}
