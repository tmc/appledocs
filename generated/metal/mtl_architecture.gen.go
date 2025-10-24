// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLArchitecture */


/* debug [class_header]: Header for MTLArchitecture */
// The class instance for the [Architecture] class.
var (
	ArchitectureClass     _ArchitectureClass
	ArchitectureClassOnce sync.Once
)

func getArchitectureClass() _ArchitectureClass {
	ArchitectureClassOnce.Do(func() {
		ArchitectureClass = _ArchitectureClass{objc.GetClass("MTLArchitecture")}
	})
	return ArchitectureClass
}

type _ArchitectureClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Architecture */
// An interface definition for the [Architecture] class.
type IArchitecture interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Architecture */
	// properties:
	Name() objc.IObject /* cross-framework: NSString */
	Architecture() IMTLArchitecture
	SetArchitecture(value IMTLArchitecture)
	IsHeadless() bool
	SetIsHeadless(value bool)
	IsLowPower() bool
	SetIsLowPower(value bool)
	IsRemovable() bool
	SetIsRemovable(value bool)
	Location() DeviceLocation
	SetLocation(value DeviceLocation)
	LocationNumber() int
	SetLocationNumber(value int)
	PeerCount() objectivec.IObject
	SetPeerCount(value objectivec.IObject)
	PeerGroupID() uint64
	SetPeerGroupID(value uint64)
	PeerIndex() objectivec.IObject
	SetPeerIndex(value objectivec.IObject)
	RegistryID() uint64
	SetRegistryID(value uint64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Architecture */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Architecture */
// Alloc allocates a new instance without initialization.
func (ac _ArchitectureClass) Alloc() Architecture {
	rv := objc.Send[Architecture](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _ArchitectureClass) New() Architecture {
	rv := objc.Send[Architecture](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ Architecture) Init() Architecture {
	rv := objc.Send[Architecture](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ Architecture) Autorelease() Architecture {
	rv := objc.Send[Architecture](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewArchitecture creates a new Architecture instance.
func NewArchitecture() Architecture {
	return getArchitectureClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Architecture */
// A class that contains the architectural details of a GPU device.


// A class that contains the architectural details of a GPU device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArchitecture
type Architecture struct {
	objectivec.Object
}

// ArchitectureFrom constructs a [Architecture] from an unsafe.Pointer.
//
// A class that contains the architectural details of a GPU device.
func ArchitectureFrom(ptr unsafe.Pointer) Architecture {
	return Architecture{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Architecture *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Architecture */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Architecture */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Architecture */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Architecture */

// The name of a GPU device’s architecture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArchitecture/name
func (a_ Architecture) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The architectural details of the GPU device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/architecture
func (a_ Architecture) Architecture() IMTLArchitecture {
	rv := objc.Send[Architecture](a_.ID, objc.Sel("architecture"))
	return rv
}/* debug [instance_properties/getter]: architecture */


// The architectural details of the GPU device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/architecture
func (a_ Architecture) SetArchitecture(value IMTLArchitecture) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setArchitecture:"), value)
}/* debug [instance_properties/setter]: architecture */


// A Boolean value that indicates whether a GPU device doesn’t have a connection to a display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/isheadless
func (a_ Architecture) IsHeadless() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isHeadless"))
	return rv
}/* debug [instance_properties/getter]: isHeadless */


// A Boolean value that indicates whether a GPU device doesn’t have a connection to a display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/isheadless
func (a_ Architecture) SetIsHeadless(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsHeadless:"), value)
}/* debug [instance_properties/setter]: isHeadless */


// A Boolean value that indicates whether the GPU lowers its performance to conserve energy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/islowpower
func (a_ Architecture) IsLowPower() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isLowPower"))
	return rv
}/* debug [instance_properties/getter]: isLowPower */


// A Boolean value that indicates whether the GPU lowers its performance to conserve energy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/islowpower
func (a_ Architecture) SetIsLowPower(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsLowPower:"), value)
}/* debug [instance_properties/setter]: isLowPower */


// A Boolean value that indicates whether the GPU is removable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/isremovable
func (a_ Architecture) IsRemovable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isRemovable"))
	return rv
}/* debug [instance_properties/getter]: isRemovable */


// A Boolean value that indicates whether the GPU is removable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/isremovable
func (a_ Architecture) SetIsRemovable(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsRemovable:"), value)
}/* debug [instance_properties/setter]: isRemovable */


// The physical location of the GPU relative to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/location
func (a_ Architecture) Location() DeviceLocation {
	rv := objc.Send[DeviceLocation](a_.ID, objc.Sel("location"))
	return rv
}/* debug [instance_properties/getter]: location */


// The physical location of the GPU relative to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/location
func (a_ Architecture) SetLocation(value DeviceLocation) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLocation:"), value)
}/* debug [instance_properties/setter]: location */


// A specific GPU position based on its general location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/locationnumber
func (a_ Architecture) LocationNumber() int {
	rv := objc.Send[int](a_.ID, objc.Sel("locationNumber"))
	return rv
}/* debug [instance_properties/getter]: locationNumber */


// A specific GPU position based on its general location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/locationnumber
func (a_ Architecture) SetLocationNumber(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLocationNumber:"), value)
}/* debug [instance_properties/setter]: locationNumber */


// The total number of GPUs in the peer group, if applicable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/peercount
func (a_ Architecture) PeerCount() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("peerCount"))
	return rv
}/* debug [instance_properties/getter]: peerCount */


// The total number of GPUs in the peer group, if applicable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/peercount
func (a_ Architecture) SetPeerCount(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPeerCount:"), value)
}/* debug [instance_properties/setter]: peerCount */


// The peer group ID the GPU belongs to, if applicable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/peergroupid
func (a_ Architecture) PeerGroupID() uint64 {
	rv := objc.Send[uint64](a_.ID, objc.Sel("peerGroupID"))
	return rv
}/* debug [instance_properties/getter]: peerGroupID */


// The peer group ID the GPU belongs to, if applicable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/peergroupid
func (a_ Architecture) SetPeerGroupID(value uint64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPeerGroupID:"), value)
}/* debug [instance_properties/setter]: peerGroupID */


// The unique identifier for a GPU in a peer group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/peerindex
func (a_ Architecture) PeerIndex() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("peerIndex"))
	return rv
}/* debug [instance_properties/getter]: peerIndex */


// The unique identifier for a GPU in a peer group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/peerindex
func (a_ Architecture) SetPeerIndex(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPeerIndex:"), value)
}/* debug [instance_properties/setter]: peerIndex */


// The GPU device’s registry identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/registryid
func (a_ Architecture) RegistryID() uint64 {
	rv := objc.Send[uint64](a_.ID, objc.Sel("registryID"))
	return rv
}/* debug [instance_properties/getter]: registryID */


// The GPU device’s registry identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/registryid
func (a_ Architecture) SetRegistryID(value uint64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRegistryID:"), value)
}/* debug [instance_properties/setter]: registryID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLArchitecture */



