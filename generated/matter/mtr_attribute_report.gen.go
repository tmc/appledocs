// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRAttributeReport] class.
var (
	MTRAttributeReportClass     _MTRAttributeReportClass
	MTRAttributeReportClassOnce sync.Once
)

func getMTRAttributeReportClass() _MTRAttributeReportClass {
	MTRAttributeReportClassOnce.Do(func() {
		MTRAttributeReportClass = _MTRAttributeReportClass{objc.GetClass("MTRAttributeReport")}
	})
	return MTRAttributeReportClass
}

type _MTRAttributeReportClass struct {
	class objc.Class
}

// An interface definition for the [MTRAttributeReport] class.
type IMTRAttributeReport interface {
	objectivec.IObject
	// properties:
	Error() objc.IObject /* cross-framework: Error */
	SetError(value objc.IObject /* cross-framework: Error */)
	Path() IMTRAttributePath
	SetPath(value IMTRAttributePath)
	Value() unsafe.Pointer
	SetValue(value unsafe.Pointer)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAttributeReport
type MTRAttributeReport struct {
	objectivec.Object
}

// MTRAttributeReportFrom constructs a [MTRAttributeReport] from an unsafe.Pointer.
func MTRAttributeReportFrom(ptr unsafe.Pointer) MTRAttributeReport {
	return MTRAttributeReport{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAttributeReportClass) Alloc() MTRAttributeReport {
	rv := objc.Send[MTRAttributeReport](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAttributeReportClass) New() MTRAttributeReport {
	rv := objc.Send[MTRAttributeReport](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAttributeReport) Init() MTRAttributeReport {
	rv := objc.Send[MTRAttributeReport](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAttributeReport) Autorelease() MTRAttributeReport {
	rv := objc.Send[MTRAttributeReport](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAttributeReport creates a new MTRAttributeReport instance.
func NewMTRAttributeReport() MTRAttributeReport {
	return getMTRAttributeReportClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrattributereport/error
func (m_ MTRAttributeReport) Error() objc.IObject /* cross-framework: Error */ {
	rv := objc.Send[coretelephony.Error](m_.ID, objc.Sel("error"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrattributereport/error
func (m_ MTRAttributeReport) SetError(value objc.IObject /* cross-framework: Error */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setError:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrattributereport/path
func (m_ MTRAttributeReport) Path() IMTRAttributePath {
	rv := objc.Send[MTRAttributePath](m_.ID, objc.Sel("path"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrattributereport/path
func (m_ MTRAttributeReport) SetPath(value IMTRAttributePath) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPath:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrattributereport/value
func (m_ MTRAttributeReport) Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("value"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrattributereport/value
func (m_ MTRAttributeReport) SetValue(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}



