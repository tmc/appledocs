// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INParameter] class.
var (
	INParameterClass     _INParameterClass
	INParameterClassOnce sync.Once
)

func getINParameterClass() _INParameterClass {
	INParameterClassOnce.Do(func() {
		INParameterClass = _INParameterClass{objc.GetClass("INParameter")}
	})
	return INParameterClass
}

type _INParameterClass struct {
	class objc.Class
}

// An interface definition for the [INParameter] class.
type IINParameter interface {
	objectivec.IObject
	IndexForSubKeyPath(subKeyPath appkit.string) uint
	IsEqualToParameter(parameter INParameter) bool
	SetIndexForSubKeyPath(index uint, subKeyPath appkit.string)
}

// A parameter of an interaction object.
//
// Use a parameter object to identify a property of an object. To fetch the value of the property, use the method of the object. You use parameters when configuring a custom user interface for your Siri or Maps interactions. SiriKit passes parameter objects to you during the configuration of your interface. When configuring your interface, you can also create parameter objects to represent properties that you display in addition to the ones that SiriKit provides.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INParameter
type INParameter struct {
	objectivec.Object
}

// INParameterFrom constructs a [INParameter] from an unsafe.Pointer.
//
// A parameter of an interaction object.
func INParameterFrom(ptr unsafe.Pointer) INParameter {
	return INParameter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INParameterClass) Alloc() INParameter {
	rv := objc.Send[INParameter](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INParameterClass) New() INParameter {
	rv := objc.Send[INParameter](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INParameter) Init() INParameter {
	rv := objc.Send[INParameter](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INParameter) Autorelease() INParameter {
	rv := objc.Send[INParameter](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINParameter creates a new INParameter instance.
func NewINParameter() INParameter {
	return getINParameterClass().New()
}




// Creates a new parameter object using the specified key path and class information.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INParameter/init(for:keyPath:)
func NewINParameterForClassKeyPath(aClass objc.Class, keyPath appkit.string) INParameter {
	rv := objc.Send[INParameter](objc.ID(getINParameterClass().class), objc.Sel("parameterForClass:keyPath:"), aClass, keyPath)
	return rv
}


// Creates a new parameter object using the specified key path and class information.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INParameter/init(for:keyPath:)
func (ic _INParameterClass) ParameterForClassKeyPath(aClass objc.Class, keyPath appkit.string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("parameterForClass:keyPath:"), aClass, keyPath)
	return rv
}

// The index into the array at the specified portion of the key path.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INParameter/index(forSubKeyPath:)
func (i_ INParameter) IndexForSubKeyPath(subKeyPath appkit.string) uint {
	rv := objc.Send[uint](i_.ID, objc.Sel("indexForSubKeyPath:"), subKeyPath)
	return rv
}

// Returns a Boolean value indicating whether the specified parameter object represents the same property as the current parameter object.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INParameter/isEqual(to:)
func (i_ INParameter) IsEqualToParameter(parameter INParameter) bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isEqualToParameter:"), parameter)
	return rv
}

// Specifies which item of an array or ordered set to use for the parameter.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INParameter/setIndex(_:forSubKeyPath:)
func (i_ INParameter) SetIndexForSubKeyPath(index uint, subKeyPath appkit.string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIndex:forSubKeyPath:"), index, subKeyPath)
}

// The type of object represented by this parameter.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INParameter/parameterClass
func (i_ INParameter) ParameterClass() objc.Class {
	rv := objc.Send[objc.Class](i_.ID, objc.Sel("parameterClass"))
	return rv
}

// The key path to a property of an interaction object.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INParameter/parameterKeyPath
func (i_ INParameter) ParameterKeyPath() appkit.string {
	rv := objc.Send[appkit.string](i_.ID, objc.Sel("parameterKeyPath"))
	return rv
}


