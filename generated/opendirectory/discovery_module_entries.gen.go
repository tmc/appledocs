// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class discoveryModuleEntries */


/* debug [class_header]: Header for discoveryModuleEntries */
// The class instance for the [discoveryModuleEntries] class.
var (
	DiscoveryModuleEntriesClass     _discoveryModuleEntriesClass
	DiscoveryModuleEntriesClassOnce sync.Once
)

func getdiscoveryModuleEntriesClass() _discoveryModuleEntriesClass {
	DiscoveryModuleEntriesClassOnce.Do(func() {
		DiscoveryModuleEntriesClass = _discoveryModuleEntriesClass{objc.GetClass("discoveryModuleEntries")}
	})
	return DiscoveryModuleEntriesClass
}

type _discoveryModuleEntriesClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for discoveryModuleEntries */
// An interface definition for the [discoveryModuleEntries] class.
type IdiscoveryModuleEntries interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for discoveryModuleEntries */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for discoveryModuleEntries */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for discoveryModuleEntries */
// Alloc allocates a new instance without initialization.
func (dc _discoveryModuleEntriesClass) Alloc() discoveryModuleEntries {
	rv := objc.Send[discoveryModuleEntries](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _discoveryModuleEntriesClass) New() discoveryModuleEntries {
	rv := objc.Send[discoveryModuleEntries](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ discoveryModuleEntries) Init() discoveryModuleEntries {
	rv := objc.Send[discoveryModuleEntries](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ discoveryModuleEntries) Autorelease() discoveryModuleEntries {
	rv := objc.Send[discoveryModuleEntries](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewdiscoveryModuleEntries creates a new discoveryModuleEntries instance.
func NewdiscoveryModuleEntries() discoveryModuleEntries {
	return getdiscoveryModuleEntriesClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for discoveryModuleEntries */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/discoveryModuleEntries-c.ivar
type discoveryModuleEntries struct {
	objectivec.Object
}

// discoveryModuleEntriesFrom constructs a [discoveryModuleEntries] from an unsafe.Pointer.
func discoveryModuleEntriesFrom(ptr unsafe.Pointer) discoveryModuleEntries {
	return discoveryModuleEntries{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for discoveryModuleEntries *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for discoveryModuleEntries */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for discoveryModuleEntries */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for discoveryModuleEntries */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for discoveryModuleEntries */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class discoveryModuleEntries */



