// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRClusterThreadNetworkDirectory */


/* debug [class_header]: Header for MTRClusterThreadNetworkDirectory */
// The class instance for the [MTRClusterThreadNetworkDirectory] class.
var (
	MTRClusterThreadNetworkDirectoryClass     _MTRClusterThreadNetworkDirectoryClass
	MTRClusterThreadNetworkDirectoryClassOnce sync.Once
)

func getMTRClusterThreadNetworkDirectoryClass() _MTRClusterThreadNetworkDirectoryClass {
	MTRClusterThreadNetworkDirectoryClassOnce.Do(func() {
		MTRClusterThreadNetworkDirectoryClass = _MTRClusterThreadNetworkDirectoryClass{objc.GetClass("MTRClusterThreadNetworkDirectory")}
	})
	return MTRClusterThreadNetworkDirectoryClass
}

type _MTRClusterThreadNetworkDirectoryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRClusterThreadNetworkDirectory */
// An interface definition for the [MTRClusterThreadNetworkDirectory] class.
type IMTRClusterThreadNetworkDirectory interface {
	IMTRGenericCluster
	
/* debug [class_interface_properties]: Properties for MTRClusterThreadNetworkDirectory */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRClusterThreadNetworkDirectory */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRClusterThreadNetworkDirectory */
// Alloc allocates a new instance without initialization.
func (mc _MTRClusterThreadNetworkDirectoryClass) Alloc() MTRClusterThreadNetworkDirectory {
	rv := objc.Send[MTRClusterThreadNetworkDirectory](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRClusterThreadNetworkDirectoryClass) New() MTRClusterThreadNetworkDirectory {
	rv := objc.Send[MTRClusterThreadNetworkDirectory](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterThreadNetworkDirectory) Init() MTRClusterThreadNetworkDirectory {
	rv := objc.Send[MTRClusterThreadNetworkDirectory](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterThreadNetworkDirectory) Autorelease() MTRClusterThreadNetworkDirectory {
	rv := objc.Send[MTRClusterThreadNetworkDirectory](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterThreadNetworkDirectory creates a new MTRClusterThreadNetworkDirectory instance.
func NewMTRClusterThreadNetworkDirectory() MTRClusterThreadNetworkDirectory {
	return getMTRClusterThreadNetworkDirectoryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRClusterThreadNetworkDirectory */
// Cluster Thread Network Directory Manages the names and credentials of Thread networks visible to the user.


// Cluster Thread Network Directory Manages the names and credentials of Thread networks visible to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadNetworkDirectory
type MTRClusterThreadNetworkDirectory struct {
	MTRGenericCluster
}

// MTRClusterThreadNetworkDirectoryFrom constructs a [MTRClusterThreadNetworkDirectory] from an unsafe.Pointer.
//
// Cluster Thread Network Directory Manages the names and credentials of Thread networks visible to the user.
func MTRClusterThreadNetworkDirectoryFrom(ptr unsafe.Pointer) MTRClusterThreadNetworkDirectory {
	return MTRClusterThreadNetworkDirectory{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRClusterThreadNetworkDirectory */

// For all instance methods that take a completion (i.e. command invocations), the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadNetworkDirectory/init(device:endpointID:queue:)
func NewMTRClusterThreadNetworkDirectoryWithDeviceEndpointIDQueue(device IMTRDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRClusterThreadNetworkDirectory {
	instance := getMTRClusterThreadNetworkDirectoryClass().Alloc()
	rv := objc.Send[MTRClusterThreadNetworkDirectory](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRClusterThreadNetworkDirectoryWithDeviceEndpointIDQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRClusterThreadNetworkDirectory */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRClusterThreadNetworkDirectory */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRClusterThreadNetworkDirectory */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRClusterThreadNetworkDirectory */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRClusterThreadNetworkDirectory */


