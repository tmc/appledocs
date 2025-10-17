// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UbiquitousKeyValueStore] class.
var UbiquitousKeyValueStoreClass objc.Class

func init() {
	UbiquitousKeyValueStoreClass = objc.GetClass("NSUbiquitousKeyValueStore")
}

type UbiquitousKeyValueStore struct {
	objc.ID
}

func UbiquitousKeyValueStoreFrom(ptr unsafe.Pointer) UbiquitousKeyValueStore {
	return UbiquitousKeyValueStore{
		ID: objc.ID(ptr),
	}
}




