// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class AUParameter */


/* debug [class_header]: Header for AUParameter */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Parameter */
// An interface definition for the [Parameter] class.
type IParameter interface {
	IParameterNode
	
/* debug [class_interface_properties]: Properties for Parameter */
	// properties:
	Address() ParameterAddress /* typedef */
	DependentParameters() []foundation.Number
	Flags() AudioUnitParameterOptions
	MaxValue() foundation.Value
	MinValue() foundation.Value
	Unit() AudioUnitParameterUnit
	UnitName() objc.IObject /* cross-framework: NSString */
	Value() foundation.Value
	SetValue(value foundation.Value)
	ValueStrings() []string
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Parameter */
	// methods:
	SetValueOriginator(value foundation.Value, originator ParameterObserverToken /* typedef */)
	SetValueOriginatorAtHostTime(value foundation.Value, originator ParameterObserverToken /* typedef */, hostTime uint64)
	SetValueOriginatorAtHostTimeEventType(value foundation.Value, originator ParameterObserverToken /* typedef */, hostTime uint64, eventType ParameterAutomationEventType)
	StringFromValue(value Value /* typedef */) foundation.String
	ValueFromString(string_ objc.IObject /* cross-framework: NSString */) foundation.Value
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Parameter */
// Alloc allocates a new instance without initialization.
func (pc _ParameterClass) Alloc() Parameter {
	rv := objc.Send[Parameter](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Parameter */
// An object that represents a single audio unit parameter.


// An object that represents a single audio unit parameter.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Parameter *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Parameter */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Parameter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Parameter */

// Sets the parameter’s value, avoiding redundant notifications to the originator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameter/setValue(_:originator:)
func (p_ Parameter) SetValueOriginator(value foundation.Value, originator ParameterObserverToken /* typedef */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setValue:originator:"), value, originator)
}/* debug [instance_methods/method]: SetValueOriginator */


// Sets the parameter’s value, preserving the host time of the gesture that initiated the change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameter/setValue(_:originator:atHostTime:)
func (p_ Parameter) SetValueOriginatorAtHostTime(value foundation.Value, originator ParameterObserverToken /* typedef */, hostTime uint64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setValue:originator:atHostTime:"), value, originator, hostTime)
}/* debug [instance_methods/method]: SetValueOriginatorAtHostTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameter/setValue(_:originator:atHostTime:eventType:)
func (p_ Parameter) SetValueOriginatorAtHostTimeEventType(value foundation.Value, originator ParameterObserverToken /* typedef */, hostTime uint64, eventType ParameterAutomationEventType) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setValue:originator:atHostTime:eventType:"), value, originator, hostTime, eventType)
}/* debug [instance_methods/method]: SetValueOriginatorAtHostTimeEventType */


// Gets the string representation of a parameter value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameter/string(fromValue:)
func (p_ Parameter) StringFromValue(value Value /* typedef */) foundation.String {
	rv := objc.Send[foundation.String](p_.ID, objc.Sel("stringFromValue:"), value)
	return rv
}/* debug [instance_methods/method]: StringFromValue */


// Converts a string into a parameter value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameter/value(from:)
func (p_ Parameter) ValueFromString(string_ objc.IObject /* cross-framework: NSString */) foundation.Value {
	rv := objc.Send[foundation.Value](p_.ID, objc.Sel("valueFromString:"), string_)
	return rv
}/* debug [instance_methods/method]: ValueFromString */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Parameter */

// The parameter’s address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameter/address
func (p_ Parameter) Address() ParameterAddress /* typedef */ {
	rv := objc.Send[uint64](p_.ID, objc.Sel("address"))
	return rv
}/* debug [instance_properties/getter]: address */


// Any other parameter’s whose values may change as a side effect of this parameter’s value changing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameter/dependentParameters
func (p_ Parameter) DependentParameters() []foundation.Number {
	rv := objc.Send[[]foundation.Number](p_.ID, objc.Sel("dependentParameters"))
	return rv
}/* debug [instance_properties/getter]: dependentParameters */


// The parameter’s characteristic details.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameter/flags
func (p_ Parameter) Flags() AudioUnitParameterOptions {
	rv := objc.Send[AudioUnitParameterOptions](p_.ID, objc.Sel("flags"))
	return rv
}/* debug [instance_properties/getter]: flags */


// The parameter’s maximum value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameter/maxValue
func (p_ Parameter) MaxValue() foundation.Value {
	rv := objc.Send[foundation.Value](p_.ID, objc.Sel("maxValue"))
	return rv
}/* debug [instance_properties/getter]: maxValue */


// The parameter’s minimum value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameter/minValue
func (p_ Parameter) MinValue() foundation.Value {
	rv := objc.Send[foundation.Value](p_.ID, objc.Sel("minValue"))
	return rv
}/* debug [instance_properties/getter]: minValue */


// The parameter’s unit of measurement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameter/unit
func (p_ Parameter) Unit() AudioUnitParameterUnit {
	rv := objc.Send[AudioUnitParameterUnit](p_.ID, objc.Sel("unit"))
	return rv
}/* debug [instance_properties/getter]: unit */


// The parameter’s localized unit name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameter/unitName
func (p_ Parameter) UnitName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("unitName"))
	return rv
}/* debug [instance_properties/getter]: unitName */


// The parameter’s current value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameter/value
func (p_ Parameter) Value() foundation.Value {
	rv := objc.Send[foundation.Value](p_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// The parameter’s current value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameter/value
func (p_ Parameter) SetValue(value foundation.Value) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setValue:"), value)
}/* debug [instance_properties/setter]: value */


// The parameter’s localized value strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameter/valueStrings
func (p_ Parameter) ValueStrings() []string {
	rv := objc.Send[[]string](p_.ID, objc.Sel("valueStrings"))
	return rv
}/* debug [instance_properties/getter]: valueStrings */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AUParameter */



