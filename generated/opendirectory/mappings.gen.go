// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mappings */


/* debug [class_header]: Header for mappings */
// The class instance for the [mappings] class.
var (
	MappingsClass     _mappingsClass
	MappingsClassOnce sync.Once
)

func getmappingsClass() _mappingsClass {
	MappingsClassOnce.Do(func() {
		MappingsClass = _mappingsClass{objc.GetClass("mappings")}
	})
	return MappingsClass
}

type _mappingsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mappings */
// An interface definition for the [mappings] class.
type Imappings interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mappings */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mappings */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mappings */
// Alloc allocates a new instance without initialization.
func (mc _mappingsClass) Alloc() mappings {
	rv := objc.Send[mappings](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mappingsClass) New() mappings {
	rv := objc.Send[mappings](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mappings) Init() mappings {
	rv := objc.Send[mappings](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mappings) Autorelease() mappings {
	rv := objc.Send[mappings](m_.ID, objc.Sel("autorelease"))
	return rv
}

// Newmappings creates a new mappings instance.
func Newmappings() mappings {
	return getmappingsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mappings */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/mappings-c.ivar
type mappings struct {
	objectivec.Object
}

// mappingsFrom constructs a [mappings] from an unsafe.Pointer.
func mappingsFrom(ptr unsafe.Pointer) mappings {
	return mappings{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mappings *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mappings */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mappings */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mappings */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mappings */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mappings */



