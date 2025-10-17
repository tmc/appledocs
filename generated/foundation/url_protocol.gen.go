// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLProtocol] class.
var uRLProtocolClass = _URLProtocolClass{objc.GetClass("NSURLProtocol")}

type _URLProtocolClass struct {
	class objc.Class
}

// An interface definition for the [URLProtocol] class.
type IURLProtocol interface {
	objectivec.IObject
}

// An abstract class that handles the loading of protocol-specific URL data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtocol

type URLProtocol struct {
	objectivec.Object
}

// URLProtocolFrom constructs a [URLProtocol] from an unsafe.Pointer.
//
// An abstract class that handles the loading of protocol-specific URL data.
func URLProtocolFrom(ptr unsafe.Pointer) URLProtocol {
	return URLProtocol{objectivec.Object{objc.ID(ptr)}}
}



