// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRCommandPath] class.
var (
	MTRCommandPathClass     _MTRCommandPathClass
	MTRCommandPathClassOnce sync.Once
)

func getMTRCommandPathClass() _MTRCommandPathClass {
	MTRCommandPathClassOnce.Do(func() {
		MTRCommandPathClass = _MTRCommandPathClass{objc.GetClass("MTRCommandPath")}
	})
	return MTRCommandPathClass
}

type _MTRCommandPathClass struct {
	class objc.Class
}

// An interface definition for the [MTRCommandPath] class.
type IMTRCommandPath interface {
	IMTRClusterPath
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommandPath
type MTRCommandPath struct {
	MTRClusterPath
}

// MTRCommandPathFrom constructs a [MTRCommandPath] from an unsafe.Pointer.
func MTRCommandPathFrom(ptr unsafe.Pointer) MTRCommandPath {
	return MTRCommandPath{
		MTRClusterPath: MTRClusterPathFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRCommandPathClass) Alloc() MTRCommandPath {
	rv := objc.Send[MTRCommandPath](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRCommandPathClass) New() MTRCommandPath {
	rv := objc.Send[MTRCommandPath](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRCommandPath) Init() MTRCommandPath {
	rv := objc.Send[MTRCommandPath](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRCommandPath) Autorelease() MTRCommandPath {
	rv := objc.Send[MTRCommandPath](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRCommandPath creates a new MTRCommandPath instance.
func NewMTRCommandPath() MTRCommandPath {
	return getMTRCommandPathClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommandpath/command
func (m_ MTRCommandPath) Command() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("command"))
	return rv
}


// SetCommand sets the value of the command property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommandpath/command
func (m_ MTRCommandPath) SetCommand(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCommand:"), value)
}



