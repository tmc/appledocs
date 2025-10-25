// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSURLSessionUploadTask */


/* debug [class_header]: Header for NSURLSessionUploadTask */
// The class instance for the [URLSessionUploadTask] class.
var (
	URLSessionUploadTaskClass     _URLSessionUploadTaskClass
	URLSessionUploadTaskClassOnce sync.Once
)

func getURLSessionUploadTaskClass() _URLSessionUploadTaskClass {
	URLSessionUploadTaskClassOnce.Do(func() {
		URLSessionUploadTaskClass = _URLSessionUploadTaskClass{objc.GetClass("NSURLSessionUploadTask")}
	})
	return URLSessionUploadTaskClass
}

type _URLSessionUploadTaskClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for URLSessionUploadTask */
// An interface definition for the [URLSessionUploadTask] class.
type IURLSessionUploadTask interface {
	IURLSessionDataTask
	
/* debug [class_interface_properties]: Properties for URLSessionUploadTask */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for URLSessionUploadTask */
	// methods:
	CancelByProducingResumeData(completionHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for URLSessionUploadTask */
// Alloc allocates a new instance without initialization.
func (uc _URLSessionUploadTaskClass) Alloc() URLSessionUploadTask {
	rv := objc.Send[URLSessionUploadTask](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _URLSessionUploadTaskClass) New() URLSessionUploadTask {
	rv := objc.Send[URLSessionUploadTask](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLSessionUploadTask) Init() URLSessionUploadTask {
	rv := objc.Send[URLSessionUploadTask](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLSessionUploadTask) Autorelease() URLSessionUploadTask {
	rv := objc.Send[URLSessionUploadTask](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLSessionUploadTask creates a new URLSessionUploadTask instance.
func NewURLSessionUploadTask() URLSessionUploadTask {
	return getURLSessionUploadTaskClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for URLSessionUploadTask */
// A URL session task that uploads data to the network in a request body.
//
// The class is a subclass of , which in turn is a concrete subclass of . The methods associated with the class are documented in . Upload tasks are used for making HTTP requests that require a request body (such as or ). They behave similarly to data tasks, but you create them by calling different methods on the session that are designed to make it easier to provide the content to upload. As with data tasks, if the server provides a response, upload tasks return that response as one or more objects in memory. When you create an upload task, you provide a instance that contains any additional headers that you might need to send alongside the upload, such as the content type, content disposition, and so on. In iOS, when you create an upload task for a file in a background session, the system copies that file to a temporary location and streams data from there. While the upload is in progress, the task calls the session delegate’s method periodically to provide you with status information. When the upload phase of the request finishes, the task behaves like a data task, calling methods on the session delegate to provide you with the server’s response—headers, status code, content data, and so on.


// A URL session task that uploads data to the network in a request body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionUploadTask
type URLSessionUploadTask struct {
	URLSessionDataTask
}

// URLSessionUploadTaskFrom constructs a [URLSessionUploadTask] from an unsafe.Pointer.
//
// A URL session task that uploads data to the network in a request body.
func URLSessionUploadTaskFrom(ptr unsafe.Pointer) URLSessionUploadTask {
	return URLSessionUploadTask{
		URLSessionDataTask: URLSessionDataTaskFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for URLSessionUploadTask */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for URLSessionUploadTask */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for URLSessionUploadTask */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for URLSessionUploadTask */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionUploadTask/cancel(byProducingResumeData:)
func (u_ URLSessionUploadTask) CancelByProducingResumeData(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("cancelByProducingResumeData:"), completionHandler)
}/* debug [instance_methods/method]: CancelByProducingResumeData */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for URLSessionUploadTask */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSURLSessionUploadTask */


