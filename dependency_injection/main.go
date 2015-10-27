package main

import (
	"github.com/go-modules/modules"

	"fmt"
)

// The KVClient interface models a simple key/value store.
type KVClient interface {
	Get(key string) string
	Put(key, value string)
}

// A MapDBClient is a simple mock KVClient implementation backed by a map and configured with a default value for missing keys.
type MapDBClient struct {
	defaultValue string
	db           map[string]string
}

func (client *MapDBClient) Get(key string) string {
	if value, ok := client.db[key]; ok {
		return value
	} else {
		return client.defaultValue
	}
}

func (client *MapDBClient) Put(key, value string) {
	client.db[key] = value
}

// A service module has a 'GetData' service which utilizes an injected DBClient.
type ServiceModule struct {
	KVClient KVClient `inject:""`
}

func (service *ServiceModule) GetData(key string) string {
	return service.KVClient.Get(key)
}

func (service *ServiceModule) StoreData(key, value string) {
	service.KVClient.Put(key, value)
}

type defaultValue string

// This data module provides a KVClient.
type DataModule struct {
	DefaultValue defaultValue
	KVClient     KVClient `provide:""`
}

func (data *DataModule) Provide() error {
	data.KVClient = &MapDBClient{defaultValue: string(data.DefaultValue), db: make(map[string]string)}
	return nil
}

func main() {
	serviceModule := &ServiceModule{}

	dataModule := &DataModule{DefaultValue: "default"}

	binder := modules.NewBinder()
	if err := binder.Bind(serviceModule, dataModule); err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", serviceModule.KVClient)

	fmt.Println(serviceModule.GetData("key"))

	serviceModule.StoreData("key", "value")
	fmt.Println(serviceModule.GetData("key"))

}
