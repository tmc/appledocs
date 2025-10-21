// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [HKAudiogramSensitivityTest] class.
var (
	HKAudiogramSensitivityTestClass     _HKAudiogramSensitivityTestClass
	HKAudiogramSensitivityTestClassOnce sync.Once
)

func getHKAudiogramSensitivityTestClass() _HKAudiogramSensitivityTestClass {
	HKAudiogramSensitivityTestClassOnce.Do(func() {
		HKAudiogramSensitivityTestClass = _HKAudiogramSensitivityTestClass{objc.GetClass("HKAudiogramSensitivityTest")}
	})
	return HKAudiogramSensitivityTestClass
}

type _HKAudiogramSensitivityTestClass struct {
	class objc.Class
}

// An interface definition for the [HKAudiogramSensitivityTest] class.
type IHKAudiogramSensitivityTest interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSensitivityTest
type HKAudiogramSensitivityTest struct {
	objectivec.Object
}

// HKAudiogramSensitivityTestFrom constructs a [HKAudiogramSensitivityTest] from an unsafe.Pointer.
func HKAudiogramSensitivityTestFrom(ptr unsafe.Pointer) HKAudiogramSensitivityTest {
	return HKAudiogramSensitivityTest{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKAudiogramSensitivityTestClass) Alloc() HKAudiogramSensitivityTest {
	rv := objc.Send[HKAudiogramSensitivityTest](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKAudiogramSensitivityTestClass) New() HKAudiogramSensitivityTest {
	rv := objc.Send[HKAudiogramSensitivityTest](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKAudiogramSensitivityTest) Init() HKAudiogramSensitivityTest {
	rv := objc.Send[HKAudiogramSensitivityTest](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKAudiogramSensitivityTest) Autorelease() HKAudiogramSensitivityTest {
	rv := objc.Send[HKAudiogramSensitivityTest](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKAudiogramSensitivityTest creates a new HKAudiogramSensitivityTest instance.
func NewHKAudiogramSensitivityTest() HKAudiogramSensitivityTest {
	return getHKAudiogramSensitivityTestClass().New()
}




