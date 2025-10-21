// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCertificates/convertMatterCertificate(_:)
func (mc _MTRCertificatesClass) ConvertMatterCertificate(matterCertificate IMTRCertificateTLVBytes) MTRCertificateDERBytes {
	rv := objc.Send[MTRCertificateDERBytes](objc.ID(mc.class), objc.Sel("convertMatterCertificate:"), matterCertificate)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCertificates/convertX509Certificate(_:)
func (mc _MTRCertificatesClass) ConvertX509Certificate(x509Certificate IMTRCertificateDERBytes) MTRCertificateTLVBytes {
	rv := objc.Send[MTRCertificateTLVBytes](objc.ID(mc.class), objc.Sel("convertX509Certificate:"), x509Certificate)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCertificates/createCertificateSigningRequest(_:)
func (mc _MTRCertificatesClass) CreateCertificateSigningRequestError(keypair objectivec.IObject, error_ unsafe.Pointer) MTRCSRDERBytes {
	rv := objc.Send[MTRCSRDERBytes](objc.ID(mc.class), objc.Sel("createCertificateSigningRequest:error:"), keypair, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCertificates/createIntermediateCertificate(_:rootCertificate:intermediatePublicKey:issuerID:fabricID:)
func (mc _MTRCertificatesClass) CreateIntermediateCertificateRootCertificateIntermediatePublicKeyIssuerIDFabricIDError(rootKeypair objectivec.IObject, rootCertificate IMTRCertificateDERBytes, intermediatePublicKey unsafe.Pointer, issuerID foundation.INumber, fabricID foundation.INumber, error_ unsafe.Pointer) MTRCertificateDERBytes {
	rv := objc.Send[MTRCertificateDERBytes](objc.ID(mc.class), objc.Sel("createIntermediateCertificate:rootCertificate:intermediatePublicKey:issuerID:fabricID:error:"), rootKeypair, rootCertificate, intermediatePublicKey, issuerID, fabricID, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCertificates/createIntermediateCertificate(_:rootCertificate:intermediatePublicKey:issuerID:fabricID:validityPeriod:)
func (mc _MTRCertificatesClass) CreateIntermediateCertificateRootCertificateIntermediatePublicKeyIssuerIDFabricIDValidityPeriodError(rootKeypair objectivec.IObject, rootCertificate IMTRCertificateDERBytes, intermediatePublicKey unsafe.Pointer, issuerID foundation.INumber, fabricID foundation.INumber, validityPeriod foundation.IDateInterval, error_ unsafe.Pointer) MTRCertificateDERBytes {
	rv := objc.Send[MTRCertificateDERBytes](objc.ID(mc.class), objc.Sel("createIntermediateCertificate:rootCertificate:intermediatePublicKey:issuerID:fabricID:validityPeriod:error:"), rootKeypair, rootCertificate, intermediatePublicKey, issuerID, fabricID, validityPeriod, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCertificates/createOperationalCertificate(_:signingCertificate:operationalPublicKey:fabricID:nodeID:caseAuthenticatedTags:)
func (mc _MTRCertificatesClass) CreateOperationalCertificateSigningCertificateOperationalPublicKeyFabricIDNodeIDCaseAuthenticatedTagsError(signingKeypair objectivec.IObject, signingCertificate IMTRCertificateDERBytes, operationalPublicKey unsafe.Pointer, fabricID foundation.INumber, nodeID foundation.INumber, caseAuthenticatedTags unsafe.Pointer, error_ unsafe.Pointer) MTRCertificateDERBytes {
	rv := objc.Send[MTRCertificateDERBytes](objc.ID(mc.class), objc.Sel("createOperationalCertificate:signingCertificate:operationalPublicKey:fabricID:nodeID:caseAuthenticatedTags:error:"), signingKeypair, signingCertificate, operationalPublicKey, fabricID, nodeID, caseAuthenticatedTags, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCertificates/createOperationalCertificate(_:signingCertificate:operationalPublicKey:fabricID:nodeID:caseAuthenticatedTags:validityPeriod:)
func (mc _MTRCertificatesClass) CreateOperationalCertificateSigningCertificateOperationalPublicKeyFabricIDNodeIDCaseAuthenticatedTagsValidityPeriodError(signingKeypair objectivec.IObject, signingCertificate IMTRCertificateDERBytes, operationalPublicKey unsafe.Pointer, fabricID foundation.INumber, nodeID foundation.INumber, caseAuthenticatedTags unsafe.Pointer, validityPeriod foundation.IDateInterval, error_ unsafe.Pointer) MTRCertificateDERBytes {
	rv := objc.Send[MTRCertificateDERBytes](objc.ID(mc.class), objc.Sel("createOperationalCertificate:signingCertificate:operationalPublicKey:fabricID:nodeID:caseAuthenticatedTags:validityPeriod:error:"), signingKeypair, signingCertificate, operationalPublicKey, fabricID, nodeID, caseAuthenticatedTags, validityPeriod, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCertificates/createRootCertificate(_:issuerID:fabricID:)
func (mc _MTRCertificatesClass) CreateRootCertificateIssuerIDFabricIDError(keypair objectivec.IObject, issuerID foundation.INumber, fabricID foundation.INumber, error_ unsafe.Pointer) MTRCertificateDERBytes {
	rv := objc.Send[MTRCertificateDERBytes](objc.ID(mc.class), objc.Sel("createRootCertificate:issuerID:fabricID:error:"), keypair, issuerID, fabricID, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCertificates/createRootCertificate(_:issuerID:fabricID:validityPeriod:)
func (mc _MTRCertificatesClass) CreateRootCertificateIssuerIDFabricIDValidityPeriodError(keypair objectivec.IObject, issuerID foundation.INumber, fabricID foundation.INumber, validityPeriod foundation.IDateInterval, error_ unsafe.Pointer) MTRCertificateDERBytes {
	rv := objc.Send[MTRCertificateDERBytes](objc.ID(mc.class), objc.Sel("createRootCertificate:issuerID:fabricID:validityPeriod:error:"), keypair, issuerID, fabricID, validityPeriod, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCertificates/generateCertificateSigningRequest(_:)
func (mc _MTRCertificatesClass) GenerateCertificateSigningRequestError(keypair objectivec.IObject, error_ unsafe.Pointer) foundation.Data {
	rv := objc.Send[foundation.Data](objc.ID(mc.class), objc.Sel("generateCertificateSigningRequest:error:"), keypair, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCertificates/generateIntermediateCertificate(_:rootCertificate:intermediatePublicKey:issuerId:fabricId:)
func (mc _MTRCertificatesClass) GenerateIntermediateCertificateRootCertificateIntermediatePublicKeyIssuerIdFabricIdError(rootKeypair objectivec.IObject, rootCertificate foundation.IData, intermediatePublicKey unsafe.Pointer, issuerId foundation.INumber, fabricId foundation.INumber, error_ unsafe.Pointer) foundation.Data {
	rv := objc.Send[foundation.Data](objc.ID(mc.class), objc.Sel("generateIntermediateCertificate:rootCertificate:intermediatePublicKey:issuerId:fabricId:error:"), rootKeypair, rootCertificate, intermediatePublicKey, issuerId, fabricId, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCertificates/generateOperationalCertificate(_:signingCertificate:operationalPublicKey:fabricId:nodeId:caseAuthenticatedTags:)
func (mc _MTRCertificatesClass) GenerateOperationalCertificateSigningCertificateOperationalPublicKeyFabricIdNodeIdCaseAuthenticatedTagsError(signingKeypair objectivec.IObject, signingCertificate foundation.IData, operationalPublicKey unsafe.Pointer, fabricId foundation.INumber, nodeId foundation.INumber, caseAuthenticatedTags []foundation.INumber, error_ unsafe.Pointer) foundation.Data {
	rv := objc.Send[foundation.Data](objc.ID(mc.class), objc.Sel("generateOperationalCertificate:signingCertificate:operationalPublicKey:fabricId:nodeId:caseAuthenticatedTags:error:"), signingKeypair, signingCertificate, operationalPublicKey, fabricId, nodeId, caseAuthenticatedTags, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCertificates/generateRootCertificate(_:issuerId:fabricId:)
func (mc _MTRCertificatesClass) GenerateRootCertificateIssuerIdFabricIdError(keypair objectivec.IObject, issuerId foundation.INumber, fabricId foundation.INumber, error_ unsafe.Pointer) foundation.Data {
	rv := objc.Send[foundation.Data](objc.ID(mc.class), objc.Sel("generateRootCertificate:issuerId:fabricId:error:"), keypair, issuerId, fabricId, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCertificates/isCertificate(_:equalTo:)
func (mc _MTRCertificatesClass) IsCertificateEqualTo(certificate1 IMTRCertificateDERBytes, certificate2 IMTRCertificateDERBytes) bool {
	rv := objc.Send[bool](objc.ID(mc.class), objc.Sel("isCertificate:equalTo:"), certificate1, certificate2)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCertificates/keypair(_:matchesCertificate:)
func (mc _MTRCertificatesClass) KeypairMatchesCertificate(keypair objectivec.IObject, certificate foundation.IData) bool {
	rv := objc.Send[bool](objc.ID(mc.class), objc.Sel("keypair:matchesCertificate:"), keypair, certificate)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCertificates/publicKey(fromCSR:)
func (mc _MTRCertificatesClass) PublicKeyFromCSRError(csr IMTRCSRDERBytes, error_ unsafe.Pointer) foundation.Data {
	rv := objc.Send[foundation.Data](objc.ID(mc.class), objc.Sel("publicKeyFromCSR:error:"), csr, error_)
	return rv
}



