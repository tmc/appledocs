// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSURLSessionTask */


/* debug [class_header]: Header for NSURLSessionTask */
// The class instance for the [URLSessionTask] class.
var (
	URLSessionTaskClass     _URLSessionTaskClass
	URLSessionTaskClassOnce sync.Once
)

func getURLSessionTaskClass() _URLSessionTaskClass {
	URLSessionTaskClassOnce.Do(func() {
		URLSessionTaskClass = _URLSessionTaskClass{objc.GetClass("NSURLSessionTask")}
	})
	return URLSessionTaskClass
}

type _URLSessionTaskClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for URLSessionTask */
// An interface definition for the [URLSessionTask] class.
type IURLSessionTask interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for URLSessionTask */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for URLSessionTask */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for URLSessionTask */
// Alloc allocates a new instance without initialization.
func (uc _URLSessionTaskClass) Alloc() URLSessionTask {
	rv := objc.Send[URLSessionTask](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _URLSessionTaskClass) New() URLSessionTask {
	rv := objc.Send[URLSessionTask](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLSessionTask) Init() URLSessionTask {
	rv := objc.Send[URLSessionTask](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLSessionTask) Autorelease() URLSessionTask {
	rv := objc.Send[URLSessionTask](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLSessionTask creates a new URLSessionTask instance.
func NewURLSessionTask() URLSessionTask {
	return getURLSessionTaskClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for URLSessionTask */
// A parent class referenced by other AVFoundation classes.


// A parent class referenced by other AVFoundation classes. [Full Topic]
type URLSessionTask struct {
	objectivec.Object
}

// URLSessionTaskFrom constructs a [URLSessionTask] from an unsafe.Pointer.
//
// A parent class referenced by other AVFoundation classes.
func URLSessionTaskFrom(ptr unsafe.Pointer) URLSessionTask {
	return URLSessionTask{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for URLSessionTask *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for URLSessionTask */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for URLSessionTask */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for URLSessionTask */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for URLSessionTask */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSURLSessionTask */



