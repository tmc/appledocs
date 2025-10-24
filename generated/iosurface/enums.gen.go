// Code generated from Apple documentation for IOSurface. DO NOT EDIT.

package iosurface

/* debug [enums.gen.go]: Generating 8 enums for IOSurface */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum IOSurfaceLockOptions (2 cases) */
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

/* debug [enums.gen.go]: Processing enum IOSurfacePurgeabilityState (4 cases) */
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

/* debug [enums.gen.go]: Processing enum IOSurfaceComponentName (8 cases) */
// IOSurfaceComponentName enum type
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceComponentName
type IOSurfaceComponentName uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceComponentName/alpha
	kIOSurfaceComponentNameAlpha IOSurfaceComponentName = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceComponentName/blue
	kIOSurfaceComponentNameBlue IOSurfaceComponentName = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceComponentName/chromaBlue
	kIOSurfaceComponentNameChromaBlue IOSurfaceComponentName = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceComponentName/chromaRed
	kIOSurfaceComponentNameChromaRed IOSurfaceComponentName = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceComponentName/green
	kIOSurfaceComponentNameGreen IOSurfaceComponentName = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceComponentName/luma
	kIOSurfaceComponentNameLuma IOSurfaceComponentName = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceComponentName/red
	kIOSurfaceComponentNameRed IOSurfaceComponentName = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceComponentName/unknown
	kIOSurfaceComponentNameUnknown IOSurfaceComponentName = 0
)

/* debug [enums.gen.go]: Processing enum IOSurfaceComponentRange (4 cases) */
// IOSurfaceComponentRange enum type
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceComponentRange
type IOSurfaceComponentRange uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceComponentRange/fullRange
	kIOSurfaceComponentRangeFullRange IOSurfaceComponentRange = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceComponentRange/unknown
	kIOSurfaceComponentRangeUnknown IOSurfaceComponentRange = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceComponentRange/videoRange
	kIOSurfaceComponentRangeVideoRange IOSurfaceComponentRange = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceComponentRange/wideRange
	kIOSurfaceComponentRangeWideRange IOSurfaceComponentRange = 0
)

/* debug [enums.gen.go]: Processing enum IOSurfaceComponentType (5 cases) */
// IOSurfaceComponentType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceComponentType
type IOSurfaceComponentType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceComponentType/float
	kIOSurfaceComponentTypeFloat IOSurfaceComponentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceComponentType/signedInteger
	kIOSurfaceComponentTypeSignedInteger IOSurfaceComponentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceComponentType/signedNormalized
	kIOSurfaceComponentTypeSignedNormalized IOSurfaceComponentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceComponentType/unknown
	kIOSurfaceComponentTypeUnknown IOSurfaceComponentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceComponentType/unsignedInteger
	kIOSurfaceComponentTypeUnsignedInteger IOSurfaceComponentType = 0
)

/* debug [enums.gen.go]: Processing enum IOSurfaceMemoryLedgerFlags (1 cases) */
// IOSurfaceMemoryLedgerFlags enum type
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceMemoryLedgerFlags
type IOSurfaceMemoryLedgerFlags uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceMemoryLedgerFlags/noFootprint
	kIOSurfaceMemoryLedgerFlagNoFootprint IOSurfaceMemoryLedgerFlags = 0
)

/* debug [enums.gen.go]: Processing enum IOSurfaceMemoryLedgerTags (5 cases) */
// IOSurfaceMemoryLedgerTags enum type
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceMemoryLedgerTags
type IOSurfaceMemoryLedgerTags uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceMemoryLedgerTags/default
	kIOSurfaceMemoryLedgerTagDefault IOSurfaceMemoryLedgerTags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceMemoryLedgerTags/graphics
	kIOSurfaceMemoryLedgerTagGraphics IOSurfaceMemoryLedgerTags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceMemoryLedgerTags/media
	kIOSurfaceMemoryLedgerTagMedia IOSurfaceMemoryLedgerTags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceMemoryLedgerTags/network
	kIOSurfaceMemoryLedgerTagNetwork IOSurfaceMemoryLedgerTags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceMemoryLedgerTags/neural
	kIOSurfaceMemoryLedgerTagNeural IOSurfaceMemoryLedgerTags = 0
)

/* debug [enums.gen.go]: Processing enum IOSurfaceSubsampling (5 cases) */
// IOSurfaceSubsampling enum type
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceSubsampling
type IOSurfaceSubsampling uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceSubsampling/subsampling411
	kIOSurfaceSubsampling411 IOSurfaceSubsampling = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceSubsampling/subsampling420
	kIOSurfaceSubsampling420 IOSurfaceSubsampling = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceSubsampling/subsampling422
	kIOSurfaceSubsampling422 IOSurfaceSubsampling = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceSubsampling/subsamplingNone
	kIOSurfaceSubsamplingNone IOSurfaceSubsampling = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceSubsampling/subsamplingUnknown
	kIOSurfaceSubsamplingUnknown IOSurfaceSubsampling = 0
)


