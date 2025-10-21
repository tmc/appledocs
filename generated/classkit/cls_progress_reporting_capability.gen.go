// Code generated from Apple documentation for ClassKit. DO NOT EDIT.

package classkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

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

// An interface definition for the [SProgressReportingCapability] class.
type ISProgressReportingCapability interface {
	ISObject
}

// A progress reporting capability supported by a context.
//
// You use activities to report metrics about a student’s progress through the task associated with a context. Every activity automatically measures time spent performing the task, but you can provide additional information, like the percentage completion, or a final score. To help teachers understand what to expect from a context, create a set of instances — one for each kind of metric the context reports. Add the complete set to the context by calling the method. When you create a reporting capability, include a brief description of the capability as a localized string in the property. Schoolwork presents this to teachers to provide additional information about the metric.
//
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

// Alloc allocates a new instance without initialization.
func (sc _SProgressReportingCapabilityClass) Alloc() SProgressReportingCapability {
	rv := objc.Send[SProgressReportingCapability](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates a new progress reporting capability of the given type with a descriptive string.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSProgressReportingCapability/init(kind:details:)
func NewSProgressReportingCapabilityWithKindDetails(kind SProgressReportingCapabilityKind, details appkit.string) SProgressReportingCapability {
	instance := getSProgressReportingCapabilityClass().Alloc()
	rv := objc.Send[SProgressReportingCapability](instance.ID, objc.Sel("initWithKind:details:"), kind, details)
	rv.Autorelease()
	return rv
}


// A description of the capability presented to teachers.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSProgressReportingCapability/details
func (s_ SProgressReportingCapability) Details() appkit.string {
	rv := objc.Send[appkit.string](s_.ID, objc.Sel("details"))
	return rv
}

// The kind of progress reporting capability.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSProgressReportingCapability/kind-swift.property
func (s_ SProgressReportingCapability) Kind() SProgressReportingCapabilityKind {
	rv := objc.Send[SProgressReportingCapabilityKind](s_.ID, objc.Sel("kind"))
	return rv
}

// The kinds of progress reporting that the context can perform.
//
// [Full Topic]: https://developer.apple.com/documentation/classkit/clscontext/progressreportingcapabilities
func (s_ SProgressReportingCapability) ProgressReportingCapabilities() CLSProgressReportingCapability {
	rv := objc.Send[CLSProgressReportingCapability](s_.ID, objc.Sel("progressReportingCapabilities"))
	return rv
}


// SetProgressReportingCapabilities sets the value of the progressReportingCapabilities property.
// The kinds of progress reporting that the context can perform.

//
// [Full Topic]: https://developer.apple.com/documentation/classkit/clscontext/progressreportingcapabilities
func (s_ SProgressReportingCapability) SetProgressReportingCapabilities(value ICLSProgressReportingCapability) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setProgressReportingCapabilities:"), value)
}


