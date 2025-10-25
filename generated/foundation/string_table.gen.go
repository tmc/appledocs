// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class stringTable */


/* debug [class_header]: Header for stringTable */
// The class instance for the [stringTable] class.
var (
	StringTableClass     _stringTableClass
	StringTableClassOnce sync.Once
)

func getstringTableClass() _stringTableClass {
	StringTableClassOnce.Do(func() {
		StringTableClass = _stringTableClass{objc.GetClass("stringTable")}
	})
	return StringTableClass
}

type _stringTableClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for stringTable */
// An interface definition for the [stringTable] class.
type IstringTable interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for stringTable */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for stringTable */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for stringTable */
// Alloc allocates a new instance without initialization.
func (sc _stringTableClass) Alloc() stringTable {
	rv := objc.Send[stringTable](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _stringTableClass) New() stringTable {
	rv := objc.Send[stringTable](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ stringTable) Init() stringTable {
	rv := objc.Send[stringTable](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ stringTable) Autorelease() stringTable {
	rv := objc.Send[stringTable](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewstringTable creates a new stringTable instance.
func NewstringTable() stringTable {
	return getstringTableClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for stringTable */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArchiver/stringTable
type stringTable struct {
	objectivec.Object
}

// stringTableFrom constructs a [stringTable] from an unsafe.Pointer.
func stringTableFrom(ptr unsafe.Pointer) stringTable {
	return stringTable{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for stringTable *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for stringTable */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for stringTable */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for stringTable */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for stringTable */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class stringTable */



