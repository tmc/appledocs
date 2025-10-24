// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSKeyValueSharedObservers */


/* debug [class_header]: Header for NSKeyValueSharedObservers */
// The class instance for the [KeyValueSharedObservers] class.
var (
	KeyValueSharedObserversClass     _KeyValueSharedObserversClass
	KeyValueSharedObserversClassOnce sync.Once
)

func getKeyValueSharedObserversClass() _KeyValueSharedObserversClass {
	KeyValueSharedObserversClassOnce.Do(func() {
		KeyValueSharedObserversClass = _KeyValueSharedObserversClass{objc.GetClass("NSKeyValueSharedObservers")}
	})
	return KeyValueSharedObserversClass
}

type _KeyValueSharedObserversClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for KeyValueSharedObservers */
// An interface definition for the [KeyValueSharedObservers] class.
type IKeyValueSharedObservers interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for KeyValueSharedObservers */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for KeyValueSharedObservers */
	// methods:
	AddSharedObserverForKeyOptionsContext(observer objc.IObject /* cross-framework: NSObject */, key IString, options uint, context objectivec.IObject)
	Snapshot() IKeyValueSharedObserversSnapshot
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for KeyValueSharedObservers */
// Alloc allocates a new instance without initialization.
func (kc _KeyValueSharedObserversClass) Alloc() KeyValueSharedObservers {
	rv := objc.Send[KeyValueSharedObservers](objc.ID(kc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (kc _KeyValueSharedObserversClass) New() KeyValueSharedObservers {
	rv := objc.Send[KeyValueSharedObservers](objc.ID(kc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (k_ KeyValueSharedObservers) Init() KeyValueSharedObservers {
	rv := objc.Send[KeyValueSharedObservers](k_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (k_ KeyValueSharedObservers) Autorelease() KeyValueSharedObservers {
	rv := objc.Send[KeyValueSharedObservers](k_.ID, objc.Sel("autorelease"))
	return rv
}

// NewKeyValueSharedObservers creates a new KeyValueSharedObservers instance.
func NewKeyValueSharedObservers() KeyValueSharedObservers {
	return getKeyValueSharedObserversClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for KeyValueSharedObservers */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueSharedObservers
type KeyValueSharedObservers struct {
	objectivec.Object
}

// KeyValueSharedObserversFrom constructs a [KeyValueSharedObservers] from an unsafe.Pointer.
func KeyValueSharedObserversFrom(ptr unsafe.Pointer) KeyValueSharedObservers {
	return KeyValueSharedObservers{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for KeyValueSharedObservers */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueSharedObservers/init(observableClass:)
func NewKeyValueSharedObserversWithObservableClass(observableClass objc.Class) KeyValueSharedObservers {
	instance := getKeyValueSharedObserversClass().Alloc()
	rv := objc.Send[KeyValueSharedObservers](instance.ID, objc.Sel("initWithObservableClass:"), observableClass)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewKeyValueSharedObserversWithObservableClass */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for KeyValueSharedObservers */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for KeyValueSharedObservers */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for KeyValueSharedObservers */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueSharedObservers/addSharedObserver(_:forKey:options:context:)
func (k_ KeyValueSharedObservers) AddSharedObserverForKeyOptionsContext(observer objc.IObject /* cross-framework: NSObject */, key IString, options uint, context objectivec.IObject) {
	objc.Send[objc.ID](k_.ID, objc.Sel("addSharedObserver:forKey:options:context:"), observer, key, options, context)
}/* debug [instance_methods/method]: AddSharedObserverForKeyOptionsContext */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueSharedObservers/snapshot()
func (k_ KeyValueSharedObservers) Snapshot() IKeyValueSharedObserversSnapshot {
	rv := objc.Send[KeyValueSharedObserversSnapshot](k_.ID, objc.Sel("snapshot"))
	return rv
}/* debug [instance_methods/method]: Snapshot */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for KeyValueSharedObservers */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSKeyValueSharedObservers */


