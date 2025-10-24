// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSURLSessionDownloadTask */


/* debug [class_header]: Header for NSURLSessionDownloadTask */
// The class instance for the [URLSessionDownloadTask] class.
var (
	URLSessionDownloadTaskClass     _URLSessionDownloadTaskClass
	URLSessionDownloadTaskClassOnce sync.Once
)

func getURLSessionDownloadTaskClass() _URLSessionDownloadTaskClass {
	URLSessionDownloadTaskClassOnce.Do(func() {
		URLSessionDownloadTaskClass = _URLSessionDownloadTaskClass{objc.GetClass("NSURLSessionDownloadTask")}
	})
	return URLSessionDownloadTaskClass
}

type _URLSessionDownloadTaskClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for URLSessionDownloadTask */
// An interface definition for the [URLSessionDownloadTask] class.
type IURLSessionDownloadTask interface {
	IURLSessionTask
	
/* debug [class_interface_properties]: Properties for URLSessionDownloadTask */
	// properties:
	Response() IURLResponse
	SetResponse(value IURLResponse)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for URLSessionDownloadTask */
	// methods:
	CancelByProducingResumeData(completionHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for URLSessionDownloadTask */
// Alloc allocates a new instance without initialization.
func (uc _URLSessionDownloadTaskClass) Alloc() URLSessionDownloadTask {
	rv := objc.Send[URLSessionDownloadTask](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _URLSessionDownloadTaskClass) New() URLSessionDownloadTask {
	rv := objc.Send[URLSessionDownloadTask](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLSessionDownloadTask) Init() URLSessionDownloadTask {
	rv := objc.Send[URLSessionDownloadTask](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLSessionDownloadTask) Autorelease() URLSessionDownloadTask {
	rv := objc.Send[URLSessionDownloadTask](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLSessionDownloadTask creates a new URLSessionDownloadTask instance.
func NewURLSessionDownloadTask() URLSessionDownloadTask {
	return getURLSessionDownloadTaskClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for URLSessionDownloadTask */
// A URL session task that stores downloaded data to a file.
//
// An is a concrete subclass of , which provides most of the methods for this class. Download tasks directly write the server’s response data to a temporary file, providing your app with progress updates as data arrives from the server. When you use download tasks in background sessions, these downloads continue even when your app is in the suspended state or otherwise not running. You can pause (cancel) download tasks and resume them later (assuming the server supports doing so). You can also resume downloads that failed because of network connectivity problems.


// A URL session task that stores downloaded data to a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionDownloadTask
type URLSessionDownloadTask struct {
	URLSessionTask
}

// URLSessionDownloadTaskFrom constructs a [URLSessionDownloadTask] from an unsafe.Pointer.
//
// A URL session task that stores downloaded data to a file.
func URLSessionDownloadTaskFrom(ptr unsafe.Pointer) URLSessionDownloadTask {
	return URLSessionDownloadTask{
		URLSessionTask: URLSessionTaskFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for URLSessionDownloadTask */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for URLSessionDownloadTask */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for URLSessionDownloadTask */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for URLSessionDownloadTask */

// Cancels a download and calls a callback with resume data for later use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionDownloadTask/cancel(byProducingResumeData:)
func (u_ URLSessionDownloadTask) CancelByProducingResumeData(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("cancelByProducingResumeData:"), completionHandler)
}/* debug [instance_methods/method]: CancelByProducingResumeData */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for URLSessionDownloadTask */

// The server’s response to the currently active request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/response
func (u_ URLSessionDownloadTask) Response() IURLResponse {
	rv := objc.Send[URLResponse](u_.ID, objc.Sel("response"))
	return rv
}/* debug [instance_properties/getter]: response */


// The server’s response to the currently active request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/response
func (u_ URLSessionDownloadTask) SetResponse(value IURLResponse) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setResponse:"), value)
}/* debug [instance_properties/setter]: response */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSURLSessionDownloadTask */


