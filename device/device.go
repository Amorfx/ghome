package device

import (
	"clementdecou/ghome/bosesoundtouch"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Device struct {
	Type        string
	Name        string
	IP          string
	isConnected bool
}

type DeviceManager struct {
	Devices []Device
	Types   map[string]string
}

func InitDeviceManager() DeviceManager {
	return DeviceManager{Types: map[string]string{
		bosesoundtouch.GetType(): bosesoundtouch.GetName(),
	}}
}

func GetAllDevices(db *mongo.Client) []Device {
	// Find All
	cursor, err := db.Database("ghome").Collection("devices").Find(context.TODO(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	var results []Device

	if err = cursor.All(context.Background(), &results); err != nil {
		log.Fatal(err)
	}

	return results
}

func AddDevice(db *mongo.Client, device_type string, device_name string, device_ip string) {
	db.Database("ghome").Collection("devices").InsertOne(context.TODO(), bson.M{"type": device_type, "ip": device_ip, "name": device_name})
}
