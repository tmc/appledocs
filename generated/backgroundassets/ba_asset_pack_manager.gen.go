// Code generated from Apple documentation for BackgroundAssets. DO NOT EDIT.

package backgroundassets

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class BAAssetPackManager */


/* debug [class_header]: Header for BAAssetPackManager */
// The class instance for the [BAAssetPackManager] class.
var (
	BAAssetPackManagerClass     _BAAssetPackManagerClass
	BAAssetPackManagerClassOnce sync.Once
)

func getBAAssetPackManagerClass() _BAAssetPackManagerClass {
	BAAssetPackManagerClassOnce.Do(func() {
		BAAssetPackManagerClass = _BAAssetPackManagerClass{objc.GetClass("BAAssetPackManager")}
	})
	return BAAssetPackManagerClass
}

type _BAAssetPackManagerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BAAssetPackManager */
// An interface definition for the [BAAssetPackManager] class.
type IBAAssetPackManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for BAAssetPackManager */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BAAssetPackManager */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BAAssetPackManager */
// Alloc allocates a new instance without initialization.
func (bc _BAAssetPackManagerClass) Alloc() BAAssetPackManager {
	rv := objc.Send[BAAssetPackManager](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BAAssetPackManagerClass) New() BAAssetPackManager {
	rv := objc.Send[BAAssetPackManager](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BAAssetPackManager) Init() BAAssetPackManager {
	rv := objc.Send[BAAssetPackManager](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BAAssetPackManager) Autorelease() BAAssetPackManager {
	rv := objc.Send[BAAssetPackManager](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBAAssetPackManager creates a new BAAssetPackManager instance.
func NewBAAssetPackManager() BAAssetPackManager {
	return getBAAssetPackManagerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BAAssetPackManager */
// A class that manages asset packs.
//
// The first time that your code refers to the shared manager, Background Assets considers that your app is opting into automatic system management of your asset packs.


// A class that manages asset packs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPackManager
type BAAssetPackManager struct {
	objectivec.Object
}

// BAAssetPackManagerFrom constructs a [BAAssetPackManager] from an unsafe.Pointer.
//
// A class that manages asset packs.
func BAAssetPackManagerFrom(ptr unsafe.Pointer) BAAssetPackManager {
	return BAAssetPackManager{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BAAssetPackManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BAAssetPackManager */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BAAssetPackManager */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BAAssetPackManager */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BAAssetPackManager */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class BAAssetPackManager */



