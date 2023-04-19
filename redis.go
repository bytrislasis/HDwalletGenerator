package main

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-redis/redis/v8"
)

func SetAddress(rdb *redis.Client, address common.Address) error {
	addressStr := address.Hex()
	err := rdb.Set(context.Background(), addressStr, "0", 0).Err()
	return err
}

func GetAddressValue(rdb *redis.Client, address common.Address) (string, error) {
	addressStr := address.Hex()
	value, err := rdb.Get(context.Background(), addressStr).Result()
	return value, err
}

func CheckKeyExists(rdb *redis.Client, address common.Address) (bool, error) {
	addressStr := address.Hex()
	keyExists, err := rdb.Exists(context.Background(), addressStr).Result()
	return keyExists == 1, err
}
