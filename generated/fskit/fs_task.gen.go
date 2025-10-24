// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class FSTask */


/* debug [class_header]: Header for FSTask */
// The class instance for the [FSTask] class.
var (
	FSTaskClass     _FSTaskClass
	FSTaskClassOnce sync.Once
)

func getFSTaskClass() _FSTaskClass {
	FSTaskClassOnce.Do(func() {
		FSTaskClass = _FSTaskClass{objc.GetClass("FSTask")}
	})
	return FSTaskClass
}

type _FSTaskClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FSTask */
// An interface definition for the [FSTask] class.
type IFSTask interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FSTask */
	// properties:
	CancellationHandler() unsafe.Pointer
	SetCancellationHandler(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FSTask */
	// methods:
	DidCompleteWithError(error_ objc.IObject /* cross-framework: Error */)
	LogMessage(str objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FSTask */
// Alloc allocates a new instance without initialization.
func (fc _FSTaskClass) Alloc() FSTask {
	rv := objc.Send[FSTask](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FSTaskClass) New() FSTask {
	rv := objc.Send[FSTask](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSTask) Init() FSTask {
	rv := objc.Send[FSTask](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSTask) Autorelease() FSTask {
	rv := objc.Send[FSTask](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSTask creates a new FSTask instance.
func NewFSTask() FSTask {
	return getFSTaskClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FSTask */
// A class that enables a file system module to pass log messages and completion notifications to clients.
//
// FSKit creates an instance of this class for each long-running operations.


// A class that enables a file system module to pass log messages and completion notifications to clients.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSTask
type FSTask struct {
	objectivec.Object
}

// FSTaskFrom constructs a [FSTask] from an unsafe.Pointer.
//
// A class that enables a file system module to pass log messages and completion notifications to clients.
func FSTaskFrom(ptr unsafe.Pointer) FSTask {
	return FSTask{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FSTask *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FSTask */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FSTask */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FSTask */

// Informs the client that the task completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSTask/didComplete(error:)
func (f_ FSTask) DidCompleteWithError(error_ objc.IObject /* cross-framework: Error */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("didCompleteWithError:"), error_)
}/* debug [instance_methods/method]: DidCompleteWithError */


// Logs the given string to the initiating client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSTask/logMessage(_:)
func (f_ FSTask) LogMessage(str objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("logMessage:"), str)
}/* debug [instance_methods/method]: LogMessage */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FSTask */

// A handler called by FSKit upon canceling the task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSTask/cancellationHandler
func (f_ FSTask) CancellationHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("cancellationHandler"))
	return rv
}/* debug [instance_properties/getter]: cancellationHandler */


// A handler called by FSKit upon canceling the task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSTask/cancellationHandler
func (f_ FSTask) SetCancellationHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setCancellationHandler:"), value)
}/* debug [instance_properties/setter]: cancellationHandler */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class FSTask */



