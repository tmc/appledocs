// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DistantObject] class.
var distantObjectClass = _DistantObjectClass{objc.GetClass("NSDistantObject")}

type _DistantObjectClass struct {
	class objc.Class
}

// A proxy for objects in other applications or threads. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDistantObject

type DistantObject struct {
	Proxy
}

// DistantObjectFrom constructs a [DistantObject] from an unsafe.Pointer.
//
// A proxy for objects in other applications or threads.
func DistantObjectFrom(ptr unsafe.Pointer) DistantObject {
	return DistantObject{
		Proxy: ProxyFrom(ptr),
	}
}



