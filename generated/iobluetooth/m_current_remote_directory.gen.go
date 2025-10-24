// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mCurrentRemoteDirectory */


/* debug [class_header]: Header for mCurrentRemoteDirectory */
// The class instance for the [mCurrentRemoteDirectory] class.
var (
	MCurrentRemoteDirectoryClass     _mCurrentRemoteDirectoryClass
	MCurrentRemoteDirectoryClassOnce sync.Once
)

func getmCurrentRemoteDirectoryClass() _mCurrentRemoteDirectoryClass {
	MCurrentRemoteDirectoryClassOnce.Do(func() {
		MCurrentRemoteDirectoryClass = _mCurrentRemoteDirectoryClass{objc.GetClass("mCurrentRemoteDirectory")}
	})
	return MCurrentRemoteDirectoryClass
}

type _mCurrentRemoteDirectoryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mCurrentRemoteDirectory */
// An interface definition for the [mCurrentRemoteDirectory] class.
type ImCurrentRemoteDirectory interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mCurrentRemoteDirectory */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mCurrentRemoteDirectory */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mCurrentRemoteDirectory */
// Alloc allocates a new instance without initialization.
func (mc _mCurrentRemoteDirectoryClass) Alloc() mCurrentRemoteDirectory {
	rv := objc.Send[mCurrentRemoteDirectory](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mCurrentRemoteDirectoryClass) New() mCurrentRemoteDirectory {
	rv := objc.Send[mCurrentRemoteDirectory](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mCurrentRemoteDirectory) Init() mCurrentRemoteDirectory {
	rv := objc.Send[mCurrentRemoteDirectory](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mCurrentRemoteDirectory) Autorelease() mCurrentRemoteDirectory {
	rv := objc.Send[mCurrentRemoteDirectory](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmCurrentRemoteDirectory creates a new mCurrentRemoteDirectory instance.
func NewmCurrentRemoteDirectory() mCurrentRemoteDirectory {
	return getmCurrentRemoteDirectoryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mCurrentRemoteDirectory */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mCurrentRemoteDirectory
type mCurrentRemoteDirectory struct {
	objectivec.Object
}

// mCurrentRemoteDirectoryFrom constructs a [mCurrentRemoteDirectory] from an unsafe.Pointer.
func mCurrentRemoteDirectoryFrom(ptr unsafe.Pointer) mCurrentRemoteDirectory {
	return mCurrentRemoteDirectory{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mCurrentRemoteDirectory *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mCurrentRemoteDirectory */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mCurrentRemoteDirectory */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mCurrentRemoteDirectory */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mCurrentRemoteDirectory */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mCurrentRemoteDirectory */



