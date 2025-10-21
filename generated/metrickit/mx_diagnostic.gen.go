// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MXDiagnostic] class.
var (
	MXDiagnosticClass     _MXDiagnosticClass
	MXDiagnosticClassOnce sync.Once
)

func getMXDiagnosticClass() _MXDiagnosticClass {
	MXDiagnosticClassOnce.Do(func() {
		MXDiagnosticClass = _MXDiagnosticClass{objc.GetClass("MXDiagnostic")}
	})
	return MXDiagnosticClass
}

type _MXDiagnosticClass struct {
	class objc.Class
}

// An interface definition for the [MXDiagnostic] class.
type IMXDiagnostic interface {
	objectivec.IObject
	DictionaryRepresentation() unsafe.Pointer
	JSONRepresentation() unsafe.Pointer
}

// An abstract data class for a diagnostic.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnostic
type MXDiagnostic struct {
	objectivec.Object
}

// MXDiagnosticFrom constructs a [MXDiagnostic] from an unsafe.Pointer.
//
// An abstract data class for a diagnostic.
func MXDiagnosticFrom(ptr unsafe.Pointer) MXDiagnostic {
	return MXDiagnostic{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MXDiagnosticClass) Alloc() MXDiagnostic {
	rv := objc.Send[MXDiagnostic](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MXDiagnosticClass) New() MXDiagnostic {
	rv := objc.Send[MXDiagnostic](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXDiagnostic) Init() MXDiagnostic {
	rv := objc.Send[MXDiagnostic](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXDiagnostic) Autorelease() MXDiagnostic {
	rv := objc.Send[MXDiagnostic](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXDiagnostic creates a new MXDiagnostic instance.
func NewMXDiagnostic() MXDiagnostic {
	return getMXDiagnosticClass().New()
}


// Returns the contents of a diagnostic as a dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnostic/dictionaryRepresentation()
func (m_ MXDiagnostic) DictionaryRepresentation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("dictionaryRepresentation"))
	return rv
}

// Returns the contents of the diagnostic in JSON format.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnostic/jsonRepresentation()
func (m_ MXDiagnostic) JSONRepresentation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("JSONRepresentation"))
	return rv
}

// The value of the bundle version key, short form, in the app’s property list.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnostic/applicationVersion
func (m_ MXDiagnostic) ApplicationVersion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("applicationVersion"))
	return rv
}

// A set of system-level information for the device.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnostic/metaData
func (m_ MXDiagnostic) MetaData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("metaData"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnostic/signpostData
func (m_ MXDiagnostic) SignpostData() []MXSignpostRecord {
	rv := objc.Send[[]MXSignpostRecord](m_.ID, objc.Sel("signpostData"))
	return rv
}



