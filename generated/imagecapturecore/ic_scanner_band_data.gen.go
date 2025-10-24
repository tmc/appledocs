// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ICScannerBandData */


/* debug [class_header]: Header for ICScannerBandData */
// The class instance for the [ICScannerBandData] class.
var (
	ICScannerBandDataClass     _ICScannerBandDataClass
	ICScannerBandDataClassOnce sync.Once
)

func getICScannerBandDataClass() _ICScannerBandDataClass {
	ICScannerBandDataClassOnce.Do(func() {
		ICScannerBandDataClass = _ICScannerBandDataClass{objc.GetClass("ICScannerBandData")}
	})
	return ICScannerBandDataClass
}

type _ICScannerBandDataClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ICScannerBandData */
// An interface definition for the [ICScannerBandData] class.
type IICScannerBandData interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ICScannerBandData */
	// properties:
	FullImageWidth() unsafe.Pointer
	SetFullImageWidth(value unsafe.Pointer)
	FullImageHeight() unsafe.Pointer
	SetFullImageHeight(value unsafe.Pointer)
	BitsPerComponent() unsafe.Pointer
	SetBitsPerComponent(value unsafe.Pointer)
	IsBigEndian() unsafe.Pointer
	SetIsBigEndian(value unsafe.Pointer)
	DataNumRows() unsafe.Pointer
	SetDataNumRows(value unsafe.Pointer)
	NumComponents() unsafe.Pointer
	SetNumComponents(value unsafe.Pointer)
	ColorSyncProfilePath() unsafe.Pointer
	SetColorSyncProfilePath(value unsafe.Pointer)
	BytesPerRow() unsafe.Pointer
	SetBytesPerRow(value unsafe.Pointer)
	PixelDataType() unsafe.Pointer
	SetPixelDataType(value unsafe.Pointer)
	DataSize() unsafe.Pointer
	SetDataSize(value unsafe.Pointer)
	BitsPerPixel() unsafe.Pointer
	SetBitsPerPixel(value unsafe.Pointer)
	DataBuffer() unsafe.Pointer
	SetDataBuffer(value unsafe.Pointer)
	DataStartRow() unsafe.Pointer
	SetDataStartRow(value unsafe.Pointer)
	BigEndian() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ICScannerBandData */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ICScannerBandData */
// Alloc allocates a new instance without initialization.
func (ic _ICScannerBandDataClass) Alloc() ICScannerBandData {
	rv := objc.Send[ICScannerBandData](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ICScannerBandDataClass) New() ICScannerBandData {
	rv := objc.Send[ICScannerBandData](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ICScannerBandData) Init() ICScannerBandData {
	rv := objc.Send[ICScannerBandData](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ICScannerBandData) Autorelease() ICScannerBandData {
	rv := objc.Send[ICScannerBandData](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewICScannerBandData creates a new ICScannerBandData instance.
func NewICScannerBandData() ICScannerBandData {
	return getICScannerBandDataClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ICScannerBandData */
// The options for each band of data that the scanner reads.


// The options for each band of data that the scanner reads.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerBandData
type ICScannerBandData struct {
	objectivec.Object
}

// ICScannerBandDataFrom constructs a [ICScannerBandData] from an unsafe.Pointer.
//
// The options for each band of data that the scanner reads.
func ICScannerBandDataFrom(ptr unsafe.Pointer) ICScannerBandData {
	return ICScannerBandData{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ICScannerBandData *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ICScannerBandData */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ICScannerBandData */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ICScannerBandData */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ICScannerBandData */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerbanddata/1507632-fullimagewidth
func (i_ ICScannerBandData) FullImageWidth() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("fullImageWidth"))
	return rv
}/* debug [instance_properties/getter]: fullImageWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerbanddata/1507632-fullimagewidth
func (i_ ICScannerBandData) SetFullImageWidth(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setFullImageWidth:"), value)
}/* debug [instance_properties/setter]: fullImageWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerbanddata/1507667-fullimageheight
func (i_ ICScannerBandData) FullImageHeight() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("fullImageHeight"))
	return rv
}/* debug [instance_properties/getter]: fullImageHeight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerbanddata/1507667-fullimageheight
func (i_ ICScannerBandData) SetFullImageHeight(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setFullImageHeight:"), value)
}/* debug [instance_properties/setter]: fullImageHeight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerbanddata/1507671-bitspercomponent
func (i_ ICScannerBandData) BitsPerComponent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("bitsPerComponent"))
	return rv
}/* debug [instance_properties/getter]: bitsPerComponent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerbanddata/1507671-bitspercomponent
func (i_ ICScannerBandData) SetBitsPerComponent(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBitsPerComponent:"), value)
}/* debug [instance_properties/setter]: bitsPerComponent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerbanddata/1507710-isbigendian
func (i_ ICScannerBandData) IsBigEndian() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("isBigEndian"))
	return rv
}/* debug [instance_properties/getter]: isBigEndian */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerbanddata/1507710-isbigendian
func (i_ ICScannerBandData) SetIsBigEndian(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsBigEndian:"), value)
}/* debug [instance_properties/setter]: isBigEndian */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerbanddata/1507714-datanumrows
func (i_ ICScannerBandData) DataNumRows() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("dataNumRows"))
	return rv
}/* debug [instance_properties/getter]: dataNumRows */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerbanddata/1507714-datanumrows
func (i_ ICScannerBandData) SetDataNumRows(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDataNumRows:"), value)
}/* debug [instance_properties/setter]: dataNumRows */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerbanddata/1507728-numcomponents
func (i_ ICScannerBandData) NumComponents() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("numComponents"))
	return rv
}/* debug [instance_properties/getter]: numComponents */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerbanddata/1507728-numcomponents
func (i_ ICScannerBandData) SetNumComponents(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setNumComponents:"), value)
}/* debug [instance_properties/setter]: numComponents */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerbanddata/1507730-colorsyncprofilepath
func (i_ ICScannerBandData) ColorSyncProfilePath() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("colorSyncProfilePath"))
	return rv
}/* debug [instance_properties/getter]: colorSyncProfilePath */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerbanddata/1507730-colorsyncprofilepath
func (i_ ICScannerBandData) SetColorSyncProfilePath(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setColorSyncProfilePath:"), value)
}/* debug [instance_properties/setter]: colorSyncProfilePath */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerbanddata/1507751-bytesperrow
func (i_ ICScannerBandData) BytesPerRow() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("bytesPerRow"))
	return rv
}/* debug [instance_properties/getter]: bytesPerRow */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerbanddata/1507751-bytesperrow
func (i_ ICScannerBandData) SetBytesPerRow(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBytesPerRow:"), value)
}/* debug [instance_properties/setter]: bytesPerRow */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerbanddata/1507776-pixeldatatype
func (i_ ICScannerBandData) PixelDataType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("pixelDataType"))
	return rv
}/* debug [instance_properties/getter]: pixelDataType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerbanddata/1507776-pixeldatatype
func (i_ ICScannerBandData) SetPixelDataType(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPixelDataType:"), value)
}/* debug [instance_properties/setter]: pixelDataType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerbanddata/1507877-datasize
func (i_ ICScannerBandData) DataSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("dataSize"))
	return rv
}/* debug [instance_properties/getter]: dataSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerbanddata/1507877-datasize
func (i_ ICScannerBandData) SetDataSize(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDataSize:"), value)
}/* debug [instance_properties/setter]: dataSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerbanddata/1507886-bitsperpixel
func (i_ ICScannerBandData) BitsPerPixel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("bitsPerPixel"))
	return rv
}/* debug [instance_properties/getter]: bitsPerPixel */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerbanddata/1507886-bitsperpixel
func (i_ ICScannerBandData) SetBitsPerPixel(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBitsPerPixel:"), value)
}/* debug [instance_properties/setter]: bitsPerPixel */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerbanddata/1507888-databuffer
func (i_ ICScannerBandData) DataBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("dataBuffer"))
	return rv
}/* debug [instance_properties/getter]: dataBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerbanddata/1507888-databuffer
func (i_ ICScannerBandData) SetDataBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDataBuffer:"), value)
}/* debug [instance_properties/setter]: dataBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerbanddata/1508005-datastartrow
func (i_ ICScannerBandData) DataStartRow() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("dataStartRow"))
	return rv
}/* debug [instance_properties/getter]: dataStartRow */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerbanddata/1508005-datastartrow
func (i_ ICScannerBandData) SetDataStartRow(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDataStartRow:"), value)
}/* debug [instance_properties/setter]: dataStartRow */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerBandData/isBigEndian
func (i_ ICScannerBandData) BigEndian() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("bigEndian"))
	return rv
}/* debug [instance_properties/getter]: bigEndian */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ICScannerBandData */



