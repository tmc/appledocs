// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var uRLQueryItemClass _URLQueryItemClass

func init() {
	uRLQueryItemClass = _URLQueryItemClass{objc.GetClass("NSURLQueryItem")}
}

type _URLQueryItemClass struct {
	class objc.Class
}

type URLQueryItem struct {
	objc.ID
}

func URLQueryItemFrom(ptr unsafe.Pointer) URLQueryItem {
	return URLQueryItem{
		ID: objc.ID(ptr),
	}
}




