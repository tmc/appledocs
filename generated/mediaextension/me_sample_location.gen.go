// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MESampleLocation */


/* debug [class_header]: Header for MESampleLocation */
// The class instance for the [MESampleLocation] class.
var (
	MESampleLocationClass     _MESampleLocationClass
	MESampleLocationClassOnce sync.Once
)

func getMESampleLocationClass() _MESampleLocationClass {
	MESampleLocationClassOnce.Do(func() {
		MESampleLocationClass = _MESampleLocationClass{objc.GetClass("MESampleLocation")}
	})
	return MESampleLocationClass
}

type _MESampleLocationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MESampleLocation */
// An interface definition for the [MESampleLocation] class.
type IMESampleLocation interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MESampleLocation */
	// properties:
	ByteSource() IMEByteSource
	SampleLocation() objc.IObject /* cross-framework: SampleCursorStorageRange */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MESampleLocation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MESampleLocation */
// Alloc allocates a new instance without initialization.
func (mc _MESampleLocationClass) Alloc() MESampleLocation {
	rv := objc.Send[MESampleLocation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MESampleLocationClass) New() MESampleLocation {
	rv := objc.Send[MESampleLocation](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MESampleLocation) Init() MESampleLocation {
	rv := objc.Send[MESampleLocation](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MESampleLocation) Autorelease() MESampleLocation {
	rv := objc.Send[MESampleLocation](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMESampleLocation creates a new MESampleLocation instance.
func NewMESampleLocation() MESampleLocation {
	return getMESampleLocationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MESampleLocation */
// An object that provides information about the sample location with the media.


// An object that provides information about the sample location with the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MESampleLocation
type MESampleLocation struct {
	objectivec.Object
}

// MESampleLocationFrom constructs a [MESampleLocation] from an unsafe.Pointer.
//
// An object that provides information about the sample location with the media.
func MESampleLocationFrom(ptr unsafe.Pointer) MESampleLocation {
	return MESampleLocation{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MESampleLocation */

// Creates a sample location object with the byte source and sample location that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MESampleLocation/init(byteSource:sampleLocation:)
func NewMESampleLocationWithByteSourceSampleLocation(byteSource IMEByteSource, sampleLocation objc.IObject /* cross-framework: SampleCursorStorageRange */) MESampleLocation {
	instance := getMESampleLocationClass().Alloc()
	rv := objc.Send[MESampleLocation](instance.ID, objc.Sel("initWithByteSource:sampleLocation:"), byteSource, sampleLocation)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMESampleLocationWithByteSourceSampleLocation */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MESampleLocation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MESampleLocation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MESampleLocation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MESampleLocation */

// The byte source to use to read the data for the sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MESampleLocation/byteSource
func (m_ MESampleLocation) ByteSource() IMEByteSource {
	rv := objc.Send[MEByteSource](m_.ID, objc.Sel("byteSource"))
	return rv
}/* debug [instance_properties/getter]: byteSource */


// The starting file offset and size in bytes of the sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MESampleLocation/sampleLocation
func (m_ MESampleLocation) SampleLocation() objc.IObject /* cross-framework: SampleCursorStorageRange */ {
	rv := objc.Send[avfoundation.SampleCursorStorageRange](m_.ID, objc.Sel("sampleLocation"))
	return rv
}/* debug [instance_properties/getter]: sampleLocation */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MESampleLocation */


