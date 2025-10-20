// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var uRLProtocolClass _URLProtocolClass

func init() {
	uRLProtocolClass = _URLProtocolClass{objc.GetClass("NSURLProtocol")}
}

type _URLProtocolClass struct {
	class objc.Class
}

type URLProtocol struct {
	objc.ID
}

func URLProtocolFrom(ptr unsafe.Pointer) URLProtocol {
	return URLProtocol{
		ID: objc.ID(ptr),
	}
}




