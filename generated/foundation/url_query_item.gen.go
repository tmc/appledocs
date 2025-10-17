// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLQueryItem] class.
var uRLQueryItemClass = _URLQueryItemClass{objc.GetClass("NSURLQueryItem")}

type _URLQueryItemClass struct {
	class objc.Class
}

// An object representing a single name/value pair for an item in the query portion of a URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLQueryItem

type URLQueryItem struct {
	objectivec.Object
}

// URLQueryItemFrom constructs a [URLQueryItem] from an unsafe.Pointer.
//
// An object representing a single name/value pair for an item in the query portion of a URL.
func URLQueryItemFrom(ptr unsafe.Pointer) URLQueryItem {
	return URLQueryItem{objectivec.Object{objc.ID(ptr)}}
}



