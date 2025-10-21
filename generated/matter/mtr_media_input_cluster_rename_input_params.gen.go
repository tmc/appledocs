// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRMediaInputClusterRenameInputParams] class.
var (
	MTRMediaInputClusterRenameInputParamsClass     _MTRMediaInputClusterRenameInputParamsClass
	MTRMediaInputClusterRenameInputParamsClassOnce sync.Once
)

func getMTRMediaInputClusterRenameInputParamsClass() _MTRMediaInputClusterRenameInputParamsClass {
	MTRMediaInputClusterRenameInputParamsClassOnce.Do(func() {
		MTRMediaInputClusterRenameInputParamsClass = _MTRMediaInputClusterRenameInputParamsClass{objc.GetClass("MTRMediaInputClusterRenameInputParams")}
	})
	return MTRMediaInputClusterRenameInputParamsClass
}

type _MTRMediaInputClusterRenameInputParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRMediaInputClusterRenameInputParams] class.
type IMTRMediaInputClusterRenameInputParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaInputClusterRenameInputParams
type MTRMediaInputClusterRenameInputParams struct {
	objectivec.Object
}

// MTRMediaInputClusterRenameInputParamsFrom constructs a [MTRMediaInputClusterRenameInputParams] from an unsafe.Pointer.
func MTRMediaInputClusterRenameInputParamsFrom(ptr unsafe.Pointer) MTRMediaInputClusterRenameInputParams {
	return MTRMediaInputClusterRenameInputParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMediaInputClusterRenameInputParamsClass) Alloc() MTRMediaInputClusterRenameInputParams {
	rv := objc.Send[MTRMediaInputClusterRenameInputParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMediaInputClusterRenameInputParamsClass) New() MTRMediaInputClusterRenameInputParams {
	rv := objc.Send[MTRMediaInputClusterRenameInputParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMediaInputClusterRenameInputParams) Init() MTRMediaInputClusterRenameInputParams {
	rv := objc.Send[MTRMediaInputClusterRenameInputParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMediaInputClusterRenameInputParams) Autorelease() MTRMediaInputClusterRenameInputParams {
	rv := objc.Send[MTRMediaInputClusterRenameInputParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMediaInputClusterRenameInputParams creates a new MTRMediaInputClusterRenameInputParams instance.
func NewMTRMediaInputClusterRenameInputParams() MTRMediaInputClusterRenameInputParams {
	return getMTRMediaInputClusterRenameInputParamsClass().New()
}




