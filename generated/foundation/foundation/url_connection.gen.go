// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [URLConnection] class.
var URLConnectionClass objc.Class

func init() {
	URLConnectionClass = objc.GetClass("NSURLConnection")
}

type URLConnection struct {
	objc.ID
}

func URLConnectionFrom(ptr unsafe.Pointer) URLConnection {
	return URLConnection{
		ID: objc.ID(ptr),
	}
}




