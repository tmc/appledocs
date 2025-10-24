// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class FSProbeResult */


/* debug [class_header]: Header for FSProbeResult */
// The class instance for the [FSProbeResult] class.
var (
	FSProbeResultClass     _FSProbeResultClass
	FSProbeResultClassOnce sync.Once
)

func getFSProbeResultClass() _FSProbeResultClass {
	FSProbeResultClassOnce.Do(func() {
		FSProbeResultClass = _FSProbeResultClass{objc.GetClass("FSProbeResult")}
	})
	return FSProbeResultClass
}

type _FSProbeResultClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FSProbeResult */
// An interface definition for the [FSProbeResult] class.
type IFSProbeResult interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FSProbeResult */
	// properties:
	ContainerID() IFSContainerIdentifier
	Name() objc.IObject /* cross-framework: NSString */
	Result() FSMatchResult
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FSProbeResult */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FSProbeResult */
// Alloc allocates a new instance without initialization.
func (fc _FSProbeResultClass) Alloc() FSProbeResult {
	rv := objc.Send[FSProbeResult](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FSProbeResultClass) New() FSProbeResult {
	rv := objc.Send[FSProbeResult](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSProbeResult) Init() FSProbeResult {
	rv := objc.Send[FSProbeResult](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSProbeResult) Autorelease() FSProbeResult {
	rv := objc.Send[FSProbeResult](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSProbeResult creates a new FSProbeResult instance.
func NewFSProbeResult() FSProbeResult {
	return getFSProbeResultClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FSProbeResult */
// An object that represents the results of a specific probe.
//
// For any value other than , ensure the and values are non- . When a container or volume format doesn’t use a name, return an empty string. Also use an empty string in the case in which the format supports a name, but the value isn’t set yet. Some container or volume formats may lack a durable UUID on which to base a container identifier. This situation is only valid for unary file systems. In such a case, return a random UUID. With a block device resource, a probe operation may successfully get a result but encounter an error reading the name or UUID. If this happens, use whatever information is available, and provide an empty string or random UUID for the name or container ID, respectively.


// An object that represents the results of a specific probe.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSProbeResult
type FSProbeResult struct {
	objectivec.Object
}

// FSProbeResultFrom constructs a [FSProbeResult] from an unsafe.Pointer.
//
// An object that represents the results of a specific probe.
func FSProbeResultFrom(ptr unsafe.Pointer) FSProbeResult {
	return FSProbeResult{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FSProbeResult *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FSProbeResult */

// Creates a probe result for a recognized file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSProbeResult/recognized(name:containerID:)
func (fc _FSProbeResultClass) RecognizedProbeResultWithNameContainerID(name objc.IObject /* cross-framework: NSString */, containerID IFSContainerIdentifier) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("recognizedProbeResultWithName:containerID:"), name, containerID)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RecognizedProbeResultWithNameContainerID) */


// Creates a probe result for a recognized and usable file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSProbeResult/usable(name:containerID:)
func (fc _FSProbeResultClass) UsableProbeResultWithNameContainerID(name objc.IObject /* cross-framework: NSString */, containerID IFSContainerIdentifier) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("usableProbeResultWithName:containerID:"), name, containerID)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=UsableProbeResultWithNameContainerID) */


// Creates a probe result for a recognized file system that is usable, but with limited capabilities.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSProbeResult/usableButLimited(name:containerID:)
func (fc _FSProbeResultClass) UsableButLimitedProbeResultWithNameContainerID(name objc.IObject /* cross-framework: NSString */, containerID IFSContainerIdentifier) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("usableButLimitedProbeResultWithName:containerID:"), name, containerID)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=UsableButLimitedProbeResultWithNameContainerID) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FSProbeResult */

// A probe result for an unrecognized file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSProbeResult/notRecognized
func (fc _FSProbeResultClass) NotRecognizedProbeResult() FSProbeResult {
	rv := objc.Send[FSProbeResult](objc.ID(fc.class), objc.Sel("notRecognizedProbeResult"))
	return rv
}/* debug [class_properties_class/property]: notRecognizedProbeResult */

// A probe result for a recognized file system that is usable, but with limited capabilities.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSProbeResult/usableButLimited
func (fc _FSProbeResultClass) UsableButLimitedProbeResult() FSProbeResult {
	rv := objc.Send[FSProbeResult](objc.ID(fc.class), objc.Sel("usableButLimitedProbeResult"))
	return rv
}/* debug [class_properties_class/property]: usableButLimitedProbeResult */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FSProbeResult */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FSProbeResult */

// The container identifier, as found during the probe operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSProbeResult/containerID
func (f_ FSProbeResult) ContainerID() IFSContainerIdentifier {
	rv := objc.Send[FSContainerIdentifier](f_.ID, objc.Sel("containerID"))
	return rv
}/* debug [instance_properties/getter]: containerID */


// The resource name, as found during the probe operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSProbeResult/name
func (f_ FSProbeResult) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// A probe result for an unrecognized file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSProbeResult/notRecognized
func (f_ FSProbeResult) NotRecognizedProbeResult() IFSProbeResult {
	rv := objc.Send[FSProbeResult](f_.ID, objc.Sel("notRecognizedProbeResult"))
	return rv
}/* debug [instance_properties/getter]: notRecognizedProbeResult */


// The match result, representing the recognition and usability of a probed resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSProbeResult/result
func (f_ FSProbeResult) Result() FSMatchResult {
	rv := objc.Send[FSMatchResult](f_.ID, objc.Sel("result"))
	return rv
}/* debug [instance_properties/getter]: result */


// A probe result for a recognized file system that is usable, but with limited capabilities.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSProbeResult/usableButLimited
func (f_ FSProbeResult) UsableButLimitedProbeResult() IFSProbeResult {
	rv := objc.Send[FSProbeResult](f_.ID, objc.Sel("usableButLimitedProbeResult"))
	return rv
}/* debug [instance_properties/getter]: usableButLimitedProbeResult */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class FSProbeResult */



