// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ICScannerFunctionalUnit */


/* debug [class_header]: Header for ICScannerFunctionalUnit */
// The class instance for the [ICScannerFunctionalUnit] class.
var (
	ICScannerFunctionalUnitClass     _ICScannerFunctionalUnitClass
	ICScannerFunctionalUnitClassOnce sync.Once
)

func getICScannerFunctionalUnitClass() _ICScannerFunctionalUnitClass {
	ICScannerFunctionalUnitClassOnce.Do(func() {
		ICScannerFunctionalUnitClass = _ICScannerFunctionalUnitClass{objc.GetClass("ICScannerFunctionalUnit")}
	})
	return ICScannerFunctionalUnitClass
}

type _ICScannerFunctionalUnitClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ICScannerFunctionalUnit */
// An interface definition for the [ICScannerFunctionalUnit] class.
type IICScannerFunctionalUnit interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ICScannerFunctionalUnit */
	// properties:
	ScanInProgress() unsafe.Pointer
	SetScanInProgress(value unsafe.Pointer)
	PreferredResolutions() unsafe.Pointer
	SetPreferredResolutions(value unsafe.Pointer)
	OverviewImage() Image get /* not a class type */
	SetOverviewImage(value Image get /* not a class type */)
	OverviewScanInProgress() unsafe.Pointer
	SetOverviewScanInProgress(value unsafe.Pointer)
	PreferredScaleFactors() unsafe.Pointer
	SetPreferredScaleFactors(value unsafe.Pointer)
	ScanProgressPercentDone() Float get /* not a class type */
	SetScanProgressPercentDone(value Float get /* not a class type */)
	MeasurementUnit() unsafe.Pointer
	SetMeasurementUnit(value unsafe.Pointer)
	SupportedMeasurementUnits() unsafe.Pointer
	SetSupportedMeasurementUnits(value unsafe.Pointer)
	ScanAreaOrientation() unsafe.Pointer
	SetScanAreaOrientation(value unsafe.Pointer)
	State() unsafe.Pointer
	SetState(value unsafe.Pointer)
	AcceptsThresholdForBlackAndWhiteScanning() unsafe.Pointer
	SetAcceptsThresholdForBlackAndWhiteScanning(value unsafe.Pointer)
	DefaultThresholdForBlackAndWhiteScanning() unsafe.Pointer
	SetDefaultThresholdForBlackAndWhiteScanning(value unsafe.Pointer)
	PixelDataType() unsafe.Pointer
	SetPixelDataType(value unsafe.Pointer)
	VendorFeatures() ICScannerFeature
	SetVendorFeatures(value ICScannerFeature)
	SupportedBitDepths() unsafe.Pointer
	SetSupportedBitDepths(value unsafe.Pointer)
	SupportedResolutions() unsafe.Pointer
	SetSupportedResolutions(value unsafe.Pointer)
	SupportedScaleFactors() unsafe.Pointer
	SetSupportedScaleFactors(value unsafe.Pointer)
	ScaleFactor() unsafe.Pointer
	SetScaleFactor(value unsafe.Pointer)
	BitDepth() unsafe.Pointer
	SetBitDepth(value unsafe.Pointer)
	ThresholdForBlackAndWhiteScanning() unsafe.Pointer
	SetThresholdForBlackAndWhiteScanning(value unsafe.Pointer)
	CanPerformOverviewScan() unsafe.Pointer
	SetCanPerformOverviewScan(value unsafe.Pointer)
	Type() unsafe.Pointer
	SetType(value unsafe.Pointer)
	UsesThresholdForBlackAndWhiteScanning() unsafe.Pointer
	SetUsesThresholdForBlackAndWhiteScanning(value unsafe.Pointer)
	PhysicalSize() Size get /* not a class type */
	SetPhysicalSize(value Size get /* not a class type */)
	OverviewResolution() unsafe.Pointer
	SetOverviewResolution(value unsafe.Pointer)
	Resolution() unsafe.Pointer
	SetResolution(value unsafe.Pointer)
	ScanArea() Rect get set /* not a class type */
	SetScanArea(value Rect get set /* not a class type */)
	Templates() ICScannerFeatureTemplate
	SetTemplates(value ICScannerFeatureTemplate)
	NativeYResolution() unsafe.Pointer
	SetNativeYResolution(value unsafe.Pointer)
	NativeXResolution() unsafe.Pointer
	SetNativeXResolution(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ICScannerFunctionalUnit */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ICScannerFunctionalUnit */
// Alloc allocates a new instance without initialization.
func (ic _ICScannerFunctionalUnitClass) Alloc() ICScannerFunctionalUnit {
	rv := objc.Send[ICScannerFunctionalUnit](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ICScannerFunctionalUnitClass) New() ICScannerFunctionalUnit {
	rv := objc.Send[ICScannerFunctionalUnit](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ICScannerFunctionalUnit) Init() ICScannerFunctionalUnit {
	rv := objc.Send[ICScannerFunctionalUnit](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ICScannerFunctionalUnit) Autorelease() ICScannerFunctionalUnit {
	rv := objc.Send[ICScannerFunctionalUnit](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewICScannerFunctionalUnit creates a new ICScannerFunctionalUnit instance.
func NewICScannerFunctionalUnit() ICScannerFunctionalUnit {
	return getICScannerFunctionalUnitClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ICScannerFunctionalUnit */
// An abstract class that represents a scanner functional unit.
//
// The ImageCaptureCore framework defines four concrete subclasses of functional units: creates instances of these subclasses.


// An abstract class that represents a scanner functional unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnit
type ICScannerFunctionalUnit struct {
	objectivec.Object
}

// ICScannerFunctionalUnitFrom constructs a [ICScannerFunctionalUnit] from an unsafe.Pointer.
//
// An abstract class that represents a scanner functional unit.
func ICScannerFunctionalUnitFrom(ptr unsafe.Pointer) ICScannerFunctionalUnit {
	return ICScannerFunctionalUnit{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ICScannerFunctionalUnit *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ICScannerFunctionalUnit */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ICScannerFunctionalUnit */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ICScannerFunctionalUnit */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ICScannerFunctionalUnit */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507566-scaninprogress
func (i_ ICScannerFunctionalUnit) ScanInProgress() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("scanInProgress"))
	return rv
}/* debug [instance_properties/getter]: scanInProgress */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507566-scaninprogress
func (i_ ICScannerFunctionalUnit) SetScanInProgress(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setScanInProgress:"), value)
}/* debug [instance_properties/setter]: scanInProgress */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507635-preferredresolutions
func (i_ ICScannerFunctionalUnit) PreferredResolutions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("preferredResolutions"))
	return rv
}/* debug [instance_properties/getter]: preferredResolutions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507635-preferredresolutions
func (i_ ICScannerFunctionalUnit) SetPreferredResolutions(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPreferredResolutions:"), value)
}/* debug [instance_properties/setter]: preferredResolutions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507637-overviewimage
func (i_ ICScannerFunctionalUnit) OverviewImage() Image get /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("overviewImage"))
	return rv
}/* debug [instance_properties/getter]: overviewImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507637-overviewimage
func (i_ ICScannerFunctionalUnit) SetOverviewImage(value Image get /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setOverviewImage:"), value)
}/* debug [instance_properties/setter]: overviewImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507709-overviewscaninprogress
func (i_ ICScannerFunctionalUnit) OverviewScanInProgress() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("overviewScanInProgress"))
	return rv
}/* debug [instance_properties/getter]: overviewScanInProgress */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507709-overviewscaninprogress
func (i_ ICScannerFunctionalUnit) SetOverviewScanInProgress(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setOverviewScanInProgress:"), value)
}/* debug [instance_properties/setter]: overviewScanInProgress */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507734-preferredscalefactors
func (i_ ICScannerFunctionalUnit) PreferredScaleFactors() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("preferredScaleFactors"))
	return rv
}/* debug [instance_properties/getter]: preferredScaleFactors */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507734-preferredscalefactors
func (i_ ICScannerFunctionalUnit) SetPreferredScaleFactors(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPreferredScaleFactors:"), value)
}/* debug [instance_properties/setter]: preferredScaleFactors */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507754-scanprogresspercentdone
func (i_ ICScannerFunctionalUnit) ScanProgressPercentDone() Float get /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("scanProgressPercentDone"))
	return rv
}/* debug [instance_properties/getter]: scanProgressPercentDone */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507754-scanprogresspercentdone
func (i_ ICScannerFunctionalUnit) SetScanProgressPercentDone(value Float get /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setScanProgressPercentDone:"), value)
}/* debug [instance_properties/setter]: scanProgressPercentDone */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507832-measurementunit
func (i_ ICScannerFunctionalUnit) MeasurementUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("measurementUnit"))
	return rv
}/* debug [instance_properties/getter]: measurementUnit */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507832-measurementunit
func (i_ ICScannerFunctionalUnit) SetMeasurementUnit(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMeasurementUnit:"), value)
}/* debug [instance_properties/setter]: measurementUnit */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507833-supportedmeasurementunits
func (i_ ICScannerFunctionalUnit) SupportedMeasurementUnits() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("supportedMeasurementUnits"))
	return rv
}/* debug [instance_properties/getter]: supportedMeasurementUnits */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507833-supportedmeasurementunits
func (i_ ICScannerFunctionalUnit) SetSupportedMeasurementUnits(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSupportedMeasurementUnits:"), value)
}/* debug [instance_properties/setter]: supportedMeasurementUnits */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507849-scanareaorientation
func (i_ ICScannerFunctionalUnit) ScanAreaOrientation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("scanAreaOrientation"))
	return rv
}/* debug [instance_properties/getter]: scanAreaOrientation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507849-scanareaorientation
func (i_ ICScannerFunctionalUnit) SetScanAreaOrientation(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setScanAreaOrientation:"), value)
}/* debug [instance_properties/setter]: scanAreaOrientation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507853-state
func (i_ ICScannerFunctionalUnit) State() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507853-state
func (i_ ICScannerFunctionalUnit) SetState(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setState:"), value)
}/* debug [instance_properties/setter]: state */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507881-acceptsthresholdforblackandwhite
func (i_ ICScannerFunctionalUnit) AcceptsThresholdForBlackAndWhiteScanning() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("acceptsThresholdForBlackAndWhiteScanning"))
	return rv
}/* debug [instance_properties/getter]: acceptsThresholdForBlackAndWhiteScanning */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507881-acceptsthresholdforblackandwhite
func (i_ ICScannerFunctionalUnit) SetAcceptsThresholdForBlackAndWhiteScanning(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAcceptsThresholdForBlackAndWhiteScanning:"), value)
}/* debug [instance_properties/setter]: acceptsThresholdForBlackAndWhiteScanning */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507900-defaultthresholdforblackandwhite
func (i_ ICScannerFunctionalUnit) DefaultThresholdForBlackAndWhiteScanning() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("defaultThresholdForBlackAndWhiteScanning"))
	return rv
}/* debug [instance_properties/getter]: defaultThresholdForBlackAndWhiteScanning */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507900-defaultthresholdforblackandwhite
func (i_ ICScannerFunctionalUnit) SetDefaultThresholdForBlackAndWhiteScanning(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDefaultThresholdForBlackAndWhiteScanning:"), value)
}/* debug [instance_properties/setter]: defaultThresholdForBlackAndWhiteScanning */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507901-pixeldatatype
func (i_ ICScannerFunctionalUnit) PixelDataType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("pixelDataType"))
	return rv
}/* debug [instance_properties/getter]: pixelDataType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507901-pixeldatatype
func (i_ ICScannerFunctionalUnit) SetPixelDataType(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPixelDataType:"), value)
}/* debug [instance_properties/setter]: pixelDataType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507905-vendorfeatures
func (i_ ICScannerFunctionalUnit) VendorFeatures() ICScannerFeature {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("vendorFeatures"))
	return rv
}/* debug [instance_properties/getter]: vendorFeatures */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507905-vendorfeatures
func (i_ ICScannerFunctionalUnit) SetVendorFeatures(value ICScannerFeature) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setVendorFeatures:"), value)
}/* debug [instance_properties/setter]: vendorFeatures */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507910-supportedbitdepths
func (i_ ICScannerFunctionalUnit) SupportedBitDepths() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("supportedBitDepths"))
	return rv
}/* debug [instance_properties/getter]: supportedBitDepths */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507910-supportedbitdepths
func (i_ ICScannerFunctionalUnit) SetSupportedBitDepths(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSupportedBitDepths:"), value)
}/* debug [instance_properties/setter]: supportedBitDepths */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507913-supportedresolutions
func (i_ ICScannerFunctionalUnit) SupportedResolutions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("supportedResolutions"))
	return rv
}/* debug [instance_properties/getter]: supportedResolutions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507913-supportedresolutions
func (i_ ICScannerFunctionalUnit) SetSupportedResolutions(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSupportedResolutions:"), value)
}/* debug [instance_properties/setter]: supportedResolutions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507919-supportedscalefactors
func (i_ ICScannerFunctionalUnit) SupportedScaleFactors() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("supportedScaleFactors"))
	return rv
}/* debug [instance_properties/getter]: supportedScaleFactors */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507919-supportedscalefactors
func (i_ ICScannerFunctionalUnit) SetSupportedScaleFactors(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSupportedScaleFactors:"), value)
}/* debug [instance_properties/setter]: supportedScaleFactors */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507928-scalefactor
func (i_ ICScannerFunctionalUnit) ScaleFactor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("scaleFactor"))
	return rv
}/* debug [instance_properties/getter]: scaleFactor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507928-scalefactor
func (i_ ICScannerFunctionalUnit) SetScaleFactor(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setScaleFactor:"), value)
}/* debug [instance_properties/setter]: scaleFactor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507954-bitdepth
func (i_ ICScannerFunctionalUnit) BitDepth() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("bitDepth"))
	return rv
}/* debug [instance_properties/getter]: bitDepth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507954-bitdepth
func (i_ ICScannerFunctionalUnit) SetBitDepth(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBitDepth:"), value)
}/* debug [instance_properties/setter]: bitDepth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507975-thresholdforblackandwhitescannin
func (i_ ICScannerFunctionalUnit) ThresholdForBlackAndWhiteScanning() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("thresholdForBlackAndWhiteScanning"))
	return rv
}/* debug [instance_properties/getter]: thresholdForBlackAndWhiteScanning */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507975-thresholdforblackandwhitescannin
func (i_ ICScannerFunctionalUnit) SetThresholdForBlackAndWhiteScanning(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setThresholdForBlackAndWhiteScanning:"), value)
}/* debug [instance_properties/setter]: thresholdForBlackAndWhiteScanning */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507998-canperformoverviewscan
func (i_ ICScannerFunctionalUnit) CanPerformOverviewScan() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("canPerformOverviewScan"))
	return rv
}/* debug [instance_properties/getter]: canPerformOverviewScan */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1507998-canperformoverviewscan
func (i_ ICScannerFunctionalUnit) SetCanPerformOverviewScan(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCanPerformOverviewScan:"), value)
}/* debug [instance_properties/setter]: canPerformOverviewScan */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1508014-type
func (i_ ICScannerFunctionalUnit) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1508014-type
func (i_ ICScannerFunctionalUnit) SetType(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setType:"), value)
}/* debug [instance_properties/setter]: type */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1508028-usesthresholdforblackandwhitesca
func (i_ ICScannerFunctionalUnit) UsesThresholdForBlackAndWhiteScanning() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("usesThresholdForBlackAndWhiteScanning"))
	return rv
}/* debug [instance_properties/getter]: usesThresholdForBlackAndWhiteScanning */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1508028-usesthresholdforblackandwhitesca
func (i_ ICScannerFunctionalUnit) SetUsesThresholdForBlackAndWhiteScanning(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setUsesThresholdForBlackAndWhiteScanning:"), value)
}/* debug [instance_properties/setter]: usesThresholdForBlackAndWhiteScanning */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1508033-physicalsize
func (i_ ICScannerFunctionalUnit) PhysicalSize() Size get /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("physicalSize"))
	return rv
}/* debug [instance_properties/getter]: physicalSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1508033-physicalsize
func (i_ ICScannerFunctionalUnit) SetPhysicalSize(value Size get /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPhysicalSize:"), value)
}/* debug [instance_properties/setter]: physicalSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1508054-overviewresolution
func (i_ ICScannerFunctionalUnit) OverviewResolution() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("overviewResolution"))
	return rv
}/* debug [instance_properties/getter]: overviewResolution */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1508054-overviewresolution
func (i_ ICScannerFunctionalUnit) SetOverviewResolution(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setOverviewResolution:"), value)
}/* debug [instance_properties/setter]: overviewResolution */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1508082-resolution
func (i_ ICScannerFunctionalUnit) Resolution() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("resolution"))
	return rv
}/* debug [instance_properties/getter]: resolution */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1508082-resolution
func (i_ ICScannerFunctionalUnit) SetResolution(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setResolution:"), value)
}/* debug [instance_properties/setter]: resolution */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1508107-scanarea
func (i_ ICScannerFunctionalUnit) ScanArea() Rect get set /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("scanArea"))
	return rv
}/* debug [instance_properties/getter]: scanArea */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1508107-scanarea
func (i_ ICScannerFunctionalUnit) SetScanArea(value Rect get set /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setScanArea:"), value)
}/* debug [instance_properties/setter]: scanArea */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1508141-templates
func (i_ ICScannerFunctionalUnit) Templates() ICScannerFeatureTemplate {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("templates"))
	return rv
}/* debug [instance_properties/getter]: templates */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1508141-templates
func (i_ ICScannerFunctionalUnit) SetTemplates(value ICScannerFeatureTemplate) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTemplates:"), value)
}/* debug [instance_properties/setter]: templates */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1508168-nativeyresolution
func (i_ ICScannerFunctionalUnit) NativeYResolution() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("nativeYResolution"))
	return rv
}/* debug [instance_properties/getter]: nativeYResolution */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1508168-nativeyresolution
func (i_ ICScannerFunctionalUnit) SetNativeYResolution(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setNativeYResolution:"), value)
}/* debug [instance_properties/setter]: nativeYResolution */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1508170-nativexresolution
func (i_ ICScannerFunctionalUnit) NativeXResolution() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("nativeXResolution"))
	return rv
}/* debug [instance_properties/getter]: nativeXResolution */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunit/1508170-nativexresolution
func (i_ ICScannerFunctionalUnit) SetNativeXResolution(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setNativeXResolution:"), value)
}/* debug [instance_properties/setter]: nativeXResolution */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ICScannerFunctionalUnit */



