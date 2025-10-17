// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [UbiquitousKeyValueStore] class.
var ubiquitousKeyValueStoreClass = _UbiquitousKeyValueStoreClass{objc.GetClass("NSUbiquitousKeyValueStore")}

type _UbiquitousKeyValueStoreClass struct {
	class objc.Class
}

// An interface definition for the [UbiquitousKeyValueStore] class.
type IUbiquitousKeyValueStore interface {
	objectivec.IObject
}

// An iCloud-based container of key-value pairs you use to share data among instances of your app running on a user’s connected devices. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore

type UbiquitousKeyValueStore struct {
	objectivec.Object
}

// UbiquitousKeyValueStoreFrom constructs a [UbiquitousKeyValueStore] from an unsafe.Pointer.
//
// An iCloud-based container of key-value pairs you use to share data among instances of your app running on a user’s connected devices.
func UbiquitousKeyValueStoreFrom(ptr unsafe.Pointer) UbiquitousKeyValueStore {
	return UbiquitousKeyValueStore{objectivec.Object{objc.ID(ptr)}}
}



