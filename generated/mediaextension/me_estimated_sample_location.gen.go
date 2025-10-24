// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MEEstimatedSampleLocation */


/* debug [class_header]: Header for MEEstimatedSampleLocation */
// The class instance for the [MEEstimatedSampleLocation] class.
var (
	MEEstimatedSampleLocationClass     _MEEstimatedSampleLocationClass
	MEEstimatedSampleLocationClassOnce sync.Once
)

func getMEEstimatedSampleLocationClass() _MEEstimatedSampleLocationClass {
	MEEstimatedSampleLocationClassOnce.Do(func() {
		MEEstimatedSampleLocationClass = _MEEstimatedSampleLocationClass{objc.GetClass("MEEstimatedSampleLocation")}
	})
	return MEEstimatedSampleLocationClass
}

type _MEEstimatedSampleLocationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MEEstimatedSampleLocation */
// An interface definition for the [MEEstimatedSampleLocation] class.
type IMEEstimatedSampleLocation interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MEEstimatedSampleLocation */
	// properties:
	ByteSource() IMEByteSource
	EstimatedSampleLocation() objc.IObject /* cross-framework: SampleCursorStorageRange */
	RefinementDataLocation() objc.IObject /* cross-framework: SampleCursorStorageRange */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MEEstimatedSampleLocation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MEEstimatedSampleLocation */
// Alloc allocates a new instance without initialization.
func (mc _MEEstimatedSampleLocationClass) Alloc() MEEstimatedSampleLocation {
	rv := objc.Send[MEEstimatedSampleLocation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MEEstimatedSampleLocationClass) New() MEEstimatedSampleLocation {
	rv := objc.Send[MEEstimatedSampleLocation](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MEEstimatedSampleLocation) Init() MEEstimatedSampleLocation {
	rv := objc.Send[MEEstimatedSampleLocation](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MEEstimatedSampleLocation) Autorelease() MEEstimatedSampleLocation {
	rv := objc.Send[MEEstimatedSampleLocation](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMEEstimatedSampleLocation creates a new MEEstimatedSampleLocation instance.
func NewMEEstimatedSampleLocation() MEEstimatedSampleLocation {
	return getMEEstimatedSampleLocationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MEEstimatedSampleLocation */
// An object that provides information about the estimated sample location with the media.


// An object that provides information about the estimated sample location with the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEEstimatedSampleLocation
type MEEstimatedSampleLocation struct {
	objectivec.Object
}

// MEEstimatedSampleLocationFrom constructs a [MEEstimatedSampleLocation] from an unsafe.Pointer.
//
// An object that provides information about the estimated sample location with the media.
func MEEstimatedSampleLocationFrom(ptr unsafe.Pointer) MEEstimatedSampleLocation {
	return MEEstimatedSampleLocation{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MEEstimatedSampleLocation */

// Creates an estimated sample location object with the byte source, sample location, and data location that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEEstimatedSampleLocation/init(byteSource:estimatedSampleLocation:refinementDataLocation:)
func NewMEEstimatedSampleLocationWithByteSourceEstimatedSampleLocationRefinementDataLocation(byteSource IMEByteSource, estimatedSampleLocation objc.IObject /* cross-framework: SampleCursorStorageRange */, refinementDataLocation objc.IObject /* cross-framework: SampleCursorStorageRange */) MEEstimatedSampleLocation {
	instance := getMEEstimatedSampleLocationClass().Alloc()
	rv := objc.Send[MEEstimatedSampleLocation](instance.ID, objc.Sel("initWithByteSource:estimatedSampleLocation:refinementDataLocation:"), byteSource, estimatedSampleLocation, refinementDataLocation)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMEEstimatedSampleLocationWithByteSourceEstimatedSampleLocationRefinementDataLocation */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MEEstimatedSampleLocation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MEEstimatedSampleLocation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MEEstimatedSampleLocation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MEEstimatedSampleLocation */

// The byte source to use to read the data for the sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEEstimatedSampleLocation/byteSource
func (m_ MEEstimatedSampleLocation) ByteSource() IMEByteSource {
	rv := objc.Send[MEByteSource](m_.ID, objc.Sel("byteSource"))
	return rv
}/* debug [instance_properties/getter]: byteSource */


// The estimated starting file offset and size in bytes of the sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEEstimatedSampleLocation/estimatedSampleLocation
func (m_ MEEstimatedSampleLocation) EstimatedSampleLocation() objc.IObject /* cross-framework: SampleCursorStorageRange */ {
	rv := objc.Send[avfoundation.SampleCursorStorageRange](m_.ID, objc.Sel("estimatedSampleLocation"))
	return rv
}/* debug [instance_properties/getter]: estimatedSampleLocation */


// The starting file offset and size in bytes of the data necessary to provide an accurate sample location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEEstimatedSampleLocation/refinementDataLocation
func (m_ MEEstimatedSampleLocation) RefinementDataLocation() objc.IObject /* cross-framework: SampleCursorStorageRange */ {
	rv := objc.Send[avfoundation.SampleCursorStorageRange](m_.ID, objc.Sel("refinementDataLocation"))
	return rv
}/* debug [instance_properties/getter]: refinementDataLocation */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MEEstimatedSampleLocation */


