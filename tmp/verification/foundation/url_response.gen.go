// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var uRLResponseClass _URLResponseClass

func init() {
	uRLResponseClass = _URLResponseClass{objc.GetClass("NSURLResponse")}
}

type _URLResponseClass struct {
	class objc.Class
}

type URLResponse struct {
	objc.ID
}

func URLResponseFrom(ptr unsafe.Pointer) URLResponse {
	return URLResponse{
		ID: objc.ID(ptr),
	}
}




