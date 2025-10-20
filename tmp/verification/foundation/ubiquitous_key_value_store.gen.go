// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var ubiquitousKeyValueStoreClass _UbiquitousKeyValueStoreClass

func init() {
	ubiquitousKeyValueStoreClass = _UbiquitousKeyValueStoreClass{objc.GetClass("NSUbiquitousKeyValueStore")}
}

type _UbiquitousKeyValueStoreClass struct {
	class objc.Class
}

type UbiquitousKeyValueStore struct {
	objc.ID
}

func UbiquitousKeyValueStoreFrom(ptr unsafe.Pointer) UbiquitousKeyValueStore {
	return UbiquitousKeyValueStore{
		ID: objc.ID(ptr),
	}
}




