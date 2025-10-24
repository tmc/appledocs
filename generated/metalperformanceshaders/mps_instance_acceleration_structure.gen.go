// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSInstanceAccelerationStructure */


/* debug [class_header]: Header for MPSInstanceAccelerationStructure */
// The class instance for the [InstanceAccelerationStructure] class.
var (
	InstanceAccelerationStructureClass     _InstanceAccelerationStructureClass
	InstanceAccelerationStructureClassOnce sync.Once
)

func getInstanceAccelerationStructureClass() _InstanceAccelerationStructureClass {
	InstanceAccelerationStructureClassOnce.Do(func() {
		InstanceAccelerationStructureClass = _InstanceAccelerationStructureClass{objc.GetClass("MPSInstanceAccelerationStructure")}
	})
	return InstanceAccelerationStructureClass
}

type _InstanceAccelerationStructureClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for InstanceAccelerationStructure */
// An interface definition for the [InstanceAccelerationStructure] class.
type IInstanceAccelerationStructure interface {
	IAccelerationStructure
	
/* debug [class_interface_properties]: Properties for InstanceAccelerationStructure */
	// properties:
	AccelerationStructures() IMPSPolygonAccelerationStructure
	SetAccelerationStructures(value IMPSPolygonAccelerationStructure)
	InstanceBuffer() Buffer get set /* not a class type */
	SetInstanceBuffer(value Buffer get set /* not a class type */)
	InstanceBufferOffset() objectivec.IObject
	SetInstanceBufferOffset(value objectivec.IObject)
	InstanceCount() objectivec.IObject
	SetInstanceCount(value objectivec.IObject)
	MaskBuffer() Buffer get set /* not a class type */
	SetMaskBuffer(value Buffer get set /* not a class type */)
	MaskBufferOffset() objectivec.IObject
	SetMaskBufferOffset(value objectivec.IObject)
	TransformBuffer() Buffer get set /* not a class type */
	SetTransformBuffer(value Buffer get set /* not a class type */)
	TransformBufferOffset() objectivec.IObject
	SetTransformBufferOffset(value objectivec.IObject)
	TransformType() TransformType get set /* not a class type */
	SetTransformType(value TransformType get set /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for InstanceAccelerationStructure */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for InstanceAccelerationStructure */
// Alloc allocates a new instance without initialization.
func (ic _InstanceAccelerationStructureClass) Alloc() InstanceAccelerationStructure {
	rv := objc.Send[InstanceAccelerationStructure](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _InstanceAccelerationStructureClass) New() InstanceAccelerationStructure {
	rv := objc.Send[InstanceAccelerationStructure](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ InstanceAccelerationStructure) Init() InstanceAccelerationStructure {
	rv := objc.Send[InstanceAccelerationStructure](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ InstanceAccelerationStructure) Autorelease() InstanceAccelerationStructure {
	rv := objc.Send[InstanceAccelerationStructure](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewInstanceAccelerationStructure creates a new InstanceAccelerationStructure instance.
func NewInstanceAccelerationStructure() InstanceAccelerationStructure {
	return getInstanceAccelerationStructureClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for InstanceAccelerationStructure */
// An acceleration structure built over instances of other acceleration structures.


// An acceleration structure built over instances of other acceleration structures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSInstanceAccelerationStructure
type InstanceAccelerationStructure struct {
	AccelerationStructure
}

// InstanceAccelerationStructureFrom constructs a [InstanceAccelerationStructure] from an unsafe.Pointer.
//
// An acceleration structure built over instances of other acceleration structures.
func InstanceAccelerationStructureFrom(ptr unsafe.Pointer) InstanceAccelerationStructure {
	return InstanceAccelerationStructure{
		AccelerationStructure: AccelerationStructureFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for InstanceAccelerationStructure *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for InstanceAccelerationStructure */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for InstanceAccelerationStructure */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for InstanceAccelerationStructure */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for InstanceAccelerationStructure */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsinstanceaccelerationstructure/2980787-accelerationstructures
func (i_ InstanceAccelerationStructure) AccelerationStructures() IMPSPolygonAccelerationStructure {
	rv := objc.Send[PolygonAccelerationStructure](i_.ID, objc.Sel("accelerationStructures"))
	return rv
}/* debug [instance_properties/getter]: accelerationStructures */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsinstanceaccelerationstructure/2980787-accelerationstructures
func (i_ InstanceAccelerationStructure) SetAccelerationStructures(value IMPSPolygonAccelerationStructure) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAccelerationStructures:"), value)
}/* debug [instance_properties/setter]: accelerationStructures */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsinstanceaccelerationstructure/2980788-instancebuffer
func (i_ InstanceAccelerationStructure) InstanceBuffer() Buffer get set /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("instanceBuffer"))
	return rv
}/* debug [instance_properties/getter]: instanceBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsinstanceaccelerationstructure/2980788-instancebuffer
func (i_ InstanceAccelerationStructure) SetInstanceBuffer(value Buffer get set /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInstanceBuffer:"), value)
}/* debug [instance_properties/setter]: instanceBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsinstanceaccelerationstructure/2980789-instancebufferoffset
func (i_ InstanceAccelerationStructure) InstanceBufferOffset() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("instanceBufferOffset"))
	return rv
}/* debug [instance_properties/getter]: instanceBufferOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsinstanceaccelerationstructure/2980789-instancebufferoffset
func (i_ InstanceAccelerationStructure) SetInstanceBufferOffset(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInstanceBufferOffset:"), value)
}/* debug [instance_properties/setter]: instanceBufferOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsinstanceaccelerationstructure/2980790-instancecount
func (i_ InstanceAccelerationStructure) InstanceCount() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("instanceCount"))
	return rv
}/* debug [instance_properties/getter]: instanceCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsinstanceaccelerationstructure/2980790-instancecount
func (i_ InstanceAccelerationStructure) SetInstanceCount(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInstanceCount:"), value)
}/* debug [instance_properties/setter]: instanceCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsinstanceaccelerationstructure/2980791-maskbuffer
func (i_ InstanceAccelerationStructure) MaskBuffer() Buffer get set /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("maskBuffer"))
	return rv
}/* debug [instance_properties/getter]: maskBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsinstanceaccelerationstructure/2980791-maskbuffer
func (i_ InstanceAccelerationStructure) SetMaskBuffer(value Buffer get set /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMaskBuffer:"), value)
}/* debug [instance_properties/setter]: maskBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsinstanceaccelerationstructure/2980792-maskbufferoffset
func (i_ InstanceAccelerationStructure) MaskBufferOffset() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("maskBufferOffset"))
	return rv
}/* debug [instance_properties/getter]: maskBufferOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsinstanceaccelerationstructure/2980792-maskbufferoffset
func (i_ InstanceAccelerationStructure) SetMaskBufferOffset(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMaskBufferOffset:"), value)
}/* debug [instance_properties/setter]: maskBufferOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsinstanceaccelerationstructure/2980793-transformbuffer
func (i_ InstanceAccelerationStructure) TransformBuffer() Buffer get set /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("transformBuffer"))
	return rv
}/* debug [instance_properties/getter]: transformBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsinstanceaccelerationstructure/2980793-transformbuffer
func (i_ InstanceAccelerationStructure) SetTransformBuffer(value Buffer get set /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTransformBuffer:"), value)
}/* debug [instance_properties/setter]: transformBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsinstanceaccelerationstructure/2980794-transformbufferoffset
func (i_ InstanceAccelerationStructure) TransformBufferOffset() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("transformBufferOffset"))
	return rv
}/* debug [instance_properties/getter]: transformBufferOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsinstanceaccelerationstructure/2980794-transformbufferoffset
func (i_ InstanceAccelerationStructure) SetTransformBufferOffset(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTransformBufferOffset:"), value)
}/* debug [instance_properties/setter]: transformBufferOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsinstanceaccelerationstructure/2980795-transformtype
func (i_ InstanceAccelerationStructure) TransformType() TransformType get set /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("transformType"))
	return rv
}/* debug [instance_properties/getter]: transformType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsinstanceaccelerationstructure/2980795-transformtype
func (i_ InstanceAccelerationStructure) SetTransformType(value TransformType get set /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTransformType:"), value)
}/* debug [instance_properties/setter]: transformType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSInstanceAccelerationStructure */



