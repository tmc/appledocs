// Code generated from Apple documentation for IOSurface. DO NOT EDIT.

package iosurface

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Surface] class.
var (
	SurfaceClass     _SurfaceClass
	SurfaceClassOnce sync.Once
)

func getSurfaceClass() _SurfaceClass {
	SurfaceClassOnce.Do(func() {
		SurfaceClass = _SurfaceClass{objc.GetClass("IOSurface")}
	})
	return SurfaceClass
}

type _SurfaceClass struct {
	class objc.Class
}

// An interface definition for the [Surface] class.
type ISurface interface {
	objectivec.IObject
	// properties:
	AllocationSize() int
	AllowsPixelSizeCasting() bool
	BaseAddress() unsafe.Pointer
	BytesPerElement() int
	BytesPerRow() int
	ElementHeight() int
	ElementWidth() int
	Height() int
	InUse() bool
	LocalUseCount() int32 /* not a class type */
	PixelFormat() uint32 /* not a class type */
	PlaneCount() uint
	Seed() uint32 /* not a class type */
	SurfaceID() uint32 /* not a class type */
	Width() int
	IsInUse() bool
	SetIsInUse(value bool)
	// methods:
	AllAttachments() foundation.IDictionary
	AttachmentForKey(key objc.IObject /* cross-framework: NSString */) objc.ID
	BaseAddressOfPlaneAtIndex(planeIndex uint)
	BytesPerElementOfPlaneAtIndex(planeIndex uint) int
	BytesPerRowOfPlaneAtIndex(planeIndex uint) int
	DecrementUseCount()
	ElementHeightOfPlaneAtIndex(planeIndex uint) int
	ElementWidthOfPlaneAtIndex(planeIndex uint) int
	HeightOfPlaneAtIndex(planeIndex uint) int
	IncrementUseCount()
	LockWithOptionsSeed(options SurfaceLockOptions, seed unsafe.Pointer) unsafe.Pointer
	RemoveAllAttachments()
	RemoveAttachmentForKey(key objc.IObject /* cross-framework: NSString */)
	SetAllAttachments(dict foundation.IDictionary)
	SetAttachmentForKey(anObject objectivec.IObject, key objc.IObject /* cross-framework: NSString */)
	SetPurgeableOldState(newState SurfacePurgeabilityState, oldState SurfacePurgeabilityState) unsafe.Pointer
	UnlockWithOptionsSeed(options SurfaceLockOptions, seed unsafe.Pointer) unsafe.Pointer
	WidthOfPlaneAtIndex(planeIndex uint) int
}

// Data type representing an IOSurface opaque object.


// Data type representing an IOSurface opaque object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurface
type Surface struct {
	objectivec.Object
}

// SurfaceFrom constructs a [Surface] from an unsafe.Pointer.
//
// Data type representing an IOSurface opaque object.
func SurfaceFrom(ptr unsafe.Pointer) Surface {
	return Surface{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SurfaceClass) Alloc() Surface {
	rv := objc.Send[Surface](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SurfaceClass) New() Surface {
	rv := objc.Send[Surface](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ Surface) Init() Surface {
	rv := objc.Send[Surface](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ Surface) Autorelease() Surface {
	rv := objc.Send[Surface](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSurface creates a new Surface instance.
func NewSurface() Surface {
	return getSurfaceClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurface/init(properties:)
func NewSurfaceWithProperties(properties foundation.IDictionary) Surface {
	instance := getSurfaceClass().Alloc()
	rv := objc.Send[Surface](instance.ID, objc.Sel("initWithProperties:"), properties)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurface/allAttachments()
func (s_ Surface) AllAttachments() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](s_.ID, objc.Sel("allAttachments"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurface/attachment(forKey:)
func (s_ Surface) AttachmentForKey(key objc.IObject /* cross-framework: NSString */) objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("attachmentForKey:"), key)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurface/baseAddressOfPlane(at:)
func (s_ Surface) BaseAddressOfPlaneAtIndex(planeIndex uint) {
	objc.Send[objc.ID](s_.ID, objc.Sel("baseAddressOfPlaneAtIndex:"), planeIndex)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurface/bytesPerElementOfPlane(at:)
func (s_ Surface) BytesPerElementOfPlaneAtIndex(planeIndex uint) int {
	rv := objc.Send[int](s_.ID, objc.Sel("bytesPerElementOfPlaneAtIndex:"), planeIndex)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurface/bytesPerRowOfPlane(at:)
func (s_ Surface) BytesPerRowOfPlaneAtIndex(planeIndex uint) int {
	rv := objc.Send[int](s_.ID, objc.Sel("bytesPerRowOfPlaneAtIndex:"), planeIndex)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurface/decrementUseCount()
func (s_ Surface) DecrementUseCount() {
	objc.Send[objc.ID](s_.ID, objc.Sel("decrementUseCount"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurface/elementHeightOfPlane(at:)
func (s_ Surface) ElementHeightOfPlaneAtIndex(planeIndex uint) int {
	rv := objc.Send[int](s_.ID, objc.Sel("elementHeightOfPlaneAtIndex:"), planeIndex)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurface/elementWidthOfPlane(at:)
func (s_ Surface) ElementWidthOfPlaneAtIndex(planeIndex uint) int {
	rv := objc.Send[int](s_.ID, objc.Sel("elementWidthOfPlaneAtIndex:"), planeIndex)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurface/heightOfPlane(at:)
func (s_ Surface) HeightOfPlaneAtIndex(planeIndex uint) int {
	rv := objc.Send[int](s_.ID, objc.Sel("heightOfPlaneAtIndex:"), planeIndex)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurface/incrementUseCount()
func (s_ Surface) IncrementUseCount() {
	objc.Send[objc.ID](s_.ID, objc.Sel("incrementUseCount"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurface/lock(options:seed:)
func (s_ Surface) LockWithOptionsSeed(options SurfaceLockOptions, seed unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("lockWithOptions:seed:"), options, seed)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurface/removeAllAttachments()
func (s_ Surface) RemoveAllAttachments() {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeAllAttachments"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurface/removeAttachment(forKey:)
func (s_ Surface) RemoveAttachmentForKey(key objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeAttachmentForKey:"), key)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurface/setAllAttachments(_:)
func (s_ Surface) SetAllAttachments(dict foundation.IDictionary) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAllAttachments:"), dict)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurface/setAttachment(_:forKey:)
func (s_ Surface) SetAttachmentForKey(anObject objectivec.IObject, key objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAttachment:forKey:"), anObject, key)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurface/setPurgeable(_:oldState:)
func (s_ Surface) SetPurgeableOldState(newState SurfacePurgeabilityState, oldState SurfacePurgeabilityState) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("setPurgeable:oldState:"), newState, oldState)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurface/unlock(options:seed:)
func (s_ Surface) UnlockWithOptionsSeed(options SurfaceLockOptions, seed unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("unlockWithOptions:seed:"), options, seed)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurface/widthOfPlane(at:)
func (s_ Surface) WidthOfPlaneAtIndex(planeIndex uint) int {
	rv := objc.Send[int](s_.ID, objc.Sel("widthOfPlaneAtIndex:"), planeIndex)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurface/allocationSize
func (s_ Surface) AllocationSize() int {
	rv := objc.Send[int](s_.ID, objc.Sel("allocationSize"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurface/allowsPixelSizeCasting
func (s_ Surface) AllowsPixelSizeCasting() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("allowsPixelSizeCasting"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurface/baseAddress
func (s_ Surface) BaseAddress() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("baseAddress"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurface/bytesPerElement
func (s_ Surface) BytesPerElement() int {
	rv := objc.Send[int](s_.ID, objc.Sel("bytesPerElement"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurface/bytesPerRow
func (s_ Surface) BytesPerRow() int {
	rv := objc.Send[int](s_.ID, objc.Sel("bytesPerRow"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurface/elementHeight
func (s_ Surface) ElementHeight() int {
	rv := objc.Send[int](s_.ID, objc.Sel("elementHeight"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurface/elementWidth
func (s_ Surface) ElementWidth() int {
	rv := objc.Send[int](s_.ID, objc.Sel("elementWidth"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurface/height
func (s_ Surface) Height() int {
	rv := objc.Send[int](s_.ID, objc.Sel("height"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurface/isInUse
func (s_ Surface) InUse() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("inUse"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurface/localUseCount
func (s_ Surface) LocalUseCount() int32 /* not a class type */ {
	rv := objc.Send[int32](s_.ID, objc.Sel("localUseCount"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurface/pixelFormat
func (s_ Surface) PixelFormat() uint32 /* not a class type */ {
	rv := objc.Send[uint32](s_.ID, objc.Sel("pixelFormat"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurface/planeCount
func (s_ Surface) PlaneCount() uint {
	rv := objc.Send[uint](s_.ID, objc.Sel("planeCount"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurface/seed
func (s_ Surface) Seed() uint32 /* not a class type */ {
	rv := objc.Send[uint32](s_.ID, objc.Sel("seed"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurface/surfaceID
func (s_ Surface) SurfaceID() uint32 /* not a class type */ {
	rv := objc.Send[uint32](s_.ID, objc.Sel("surfaceID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurface/width
func (s_ Surface) Width() int {
	rv := objc.Send[int](s_.ID, objc.Sel("width"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurface/isinuse
func (s_ Surface) IsInUse() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isInUse"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iosurface/iosurface/isinuse
func (s_ Surface) SetIsInUse(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsInUse:"), value)
}


