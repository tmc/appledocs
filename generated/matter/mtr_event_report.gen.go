// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTREventReport */


/* debug [class_header]: Header for MTREventReport */
// The class instance for the [MTREventReport] class.
var (
	MTREventReportClass     _MTREventReportClass
	MTREventReportClassOnce sync.Once
)

func getMTREventReportClass() _MTREventReportClass {
	MTREventReportClassOnce.Do(func() {
		MTREventReportClass = _MTREventReportClass{objc.GetClass("MTREventReport")}
	})
	return MTREventReportClass
}

type _MTREventReportClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTREventReport */
// An interface definition for the [MTREventReport] class.
type IMTREventReport interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTREventReport */
	// properties:
	Error() objc.IObject /* cross-framework: Error */
	EventNumber() objc.IObject /* cross-framework: NSNumber */
	EventTimeType() unsafe.Pointer
	Path() IMTREventPath
	Priority() objc.IObject /* cross-framework: NSNumber */
	SystemUpTime() float64
	Timestamp() objc.IObject /* cross-framework: NSNumber */
	TimestampDate() objc.IObject /* cross-framework: NSDate */
	Value() objc.ID
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTREventReport */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTREventReport */
// Alloc allocates a new instance without initialization.
func (mc _MTREventReportClass) Alloc() MTREventReport {
	rv := objc.Send[MTREventReport](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTREventReportClass) New() MTREventReport {
	rv := objc.Send[MTREventReport](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREventReport) Init() MTREventReport {
	rv := objc.Send[MTREventReport](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREventReport) Autorelease() MTREventReport {
	rv := objc.Send[MTREventReport](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREventReport creates a new MTREventReport instance.
func NewMTREventReport() MTREventReport {
	return getMTREventReportClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTREventReport */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREventReport
type MTREventReport struct {
	objectivec.Object
}

// MTREventReportFrom constructs a [MTREventReport] from an unsafe.Pointer.
func MTREventReportFrom(ptr unsafe.Pointer) MTREventReport {
	return MTREventReport{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTREventReport */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREventReport/init(responseValue:)
func NewMTREventReportWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTREventReport {
	instance := getMTREventReportClass().Alloc()
	rv := objc.Send[MTREventReport](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTREventReportWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTREventReport */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTREventReport */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTREventReport */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTREventReport */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREventReport/error
func (m_ MTREventReport) Error() objc.IObject /* cross-framework: Error */ {
	rv := objc.Send[coretelephony.Error](m_.ID, objc.Sel("error"))
	return rv
}/* debug [instance_properties/getter]: error */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREventReport/eventNumber
func (m_ MTREventReport) EventNumber() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("eventNumber"))
	return rv
}/* debug [instance_properties/getter]: eventNumber */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREventReport/eventTimeType
func (m_ MTREventReport) EventTimeType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("eventTimeType"))
	return rv
}/* debug [instance_properties/getter]: eventTimeType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREventReport/path
func (m_ MTREventReport) Path() IMTREventPath {
	rv := objc.Send[MTREventPath](m_.ID, objc.Sel("path"))
	return rv
}/* debug [instance_properties/getter]: path */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREventReport/priority
func (m_ MTREventReport) Priority() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("priority"))
	return rv
}/* debug [instance_properties/getter]: priority */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREventReport/systemUpTime
func (m_ MTREventReport) SystemUpTime() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("systemUpTime"))
	return rv
}/* debug [instance_properties/getter]: systemUpTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREventReport/timestamp
func (m_ MTREventReport) Timestamp() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timestamp"))
	return rv
}/* debug [instance_properties/getter]: timestamp */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREventReport/timestampDate
func (m_ MTREventReport) TimestampDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("timestampDate"))
	return rv
}/* debug [instance_properties/getter]: timestampDate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREventReport/value
func (m_ MTREventReport) Value() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTREventReport */


