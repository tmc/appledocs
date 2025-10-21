// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SpecifierTest] class.
var (
	SpecifierTestClass     _SpecifierTestClass
	SpecifierTestClassOnce sync.Once
)

func getSpecifierTestClass() _SpecifierTestClass {
	SpecifierTestClassOnce.Do(func() {
		SpecifierTestClass = _SpecifierTestClass{objc.GetClass("NSSpecifierTest")}
	})
	return SpecifierTestClass
}

type _SpecifierTestClass struct {
	class objc.Class
}

// An interface definition for the [SpecifierTest] class.
type ISpecifierTest interface {
	IScriptWhoseTest
}

// A comparison between an object specifier and a test object.
//
// Instances of this class represent a Boolean expression; they evaluate an object specifier and compare the resulting object to another object using a given comparison method. For more information on , see the method description for its sole public method, its initializer, . When an object is properly initialized, it holds two objects: A “value” or “test” object used as the basis of the comparison; this object can be a regular object or object specifier (such as “blue” in “words whose color is blue”). An object specifier evaluating to the container (“words”). The instance also encapsulates a selector identifying the method performing this comparison. The informal protocol defines a set of comparison methods useful for this purpose, while describes additional methods you may need to use for scripting. The test object is compared, using the selector, against each object in the container. Specifiers in these tests usually have invoked on their topmost container. You should rarely need to subclass .
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getSpecifierTestClass().New()
}




// Returns a specifier test initialized to evaluate a test object against an object specified by an object specifier using a given comparison operation.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSpecifierTest/init(objectSpecifier:comparisonOperator:test:)
func NewSpecifierTestWithObjectSpecifierComparisonOperatorTestObject(obj1 unsafe.Pointer, compOp unsafe.Pointer, obj2 objc.ID) SpecifierTest {
	instance := getSpecifierTestClass().Alloc()
	rv := objc.Send[SpecifierTest](instance.ID, objc.Sel("initWithObjectSpecifier:comparisonOperator:testObject:"), obj1, compOp, obj2)
	rv.Autorelease()
	return rv
}



