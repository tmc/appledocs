// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRRVCOperationalStateClusterResumeParams] class.
var (
	MTRRVCOperationalStateClusterResumeParamsClass     _MTRRVCOperationalStateClusterResumeParamsClass
	MTRRVCOperationalStateClusterResumeParamsClassOnce sync.Once
)

func getMTRRVCOperationalStateClusterResumeParamsClass() _MTRRVCOperationalStateClusterResumeParamsClass {
	MTRRVCOperationalStateClusterResumeParamsClassOnce.Do(func() {
		MTRRVCOperationalStateClusterResumeParamsClass = _MTRRVCOperationalStateClusterResumeParamsClass{objc.GetClass("MTRRVCOperationalStateClusterResumeParams")}
	})
	return MTRRVCOperationalStateClusterResumeParamsClass
}

type _MTRRVCOperationalStateClusterResumeParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRRVCOperationalStateClusterResumeParams] class.
type IMTRRVCOperationalStateClusterResumeParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRVCOperationalStateClusterResumeParams
type MTRRVCOperationalStateClusterResumeParams struct {
	objectivec.Object
}

// MTRRVCOperationalStateClusterResumeParamsFrom constructs a [MTRRVCOperationalStateClusterResumeParams] from an unsafe.Pointer.
func MTRRVCOperationalStateClusterResumeParamsFrom(ptr unsafe.Pointer) MTRRVCOperationalStateClusterResumeParams {
	return MTRRVCOperationalStateClusterResumeParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRRVCOperationalStateClusterResumeParamsClass) Alloc() MTRRVCOperationalStateClusterResumeParams {
	rv := objc.Send[MTRRVCOperationalStateClusterResumeParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRRVCOperationalStateClusterResumeParamsClass) New() MTRRVCOperationalStateClusterResumeParams {
	rv := objc.Send[MTRRVCOperationalStateClusterResumeParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRRVCOperationalStateClusterResumeParams) Init() MTRRVCOperationalStateClusterResumeParams {
	rv := objc.Send[MTRRVCOperationalStateClusterResumeParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRRVCOperationalStateClusterResumeParams) Autorelease() MTRRVCOperationalStateClusterResumeParams {
	rv := objc.Send[MTRRVCOperationalStateClusterResumeParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRRVCOperationalStateClusterResumeParams creates a new MTRRVCOperationalStateClusterResumeParams instance.
func NewMTRRVCOperationalStateClusterResumeParams() MTRRVCOperationalStateClusterResumeParams {
	return getMTRRVCOperationalStateClusterResumeParamsClass().New()
}




