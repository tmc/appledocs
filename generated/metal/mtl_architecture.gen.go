// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [Architecture] class.
type IArchitecture interface {
	objectivec.IObject
	

	// properties:
	Name() foundation.foundation.INSString
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


	

	// methods:


}





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

























// The name of a GPU device’s architecture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArchitecture/name
func (a_ Architecture) Name() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("name"))
	return rv
}


// The architectural details of the GPU device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/architecture
func (a_ Architecture) Architecture() IMTLArchitecture {
	rv := objc.Send[Architecture](a_.ID, objc.Sel("architecture"))
	return rv
}


// The architectural details of the GPU device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/architecture
func (a_ Architecture) SetArchitecture(value IMTLArchitecture) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setArchitecture:"), value)
}


// A Boolean value that indicates whether a GPU device doesn’t have a connection to a display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/isheadless
func (a_ Architecture) IsHeadless() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isHeadless"))
	return rv
}


// A Boolean value that indicates whether a GPU device doesn’t have a connection to a display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/isheadless
func (a_ Architecture) SetIsHeadless(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsHeadless:"), value)
}


// A Boolean value that indicates whether the GPU lowers its performance to conserve energy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/islowpower
func (a_ Architecture) IsLowPower() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isLowPower"))
	return rv
}


// A Boolean value that indicates whether the GPU lowers its performance to conserve energy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/islowpower
func (a_ Architecture) SetIsLowPower(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsLowPower:"), value)
}


// A Boolean value that indicates whether the GPU is removable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/isremovable
func (a_ Architecture) IsRemovable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isRemovable"))
	return rv
}


// A Boolean value that indicates whether the GPU is removable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/isremovable
func (a_ Architecture) SetIsRemovable(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsRemovable:"), value)
}


// The physical location of the GPU relative to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/location
func (a_ Architecture) Location() DeviceLocation {
	rv := objc.Send[DeviceLocation](a_.ID, objc.Sel("location"))
	return rv
}


// The physical location of the GPU relative to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/location
func (a_ Architecture) SetLocation(value DeviceLocation) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLocation:"), value)
}


// A specific GPU position based on its general location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/locationnumber
func (a_ Architecture) LocationNumber() int {
	rv := objc.Send[int](a_.ID, objc.Sel("locationNumber"))
	return rv
}


// A specific GPU position based on its general location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/locationnumber
func (a_ Architecture) SetLocationNumber(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLocationNumber:"), value)
}


// The total number of GPUs in the peer group, if applicable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/peercount
func (a_ Architecture) PeerCount() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("peerCount"))
	return rv
}


// The total number of GPUs in the peer group, if applicable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/peercount
func (a_ Architecture) SetPeerCount(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPeerCount:"), value)
}


// The peer group ID the GPU belongs to, if applicable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/peergroupid
func (a_ Architecture) PeerGroupID() uint64 {
	rv := objc.Send[uint64](a_.ID, objc.Sel("peerGroupID"))
	return rv
}


// The peer group ID the GPU belongs to, if applicable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/peergroupid
func (a_ Architecture) SetPeerGroupID(value uint64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPeerGroupID:"), value)
}


// The unique identifier for a GPU in a peer group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/peerindex
func (a_ Architecture) PeerIndex() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("peerIndex"))
	return rv
}


// The unique identifier for a GPU in a peer group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/peerindex
func (a_ Architecture) SetPeerIndex(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPeerIndex:"), value)
}


// The GPU device’s registry identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/registryid
func (a_ Architecture) RegistryID() uint64 {
	rv := objc.Send[uint64](a_.ID, objc.Sel("registryID"))
	return rv
}


// The GPU device’s registry identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/registryid
func (a_ Architecture) SetRegistryID(value uint64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRegistryID:"), value)
}








