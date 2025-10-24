// Code generated from Apple documentation for CoreAudioKit. DO NOT EDIT.

package coreaudiokit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/audiotoolbox"
)

/* debug [class.gen.go]: Generating class AUGenericViewController */


/* debug [class_header]: Header for AUGenericViewController */
// The class instance for the [GenericViewController] class.
var (
	GenericViewControllerClass     _GenericViewControllerClass
	GenericViewControllerClassOnce sync.Once
)

func getGenericViewControllerClass() _GenericViewControllerClass {
	GenericViewControllerClassOnce.Do(func() {
		GenericViewControllerClass = _GenericViewControllerClass{objc.GetClass("AUGenericViewController")}
	})
	return GenericViewControllerClass
}

type _GenericViewControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GenericViewController */
// An interface definition for the [GenericViewController] class.
type IGenericViewController interface {
	IViewController
	
/* debug [class_interface_properties]: Properties for GenericViewController */
	// properties:
	AuAudioUnit() audiotoolbox.AudioUnit
	SetAuAudioUnit(value audiotoolbox.AudioUnit)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GenericViewController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GenericViewController */
// Alloc allocates a new instance without initialization.
func (gc _GenericViewControllerClass) Alloc() GenericViewController {
	rv := objc.Send[GenericViewController](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GenericViewControllerClass) New() GenericViewController {
	rv := objc.Send[GenericViewController](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GenericViewController) Init() GenericViewController {
	rv := objc.Send[GenericViewController](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GenericViewController) Autorelease() GenericViewController {
	rv := objc.Send[GenericViewController](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGenericViewController creates a new GenericViewController instance.
func NewGenericViewController() GenericViewController {
	return getGenericViewControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GenericViewController */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/AUGenericViewController
type GenericViewController struct {
	ViewController
}

// GenericViewControllerFrom constructs a [GenericViewController] from an unsafe.Pointer.
func GenericViewControllerFrom(ptr unsafe.Pointer) GenericViewController {
	return GenericViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GenericViewController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GenericViewController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GenericViewController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GenericViewController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GenericViewController */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/AUGenericViewController/auAudioUnit
func (g_ GenericViewController) AuAudioUnit() audiotoolbox.AudioUnit {
	rv := objc.Send[audiotoolbox.AudioUnit](g_.ID, objc.Sel("auAudioUnit"))
	return rv
}/* debug [instance_properties/getter]: auAudioUnit */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/AUGenericViewController/auAudioUnit
func (g_ GenericViewController) SetAuAudioUnit(value audiotoolbox.AudioUnit) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setAuAudioUnit:"), value)
}/* debug [instance_properties/setter]: auAudioUnit */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AUGenericViewController */



