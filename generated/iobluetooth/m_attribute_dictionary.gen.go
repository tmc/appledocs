// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mAttributeDictionary */


/* debug [class_header]: Header for mAttributeDictionary */
// The class instance for the [mAttributeDictionary] class.
var (
	MAttributeDictionaryClass     _mAttributeDictionaryClass
	MAttributeDictionaryClassOnce sync.Once
)

func getmAttributeDictionaryClass() _mAttributeDictionaryClass {
	MAttributeDictionaryClassOnce.Do(func() {
		MAttributeDictionaryClass = _mAttributeDictionaryClass{objc.GetClass("mAttributeDictionary")}
	})
	return MAttributeDictionaryClass
}

type _mAttributeDictionaryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mAttributeDictionary */
// An interface definition for the [mAttributeDictionary] class.
type ImAttributeDictionary interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mAttributeDictionary */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mAttributeDictionary */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mAttributeDictionary */
// Alloc allocates a new instance without initialization.
func (mc _mAttributeDictionaryClass) Alloc() mAttributeDictionary {
	rv := objc.Send[mAttributeDictionary](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mAttributeDictionaryClass) New() mAttributeDictionary {
	rv := objc.Send[mAttributeDictionary](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mAttributeDictionary) Init() mAttributeDictionary {
	rv := objc.Send[mAttributeDictionary](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mAttributeDictionary) Autorelease() mAttributeDictionary {
	rv := objc.Send[mAttributeDictionary](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmAttributeDictionary creates a new mAttributeDictionary instance.
func NewmAttributeDictionary() mAttributeDictionary {
	return getmAttributeDictionaryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mAttributeDictionary */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/mAttributeDictionary
type mAttributeDictionary struct {
	objectivec.Object
}

// mAttributeDictionaryFrom constructs a [mAttributeDictionary] from an unsafe.Pointer.
func mAttributeDictionaryFrom(ptr unsafe.Pointer) mAttributeDictionary {
	return mAttributeDictionary{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mAttributeDictionary *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mAttributeDictionary */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mAttributeDictionary */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mAttributeDictionary */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mAttributeDictionary */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mAttributeDictionary */



