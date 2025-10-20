# QuartzCore (Core Animation) Framework Example

This example demonstrates the QuartzCore/Core Animation framework for hardware-accelerated graphics and animations in Go.

## What it demonstrates

- CALayer creation and properties
- Animation types and capabilities
- Animatable properties
- Timing functions
- Layer types and specializations
- Performance features
- Common use cases
- Framework integration

## Running the example

```bash
go run main.go
# or with e2e flag
go run main.go -e2e
```

## Key Concepts

### QuartzCore/Core Animation

Core Animation (QuartzCore framework) is Apple's graphics rendering and animation infrastructure:
- Hardware-accelerated layer compositing
- Smooth 60fps animations with minimal code
- Rich layer hierarchy system
- Integration with all graphics frameworks
- Runs on separate rendering thread

### CALayer - The Foundation

Every view has a backing CALayer:

```go
// Create a layer
layer := quartzcore.NewLayer()

// Set properties
layer.SetFrame(rect)
layer.SetBackgroundColor(color)
layer.SetOpacity(0.8)
layer.SetCornerRadius(10.0)
```

### Layer Hierarchy

Layers form a tree structure:

```go
// Add sublayers
parentLayer.AddSublayer(childLayer)

// Access hierarchy
sublayers := layer.Sublayers()
superlayer := layer.Superlayer()

// Remove from hierarchy
layer.RemoveFromSuperlayer()
```

## Animation Types

### CABasicAnimation

Animate a property from one value to another:

```go
// Create animation
animation := quartzcore.NewBasicAnimation()

// Configure
animation.SetKeyPath("opacity")
animation.SetFromValue(1.0)
animation.SetToValue(0.0)
animation.SetDuration(1.0)

// Add to layer
layer.AddAnimationForKey(animation, "fadeOut")
```

### CAKeyframeAnimation

Animate through multiple keyframe values:

```go
// Create keyframe animation
animation := quartzcore.NewKeyframeAnimation()

// Set path or values
animation.SetKeyPath("position")
animation.SetValues([]interface{}{
    point1, point2, point3, point4,
})
animation.SetDuration(2.0)

// Add to layer
layer.AddAnimationForKey(animation, "path")
```

### CAAnimationGroup

Combine multiple animations:

```go
// Create group
group := quartzcore.NewAnimationGroup()

// Add animations
fadeAnim := quartzcore.NewBasicAnimation()
scaleAnim := quartzcore.NewBasicAnimation()

group.SetAnimations([]Animation{fadeAnim, scaleAnim})
group.SetDuration(1.0)

// Animate
layer.AddAnimationForKey(group, "combo")
```

### CATransition

Transition effects between layer states:

```go
// Create transition
transition := quartzcore.NewTransition()

// Configure effect
transition.SetType("push")
transition.SetSubtype("fromRight")
transition.SetDuration(0.5)

// Apply
layer.AddAnimationForKey(transition, "transition")

// Update layer content
layer.SetContents(newImage)
```

### CASpringAnimation

Physics-based spring animations:

```go
// Create spring animation
spring := quartzcore.NewSpringAnimation()

// Configure spring physics
spring.SetKeyPath("position.y")
spring.SetDamping(10.0)
spring.SetMass(1.0)
spring.SetStiffness(100.0)
spring.SetInitialVelocity(0.0)

// Animate
layer.AddAnimationForKey(spring, "bounce")
```

## Animatable Properties

### Position and Size

```go
// Position
animation.SetKeyPath("position")
animation.SetKeyPath("position.x")
animation.SetKeyPath("position.y")

// Size and bounds
animation.SetKeyPath("bounds")
animation.SetKeyPath("bounds.size")
animation.SetKeyPath("frame") // Not directly animatable, use bounds + position
```

### Visual Properties

```go
// Opacity
animation.SetKeyPath("opacity")

// Background color
animation.SetKeyPath("backgroundColor")

// Corner radius
animation.SetKeyPath("cornerRadius")

// Border
animation.SetKeyPath("borderWidth")
animation.SetKeyPath("borderColor")
```

### Transform Properties

```go
// 2D transforms
animation.SetKeyPath("transform.scale")
animation.SetKeyPath("transform.scale.x")
animation.SetKeyPath("transform.scale.y")
animation.SetKeyPath("transform.rotation")
animation.SetKeyPath("transform.translation")

// 3D transforms
animation.SetKeyPath("transform.rotation.x")
animation.SetKeyPath("transform.rotation.y")
animation.SetKeyPath("transform.rotation.z")
```

### Shadow Properties

```go
animation.SetKeyPath("shadowOpacity")
animation.SetKeyPath("shadowRadius")
animation.SetKeyPath("shadowOffset")
animation.SetKeyPath("shadowColor")
```

## Timing Functions

Control animation pacing:

```go
// Create timing function
timingFunc := quartzcore.NewMediaTimingFunction()

// Built-in timing functions
timingFunc.NewMediaTimingFunctionWithName("linear")
timingFunc.NewMediaTimingFunctionWithName("easeIn")
timingFunc.NewMediaTimingFunctionWithName("easeOut")
timingFunc.NewMediaTimingFunctionWithName("easeInEaseOut")
timingFunc.NewMediaTimingFunctionWithName("default")

// Custom cubic bezier
timingFunc.NewMediaTimingFunctionWithControlPoints(0.5, 0.0, 0.5, 1.0)

// Apply to animation
animation.SetTimingFunction(timingFunc)
```

## Specialized Layer Types

### CAShapeLayer

Vector graphics with CGPath:

```go
// Create shape layer
shapeLayer := quartzcore.NewShapeLayer()

// Set path
path := coregraphics.NewPath()
// ... configure path ...
shapeLayer.SetPath(path)

// Styling
shapeLayer.SetFillColor(color)
shapeLayer.SetStrokeColor(color)
shapeLayer.SetLineWidth(2.0)

// Animate path
animation := quartzcore.NewBasicAnimation()
animation.SetKeyPath("strokeEnd")
animation.SetFromValue(0.0)
animation.SetToValue(1.0)
shapeLayer.AddAnimationForKey(animation, "draw")
```

### CATextLayer

Rendered text content:

```go
// Create text layer
textLayer := quartzcore.NewTextLayer()

// Configure text
textLayer.SetString("Hello, World!")
textLayer.SetFont(font)
textLayer.SetFontSize(24.0)
textLayer.SetForegroundColor(color)
textLayer.SetAlignmentMode("center")

// Enable anti-aliasing
textLayer.SetIsWrapped(true)
```

### CAGradientLayer

Color gradient fills:

```go
// Create gradient layer
gradientLayer := quartzcore.NewGradientLayer()

// Set colors
gradientLayer.SetColors([]interface{}{
    color1, color2, color3,
})

// Set gradient direction
gradientLayer.SetStartPoint(CGPoint{0, 0})  // Top
gradientLayer.SetEndPoint(CGPoint{0, 1})    // Bottom

// Set color stops
gradientLayer.SetLocations([]float64{0.0, 0.5, 1.0})
```

### CAEmitterLayer

Particle effects system:

```go
// Create emitter layer
emitterLayer := quartzcore.NewEmitterLayer()

// Configure emitter position
emitterLayer.SetEmitterPosition(point)
emitterLayer.SetEmitterSize(size)
emitterLayer.SetEmitterShape("circle")

// Create emitter cell
cell := quartzcore.NewEmitterCell()
cell.SetContents(image)
cell.SetBirthRate(10.0)
cell.SetLifetime(5.0)
cell.SetVelocity(100.0)
cell.SetScale(0.1)
cell.SetEmissionRange(2 * math.Pi)

// Add cells
emitterLayer.SetEmitterCells([]EmitterCell{cell})
```

### CAReplicatorLayer

Duplicate sublayers with transforms:

```go
// Create replicator layer
replicatorLayer := quartzcore.NewReplicatorLayer()

// Configure replication
replicatorLayer.SetInstanceCount(10)
replicatorLayer.SetInstanceDelay(0.1)

// Set transform for each instance
transform := CATransform3DMakeTranslation(20, 0, 0)
replicatorLayer.SetInstanceTransform(transform)

// Set color/opacity offset
replicatorLayer.SetInstanceRedOffset(-0.1)
replicatorLayer.SetInstanceAlphaOffset(-0.1)

// Add base layer
replicatorLayer.AddSublayer(baseLayer)
```

## Animation Workflow

### Implicit Animations

Changes to animatable properties automatically animate:

```go
// Disable implicit animations
quartzcore.CATransaction.Begin()
quartzcore.CATransaction.SetDisableActions(true)

layer.SetOpacity(0.5) // No animation

quartzcore.CATransaction.Commit()

// With implicit animation (default)
layer.SetOpacity(0.5) // Animates automatically
```

### Explicit Animations

Full control over animation:

```go
// Create animation
animation := quartzcore.NewBasicAnimation()
animation.SetKeyPath("position")
animation.SetFromValue(startPoint)
animation.SetToValue(endPoint)
animation.SetDuration(1.0)
animation.SetTimingFunction(timingFunc)

// Add to layer
layer.AddAnimationForKey(animation, "move")

// Note: Must also set final value
layer.SetPosition(endPoint)
```

### Animation Callbacks

Handle animation completion:

```go
// Set delegate (requires implementing CAAnimationDelegate protocol)
animation.SetDelegate(delegate)

// Delegate methods:
// - animationDidStart(anim)
// - animationDidStop(anim, finished)
```

## Transactions

Batch property changes:

```go
// Begin transaction
quartzcore.CATransaction.Begin()

// Set transaction properties
quartzcore.CATransaction.SetAnimationDuration(2.0)
quartzcore.CATransaction.SetAnimationTimingFunction(timingFunc)

// Make changes
layer1.SetOpacity(0.5)
layer2.SetPosition(newPosition)
layer3.SetTransform(transform)

// Commit transaction
quartzcore.CATransaction.Commit()
```

## 3D Transforms

Perspective and depth:

```go
// Create 3D transform
var transform CATransform3D

// Rotation
transform = CATransform3DMakeRotation(angle, x, y, z)

// Scale
transform = CATransform3DMakeScale(sx, sy, sz)

// Translation
transform = CATransform3DMakeTranslation(tx, ty, tz)

// Combine transforms
transform = CATransform3DRotate(transform, angle, 0, 1, 0)
transform = CATransform3DScale(transform, 2.0, 2.0, 1.0)

// Apply to layer
layer.SetTransform(transform)

// Set perspective on parent
var perspective CATransform3D
perspective.m34 = -1.0 / 500.0 // Perspective strength
parentLayer.SetSublayerTransform(perspective)
```

## Performance Optimization

### Rendering Optimization

```go
// Rasterize complex layer hierarchies
layer.SetShouldRasterize(true)
layer.SetRasterizationScale(scale) // Match screen scale

// Draw asynchronously
layer.SetDrawsAsynchronously(true)

// Opaque layers render faster
layer.SetOpaque(true)
```

### Animation Optimization

```go
// Use smaller layer trees
// Prefer transform over frame changes
// Use CAShapeLayer instead of redrawing
// Cache static content
// Disable off-screen rendering when possible
```

## Use Cases

### UI Transitions

```go
// Fade transition
transition := quartzcore.NewTransition()
transition.SetType("fade")
transition.SetDuration(0.3)
view.Layer().AddAnimationForKey(transition, "fade")

// Update content
view.Layer().SetContents(newImage)
```

### Custom Drawing

```go
// Override drawInContext in custom layer
// Use Core Graphics to draw
// Layer caches result automatically
```

### Video Playback

```go
// AVPlayerLayer is a CALayer subclass
playerLayer := avfoundation.NewPlayerLayer()
playerLayer.SetPlayer(player)
view.Layer().AddSublayer(playerLayer)
```

### Particle Effects

```go
// Create snow effect
emitterLayer := quartzcore.NewEmitterLayer()
// Configure with white particles
// Set velocity, lifetime, birth rate
view.Layer().AddSublayer(emitterLayer)
```

## Best Practices

1. **Use Implicit Animations**: Let Core Animation handle simple property changes
2. **Transform over Frame**: Animating transform is faster than frame
3. **Opaque Layers**: Set opaque=true when possible for better performance
4. **Layer Backing**: Use layer-backed views for better performance
5. **Reduce Overdraw**: Minimize transparent overlapping layers
6. **Rasterize Complex Hierarchies**: Cache static complex layer trees
7. **Match Screen Scale**: Set rasterizationScale to screen scale
8. **Disable Actions When Needed**: Use transactions to disable implicit animations

## Common Patterns

### Fade In/Out

```go
// Fade out
animation := quartzcore.NewBasicAnimation()
animation.SetKeyPath("opacity")
animation.SetToValue(0.0)
animation.SetDuration(0.3)
layer.AddAnimationForKey(animation, "fadeOut")
layer.SetOpacity(0.0)
```

### Pulse Effect

```go
// Scale pulse
animation := quartzcore.NewBasicAnimation()
animation.SetKeyPath("transform.scale")
animation.SetFromValue(1.0)
animation.SetToValue(1.2)
animation.SetDuration(0.3)
animation.SetAutoreverses(true)
animation.SetRepeatCount(math.Inf(1)) // Infinite
layer.AddAnimationForKey(animation, "pulse")
```

### Rotation

```go
// Continuous rotation
animation := quartzcore.NewBasicAnimation()
animation.SetKeyPath("transform.rotation")
animation.SetFromValue(0.0)
animation.SetToValue(2 * math.Pi)
animation.SetDuration(1.0)
animation.SetRepeatCount(math.Inf(1))
layer.AddAnimationForKey(animation, "spin")
```

## Integration with Other Frameworks

### UIKit/AppKit

Every view has a layer:

```go
// Access view's layer
layer := view.Layer()

// Animate view
quartzcore.CATransaction.Begin()
view.Frame = newFrame
quartzcore.CATransaction.Commit()
```

### Metal

Direct GPU rendering:

```go
// Create Metal layer
metalLayer := quartzcore.NewMetalLayer()
metalLayer.SetDevice(device)
metalLayer.SetPixelFormat(pixelFormat)

// Add to view
view.Layer().AddSublayer(metalLayer)

// Render directly with Metal
```

### Core Image

Apply filters to layers:

```go
// Set Core Image filter
layer.SetFilters([]interface{}{filter})

// Background filters
layer.SetBackgroundFilters([]interface{}{blurFilter})
```

## References

- [Core Animation Programming Guide](https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/CoreAnimation_guide/)
- [CALayer Documentation](https://developer.apple.com/documentation/quartzcore/calayer)
- [CAAnimation Documentation](https://developer.apple.com/documentation/quartzcore/caanimation)
- [Core Animation Basics](https://developer.apple.com/documentation/quartzcore)
