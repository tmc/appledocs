// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HKSource] class.
var (
	HKSourceClass     _HKSourceClass
	HKSourceClassOnce sync.Once
)

func getHKSourceClass() _HKSourceClass {
	HKSourceClassOnce.Do(func() {
		HKSourceClass = _HKSourceClass{objc.GetClass("HKSource")}
	})
	return HKSourceClass
}

type _HKSourceClass struct {
	class objc.Class
}

// An interface definition for the [HKSource] class.
type IHKSource interface {
	objectivec.IObject
	// properties:
	BundleIdentifier() string /* primitive/slice/pointer. */
	SetBundleIdentifier(value string /* primitive/slice/pointer. */)
	Name() string /* primitive/slice/pointer. */
	SetName(value string /* primitive/slice/pointer. */)
	// methods:
}

// An object indicating the app or device that created a HealthKit sample
//
// Sources include apps and devices that save data to the HealthKit store. Currently, HealthKit supports only the direct import of data from Bluetooth LE heart rate monitors. All other devices need a companion app to collect and save the data to HealthKit.


// An object indicating the app or device that created a HealthKit sample
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSource
type HKSource struct {
	objectivec.Object
}

// HKSourceFrom constructs a [HKSource] from an unsafe.Pointer.
//
// An object indicating the app or device that created a HealthKit sample
func HKSourceFrom(ptr unsafe.Pointer) HKSource {
	return HKSource{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKSourceClass) Alloc() HKSource {
	rv := objc.Send[HKSource](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKSourceClass) New() HKSource {
	rv := objc.Send[HKSource](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKSource) Init() HKSource {
	rv := objc.Send[HKSource](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKSource) Autorelease() HKSource {
	rv := objc.Send[HKSource](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKSource creates a new HKSource instance.
func NewHKSource() HKSource {
	return getHKSourceClass().New()
}



// The source’s bundle identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hksource/bundleidentifier
func (h_ HKSource) BundleIdentifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](h_.ID, objc.Sel("bundleIdentifier"))
	return rv
}


// The source’s bundle identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hksource/bundleidentifier
func (h_ HKSource) SetBundleIdentifier(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setBundleIdentifier:"), objc.String(value))
}


// The source’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hksource/name
func (h_ HKSource) Name() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](h_.ID, objc.Sel("name"))
	return rv
}


// The source’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hksource/name
func (h_ HKSource) SetName(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setName:"), objc.String(value))
}



