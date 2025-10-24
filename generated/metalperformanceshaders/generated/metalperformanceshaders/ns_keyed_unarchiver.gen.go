// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSKeyedUnarchiver */


/* debug [class_header]: Header for NSKeyedUnarchiver */
// The class instance for the [KeyedUnarchiver] class.
var (
	KeyedUnarchiverClass     _KeyedUnarchiverClass
	KeyedUnarchiverClassOnce sync.Once
)

func getKeyedUnarchiverClass() _KeyedUnarchiverClass {
	KeyedUnarchiverClassOnce.Do(func() {
		KeyedUnarchiverClass = _KeyedUnarchiverClass{objc.GetClass("NSKeyedUnarchiver")}
	})
	return KeyedUnarchiverClass
}

type _KeyedUnarchiverClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for KeyedUnarchiver */
// An interface definition for the [KeyedUnarchiver] class.
type IKeyedUnarchiver interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for KeyedUnarchiver */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for KeyedUnarchiver */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for KeyedUnarchiver */
// Alloc allocates a new instance without initialization.
func (kc _KeyedUnarchiverClass) Alloc() KeyedUnarchiver {
	rv := objc.Send[KeyedUnarchiver](objc.ID(kc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (kc _KeyedUnarchiverClass) New() KeyedUnarchiver {
	rv := objc.Send[KeyedUnarchiver](objc.ID(kc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (k_ KeyedUnarchiver) Init() KeyedUnarchiver {
	rv := objc.Send[KeyedUnarchiver](k_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (k_ KeyedUnarchiver) Autorelease() KeyedUnarchiver {
	rv := objc.Send[KeyedUnarchiver](k_.ID, objc.Sel("autorelease"))
	return rv
}

// NewKeyedUnarchiver creates a new KeyedUnarchiver instance.
func NewKeyedUnarchiver() KeyedUnarchiver {
	return getKeyedUnarchiverClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for KeyedUnarchiver */
// A parent class referenced by other MetalPerformanceShaders classes.


// A parent class referenced by other MetalPerformanceShaders classes. [Full Topic]
type KeyedUnarchiver struct {
	objectivec.Object
}

// KeyedUnarchiverFrom constructs a [KeyedUnarchiver] from an unsafe.Pointer.
//
// A parent class referenced by other MetalPerformanceShaders classes.
func KeyedUnarchiverFrom(ptr unsafe.Pointer) KeyedUnarchiver {
	return KeyedUnarchiver{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for KeyedUnarchiver *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for KeyedUnarchiver */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for KeyedUnarchiver */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for KeyedUnarchiver */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for KeyedUnarchiver */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSKeyedUnarchiver */



