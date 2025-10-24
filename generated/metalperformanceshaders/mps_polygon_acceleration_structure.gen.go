// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSPolygonAccelerationStructure */


/* debug [class_header]: Header for MPSPolygonAccelerationStructure */
// The class instance for the [PolygonAccelerationStructure] class.
var (
	PolygonAccelerationStructureClass     _PolygonAccelerationStructureClass
	PolygonAccelerationStructureClassOnce sync.Once
)

func getPolygonAccelerationStructureClass() _PolygonAccelerationStructureClass {
	PolygonAccelerationStructureClassOnce.Do(func() {
		PolygonAccelerationStructureClass = _PolygonAccelerationStructureClass{objc.GetClass("MPSPolygonAccelerationStructure")}
	})
	return PolygonAccelerationStructureClass
}

type _PolygonAccelerationStructureClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PolygonAccelerationStructure */
// An interface definition for the [PolygonAccelerationStructure] class.
type IPolygonAccelerationStructure interface {
	IAccelerationStructure
	
/* debug [class_interface_properties]: Properties for PolygonAccelerationStructure */
	// properties:
	IndexBuffer() Buffer get set /* not a class type */
	SetIndexBuffer(value Buffer get set /* not a class type */)
	IndexBufferOffset() objectivec.IObject
	SetIndexBufferOffset(value objectivec.IObject)
	IndexType() DataType get set /* not a class type */
	SetIndexType(value DataType get set /* not a class type */)
	MaskBuffer() Buffer get set /* not a class type */
	SetMaskBuffer(value Buffer get set /* not a class type */)
	MaskBufferOffset() objectivec.IObject
	SetMaskBufferOffset(value objectivec.IObject)
	PolygonCount() objectivec.IObject
	SetPolygonCount(value objectivec.IObject)
	PolygonType() PolygonType get set /* not a class type */
	SetPolygonType(value PolygonType get set /* not a class type */)
	VertexBuffer() Buffer get set /* not a class type */
	SetVertexBuffer(value Buffer get set /* not a class type */)
	VertexBufferOffset() objectivec.IObject
	SetVertexBufferOffset(value objectivec.IObject)
	VertexStride() objectivec.IObject
	SetVertexStride(value objectivec.IObject)
	PolygonBuffers() IMPSPolygonBuffer
	SetPolygonBuffers(value IMPSPolygonBuffer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PolygonAccelerationStructure */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PolygonAccelerationStructure */
// Alloc allocates a new instance without initialization.
func (pc _PolygonAccelerationStructureClass) Alloc() PolygonAccelerationStructure {
	rv := objc.Send[PolygonAccelerationStructure](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PolygonAccelerationStructureClass) New() PolygonAccelerationStructure {
	rv := objc.Send[PolygonAccelerationStructure](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PolygonAccelerationStructure) Init() PolygonAccelerationStructure {
	rv := objc.Send[PolygonAccelerationStructure](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PolygonAccelerationStructure) Autorelease() PolygonAccelerationStructure {
	rv := objc.Send[PolygonAccelerationStructure](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPolygonAccelerationStructure creates a new PolygonAccelerationStructure instance.
func NewPolygonAccelerationStructure() PolygonAccelerationStructure {
	return getPolygonAccelerationStructureClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PolygonAccelerationStructure */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPolygonAccelerationStructure
type PolygonAccelerationStructure struct {
	AccelerationStructure
}

// PolygonAccelerationStructureFrom constructs a [PolygonAccelerationStructure] from an unsafe.Pointer.
func PolygonAccelerationStructureFrom(ptr unsafe.Pointer) PolygonAccelerationStructure {
	return PolygonAccelerationStructure{
		AccelerationStructure: AccelerationStructureFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PolygonAccelerationStructure *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PolygonAccelerationStructure */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PolygonAccelerationStructure */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PolygonAccelerationStructure */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PolygonAccelerationStructure */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/3088894-indexbuffer
func (p_ PolygonAccelerationStructure) IndexBuffer() Buffer get set /* not a class type */ {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("indexBuffer"))
	return rv
}/* debug [instance_properties/getter]: indexBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/3088894-indexbuffer
func (p_ PolygonAccelerationStructure) SetIndexBuffer(value Buffer get set /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIndexBuffer:"), value)
}/* debug [instance_properties/setter]: indexBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/3088895-indexbufferoffset
func (p_ PolygonAccelerationStructure) IndexBufferOffset() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("indexBufferOffset"))
	return rv
}/* debug [instance_properties/getter]: indexBufferOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/3088895-indexbufferoffset
func (p_ PolygonAccelerationStructure) SetIndexBufferOffset(value objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIndexBufferOffset:"), value)
}/* debug [instance_properties/setter]: indexBufferOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/3088896-indextype
func (p_ PolygonAccelerationStructure) IndexType() DataType get set /* not a class type */ {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("indexType"))
	return rv
}/* debug [instance_properties/getter]: indexType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/3088896-indextype
func (p_ PolygonAccelerationStructure) SetIndexType(value DataType get set /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIndexType:"), value)
}/* debug [instance_properties/setter]: indexType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/3088897-maskbuffer
func (p_ PolygonAccelerationStructure) MaskBuffer() Buffer get set /* not a class type */ {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("maskBuffer"))
	return rv
}/* debug [instance_properties/getter]: maskBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/3088897-maskbuffer
func (p_ PolygonAccelerationStructure) SetMaskBuffer(value Buffer get set /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMaskBuffer:"), value)
}/* debug [instance_properties/setter]: maskBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/3088898-maskbufferoffset
func (p_ PolygonAccelerationStructure) MaskBufferOffset() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("maskBufferOffset"))
	return rv
}/* debug [instance_properties/getter]: maskBufferOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/3088898-maskbufferoffset
func (p_ PolygonAccelerationStructure) SetMaskBufferOffset(value objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMaskBufferOffset:"), value)
}/* debug [instance_properties/setter]: maskBufferOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/3088899-polygoncount
func (p_ PolygonAccelerationStructure) PolygonCount() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("polygonCount"))
	return rv
}/* debug [instance_properties/getter]: polygonCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/3088899-polygoncount
func (p_ PolygonAccelerationStructure) SetPolygonCount(value objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPolygonCount:"), value)
}/* debug [instance_properties/setter]: polygonCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/3088900-polygontype
func (p_ PolygonAccelerationStructure) PolygonType() PolygonType get set /* not a class type */ {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("polygonType"))
	return rv
}/* debug [instance_properties/getter]: polygonType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/3088900-polygontype
func (p_ PolygonAccelerationStructure) SetPolygonType(value PolygonType get set /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPolygonType:"), value)
}/* debug [instance_properties/setter]: polygonType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/3088901-vertexbuffer
func (p_ PolygonAccelerationStructure) VertexBuffer() Buffer get set /* not a class type */ {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("vertexBuffer"))
	return rv
}/* debug [instance_properties/getter]: vertexBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/3088901-vertexbuffer
func (p_ PolygonAccelerationStructure) SetVertexBuffer(value Buffer get set /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVertexBuffer:"), value)
}/* debug [instance_properties/setter]: vertexBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/3088902-vertexbufferoffset
func (p_ PolygonAccelerationStructure) VertexBufferOffset() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("vertexBufferOffset"))
	return rv
}/* debug [instance_properties/getter]: vertexBufferOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/3088902-vertexbufferoffset
func (p_ PolygonAccelerationStructure) SetVertexBufferOffset(value objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVertexBufferOffset:"), value)
}/* debug [instance_properties/setter]: vertexBufferOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/3088903-vertexstride
func (p_ PolygonAccelerationStructure) VertexStride() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("vertexStride"))
	return rv
}/* debug [instance_properties/getter]: vertexStride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/3088903-vertexstride
func (p_ PolygonAccelerationStructure) SetVertexStride(value objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVertexStride:"), value)
}/* debug [instance_properties/setter]: vertexStride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/3152577-polygonbuffers
func (p_ PolygonAccelerationStructure) PolygonBuffers() IMPSPolygonBuffer {
	rv := objc.Send[PolygonBuffer](p_.ID, objc.Sel("polygonBuffers"))
	return rv
}/* debug [instance_properties/getter]: polygonBuffers */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/3152577-polygonbuffers
func (p_ PolygonAccelerationStructure) SetPolygonBuffers(value IMPSPolygonBuffer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPolygonBuffers:"), value)
}/* debug [instance_properties/setter]: polygonBuffers */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSPolygonAccelerationStructure */



