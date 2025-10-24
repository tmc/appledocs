// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRCertificates] class.
var (
	MTRCertificatesClass     _MTRCertificatesClass
	MTRCertificatesClassOnce sync.Once
)

func getMTRCertificatesClass() _MTRCertificatesClass {
	MTRCertificatesClassOnce.Do(func() {
		MTRCertificatesClass = _MTRCertificatesClass{objc.GetClass("MTRCertificates")}
	})
	return MTRCertificatesClass
}

type _MTRCertificatesClass struct {
	class objc.Class
}

// An interface definition for the [MTRCertificates] class.
type IMTRCertificates interface {
	objectivec.IObject
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCertificates
type MTRCertificates struct {
	objectivec.Object
}

// MTRCertificatesFrom constructs a [MTRCertificates] from an unsafe.Pointer.
func MTRCertificatesFrom(ptr unsafe.Pointer) MTRCertificates {
	return MTRCertificates{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRCertificatesClass) Alloc() MTRCertificates {
	rv := objc.Send[MTRCertificates](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRCertificatesClass) New() MTRCertificates {
	rv := objc.Send[MTRCertificates](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRCertificates) Init() MTRCertificates {
	rv := objc.Send[MTRCertificates](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRCertificates) Autorelease() MTRCertificates {
	rv := objc.Send[MTRCertificates](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRCertificates creates a new MTRCertificates instance.
func NewMTRCertificates() MTRCertificates {
	return getMTRCertificatesClass().New()
}
