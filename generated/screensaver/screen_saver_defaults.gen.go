// Code generated from Apple documentation for ScreenSaver. DO NOT EDIT.

package screensaver

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class ScreenSaverDefaults */


/* debug [class_header]: Header for ScreenSaverDefaults */
// The class instance for the [ScreenSaverDefaults] class.
var (
	ScreenSaverDefaultsClass     _ScreenSaverDefaultsClass
	ScreenSaverDefaultsClassOnce sync.Once
)

func getScreenSaverDefaultsClass() _ScreenSaverDefaultsClass {
	ScreenSaverDefaultsClassOnce.Do(func() {
		ScreenSaverDefaultsClass = _ScreenSaverDefaultsClass{objc.GetClass("ScreenSaverDefaults")}
	})
	return ScreenSaverDefaultsClass
}

type _ScreenSaverDefaultsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ScreenSaverDefaults */
// An interface definition for the [ScreenSaverDefaults] class.
type IScreenSaverDefaults interface {
	foundation.IUserDefaults
	
/* debug [class_interface_properties]: Properties for ScreenSaverDefaults */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ScreenSaverDefaults */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ScreenSaverDefaults */
// Alloc allocates a new instance without initialization.
func (sc _ScreenSaverDefaultsClass) Alloc() ScreenSaverDefaults {
	rv := objc.Send[ScreenSaverDefaults](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _ScreenSaverDefaultsClass) New() ScreenSaverDefaults {
	rv := objc.Send[ScreenSaverDefaults](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScreenSaverDefaults) Init() ScreenSaverDefaults {
	rv := objc.Send[ScreenSaverDefaults](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScreenSaverDefaults) Autorelease() ScreenSaverDefaults {
	rv := objc.Send[ScreenSaverDefaults](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScreenSaverDefaults creates a new ScreenSaverDefaults instance.
func NewScreenSaverDefaults() ScreenSaverDefaults {
	return getScreenSaverDefaultsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ScreenSaverDefaults */
// A class that defines a set of methods for saving and restoring user defaults for screen savers.
//
// gives you access to preference values you need to configure your screen saver. Because multiple apps can load a screen saver, you can’t use the standard object to store preferences. Instead, instantiate this class using the method, which takes your screen saver’s bundle identifier as a parameter. The resulting object gives you a way to store your preference values and associate them only with your screen saver. Use the inherited methods to load, store, or modify values.


// A class that defines a set of methods for saving and restoring user defaults for screen savers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenSaver/ScreenSaverDefaults
type ScreenSaverDefaults struct {
	foundation.UserDefaults
}

// ScreenSaverDefaultsFrom constructs a [ScreenSaverDefaults] from an unsafe.Pointer.
//
// A class that defines a set of methods for saving and restoring user defaults for screen savers.
func ScreenSaverDefaultsFrom(ptr unsafe.Pointer) ScreenSaverDefaults {
	return ScreenSaverDefaults{
		UserDefaults: foundation.UserDefaultsFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ScreenSaverDefaults */

// Returns a screen saver defaults instance that reads and writes defaults for the specified module.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenSaver/ScreenSaverDefaults/init(forModuleWithName:)
func NewScreenSaverDefaultsForModuleWithName(inModuleName objc.IObject /* cross-framework: NSString */) ScreenSaverDefaults {
	rv := objc.Send[ScreenSaverDefaults](objc.ID(getScreenSaverDefaultsClass().class), objc.Sel("defaultsForModuleWithName:"), inModuleName)
	return rv
}/* debug [class_init_methods/constructor]: NewScreenSaverDefaultsForModuleWithName */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ScreenSaverDefaults */

// Returns a screen saver defaults instance that reads and writes defaults for the specified module.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenSaver/ScreenSaverDefaults/init(forModuleWithName:)
func (sc _ScreenSaverDefaultsClass) DefaultsForModuleWithName(inModuleName objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("defaultsForModuleWithName:"), inModuleName)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DefaultsForModuleWithName) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ScreenSaverDefaults */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ScreenSaverDefaults */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ScreenSaverDefaults */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ScreenSaverDefaults */


