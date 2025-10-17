// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PurgeableData] class.
var purgeableDataClass = _PurgeableDataClass{objc.GetClass("NSPurgeableData")}

type _PurgeableDataClass struct {
	class objc.Class
}

// An interface definition for the [PurgeableData] class.
type IPurgeableData interface {
	IMutableData
}

// A mutable data object containing bytes that can be discarded when they’re no longer needed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPurgeableData

type PurgeableData struct {
	MutableData
}

// PurgeableDataFrom constructs a [PurgeableData] from an unsafe.Pointer.
//
// A mutable data object containing bytes that can be discarded when they’re no longer needed.
func PurgeableDataFrom(ptr unsafe.Pointer) PurgeableData {
	return PurgeableData{
		MutableData: MutableDataFrom(ptr),
	}
}



