// Code generated from Apple documentation for IOSurface. DO NOT EDIT.

package iosurface

// Enum types and constants
// IOSurfaceComponentName enum type
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceComponentName
type IOSurfaceComponentName uint

// IOSurfaceComponentRange enum type
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceComponentRange
type IOSurfaceComponentRange uint

// IOSurfaceComponentType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceComponentType
type IOSurfaceComponentType uint

// IOSurfaceLockOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceLockOptions
type IOSurfaceLockOptions uint

const (
	// kIOSurfaceLockAvoidSync - If you want to detect/avoid a potentially expensive paging operation (such as readback from a GPU to system memory) when you lock the buffer, you may include this flag. If locking the buffer requires a readback, the lock will fail with an error return of  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceLockOptions/avoidSync
	kIOSurfaceLockAvoidSync IOSurfaceLockOptions = 0
	// kIOSurfaceLockReadOnly - If you are not going to modify the data while you hold the lock, you should set this flag to avoid invalidating any existing caches of the buffer contents. This flag should be passed both to the lock and unlock functions. Non-symmentrical usage of this flag will result in undefined behavior.
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceLockOptions/readOnly
	kIOSurfaceLockReadOnly IOSurfaceLockOptions = 0
)

// IOSurfaceMemoryLedgerFlags enum type
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceMemoryLedgerFlags
type IOSurfaceMemoryLedgerFlags uint

// IOSurfaceMemoryLedgerTags enum type
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceMemoryLedgerTags
type IOSurfaceMemoryLedgerTags uint

// IOSurfacePurgeabilityState enum type
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfacePurgeabilityState
type IOSurfacePurgeabilityState uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfacePurgeabilityState/kIOSurfacePurgeableNonVolatile
	kIOSurfacePurgeableNonVolatile IOSurfacePurgeabilityState = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfacePurgeabilityState/purgeableEmpty
	kIOSurfacePurgeableEmpty IOSurfacePurgeabilityState = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfacePurgeabilityState/purgeableKeepCurrent
	kIOSurfacePurgeableKeepCurrent IOSurfacePurgeabilityState = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfacePurgeabilityState/purgeableVolatile
	kIOSurfacePurgeableVolatile IOSurfacePurgeabilityState = 0
)

// IOSurfaceSubsampling enum type
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceSubsampling
type IOSurfaceSubsampling uint


