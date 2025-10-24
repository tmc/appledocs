// Code generated from Apple documentation for IOBluetoothUI. DO NOT EDIT.

package iobluetoothui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mPasskeyCharacters */


/* debug [class_header]: Header for mPasskeyCharacters */
// The class instance for the [mPasskeyCharacters] class.
var (
	MPasskeyCharactersClass     _mPasskeyCharactersClass
	MPasskeyCharactersClassOnce sync.Once
)

func getmPasskeyCharactersClass() _mPasskeyCharactersClass {
	MPasskeyCharactersClassOnce.Do(func() {
		MPasskeyCharactersClass = _mPasskeyCharactersClass{objc.GetClass("mPasskeyCharacters")}
	})
	return MPasskeyCharactersClass
}

type _mPasskeyCharactersClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mPasskeyCharacters */
// An interface definition for the [mPasskeyCharacters] class.
type ImPasskeyCharacters interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mPasskeyCharacters */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mPasskeyCharacters */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mPasskeyCharacters */
// Alloc allocates a new instance without initialization.
func (mc _mPasskeyCharactersClass) Alloc() mPasskeyCharacters {
	rv := objc.Send[mPasskeyCharacters](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mPasskeyCharactersClass) New() mPasskeyCharacters {
	rv := objc.Send[mPasskeyCharacters](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mPasskeyCharacters) Init() mPasskeyCharacters {
	rv := objc.Send[mPasskeyCharacters](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mPasskeyCharacters) Autorelease() mPasskeyCharacters {
	rv := objc.Send[mPasskeyCharacters](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmPasskeyCharacters creates a new mPasskeyCharacters instance.
func NewmPasskeyCharacters() mPasskeyCharacters {
	return getmPasskeyCharactersClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mPasskeyCharacters */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay/mPasskeyCharacters
type mPasskeyCharacters struct {
	objectivec.Object
}

// mPasskeyCharactersFrom constructs a [mPasskeyCharacters] from an unsafe.Pointer.
func mPasskeyCharactersFrom(ptr unsafe.Pointer) mPasskeyCharacters {
	return mPasskeyCharacters{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mPasskeyCharacters *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mPasskeyCharacters */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mPasskeyCharacters */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mPasskeyCharacters */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mPasskeyCharacters */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mPasskeyCharacters */



