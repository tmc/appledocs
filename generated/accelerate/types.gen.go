// Code generated from Apple documentation for Accelerate. DO NOT EDIT.

package accelerate
import (
	"unsafe"
)


// C struct types
// SparseOpaquePreconditioner_Complex_Double - Represents a preconditioner for matrices of complex double values .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseOpaquePreconditioner_Complex_Double
type SparseOpaquePreconditioner_Complex_Double struct {
	Apply unsafe.Pointer
	Mem unsafe.Pointer
	Type unsafe.Pointer // Types of preconditioner.
}/* debug [types.gen.go/struct]: SparseOpaquePreconditioner_Complex_Double */

// SparseOpaquePreconditioner_Complex_Float - Represents a preconditioner for matrices of complex float values .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseOpaquePreconditioner_Complex_Float
type SparseOpaquePreconditioner_Complex_Float struct {
	Apply unsafe.Pointer
	Mem unsafe.Pointer
	Type unsafe.Pointer // Types of preconditioner.
}/* debug [types.gen.go/struct]: SparseOpaquePreconditioner_Complex_Float */

// SparseOpaqueSubfactor_Complex_Double - Represents a sub-factor of the factorization (for example,  
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseOpaqueSubfactor_Complex_Double
type SparseOpaqueSubfactor_Complex_Double struct {
	Attributes unsafe.Pointer // A type representing the attributes of a matrix.
	Contents unsafe.Pointer // Types of sub-factor object.
	Factor unsafe.Pointer // A semi-opaque type representing a matrix factorization in complex double.
	WorkspaceRequiredPerRHS uintptr
	WorkspaceRequiredStatic uintptr
}/* debug [types.gen.go/struct]: SparseOpaqueSubfactor_Complex_Double */

// SparseOpaqueSubfactor_Complex_Float - Represents a sub-factor of the factorization (for example,  
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseOpaqueSubfactor_Complex_Float
type SparseOpaqueSubfactor_Complex_Float struct {
	Attributes unsafe.Pointer // A type representing the attributes of a matrix.
	Contents unsafe.Pointer // Types of sub-factor object.
	Factor unsafe.Pointer // A semi-opaque type representing a matrix factorization in complex float.
	WorkspaceRequiredPerRHS uintptr
	WorkspaceRequiredStatic uintptr
}/* debug [types.gen.go/struct]: SparseOpaqueSubfactor_Complex_Float */

// bnns_graph_argument_t - Describes data associated with an input or output argument
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/bnns_graph_argument_t
type bnns_graph_argument_t struct {
	Data_ptr unsafe.Pointer // Direct pointer to numerical data
	Data_ptr_size uintptr // size in bytes of  , if set
	Descriptor BNNSNDArrayDescriptor // Pointer to BNNSNDArrayDescriptor (deprecated, use BNNSTensor instead)
	Tensor BNNSTensor // Pointer to BNNSTensor
}/* debug [types.gen.go/struct]: bnns_graph_argument_t */

// data_ptr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/bnns_graph_argument_t/data_ptr-89cqn
type data_ptr struct {
}/* debug [types.gen.go/struct]: data_ptr */

// descriptor
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/bnns_graph_argument_t/descriptor-8d2bd
type descriptor struct {
}/* debug [types.gen.go/struct]: descriptor */

// tensor
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/bnns_graph_argument_t/tensor-6l2lt
type tensor struct {
}/* debug [types.gen.go/struct]: tensor */

// bnns_graph_compile_options_t - The compilation options that BNNS uses when compiling a source mlmodelc file to a graph object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/bnns_graph_compile_options_t
type bnns_graph_compile_options_t struct {
	Data unsafe.Pointer // A pointer to the opaque compilation options object.
	Size uintptr // The size, in bytes, of the opaque compilation options object.
}/* debug [types.gen.go/struct]: bnns_graph_compile_options_t */

// bnns_graph_context_t - An object that wraps a compiled graph object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/bnns_graph_context_t
type bnns_graph_context_t struct {
	Data unsafe.Pointer // A pointer to the opaque graph context object.
	Size uintptr // The size, in bytes, of the opaque graph context object.
}/* debug [types.gen.go/struct]: bnns_graph_context_t */

// bnns_graph_shape_t - The specification of the shape of an argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/bnns_graph_shape_t
type bnns_graph_shape_t struct {
	Rank uintptr // The rank of the shape.
	Shape []uint64 // An array of unsigned-integer elements that specify the size of the shape.
}/* debug [types.gen.go/struct]: bnns_graph_shape_t */

// bnns_graph_t - The compiled graph object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/bnns_graph_t
type bnns_graph_t struct {
	Data unsafe.Pointer // A pointer to opaque graph object.
	Size uintptr // The size, in bytes, of the opaque graph object.
}/* debug [types.gen.go/struct]: bnns_graph_t */

// bnns_user_message_data_t - Additional user-defined logging argument for message-logging callbacks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/bnns_user_message_data_t
type bnns_user_message_data_t struct {
	Data unsafe.Pointer // A pointer to the additional logging data.
	Size uintptr // The size of the additional logging data.
}/* debug [types.gen.go/struct]: bnns_user_message_data_t */

// BNNSActivation - A set of parameters that describe common activation functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSActivation
type BNNSActivation struct {
	Alpha float32 // The parameter for the alpha of the activation function.
	Beta float32 // The parameter for the beta of the activation function.
	Function unsafe.Pointer // The activation function that the layer applies to its output.
	Ioffset int32 // Offset for integer functions.
	Ioffset_per_channel unsafe.Pointer // Offset per channel for integer functions.
	Iscale int32 // Scale for integer functions.
	Iscale_per_channel unsafe.Pointer // Scale per channel for integer functions.
	Ishift int32 // Shift for integer functions.
	Ishift_per_channel unsafe.Pointer // Shift per channel for integer functions.
}/* debug [types.gen.go/struct]: BNNSActivation */

// BNNSArithmeticBinary - A structure that contains the inputs and output of an arithmetic operation with two inputs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSArithmeticBinary
type BNNSArithmeticBinary struct {
	In1 BNNSNDArrayDescriptor // The descriptor of the first input.
	In1_type unsafe.Pointer // The descriptor type of the first input.
	In2 BNNSNDArrayDescriptor // The descriptor of the second input.
	In2_type unsafe.Pointer // The descriptor type of the second input.
	Out BNNSNDArrayDescriptor // The descriptor of the output.
	Out_type unsafe.Pointer // The descriptor type of the output.
}/* debug [types.gen.go/struct]: BNNSArithmeticBinary */

// BNNSArithmeticTernary - A structure that contains the inputs and output of an arithmetic operation with three inputs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSArithmeticTernary
type BNNSArithmeticTernary struct {
	In1 BNNSNDArrayDescriptor // The descriptor of the first input.
	In1_type unsafe.Pointer // The descriptor type of the first input.
	In2 BNNSNDArrayDescriptor // The descriptor of the second input.
	In2_type unsafe.Pointer // The descriptor type of the second input.
	In3 BNNSNDArrayDescriptor // The descriptor of the third input.
	In3_type unsafe.Pointer // The descriptor type of the third input.
	Out BNNSNDArrayDescriptor // The descriptor of the output.
	Out_type unsafe.Pointer // The descriptor type of the output.
}/* debug [types.gen.go/struct]: BNNSArithmeticTernary */

// BNNSArithmeticUnary - A structure that contains the input and output of an arithmetic operation with a single input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSArithmeticUnary
type BNNSArithmeticUnary struct {
	In BNNSNDArrayDescriptor // The descriptor of the input.
	In_type unsafe.Pointer // The descriptor type of the input.
	Out BNNSNDArrayDescriptor // The descriptor of the output.
	Out_type unsafe.Pointer // The descriptor type of the output.
}/* debug [types.gen.go/struct]: BNNSArithmeticUnary */

// BNNSConvolutionLayerParameters - A structure containing convolution parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSConvolutionLayerParameters
type BNNSConvolutionLayerParameters struct {
	Activation BNNSActivation // The layer activation function.
	Bias BNNSLayerData // Layer bias, one for each output channel.
	In_channels uintptr // The number of input channels.
	K_height uintptr // The height of the convolution kernel.
	K_width uintptr // The width of the convolution kernel.
	Out_channels uintptr // The number of output channels.
	Weights BNNSLayerData // Convolution weights.
	X_padding uintptr // The X padding.
	X_stride uintptr // The X increment in the input image.
	Y_padding uintptr // The Y padding.
	Y_stride uintptr // The Y increment in the input image.
}/* debug [types.gen.go/struct]: BNNSConvolutionLayerParameters */

// BNNSFilterParameters - A structure that contains common filter parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSFilterParameters
type BNNSFilterParameters struct {
	Alloc_memory BNNSAlloc // The function called to allocate memory.
	Flags uint32 // A logical OR of zero or more values from BNNS flags.
	Free_memory BNNSFree // The function called to deallocate memory.
	N_threads uintptr // The number of worker threads to execute.
}/* debug [types.gen.go/struct]: BNNSFilterParameters */

// BNNSFullyConnectedLayerParameters - A structure containing fully connected layer parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSFullyConnectedLayerParameters
type BNNSFullyConnectedLayerParameters struct {
	Activation BNNSActivation // The layer activation function.
	Bias BNNSLayerData // Layer bias, one for each output component.
	In_size uintptr // The size of the input vector.
	Out_size uintptr // The size of the output vector.
	Weights BNNSLayerData // Matrix coefficients.
}/* debug [types.gen.go/struct]: BNNSFullyConnectedLayerParameters */

// BNNSImageStackDescriptor
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSImageStackDescriptor
type BNNSImageStackDescriptor struct {
	Channels uintptr
	Data_bias float32
	Data_scale float32
	Data_type unsafe.Pointer
	Height uintptr
	Image_stride uintptr
	Row_stride uintptr
	Width uintptr
}/* debug [types.gen.go/struct]: BNNSImageStackDescriptor */

// BNNSLayerData - A structure containing common layer parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerData
type BNNSLayerData struct {
	Data unsafe.Pointer // Pointer to layer values (weights, bias), layout and size are specific to each layer.
	Data_bias float32 // Conversion bias for values, used for integer data types only, ignored for indexed and float data types.
	Data_scale float32 // Conversion scale for values, used for integer data types only, ignored for indexed and float data types.
	Data_table []float32 // Conversion table (256 values) for indexed floating point data, used for indexed data types only.
	Data_type unsafe.Pointer // Storage data type for the values stored in data.
}/* debug [types.gen.go/struct]: BNNSLayerData */

// BNNSLayerParametersActivation - A set of parameters that define an activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersActivation
type BNNSLayerParametersActivation struct {
	Activation BNNSActivation // The activation function that the layer applies to the output.
	Axis_flags uint32 // Flags that indicate axes on which to apply certain activation functions.
	I_desc BNNSNDArrayDescriptor // The descriptor of the input.
	O_desc BNNSNDArrayDescriptor // The descriptor of the output.
}/* debug [types.gen.go/struct]: BNNSLayerParametersActivation */

// BNNSLayerParametersArithmetic - A structure that contains the parameters of an arithmetic layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersArithmetic
type BNNSLayerParametersArithmetic struct {
	Activation BNNSActivation // The activation function that the layer applies to the output.
	Arithmetic_function unsafe.Pointer // The arithmetic operation of the layer.
	Arithmetic_function_fields unsafe.Pointer // A pointer to an arithmetic function field structure.
}/* debug [types.gen.go/struct]: BNNSLayerParametersArithmetic */

// BNNSLayerParametersBroadcastMatMul - A set of parameters that define a broadcast matrix multiply layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersBroadcastMatMul
type BNNSLayerParametersBroadcastMatMul struct {
	A_is_weights bool // A Boolean value that determines whether to treat matrix   as weights.
	Alpha float32 // A value to scale the result.
	B_is_weights bool // A Boolean value that determines whether to treat matrix   as weights.
	Beta float32 // A value, that must be either 0.0 or 1.0, you use to scale the existing output before the operation adds it to the result.
	IA_desc BNNSNDArrayDescriptor // The descriptor of matrix  .
	IB_desc BNNSNDArrayDescriptor // The descriptor of matrix  .
	O_desc BNNSNDArrayDescriptor // The descriptor of the output.
	Quadratic bool // A Boolean value that determines whether the operation multiplies matrix   by itself.
	TransA bool // A Boolean value that transposes the last two dimensions of matrix  .
	TransB bool // A Boolean value that transposes the last two dimensions of matrix  .
}/* debug [types.gen.go/struct]: BNNSLayerParametersBroadcastMatMul */

// BNNSLayerParametersConvolution - A structure that contains the parameters of a convolution layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersConvolution
type BNNSLayerParametersConvolution struct {
	Activation BNNSActivation // The activation function that the layer applies to the output.
	Bias BNNSNDArrayDescriptor // The bias descriptor.
	Groups uintptr // Convolution group size.
	I_desc BNNSNDArrayDescriptor // The descriptor of the input.
	O_desc BNNSNDArrayDescriptor // The descriptor of the output.
	Pad uintptr // Padding which is asymmetric and ignored if the width or height padding values are greater than zero.
	W_desc BNNSNDArrayDescriptor // The descriptor of the weights.
	X_dilation_stride uintptr // The width increment between elements in the input image during convolution.
	X_padding uintptr // The width padding, which is the number of virtual zeros added to the left and right of each channel.
	X_stride uintptr // The width increment of the input image.
	Y_dilation_stride uintptr // The height increment between elements in the input image during convolution.
	Y_padding uintptr // The height padding, which is the number of virtual zeros added to the top and bottom of each channel.
	Y_stride uintptr // The height increment of the input image.
}/* debug [types.gen.go/struct]: BNNSLayerParametersConvolution */

// BNNSLayerParametersCropResize - A set of parameters that describe a crop-resize operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersCropResize
type BNNSLayerParametersCropResize struct {
	Box_coordinate_mode unsafe.Pointer // A constant that defines the convention that the operation uses to specify the four bounding box coordinates.
	Extrapolation_value float32 // A value that the operation uses for extrapolation. Default value is  .
	Method unsafe.Pointer // The interpolation method.
	Normalized_coordinates bool // A Boolean value that specifies whether the operation treats the coordinates as normalized to  .
	Sampling_mode unsafe.Pointer // The sampling mode that the operation uses to select sample points.
	Spatial_scale float32 // An additional spatial scale that mutliplies the bounding box coordinates.
}/* debug [types.gen.go/struct]: BNNSLayerParametersCropResize */

// BNNSLayerParametersDropout - A structure that contains the parameters of a dropout layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersDropout
type BNNSLayerParametersDropout struct {
	Control uint8 // An 8-bit bit mask that indicates the dimension of the grouping of the dropout decision.
	I_desc BNNSNDArrayDescriptor // The descriptor of the input.
	O_desc BNNSNDArrayDescriptor // The descriptor of the output.
	Rate float32 // The probability that the layer drops out an element or a group of elements.
	Seed uint32 // The seed for the random number generator which is ignored if zero.
}/* debug [types.gen.go/struct]: BNNSLayerParametersDropout */

// BNNSLayerParametersEmbedding - A structure that contains the parameters of an embedding layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersEmbedding
type BNNSLayerParametersEmbedding struct {
	Dictionary BNNSNDArrayDescriptor // The descriptor of the dictionary.
	Flags unsafe.Pointer // A bit field for flags that specify additional behavior, such as scaling gradient by frequency.
	I_desc BNNSNDArrayDescriptor // The signed or unsigned integer descriptor of the input.
	Max_norm float32 // The maximum norm.
	Norm_type float32 // The norm type.
	O_desc BNNSNDArrayDescriptor // The descriptor of the output.
	Padding_idx uintptr // The padding index.
}/* debug [types.gen.go/struct]: BNNSLayerParametersEmbedding */

// BNNSLayerParametersFullyConnected - A structure that contains the parameters of a fully connected layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersFullyConnected
type BNNSLayerParametersFullyConnected struct {
	Activation BNNSActivation // The activation function that the layer applies to the output.
	Bias BNNSNDArrayDescriptor // The descriptor of the bias.
	I_desc BNNSNDArrayDescriptor // The descriptor of the input.
	O_desc BNNSNDArrayDescriptor // The descriptor of the output.
	W_desc BNNSNDArrayDescriptor // The descriptor of the weights.
}/* debug [types.gen.go/struct]: BNNSLayerParametersFullyConnected */

// BNNSLayerParametersGram - A set of parameters that define a Gram matrix layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersGram
type BNNSLayerParametersGram struct {
	Alpha float32 // A value to scale the result.
	I_desc BNNSNDArrayDescriptor // The descriptor of the input.
	O_desc BNNSNDArrayDescriptor // The descriptor of the output.
}/* debug [types.gen.go/struct]: BNNSLayerParametersGram */

// BNNSLayerParametersLossBase - A structure that contains the parameters of a loss layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersLossBase
type BNNSLayerParametersLossBase struct {
	Function unsafe.Pointer // The function that’s used to compute loss.
	I_desc BNNSNDArrayDescriptor // The descriptor of the input.
	O_desc BNNSNDArrayDescriptor // The descriptor of the output.
	Reduction unsafe.Pointer // The function that’s used to reduce the computed loss.
}/* debug [types.gen.go/struct]: BNNSLayerParametersLossBase */

// BNNSLayerParametersLossHuber - A structure that contains the parameters of a Huber loss layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersLossHuber
type BNNSLayerParametersLossHuber struct {
	Function unsafe.Pointer // The function that’s used to compute loss.
	Huber_delta float32 // The boundary value that defines where Huber loss returns mean absolute error or mean square error.
	I_desc BNNSNDArrayDescriptor // The descriptor of the input.
	O_desc BNNSNDArrayDescriptor // The descriptor of the output.
	Reduction unsafe.Pointer // The function that’s used to reduce the computed loss.
}/* debug [types.gen.go/struct]: BNNSLayerParametersLossHuber */

// BNNSLayerParametersLossSigmoidCrossEntropy - A structure that contains the parameters of a sigmoid cross entropy loss layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersLossSigmoidCrossEntropy
type BNNSLayerParametersLossSigmoidCrossEntropy struct {
	Function unsafe.Pointer // The function that’s used to compute loss.
	I_desc BNNSNDArrayDescriptor // The descriptor of the input.
	Label_smooth float32 // A value that defines the smoothing that the loss function applies to the labels.
	O_desc BNNSNDArrayDescriptor // The descriptor of the output.
	Reduction unsafe.Pointer // The function that’s used to reduce the computed loss.
}/* debug [types.gen.go/struct]: BNNSLayerParametersLossSigmoidCrossEntropy */

// BNNSLayerParametersLossSoftmaxCrossEntropy - A structure that contains the parameters of a softmax cross entropy loss layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersLossSoftmaxCrossEntropy
type BNNSLayerParametersLossSoftmaxCrossEntropy struct {
	Function unsafe.Pointer // The function that’s used to compute loss.
	I_desc BNNSNDArrayDescriptor // The descriptor of the input.
	Label_smooth float32 // A value that defines the smoothing that the loss function applies to the labels.
	O_desc BNNSNDArrayDescriptor // The descriptor of the output.
	Reduction unsafe.Pointer // The function that’s used to reduce the computed loss.
}/* debug [types.gen.go/struct]: BNNSLayerParametersLossSoftmaxCrossEntropy */

// BNNSLayerParametersLossYolo - A structure that contains the parameters of a You Only Look Once (YOLO) loss layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersLossYolo
type BNNSLayerParametersLossYolo struct {
	Anchor_box_size uintptr // The size of the anchor box.
	Anchors_data []float32 // Maximum IOU for treating as no object.
	Function unsafe.Pointer // The function that’s used to compute loss.
	Huber_delta float32 // A value that’s interpreted as width-height loss.
	I_desc BNNSNDArrayDescriptor // The descriptor of the input.
	No_object_maximum_iou float32 // The value that specifies intersection over union (IOU) that’s the maximum the function treats as not an object.
	Number_of_anchor_boxes uintptr // The number of anchor boxes in each cell.
	Number_of_grid_columns uintptr // The number of columns in the grid.
	Number_of_grid_rows uintptr // The number of rows in the grid.
	O_desc BNNSNDArrayDescriptor // The descriptor of the output.
	Object_minimum_iou float32 // The value that specifies intersection over union (IOU) that’s the minimum the function treats as an object.
	Reduction unsafe.Pointer // The function that’s used to reduce the computed loss (must be sum reduction for YOLO).
	Rescore bool // A Boolean value that determines whether to rescore confidence according to prediction verus ground truth Intersection Over Union (IOU).
	Scale_classification float32 // The value that specifies the classification scaling factor.
	Scale_no_object float32 // The value that specifies the no-object confidence scaling factor.
	Scale_object float32 // The value that specifies the object confidence loss-scaling factor.
	Scale_wh float32 // A Boolean value that determines whether to rescore confidence according to prediction verus ground truth Intersection Over Union (IOU).
	Scale_xy float32 // The value that specifies the x, y loss-scaling factor.
}/* debug [types.gen.go/struct]: BNNSLayerParametersLossYolo */

// BNNSLayerParametersLSTM - A structure that contains the parameters of a long short-term memory (LSTM) layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersLSTM
type BNNSLayerParametersLSTM struct {
	Batch_size uintptr // The number of input and output samples.
	Candidate_gate BNNSLSTMGateDescriptor // The descriptor of the candidate gate, which uses default tanh activation.
	Dropout float32 // The dropout ratio to apply between long short-term memory (LSTM) layers.
	Forget_gate BNNSLSTMGateDescriptor // The descriptor of the forget gate, which uses default sigmoid activation.
	Hidden_activation BNNSActivation // Hidden activation function, which uses default tanh activation.
	Hidden_size uintptr // The number of elements in the hidden state.
	Input_descriptor BNNSLSTMDataDescriptor // Descriptors of the input, hidden input, and cell-state input data.
	Input_gate BNNSLSTMGateDescriptor // The descriptor of the input gate, which uses default sigmoid activation.
	Input_size uintptr // The number of elements in the input.
	Lstm_flags uint32 // Flags that control the behavior of a long short-term memory (LSTM) layer.
	Num_layers uintptr // The number of stacked long short-term memory (LSTM) layers.
	Output_descriptor BNNSLSTMDataDescriptor // Descriptors of the output, hidden output, and cell-state output data.
	Output_gate BNNSLSTMGateDescriptor // The descriptor of the output gate, which uses default sigmoid activation.
	Seq_len uintptr // The size of the sequential input.
	Sequence_descriptor BNNSNDArrayDescriptor // A 1D array of unsigned-integer elements that determines the batch size for each step.
}/* debug [types.gen.go/struct]: BNNSLayerParametersLSTM */

// BNNSLayerParametersMultiheadAttention - A structure that contains the parameters of a multihead attention layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersMultiheadAttention
type BNNSLayerParametersMultiheadAttention struct {
	Add_zero_attn bool // A Boolean value that, if true, adds a row of zeroes to the projected   and   inputs to the calculation.
	Dropout float32 // The seed for the dropout layer’s random number generator.
	Key BNNSMHAProjectionParameters // A projection parameter structure that describes the key-related input parameters and projection.
	Key_attn_bias BNNSNDArrayDescriptor // A 2D tensor that’s added to the value as part of the attention calculation.
	Output BNNSMHAProjectionParameters // A projection parameter structure that describes the output tensor and associated projection.
	Query BNNSMHAProjectionParameters // A projection parameter structure that describes the query-related input parameters and projection.
	Seed uint32 // A random seed for the dropout layer.
	Value BNNSMHAProjectionParameters // A projection parameter structure that describes the value-related input parameters and projection.
	Value_attn_bias BNNSNDArrayDescriptor // An optional   x   2D tensor that’s added as part of the attention calculation.
}/* debug [types.gen.go/struct]: BNNSLayerParametersMultiheadAttention */

// BNNSLayerParametersNormalization - A structure that contains the parameters of a normalization layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersNormalization
type BNNSLayerParametersNormalization struct {
	Activation BNNSActivation // The activation function that the layer applies to the output.
	Beta_desc BNNSNDArrayDescriptor // The descriptor of the beta or bias.
	Epsilon float32 // The epsilon in the computation of the standard deviation.
	Gamma_desc BNNSNDArrayDescriptor // The descriptor of the gamma or scale.
	I_desc BNNSNDArrayDescriptor // The descriptor of the input.
	Momentum float32 // A value, between 0 and 1, the normalization operation uses to update the moving mean and moving variance during training.
	Moving_mean_desc BNNSNDArrayDescriptor // The descriptor of the moving mean.
	Moving_variance_desc BNNSNDArrayDescriptor // The descriptor of the moving variance.
	Normalization_axis uintptr // The axis on which a layer normalization operation starts normalization.
	Num_groups uintptr // The number of groups over which the layer computes normalization statistics.
	O_desc BNNSNDArrayDescriptor // The descriptor of the output.
}/* debug [types.gen.go/struct]: BNNSLayerParametersNormalization */

// BNNSLayerParametersPadding - A structure that contains the parameters of a padding layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersPadding
type BNNSLayerParametersPadding struct {
	I_desc BNNSNDArrayDescriptor // The descriptor of the input.
	O_desc BNNSNDArrayDescriptor // The descriptor of the output.
	Padding_mode unsafe.Pointer // The mode the operation uses to pad.
	Padding_size uintptr // The number of padding elements to add before and after the original data.
	Padding_value uint32 // The value the operation uses to fill the padding area when the mode is constant.
}/* debug [types.gen.go/struct]: BNNSLayerParametersPadding */

// BNNSLayerParametersPermute - A structure that contains the parameters of a permute layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersPermute
type BNNSLayerParametersPermute struct {
	I_desc BNNSNDArrayDescriptor // The descriptor of the input.
	O_desc BNNSNDArrayDescriptor // The descriptor of the output.
	Permutation uintptr // The tuple that defines the permutation.
}/* debug [types.gen.go/struct]: BNNSLayerParametersPermute */

// BNNSLayerParametersPooling - A structure that contains the parameters of a pooling layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersPooling
type BNNSLayerParametersPooling struct {
	Activation BNNSActivation // The activation function that the layer applies to the output.
	Bias BNNSNDArrayDescriptor // The descriptor of the bias.
	I_desc BNNSNDArrayDescriptor // The descriptor of the input.
	K_height uintptr // The height of the kernel.
	K_width uintptr // The width of the kernel.
	O_desc BNNSNDArrayDescriptor // The descriptor of the output.
	Pad uintptr // Asymmetric padding, ignored if   or   are greater than zero.
	Pooling_function unsafe.Pointer // The variable that specifies the pooling function.
	X_dilation_stride uintptr // The width increment between elements in the input image during convolution.
	X_padding uintptr // The width padding, which is the number of virtual zeros added to the left and right of each channel.
	X_stride uintptr // The width increment of the input image.
	Y_dilation_stride uintptr // The height increment between elements in the input image during convolution.
	Y_padding uintptr // The height padding, which is the number of virtual zeros added to the top and bottom of each channel.
	Y_stride uintptr // The height increment of the input image.
}/* debug [types.gen.go/struct]: BNNSLayerParametersPooling */

// BNNSLayerParametersQuantization - A structure that contains the parameters of a quantization layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersQuantization
type BNNSLayerParametersQuantization struct {
	Axis_mask uintptr // A bitmask that defines the axis  to which the function applies scale and bias.
	Bias BNNSNDArrayDescriptor // The descriptor of the bias.
	Function unsafe.Pointer // The quantize function.
	I_desc BNNSNDArrayDescriptor // The descriptor of the input.
	O_desc BNNSNDArrayDescriptor // The descriptor of the output.
	Scale BNNSNDArrayDescriptor // The descriptor of the scale.
}/* debug [types.gen.go/struct]: BNNSLayerParametersQuantization */

// BNNSLayerParametersReduction - A set of parameters that define a reduction layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersReduction
type BNNSLayerParametersReduction struct {
	Epsilon float32 // A value that the operation adds to each element when computing the sum of logarithms.
	I_desc BNNSNDArrayDescriptor // The descriptor of the input.
	O_desc BNNSNDArrayDescriptor // The descriptor of the output.
	Reduce_func unsafe.Pointer // The variable that specifies the reduction function.
	W_desc BNNSNDArrayDescriptor // The descriptor of the weights.
}/* debug [types.gen.go/struct]: BNNSLayerParametersReduction */

// BNNSLayerParametersResize - A structure that contains the parameters of a resize layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersResize
type BNNSLayerParametersResize struct {
	Align_corners bool // A Boolean value that specifies whether to align the corners of the upscaling grid to the center of scaling dimensions instead of to the edges.
	I_desc BNNSNDArrayDescriptor // The descriptor of the input.
	Method unsafe.Pointer // The interpolation method for resizing.
	O_desc BNNSNDArrayDescriptor // The descriptor of the output.
}/* debug [types.gen.go/struct]: BNNSLayerParametersResize */

// BNNSLayerParametersTensorContraction - A structure that contains the parameters of a tensor-contraction layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersTensorContraction
type BNNSLayerParametersTensorContraction struct {
	Alpha float32 // Scaling that the operation applies to the result.
	Beta float32 // A value, that must be either 0.0 or 1.0, you use to scale the existing output before the operation adds it to the result.
	IA_desc BNNSNDArrayDescriptor // The descriptor of input matrix  .
	IB_desc BNNSNDArrayDescriptor // The descriptor of input matrix  .
	O_desc BNNSNDArrayDescriptor // The descriptor of the output.
	Operation unsafe.Pointer // The string that describes the operation.
}/* debug [types.gen.go/struct]: BNNSLayerParametersTensorContraction */

// BNNSLSTMDataDescriptor - A structure that contains the input-output, hidden, and cell state n-dimensional array descriptors for a long short-term memory (LSTM) layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLSTMDataDescriptor
type BNNSLSTMDataDescriptor struct {
	Cell_state_desc BNNSNDArrayDescriptor // The descriptor of the cell-state input-output.
	Data_desc BNNSNDArrayDescriptor // The descriptor of the input-output.
	Hidden_desc BNNSNDArrayDescriptor // The descriptor of the hidden input-output.
}/* debug [types.gen.go/struct]: BNNSLSTMDataDescriptor */

// BNNSLSTMGateDescriptor - A structure that describes a long short-term memory (LSTM) gate layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLSTMGateDescriptor
type BNNSLSTMGateDescriptor struct {
	Activation BNNSActivation // The activation function that the layer applies to the output.
	B_desc BNNSNDArrayDescriptor // The descriptor of the bias.
	Cw_desc BNNSNDArrayDescriptor // The descriptor of the cell weights.
	Hw_desc BNNSNDArrayDescriptor // The descriptor of the hidden weights.
	Iw_desc BNNSNDArrayDescriptor // The descriptor of the input weights.
}/* debug [types.gen.go/struct]: BNNSLSTMGateDescriptor */

// BNNSMHAProjectionParameters - A structure that contains multihead attention projection parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSMHAProjectionParameters
type BNNSMHAProjectionParameters struct {
	Bias BNNSNDArrayDescriptor // The descriptor of the initial projection’s bias.
	Target_desc BNNSNDArrayDescriptor // The descriptor—which is either an input query, key, or value, or an output—of the main target of the operation.
	Weights BNNSNDArrayDescriptor // The descriptor of the initial projection’s weights.
}/* debug [types.gen.go/struct]: BNNSMHAProjectionParameters */

// BNNSNDArrayDescriptor - A structure that describes the shape, stride, data type, and, optionally, the memory location of an n-dimensional array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSNDArrayDescriptor
type BNNSNDArrayDescriptor struct {
	Data unsafe.Pointer // A pointer that is optional and points to the underlying data.
	Data_bias float32 // The bias you use to convert integer and unsigned integer data to floating point.
	Data_scale float32 // The scale you use to convert integer and unsigned integer data to floating point.
	Data_type unsafe.Pointer // The data type of the n-dimensional array.
	Flags unsafe.Pointer // Flags that control some behaviors of the n-dimensional array.
	Layout unsafe.Pointer // The dimension of the n-dimensional array.
	Size uintptr // The number of values in each dimension.
	Stride uintptr // The increment, in values, between consecutive elements in each dimension.
	Table_data unsafe.Pointer // The lookup table for indexed data types.
	Table_data_type unsafe.Pointer // The data type of the lookup table.
}/* debug [types.gen.go/struct]: BNNSNDArrayDescriptor */

// BNNSOptimizerAdamFields - A structure that contains the fields of an Adam optimizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSOptimizerAdamFields
type BNNSOptimizerAdamFields struct {
	Beta1 float32 // A value that specifies the first moment constant in the range 0 to 1.
	Beta2 float32 // A value that specifies the second moment constant in the range 0 to 1.
	Clip_gradients bool // A Boolean value that specifies whether to clip the gradient between minimum and maximum values.
	Clip_gradients_max float32 // The values for the maximum gradient.
	Clip_gradients_min float32 // The values for the minimum gradient.
	Epsilon float32 // An addition for the division in the parameter update stage.
	Gradient_scale float32 // A value that specifies the gradient scaling factor.
	Learning_rate float32 // A value that specifies the learning rate.
	Regularization_func unsafe.Pointer // The variable that specifies the regularization function.
	Regularization_scale float32 // A value that specifies the regularization scaling factor.
	Time_step float32 // A value that represents the optimizer’s current time and you’re responsible for updating after optimizing all the layer parameters in your network.
}/* debug [types.gen.go/struct]: BNNSOptimizerAdamFields */

// BNNSOptimizerAdamWithClippingFields - A structure that contains the fields of an Adam or AdamW optimizer that optionally clips the gradient by value or by norm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSOptimizerAdamWithClippingFields
type BNNSOptimizerAdamWithClippingFields struct {
	Beta1 float32 // A value that specifies the first moment constant in the range 0 to 1.
	Beta2 float32 // A value that specifies the second moment constant in the range 0 to 1.
	Clip_gradients_max float32 // The maximum clipping value for clipping by value.
	Clip_gradients_max_norm float32 // The maximum Euclidean norm for clipping by norm and clipping by global norm.
	Clip_gradients_min float32 // The minimum clipping value for clipping by value.
	Clip_gradients_use_norm float32 // An optional value for a known Euclidean norm for clipping by global norm.
	Clipping_func unsafe.Pointer // The clipping function.
	Epsilon float32 // An addition for the division in the parameter update stage.
	Gradient_scale float32 // A value that specifies the gradient scaling factor.
	Learning_rate float32 // A value that specifies the learning rate.
	Regularization_func unsafe.Pointer // The variable that specifies the regularization function.
	Regularization_scale float32 // A value that specifies the regularization scaling factor.
	Time_step float32 // A value that’s at least 1 and represents the optimizer’s current time.
}/* debug [types.gen.go/struct]: BNNSOptimizerAdamWithClippingFields */

// BNNSOptimizerRMSPropFields - A structure that contains the fields of a root mean square propagation (RMSProp) optimizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSOptimizerRMSPropFields
type BNNSOptimizerRMSPropFields struct {
	Alpha float32 // A constant that specifies smoothing.
	Centered bool // A Boolean value that specifies whether to use the centered variant.
	Clip_gradients bool // A Boolean value that specifies whether to clip the gradient between minimum and maximum values.
	Clip_gradients_max float32 // The values for the maximum gradient.
	Clip_gradients_min float32 // The values for the minimum gradient.
	Epsilon float32 // A term that the optimizer adds to the denominator.
	Gradient_scale float32 // A value that specifies the gradient scaling factor.
	Learning_rate float32 // A value that specifies the learning rate.
	Momentum float32 // The rate of momentum decay.
	Regularization_func unsafe.Pointer // The variable that specifies the regularization function.
	Regularization_scale float32 // A value that specifies the regularization scaling factor.
}/* debug [types.gen.go/struct]: BNNSOptimizerRMSPropFields */

// BNNSOptimizerRMSPropWithClippingFields - A structure that contains the fields of a root mean square propagation (RMSProp) optimizer that optionally clips the gradient by value or by norm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSOptimizerRMSPropWithClippingFields
type BNNSOptimizerRMSPropWithClippingFields struct {
	Alpha float32 // A constant that specifies smoothing.
	Centered bool // A Boolean value that specifies whether to use the centered variant.
	Clip_gradients_max float32 // The maximum clipping value for clipping by value.
	Clip_gradients_max_norm float32 // The maximum Euclidean norm for clipping by norm and clipping by global norm.
	Clip_gradients_min float32 // The minimum clipping value for clipping by value.
	Clip_gradients_use_norm float32 // An optional value for a known Euclidean norm for clipping by global norm.
	Clipping_func unsafe.Pointer // The clipping function.
	Epsilon float32 // A term that the optimizer adds to the denominator.
	Gradient_scale float32 // A value that specifies the gradient scaling factor.
	Learning_rate float32 // A value that specifies the learning rate.
	Momentum float32 // The rate of momentum decay.
	Regularization_func unsafe.Pointer // The variable that specifies the regularization function.
	Regularization_scale float32 // A value that specifies the regularization scaling factor.
}/* debug [types.gen.go/struct]: BNNSOptimizerRMSPropWithClippingFields */

// BNNSOptimizerSGDMomentumFields - A structure that contains the fields of a stochastic gradient descent (SGD) with momentum optimizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSOptimizerSGDMomentumFields
type BNNSOptimizerSGDMomentumFields struct {
	Clip_gradients bool // A Boolean value that specifies whether to clip the gradient between minimum and maximum values.
	Clip_gradients_max float32 // The values for the maximum gradient.
	Clip_gradients_min float32 // The values for the minimum gradient.
	Gradient_scale float32 // A value that specifies the gradient scaling factor.
	Learning_rate float32 // A value that specifies the learning rate.
	Momentum float32 // The rate of momentum decay.
	Nesterov bool // A Boolean value that specifies whether to use Nesterov momentum update.
	Regularization_func unsafe.Pointer // The variable that specifies the regularization function.
	Regularization_scale float32 // A value that specifies the regularization scaling factor.
	Sgd_momentum_variant unsafe.Pointer // The variable that specifies the momentum variant.
}/* debug [types.gen.go/struct]: BNNSOptimizerSGDMomentumFields */

// BNNSOptimizerSGDMomentumWithClippingFields - A structure that contains the fields of a stochastic gradient descent (SGD) with momentum optimizer that optionally clips the gradient by value or by norm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSOptimizerSGDMomentumWithClippingFields
type BNNSOptimizerSGDMomentumWithClippingFields struct {
	Clip_gradients_max float32 // The maximum clipping value for clipping by value.
	Clip_gradients_max_norm float32 // The maximum Euclidean norm for clipping by norm and clipping by global norm.
	Clip_gradients_min float32 // The minimum clipping value for clipping by value.
	Clip_gradients_use_norm float32 // An optional value for a known Euclidean norm for clipping by global norm.
	Clipping_func unsafe.Pointer // The clipping function.
	Gradient_scale float32 // A value that specifies the gradient scaling factor.
	Learning_rate float32 // A value that specifies the learning rate.
	Momentum float32 // The rate of momentum decay.
	Nesterov bool // A Boolean value that specifies whether to use Nesterov momentum update.
	Regularization_func unsafe.Pointer // The variable that specifies the regularization function.
	Regularization_scale float32 // A value that specifies the regularization scaling factor.
	Sgd_momentum_variant unsafe.Pointer // The variable that specifies the momentum variant.
}/* debug [types.gen.go/struct]: BNNSOptimizerSGDMomentumWithClippingFields */

// BNNSPoolingLayerParameters - A structure containing pooling layer parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSPoolingLayerParameters
type BNNSPoolingLayerParameters struct {
	Activation BNNSActivation // The layer activation function.
	Bias BNNSLayerData // Layer bias, one for each output channel.
	In_channels uintptr // The number of input channels.
	K_height uintptr // The height of the convolution kernel.
	K_width uintptr // The width of the convolution kernel.
	Out_channels uintptr // The number of output channels.
	Pooling_function unsafe.Pointer // The pooling function to apply to each sample.
	X_padding uintptr // The X padding.
	X_stride uintptr // The X increment in the input image.
	Y_padding uintptr // The Y padding.
	Y_stride uintptr // The Y increment in the input image.
}/* debug [types.gen.go/struct]: BNNSPoolingLayerParameters */

// BNNSSparsityParameters
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSSparsityParameters
type BNNSSparsityParameters struct {
	Flags uint64
	Sparsity_ratio uint32
	Sparsity_type unsafe.Pointer
	Target_system unsafe.Pointer
}/* debug [types.gen.go/struct]: BNNSSparsityParameters */

// BNNSTensor - A structure that describes the shape, stride, data type, and, optionally, the memory location of an n-dimensional array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSTensor
type BNNSTensor struct {
	Data unsafe.Pointer // A pointer to the memory that contains the tensor values.
	Data_size_in_bytes uintptr // The size, in bytes, of the memory that contains the tensor values.
	Data_type unsafe.Pointer // The data type of the tensor.
	Name unsafe.Pointer // An optional name for the tensor that you can use for debugging.
	Rank uint8 // The rank of the tensor.
	Shape unsafe.Pointer // A tuple of unsigned-integer elements that specify the size of the tensor.
	Stride unsafe.Pointer // A tuple of unsigned-integer elements that specify the stride of the tensor.
}/* debug [types.gen.go/struct]: BNNSTensor */

// BNNSVectorDescriptor
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSVectorDescriptor
type BNNSVectorDescriptor struct {
	Data_bias float32
	Data_scale float32
	Data_type unsafe.Pointer
	Size uintptr
}/* debug [types.gen.go/struct]: BNNSVectorDescriptor */

// DenseMatrix_Double - A structure that contains a dense matrix of double-precision, floating-point values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/DenseMatrix_Double
type DenseMatrix_Double struct {
	Attributes SparseAttributes_t // The attributes of the matrix, such as whether it’s symmetrical or triangular.
	ColumnCount int // The number of columns in the matrix.
	ColumnStride int // The stride between matrix columns, in elements.
	Data []float64 // The array of double-precision, floating-point values in column-major order.
	RowCount int // The number of rows in the matrix.
}/* debug [types.gen.go/struct]: DenseMatrix_Double */

// DenseMatrix_Float - A structure that contains a dense matrix of single-precision, floating-point values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/DenseMatrix_Float
type DenseMatrix_Float struct {
	Attributes SparseAttributes_t // The attributes of the matrix, such as whether it’s symmetrical or triangular.
	ColumnCount int // The number of columns in the matrix.
	ColumnStride int // The stride between matrix columns, in elements.
	Data []float32 // The array of single-precision, floating-point values in column-major order.
	RowCount int // The number of rows in the matrix.
}/* debug [types.gen.go/struct]: DenseMatrix_Float */

// DenseVector_Double - A structure that contains a dense vector of double-precision, floating-point values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/DenseVector_Double
type DenseVector_Double struct {
	Count int // The number of items in the vector.
	Data []float64 // The array of double-precision, floating-point values.
}/* debug [types.gen.go/struct]: DenseVector_Double */

// DenseVector_Float - A structure that contains a dense vector of single-precision, floating-point values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/DenseVector_Float
type DenseVector_Float struct {
	Count int // The number of items in the vector.
	Data []float32 // The array of single-precision, floating-point values.
}/* debug [types.gen.go/struct]: DenseVector_Float */

// DSPComplex - A structure that represents a single-precision complex value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/DSPComplex
type DSPComplex struct {
	Imag float32 // The imaginary part of the value.
	Real float32 // The real part of the value.
}/* debug [types.gen.go/struct]: DSPComplex */

// DSPDoubleComplex - A structure that represents a double-precision complex value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/DSPDoubleComplex
type DSPDoubleComplex struct {
	Imag float64 // The imaginary part of the value.
	Real float64 // The real part of the value.
}/* debug [types.gen.go/struct]: DSPDoubleComplex */

// DSPDoubleSplitComplex - A structure that represents a double-precision complex vector with the real and imaginary parts stored in separate arrays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/DSPDoubleSplitComplex
type DSPDoubleSplitComplex struct {
	Imagp []float64 // An array of imaginary parts of the complex numbers.
	Realp []float64 // An array of real parts of the complex numbers.
}/* debug [types.gen.go/struct]: DSPDoubleSplitComplex */

// DSPSplitComplex - A structure that represents a single-precision complex vector with the real and imaginary parts stored in separate arrays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/DSPSplitComplex
type DSPSplitComplex struct {
	Imagp []float32 // An array of imaginary parts of the complex numbers.
	Realp []float32 // An array of real parts of the complex numbers.
}/* debug [types.gen.go/struct]: DSPSplitComplex */

// quadrature_integrate_function
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/quadrature_integrate_function
type quadrature_integrate_function struct {
	Fun Quadrature_function_array
	Fun_arg unsafe.Pointer
}/* debug [types.gen.go/struct]: quadrature_integrate_function */

// quadrature_integrate_options
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/quadrature_integrate_options
type quadrature_integrate_options struct {
	Abs_tolerance float64
	Integrator unsafe.Pointer
	Max_intervals uintptr
	Qag_points_per_interval uintptr
	Rel_tolerance float64
}/* debug [types.gen.go/struct]: quadrature_integrate_options */

// SparseAttributes_t - A structure that represents the attributes of a matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseAttributes_t
type SparseAttributes_t struct {
	Kind unsafe.Pointer // An eumeration that specifies whether the matrix is ordinary, unit-triangular, triangular, or symmetric.
	Transpose bool // A Boolean value that specifies whether to implicitly transpose the matrix.
	Triangle unsafe.Pointer // An enumeration that specifies which triangle unit-triangular, triangular, and symmetric matrices need to use.
}/* debug [types.gen.go/struct]: SparseAttributes_t */

// SparseCGOptions - Options for creating a conjugate gradient (CG) method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseCGOptions
type SparseCGOptions struct {
	Atol float64 // The absolute convergence tolerance.
	MaxIterations int // The maximum number of iterations to perform.
	ReportError unsafe.Pointer // An optional error-reporting routine.
	ReportStatus unsafe.Pointer // The function to report status.
	Rtol float64 // The relative convergence tolerance.
}/* debug [types.gen.go/struct]: SparseCGOptions */

// SparseGMRESOptions - Options for creating a generalized minimal residual (GMRES) method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseGMRESOptions
type SparseGMRESOptions struct {
	Atol float64 // The absolute convergence tolerance.
	MaxIterations int // The maximum number of iterations to perform.
	Nvec int // The number of orthogonal vectors the operation maintains.
	ReportError unsafe.Pointer // An optional error-reporting routine.
	ReportStatus unsafe.Pointer // The function to report status.
	Rtol float64 // The relative convergence tolerance.
	Variant unsafe.Pointer // The exact variant of GMRES to implement.
}/* debug [types.gen.go/struct]: SparseGMRESOptions */

// SparseIterativeMethod - The base type for all iterative methods.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseIterativeMethod
type SparseIterativeMethod struct {
	Method int // The iterative method this structure represents.
	Options unsafe.Pointer // The options for the method.
	Base unsafe.Pointer
	Cg SparseCGOptions // Conjugate Gradient Options.
	Gmres SparseGMRESOptions // Right-preconditioned (F/DQ)GMRES Parameters Options.
	Lsmr SparseLSMROptions // LSMR is MINRES specialised for solving least squares.
	Padding unsafe.Pointer
}/* debug [types.gen.go/struct]: SparseIterativeMethod */

// SparseLSMROptions - Options for creating a least squares minimum residual method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseLSMROptions
type SparseLSMROptions struct {
	Atol float64 // The absolute tolerance (default test) or   tolerance (Fong-Saunders test).
	Btol float64 // The   tolerance (Fong-Saunders test only).
	ConditionLimit float64 // The condition number limit (Fong-Saunders test only).
	ConvergenceTest unsafe.Pointer // The convergence test to use for iterative solve methods.
	Lambda float64 // The damping parameter lambda for regularized least squares.
	MaxIterations int // The maximum number of iterations.
	Nvec int // The number of vectors to use for local reorthogonalization.
	ReportError unsafe.Pointer // An optional error-reporting routine.
	ReportStatus unsafe.Pointer // An optional status-reporting routine.
	Rtol float64 // The relative convergence tolerance (default test only).
}/* debug [types.gen.go/struct]: SparseLSMROptions */

// SparseMatrix_Double - A structure that contains a sparse matrix of double-precision, floating-point values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseMatrix_Double
type SparseMatrix_Double struct {
	Data []float64 // The array of contiguous values in the nonzero blocks of the matrix.
	Structure SparseMatrixStructure // The sparsity structure of the matrix.
}/* debug [types.gen.go/struct]: SparseMatrix_Double */

// SparseMatrix_Float - A structure that contains a sparse matrix of single-precision, floating-point values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseMatrix_Float
type SparseMatrix_Float struct {
	Data []float32 // The array of contiguous values in the nonzero blocks of the matrix.
	Structure SparseMatrixStructure // The sparsity structure of the matrix.
}/* debug [types.gen.go/struct]: SparseMatrix_Float */

// SparseMatrixStructure - A description of the sparsity structure of a sparse matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseMatrixStructure
type SparseMatrixStructure struct {
	Attributes SparseAttributes_t // The attributes of the matrix, such as whether it’s symmetrical or triangular.
	BlockSize uint8 // The block size of the matrix.
	ColumnCount int // The number of columns in the matrix.
	ColumnStarts unsafe.Pointer // The starting index for each column in the row indices array.
	RowCount int // The number of rows in the matrix.
	RowIndices []int // The row indices of the matrix.
}/* debug [types.gen.go/struct]: SparseMatrixStructure */

// SparseNumericFactorOptions - A structure that contains options that affect the numerical stage of a sparse factorization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseNumericFactorOptions
type SparseNumericFactorOptions struct {
	Control unsafe.Pointer // The flags that control the computation.
	PivotTolerance float64 // The pivot tolerance that threshold partial pivoting uses.
	Scaling unsafe.Pointer // An array that scales the matrix before factorization.
	ScalingMethod unsafe.Pointer // The scaling method.
	ZeroTolerance float64 // The zero tolerance that some pivoting modes use.
}/* debug [types.gen.go/struct]: SparseNumericFactorOptions */

// SparseOpaqueFactorization_Double - A structure that represents the factorization of a matrix of double-precision, floating-point values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseOpaqueFactorization_Double
type SparseOpaqueFactorization_Double struct {
	Attributes SparseAttributes_t // The attributes of a factorization object.
	NumericFactorization unsafe.Pointer // The pointer to a private internal representation of a numeric factor.
	SolveWorkspaceRequiredPerRHS uintptr // The required size of the per-right-hand-side workspace for a call to a sparse solve function.
	SolveWorkspaceRequiredStatic uintptr // The required size of the static workspace for a call to a sparse solve function.
	Status unsafe.Pointer // The status of the factorization object.
	SymbolicFactorization SparseOpaqueSymbolicFactorization // The symbolic factorization that this numeric factorization depends on.
	UserFactorStorage bool // A Boolean value that indicates whether user-provided storage backs this object.
}/* debug [types.gen.go/struct]: SparseOpaqueFactorization_Double */

// SparseOpaqueFactorization_Float - A structure that represents the factorization of a matrix of single-precision, floating-point values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseOpaqueFactorization_Float
type SparseOpaqueFactorization_Float struct {
	Attributes SparseAttributes_t // The attributes of a factorization object.
	NumericFactorization unsafe.Pointer // The pointer to a private internal representation of a numeric factor.
	SolveWorkspaceRequiredPerRHS uintptr // The required size of the per-right-hand-side workspace for a call to a sparse solve function.
	SolveWorkspaceRequiredStatic uintptr // The required size of the static workspace for a call to a sparse solve function.
	Status unsafe.Pointer // The status of the factorization object.
	SymbolicFactorization SparseOpaqueSymbolicFactorization // The symbolic factorization that this numeric factorization depends on.
	UserFactorStorage bool // A Boolean value that indicates whether user-provided storage backs this object.
}/* debug [types.gen.go/struct]: SparseOpaqueFactorization_Float */

// SparseOpaquePreconditioner_Double - A structure that represents a double-precision preconditioner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseOpaquePreconditioner_Double
type SparseOpaquePreconditioner_Double struct {
	Apply unsafe.Pointer // A function that calculates  , where   is the preconditioner.
	Mem unsafe.Pointer // The unaltered memory pointer that passes as the first parameter of the apply function.
	Type unsafe.Pointer // The preconditioner type.
}/* debug [types.gen.go/struct]: SparseOpaquePreconditioner_Double */

// SparseOpaquePreconditioner_Float - A structure that represents a single-precision preconditioner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseOpaquePreconditioner_Float
type SparseOpaquePreconditioner_Float struct {
	Apply unsafe.Pointer // A function that calculates  , where   is the preconditioner.
	Mem unsafe.Pointer // The unaltered memory pointer that passes as the first parameter of the apply function.
	Type unsafe.Pointer // The preconditioner type.
}/* debug [types.gen.go/struct]: SparseOpaquePreconditioner_Float */

// SparseOpaqueSubfactor_Double - Represents a sub-factor of the factorization (for example,  
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseOpaqueSubfactor_Double
type SparseOpaqueSubfactor_Double struct {
	Attributes SparseAttributes_t // A type representing the attributes of a matrix.
	Contents unsafe.Pointer // Types of sub-factor object.
	Factor SparseOpaqueFactorization_Double // A semi-opaque type representing a matrix factorization in double.
	WorkspaceRequiredPerRHS uintptr
	WorkspaceRequiredStatic uintptr
}/* debug [types.gen.go/struct]: SparseOpaqueSubfactor_Double */

// SparseOpaqueSubfactor_Float - Represents a sub-factor of the factorization (for example,  
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseOpaqueSubfactor_Float
type SparseOpaqueSubfactor_Float struct {
	Attributes SparseAttributes_t // A type representing the attributes of a matrix.
	Contents unsafe.Pointer // Types of sub-factor object.
	Factor SparseOpaqueFactorization_Float // A semi-opaque type representing a matrix factorization in float.
	WorkspaceRequiredPerRHS uintptr
	WorkspaceRequiredStatic uintptr
}/* debug [types.gen.go/struct]: SparseOpaqueSubfactor_Float */

// SparseOpaqueSymbolicFactorization - A semi-opaque type that represents symbolic matrix factorization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseOpaqueSymbolicFactorization
type SparseOpaqueSymbolicFactorization struct {
	Attributes SparseAttributes_t // The attributes of the factorization.
	BlockSize uint8 // The block size.
	ColumnCount int // The number of columns.
	Factorization unsafe.Pointer // A pointer to a private internal representation of the symbolic factor.
	FactorSize_Double uintptr // Minimum size, in bytes, required to store numerical factors in doubles.
	FactorSize_Float uintptr // Minimum size, in bytes, required to store numerical factors in float.
	RowCount int // The number of rows.
	Status unsafe.Pointer // The status of the factorization.
	Type unsafe.Pointer // The factorization type.
	WorkspaceSize_Double uintptr // Size, in bytes, of workspace required to perform numerical factorization in doubles.
	WorkspaceSize_Float uintptr // Size, in bytes, of workspace required to perform numerical factorization in floats.
}/* debug [types.gen.go/struct]: SparseOpaqueSymbolicFactorization */

// SparseSymbolicFactorOptions - A structure that contains options that affect the symbolic stage of a sparse factorization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseSymbolicFactorOptions
type SparseSymbolicFactorOptions struct {
	Control unsafe.Pointer // The flags that control the computation.
	Free unsafe.Pointer // The function for freeing allocated storage.
	IgnoreRowsAndColumns []int // An array that contains row and column indices to ignore.
	Malloc unsafe.Pointer // The function for allocating any necessary storage.
	Order []int // The user-supplied array for ordering.
	OrderMethod unsafe.Pointer // The ordering algorithm.
	ReportError unsafe.Pointer // The function for reporting parameter errors.
}/* debug [types.gen.go/struct]: SparseSymbolicFactorOptions */

// vDSP_int24 - A data structure that holds a 24-bit signed integer value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_int24
type vDSP_int24 struct {
	Bytes uint8 // The bytes that represent the value.
}/* debug [types.gen.go/struct]: vDSP_int24 */

// vDSP_uint24 - A data structure that holds a 24-bit unsigned integer value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_uint24
type vDSP_uint24 struct {
	Bytes uint8 // The bytes that represent the value.
}/* debug [types.gen.go/struct]: vDSP_uint24 */

// vImage_AffineTransform - A structure for values that represent an affine transformation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImage_AffineTransform
type vImage_AffineTransform struct {
	A float32 // The entry at position   in the matrix.
	B float32 // The entry at position   in the matrix.
	C float32 // The entry at position   in the matrix.
	D float32 // The entry at position   in the matrix.
	Tx float32 // The entry at position   in the matrix.
	Ty float32 // The entry at position   in the matrix.
}/* debug [types.gen.go/struct]: vImage_AffineTransform */

// vImage_AffineTransform_Double - A structure for values that represent a double-precision affine transformation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImage_AffineTransform_Double
type vImage_AffineTransform_Double struct {
	A float64 // The entry at position   in the matrix.
	B float64 // The entry at position   in the matrix.
	C float64 // The entry at position   in the matrix.
	D float64 // The entry at position   in the matrix.
	Tx float64 // The entry at position   in the matrix.
	Ty float64 // The entry at position   in the matrix.
}/* debug [types.gen.go/struct]: vImage_AffineTransform_Double */

// vImage_ARGBToYpCbCr - The information that describes the conversion from ARGB to YpCbCr.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImage_ARGBToYpCbCr
type vImage_ARGBToYpCbCr struct {
	Opaque uint8
}/* debug [types.gen.go/struct]: vImage_ARGBToYpCbCr */

// vImage_ARGBToYpCbCrMatrix - The 3 x 3 matrix that the vImage library uses to convert from RGB to YpCbCr.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImage_ARGBToYpCbCrMatrix
type vImage_ARGBToYpCbCrMatrix struct {
	B_Cb_R_Cr float32 // The   value in the conversion matrix.
	B_Cr float32 // The   value in the conversion matrix.
	B_Yp float32 // The   value in the conversion matrix.
	G_Cb float32 // The   value in the conversion matrix.
	G_Cr float32 // The   value in the conversion matrix.
	G_Yp float32 // The   value in the conversion matrix.
	R_Cb float32 // The   value in the conversion matrix.
	R_Yp float32 // The   value in the conversion matrix.
}/* debug [types.gen.go/struct]: vImage_ARGBToYpCbCrMatrix */

// vImage_Buffer - An image buffer that stores an image’s pixel data, dimensions, and row stride.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImage_Buffer
type vImage_Buffer struct {
	Data unsafe.Pointer // A pointer to the top-left pixel of the image.
	Height VImagePixelCount // The height of the image, in pixels.
	RowBytes uintptr // The distance, in bytes, between the start of one pixel row and the next in an image, including any unused space between them.
	Width VImagePixelCount // The width of the image, in pixels.
}/* debug [types.gen.go/struct]: vImage_Buffer */

// vImage_CGImageFormat - The description of a Core Graphics image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImage_CGImageFormat
type vImage_CGImageFormat struct {
	BitmapInfo BitmapInfo // The component information that describes the color channels.
	BitsPerComponent uint32 // The number of bits that represents one channel of data in one pixel.
	BitsPerPixel uint32 // The number of bits that represents one pixel.
	ColorSpace ColorSpaceRef // A description of the position of the pixel data in the image, relative to a reference XYZ color space.
	Decode *float64 // The decode array for the image.
	RenderingIntent ColorRenderingIntent // A rendering intent constant that specifies how Core Graphics handles colors that aren’t within the destination color space gamut.
	Version uint32 // The version number.
}/* debug [types.gen.go/struct]: vImage_CGImageFormat */

// vImage_PerpsectiveTransform - A projective-transformation matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImage_PerpsectiveTransform
type vImage_PerpsectiveTransform struct {
	A float32 // The top-left cell in the 3 x 3 transformation matrix.
	B float32 // The top-middle cell in the 3 x 3 transformation matrix.
	C float32 // The middle-left cell in the 3 x 3 transformation matrix.
	D float32 // The middle-middle cell in the 3 x 3 transformation matrix.
	Tx float32 // The x-coordinate translation.
	Ty float32 // The y-coordinate translation.
	V float32 // The homogeneous scale factor.
	Vx float32 // The x-component of the projective vector.
	Vy float32 // The y-component of the projective vector.
}/* debug [types.gen.go/struct]: vImage_PerpsectiveTransform */

// vImage_YpCbCrPixelRange - The description of range and clamping information for YpCbCr pixel formats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImage_YpCbCrPixelRange
type vImage_YpCbCrPixelRange struct {
	CbCr_bias int32 // The encoding for   for this video format.
	CbCrMax int32 // The encoding of the maximum allowed   value.
	CbCrMin int32 // The encoding of the minimum allowed   value.
	CbCrRangeMax int32 // The encoding for   for this video format.
	Yp_bias int32 // The encoding for   for this video format (varies by bit depth).
	YpMax int32 // The encoding for the maximum allowed Y’ value.
	YpMin int32 // The encoding of the minimum allowed Y’ value.
	YpRangeMax int32 // The encoding for   for this video format.
}/* debug [types.gen.go/struct]: vImage_YpCbCrPixelRange */

// vImage_YpCbCrToARGB - The information that describes the conversion from YpCbCr to ARGB.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImage_YpCbCrToARGB
type vImage_YpCbCrToARGB struct {
	Opaque uint8 // The bytes of the opaque representation.
}/* debug [types.gen.go/struct]: vImage_YpCbCrToARGB */

// vImage_YpCbCrToARGBMatrix - The 3 x 3 matrix that the vImage library uses to convert from YpCbCr to RGB.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImage_YpCbCrToARGBMatrix
type vImage_YpCbCrToARGBMatrix struct {
	Cb_B float32 // The   value in the conversion matrix.
	Cb_G float32 // The   value in the conversion matrix.
	Cr_G float32 // The   value in the conversion matrix.
	Cr_R float32 // The   value in the conversion matrix.
	Yp float32 // The   value in the conversion matrix.
}/* debug [types.gen.go/struct]: vImage_YpCbCrToARGBMatrix */

// vImageChannelDescription - A description of the range and clamp limits for a pixel format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageChannelDescription
type vImageChannelDescription struct {
	Full float64 // The encoding for the value one.
	Max float64 // The maximum encoded value.
	Min float64 // The minimum encoded value.
	Zero float64 // The encoding for the value zero.
}/* debug [types.gen.go/struct]: vImageChannelDescription */

// vImageRGBPrimaries - A representation of the chromaticity of primaries that define a color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageRGBPrimaries
type vImageRGBPrimaries struct {
	Blue_x float32 // The blue   value according to the CIE 1931 color space.
	Blue_y float32 // The blue   value according to the CIE 1931 color space.
	Green_x float32 // The green   value according to the CIE 1931 color space.
	Green_y float32 // The green_ _  value according to the CIE 1931 color space.
	Red_x float32 // The red   value according to the CIE 1931 color space.
	Red_y float32 // The red   value according to the CIE 1931 color space.
	White_x float32 // The white point   value according to the CIE 1931 color space.
	White_y float32 // The white point   value according to the CIE 1931 color space.
}/* debug [types.gen.go/struct]: vImageRGBPrimaries */

// vImageTransferFunction - A transfer function to convert from linear to nonlinear RGB.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageTransferFunction
type vImageTransferFunction struct {
	C0 float64 // The   in the transfer function.
	C1 float64 // The   in the transfer function.
	C2 float64 // The   in the transfer function.
	C3 float64 // The   in the transfer function.
	C4 float64 // The   in the transfer function.
	C5 float64 // The   in the transfer function.
	Cutoff float64 // The   in the transfer function.
	Gamma float64 // The   in the transfer function.
}/* debug [types.gen.go/struct]: vImageTransferFunction */

// vImageWhitePoint - A representation of a white point according to the CIE 1931 color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageWhitePoint
type vImageWhitePoint struct {
	White_x float32 // The white point   value according to the CIE 1931 color space.
	White_y float32 // The white point   value according to the CIE 1931 color space.
}/* debug [types.gen.go/struct]: vImageWhitePoint */





