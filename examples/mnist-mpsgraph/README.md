# MNIST Neural Network Training with MPSGraph

This example demonstrates training a neural network on the MNIST dataset using Apple's MPSGraph framework for GPU-accelerated machine learning.

## Features

- **Real MNIST Data**: Downloads and caches the MNIST handwritten digit dataset
- **GPU Training**: Uses Metal Performance Shaders Graph (MPSGraph) for GPU acceleration
- **Live Visualization**: Real-time display of training progress with digit samples
- **Interactive UI**: macOS native window showing accuracy, loss, and prediction results

## Architecture

### Neural Network
- **Input**: 784 features (28×28 pixel images)
- **Hidden Layer**: 256 neurons with ReLU activation
- **Output**: 10 classes (digits 0-9)
- **Loss**: Softmax cross-entropy
- **Optimizer**: Stochastic Gradient Descent (SGD)

### Training Parameters
- Batch size: 40
- Learning rate: 0.01
- Iterations: 5,000
- Hidden size: 256

## Files

- `main.go` - Application entry point and UI setup
- `trainer.go` - Neural network training logic with MPSGraph
- `mnist_loader.go` - MNIST dataset downloading and parsing
- `visualization.go` - Real-time training visualization

## Usage

### Build and Run
```bash
go build
./mnist-mpsgraph
```

### Command-line Options

```bash
# Run in headless test mode
./mnist-mpsgraph -e2e

# Add delay between iterations for slower visualization
./mnist-mpsgraph -delay 100ms

# Update UI less frequently (every N iterations)
./mnist-mpsgraph -refresh 10
```

## How It Works

1. **Data Loading**: Downloads MNIST dataset from Google Cloud Storage (cached in `/tmp/mnist_cache`)
2. **Metal Setup**: Creates Metal device and command queue for GPU operations
3. **Graph Construction**: Builds computational graph with MPSGraph API
   - Placeholders for inputs and labels
   - Variable tensors for weights and biases
   - Forward pass operations (matrix multiply, add, ReLU)
   - Loss computation (softmax cross-entropy)
4. **Automatic Differentiation**: Computes gradients using MPSGraph's gradient API
5. **Training Loop**: Iteratively:
   - Samples random batch from dataset
   - Runs forward and backward pass on GPU
   - Updates weights with SGD
   - Reports metrics to UI

## Frameworks Used

- **AppKit**: macOS UI framework
- **Foundation**: Core data structures
- **Metal**: GPU device management
- **MPS**: Metal Performance Shaders base types
- **MPSGraph**: High-level graph-based neural network API

## Implementation Notes

### Metal Device Creation
This example includes a workaround for `metal.CreateSystemDefaultDevice()` which may not be available in all darwinkit versions:

```go
func createSystemDefaultDevice() metal.DeviceObject {
    metalLib, _ := purego.Dlopen("/System/Library/Frameworks/Metal.framework/Metal", ...)
    var mtlCreateSystemDefaultDevice func() unsafe.Pointer
    purego.RegisterLibFunc(&mtlCreateSystemDefaultDevice, metalLib, "MTLCreateSystemDefaultDevice")
    device := mtlCreateSystemDefaultDevice()
    return metal.DeviceObject{Object: objc.ObjectFrom(device)}
}
```

### Memory Safety
- Uses `unsafe.Slice` to convert Go slices to byte slices for Metal
- Manual object retention with `objc.Retain()` for Objective-C objects
- Proper cleanup of Metal resources

## Expected Results

After 5,000 iterations, the model typically achieves:
- **Test Accuracy**: ~90-95%
- **Training Loss**: <0.3

The visualization shows:
- 16 sample digits (4×4 grid)
- Green labels for correct predictions
- Red labels for incorrect predictions
- Confidence scores for each prediction
- Smoothed accuracy and loss metrics

## Requirements

- macOS with Metal support
- Go 1.24+
- Internet connection (for first run to download MNIST)
- GPU with Metal support

## See Also

- Original darwinkit example: `/Volumes/tmc/go/src/github.com/progrium/darwinkit/macos/_examples/mnist-mpsgraph/`
- Apple MPSGraph documentation: https://developer.apple.com/documentation/metalperformanceshadersgraph
