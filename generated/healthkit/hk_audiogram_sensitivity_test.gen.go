// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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


//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkaudiogramsensitivitytest/side
func (h_ HKAudiogramSensitivityTest) Side() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("side"))
	return rv
}


// SetSide sets the value of the side property.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkaudiogramsensitivitytest/side
func (h_ HKAudiogramSensitivityTest) SetSide(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSide:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkaudiogramsensitivitytest/sensitivity
func (h_ HKAudiogramSensitivityTest) Sensitivity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("sensitivity"))
	return rv
}


// SetSensitivity sets the value of the sensitivity property.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkaudiogramsensitivitytest/sensitivity
func (h_ HKAudiogramSensitivityTest) SetSensitivity(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSensitivity:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkaudiogramsensitivitytest/type
func (h_ HKAudiogramSensitivityTest) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("type"))
	return rv
}


// SetType sets the value of the type property.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkaudiogramsensitivitytest/type
func (h_ HKAudiogramSensitivityTest) SetType(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkaudiogramsensitivitytest/masked
func (h_ HKAudiogramSensitivityTest) Masked() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("masked"))
	return rv
}


// SetMasked sets the value of the masked property.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkaudiogramsensitivitytest/masked
func (h_ HKAudiogramSensitivityTest) SetMasked(value bool) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setMasked:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkaudiogramsensitivitytest/clampingrange
func (h_ HKAudiogramSensitivityTest) ClampingRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("clampingRange"))
	return rv
}


// SetClampingRange sets the value of the clampingRange property.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkaudiogramsensitivitytest/clampingrange
func (h_ HKAudiogramSensitivityTest) SetClampingRange(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setClampingRange:"), value)
}



