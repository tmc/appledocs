// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AUAudioUnitBusArray */


/* debug [class_header]: Header for AUAudioUnitBusArray */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioUnitBusArray */
// An interface definition for the [AudioUnitBusArray] class.
type IAudioUnitBusArray interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AudioUnitBusArray */
	// properties:
	BusType() AudioUnitBusType
	Count() uint
	CountChangeable() bool
	OwnerAudioUnit() IAUAudioUnit
	InputBusses() IAUAudioUnitBusArray
	SetInputBusses(value IAUAudioUnitBusArray)
	OutputBusses() IAUAudioUnitBusArray
	SetOutputBusses(value IAUAudioUnitBusArray)
	IsCountChangeable() bool
	SetIsCountChangeable(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioUnitBusArray */
	// methods:
	AddObserverToAllBussesForKeyPathOptionsContext(observer objc.IObject /* cross-framework: NSObject */, keyPath objc.IObject /* cross-framework: NSString */, options uint, context objectivec.IObject)
	RemoveObserverFromAllBussesForKeyPathContext(observer objc.IObject /* cross-framework: NSObject */, keyPath objc.IObject /* cross-framework: NSString */, context objectivec.IObject)
	ReplaceBusses(busArray []AudioUnitBus)
	SetBusCountError(count uint, outError objectivec.IObject) bool
	ObjectAtIndexedSubscript(index uint) IAudioUnitBus
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioUnitBusArray */
// Alloc allocates a new instance without initialization.
func (ac _AudioUnitBusArrayClass) Alloc() AudioUnitBusArray {
	rv := objc.Send[AudioUnitBusArray](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioUnitBusArray */
// A class that defines a container for an audio unit’s input or output busses.
//
// Hosts can observe a bus property across all busses by using KVO on a bus array object, without having to observe it on each individual bus. Some audio units (e.g. mixers) support variable numbers of busses, via subclassing. When the bus count changes, a KVO notification is sent on the audio unit’s or property, as appropriate. This version 3 class is bridged to the version 2 API.


// A class that defines a container for an audio unit’s input or output busses.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioUnitBusArray */

// Initializes an empty bus array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray/init(audioUnit:busType:)
func NewAudioUnitBusArrayWithAudioUnitBusType(owner IAUAudioUnit, busType AudioUnitBusType) AudioUnitBusArray {
	instance := getAudioUnitBusArrayClass().Alloc()
	rv := objc.Send[AudioUnitBusArray](instance.ID, objc.Sel("initWithAudioUnit:busType:"), owner, busType)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAudioUnitBusArrayWithAudioUnitBusType */


// Initializes a bus array by making a copy of the supplied busses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray/init(audioUnit:busType:busses:)
func NewAudioUnitBusArrayWithAudioUnitBusTypeBusses(owner IAUAudioUnit, busType AudioUnitBusType, busArray []AudioUnitBus) AudioUnitBusArray {
	instance := getAudioUnitBusArrayClass().Alloc()
	rv := objc.Send[AudioUnitBusArray](instance.ID, objc.Sel("initWithAudioUnit:busType:busses:"), owner, busType, busArray)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAudioUnitBusArrayWithAudioUnitBusTypeBusses */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioUnitBusArray */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioUnitBusArray */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioUnitBusArray */

// Adds a KVO observer for a given property on all busses in the array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray/addObserver(toAllBusses:forKeyPath:options:context:)
func (a_ AudioUnitBusArray) AddObserverToAllBussesForKeyPathOptionsContext(observer objc.IObject /* cross-framework: NSObject */, keyPath objc.IObject /* cross-framework: NSString */, options uint, context objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("addObserverToAllBusses:forKeyPath:options:context:"), observer, keyPath, options, context)
}/* debug [instance_methods/method]: AddObserverToAllBussesForKeyPathOptionsContext */


// Removes a KVO observer for a given property on all busses in the array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray/removeObserver(fromAllBusses:forKeyPath:context:)
func (a_ AudioUnitBusArray) RemoveObserverFromAllBussesForKeyPathContext(observer objc.IObject /* cross-framework: NSObject */, keyPath objc.IObject /* cross-framework: NSString */, context objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("removeObserverFromAllBusses:forKeyPath:context:"), observer, keyPath, context)
}/* debug [instance_methods/method]: RemoveObserverFromAllBussesForKeyPathContext */


// Replaces the current bus array with a copy of the supplied bus array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray/replaceBusses(_:)
func (a_ AudioUnitBusArray) ReplaceBusses(busArray []AudioUnitBus) {
	objc.Send[objc.ID](a_.ID, objc.Sel("replaceBusses:"), busArray)
}/* debug [instance_methods/method]: ReplaceBusses */


// Changes the number of busses in the array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray/setBusCount(_:)
func (a_ AudioUnitBusArray) SetBusCountError(count uint, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setBusCount:error:"), count, outError)
	return rv
}/* debug [instance_methods/method]: SetBusCountError */


// Returns the bus at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray/subscript(_:)
func (a_ AudioUnitBusArray) ObjectAtIndexedSubscript(index uint) IAudioUnitBus {
	rv := objc.Send[AudioUnitBus](a_.ID, objc.Sel("objectAtIndexedSubscript:"), index)
	return rv
}/* debug [instance_methods/method]: ObjectAtIndexedSubscript */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioUnitBusArray */

// Determines whether the bus array is for input or output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray/busType
func (a_ AudioUnitBusArray) BusType() AudioUnitBusType {
	rv := objc.Send[AudioUnitBusType](a_.ID, objc.Sel("busType"))
	return rv
}/* debug [instance_properties/getter]: busType */


// The number of busses in the array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray/count
func (a_ AudioUnitBusArray) Count() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("count"))
	return rv
}/* debug [instance_properties/getter]: count */


// Determines whether the array can have a variable number of busses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray/isCountChangeable
func (a_ AudioUnitBusArray) CountChangeable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("countChangeable"))
	return rv
}/* debug [instance_properties/getter]: countChangeable */


// The audio unit that owns the bus array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray/ownerAudioUnit
func (a_ AudioUnitBusArray) OwnerAudioUnit() IAUAudioUnit {
	rv := objc.Send[AudioUnit](a_.ID, objc.Sel("ownerAudioUnit"))
	return rv
}/* debug [instance_properties/getter]: ownerAudioUnit */


// An array containing the audio unit’s input connection points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/inputbusses
func (a_ AudioUnitBusArray) InputBusses() IAUAudioUnitBusArray {
	rv := objc.Send[AudioUnitBusArray](a_.ID, objc.Sel("inputBusses"))
	return rv
}/* debug [instance_properties/getter]: inputBusses */


// An array containing the audio unit’s input connection points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/inputbusses
func (a_ AudioUnitBusArray) SetInputBusses(value IAUAudioUnitBusArray) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setInputBusses:"), value)
}/* debug [instance_properties/setter]: inputBusses */


// An array containing the audio unit’s output connection points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/outputbusses
func (a_ AudioUnitBusArray) OutputBusses() IAUAudioUnitBusArray {
	rv := objc.Send[AudioUnitBusArray](a_.ID, objc.Sel("outputBusses"))
	return rv
}/* debug [instance_properties/getter]: outputBusses */


// An array containing the audio unit’s output connection points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/outputbusses
func (a_ AudioUnitBusArray) SetOutputBusses(value IAUAudioUnitBusArray) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputBusses:"), value)
}/* debug [instance_properties/setter]: outputBusses */


// Determines whether the array can have a variable number of busses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounitbusarray/iscountchangeable
func (a_ AudioUnitBusArray) IsCountChangeable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isCountChangeable"))
	return rv
}/* debug [instance_properties/getter]: isCountChangeable */


// Determines whether the array can have a variable number of busses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounitbusarray/iscountchangeable
func (a_ AudioUnitBusArray) SetIsCountChangeable(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsCountChangeable:"), value)
}/* debug [instance_properties/setter]: isCountChangeable */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AUAudioUnitBusArray */


