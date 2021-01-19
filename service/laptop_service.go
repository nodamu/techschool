package service

import (
	"errors"
	"fmt"
	"sync"

	"github.com/nodamu/techschool/pb"

	"github.com/jinzhu/copier"
)

//ErrorAlreadyExists return error when record alread exists
var ErrorAlreadyExists = errors.New("record already exists")

// LaptopStore is an interface to store laptop
type LaptopStore interface {
	// Save saves the laptop to the store
	Save(laptop *pb.Laptop) error

	// Find laptop by Id
	Find(id string) (*pb.Laptop, error)
}

// InMemoryLaptopStore stores laptop in memory
type InMemoryLaptopStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Laptop
}

//NewInMemoryLaptopStore returns a new InMemoryLaptopStore
func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

// Save saves the laptop to the store
func (store *InMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[laptop.Id] != nil {
		return ErrorAlreadyExists
	}

	//Deep copy
	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return fmt.Errorf("cannot copy laptop data: %w", err)
	}
	store.data[other.Id] = other
	return nil
}

// Find laptop by Id
func (store *InMemoryLaptopStore) Find(id string) (*pb.Laptop, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	laptop := store.data[id]

	if laptop == nil {
		return nil, nil
	}

	//deep copy
	other := &pb.Laptop{}

	err := copier.Copy(other, laptop)

	if err != nil {
		return nil, fmt.Errorf("Cannot copy laptop data: %w", err)
	}

	return other, nil

}
