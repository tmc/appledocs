// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class replacementTable */


/* debug [class_header]: Header for replacementTable */
// The class instance for the [replacementTable] class.
var (
	ReplacementTableClass     _replacementTableClass
	ReplacementTableClassOnce sync.Once
)

func getreplacementTableClass() _replacementTableClass {
	ReplacementTableClassOnce.Do(func() {
		ReplacementTableClass = _replacementTableClass{objc.GetClass("replacementTable")}
	})
	return ReplacementTableClass
}

type _replacementTableClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for replacementTable */
// An interface definition for the [replacementTable] class.
type IreplacementTable interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for replacementTable */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for replacementTable */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for replacementTable */
// Alloc allocates a new instance without initialization.
func (rc _replacementTableClass) Alloc() replacementTable {
	rv := objc.Send[replacementTable](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _replacementTableClass) New() replacementTable {
	rv := objc.Send[replacementTable](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ replacementTable) Init() replacementTable {
	rv := objc.Send[replacementTable](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ replacementTable) Autorelease() replacementTable {
	rv := objc.Send[replacementTable](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewreplacementTable creates a new replacementTable instance.
func NewreplacementTable() replacementTable {
	return getreplacementTableClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for replacementTable */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArchiver/replacementTable
type replacementTable struct {
	objectivec.Object
}

// replacementTableFrom constructs a [replacementTable] from an unsafe.Pointer.
func replacementTableFrom(ptr unsafe.Pointer) replacementTable {
	return replacementTable{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for replacementTable *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for replacementTable */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for replacementTable */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for replacementTable */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for replacementTable */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class replacementTable */



