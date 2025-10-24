// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion


// C struct types
// CMAcceleration - The type of a structure containing 3-axis acceleration values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAcceleration
type CMAcceleration struct {
	X float64 // X-axis acceleration in G’s (gravitational force).
	Y float64 // Y-axis acceleration in G’s (gravitational force).
	Z float64 // Z-axis acceleration in G’s (gravitational force).
}// CMCalibratedMagneticField - Calibrated magnetic field data and an estimate of the accuracy of the calibration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMCalibratedMagneticField
type CMCalibratedMagneticField struct {
}// CMMagneticField - A structure containing 3-axis magnetometer data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMagneticField
type CMMagneticField struct {
	X float64 // X-axis magnetic field in microteslas.
	Y float64 // Y-axis magnetic field in microteslas.
	Z float64 // Z-axis magnetic field in microteslas.
}// CMQuaternion - The type for a quaternion representing a measurement of attitude.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMQuaternion
type CMQuaternion struct {
}// CMRotationMatrix - The type of a structure representing a rotation matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMRotationMatrix
type CMRotationMatrix struct {
}// CMRotationRate - The type of structures representing a measurement of rotation rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMRotationRate
type CMRotationRate struct {
	X float64 // The value for the X-axis.
	Y float64 // The value for the Y-axis.
	Z float64 // The value for the Z-axis.
}



