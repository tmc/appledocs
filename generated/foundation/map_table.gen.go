// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MapTable] class.
var mapTableClass = _MapTableClass{objc.GetClass("NSMapTable")}

type _MapTableClass struct {
	class objc.Class
}

// A collection similar to a dictionary, but with a broader range of available memory semantics. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapTable

type MapTable struct {
	objectivec.Object
}

// MapTableFrom constructs a [MapTable] from an unsafe.Pointer.
//
// A collection similar to a dictionary, but with a broader range of available memory semantics.
func MapTableFrom(ptr unsafe.Pointer) MapTable {
	return MapTable{objectivec.Object{objc.ID(ptr)}}
}



