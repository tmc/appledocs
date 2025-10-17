// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [URLQueryItem] class.
var URLQueryItemClass objc.Class

func init() {
	URLQueryItemClass = objc.GetClass("NSURLQueryItem")
}

type URLQueryItem struct {
	objc.ID
}

func URLQueryItemFrom(ptr unsafe.Pointer) URLQueryItem {
	return URLQueryItem{
		ID: objc.ID(ptr),
	}
}




