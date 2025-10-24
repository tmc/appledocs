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
	// properties:
	AdminSubject() objc.IObject /* cross-framework: NSNumber */
	SetAdminSubject(value objc.IObject /* cross-framework: NSNumber */)
	IntermediateCertificate() MTRCertificateDERBytes /* typedef */
	SetIntermediateCertificate(value MTRCertificateDERBytes /* typedef */)
	OperationalCertificate() MTRCertificateDERBytes /* typedef */
	SetOperationalCertificate(value MTRCertificateDERBytes /* typedef */)
	RootCertificate() MTRCertificateDERBytes /* typedef */
	SetRootCertificate(value MTRCertificateDERBytes /* typedef */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCertificateChain/init(operationalCertificate:intermediateCertificate:rootCertificate:adminSubject:)
func NewMTROperationalCertificateChainWithOperationalCertificateIntermediateCertificateRootCertificateAdminSubject(operationalCertificate MTRCertificateDERBytes /* typedef */, intermediateCertificate MTRCertificateDERBytes /* typedef */, rootCertificate MTRCertificateDERBytes /* typedef */, adminSubject objc.IObject /* cross-framework: NSNumber */) MTROperationalCertificateChain {
	instance := getMTROperationalCertificateChainClass().Alloc()
	rv := objc.Send[MTROperationalCertificateChain](instance.ID, objc.Sel("initWithOperationalCertificate:intermediateCertificate:rootCertificate:adminSubject:"), operationalCertificate, intermediateCertificate, rootCertificate, adminSubject)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCertificateChain/adminSubject
func (m_ MTROperationalCertificateChain) AdminSubject() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("adminSubject"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCertificateChain/adminSubject
func (m_ MTROperationalCertificateChain) SetAdminSubject(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAdminSubject:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCertificateChain/intermediateCertificate
func (m_ MTROperationalCertificateChain) IntermediateCertificate() MTRCertificateDERBytes /* typedef */ {
	rv := objc.Send[MTRCertificateDERBytes](m_.ID, objc.Sel("intermediateCertificate"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCertificateChain/intermediateCertificate
func (m_ MTROperationalCertificateChain) SetIntermediateCertificate(value MTRCertificateDERBytes /* typedef */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIntermediateCertificate:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCertificateChain/operationalCertificate
func (m_ MTROperationalCertificateChain) OperationalCertificate() MTRCertificateDERBytes /* typedef */ {
	rv := objc.Send[MTRCertificateDERBytes](m_.ID, objc.Sel("operationalCertificate"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCertificateChain/operationalCertificate
func (m_ MTROperationalCertificateChain) SetOperationalCertificate(value MTRCertificateDERBytes /* typedef */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationalCertificate:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCertificateChain/rootCertificate
func (m_ MTROperationalCertificateChain) RootCertificate() MTRCertificateDERBytes /* typedef */ {
	rv := objc.Send[MTRCertificateDERBytes](m_.ID, objc.Sel("rootCertificate"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCertificateChain/rootCertificate
func (m_ MTROperationalCertificateChain) SetRootCertificate(value MTRCertificateDERBytes /* typedef */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRootCertificate:"), value)
}


