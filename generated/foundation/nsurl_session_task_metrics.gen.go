// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSURLSessionTaskMetrics */


/* debug [class_header]: Header for NSURLSessionTaskMetrics */
// The class instance for the [URLSessionTaskMetrics] class.
var (
	URLSessionTaskMetricsClass     _URLSessionTaskMetricsClass
	URLSessionTaskMetricsClassOnce sync.Once
)

func getURLSessionTaskMetricsClass() _URLSessionTaskMetricsClass {
	URLSessionTaskMetricsClassOnce.Do(func() {
		URLSessionTaskMetricsClass = _URLSessionTaskMetricsClass{objc.GetClass("NSURLSessionTaskMetrics")}
	})
	return URLSessionTaskMetricsClass
}

type _URLSessionTaskMetricsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for URLSessionTaskMetrics */
// An interface definition for the [URLSessionTaskMetrics] class.
type IURLSessionTaskMetrics interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for URLSessionTaskMetrics */
	// properties:
	RedirectCount() uint
	TaskInterval() IDateInterval
	TransactionMetrics() []URLSessionTaskTransactionMetrics
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for URLSessionTaskMetrics */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for URLSessionTaskMetrics */
// Alloc allocates a new instance without initialization.
func (uc _URLSessionTaskMetricsClass) Alloc() URLSessionTaskMetrics {
	rv := objc.Send[URLSessionTaskMetrics](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _URLSessionTaskMetricsClass) New() URLSessionTaskMetrics {
	rv := objc.Send[URLSessionTaskMetrics](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLSessionTaskMetrics) Init() URLSessionTaskMetrics {
	rv := objc.Send[URLSessionTaskMetrics](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLSessionTaskMetrics) Autorelease() URLSessionTaskMetrics {
	rv := objc.Send[URLSessionTaskMetrics](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLSessionTaskMetrics creates a new URLSessionTaskMetrics instance.
func NewURLSessionTaskMetrics() URLSessionTaskMetrics {
	return getURLSessionTaskMetricsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for URLSessionTaskMetrics */
// An object encapsulating the metrics for a session task.
//
// Each object contains the and , as well as metrics for each request-and-response transaction made during the execution of the task.


// An object encapsulating the metrics for a session task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskMetrics
type URLSessionTaskMetrics struct {
	objectivec.Object
}

// URLSessionTaskMetricsFrom constructs a [URLSessionTaskMetrics] from an unsafe.Pointer.
//
// An object encapsulating the metrics for a session task.
func URLSessionTaskMetricsFrom(ptr unsafe.Pointer) URLSessionTaskMetrics {
	return URLSessionTaskMetrics{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for URLSessionTaskMetrics */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for URLSessionTaskMetrics */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for URLSessionTaskMetrics */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for URLSessionTaskMetrics */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for URLSessionTaskMetrics */

// The number of redirects that occurred during the execution of the task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskMetrics/redirectCount
func (u_ URLSessionTaskMetrics) RedirectCount() uint {
	rv := objc.Send[uint](u_.ID, objc.Sel("redirectCount"))
	return rv
}/* debug [instance_properties/getter]: redirectCount */


// The time interval between when a task is instantiated and when the task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskMetrics/taskInterval
func (u_ URLSessionTaskMetrics) TaskInterval() IDateInterval {
	rv := objc.Send[DateInterval](u_.ID, objc.Sel("taskInterval"))
	return rv
}/* debug [instance_properties/getter]: taskInterval */


// An array of metrics for each individual request-response transaction made during the execution of the task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskMetrics/transactionMetrics
func (u_ URLSessionTaskMetrics) TransactionMetrics() []URLSessionTaskTransactionMetrics {
	rv := objc.Send[[]URLSessionTaskTransactionMetrics](u_.ID, objc.Sel("transactionMetrics"))
	return rv
}/* debug [instance_properties/getter]: transactionMetrics */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSURLSessionTaskMetrics */


