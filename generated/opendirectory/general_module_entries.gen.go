// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class generalModuleEntries */


/* debug [class_header]: Header for generalModuleEntries */
// The class instance for the [generalModuleEntries] class.
var (
	GeneralModuleEntriesClass     _generalModuleEntriesClass
	GeneralModuleEntriesClassOnce sync.Once
)

func getgeneralModuleEntriesClass() _generalModuleEntriesClass {
	GeneralModuleEntriesClassOnce.Do(func() {
		GeneralModuleEntriesClass = _generalModuleEntriesClass{objc.GetClass("generalModuleEntries")}
	})
	return GeneralModuleEntriesClass
}

type _generalModuleEntriesClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for generalModuleEntries */
// An interface definition for the [generalModuleEntries] class.
type IgeneralModuleEntries interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for generalModuleEntries */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for generalModuleEntries */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for generalModuleEntries */
// Alloc allocates a new instance without initialization.
func (gc _generalModuleEntriesClass) Alloc() generalModuleEntries {
	rv := objc.Send[generalModuleEntries](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _generalModuleEntriesClass) New() generalModuleEntries {
	rv := objc.Send[generalModuleEntries](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ generalModuleEntries) Init() generalModuleEntries {
	rv := objc.Send[generalModuleEntries](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ generalModuleEntries) Autorelease() generalModuleEntries {
	rv := objc.Send[generalModuleEntries](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewgeneralModuleEntries creates a new generalModuleEntries instance.
func NewgeneralModuleEntries() generalModuleEntries {
	return getgeneralModuleEntriesClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for generalModuleEntries */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/generalModuleEntries-c.ivar
type generalModuleEntries struct {
	objectivec.Object
}

// generalModuleEntriesFrom constructs a [generalModuleEntries] from an unsafe.Pointer.
func generalModuleEntriesFrom(ptr unsafe.Pointer) generalModuleEntries {
	return generalModuleEntries{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for generalModuleEntries *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for generalModuleEntries */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for generalModuleEntries */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for generalModuleEntries */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for generalModuleEntries */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class generalModuleEntries */



