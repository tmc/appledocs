// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var proxyClass _ProxyClass

func init() {
	proxyClass = _ProxyClass{objc.GetClass("NSProxy")}
}

type _ProxyClass struct {
	class objc.Class
}

type Proxy struct {
	objc.ID
}

func ProxyFrom(ptr unsafe.Pointer) Proxy {
	return Proxy{
		ID: objc.ID(ptr),
	}
}




