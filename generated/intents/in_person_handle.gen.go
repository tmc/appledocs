// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INPersonHandle] class.
var (
	INPersonHandleClass     _INPersonHandleClass
	INPersonHandleClassOnce sync.Once
)

func getINPersonHandleClass() _INPersonHandleClass {
	INPersonHandleClassOnce.Do(func() {
		INPersonHandleClass = _INPersonHandleClass{objc.GetClass("INPersonHandle")}
	})
	return INPersonHandleClass
}

type _INPersonHandleClass struct {
	class objc.Class
}

// An interface definition for the [INPersonHandle] class.
type IINPersonHandle interface {
	objectivec.IObject
	Label() unsafe.Pointer
	SetLabel(value unsafe.Pointer)
	Type() unsafe.Pointer
	SetType(value unsafe.Pointer)
	Value() string
	SetValue(value string)
}

// The identifying information for a user of your app.
//
// An object contains information that you use to uniquely identify a user of your app. When resolving a person associated with an intent, you might create instances of this class and add them to an object when resolving an intent involving that person. Handles contain unique information such as an email address or phone number.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INPersonHandle
type INPersonHandle struct {
	objectivec.Object
}

// INPersonHandleFrom constructs a [INPersonHandle] from an unsafe.Pointer.
//
// The identifying information for a user of your app.
func INPersonHandleFrom(ptr unsafe.Pointer) INPersonHandle {
	return INPersonHandle{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INPersonHandleClass) Alloc() INPersonHandle {
	rv := objc.Send[INPersonHandle](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INPersonHandleClass) New() INPersonHandle {
	rv := objc.Send[INPersonHandle](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INPersonHandle) Init() INPersonHandle {
	rv := objc.Send[INPersonHandle](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INPersonHandle) Autorelease() INPersonHandle {
	rv := objc.Send[INPersonHandle](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINPersonHandle creates a new INPersonHandle instance.
func NewINPersonHandle() INPersonHandle {
	return getINPersonHandleClass().New()
}


// A standard label that describes the meaning of the information.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inpersonhandle/label
func (i_ INPersonHandle) Label() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
// A standard label that describes the meaning of the information.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inpersonhandle/label
func (i_ INPersonHandle) SetLabel(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setLabel:"), value)
}

// The type of information contained in the handle.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inpersonhandle/type
func (i_ INPersonHandle) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("type"))
	return rv
}


// SetType sets the value of the type property.
// The type of information contained in the handle.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inpersonhandle/type
func (i_ INPersonHandle) SetType(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setType:"), value)
}

// The data for the handle.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inpersonhandle/value
func (i_ INPersonHandle) Value() string {
	rv := objc.Send[string](i_.ID, objc.Sel("value"))
	return rv
}


// SetValue sets the value of the value property.
// The data for the handle.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inpersonhandle/value
func (i_ INPersonHandle) SetValue(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setValue:"), objc.String(value))
}



