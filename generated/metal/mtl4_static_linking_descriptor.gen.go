// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MTL4StaticLinkingDescriptor] class.
var (
	MTL4StaticLinkingDescriptorClass     _MTL4StaticLinkingDescriptorClass
	MTL4StaticLinkingDescriptorClassOnce sync.Once
)

func getMTL4StaticLinkingDescriptorClass() _MTL4StaticLinkingDescriptorClass {
	MTL4StaticLinkingDescriptorClassOnce.Do(func() {
		MTL4StaticLinkingDescriptorClass = _MTL4StaticLinkingDescriptorClass{objc.GetClass("MTL4StaticLinkingDescriptor")}
	})
	return MTL4StaticLinkingDescriptorClass
}

type _MTL4StaticLinkingDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [MTL4StaticLinkingDescriptor] class.
type IMTL4StaticLinkingDescriptor interface {
	objectivec.IObject
	

	// properties:
	FunctionDescriptors() []MTL4FunctionDescriptor
	SetFunctionDescriptors(value []MTL4FunctionDescriptor)
	Groups() foundation.IDictionary
	SetGroups(value foundation.IDictionary)
	PrivateFunctionDescriptors() []MTL4FunctionDescriptor
	SetPrivateFunctionDescriptors(value []MTL4FunctionDescriptor)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MTL4StaticLinkingDescriptorClass) Alloc() MTL4StaticLinkingDescriptor {
	rv := objc.Send[MTL4StaticLinkingDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4StaticLinkingDescriptorClass) New() MTL4StaticLinkingDescriptor {
	rv := objc.Send[MTL4StaticLinkingDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4StaticLinkingDescriptor) Init() MTL4StaticLinkingDescriptor {
	rv := objc.Send[MTL4StaticLinkingDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4StaticLinkingDescriptor) Autorelease() MTL4StaticLinkingDescriptor {
	rv := objc.Send[MTL4StaticLinkingDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4StaticLinkingDescriptor creates a new MTL4StaticLinkingDescriptor instance.
func NewMTL4StaticLinkingDescriptor() MTL4StaticLinkingDescriptor {
	return getMTL4StaticLinkingDescriptorClass().New()
}





// Groups together properties to drive a static linking process.


// Groups together properties to drive a static linking process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4StaticLinkingDescriptor
type MTL4StaticLinkingDescriptor struct {
	objectivec.Object
}

// MTL4StaticLinkingDescriptorFrom constructs a [MTL4StaticLinkingDescriptor] from an unsafe.Pointer.
//
// Groups together properties to drive a static linking process.
func MTL4StaticLinkingDescriptorFrom(ptr unsafe.Pointer) MTL4StaticLinkingDescriptor {
	return MTL4StaticLinkingDescriptor{objectivec.Object{objc.ID(ptr)}}
}

























// Provides an array of functions to link at the Metal IR level.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4StaticLinkingDescriptor/functionDescriptors
func (m_ MTL4StaticLinkingDescriptor) FunctionDescriptors() []MTL4FunctionDescriptor {
	rv := objc.Send[[]MTL4FunctionDescriptor](m_.ID, objc.Sel("functionDescriptors"))
	return rv
}


// Provides an array of functions to link at the Metal IR level.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4StaticLinkingDescriptor/functionDescriptors
func (m_ MTL4StaticLinkingDescriptor) SetFunctionDescriptors(value []MTL4FunctionDescriptor) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setFunctionDescriptors:"), nsArray)
}


// Assigns groups of functions to match call-site attributes in shader code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4StaticLinkingDescriptor/groups
func (m_ MTL4StaticLinkingDescriptor) Groups() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("groups"))
	return rv
}


// Assigns groups of functions to match call-site attributes in shader code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4StaticLinkingDescriptor/groups
func (m_ MTL4StaticLinkingDescriptor) SetGroups(value foundation.IDictionary) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroups:"), value)
}


// Provides an array of private functions to link at the Metal IR level.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4StaticLinkingDescriptor/privateFunctionDescriptors
func (m_ MTL4StaticLinkingDescriptor) PrivateFunctionDescriptors() []MTL4FunctionDescriptor {
	rv := objc.Send[[]MTL4FunctionDescriptor](m_.ID, objc.Sel("privateFunctionDescriptors"))
	return rv
}


// Provides an array of private functions to link at the Metal IR level.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4StaticLinkingDescriptor/privateFunctionDescriptors
func (m_ MTL4StaticLinkingDescriptor) SetPrivateFunctionDescriptors(value []MTL4FunctionDescriptor) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setPrivateFunctionDescriptors:"), nsArray)
}








