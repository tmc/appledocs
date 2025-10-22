// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AudioUnitBusArray] class.
var (
	AudioUnitBusArrayClass     _AudioUnitBusArrayClass
	AudioUnitBusArrayClassOnce sync.Once
)

func getAudioUnitBusArrayClass() _AudioUnitBusArrayClass {
	AudioUnitBusArrayClassOnce.Do(func() {
		AudioUnitBusArrayClass = _AudioUnitBusArrayClass{objc.GetClass("AUAudioUnitBusArray")}
	})
	return AudioUnitBusArrayClass
}

type _AudioUnitBusArrayClass struct {
	class objc.Class
}

// An interface definition for the [AudioUnitBusArray] class.
type IAudioUnitBusArray interface {
	objectivec.IObject
	AddObserverToAllBussesForKeyPathOptionsContext(observer foundation.IObject, keyPath string, options unsafe.Pointer, context unsafe.Pointer)
	RemoveObserverFromAllBussesForKeyPathContext(observer foundation.IObject, keyPath string, context unsafe.Pointer)
	ReplaceBusses(busArray []AudioUnitBus)
	SetBusCountError(count uint, outError unsafe.Pointer) bool
	ObjectAtIndexedSubscript(index uint) AudioUnitBus
	BusType() AudioUnitBusType
	Count() uint
	CountChangeable() bool
	OwnerAudioUnit() AUAudioUnit
	InputBusses() AUAudioUnitBusArray
	SetInputBusses(value IAUAudioUnitBusArray)
	OutputBusses() AUAudioUnitBusArray
	SetOutputBusses(value IAUAudioUnitBusArray)
	IsCountChangeable() bool
	SetIsCountChangeable(value bool)
}

// A class that defines a container for an audio unit’s input or output busses.
//
// Hosts can observe a bus property across all busses by using KVO on a bus array object, without having to observe it on each individual bus. Some audio units (e.g. mixers) support variable numbers of busses, via subclassing. When the bus count changes, a KVO notification is sent on the audio unit’s or property, as appropriate. This version 3 class is bridged to the version 2 API.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray
type AudioUnitBusArray struct {
	objectivec.Object
}

// AudioUnitBusArrayFrom constructs a [AudioUnitBusArray] from an unsafe.Pointer.
//
// A class that defines a container for an audio unit’s input or output busses.
func AudioUnitBusArrayFrom(ptr unsafe.Pointer) AudioUnitBusArray {
	return AudioUnitBusArray{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AudioUnitBusArrayClass) Alloc() AudioUnitBusArray {
	rv := objc.Send[AudioUnitBusArray](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AudioUnitBusArrayClass) New() AudioUnitBusArray {
	rv := objc.Send[AudioUnitBusArray](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioUnitBusArray) Init() AudioUnitBusArray {
	rv := objc.Send[AudioUnitBusArray](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioUnitBusArray) Autorelease() AudioUnitBusArray {
	rv := objc.Send[AudioUnitBusArray](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioUnitBusArray creates a new AudioUnitBusArray instance.
func NewAudioUnitBusArray() AudioUnitBusArray {
	return getAudioUnitBusArrayClass().New()
}




// Initializes an empty bus array.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray/init(audioUnit:busType:)
func NewAudioUnitBusArrayWithAudioUnitBusType(owner IAUAudioUnit, busType AudioUnitBusType) AudioUnitBusArray {
	instance := getAudioUnitBusArrayClass().Alloc()
	rv := objc.Send[AudioUnitBusArray](instance.ID, objc.Sel("initWithAudioUnit:busType:"), owner, busType)
	rv.Autorelease()
	return rv
}



// Initializes a bus array by making a copy of the supplied busses.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray/init(audioUnit:busType:busses:)
func NewAudioUnitBusArrayWithAudioUnitBusTypeBusses(owner IAUAudioUnit, busType AudioUnitBusType, busArray []AudioUnitBus) AudioUnitBusArray {
	instance := getAudioUnitBusArrayClass().Alloc()
	rv := objc.Send[AudioUnitBusArray](instance.ID, objc.Sel("initWithAudioUnit:busType:busses:"), owner, busType, busArray)
	rv.Autorelease()
	return rv
}


// Adds a KVO observer for a given property on all busses in the array.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray/addObserver(toAllBusses:forKeyPath:options:context:)
func (a_ AudioUnitBusArray) AddObserverToAllBussesForKeyPathOptionsContext(observer foundation.IObject, keyPath string, options unsafe.Pointer, context unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("addObserverToAllBusses:forKeyPath:options:context:"), observer, objc.String(keyPath), options, context)
}

// Removes a KVO observer for a given property on all busses in the array.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray/removeObserver(fromAllBusses:forKeyPath:context:)
func (a_ AudioUnitBusArray) RemoveObserverFromAllBussesForKeyPathContext(observer foundation.IObject, keyPath string, context unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("removeObserverFromAllBusses:forKeyPath:context:"), observer, objc.String(keyPath), context)
}

// Replaces the current bus array with a copy of the supplied bus array.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray/replaceBusses(_:)
func (a_ AudioUnitBusArray) ReplaceBusses(busArray []AudioUnitBus) {
	objc.Send[objc.ID](a_.ID, objc.Sel("replaceBusses:"), busArray)
}

// Changes the number of busses in the array.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray/setBusCount(_:)
func (a_ AudioUnitBusArray) SetBusCountError(count uint, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setBusCount:error:"), count, outError)
	return rv
}

// Returns the bus at the specified index.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray/subscript(_:)
func (a_ AudioUnitBusArray) ObjectAtIndexedSubscript(index uint) AudioUnitBus {
	rv := objc.Send[AudioUnitBus](a_.ID, objc.Sel("objectAtIndexedSubscript:"), index)
	return rv
}

// Determines whether the bus array is for input or output.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray/busType
func (a_ AudioUnitBusArray) BusType() AudioUnitBusType {
	rv := objc.Send[AudioUnitBusType](a_.ID, objc.Sel("busType"))
	return rv
}

// The number of busses in the array.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray/count
func (a_ AudioUnitBusArray) Count() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("count"))
	return rv
}

// Determines whether the array can have a variable number of busses.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray/isCountChangeable
func (a_ AudioUnitBusArray) CountChangeable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("countChangeable"))
	return rv
}

// The audio unit that owns the bus array.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray/ownerAudioUnit
func (a_ AudioUnitBusArray) OwnerAudioUnit() AUAudioUnit {
	rv := objc.Send[AUAudioUnit](a_.ID, objc.Sel("ownerAudioUnit"))
	return rv
}

// An array containing the audio unit’s input connection points.
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/inputbusses
func (a_ AudioUnitBusArray) InputBusses() AUAudioUnitBusArray {
	rv := objc.Send[AUAudioUnitBusArray](a_.ID, objc.Sel("inputBusses"))
	return rv
}


// SetInputBusses sets the value of the inputBusses property.
// An array containing the audio unit’s input connection points.

//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/inputbusses
func (a_ AudioUnitBusArray) SetInputBusses(value IAUAudioUnitBusArray) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setInputBusses:"), value)
}

// An array containing the audio unit’s output connection points.
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/outputbusses
func (a_ AudioUnitBusArray) OutputBusses() AUAudioUnitBusArray {
	rv := objc.Send[AUAudioUnitBusArray](a_.ID, objc.Sel("outputBusses"))
	return rv
}


// SetOutputBusses sets the value of the outputBusses property.
// An array containing the audio unit’s output connection points.

//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/outputbusses
func (a_ AudioUnitBusArray) SetOutputBusses(value IAUAudioUnitBusArray) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputBusses:"), value)
}

// Determines whether the array can have a variable number of busses.
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounitbusarray/iscountchangeable
func (a_ AudioUnitBusArray) IsCountChangeable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isCountChangeable"))
	return rv
}


// SetIsCountChangeable sets the value of the isCountChangeable property.
// Determines whether the array can have a variable number of busses.

//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounitbusarray/iscountchangeable
func (a_ AudioUnitBusArray) SetIsCountChangeable(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsCountChangeable:"), value)
}


