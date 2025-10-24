// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MXCrashDiagnosticObjectiveCExceptionReason */


/* debug [class_header]: Header for MXCrashDiagnosticObjectiveCExceptionReason */
// The class instance for the [MXCrashDiagnosticObjectiveCExceptionReason] class.
var (
	MXCrashDiagnosticObjectiveCExceptionReasonClass     _MXCrashDiagnosticObjectiveCExceptionReasonClass
	MXCrashDiagnosticObjectiveCExceptionReasonClassOnce sync.Once
)

func getMXCrashDiagnosticObjectiveCExceptionReasonClass() _MXCrashDiagnosticObjectiveCExceptionReasonClass {
	MXCrashDiagnosticObjectiveCExceptionReasonClassOnce.Do(func() {
		MXCrashDiagnosticObjectiveCExceptionReasonClass = _MXCrashDiagnosticObjectiveCExceptionReasonClass{objc.GetClass("MXCrashDiagnosticObjectiveCExceptionReason")}
	})
	return MXCrashDiagnosticObjectiveCExceptionReasonClass
}

type _MXCrashDiagnosticObjectiveCExceptionReasonClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MXCrashDiagnosticObjectiveCExceptionReason */
// An interface definition for the [MXCrashDiagnosticObjectiveCExceptionReason] class.
type IMXCrashDiagnosticObjectiveCExceptionReason interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MXCrashDiagnosticObjectiveCExceptionReason */
	// properties:
	Arguments() []string
	ClassName() objc.IObject /* cross-framework: NSString */
	ComposedMessage() objc.IObject /* cross-framework: NSString */
	ExceptionName() objc.IObject /* cross-framework: NSString */
	ExceptionType() objc.IObject /* cross-framework: NSString */
	FormatString() objc.IObject /* cross-framework: NSString */
	MXErrorDomain() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MXCrashDiagnosticObjectiveCExceptionReason */
	// methods:
	DictionaryRepresentation() foundation.Dictionary
	JSONRepresentation() foundation.Data
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MXCrashDiagnosticObjectiveCExceptionReason */
// Alloc allocates a new instance without initialization.
func (mc _MXCrashDiagnosticObjectiveCExceptionReasonClass) Alloc() MXCrashDiagnosticObjectiveCExceptionReason {
	rv := objc.Send[MXCrashDiagnosticObjectiveCExceptionReason](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MXCrashDiagnosticObjectiveCExceptionReasonClass) New() MXCrashDiagnosticObjectiveCExceptionReason {
	rv := objc.Send[MXCrashDiagnosticObjectiveCExceptionReason](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXCrashDiagnosticObjectiveCExceptionReason) Init() MXCrashDiagnosticObjectiveCExceptionReason {
	rv := objc.Send[MXCrashDiagnosticObjectiveCExceptionReason](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXCrashDiagnosticObjectiveCExceptionReason) Autorelease() MXCrashDiagnosticObjectiveCExceptionReason {
	rv := objc.Send[MXCrashDiagnosticObjectiveCExceptionReason](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXCrashDiagnosticObjectiveCExceptionReason creates a new MXCrashDiagnosticObjectiveCExceptionReason instance.
func NewMXCrashDiagnosticObjectiveCExceptionReason() MXCrashDiagnosticObjectiveCExceptionReason {
	return getMXCrashDiagnosticObjectiveCExceptionReasonClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MXCrashDiagnosticObjectiveCExceptionReason */
// An object that represents the exception reason for an uncaught ObjC exception.
//
// The crash report for an uncaught Objective-C can contain detailed information about the type, name and description of the exception object. Use the properties and methods on to access this information.


// An object that represents the exception reason for an uncaught ObjC exception.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCrashDiagnosticObjectiveCExceptionReason
type MXCrashDiagnosticObjectiveCExceptionReason struct {
	objectivec.Object
}

// MXCrashDiagnosticObjectiveCExceptionReasonFrom constructs a [MXCrashDiagnosticObjectiveCExceptionReason] from an unsafe.Pointer.
//
// An object that represents the exception reason for an uncaught ObjC exception.
func MXCrashDiagnosticObjectiveCExceptionReasonFrom(ptr unsafe.Pointer) MXCrashDiagnosticObjectiveCExceptionReason {
	return MXCrashDiagnosticObjectiveCExceptionReason{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MXCrashDiagnosticObjectiveCExceptionReason *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MXCrashDiagnosticObjectiveCExceptionReason */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MXCrashDiagnosticObjectiveCExceptionReason */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MXCrashDiagnosticObjectiveCExceptionReason */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCrashDiagnosticObjectiveCExceptionReason/dictionaryRepresentation()
func (m_ MXCrashDiagnosticObjectiveCExceptionReason) DictionaryRepresentation() foundation.Dictionary {
	rv := objc.Send[foundation.Dictionary](m_.ID, objc.Sel("dictionaryRepresentation"))
	return rv
}/* debug [instance_methods/method]: DictionaryRepresentation */


// Returns the contents of the exception reason in JSON format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCrashDiagnosticObjectiveCExceptionReason/jsonRepresentation()
func (m_ MXCrashDiagnosticObjectiveCExceptionReason) JSONRepresentation() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("JSONRepresentation"))
	return rv
}/* debug [instance_methods/method]: JSONRepresentation */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MXCrashDiagnosticObjectiveCExceptionReason */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCrashDiagnosticObjectiveCExceptionReason/arguments
func (m_ MXCrashDiagnosticObjectiveCExceptionReason) Arguments() []string {
	rv := objc.Send[[]string](m_.ID, objc.Sel("arguments"))
	return rv
}/* debug [instance_properties/getter]: arguments */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCrashDiagnosticObjectiveCExceptionReason/className
func (m_ MXCrashDiagnosticObjectiveCExceptionReason) ClassName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("className"))
	return rv
}/* debug [instance_properties/getter]: className */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCrashDiagnosticObjectiveCExceptionReason/composedMessage
func (m_ MXCrashDiagnosticObjectiveCExceptionReason) ComposedMessage() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("composedMessage"))
	return rv
}/* debug [instance_properties/getter]: composedMessage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCrashDiagnosticObjectiveCExceptionReason/exceptionName
func (m_ MXCrashDiagnosticObjectiveCExceptionReason) ExceptionName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("exceptionName"))
	return rv
}/* debug [instance_properties/getter]: exceptionName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCrashDiagnosticObjectiveCExceptionReason/exceptionType
func (m_ MXCrashDiagnosticObjectiveCExceptionReason) ExceptionType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("exceptionType"))
	return rv
}/* debug [instance_properties/getter]: exceptionType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCrashDiagnosticObjectiveCExceptionReason/formatString
func (m_ MXCrashDiagnosticObjectiveCExceptionReason) FormatString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("formatString"))
	return rv
}/* debug [instance_properties/getter]: formatString */


// Error domain for error values from app metrics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxerrordomain
func (m_ MXCrashDiagnosticObjectiveCExceptionReason) MXErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MXErrorDomain"))
	return rv
}/* debug [instance_properties/getter]: MXErrorDomain */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MXCrashDiagnosticObjectiveCExceptionReason */



