// Code generated from Apple documentation for ClassKit. DO NOT EDIT.

package classkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CLSProgressReportingCapability */


/* debug [class_header]: Header for CLSProgressReportingCapability */
// The class instance for the [SProgressReportingCapability] class.
var (
	SProgressReportingCapabilityClass     _SProgressReportingCapabilityClass
	SProgressReportingCapabilityClassOnce sync.Once
)

func getSProgressReportingCapabilityClass() _SProgressReportingCapabilityClass {
	SProgressReportingCapabilityClassOnce.Do(func() {
		SProgressReportingCapabilityClass = _SProgressReportingCapabilityClass{objc.GetClass("CLSProgressReportingCapability")}
	})
	return SProgressReportingCapabilityClass
}

type _SProgressReportingCapabilityClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SProgressReportingCapability */
// An interface definition for the [SProgressReportingCapability] class.
type ISProgressReportingCapability interface {
	ISObject
	
/* debug [class_interface_properties]: Properties for SProgressReportingCapability */
	// properties:
	Details() objc.IObject /* cross-framework: NSString */
	Kind() SProgressReportingCapabilityKind
	ProgressReportingCapabilities() ICLSProgressReportingCapability
	SetProgressReportingCapabilities(value ICLSProgressReportingCapability)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SProgressReportingCapability */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SProgressReportingCapability */
// Alloc allocates a new instance without initialization.
func (sc _SProgressReportingCapabilityClass) Alloc() SProgressReportingCapability {
	rv := objc.Send[SProgressReportingCapability](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SProgressReportingCapabilityClass) New() SProgressReportingCapability {
	rv := objc.Send[SProgressReportingCapability](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SProgressReportingCapability) Init() SProgressReportingCapability {
	rv := objc.Send[SProgressReportingCapability](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SProgressReportingCapability) Autorelease() SProgressReportingCapability {
	rv := objc.Send[SProgressReportingCapability](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSProgressReportingCapability creates a new SProgressReportingCapability instance.
func NewSProgressReportingCapability() SProgressReportingCapability {
	return getSProgressReportingCapabilityClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SProgressReportingCapability */
// A progress reporting capability supported by a context.
//
// You use activities to report metrics about a student’s progress through the task associated with a context. Every activity automatically measures time spent performing the task, but you can provide additional information, like the percentage completion, or a final score. To help teachers understand what to expect from a context, create a set of instances — one for each kind of metric the context reports. Add the complete set to the context by calling the method. When you create a reporting capability, include a brief description of the capability as a localized string in the property. Schoolwork presents this to teachers to provide additional information about the metric.


// A progress reporting capability supported by a context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSProgressReportingCapability
type SProgressReportingCapability struct {
	SObject
}

// SProgressReportingCapabilityFrom constructs a [SProgressReportingCapability] from an unsafe.Pointer.
//
// A progress reporting capability supported by a context.
func SProgressReportingCapabilityFrom(ptr unsafe.Pointer) SProgressReportingCapability {
	return SProgressReportingCapability{
		SObject: SObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SProgressReportingCapability */

// Creates a new progress reporting capability of the given type with a descriptive string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSProgressReportingCapability/init(kind:details:)
func NewSProgressReportingCapabilityWithKindDetails(kind SProgressReportingCapabilityKind, details objc.IObject /* cross-framework: NSString */) SProgressReportingCapability {
	instance := getSProgressReportingCapabilityClass().Alloc()
	rv := objc.Send[SProgressReportingCapability](instance.ID, objc.Sel("initWithKind:details:"), kind, details)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSProgressReportingCapabilityWithKindDetails */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SProgressReportingCapability */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SProgressReportingCapability */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SProgressReportingCapability */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SProgressReportingCapability */

// A description of the capability presented to teachers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSProgressReportingCapability/details
func (s_ SProgressReportingCapability) Details() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("details"))
	return rv
}/* debug [instance_properties/getter]: details */


// The kind of progress reporting capability.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSProgressReportingCapability/kind-swift.property
func (s_ SProgressReportingCapability) Kind() SProgressReportingCapabilityKind {
	rv := objc.Send[SProgressReportingCapabilityKind](s_.ID, objc.Sel("kind"))
	return rv
}/* debug [instance_properties/getter]: kind */


// The kinds of progress reporting that the context can perform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/classkit/clscontext/progressreportingcapabilities
func (s_ SProgressReportingCapability) ProgressReportingCapabilities() ICLSProgressReportingCapability {
	rv := objc.Send[SProgressReportingCapability](s_.ID, objc.Sel("progressReportingCapabilities"))
	return rv
}/* debug [instance_properties/getter]: progressReportingCapabilities */


// The kinds of progress reporting that the context can perform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/classkit/clscontext/progressreportingcapabilities
func (s_ SProgressReportingCapability) SetProgressReportingCapabilities(value ICLSProgressReportingCapability) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setProgressReportingCapabilities:"), value)
}/* debug [instance_properties/setter]: progressReportingCapabilities */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CLSProgressReportingCapability */


