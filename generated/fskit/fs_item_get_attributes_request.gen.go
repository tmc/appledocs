// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FSItemGetAttributesRequest] class.
var (
	FSItemGetAttributesRequestClass     _FSItemGetAttributesRequestClass
	FSItemGetAttributesRequestClassOnce sync.Once
)

func getFSItemGetAttributesRequestClass() _FSItemGetAttributesRequestClass {
	FSItemGetAttributesRequestClassOnce.Do(func() {
		FSItemGetAttributesRequestClass = _FSItemGetAttributesRequestClass{objc.GetClass("FSItemGetAttributesRequest")}
	})
	return FSItemGetAttributesRequestClass
}

type _FSItemGetAttributesRequestClass struct {
	class objc.Class
}

// An interface definition for the [FSItemGetAttributesRequest] class.
type IFSItemGetAttributesRequest interface {
	objectivec.IObject
	IsAttributeWanted(attribute FSItemAttribute) bool
	WantedAttributes() FSItemAttribute
	SetWantedAttributes(value FSItemAttribute)
}

// A request to get attributes from an item.
//
// Methods that retrieve attributes use this type and inspect the property to determine which attributes to provide. FSKit calls the method to determine whether the request requires a given attribute.


// A request to get attributes from an item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/GetAttributesRequest

type FSItemGetAttributesRequest struct {
	objectivec.Object
}

// FSItemGetAttributesRequestFrom constructs a [FSItemGetAttributesRequest] from an unsafe.Pointer.
//
// A request to get attributes from an item.
func FSItemGetAttributesRequestFrom(ptr unsafe.Pointer) FSItemGetAttributesRequest {
	return FSItemGetAttributesRequest{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FSItemGetAttributesRequestClass) Alloc() FSItemGetAttributesRequest {
	rv := objc.Send[FSItemGetAttributesRequest](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FSItemGetAttributesRequestClass) New() FSItemGetAttributesRequest {
	rv := objc.Send[FSItemGetAttributesRequest](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSItemGetAttributesRequest) Init() FSItemGetAttributesRequest {
	rv := objc.Send[FSItemGetAttributesRequest](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSItemGetAttributesRequest) Autorelease() FSItemGetAttributesRequest {
	rv := objc.Send[FSItemGetAttributesRequest](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSItemGetAttributesRequest creates a new FSItemGetAttributesRequest instance.
func NewFSItemGetAttributesRequest() FSItemGetAttributesRequest {
	return getFSItemGetAttributesRequestClass().New()
}




// A method that indicates whether the request wants given attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/GetAttributesRequest/isAttributeWanted(_:)

func (f_ FSItemGetAttributesRequest) IsAttributeWanted(attribute FSItemAttribute) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isAttributeWanted:"), attribute)
	return rv
}


// The attributes requested by the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/GetAttributesRequest/wantedAttributes

func (f_ FSItemGetAttributesRequest) WantedAttributes() FSItemAttribute {
	rv := objc.Send[FSItemAttribute](f_.ID, objc.Sel("wantedAttributes"))
	return rv
}


// The attributes requested by the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/GetAttributesRequest/wantedAttributes

func (f_ FSItemGetAttributesRequest) SetWantedAttributes(value FSItemAttribute) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setWantedAttributes:"), value)
}



