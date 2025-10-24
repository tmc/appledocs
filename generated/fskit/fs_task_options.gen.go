// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class FSTaskOptions */


/* debug [class_header]: Header for FSTaskOptions */
// The class instance for the [FSTaskOptions] class.
var (
	FSTaskOptionsClass     _FSTaskOptionsClass
	FSTaskOptionsClassOnce sync.Once
)

func getFSTaskOptionsClass() _FSTaskOptionsClass {
	FSTaskOptionsClassOnce.Do(func() {
		FSTaskOptionsClass = _FSTaskOptionsClass{objc.GetClass("FSTaskOptions")}
	})
	return FSTaskOptionsClass
}

type _FSTaskOptionsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FSTaskOptions */
// An interface definition for the [FSTaskOptions] class.
type IFSTaskOptions interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FSTaskOptions */
	// properties:
	TaskOptions() []string
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FSTaskOptions */
	// methods:
	UrlForOption(option objc.IObject /* cross-framework: NSString */) foundation.URL
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FSTaskOptions */
// Alloc allocates a new instance without initialization.
func (fc _FSTaskOptionsClass) Alloc() FSTaskOptions {
	rv := objc.Send[FSTaskOptions](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FSTaskOptionsClass) New() FSTaskOptions {
	rv := objc.Send[FSTaskOptions](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSTaskOptions) Init() FSTaskOptions {
	rv := objc.Send[FSTaskOptions](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSTaskOptions) Autorelease() FSTaskOptions {
	rv := objc.Send[FSTaskOptions](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSTaskOptions creates a new FSTaskOptions instance.
func NewFSTaskOptions() FSTaskOptions {
	return getFSTaskOptionsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FSTaskOptions */
// A class that passes command options to a task, optionally providing security-scoped URLs.


// A class that passes command options to a task, optionally providing security-scoped URLs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSTaskOptions
type FSTaskOptions struct {
	objectivec.Object
}

// FSTaskOptionsFrom constructs a [FSTaskOptions] from an unsafe.Pointer.
//
// A class that passes command options to a task, optionally providing security-scoped URLs.
func FSTaskOptionsFrom(ptr unsafe.Pointer) FSTaskOptions {
	return FSTaskOptions{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FSTaskOptions *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FSTaskOptions */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FSTaskOptions */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FSTaskOptions */

// Retrieves a URL for a given option.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSTaskOptions/url(forOption:)
func (f_ FSTaskOptions) UrlForOption(option objc.IObject /* cross-framework: NSString */) foundation.URL {
	rv := objc.Send[foundation.URL](f_.ID, objc.Sel("urlForOption:"), option)
	return rv
}/* debug [instance_methods/method]: UrlForOption */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FSTaskOptions */

// An array of strings that represent command-line options for the task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSTaskOptions/taskOptions
func (f_ FSTaskOptions) TaskOptions() []string {
	rv := objc.Send[[]string](f_.ID, objc.Sel("taskOptions"))
	return rv
}/* debug [instance_properties/getter]: taskOptions */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class FSTaskOptions */



