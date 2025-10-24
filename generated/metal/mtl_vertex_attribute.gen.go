// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLVertexAttribute */


/* debug [class_header]: Header for MTLVertexAttribute */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VertexAttribute */
// An interface definition for the [VertexAttribute] class.
type IVertexAttribute interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VertexAttribute */
	// properties:
	AttributeIndex() uint
	AttributeType() DataType
	Active() bool
	PatchControlPointData() bool
	PatchData() bool
	Name() objc.IObject /* cross-framework: NSString */
	VertexAttributes() IMTLVertexAttribute
	SetVertexAttributes(value IMTLVertexAttribute)
	IsActive() bool
	SetIsActive(value bool)
	IsPatchControlPointData() bool
	SetIsPatchControlPointData(value bool)
	IsPatchData() bool
	SetIsPatchData(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VertexAttribute */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VertexAttribute */
// Alloc allocates a new instance without initialization.
func (vc _VertexAttributeClass) Alloc() VertexAttribute {
	rv := objc.Send[VertexAttribute](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VertexAttribute */
// An instance that represents an attribute of a vertex function.
//
// An instance represents an attribute for per-vertex input in a vertex function. You use vertex attribute instances to inspect the inputs of a vertex function by examining the property of the corresponding instance.


// An instance that represents an attribute of a vertex function.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VertexAttribute *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VertexAttribute */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VertexAttribute */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VertexAttribute */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VertexAttribute */

// The index of the attribute, as declared in Metal shader source code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexAttribute/attributeIndex
func (v_ VertexAttribute) AttributeIndex() uint {
	rv := objc.Send[uint](v_.ID, objc.Sel("attributeIndex"))
	return rv
}/* debug [instance_properties/getter]: attributeIndex */


// The data type for the attribute, as declared in Metal shader source code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexAttribute/attributeType
func (v_ VertexAttribute) AttributeType() DataType {
	rv := objc.Send[DataType](v_.ID, objc.Sel("attributeType"))
	return rv
}/* debug [instance_properties/getter]: attributeType */


// A Boolean value that indicates whether this vertex attribute is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexAttribute/isActive
func (v_ VertexAttribute) Active() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("active"))
	return rv
}/* debug [instance_properties/getter]: active */


// A Boolean value that indicates whether this vertex attribute represents control point data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexAttribute/isPatchControlPointData
func (v_ VertexAttribute) PatchControlPointData() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("patchControlPointData"))
	return rv
}/* debug [instance_properties/getter]: patchControlPointData */


// A Boolean value that indicates whether this vertex attribute represents patch data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexAttribute/isPatchData
func (v_ VertexAttribute) PatchData() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("patchData"))
	return rv
}/* debug [instance_properties/getter]: patchData */


// The name of the attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexAttribute/name
func (v_ VertexAttribute) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](v_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// An array that describes the vertex input attributes to a vertex function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunction/vertexattributes
func (v_ VertexAttribute) VertexAttributes() IMTLVertexAttribute {
	rv := objc.Send[VertexAttribute](v_.ID, objc.Sel("vertexAttributes"))
	return rv
}/* debug [instance_properties/getter]: vertexAttributes */


// An array that describes the vertex input attributes to a vertex function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunction/vertexattributes
func (v_ VertexAttribute) SetVertexAttributes(value IMTLVertexAttribute) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setVertexAttributes:"), value)
}/* debug [instance_properties/setter]: vertexAttributes */


// A Boolean value that indicates whether this vertex attribute is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexattribute/isactive
func (v_ VertexAttribute) IsActive() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isActive"))
	return rv
}/* debug [instance_properties/getter]: isActive */


// A Boolean value that indicates whether this vertex attribute is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexattribute/isactive
func (v_ VertexAttribute) SetIsActive(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsActive:"), value)
}/* debug [instance_properties/setter]: isActive */


// A Boolean value that indicates whether this vertex attribute represents control point data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexattribute/ispatchcontrolpointdata
func (v_ VertexAttribute) IsPatchControlPointData() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isPatchControlPointData"))
	return rv
}/* debug [instance_properties/getter]: isPatchControlPointData */


// A Boolean value that indicates whether this vertex attribute represents control point data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexattribute/ispatchcontrolpointdata
func (v_ VertexAttribute) SetIsPatchControlPointData(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsPatchControlPointData:"), value)
}/* debug [instance_properties/setter]: isPatchControlPointData */


// A Boolean value that indicates whether this vertex attribute represents patch data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexattribute/ispatchdata
func (v_ VertexAttribute) IsPatchData() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isPatchData"))
	return rv
}/* debug [instance_properties/getter]: isPatchData */


// A Boolean value that indicates whether this vertex attribute represents patch data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexattribute/ispatchdata
func (v_ VertexAttribute) SetIsPatchData(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsPatchData:"), value)
}/* debug [instance_properties/setter]: isPatchData */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLVertexAttribute */



