package config

import (
	"unsafe"

	mgo "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/fluent/fluent-bit-go/output"
)

func GetAddress(ctx unsafe.Pointer) string {
	return output.FLBPluginConfigKey(ctx, "host_port")
}

func GetCollection(ctx unsafe.Pointer) string {
	return output.FLBPluginConfigKey(ctx, "collection")
}

func GetUsername(ctx unsafe.Pointer) string {
	return output.FLBPluginConfigKey(ctx, "username")
}

func GetPassword(ctx unsafe.Pointer) string {
	return output.FLBPluginConfigKey(ctx, "password")
}

func GetSource(ctx unsafe.Pointer) string {
	return output.FLBPluginConfigKey(ctx, "auth_database")
}

func GetDatabase(ctx unsafe.Pointer) string {
	return output.FLBPluginConfigKey(ctx, "database")
}

func GetConfig(ctx unsafe.Pointer) *mgo.DialInfo {
	return &mgo.DialInfo{
		Addrs:    []string{GetAddress(ctx)},
		Username: GetUsername(ctx),
		Password: GetPassword(ctx),
		Source:   GetSource(ctx),
		Database: GetDatabase(ctx),
	}
}
