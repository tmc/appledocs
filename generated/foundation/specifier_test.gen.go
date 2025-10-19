// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SpecifierTest] class.
var specifierTestClass = _SpecifierTestClass{objc.GetClass("NSSpecifierTest")}

type _SpecifierTestClass struct {
	class objc.Class
}

// An interface definition for the [SpecifierTest] class.
type ISpecifierTest interface {
	IScriptWhoseTest
}

// A comparison between an object specifier and a test object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSpecifierTest

type SpecifierTest struct {
	ScriptWhoseTest
}

// SpecifierTestFrom constructs a [SpecifierTest] from an unsafe.Pointer.
//
// A comparison between an object specifier and a test object.
func SpecifierTestFrom(ptr unsafe.Pointer) SpecifierTest {
	return SpecifierTest{
		ScriptWhoseTest: ScriptWhoseTestFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (sc _SpecifierTestClass) Alloc() SpecifierTest {
	rv := objc.Send[SpecifierTest](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (sc _SpecifierTestClass) New() SpecifierTest {
	rv := objc.Send[SpecifierTest](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SpecifierTest) Init() SpecifierTest {
	rv := objc.Send[SpecifierTest](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SpecifierTest) Autorelease() SpecifierTest {
	rv := objc.Send[SpecifierTest](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSpecifierTest creates a new SpecifierTest instance.
func NewSpecifierTest() SpecifierTest {
	return specifierTestClass.New()
}




