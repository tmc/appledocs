// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class FSClient */


/* debug [class_header]: Header for FSClient */
// The class instance for the [FSClient] class.
var (
	FSClientClass     _FSClientClass
	FSClientClassOnce sync.Once
)

func getFSClientClass() _FSClientClass {
	FSClientClassOnce.Do(func() {
		FSClientClass = _FSClientClass{objc.GetClass("FSClient")}
	})
	return FSClientClass
}

type _FSClientClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FSClient */
// An interface definition for the [FSClient] class.
type IFSClient interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FSClient */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FSClient */
	// methods:
	FetchInstalledExtensionsWithCompletionHandler(completionHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FSClient */
// Alloc allocates a new instance without initialization.
func (fc _FSClientClass) Alloc() FSClient {
	rv := objc.Send[FSClient](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FSClientClass) New() FSClient {
	rv := objc.Send[FSClient](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSClient) Init() FSClient {
	rv := objc.Send[FSClient](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSClient) Autorelease() FSClient {
	rv := objc.Send[FSClient](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSClient creates a new FSClient instance.
func NewFSClient() FSClient {
	return getFSClientClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FSClient */
// An interface for apps and daemons to interact with FSKit.
//
// FSClient is the primary management interface for FSKit. Use this class to discover FSKit extensions installed on the system, including your own.


// An interface for apps and daemons to interact with FSKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSClient
type FSClient struct {
	objectivec.Object
}

// FSClientFrom constructs a [FSClient] from an unsafe.Pointer.
//
// An interface for apps and daemons to interact with FSKit.
func FSClientFrom(ptr unsafe.Pointer) FSClient {
	return FSClient{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FSClient *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FSClient */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FSClient */

// The shared instance of the FSKit client class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSClient/shared
func (fc _FSClientClass) SharedInstance() FSClient {
	rv := objc.Send[FSClient](objc.ID(fc.class), objc.Sel("sharedInstance"))
	return rv
}/* debug [class_properties_class/property]: sharedInstance */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FSClient */

// Asynchronously retrieves an list of installed file system modules.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSClient/fetchInstalledExtensions(completionHandler:)
func (f_ FSClient) FetchInstalledExtensionsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("fetchInstalledExtensionsWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: FetchInstalledExtensionsWithCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FSClient */

// The shared instance of the FSKit client class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSClient/shared
func (f_ FSClient) SharedInstance() IFSClient {
	rv := objc.Send[FSClient](f_.ID, objc.Sel("sharedInstance"))
	return rv
}/* debug [instance_properties/getter]: sharedInstance */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class FSClient */



