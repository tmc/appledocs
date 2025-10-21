// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLSessionTaskTransactionMetrics] class.
var (
	URLSessionTaskTransactionMetricsClass     _URLSessionTaskTransactionMetricsClass
	URLSessionTaskTransactionMetricsClassOnce sync.Once
)

func getURLSessionTaskTransactionMetricsClass() _URLSessionTaskTransactionMetricsClass {
	URLSessionTaskTransactionMetricsClassOnce.Do(func() {
		URLSessionTaskTransactionMetricsClass = _URLSessionTaskTransactionMetricsClass{objc.GetClass("NSURLSessionTaskTransactionMetrics")}
	})
	return URLSessionTaskTransactionMetricsClass
}

type _URLSessionTaskTransactionMetricsClass struct {
	class objc.Class
}

// An interface definition for the [URLSessionTaskTransactionMetrics] class.
type IURLSessionTaskTransactionMetrics interface {
	objectivec.IObject
}

// An object that encapsualtes the performance metrics collected by the URL Loading System during the execution of a session task.
//
// Each object consists of a and property, corresponding to the request and response of the corresponding task. It also contains temporal metrics, starting with and ending with , as well as other characteristics like and .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics
type URLSessionTaskTransactionMetrics struct {
	objectivec.Object
}

// URLSessionTaskTransactionMetricsFrom constructs a [URLSessionTaskTransactionMetrics] from an unsafe.Pointer.
//
// An object that encapsualtes the performance metrics collected by the URL Loading System during the execution of a session task.
func URLSessionTaskTransactionMetricsFrom(ptr unsafe.Pointer) URLSessionTaskTransactionMetrics {
	return URLSessionTaskTransactionMetrics{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _URLSessionTaskTransactionMetricsClass) Alloc() URLSessionTaskTransactionMetrics {
	rv := objc.Send[URLSessionTaskTransactionMetrics](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _URLSessionTaskTransactionMetricsClass) New() URLSessionTaskTransactionMetrics {
	rv := objc.Send[URLSessionTaskTransactionMetrics](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLSessionTaskTransactionMetrics) Init() URLSessionTaskTransactionMetrics {
	rv := objc.Send[URLSessionTaskTransactionMetrics](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLSessionTaskTransactionMetrics) Autorelease() URLSessionTaskTransactionMetrics {
	rv := objc.Send[URLSessionTaskTransactionMetrics](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLSessionTaskTransactionMetrics creates a new URLSessionTaskTransactionMetrics instance.
func NewURLSessionTaskTransactionMetrics() URLSessionTaskTransactionMetrics {
	return getURLSessionTaskTransactionMetricsClass().New()
}




