// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKProcessPool */


/* debug [class_header]: Header for WKProcessPool */
// The class instance for the [ProcessPool] class.
var (
	ProcessPoolClass     _ProcessPoolClass
	ProcessPoolClassOnce sync.Once
)

func getProcessPoolClass() _ProcessPoolClass {
	ProcessPoolClassOnce.Do(func() {
		ProcessPoolClass = _ProcessPoolClass{objc.GetClass("WKProcessPool")}
	})
	return ProcessPoolClass
}

type _ProcessPoolClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ProcessPool */
// An interface definition for the [ProcessPool] class.
type IProcessPool interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ProcessPool */
	// properties:
	ProcessPool() IWKProcessPool
	SetProcessPool(value IWKProcessPool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ProcessPool */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ProcessPool */
// Alloc allocates a new instance without initialization.
func (pc _ProcessPoolClass) Alloc() ProcessPool {
	rv := objc.Send[ProcessPool](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _ProcessPoolClass) New() ProcessPool {
	rv := objc.Send[ProcessPool](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ ProcessPool) Init() ProcessPool {
	rv := objc.Send[ProcessPool](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ ProcessPool) Autorelease() ProcessPool {
	rv := objc.Send[ProcessPool](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewProcessPool creates a new ProcessPool instance.
func NewProcessPool() ProcessPool {
	return getProcessPoolClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ProcessPool */
// An opaque token that you use to run multiple web views in a single process.
//
// A object represents a single process that WebKit uses to manage web content. To provide a more secure and stable experience, WebKit renders the content of web views in separate processes, rather than in your app’s process space. By default, WebKit gives each web view its own process space until it reaches an implementation-defined process limit. After that, web views with the same object share the same web content process. If your app creates multiple web views, assign the same object to web views that may safely share a process space. Instantiate an instance of this class and assign it to the property of each web view’s object.


// An opaque token that you use to run multiple web views in a single process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKProcessPool
type ProcessPool struct {
	objectivec.Object
}

// ProcessPoolFrom constructs a [ProcessPool] from an unsafe.Pointer.
//
// An opaque token that you use to run multiple web views in a single process.
func ProcessPoolFrom(ptr unsafe.Pointer) ProcessPool {
	return ProcessPool{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ProcessPool *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ProcessPool */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ProcessPool */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ProcessPool */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ProcessPool */

// The object that coordinates the processes the web view uses to render its web content and execute scripts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/processpool
func (p_ ProcessPool) ProcessPool() IWKProcessPool {
	rv := objc.Send[ProcessPool](p_.ID, objc.Sel("processPool"))
	return rv
}/* debug [instance_properties/getter]: processPool */


// The object that coordinates the processes the web view uses to render its web content and execute scripts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/processpool
func (p_ ProcessPool) SetProcessPool(value IWKProcessPool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setProcessPool:"), value)
}/* debug [instance_properties/setter]: processPool */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class WKProcessPool */



