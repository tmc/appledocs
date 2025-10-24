// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class FSUnaryFileSystem */


/* debug [class_header]: Header for FSUnaryFileSystem */
// The class instance for the [FSUnaryFileSystem] class.
var (
	FSUnaryFileSystemClass     _FSUnaryFileSystemClass
	FSUnaryFileSystemClassOnce sync.Once
)

func getFSUnaryFileSystemClass() _FSUnaryFileSystemClass {
	FSUnaryFileSystemClassOnce.Do(func() {
		FSUnaryFileSystemClass = _FSUnaryFileSystemClass{objc.GetClass("FSUnaryFileSystem")}
	})
	return FSUnaryFileSystemClass
}

type _FSUnaryFileSystemClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FSUnaryFileSystem */
// An interface definition for the [FSUnaryFileSystem] class.
type IFSUnaryFileSystem interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FSUnaryFileSystem */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FSUnaryFileSystem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FSUnaryFileSystem */
// Alloc allocates a new instance without initialization.
func (fc _FSUnaryFileSystemClass) Alloc() FSUnaryFileSystem {
	rv := objc.Send[FSUnaryFileSystem](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FSUnaryFileSystemClass) New() FSUnaryFileSystem {
	rv := objc.Send[FSUnaryFileSystem](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSUnaryFileSystem) Init() FSUnaryFileSystem {
	rv := objc.Send[FSUnaryFileSystem](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSUnaryFileSystem) Autorelease() FSUnaryFileSystem {
	rv := objc.Send[FSUnaryFileSystem](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSUnaryFileSystem creates a new FSUnaryFileSystem instance.
func NewFSUnaryFileSystem() FSUnaryFileSystem {
	return getFSUnaryFileSystemClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FSUnaryFileSystem */
// An abstract base class for implementing a minimal file system.
//
// is a simplified file system, which works with one and presents it as one . The one volume and its container have a shared state and lifetime, a more constrained life cycle than the design flow. Implement your app extension by providing a subclass of as a delegate object. Your delegate also needs to implement the protocol so that it can load resources.


// An abstract base class for implementing a minimal file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSUnaryFileSystem
type FSUnaryFileSystem struct {
	objectivec.Object
}

// FSUnaryFileSystemFrom constructs a [FSUnaryFileSystem] from an unsafe.Pointer.
//
// An abstract base class for implementing a minimal file system.
func FSUnaryFileSystemFrom(ptr unsafe.Pointer) FSUnaryFileSystem {
	return FSUnaryFileSystem{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FSUnaryFileSystem *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FSUnaryFileSystem */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FSUnaryFileSystem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FSUnaryFileSystem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FSUnaryFileSystem */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class FSUnaryFileSystem */



