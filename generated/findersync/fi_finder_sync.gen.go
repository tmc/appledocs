// Code generated from Apple documentation for FinderSync. DO NOT EDIT.

package findersync

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class FIFinderSync */


/* debug [class_header]: Header for FIFinderSync */
// The class instance for the [FIFinderSync] class.
var (
	FIFinderSyncClass     _FIFinderSyncClass
	FIFinderSyncClassOnce sync.Once
)

func getFIFinderSyncClass() _FIFinderSyncClass {
	FIFinderSyncClassOnce.Do(func() {
		FIFinderSyncClass = _FIFinderSyncClass{objc.GetClass("FIFinderSync")}
	})
	return FIFinderSyncClass
}

type _FIFinderSyncClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FIFinderSync */
// An interface definition for the [FIFinderSync] class.
type IFIFinderSync interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FIFinderSync */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FIFinderSync */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FIFinderSync */
// Alloc allocates a new instance without initialization.
func (fc _FIFinderSyncClass) Alloc() FIFinderSync {
	rv := objc.Send[FIFinderSync](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FIFinderSyncClass) New() FIFinderSync {
	rv := objc.Send[FIFinderSync](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FIFinderSync) Init() FIFinderSync {
	rv := objc.Send[FIFinderSync](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FIFinderSync) Autorelease() FIFinderSync {
	rv := objc.Send[FIFinderSync](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFIFinderSync creates a new FIFinderSync instance.
func NewFIFinderSync() FIFinderSync {
	return getFIFinderSyncClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FIFinderSync */
// A type to subclass to add badges, custom shortcut menus, and toolbar buttons to the Finder.
//
// Subclass the FIFinderSync class when you want to customize the appearance of the Finder. Although the FIFinderSync class doesn’t provide any developer accessible API, it does adopt the protocol. This protocol declares methods you can implement to modify the appearance of the Finder. For more information on these methods, see . To learn more about creating a Finder Sync extension, see in .


// A type to subclass to add badges, custom shortcut menus, and toolbar buttons to the Finder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FinderSync/FIFinderSync-swift.class
type FIFinderSync struct {
	objectivec.Object
}

// FIFinderSyncFrom constructs a [FIFinderSync] from an unsafe.Pointer.
//
// A type to subclass to add badges, custom shortcut menus, and toolbar buttons to the Finder.
func FIFinderSyncFrom(ptr unsafe.Pointer) FIFinderSync {
	return FIFinderSync{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FIFinderSync *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FIFinderSync */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FIFinderSync */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FIFinderSync */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FIFinderSync */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class FIFinderSync */



