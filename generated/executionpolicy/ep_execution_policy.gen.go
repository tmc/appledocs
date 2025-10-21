// Code generated from Apple documentation for ExecutionPolicy. DO NOT EDIT.

package executionpolicy

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [EPExecutionPolicy] class.
var (
	EPExecutionPolicyClass     _EPExecutionPolicyClass
	EPExecutionPolicyClassOnce sync.Once
)

func getEPExecutionPolicyClass() _EPExecutionPolicyClass {
	EPExecutionPolicyClassOnce.Do(func() {
		EPExecutionPolicyClass = _EPExecutionPolicyClass{objc.GetClass("EPExecutionPolicy")}
	})
	return EPExecutionPolicyClass
}

type _EPExecutionPolicyClass struct {
	class objc.Class
}

// An interface definition for the [EPExecutionPolicy] class.
type IEPExecutionPolicy interface {
	objectivec.IObject
	AddPolicyExceptionForURLError(url unsafe.Pointer, error_ unsafe.Pointer) bool
}

//
// [Full Topic]: https://developer.apple.com/documentation/ExecutionPolicy/EPExecutionPolicy
type EPExecutionPolicy struct {
	objectivec.Object
}

// EPExecutionPolicyFrom constructs a [EPExecutionPolicy] from an unsafe.Pointer.
func EPExecutionPolicyFrom(ptr unsafe.Pointer) EPExecutionPolicy {
	return EPExecutionPolicy{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ec _EPExecutionPolicyClass) Alloc() EPExecutionPolicy {
	rv := objc.Send[EPExecutionPolicy](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EPExecutionPolicyClass) New() EPExecutionPolicy {
	rv := objc.Send[EPExecutionPolicy](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EPExecutionPolicy) Init() EPExecutionPolicy {
	rv := objc.Send[EPExecutionPolicy](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EPExecutionPolicy) Autorelease() EPExecutionPolicy {
	rv := objc.Send[EPExecutionPolicy](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEPExecutionPolicy creates a new EPExecutionPolicy instance.
func NewEPExecutionPolicy() EPExecutionPolicy {
	return getEPExecutionPolicyClass().New()
}



//
// [Full Topic]: https://developer.apple.com/documentation/ExecutionPolicy/EPExecutionPolicy/addException(for:)
func (e_ EPExecutionPolicy) AddPolicyExceptionForURLError(url unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("addPolicyExceptionForURL:error:"), url, error_)
	return rv
}


