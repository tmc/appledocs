// Code generated from Apple documentation for ExecutionPolicy. DO NOT EDIT.

package executionpolicy

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [EPDeveloperTool] class.
var (
	EPDeveloperToolClass     _EPDeveloperToolClass
	EPDeveloperToolClassOnce sync.Once
)

func getEPDeveloperToolClass() _EPDeveloperToolClass {
	EPDeveloperToolClassOnce.Do(func() {
		EPDeveloperToolClass = _EPDeveloperToolClass{objc.GetClass("EPDeveloperTool")}
	})
	return EPDeveloperToolClass
}

type _EPDeveloperToolClass struct {
	class objc.Class
}

// An interface definition for the [EPDeveloperTool] class.
type IEPDeveloperTool interface {
	objectivec.IObject
	// properties:
	AuthorizationStatus() EPDeveloperToolStatus
	// methods:
	RequestDeveloperToolAccessWithCompletionHandler(handler unsafe.Pointer)
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExecutionPolicy/EPDeveloperTool
type EPDeveloperTool struct {
	objectivec.Object
}

// EPDeveloperToolFrom constructs a [EPDeveloperTool] from an unsafe.Pointer.
func EPDeveloperToolFrom(ptr unsafe.Pointer) EPDeveloperTool {
	return EPDeveloperTool{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ec _EPDeveloperToolClass) Alloc() EPDeveloperTool {
	rv := objc.Send[EPDeveloperTool](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EPDeveloperToolClass) New() EPDeveloperTool {
	rv := objc.Send[EPDeveloperTool](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EPDeveloperTool) Init() EPDeveloperTool {
	rv := objc.Send[EPDeveloperTool](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EPDeveloperTool) Autorelease() EPDeveloperTool {
	rv := objc.Send[EPDeveloperTool](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEPDeveloperTool creates a new EPDeveloperTool instance.
func NewEPDeveloperTool() EPDeveloperTool {
	return getEPDeveloperToolClass().New()
}




// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExecutionPolicy/EPDeveloperTool/requestAccess(completionHandler:)
func (e_ EPDeveloperTool) RequestDeveloperToolAccessWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("requestDeveloperToolAccessWithCompletionHandler:"), handler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExecutionPolicy/EPDeveloperTool/authorizationStatus
func (e_ EPDeveloperTool) AuthorizationStatus() EPDeveloperToolStatus {
	rv := objc.Send[EPDeveloperToolStatus](e_.ID, objc.Sel("authorizationStatus"))
	return rv
}


