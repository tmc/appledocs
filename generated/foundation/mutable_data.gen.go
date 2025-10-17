// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MutableData] class.
var mutableDataClass = _MutableDataClass{objc.GetClass("NSMutableData")}

type _MutableDataClass struct {
	class objc.Class
}

// An interface definition for the [MutableData] class.
type IMutableData interface {
	IData
}

// An object representing a dynamic byte buffer in memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableData

type MutableData struct {
	Data
}

// MutableDataFrom constructs a [MutableData] from an unsafe.Pointer.
//
// An object representing a dynamic byte buffer in memory.
func MutableDataFrom(ptr unsafe.Pointer) MutableData {
	return MutableData{
		Data: DataFrom(ptr),
	}
}



