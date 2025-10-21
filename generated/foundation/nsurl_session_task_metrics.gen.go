// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [URLSessionTaskMetrics] class.
type IURLSessionTaskMetrics interface {
	objectivec.IObject
}

// An object encapsulating the metrics for a session task.
//
// Each object contains the and , as well as metrics for each request-and-response transaction made during the execution of the task.
//
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

// Alloc allocates a new instance without initialization.
func (uc _URLSessionTaskMetricsClass) Alloc() URLSessionTaskMetrics {
	rv := objc.Send[URLSessionTaskMetrics](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The number of redirects that occurred during the execution of the task.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskMetrics/redirectCount
func (u_ URLSessionTaskMetrics) RedirectCount() uint {
	rv := objc.Send[uint](u_.ID, objc.Sel("redirectCount"))
	return rv
}

// The time interval between when a task is instantiated and when the task is completed.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskMetrics/taskInterval
func (u_ URLSessionTaskMetrics) TaskInterval() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("taskInterval"))
	return rv
}

// An array of metrics for each individual request-response transaction made during the execution of the task.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskMetrics/transactionMetrics
func (u_ URLSessionTaskMetrics) TransactionMetrics() []URLSessionTaskTransactionMetrics {
	rv := objc.Send[[]URLSessionTaskTransactionMetrics](u_.ID, objc.Sel("transactionMetrics"))
	return rv
}


