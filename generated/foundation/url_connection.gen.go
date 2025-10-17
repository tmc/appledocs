// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLConnection] class.
var uRLConnectionClass = _URLConnectionClass{objc.GetClass("NSURLConnection")}

type _URLConnectionClass struct {
	class objc.Class
}

// An interface definition for the [URLConnection] class.
type IURLConnection interface {
	objectivec.IObject
}

// An object that enables you to start and stop URL requests. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLConnection

type URLConnection struct {
	objectivec.Object
}

// URLConnectionFrom constructs a [URLConnection] from an unsafe.Pointer.
//
// An object that enables you to start and stop URL requests.
func URLConnectionFrom(ptr unsafe.Pointer) URLConnection {
	return URLConnection{objectivec.Object{objc.ID(ptr)}}
}



