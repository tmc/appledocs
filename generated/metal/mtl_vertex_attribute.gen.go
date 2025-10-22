// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VertexAttribute] class.
var (
	VertexAttributeClass     _VertexAttributeClass
	VertexAttributeClassOnce sync.Once
)

func getVertexAttributeClass() _VertexAttributeClass {
	VertexAttributeClassOnce.Do(func() {
		VertexAttributeClass = _VertexAttributeClass{objc.GetClass("MTLVertexAttribute")}
	})
	return VertexAttributeClass
}

type _VertexAttributeClass struct {
	class objc.Class
}

// An interface definition for the [VertexAttribute] class.
type IVertexAttribute interface {
	objectivec.IObject
	VertexAttributes() MTLVertexAttribute
	SetVertexAttributes(value MTLVertexAttribute)
	AttributeIndex() int
	SetAttributeIndex(value int)
	AttributeType() unsafe.Pointer
	SetAttributeType(value unsafe.Pointer)
	IsActive() bool
	SetIsActive(value bool)
	IsPatchControlPointData() bool
	SetIsPatchControlPointData(value bool)
	IsPatchData() bool
	SetIsPatchData(value bool)
	Name() string
	SetName(value string)
}

// An instance that represents an attribute of a vertex function.
//
// An instance represents an attribute for per-vertex input in a vertex function. You use vertex attribute instances to inspect the inputs of a vertex function by examining the property of the corresponding instance.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexAttribute
type VertexAttribute struct {
	objectivec.Object
}

// VertexAttributeFrom constructs a [VertexAttribute] from an unsafe.Pointer.
//
// An instance that represents an attribute of a vertex function.
func VertexAttributeFrom(ptr unsafe.Pointer) VertexAttribute {
	return VertexAttribute{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VertexAttributeClass) Alloc() VertexAttribute {
	rv := objc.Send[VertexAttribute](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VertexAttributeClass) New() VertexAttribute {
	rv := objc.Send[VertexAttribute](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VertexAttribute) Init() VertexAttribute {
	rv := objc.Send[VertexAttribute](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VertexAttribute) Autorelease() VertexAttribute {
	rv := objc.Send[VertexAttribute](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVertexAttribute creates a new VertexAttribute instance.
func NewVertexAttribute() VertexAttribute {
	return getVertexAttributeClass().New()
}


// An array that describes the vertex input attributes to a vertex function.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunction/vertexattributes
func (v_ VertexAttribute) VertexAttributes() MTLVertexAttribute {
	rv := objc.Send[MTLVertexAttribute](v_.ID, objc.Sel("vertexAttributes"))
	return rv
}


// SetVertexAttributes sets the value of the vertexAttributes property.
// An array that describes the vertex input attributes to a vertex function.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunction/vertexattributes
func (v_ VertexAttribute) SetVertexAttributes(value MTLVertexAttribute) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setVertexAttributes:"), value)
}

// The index of the attribute, as declared in Metal shader source code.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexattribute/attributeindex
func (v_ VertexAttribute) AttributeIndex() int {
	rv := objc.Send[int](v_.ID, objc.Sel("attributeIndex"))
	return rv
}


// SetAttributeIndex sets the value of the attributeIndex property.
// The index of the attribute, as declared in Metal shader source code.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexattribute/attributeindex
func (v_ VertexAttribute) SetAttributeIndex(value int) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAttributeIndex:"), value)
}

// The data type for the attribute, as declared in Metal shader source code.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexattribute/attributetype
func (v_ VertexAttribute) AttributeType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("attributeType"))
	return rv
}


// SetAttributeType sets the value of the attributeType property.
// The data type for the attribute, as declared in Metal shader source code.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexattribute/attributetype
func (v_ VertexAttribute) SetAttributeType(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAttributeType:"), value)
}

// A Boolean value that indicates whether this vertex attribute is active.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexattribute/isactive
func (v_ VertexAttribute) IsActive() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isActive"))
	return rv
}


// SetIsActive sets the value of the isActive property.
// A Boolean value that indicates whether this vertex attribute is active.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexattribute/isactive
func (v_ VertexAttribute) SetIsActive(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsActive:"), value)
}

// A Boolean value that indicates whether this vertex attribute represents control point data.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexattribute/ispatchcontrolpointdata
func (v_ VertexAttribute) IsPatchControlPointData() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isPatchControlPointData"))
	return rv
}


// SetIsPatchControlPointData sets the value of the isPatchControlPointData property.
// A Boolean value that indicates whether this vertex attribute represents control point data.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexattribute/ispatchcontrolpointdata
func (v_ VertexAttribute) SetIsPatchControlPointData(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsPatchControlPointData:"), value)
}

// A Boolean value that indicates whether this vertex attribute represents patch data.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexattribute/ispatchdata
func (v_ VertexAttribute) IsPatchData() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isPatchData"))
	return rv
}


// SetIsPatchData sets the value of the isPatchData property.
// A Boolean value that indicates whether this vertex attribute represents patch data.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexattribute/ispatchdata
func (v_ VertexAttribute) SetIsPatchData(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsPatchData:"), value)
}

// The name of the attribute.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexattribute/name
func (v_ VertexAttribute) Name() string {
	rv := objc.Send[string](v_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// The name of the attribute.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexattribute/name
func (v_ VertexAttribute) SetName(value string) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setName:"), objc.String(value))
}



