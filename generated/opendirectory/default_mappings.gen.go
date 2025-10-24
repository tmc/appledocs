// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class defaultMappings */


/* debug [class_header]: Header for defaultMappings */
// The class instance for the [defaultMappings] class.
var (
	DefaultMappingsClass     _defaultMappingsClass
	DefaultMappingsClassOnce sync.Once
)

func getdefaultMappingsClass() _defaultMappingsClass {
	DefaultMappingsClassOnce.Do(func() {
		DefaultMappingsClass = _defaultMappingsClass{objc.GetClass("defaultMappings")}
	})
	return DefaultMappingsClass
}

type _defaultMappingsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for defaultMappings */
// An interface definition for the [defaultMappings] class.
type IdefaultMappings interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for defaultMappings */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for defaultMappings */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for defaultMappings */
// Alloc allocates a new instance without initialization.
func (dc _defaultMappingsClass) Alloc() defaultMappings {
	rv := objc.Send[defaultMappings](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _defaultMappingsClass) New() defaultMappings {
	rv := objc.Send[defaultMappings](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ defaultMappings) Init() defaultMappings {
	rv := objc.Send[defaultMappings](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ defaultMappings) Autorelease() defaultMappings {
	rv := objc.Send[defaultMappings](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewdefaultMappings creates a new defaultMappings instance.
func NewdefaultMappings() defaultMappings {
	return getdefaultMappingsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for defaultMappings */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/defaultMappings-c.ivar
type defaultMappings struct {
	objectivec.Object
}

// defaultMappingsFrom constructs a [defaultMappings] from an unsafe.Pointer.
func defaultMappingsFrom(ptr unsafe.Pointer) defaultMappings {
	return defaultMappings{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for defaultMappings *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for defaultMappings */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for defaultMappings */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for defaultMappings */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for defaultMappings */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class defaultMappings */



