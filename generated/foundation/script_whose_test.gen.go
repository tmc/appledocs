// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ScriptWhoseTest] class.
var (
	scriptWhoseTestClass     _ScriptWhoseTestClass
	scriptWhoseTestClassOnce sync.Once
)

func getScriptWhoseTestClass() _ScriptWhoseTestClass {
	scriptWhoseTestClassOnce.Do(func() {
		scriptWhoseTestClass = _ScriptWhoseTestClass{objc.GetClass("NSScriptWhoseTest")}
	})
	return scriptWhoseTestClass
}

type _ScriptWhoseTestClass struct {
	class objc.Class
}

// An interface definition for the [ScriptWhoseTest] class.
type IScriptWhoseTest interface {
	objectivec.IObject
}

// An abstract class that provides the basis for testing specifiers one at a time or in groups.
//
// is an abstract class whose sole method is . Two concrete subclasses of generate objects representing Boolean expressions comparing one object with another and objects representing multiple Boolean expressions connected by logical operators ( , , ). These classes are, respectively, and . In evaluating itself, an invokes the method of its “test” object. You shouldn’t need to subclass , and you should rarely need to subclass one of its subclasses.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptWhoseTest
type ScriptWhoseTest struct {
	objectivec.Object
}

// ScriptWhoseTestFrom constructs a [ScriptWhoseTest] from an unsafe.Pointer.
//
// An abstract class that provides the basis for testing specifiers one at a time or in groups.
func ScriptWhoseTestFrom(ptr unsafe.Pointer) ScriptWhoseTest {
	return ScriptWhoseTest{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _ScriptWhoseTestClass) Alloc() ScriptWhoseTest {
	rv := objc.Send[ScriptWhoseTest](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _ScriptWhoseTestClass) New() ScriptWhoseTest {
	rv := objc.Send[ScriptWhoseTest](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScriptWhoseTest) Init() ScriptWhoseTest {
	rv := objc.Send[ScriptWhoseTest](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScriptWhoseTest) Autorelease() ScriptWhoseTest {
	rv := objc.Send[ScriptWhoseTest](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScriptWhoseTest creates a new ScriptWhoseTest instance.
func NewScriptWhoseTest() ScriptWhoseTest {
	return getScriptWhoseTestClass().New()
}




