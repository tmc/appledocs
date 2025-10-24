// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRThermostatClusterScheduleStruct] class.
var (
	MTRThermostatClusterScheduleStructClass     _MTRThermostatClusterScheduleStructClass
	MTRThermostatClusterScheduleStructClassOnce sync.Once
)

func getMTRThermostatClusterScheduleStructClass() _MTRThermostatClusterScheduleStructClass {
	MTRThermostatClusterScheduleStructClassOnce.Do(func() {
		MTRThermostatClusterScheduleStructClass = _MTRThermostatClusterScheduleStructClass{objc.GetClass("MTRThermostatClusterScheduleStruct")}
	})
	return MTRThermostatClusterScheduleStructClass
}

type _MTRThermostatClusterScheduleStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRThermostatClusterScheduleStruct] class.
type IMTRThermostatClusterScheduleStruct interface {
	objectivec.IObject
	// properties:
	BuiltIn() objc.IObject /* cross-framework: NSNumber */
	SetBuiltIn(value objc.IObject /* cross-framework: NSNumber */)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	PresetHandle() objc.IObject /* cross-framework: NSData */
	SetPresetHandle(value objc.IObject /* cross-framework: NSData */)
	ScheduleHandle() objc.IObject /* cross-framework: NSData */
	SetScheduleHandle(value objc.IObject /* cross-framework: NSData */)
	SystemMode() objc.IObject /* cross-framework: NSNumber */
	SetSystemMode(value objc.IObject /* cross-framework: NSNumber */)
	Transitions() objc.IObject /* cross-framework: NSArray */
	SetTransitions(value objc.IObject /* cross-framework: NSArray */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleStruct
type MTRThermostatClusterScheduleStruct struct {
	objectivec.Object
}

// MTRThermostatClusterScheduleStructFrom constructs a [MTRThermostatClusterScheduleStruct] from an unsafe.Pointer.
func MTRThermostatClusterScheduleStructFrom(ptr unsafe.Pointer) MTRThermostatClusterScheduleStruct {
	return MTRThermostatClusterScheduleStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThermostatClusterScheduleStructClass) Alloc() MTRThermostatClusterScheduleStruct {
	rv := objc.Send[MTRThermostatClusterScheduleStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThermostatClusterScheduleStructClass) New() MTRThermostatClusterScheduleStruct {
	rv := objc.Send[MTRThermostatClusterScheduleStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThermostatClusterScheduleStruct) Init() MTRThermostatClusterScheduleStruct {
	rv := objc.Send[MTRThermostatClusterScheduleStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThermostatClusterScheduleStruct) Autorelease() MTRThermostatClusterScheduleStruct {
	rv := objc.Send[MTRThermostatClusterScheduleStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThermostatClusterScheduleStruct creates a new MTRThermostatClusterScheduleStruct instance.
func NewMTRThermostatClusterScheduleStruct() MTRThermostatClusterScheduleStruct {
	return getMTRThermostatClusterScheduleStructClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleStruct/builtIn
func (m_ MTRThermostatClusterScheduleStruct) BuiltIn() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("builtIn"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleStruct/builtIn
func (m_ MTRThermostatClusterScheduleStruct) SetBuiltIn(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBuiltIn:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleStruct/name
func (m_ MTRThermostatClusterScheduleStruct) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleStruct/name
func (m_ MTRThermostatClusterScheduleStruct) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleStruct/presetHandle
func (m_ MTRThermostatClusterScheduleStruct) PresetHandle() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("presetHandle"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleStruct/presetHandle
func (m_ MTRThermostatClusterScheduleStruct) SetPresetHandle(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPresetHandle:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleStruct/scheduleHandle
func (m_ MTRThermostatClusterScheduleStruct) ScheduleHandle() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("scheduleHandle"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleStruct/scheduleHandle
func (m_ MTRThermostatClusterScheduleStruct) SetScheduleHandle(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setScheduleHandle:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleStruct/systemMode
func (m_ MTRThermostatClusterScheduleStruct) SystemMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("systemMode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleStruct/systemMode
func (m_ MTRThermostatClusterScheduleStruct) SetSystemMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSystemMode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleStruct/transitions
func (m_ MTRThermostatClusterScheduleStruct) Transitions() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("transitions"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleStruct/transitions
func (m_ MTRThermostatClusterScheduleStruct) SetTransitions(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitions:"), value)
}



