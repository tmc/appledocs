// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [Attribute] class.
var (
	AttributeClass     _AttributeClass
	AttributeClassOnce sync.Once
)

func getAttributeClass() _AttributeClass {
	AttributeClassOnce.Do(func() {
		AttributeClass = _AttributeClass{objc.GetClass("MTLAttribute")}
	})
	return AttributeClass
}

type _AttributeClass struct {
	class objc.Class
}





// An interface definition for the [Attribute] class.
type IAttribute interface {
	objectivec.IObject
	

	// properties:
	AttributeIndex() uint
	AttributeType() DataType
	Active() bool
	PatchControlPointData() bool
	PatchData() bool
	Name() foundation.foundation.INSString
	IsActive() bool
	SetIsActive(value bool)
	IsPatchControlPointData() bool
	SetIsPatchControlPointData(value bool)
	IsPatchData() bool
	SetIsPatchData(value bool)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ac _AttributeClass) Alloc() Attribute {
	rv := objc.Send[Attribute](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AttributeClass) New() Attribute {
	rv := objc.Send[Attribute](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ Attribute) Init() Attribute {
	rv := objc.Send[Attribute](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ Attribute) Autorelease() Attribute {
	rv := objc.Send[Attribute](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAttribute creates a new Attribute instance.
func NewAttribute() Attribute {
	return getAttributeClass().New()
}





// An object that describes an attribute defined in the stage-in argument for a shader.


// An object that describes an attribute defined in the stage-in argument for a shader.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttribute
type Attribute struct {
	objectivec.Object
}

// AttributeFrom constructs a [Attribute] from an unsafe.Pointer.
//
// An object that describes an attribute defined in the stage-in argument for a shader.
func AttributeFrom(ptr unsafe.Pointer) Attribute {
	return Attribute{objectivec.Object{objc.ID(ptr)}}
}

























// The index of the attribute, as declared in Metal shader source code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttribute/attributeIndex
func (a_ Attribute) AttributeIndex() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("attributeIndex"))
	return rv
}


// The data type for the attribute, as declared in Metal shader source code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttribute/attributeType
func (a_ Attribute) AttributeType() DataType {
	rv := objc.Send[DataType](a_.ID, objc.Sel("attributeType"))
	return rv
}


// A Boolean value that indicates whether the attribute is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttribute/isActive
func (a_ Attribute) Active() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("active"))
	return rv
}


// A Boolean value that indicates whether the attribute represents control point data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttribute/isPatchControlPointData
func (a_ Attribute) PatchControlPointData() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("patchControlPointData"))
	return rv
}


// A Boolean value that indicates whether the attribute represents tessellation patch data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttribute/isPatchData
func (a_ Attribute) PatchData() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("patchData"))
	return rv
}


// The name of the attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttribute/name
func (a_ Attribute) Name() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("name"))
	return rv
}


// A Boolean value that indicates whether the attribute is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlattribute/isactive
func (a_ Attribute) IsActive() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isActive"))
	return rv
}


// A Boolean value that indicates whether the attribute is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlattribute/isactive
func (a_ Attribute) SetIsActive(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsActive:"), value)
}


// A Boolean value that indicates whether the attribute represents control point data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlattribute/ispatchcontrolpointdata
func (a_ Attribute) IsPatchControlPointData() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isPatchControlPointData"))
	return rv
}


// A Boolean value that indicates whether the attribute represents control point data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlattribute/ispatchcontrolpointdata
func (a_ Attribute) SetIsPatchControlPointData(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsPatchControlPointData:"), value)
}


// A Boolean value that indicates whether the attribute represents tessellation patch data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlattribute/ispatchdata
func (a_ Attribute) IsPatchData() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isPatchData"))
	return rv
}


// A Boolean value that indicates whether the attribute represents tessellation patch data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlattribute/ispatchdata
func (a_ Attribute) SetIsPatchData(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsPatchData:"), value)
}








