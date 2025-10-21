// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [Parameter] class.
var (
	ParameterClass     _ParameterClass
	ParameterClassOnce sync.Once
)

func getParameterClass() _ParameterClass {
	ParameterClassOnce.Do(func() {
		ParameterClass = _ParameterClass{objc.GetClass("AUParameter")}
	})
	return ParameterClass
}

type _ParameterClass struct {
	class objc.Class
}

// An interface definition for the [Parameter] class.
type IParameter interface {
	IParameterNode
	SetValueOriginator(value IValue, originator IParameterObserverToken)
	SetValueOriginatorAtHostTime(value IValue, originator IParameterObserverToken, hostTime uint64)
	SetValueOriginatorAtHostTimeEventType(value IValue, originator IParameterObserverToken, hostTime uint64, eventType ParameterAutomationEventType)
	StringFromValue(value IAUValue) foundation.String
	ValueFromString(string_ appkit.string) Value
}

// An object that represents a single audio unit parameter.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameter
type Parameter struct {
	ParameterNode
}

// ParameterFrom constructs a [Parameter] from an unsafe.Pointer.
//
// An object that represents a single audio unit parameter.
func ParameterFrom(ptr unsafe.Pointer) Parameter {
	return Parameter{
		ParameterNode: ParameterNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _ParameterClass) Alloc() Parameter {
	rv := objc.Send[Parameter](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _ParameterClass) New() Parameter {
	rv := objc.Send[Parameter](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ Parameter) Init() Parameter {
	rv := objc.Send[Parameter](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ Parameter) Autorelease() Parameter {
	rv := objc.Send[Parameter](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewParameter creates a new Parameter instance.
func NewParameter() Parameter {
	return getParameterClass().New()
}


// Sets the parameter’s value, avoiding redundant notifications to the originator.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameter/setValue(_:originator:)
func (p_ Parameter) SetValueOriginator(value IValue, originator IParameterObserverToken) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setValue:originator:"), value, originator)
}

// Sets the parameter’s value, preserving the host time of the gesture that initiated the change.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameter/setValue(_:originator:atHostTime:)
func (p_ Parameter) SetValueOriginatorAtHostTime(value IValue, originator IParameterObserverToken, hostTime uint64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setValue:originator:atHostTime:"), value, originator, hostTime)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameter/setValue(_:originator:atHostTime:eventType:)
func (p_ Parameter) SetValueOriginatorAtHostTimeEventType(value IValue, originator IParameterObserverToken, hostTime uint64, eventType ParameterAutomationEventType) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setValue:originator:atHostTime:eventType:"), value, originator, hostTime, eventType)
}

// Gets the string representation of a parameter value.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameter/string(fromValue:)
func (p_ Parameter) StringFromValue(value IAUValue) foundation.String {
	rv := objc.Send[foundation.String](p_.ID, objc.Sel("stringFromValue:"), value)
	return rv
}

// Converts a string into a parameter value.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameter/value(from:)
func (p_ Parameter) ValueFromString(string_ appkit.string) Value {
	rv := objc.Send[Value](p_.ID, objc.Sel("valueFromString:"), string_)
	return rv
}

// The parameter’s address.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameter/address
func (p_ Parameter) Address() ParameterAddress {
	rv := objc.Send[ParameterAddress](p_.ID, objc.Sel("address"))
	return rv
}

// Any other parameter’s whose values may change as a side effect of this parameter’s value changing.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameter/dependentParameters
func (p_ Parameter) DependentParameters() []foundation.Number {
	rv := objc.Send[[]foundation.Number](p_.ID, objc.Sel("dependentParameters"))
	return rv
}

// The parameter’s characteristic details.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameter/flags
func (p_ Parameter) Flags() AudioUnitParameterOptions {
	rv := objc.Send[AudioUnitParameterOptions](p_.ID, objc.Sel("flags"))
	return rv
}

// The parameter’s maximum value.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameter/maxValue
func (p_ Parameter) MaxValue() Value {
	rv := objc.Send[Value](p_.ID, objc.Sel("maxValue"))
	return rv
}

// The parameter’s minimum value.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameter/minValue
func (p_ Parameter) MinValue() Value {
	rv := objc.Send[Value](p_.ID, objc.Sel("minValue"))
	return rv
}

// The parameter’s unit of measurement.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameter/unit
func (p_ Parameter) Unit() AudioUnitParameterUnit {
	rv := objc.Send[AudioUnitParameterUnit](p_.ID, objc.Sel("unit"))
	return rv
}

// The parameter’s localized unit name.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameter/unitName
func (p_ Parameter) UnitName() appkit.string {
	rv := objc.Send[appkit.string](p_.ID, objc.Sel("unitName"))
	return rv
}

// The parameter’s current value.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameter/value
func (p_ Parameter) Value() Value {
	rv := objc.Send[Value](p_.ID, objc.Sel("value"))
	return rv
}


// SetValue sets the value of the value property.
// The parameter’s current value.

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameter/value
func (p_ Parameter) SetValue(value IValue) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setValue:"), value)
}

// The parameter’s localized value strings.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameter/valueStrings
func (p_ Parameter) ValueStrings() []string {
	rv := objc.Send[[]string](p_.ID, objc.Sel("valueStrings"))
	return rv
}



