// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AppleEventDescriptor] class.
var (
	appleEventDescriptorClass     _AppleEventDescriptorClass
	appleEventDescriptorClassOnce sync.Once
)

func getAppleEventDescriptorClass() _AppleEventDescriptorClass {
	appleEventDescriptorClassOnce.Do(func() {
		appleEventDescriptorClass = _AppleEventDescriptorClass{objc.GetClass("NSAppleEventDescriptor")}
	})
	return appleEventDescriptorClass
}

type _AppleEventDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [AppleEventDescriptor] class.
type IAppleEventDescriptor interface {
	objectivec.IObject
}

// A wrapper for the Apple event descriptor data type.
//
// An instance of represents a descriptor—the basic building block for Apple events. This class is a wrapper for the underlying Apple event descriptor data type, . Scriptable Cocoa applications frequently work with instances of , but should rarely need to work directly with the data structure. A is a data structure that stores data and an accompanying four-character code. A descriptor can store a value, or it can store a list of other descriptors (which may also be lists). All the information in an Apple event is stored in descriptors and lists of descriptors, and every Apple event is itself a descriptor list that matches certain criteria. Descriptors can be used to build arbitrarily complex containers, so that one Apple event can represent a script statement such as . In working with Apple event descriptors, it can be useful to understand some of the underlying data types. You’ll find terms such as descriptor, descriptor list, Apple event record, and Apple event defined in Building an Apple Event in Apple Events Programming Guide. You’ll also find information on the four-character codes used to identify information within a descriptor. Apple event data types are defined in . The values of many four-character codes used by Apple (and in some cases reused by developers) can be found in . The most common reason to construct an Apple event with an instance of is to supply information in a return Apple event. The most common situation where you might need to extract information from an Apple event (as an instance of ) is when an Apple event handler installed by your application is invoked, as described in “Installing an Apple Event Handler” in . In addition, if you execute an AppleScript script using the class, you get an instance of as the return value, from which you can extract any required information. When you work with an instance of , you can access the underlying descriptor directly, if necessary, with the method. Other methods, including make it possible to create and initialize instances of without creating temporary instances of . The designated initializer for is . However, it is unlikely that you will need to create a subclass of . Cocoa doesn’t currently provide a mechanism for applications to directly send raw Apple events (though compiling and executing an AppleScript script with may result in Apple events being sent). However, Cocoa applications have full access to the Apple Event Manager C APIs for working with Apple events. So, for example, you might use an instance of to assemble an Apple event and call the Apple Event Manager function to send it. If you need to send Apple events, or if you need more information on some of the Apple event concepts described here, see Apple Events Programming Guide and .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor
type AppleEventDescriptor struct {
	objectivec.Object
}

// AppleEventDescriptorFrom constructs a [AppleEventDescriptor] from an unsafe.Pointer.
//
// A wrapper for the Apple event descriptor data type.
func AppleEventDescriptorFrom(ptr unsafe.Pointer) AppleEventDescriptor {
	return AppleEventDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AppleEventDescriptorClass) Alloc() AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AppleEventDescriptorClass) New() AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AppleEventDescriptor) Init() AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AppleEventDescriptor) Autorelease() AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAppleEventDescriptor creates a new AppleEventDescriptor instance.
func NewAppleEventDescriptor() AppleEventDescriptor {
	return getAppleEventDescriptorClass().New()
}




