// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [UUID] class.
var uUIDClass = _UUIDClass{objc.GetClass("NSUUID")}

type _UUIDClass struct {
	class objc.Class
}

// An interface definition for the [UUID] class.
type IUUID interface {
	objectivec.IObject
}

// A universally unique value that can be used to identify types, interfaces, and other items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUUID

type UUID struct {
	objectivec.Object
}

// UUIDFrom constructs a [UUID] from an unsafe.Pointer.
//
// A universally unique value that can be used to identify types, interfaces, and other items.
func UUIDFrom(ptr unsafe.Pointer) UUID {
	return UUID{objectivec.Object{objc.ID(ptr)}}
}



