// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKObject */


/* debug [class_header]: Header for HKObject */
// The class instance for the [HKObject] class.
var (
	HKObjectClass     _HKObjectClass
	HKObjectClassOnce sync.Once
)

func getHKObjectClass() _HKObjectClass {
	HKObjectClassOnce.Do(func() {
		HKObjectClass = _HKObjectClass{objc.GetClass("HKObject")}
	})
	return HKObjectClass
}

type _HKObjectClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKObject */
// An interface definition for the [HKObject] class.
type IHKObject interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKObject */
	// properties:
	Device() IHKDevice
	Metadata() foundation.IDictionary
	Source() IHKSource
	SourceRevision() IHKSourceRevision
	UUID() foundation.UUID
	HKPredicateKeyPathMetadata() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathUUID() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKObject */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKObject */
// Alloc allocates a new instance without initialization.
func (hc _HKObjectClass) Alloc() HKObject {
	rv := objc.Send[HKObject](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKObjectClass) New() HKObject {
	rv := objc.Send[HKObject](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKObject) Init() HKObject {
	rv := objc.Send[HKObject](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKObject) Autorelease() HKObject {
	rv := objc.Send[HKObject](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKObject creates a new HKObject instance.
func NewHKObject() HKObject {
	return getHKObjectClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKObject */
// A piece of data that can be stored inside the HealthKit store.
//
// The class is an abstract class. You should never instantiate a object directly. Instead, always work with one of its concrete subclasses: , , , or . HealthKit objects are all immutable. With a few exceptions (such as the object’s source revision), the object’s properties are set when the object is first created and they cannot change.


// A piece of data that can be stored inside the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObject
type HKObject struct {
	objectivec.Object
}

// HKObjectFrom constructs a [HKObject] from an unsafe.Pointer.
//
// A piece of data that can be stored inside the HealthKit store.
func HKObjectFrom(ptr unsafe.Pointer) HKObject {
	return HKObject{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKObject *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKObject */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKObject */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKObject */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKObject */

// The device that generated the data for this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObject/device
func (h_ HKObject) Device() IHKDevice {
	rv := objc.Send[HKDevice](h_.ID, objc.Sel("device"))
	return rv
}/* debug [instance_properties/getter]: device */


// The metadata for this HealthKit object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObject/metadata
func (h_ HKObject) Metadata() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](h_.ID, objc.Sel("metadata"))
	return rv
}/* debug [instance_properties/getter]: metadata */


// A HealthKit source, representing the app or device that created this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObject/source
func (h_ HKObject) Source() IHKSource {
	rv := objc.Send[HKSource](h_.ID, objc.Sel("source"))
	return rv
}/* debug [instance_properties/getter]: source */


// The app or device that created this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObject/sourceRevision
func (h_ HKObject) SourceRevision() IHKSourceRevision {
	rv := objc.Send[HKSourceRevision](h_.ID, objc.Sel("sourceRevision"))
	return rv
}/* debug [instance_properties/getter]: sourceRevision */


// The universally unique identifier (UUID) for this HealthKit object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObject/uuid
func (h_ HKObject) UUID() foundation.UUID {
	rv := objc.Send[foundation.UUID](h_.ID, objc.Sel("UUID"))
	return rv
}/* debug [instance_properties/getter]: UUID */


// The key path for accessing the object’s metadata dictionary inside a predicate format string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathmetadata
func (h_ HKObject) HKPredicateKeyPathMetadata() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathMetadata"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathMetadata */


// The key path for accessing the object’s UUID inside a predicate format string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathuuid
func (h_ HKObject) HKPredicateKeyPathUUID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathUUID"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathUUID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKObject */



