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

/* debug [class.gen.go]: Generating class MTRAttributeReport */


/* debug [class_header]: Header for MTRAttributeReport */
// The class instance for the [MTRAttributeReport] class.
var (
	MTRAttributeReportClass     _MTRAttributeReportClass
	MTRAttributeReportClassOnce sync.Once
)

func getMTRAttributeReportClass() _MTRAttributeReportClass {
	MTRAttributeReportClassOnce.Do(func() {
		MTRAttributeReportClass = _MTRAttributeReportClass{objc.GetClass("MTRAttributeReport")}
	})
	return MTRAttributeReportClass
}

type _MTRAttributeReportClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRAttributeReport */
// An interface definition for the [MTRAttributeReport] class.
type IMTRAttributeReport interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRAttributeReport */
	// properties:
	Error() objc.IObject /* cross-framework: Error */
	Path() IMTRAttributePath
	Value() objc.ID
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRAttributeReport */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRAttributeReport */
// Alloc allocates a new instance without initialization.
func (mc _MTRAttributeReportClass) Alloc() MTRAttributeReport {
	rv := objc.Send[MTRAttributeReport](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRAttributeReportClass) New() MTRAttributeReport {
	rv := objc.Send[MTRAttributeReport](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAttributeReport) Init() MTRAttributeReport {
	rv := objc.Send[MTRAttributeReport](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAttributeReport) Autorelease() MTRAttributeReport {
	rv := objc.Send[MTRAttributeReport](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAttributeReport creates a new MTRAttributeReport instance.
func NewMTRAttributeReport() MTRAttributeReport {
	return getMTRAttributeReportClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRAttributeReport */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAttributeReport
type MTRAttributeReport struct {
	objectivec.Object
}

// MTRAttributeReportFrom constructs a [MTRAttributeReport] from an unsafe.Pointer.
func MTRAttributeReportFrom(ptr unsafe.Pointer) MTRAttributeReport {
	return MTRAttributeReport{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRAttributeReport */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAttributeReport/init(responseValue:)
func NewMTRAttributeReportWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRAttributeReport {
	instance := getMTRAttributeReportClass().Alloc()
	rv := objc.Send[MTRAttributeReport](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRAttributeReportWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRAttributeReport */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRAttributeReport */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRAttributeReport */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRAttributeReport */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAttributeReport/error
func (m_ MTRAttributeReport) Error() objc.IObject /* cross-framework: Error */ {
	rv := objc.Send[coretelephony.Error](m_.ID, objc.Sel("error"))
	return rv
}/* debug [instance_properties/getter]: error */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAttributeReport/path
func (m_ MTRAttributeReport) Path() IMTRAttributePath {
	rv := objc.Send[MTRAttributePath](m_.ID, objc.Sel("path"))
	return rv
}/* debug [instance_properties/getter]: path */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAttributeReport/value
func (m_ MTRAttributeReport) Value() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRAttributeReport */


