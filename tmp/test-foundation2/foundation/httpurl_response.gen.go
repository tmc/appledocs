// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var hTTPURLResponseClass _HTTPURLResponseClass

func init() {
	hTTPURLResponseClass = _HTTPURLResponseClass{objc.GetClass("NSHTTPURLResponse")}
}

type _HTTPURLResponseClass struct {
	class objc.Class
}

type HTTPURLResponse struct {
	objc.ID
}

func HTTPURLResponseFrom(ptr unsafe.Pointer) HTTPURLResponse {
	return HTTPURLResponse{
		ID: objc.ID(ptr),
	}
}




