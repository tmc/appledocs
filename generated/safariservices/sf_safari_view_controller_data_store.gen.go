// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SFSafariViewControllerDataStore */


/* debug [class_header]: Header for SFSafariViewControllerDataStore */
// The class instance for the [SFSafariViewControllerDataStore] class.
var (
	SFSafariViewControllerDataStoreClass     _SFSafariViewControllerDataStoreClass
	SFSafariViewControllerDataStoreClassOnce sync.Once
)

func getSFSafariViewControllerDataStoreClass() _SFSafariViewControllerDataStoreClass {
	SFSafariViewControllerDataStoreClassOnce.Do(func() {
		SFSafariViewControllerDataStoreClass = _SFSafariViewControllerDataStoreClass{objc.GetClass("SFSafariViewControllerDataStore")}
	})
	return SFSafariViewControllerDataStoreClass
}

type _SFSafariViewControllerDataStoreClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SFSafariViewControllerDataStore */
// An interface definition for the [SFSafariViewControllerDataStore] class.
type ISFSafariViewControllerDataStore interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SFSafariViewControllerDataStore */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SFSafariViewControllerDataStore */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SFSafariViewControllerDataStore */
// Alloc allocates a new instance without initialization.
func (sc _SFSafariViewControllerDataStoreClass) Alloc() SFSafariViewControllerDataStore {
	rv := objc.Send[SFSafariViewControllerDataStore](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SFSafariViewControllerDataStoreClass) New() SFSafariViewControllerDataStore {
	rv := objc.Send[SFSafariViewControllerDataStore](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFSafariViewControllerDataStore) Init() SFSafariViewControllerDataStore {
	rv := objc.Send[SFSafariViewControllerDataStore](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFSafariViewControllerDataStore) Autorelease() SFSafariViewControllerDataStore {
	rv := objc.Send[SFSafariViewControllerDataStore](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSafariViewControllerDataStore creates a new SFSafariViewControllerDataStore instance.
func NewSFSafariViewControllerDataStore() SFSafariViewControllerDataStore {
	return getSFSafariViewControllerDataStoreClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SFSafariViewControllerDataStore */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/DataStore
type SFSafariViewControllerDataStore struct {
	objectivec.Object
}

// SFSafariViewControllerDataStoreFrom constructs a [SFSafariViewControllerDataStore] from an unsafe.Pointer.
func SFSafariViewControllerDataStoreFrom(ptr unsafe.Pointer) SFSafariViewControllerDataStore {
	return SFSafariViewControllerDataStore{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SFSafariViewControllerDataStore *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SFSafariViewControllerDataStore */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SFSafariViewControllerDataStore */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/DataStore/default
func (sc _SFSafariViewControllerDataStoreClass) DefaultDataStore() SFSafariViewControllerDataStore {
	rv := objc.Send[SFSafariViewControllerDataStore](objc.ID(sc.class), objc.Sel("defaultDataStore"))
	return rv
}/* debug [class_properties_class/property]: defaultDataStore */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SFSafariViewControllerDataStore */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SFSafariViewControllerDataStore */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SFSafariViewControllerDataStore */


