// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROperationalCertificateChain] class.
var (
	MTROperationalCertificateChainClass     _MTROperationalCertificateChainClass
	MTROperationalCertificateChainClassOnce sync.Once
)

func getMTROperationalCertificateChainClass() _MTROperationalCertificateChainClass {
	MTROperationalCertificateChainClassOnce.Do(func() {
		MTROperationalCertificateChainClass = _MTROperationalCertificateChainClass{objc.GetClass("MTROperationalCertificateChain")}
	})
	return MTROperationalCertificateChainClass
}

type _MTROperationalCertificateChainClass struct {
	class objc.Class
}

// An interface definition for the [MTROperationalCertificateChain] class.
type IMTROperationalCertificateChain interface {
	objectivec.IObject
	AdminSubject() foundation.Number
	SetAdminSubject(value foundation.INumber)
	IntermediateCertificate() foundation.Data
	SetIntermediateCertificate(value foundation.IData)
	OperationalCertificate() foundation.Data
	SetOperationalCertificate(value foundation.IData)
	RootCertificate() foundation.Data
	SetRootCertificate(value foundation.IData)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCertificateChain
type MTROperationalCertificateChain struct {
	objectivec.Object
}

// MTROperationalCertificateChainFrom constructs a [MTROperationalCertificateChain] from an unsafe.Pointer.
func MTROperationalCertificateChainFrom(ptr unsafe.Pointer) MTROperationalCertificateChain {
	return MTROperationalCertificateChain{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCertificateChainClass) Alloc() MTROperationalCertificateChain {
	rv := objc.Send[MTROperationalCertificateChain](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROperationalCertificateChainClass) New() MTROperationalCertificateChain {
	rv := objc.Send[MTROperationalCertificateChain](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCertificateChain) Init() MTROperationalCertificateChain {
	rv := objc.Send[MTROperationalCertificateChain](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCertificateChain) Autorelease() MTROperationalCertificateChain {
	rv := objc.Send[MTROperationalCertificateChain](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCertificateChain creates a new MTROperationalCertificateChain instance.
func NewMTROperationalCertificateChain() MTROperationalCertificateChain {
	return getMTROperationalCertificateChainClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcertificatechain/adminsubject
func (m_ MTROperationalCertificateChain) AdminSubject() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("adminSubject"))
	return rv
}


// SetAdminSubject sets the value of the adminSubject property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcertificatechain/adminsubject
func (m_ MTROperationalCertificateChain) SetAdminSubject(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAdminSubject:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcertificatechain/intermediatecertificate
func (m_ MTROperationalCertificateChain) IntermediateCertificate() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("intermediateCertificate"))
	return rv
}


// SetIntermediateCertificate sets the value of the intermediateCertificate property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcertificatechain/intermediatecertificate
func (m_ MTROperationalCertificateChain) SetIntermediateCertificate(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIntermediateCertificate:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcertificatechain/operationalcertificate
func (m_ MTROperationalCertificateChain) OperationalCertificate() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("operationalCertificate"))
	return rv
}


// SetOperationalCertificate sets the value of the operationalCertificate property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcertificatechain/operationalcertificate
func (m_ MTROperationalCertificateChain) SetOperationalCertificate(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationalCertificate:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcertificatechain/rootcertificate
func (m_ MTROperationalCertificateChain) RootCertificate() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("rootCertificate"))
	return rv
}


// SetRootCertificate sets the value of the rootCertificate property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcertificatechain/rootcertificate
func (m_ MTROperationalCertificateChain) SetRootCertificate(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRootCertificate:"), value)
}



