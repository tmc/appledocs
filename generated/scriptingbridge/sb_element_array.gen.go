// Code generated from Apple documentation for ScriptingBridge. DO NOT EDIT.

package scriptingbridge

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SBElementArray] class.
var (
	SBElementArrayClass     _SBElementArrayClass
	SBElementArrayClassOnce sync.Once
)

func getSBElementArrayClass() _SBElementArrayClass {
	SBElementArrayClassOnce.Do(func() {
		SBElementArrayClass = _SBElementArrayClass{objc.GetClass("SBElementArray")}
	})
	return SBElementArrayClass
}

type _SBElementArrayClass struct {
	class objc.Class
}

// An interface definition for the [SBElementArray] class.
type ISBElementArray interface {
	foundation.IMutableArray
	ArrayByApplyingSelector(selector objc.SEL) []objc.ID
	ArrayByApplyingSelectorWithObject(aSelector objc.SEL, argument objectivec.IObject) []objc.ID
	Get() []objc.ID
	ObjectAtLocation(location objectivec.IObject) unsafe.Pointer
	ObjectWithID(identifier objectivec.IObject) unsafe.Pointer
	ObjectWithName(name appkit.string) unsafe.Pointer
}

// is subclass of that manages collections of related objects. For example, when you ask the Finder for a list of disks, or ask iTunes for a list of playlists, you get the result back as an containing Scripting Bridge objects representing those items.
//
// defines methods beyond those of for obtaining individual objects. In addition to , also defines , , and .
//
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBElementArray
type SBElementArray struct {
	foundation.MutableArray
}

// SBElementArrayFrom constructs a [SBElementArray] from an unsafe.Pointer.
//
// is subclass of that manages collections of related objects. For example, when you ask the Finder for a list of disks, or ask iTunes for a list of playlists, you get the result back as an containing Scripting Bridge objects representing those items.
func SBElementArrayFrom(ptr unsafe.Pointer) SBElementArray {
	return SBElementArray{
		MutableArray: foundation.MutableArrayFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SBElementArrayClass) Alloc() SBElementArray {
	rv := objc.Send[SBElementArray](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SBElementArrayClass) New() SBElementArray {
	rv := objc.Send[SBElementArray](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SBElementArray) Init() SBElementArray {
	rv := objc.Send[SBElementArray](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SBElementArray) Autorelease() SBElementArray {
	rv := objc.Send[SBElementArray](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSBElementArray creates a new SBElementArray instance.
func NewSBElementArray() SBElementArray {
	return getSBElementArrayClass().New()
}


// Returns an array containing the results of sending the specified message to each object in the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBElementArray/array(byApplying:)
func (s_ SBElementArray) ArrayByApplyingSelector(selector objc.SEL) []objc.ID {
	rv := objc.Send[[]objc.ID](s_.ID, objc.Sel("arrayByApplyingSelector:"), selector)
	return rv
}

// Returns an array containing the results of sending the specified message to each object in the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBElementArray/array(byApplying:with:)
func (s_ SBElementArray) ArrayByApplyingSelectorWithObject(aSelector objc.SEL, argument objectivec.IObject) []objc.ID {
	rv := objc.Send[[]objc.ID](s_.ID, objc.Sel("arrayByApplyingSelector:withObject:"), aSelector, argument)
	return rv
}

// Forces evaluation of the receiver, causing the real object to be returned immediately.
//
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBElementArray/get()
func (s_ SBElementArray) Get() []objc.ID {
	rv := objc.Send[[]objc.ID](s_.ID, objc.Sel("get"))
	return rv
}

// Returns the object at the given location in the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBElementArray/object(atLocation:)
func (s_ SBElementArray) ObjectAtLocation(location objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("objectAtLocation:"), location)
	return rv
}

// Returns the object in the array with the given identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBElementArray/object(withID:)
func (s_ SBElementArray) ObjectWithID(identifier objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("objectWithID:"), identifier)
	return rv
}

// Returns the object in the array with the given name.
//
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBElementArray/object(withName:)
func (s_ SBElementArray) ObjectWithName(name appkit.string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("objectWithName:"), name)
	return rv
}



