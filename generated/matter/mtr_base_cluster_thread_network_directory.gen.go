// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterThreadNetworkDirectory */


/* debug [class_header]: Header for MTRBaseClusterThreadNetworkDirectory */
// The class instance for the [MTRBaseClusterThreadNetworkDirectory] class.
var (
	MTRBaseClusterThreadNetworkDirectoryClass     _MTRBaseClusterThreadNetworkDirectoryClass
	MTRBaseClusterThreadNetworkDirectoryClassOnce sync.Once
)

func getMTRBaseClusterThreadNetworkDirectoryClass() _MTRBaseClusterThreadNetworkDirectoryClass {
	MTRBaseClusterThreadNetworkDirectoryClassOnce.Do(func() {
		MTRBaseClusterThreadNetworkDirectoryClass = _MTRBaseClusterThreadNetworkDirectoryClass{objc.GetClass("MTRBaseClusterThreadNetworkDirectory")}
	})
	return MTRBaseClusterThreadNetworkDirectoryClass
}

type _MTRBaseClusterThreadNetworkDirectoryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterThreadNetworkDirectory */
// An interface definition for the [MTRBaseClusterThreadNetworkDirectory] class.
type IMTRBaseClusterThreadNetworkDirectory interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterThreadNetworkDirectory */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterThreadNetworkDirectory */
	// methods:
	WriteAttributePreferredExtendedPanIDWithValueCompletion(value objc.IObject /* cross-framework: NSData */, completion unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterThreadNetworkDirectory */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterThreadNetworkDirectoryClass) Alloc() MTRBaseClusterThreadNetworkDirectory {
	rv := objc.Send[MTRBaseClusterThreadNetworkDirectory](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterThreadNetworkDirectoryClass) New() MTRBaseClusterThreadNetworkDirectory {
	rv := objc.Send[MTRBaseClusterThreadNetworkDirectory](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterThreadNetworkDirectory) Init() MTRBaseClusterThreadNetworkDirectory {
	rv := objc.Send[MTRBaseClusterThreadNetworkDirectory](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterThreadNetworkDirectory) Autorelease() MTRBaseClusterThreadNetworkDirectory {
	rv := objc.Send[MTRBaseClusterThreadNetworkDirectory](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterThreadNetworkDirectory creates a new MTRBaseClusterThreadNetworkDirectory instance.
func NewMTRBaseClusterThreadNetworkDirectory() MTRBaseClusterThreadNetworkDirectory {
	return getMTRBaseClusterThreadNetworkDirectoryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterThreadNetworkDirectory */
// Cluster Thread Network Directory
//
// Manages the names and credentials of Thread networks visible to the user.


// Cluster Thread Network Directory
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadNetworkDirectory
type MTRBaseClusterThreadNetworkDirectory struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterThreadNetworkDirectoryFrom constructs a [MTRBaseClusterThreadNetworkDirectory] from an unsafe.Pointer.
//
// Cluster Thread Network Directory
func MTRBaseClusterThreadNetworkDirectoryFrom(ptr unsafe.Pointer) MTRBaseClusterThreadNetworkDirectory {
	return MTRBaseClusterThreadNetworkDirectory{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterThreadNetworkDirectory *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterThreadNetworkDirectory */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterThreadNetworkDirectory */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterThreadNetworkDirectory */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadNetworkDirectory/writeAttributePreferredExtendedPanID(withValue:completion:)
func (m_ MTRBaseClusterThreadNetworkDirectory) WriteAttributePreferredExtendedPanIDWithValueCompletion(value objc.IObject /* cross-framework: NSData */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributePreferredExtendedPanIDWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributePreferredExtendedPanIDWithValueCompletion */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterThreadNetworkDirectory */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterThreadNetworkDirectory */



