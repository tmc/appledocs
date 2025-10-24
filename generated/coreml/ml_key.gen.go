// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLKey */


/* debug [class_header]: Header for MLKey */
// The class instance for the [Key] class.
var (
	KeyClass     _KeyClass
	KeyClassOnce sync.Once
)

func getKeyClass() _KeyClass {
	KeyClassOnce.Do(func() {
		KeyClass = _KeyClass{objc.GetClass("MLKey")}
	})
	return KeyClass
}

type _KeyClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Key */
// An interface definition for the [Key] class.
type IKey interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Key */
	// properties:
	Name() objc.IObject /* cross-framework: NSString */
	Scope() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Key */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Key */
// Alloc allocates a new instance without initialization.
func (kc _KeyClass) Alloc() Key {
	rv := objc.Send[Key](objc.ID(kc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (kc _KeyClass) New() Key {
	rv := objc.Send[Key](objc.ID(kc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (k_ Key) Init() Key {
	rv := objc.Send[Key](k_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (k_ Key) Autorelease() Key {
	rv := objc.Send[Key](k_.ID, objc.Sel("autorelease"))
	return rv
}

// NewKey creates a new Key instance.
func NewKey() Key {
	return getKeyClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Key */
// An abstract base class for machine learning key types.
//
// You don’t create use this class directly. Instead, use a class that inherits from this one, such as or .


// An abstract base class for machine learning key types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLKey
type Key struct {
	objectivec.Object
}

// KeyFrom constructs a [Key] from an unsafe.Pointer.
//
// An abstract base class for machine learning key types.
func KeyFrom(ptr unsafe.Pointer) Key {
	return Key{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Key *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Key */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Key */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Key */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Key */

// The name of the machine learning key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLKey/name
func (k_ Key) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](k_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The applicable scope of the machine learning key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLKey/scope
func (k_ Key) Scope() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](k_.ID, objc.Sel("scope"))
	return rv
}/* debug [instance_properties/getter]: scope */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLKey */



