// Code generated from Apple documentation for Accelerate. DO NOT EDIT.

package accelerate

/* debug [enums.gen.go]: Generating 4 enums for Accelerate */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum CBLAS_ORDER (0 cases) */
// CBLAS_ORDER enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/CBLAS_ORDER
type CBLAS_ORDER uint

/* debug [enums.gen.go]: Processing enum vDSP_DCT_Type (3 cases) */
// vDSP_DCT_Type - Constants that indicate the type for a discrete cosine transform.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_DCT_Type
type vDSP_DCT_Type uint

const (
	// vDSP_DCT_II - A constant that specifies a type II discrete cosine transform.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_DCT_Type/II
	vDSP_DCT_II vDSP_DCT_Type = 0
	// vDSP_DCT_III - A constant that specifies a type III discrete cosine transform.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_DCT_Type/III
	vDSP_DCT_III vDSP_DCT_Type = 0
	// vDSP_DCT_IV - A constant that specifies a type IV discrete cosine transform.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_DCT_Type/IV
	vDSP_DCT_IV vDSP_DCT_Type = 0
)

/* debug [enums.gen.go]: Processing enum vDSP_DFT_Direction (2 cases) */
// vDSP_DFT_Direction - An enumeration that specifies whether to perform a forward or inverse DFT.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_DFT_Direction
type vDSP_DFT_Direction uint

const (
	// vDSP_DFT_FORWARD - A constant that specifies a forward transform.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_DFT_Direction/FORWARD
	vDSP_DFT_FORWARD vDSP_DFT_Direction = 0
	// vDSP_DFT_INVERSE - A constant that specifies an inverse transform.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_DFT_Direction/INVERSE
	vDSP_DFT_INVERSE vDSP_DFT_Direction = 0
)

/* debug [enums.gen.go]: Processing enum vDSP_DFT_RealtoComplex (2 cases) */
// vDSP_DFT_RealtoComplex enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_DFT_RealtoComplex
type vDSP_DFT_RealtoComplex uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_DFT_RealtoComplex/interleaved_ComplextoComplex
	vDSP_DFT_Interleaved_ComplextoComplex vDSP_DFT_RealtoComplex = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_DFT_RealtoComplex/interleaved_RealtoComplex
	vDSP_DFT_Interleaved_RealtoComplex vDSP_DFT_RealtoComplex = 0
)


