// Code generated from Apple documentation for IOSurface. DO NOT EDIT.

package iosurface

// Enum types and constants
// IOSurfaceComponentName enum type
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceComponentName
type SurfaceComponentName uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceComponentName/alpha
	kIOSurfaceComponentNameAlpha SurfaceComponentName = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceComponentName/blue
	kIOSurfaceComponentNameBlue SurfaceComponentName = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceComponentName/chromaBlue
	kIOSurfaceComponentNameChromaBlue SurfaceComponentName = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceComponentName/chromaRed
	kIOSurfaceComponentNameChromaRed SurfaceComponentName = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceComponentName/green
	kIOSurfaceComponentNameGreen SurfaceComponentName = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceComponentName/luma
	kIOSurfaceComponentNameLuma SurfaceComponentName = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceComponentName/red
	kIOSurfaceComponentNameRed SurfaceComponentName = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceComponentName/unknown
	kIOSurfaceComponentNameUnknown SurfaceComponentName = 0
)

// IOSurfaceComponentRange enum type
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceComponentRange
type SurfaceComponentRange uint

// IOSurfaceComponentType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceComponentType
type SurfaceComponentType uint

// IOSurfaceLockOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceLockOptions
type SurfaceLockOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceLockOptions/avoidSync
	kIOSurfaceLockAvoidSync SurfaceLockOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceLockOptions/readOnly
	kIOSurfaceLockReadOnly SurfaceLockOptions = 0
)

// IOSurfaceMemoryLedgerFlags enum type
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceMemoryLedgerFlags
type SurfaceMemoryLedgerFlags uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceMemoryLedgerFlags/noFootprint
	kIOSurfaceMemoryLedgerFlagNoFootprint SurfaceMemoryLedgerFlags = 0
)

// IOSurfaceMemoryLedgerTags enum type
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceMemoryLedgerTags
type SurfaceMemoryLedgerTags uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceMemoryLedgerTags/default
	kIOSurfaceMemoryLedgerTagDefault SurfaceMemoryLedgerTags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceMemoryLedgerTags/graphics
	kIOSurfaceMemoryLedgerTagGraphics SurfaceMemoryLedgerTags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceMemoryLedgerTags/media
	kIOSurfaceMemoryLedgerTagMedia SurfaceMemoryLedgerTags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceMemoryLedgerTags/network
	kIOSurfaceMemoryLedgerTagNetwork SurfaceMemoryLedgerTags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceMemoryLedgerTags/neural
	kIOSurfaceMemoryLedgerTagNeural SurfaceMemoryLedgerTags = 0
)

// IOSurfacePurgeabilityState enum type
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfacePurgeabilityState
type SurfacePurgeabilityState uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfacePurgeabilityState/kIOSurfacePurgeableNonVolatile
	kIOSurfacePurgeableNonVolatile SurfacePurgeabilityState = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfacePurgeabilityState/purgeableEmpty
	kIOSurfacePurgeableEmpty SurfacePurgeabilityState = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfacePurgeabilityState/purgeableKeepCurrent
	kIOSurfacePurgeableKeepCurrent SurfacePurgeabilityState = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfacePurgeabilityState/purgeableVolatile
	kIOSurfacePurgeableVolatile SurfacePurgeabilityState = 0
)

// IOSurfaceSubsampling enum type
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceSubsampling
type SurfaceSubsampling uint


