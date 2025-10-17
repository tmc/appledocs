// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Proxy] class.
var ProxyClass objc.Class

func init() {
	ProxyClass = objc.GetClass("NSProxy")
}

type Proxy struct {
	objc.ID
}

func ProxyFrom(ptr unsafe.Pointer) Proxy {
	return Proxy{
		ID: objc.ID(ptr),
	}
}




