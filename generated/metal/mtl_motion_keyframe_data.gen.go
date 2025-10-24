// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLMotionKeyframeData */


/* debug [class_header]: Header for MTLMotionKeyframeData */
// The class instance for the [MotionKeyframeData] class.
var (
	MotionKeyframeDataClass     _MotionKeyframeDataClass
	MotionKeyframeDataClassOnce sync.Once
)

func getMotionKeyframeDataClass() _MotionKeyframeDataClass {
	MotionKeyframeDataClassOnce.Do(func() {
		MotionKeyframeDataClass = _MotionKeyframeDataClass{objc.GetClass("MTLMotionKeyframeData")}
	})
	return MotionKeyframeDataClass
}

type _MotionKeyframeDataClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MotionKeyframeData */
// An interface definition for the [MotionKeyframeData] class.
type IMotionKeyframeData interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MotionKeyframeData */
	// properties:
	Buffer() unsafe.Pointer
	SetBuffer(value unsafe.Pointer)
	Offset() uint
	SetOffset(value uint)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MotionKeyframeData */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MotionKeyframeData */
// Alloc allocates a new instance without initialization.
func (mc _MotionKeyframeDataClass) Alloc() MotionKeyframeData {
	rv := objc.Send[MotionKeyframeData](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MotionKeyframeDataClass) New() MotionKeyframeData {
	rv := objc.Send[MotionKeyframeData](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MotionKeyframeData) Init() MotionKeyframeData {
	rv := objc.Send[MotionKeyframeData](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MotionKeyframeData) Autorelease() MotionKeyframeData {
	rv := objc.Send[MotionKeyframeData](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMotionKeyframeData creates a new MotionKeyframeData instance.
func NewMotionKeyframeData() MotionKeyframeData {
	return getMotionKeyframeDataClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MotionKeyframeData */
// Geometry data for a specific keyframe to use in a moving instance.
//
// An instance describes the location of geometry data for a keyframe. The exact type of data can vary, depending on which kind of motion descriptor you create. For an instance, the buffer data is a list of bounding boxes. For an , the buffer data is a list of vertices.


// Geometry data for a specific keyframe to use in a moving instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMotionKeyframeData
type MotionKeyframeData struct {
	objectivec.Object
}

// MotionKeyframeDataFrom constructs a [MotionKeyframeData] from an unsafe.Pointer.
//
// Geometry data for a specific keyframe to use in a moving instance.
func MotionKeyframeDataFrom(ptr unsafe.Pointer) MotionKeyframeData {
	return MotionKeyframeData{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MotionKeyframeData *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MotionKeyframeData */

// Creates a new keyframe object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMotionKeyframeData/data
func (mc _MotionKeyframeDataClass) Data() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("data"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Data) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MotionKeyframeData */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MotionKeyframeData */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MotionKeyframeData */

// The buffer that holds the geometry data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMotionKeyframeData/buffer
func (m_ MotionKeyframeData) Buffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("buffer"))
	return rv
}/* debug [instance_properties/getter]: buffer */


// The buffer that holds the geometry data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMotionKeyframeData/buffer
func (m_ MotionKeyframeData) SetBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBuffer:"), value)
}/* debug [instance_properties/setter]: buffer */


// The offset, in bytes, to the keyframe data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMotionKeyframeData/offset
func (m_ MotionKeyframeData) Offset() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("offset"))
	return rv
}/* debug [instance_properties/getter]: offset */


// The offset, in bytes, to the keyframe data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMotionKeyframeData/offset
func (m_ MotionKeyframeData) SetOffset(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOffset:"), value)
}/* debug [instance_properties/setter]: offset */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLMotionKeyframeData */



