// Code generated from Apple documentation for ScriptingBridge. DO NOT EDIT.

package scriptingbridge

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SBObject] class.
var (
	SBObjectClass     _SBObjectClass
	SBObjectClassOnce sync.Once
)

func getSBObjectClass() _SBObjectClass {
	SBObjectClassOnce.Do(func() {
		SBObjectClass = _SBObjectClass{objc.GetClass("SBObject")}
	})
	return SBObjectClass
}

type _SBObjectClass struct {
	class objc.Class
}

// An interface definition for the [SBObject] class.
type ISBObject interface {
	objectivec.IObject
	ElementArrayWithCode(code unsafe.Pointer) SBElementArray
	Get() objc.ID
	LastError() foundation.Error
	PropertyWithClassCode(cls objc.Class, code unsafe.Pointer) SBObject
	PropertyWithCode(code unsafe.Pointer) SBObject
	SendEventIdParameters(eventClass unsafe.Pointer, eventID unsafe.Pointer, firstParamCode unsafe.Pointer) objc.ID
	SetTo(value objectivec.IObject)
}

// The class declares methods that can be invoked on any object in a scriptable application. It defines methods for getting elements and properties of an object, as well as setting a given object to a new value.
//
// Each is built around an object specifier, which tells Scripting Bridge how to locate the object. Therefore, you can think of an as a reference to an object in an target application rather than an object itself. To bypass this reference-based approach and force evaluation, use the method. Typically, rather than create instances explictly, you receive objects by calling methods of an subclass. For example, if you wanted to get an representing the current iTunes track, you would use code like this (where is a subclass of ): You can discover the names of dynamically generated classes such as and by examining the header file created by the tool. Alternatively, you give these variables the dynamic Objective-C type .
//
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBObject
type SBObject struct {
	objectivec.Object
}

// SBObjectFrom constructs a [SBObject] from an unsafe.Pointer.
//
// The class declares methods that can be invoked on any object in a scriptable application. It defines methods for getting elements and properties of an object, as well as setting a given object to a new value.
func SBObjectFrom(ptr unsafe.Pointer) SBObject {
	return SBObject{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SBObjectClass) Alloc() SBObject {
	rv := objc.Send[SBObject](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SBObjectClass) New() SBObject {
	rv := objc.Send[SBObject](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SBObject) Init() SBObject {
	rv := objc.Send[SBObject](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SBObject) Autorelease() SBObject {
	rv := objc.Send[SBObject](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSBObject creates a new SBObject instance.
func NewSBObject() SBObject {
	return getSBObjectClass().New()
}




// Returns an instance of an subclass initialized with the given data.
//
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBObject/init(data:)
func NewSBObjectWithData(data objectivec.IObject) SBObject {
	instance := getSBObjectClass().Alloc()
	rv := objc.Send[SBObject](instance.ID, objc.Sel("initWithData:"), data)
	rv.Autorelease()
	return rv
}



// Returns an instance of an subclass initialized with the specified properties and data and added to the designated element array.
//
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBObject/init(elementCode:properties:data:)
func NewSBObjectWithElementCodePropertiesData(code unsafe.Pointer, properties unsafe.Pointer, data objectivec.IObject) SBObject {
	instance := getSBObjectClass().Alloc()
	rv := objc.Send[SBObject](instance.ID, objc.Sel("initWithElementCode:properties:data:"), code, properties, data)
	rv.Autorelease()
	return rv
}



// Returns an instance of an subclass initialized with the specified properties.
//
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBObject/init(properties:)
func NewSBObjectWithProperties(properties objectivec.IObject) SBObject {
	instance := getSBObjectClass().Alloc()
	rv := objc.Send[SBObject](instance.ID, objc.Sel("initWithProperties:"), properties)
	rv.Autorelease()
	return rv
}


// Returns an array containing every child of the receiver with the given class-type code.
//
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBObject/elementArray(withCode:)
func (s_ SBObject) ElementArrayWithCode(code unsafe.Pointer) SBElementArray {
	rv := objc.Send[SBElementArray](s_.ID, objc.Sel("elementArrayWithCode:"), code)
	return rv
}

// Forces evaluation of the receiver, causing the real object to be returned immediately.
//
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBObject/get()
func (s_ SBObject) Get() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("get"))
	return rv
}

// The error from the last event this object sent, or nil if it succeeded.
//
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBObject/lastError()
func (s_ SBObject) LastError() foundation.Error {
	rv := objc.Send[foundation.Error](s_.ID, objc.Sel("lastError"))
	return rv
}

// Returns an object of the designated scripting class representing the specified property of the receiver
//
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBObject/property(with:code:)
func (s_ SBObject) PropertyWithClassCode(cls objc.Class, code unsafe.Pointer) SBObject {
	rv := objc.Send[SBObject](s_.ID, objc.Sel("propertyWithClass:code:"), cls, code)
	return rv
}

// Returns an object representing the specified property of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBObject/property(withCode:)
func (s_ SBObject) PropertyWithCode(code unsafe.Pointer) SBObject {
	rv := objc.Send[SBObject](s_.ID, objc.Sel("propertyWithCode:"), code)
	return rv
}

// Sends an Apple event with the given event class, event ID, and format to the target application.
//
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBObject/sendEvent:id:parameters:
func (s_ SBObject) SendEventIdParameters(eventClass unsafe.Pointer, eventID unsafe.Pointer, firstParamCode unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("sendEvent:id:parameters:"), eventClass, eventID, firstParamCode)
	return rv
}

// Sets the receiver to a specified value.
//
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBObject/setTo(_:)
func (s_ SBObject) SetTo(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTo:"), value)
}


