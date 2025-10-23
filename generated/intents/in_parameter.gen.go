// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	ParameterClass() unsafe.Pointer
	SetParameterClass(value unsafe.Pointer)
	ParameterKeyPath() string
	SetParameterKeyPath(value string)
	// methods:
}

// A parameter of an interaction object.
//
// Use a parameter object to identify a property of an object. To fetch the value of the property, use the method of the object. You use parameters when configuring a custom user interface for your Siri or Maps interactions. SiriKit passes parameter objects to you during the configuration of your interface. When configuring your interface, you can also create parameter objects to represent properties that you display in addition to the ones that SiriKit provides.


// A parameter of an interaction object.
//
// [Full Topic]
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



// The type of object represented by this parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inparameter/parameterclass
func (i_ INParameter) ParameterClass() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("parameterClass"))
	return rv
}


// The type of object represented by this parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inparameter/parameterclass
func (i_ INParameter) SetParameterClass(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setParameterClass:"), value)
}


// The key path to a property of an interaction object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inparameter/parameterkeypath
func (i_ INParameter) ParameterKeyPath() string {
	rv := objc.Send[string](i_.ID, objc.Sel("parameterKeyPath"))
	return rv
}


// The key path to a property of an interaction object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inparameter/parameterkeypath
func (i_ INParameter) SetParameterKeyPath(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setParameterKeyPath:"), objc.String(value))
}



