// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [URLComponents] class.
var URLComponentsClass objc.Class

func init() {
	URLComponentsClass = objc.GetClass("NSURLComponents")
}

type URLComponents struct {
	objc.ID
}

func URLComponentsFrom(ptr unsafe.Pointer) URLComponents {
	return URLComponents{
		ID: objc.ID(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc URLComponents) Alloc() URLComponents {
	ret := objc.ID(URLComponentsClass).Send(objc.RegisterName("alloc"))
	return URLComponents{ret}
}

// Init initializes the instance.
func (u_ URLComponents) Init() URLComponents {
	ret := u_.ID.Send(objc.RegisterName("init"))
	return URLComponents{ret}
}
// Creates a URL components object with all components left undefined. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURLComponents/init()
func NewURLComponents() URLComponents {
	instance := URLComponents{}.Alloc()
	instance = instance.Init()
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}



