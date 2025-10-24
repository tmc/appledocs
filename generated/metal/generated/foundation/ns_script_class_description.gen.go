// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ScriptClassDescription] class.
var (
	ScriptClassDescriptionClass     _ScriptClassDescriptionClass
	ScriptClassDescriptionClassOnce sync.Once
)

func getScriptClassDescriptionClass() _ScriptClassDescriptionClass {
	ScriptClassDescriptionClassOnce.Do(func() {
		ScriptClassDescriptionClass = _ScriptClassDescriptionClass{objc.GetClass("NSScriptClassDescription")}
	})
	return ScriptClassDescriptionClass
}

type _ScriptClassDescriptionClass struct {
	class objc.Class
}

// An interface definition for the [ScriptClassDescription] class.
type IScriptClassDescription interface {
	IClassDescription
	// properties:
	AppleEventCode() uint32 /* not a class type */
	SetAppleEventCode(value uint32 /* not a class type */)
	ClassName() IString
	SetClassName(value IString)
	DefaultSubcontainerAttributeKey() IString
	SetDefaultSubcontainerAttributeKey(value IString)
	ImplementationClassName() IString
	SetImplementationClassName(value IString)
	SuiteName() IString
	SetSuiteName(value IString)
	Superclass() IScriptClassDescription
	SetSuperclass(value IScriptClassDescription)
	// methods:
}

// A scriptable class that a macOS app supports.
//
// A scriptable application provides scriptability information that describes the commands and objects scripters can use in scripts that target the application. That includes information about the classes those scriptable objects are created from. An application’s scriptability information is collected automatically by an instance of . The registry object creates an for each class it finds and caches these objects in memory. Cocoa scripting uses registry information in handling scripting requests that target the application. A class description instance stores the name, attributes, relationships, and supported commands for a class. For example, a scriptable class for a drawing application might support attributes such as and , relationships such as collections of , , and , and commands such as and . As with many of the classes in Cocoa’s built-in scripting support, your application may never need to directly work with instances of . However, one case where you might need access to a class description is if you override in a scriptable class. For information on how to do this, see in . Another case where your application may need access to class description information is if you override in a specifier class. Although you can subclass , it is unlikely that you would need to do so, or even to create instances of it.


// A scriptable class that a macOS app supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription
type ScriptClassDescription struct {
	ClassDescription
}

// ScriptClassDescriptionFrom constructs a [ScriptClassDescription] from an unsafe.Pointer.
//
// A scriptable class that a macOS app supports.
func ScriptClassDescriptionFrom(ptr unsafe.Pointer) ScriptClassDescription {
	return ScriptClassDescription{
		ClassDescription: ClassDescriptionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _ScriptClassDescriptionClass) Alloc() ScriptClassDescription {
	rv := objc.Send[ScriptClassDescription](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _ScriptClassDescriptionClass) New() ScriptClassDescription {
	rv := objc.Send[ScriptClassDescription](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScriptClassDescription) Init() ScriptClassDescription {
	rv := objc.Send[ScriptClassDescription](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScriptClassDescription) Autorelease() ScriptClassDescription {
	rv := objc.Send[ScriptClassDescription](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScriptClassDescription creates a new ScriptClassDescription instance.
func NewScriptClassDescription() ScriptClassDescription {
	return getScriptClassDescriptionClass().New()
}



// Returns the Apple event code associated with the receiver’s class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptclassdescription/appleeventcode
func (s_ ScriptClassDescription) AppleEventCode() uint32 /* not a class type */ {
	rv := objc.Send[uint32](s_.ID, objc.Sel("appleEventCode"))
	return rv
}


// Returns the Apple event code associated with the receiver’s class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptclassdescription/appleeventcode
func (s_ ScriptClassDescription) SetAppleEventCode(value uint32 /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAppleEventCode:"), value)
}


// Returns the name of the class the receiver describes, as provided at initialization time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptclassdescription/classname
func (s_ ScriptClassDescription) ClassName() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("className"))
	return rv
}


// Returns the name of the class the receiver describes, as provided at initialization time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptclassdescription/classname
func (s_ ScriptClassDescription) SetClassName(value IString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setClassName:"), value)
}


// Returns the value of the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptclassdescription/defaultsubcontainerattributekey
func (s_ ScriptClassDescription) DefaultSubcontainerAttributeKey() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("defaultSubcontainerAttributeKey"))
	return rv
}


// Returns the value of the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptclassdescription/defaultsubcontainerattributekey
func (s_ ScriptClassDescription) SetDefaultSubcontainerAttributeKey(value IString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDefaultSubcontainerAttributeKey:"), value)
}


// Returns the name of the Objective-C class instantiated to implement the scripting class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptclassdescription/implementationclassname
func (s_ ScriptClassDescription) ImplementationClassName() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("implementationClassName"))
	return rv
}


// Returns the name of the Objective-C class instantiated to implement the scripting class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptclassdescription/implementationclassname
func (s_ ScriptClassDescription) SetImplementationClassName(value IString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setImplementationClassName:"), value)
}


// Returns the name of the receiver’s suite.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptclassdescription/suitename
func (s_ ScriptClassDescription) SuiteName() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("suiteName"))
	return rv
}


// Returns the name of the receiver’s suite.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptclassdescription/suitename
func (s_ ScriptClassDescription) SetSuiteName(value IString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSuiteName:"), value)
}


// Returns the class description instance for the superclass of the receiver’s class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptclassdescription/superclass
func (s_ ScriptClassDescription) Superclass() IScriptClassDescription {
	rv := objc.Send[ScriptClassDescription](s_.ID, objc.Sel("superclass"))
	return rv
}


// Returns the class description instance for the superclass of the receiver’s class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptclassdescription/superclass
func (s_ ScriptClassDescription) SetSuperclass(value IScriptClassDescription) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSuperclass:"), value)
}



