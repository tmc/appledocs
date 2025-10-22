// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [FSItemSetAttributesRequest] class.
var (
	FSItemSetAttributesRequestClass     _FSItemSetAttributesRequestClass
	FSItemSetAttributesRequestClassOnce sync.Once
)

func getFSItemSetAttributesRequestClass() _FSItemSetAttributesRequestClass {
	FSItemSetAttributesRequestClassOnce.Do(func() {
		FSItemSetAttributesRequestClass = _FSItemSetAttributesRequestClass{objc.GetClass("FSItemSetAttributesRequest")}
	})
	return FSItemSetAttributesRequestClass
}

type _FSItemSetAttributesRequestClass struct {
	class objc.Class
}

// An interface definition for the [FSItemSetAttributesRequest] class.
type IFSItemSetAttributesRequest interface {
	IFSItemAttributes
	WasAttributeConsumed(attribute FSItemAttribute) bool
	ConsumedAttributes() FSItemAttribute
	SetConsumedAttributes(value FSItemAttribute)
}

// A request to set attributes on an item.
//
// Methods that take attributes use this type to receive attribute values and to indicate which attributes they support. The various members of the parent type, , contain the values of the attributes to set. Modify the property to indicate which attributes your file system successfully used. FSKit calls the method to determine whether the file system successfully used a given attribute. Only set the attributes that your file system supports.


// A request to set attributes on an item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/SetAttributesRequest

type FSItemSetAttributesRequest struct {
	FSItemAttributes
}

// FSItemSetAttributesRequestFrom constructs a [FSItemSetAttributesRequest] from an unsafe.Pointer.
//
// A request to set attributes on an item.
func FSItemSetAttributesRequestFrom(ptr unsafe.Pointer) FSItemSetAttributesRequest {
	return FSItemSetAttributesRequest{
		FSItemAttributes: FSItemAttributesFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (fc _FSItemSetAttributesRequestClass) Alloc() FSItemSetAttributesRequest {
	rv := objc.Send[FSItemSetAttributesRequest](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FSItemSetAttributesRequestClass) New() FSItemSetAttributesRequest {
	rv := objc.Send[FSItemSetAttributesRequest](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSItemSetAttributesRequest) Init() FSItemSetAttributesRequest {
	rv := objc.Send[FSItemSetAttributesRequest](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSItemSetAttributesRequest) Autorelease() FSItemSetAttributesRequest {
	rv := objc.Send[FSItemSetAttributesRequest](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSItemSetAttributesRequest creates a new FSItemSetAttributesRequest instance.
func NewFSItemSetAttributesRequest() FSItemSetAttributesRequest {
	return getFSItemSetAttributesRequestClass().New()
}




// A method that indicates whether the file system used the given attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/SetAttributesRequest/wasAttributeConsumed(_:)

func (f_ FSItemSetAttributesRequest) WasAttributeConsumed(attribute FSItemAttribute) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("wasAttributeConsumed:"), attribute)
	return rv
}


// The attributes successfully used by the file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/SetAttributesRequest/consumedAttributes

func (f_ FSItemSetAttributesRequest) ConsumedAttributes() FSItemAttribute {
	rv := objc.Send[FSItemAttribute](f_.ID, objc.Sel("consumedAttributes"))
	return rv
}


// The attributes successfully used by the file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/SetAttributesRequest/consumedAttributes

func (f_ FSItemSetAttributesRequest) SetConsumedAttributes(value FSItemAttribute) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setConsumedAttributes:"), value)
}



