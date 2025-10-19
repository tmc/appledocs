// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [LogicalTest] class.
var logicalTestClass = _LogicalTestClass{objc.GetClass("NSLogicalTest")}

type _LogicalTestClass struct {
	class objc.Class
}

// An interface definition for the [LogicalTest] class.
type ILogicalTest interface {
	IScriptWhoseTest
}

// The logical combination of one or more specifier tests. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLogicalTest

type LogicalTest struct {
	ScriptWhoseTest
}

// LogicalTestFrom constructs a [LogicalTest] from an unsafe.Pointer.
//
// The logical combination of one or more specifier tests.
func LogicalTestFrom(ptr unsafe.Pointer) LogicalTest {
	return LogicalTest{
		ScriptWhoseTest: ScriptWhoseTestFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (lc _LogicalTestClass) Alloc() LogicalTest {
	rv := objc.Send[LogicalTest](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (lc _LogicalTestClass) New() LogicalTest {
	rv := objc.Send[LogicalTest](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LogicalTest) Init() LogicalTest {
	rv := objc.Send[LogicalTest](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LogicalTest) Autorelease() LogicalTest {
	rv := objc.Send[LogicalTest](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLogicalTest creates a new LogicalTest instance.
func NewLogicalTest() LogicalTest {
	return logicalTestClass.New()
}




