// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSURLSessionDataTask */


/* debug [class_header]: Header for NSURLSessionDataTask */
// The class instance for the [URLSessionDataTask] class.
var (
	URLSessionDataTaskClass     _URLSessionDataTaskClass
	URLSessionDataTaskClassOnce sync.Once
)

func getURLSessionDataTaskClass() _URLSessionDataTaskClass {
	URLSessionDataTaskClassOnce.Do(func() {
		URLSessionDataTaskClass = _URLSessionDataTaskClass{objc.GetClass("NSURLSessionDataTask")}
	})
	return URLSessionDataTaskClass
}

type _URLSessionDataTaskClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for URLSessionDataTask */
// An interface definition for the [URLSessionDataTask] class.
type IURLSessionDataTask interface {
	IURLSessionTask
	
/* debug [class_interface_properties]: Properties for URLSessionDataTask */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for URLSessionDataTask */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for URLSessionDataTask */
// Alloc allocates a new instance without initialization.
func (uc _URLSessionDataTaskClass) Alloc() URLSessionDataTask {
	rv := objc.Send[URLSessionDataTask](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _URLSessionDataTaskClass) New() URLSessionDataTask {
	rv := objc.Send[URLSessionDataTask](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLSessionDataTask) Init() URLSessionDataTask {
	rv := objc.Send[URLSessionDataTask](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLSessionDataTask) Autorelease() URLSessionDataTask {
	rv := objc.Send[URLSessionDataTask](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLSessionDataTask creates a new URLSessionDataTask instance.
func NewURLSessionDataTask() URLSessionDataTask {
	return getURLSessionDataTaskClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for URLSessionDataTask */
// A URL session task that returns downloaded data directly to the app in memory.
//
// A is a concrete subclass of . The methods in the class are documented in . A data task returns data directly to the app (in memory) as one or more objects. When you use a data task: During upload of the body data (if your app provides any), the session periodically calls its delegate’s method with status information. After receiving an initial response, the session calls its delegate’s method to let you examine the status code and headers, and optionally convert the data task into a download task. During the transfer, the session calls its delegate’s method to provide your app with the content as it arrives. Upon completion, the session calls its delegate’s method to let you determine whether the response should be cached. For examples of using data tasks for fetching and uploading data, see and .


// A URL session task that returns downloaded data directly to the app in memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionDataTask
type URLSessionDataTask struct {
	URLSessionTask
}

// URLSessionDataTaskFrom constructs a [URLSessionDataTask] from an unsafe.Pointer.
//
// A URL session task that returns downloaded data directly to the app in memory.
func URLSessionDataTaskFrom(ptr unsafe.Pointer) URLSessionDataTask {
	return URLSessionDataTask{
		URLSessionTask: URLSessionTaskFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for URLSessionDataTask */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for URLSessionDataTask */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for URLSessionDataTask */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for URLSessionDataTask */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for URLSessionDataTask */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSURLSessionDataTask */


