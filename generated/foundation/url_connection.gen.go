// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [URLConnection] class.
var URLConnectionClass = _URLConnectionClass{objc.GetClass("NSURLConnection")}

type _URLConnectionClass struct {
	class objc.Class
}

type URLConnection struct {
	objc.ID
}

func URLConnectionFrom(ptr unsafe.Pointer) URLConnection {
	return URLConnection{
		ID: objc.ID(ptr),
	}
}




