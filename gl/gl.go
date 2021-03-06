/*
A straight GL 3.2 binding.

Not idiomatic Go in any real stretch, wanted to keep it close to the original API.

*/
package gl

// #cgo linux   LDFLAGS: -lGLEW -lGL
// #cgo windows LDFLAGS: -lglew32 -lopengl32
// #cgo darwin LDFLAGS: -framework OpenGL -lGLEW
// #include "common.h"
import "C"

import (
	"unsafe"
)

type GLsync C.GLsync

func Init() {
	C.glewExperimental = 1
	C.glewInit()
}

/*
Select active texture unit

Parameters
    texture - Specifies which texture unit to make active. The number of texture units is implementation dependent, but must be at least 48. texture must be one of Texturei, where i ranges from 0 (MaxCombinedTextureImageUnits - 1). The initial value is Texture0.
*/
func ActiveTexture(texture GLenum) {
	C.glActiveTexture(C.GLenum(texture))
}

/*
Attaches a shader object to a program object

Parameters
    program - Specifies the program object to which a shader object will be attached.
    shader - Specifies the shader object that is to be attached.
*/
func AttachShader(program, shader uint32) {
	C.glAttachShader(C.GLuint(program), C.GLuint(shader))
}

/*
Start conditional rendering

Parameters
    id - Specifies the name of an occlusion query object whose results are used to determine if the rendering commands are discarded.
    mode - Specifies how BeginConditionalRender interprets the results of the occlusion query.
*/
func BeginConditionalRender(id uint32, mode GLenum) {
	C.glBeginConditionalRender(C.GLuint(id), C.GLenum(mode))
}

/*
Delimit the boundaries of a query object

Parameters
    target - Specifies the target type of query object established between BeginQuery and the subsequent EndQuery. The symbolic constant must be one of SamplesPassed, AnySamplesPassed, PrimitivesGenerated, TransformFeedbackPrimitivesWritten, or TimeElapsed.
    id - Specifies the name of a query object.
*/
func BeginQuery(target GLenum, id uint32) {
	C.glBeginQuery(C.GLenum(target), C.GLuint(id))
}

/*
Start transform feedback operation

Parameters
    primitiveMode - Specify the output type of the primitives that will be recorded into the buffer objects that are bound for transform feedback.
*/
func BeginTransformFeedback(primitiveMode GLenum) {
	C.glBeginTransformFeedback(C.GLenum(primitiveMode))
}

/*
Associates a generic vertex attribute index with a named attribute variable

Parameters
    program - Specifies the handle of the program object in which the association is to be made.
    index - Specifies the index of the generic vertex attribute to be bound.
    name - Specifies a string containing the name of the vertex shader attribute variable to which index is to be bound.
*/
func BindAttribLocation(program, index uint32, name string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	C.glBindAttribLocation(C.GLuint(program), C.GLuint(index), (*C.GLchar)(cName))
}

/*
Bind a named buffer object

Parameters
    target - Specifies the target to which the buffer object is bound. The symbolic constant must be ArrayBuffer, CopyReadBuffer, CopyWriteBuffer, ElementArrayBuffer, PixelPackBuffer, PixelUnpackBuffer, TextureBuffer, TransformFeedbackBuffer,  or UniformBuffer.
    buffer - Specifies the name of a buffer object.
*/
func BindBuffer(target GLenum, buffer uint32) {
	C.glBindBuffer(C.GLenum(target), C.GLuint(buffer))
}

/*
Bind a buffer object to an indexed buffer target

Parameters
    target - Specify the target of the bind operation. target must be either TransformFeedbackBuffer or UniformBuffer.
    index - Specify the index of the binding point within the array specified by target.
    buffer - The name of a buffer object to bind to the specified binding point.
*/
func BindBufferBase(target GLenum, index, buffer uint32) {
	C.glBindBufferBase(C.GLenum(target), C.GLuint(index), C.GLuint(buffer))
}

/*
Bind a range within a buffer object to an indexed buffer target

Parameters
    target - Specify the target of the bind operation. target must be either TransformFeedbackBuffer or UniformBuffer.
    index - Specify the index of the binding point within the array specified by target.
    buffer - The name of a buffer object to bind to the specified binding point.
    offset - The starting offset in basic machine units into the buffer object buffer.
    size - The amount of data in machine units that can be read from the buffet object while used as an indexed target.
*/
func BindBufferRange(target GLenum, index, buffer, offset, size uint32) {
	C.glBindBufferRange(C.GLenum(target), C.GLuint(index), C.GLuint(buffer), C.GLintptr(offset), C.GLsizeiptr(size))
}

/*
Bind a framebuffer to a framebuffer target

Parameters
    target - Specifies the framebuffer target of the binding operation.
    framebuffer - Specifies the name of the framebuffer object to bind.
*/
func BindFramebuffer(target GLenum, framebuffer uint32) {
	C.glBindFramebuffer(C.GLenum(target), C.GLuint(framebuffer))
}

/*
Bind a renderbuffer to a renderbuffer target

Parameters
    target - Specifies the renderbuffer target of the binding operation. target must be Renderbuffer.
    renderbuffer - Specifies the name of the renderbuffer object to bind.
*/
func BindRenderbuffer(target GLenum, renderbuffer uint32) {
	C.glBindRenderbuffer(C.GLenum(target), C.GLuint(renderbuffer))
}

/*
Bind a named texture to a texturing target

Parameters
    target - Specifies the target to which the texture is bound. Must be either Texture1d, Texture2d, Texture3d, or Texture1dArray, Texture2dArray, TextureRectangle, TextureCubeMap, TextureBuffer, Texture2dMultisample or Texture2dMultisampleArray.
    texture - Specifies the name of a texture.
*/
func BindTexture(target GLenum, texture uint32) {
	C.glBindTexture(C.GLenum(target), C.GLuint(texture))
}

/*
Bind a vertex array object

Parameters
    array - Specifies the name of the vertex array to bind.
*/
func BindVertexArray(array uint32) {
	C.glBindVertexArray(C.GLuint(array))
}

/*
Set the blend color

Parameters
    red, green, blue, alpha - Specify the components of BLEND_COLOR.
*/
func BlendColor(red, green, blue, alpha float32) {
	C.glBlendColor(C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha))
}

/*
Specify the equation used for both the RGB blend equation and the Alpha blend equation

Parameters
    mode - Specifies how source and destination colors are combined. It must be FuncAdd, FuncSubtract, FuncReverseSubtract, Min, Max.
*/
func BlendEquation(mode GLenum) {
	C.glBlendEquation(C.GLenum(mode))
}

/*
Set the RGB blend equation and the alpha blend equation separately

Parameters
    modeRGB - Specifies the RGB blend equation, how the red, green, and blue components of the source and destination colors are combined. It must be FuncAdd, FuncSubtract, FuncReverseSubtract, Min, Max.
    modeAlpha - Specifies the alpha blend equation, how the alpha component of the source and destination colors are combined. It must be FuncAdd, FuncSubtract, FuncReverseSubtract, Min, Max.
*/
func BlendEquationSeparate(modeRGB, modeAlpha GLenum) {
	C.glBlendEquationSeparate(C.GLenum(modeRGB), C.GLenum(modeAlpha))
}

/*
Specify pixel arithmetic.

Parameters
    sfactor - Specifies how the red, green, blue, and alpha source blending factors are computed. The initial value is One.
    dfactor - Specifies how the red, green, blue, and alpha destination blending factors are computed. The following symbolic constants are accepted: Zero, One, SrcColor, OneMinusSrcColor, DstColor, OneMinusDstColor, SrcAlpha, OneMinusSrcAlpha, DstAlpha, OneMinusDstAlphaConstantColor, OneMinusConstantColor, ConstantAlpha, and OneMinusConstantAlpha. The initial value is Zero.
*/
func BlendFunc(sfactor, dfactor GLenum) {
	C.glBlendFunc(C.GLenum(sfactor), C.GLenum(dfactor))
}

/*
Specify pixel arithmetic for RGB and alpha components separately

Parameters
    srcRGB - Specifies how the red, green, and blue blending factors are computed. The initial value is One.
    dstRGB - Specifies how the red, green, and blue destination blending factors are computed. The initial value is Zero.
    srcAlpha - Specified how the alpha source blending factor is computed. The initial value is One.
    dstAlpha - Specified how the alpha destination blending factor is computed. The initial value is Zero.
*/
func BlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha GLenum) {
	C.glBlendFuncSeparate(C.GLenum(srcRGB), C.GLenum(dstRGB), C.GLenum(srcAlpha), C.GLenum(dstAlpha))
}

/*
Copy a block of pixels from the read framebuffer to the draw framebuffer

Parameters
    srcX0, srcY0, srcX1, srcY1 - Specify the bounds of the source rectangle within the read buffer of the read framebuffer.
    dstX0, dstY0, dstX1, dstY1 - Specify the bounds of the destination rectangle within the write buffer of the write framebuffer.
    mask - The bitwise OR of the flags indicating which buffers are to be copied. The allowed flags are ColorBufferBit, DepthBufferBit and StencilBufferBit.
    filter - Specifies the interpolation to be applied if the image is stretched. Must be Nearest or Linear.
*/
func BlitFramebuffer(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1 int, mask GLbitfield, filter GLenum) {
	C.glBlitFramebuffer(C.GLint(srcX0), C.GLint(srcY0), C.GLint(srcX1), C.GLint(srcY1), C.GLint(dstX0), C.GLint(dstY0), C.GLint(dstX1), C.GLint(dstY1), C.GLbitfield(mask), C.GLenum(filter))
}

/*
Creates and initializes a buffer object's data store

Parameters
    target - Specifies the target buffer object. The symbolic constant must be ArrayBuffer, CopyReadBuffer, CopyWriteBuffer, ElementArrayBuffer, PixelPackBuffer, PixelUnpackBuffer, TextureBuffer, TransformFeedbackBuffer, or UniformBuffer.
    data - Specifies a slice of the data that will be copied into the data store for initialization, or nil if not data is to be copied.
    usage - Specifies the expected usage pattern of the data store. The symbolic constant must be StreamDraw, StreamRead, StreamCopy, StaticDraw, StaticRead, StaticCopy, DynamicDraw, DynamicRead, or DynamicCopy.
*/
func BufferData(target GLenum, data interface{}, usage GLenum) error {
	if data == nil {
		C.glBufferData(C.GLenum(target), 0, nil, C.GLenum(usage))
	} else {
		_, size, _, ptr, err := sliceToGLData(data)
		if err != nil {
			return err
		}

		C.glBufferData(C.GLenum(target), C.GLsizeiptr(size), ptr, C.GLenum(usage))
	}

	return nil
}

/*
Updates a subset of a buffer object's data store

Parameters
    target - Specifies the target buffer object. The symbolic constant must be ArrayBuffer, CopyReadBuffer, CopyWriteBuffer, ElementArrayBuffer, PixelPackBuffer, PixelUnpackBuffer, TextureBuffer, TransformFeedbackBuffer, or UniformBuffer.
    offset - Specifies the offset into the buffer object's data store where data replacement will begin, measured in bytes.
    data - Specifies a slice to the new data that will be copied into the data store.
*/
func BufferSubData(target GLenum, offset int32, data interface{}) error {
	_, size, _, ptr, err := sliceToGLData(data)
	if err != nil {
		return err
	}

	C.glBufferSubData(C.GLenum(target), C.GLintptr(offset), C.GLsizeiptr(size), ptr)

	return nil
}

/*
Check the completeness status of a framebuffer

Parameters
    target - Specify the target of the framebuffer completeness check.
*/
func CheckFramebufferStatus(target GLenum) GLenum {
	return GLenum(C.glCheckFramebufferStatus(C.GLenum(target)))
}

/*
Specify whether data read via ReadPixels should be clamped

Parameters
    target - arget for color clamping. target must be ClampReadColor.
    clamp - Specifies whether to apply color clamping. clamp must be True or False.
*/
func ClampColor(target, clamp GLenum) {
	C.glClampColor(C.GLenum(target), C.GLenum(clamp))
}

/*
Clear buffers to preset values.

Parameters
    mask - Bitwise OR of masks that indicate the buffers to be cleared. The four masks are ColorBufferBit, DepthBufferBit, AccumBufferBit, and StencilBufferBit.
*/
func Clear(mask GLbitfield) {
	C.glClear(C.GLbitfield(mask))
}

/*
Specify clear values for the color buffers.

Parameters
    red, green, blue, alpha - Specify the red, green, blue, and alpha values used when the color buffers are cleared. The initial values are all 0.
*/
func ClearColor(red, green, blue, alpha float32) {
	C.glClearColor(C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha))
}

/*
Specify the clear value for the depth buffer.

Parameters
    depth - Specifies the depth value used when the depth buffer is cleared. The initial value is 1.
*/
func ClearDepth(depth float32) {
	C.glClearDepth(C.GLclampd(depth))
}

/*
Specify the clear value for the stencil buffer.

Parameters
    s - Specifies the index used when the stencil buffer is cleared. The initial value is 0.
*/
func ClearStencil(s int) {
	C.glClearStencil(C.GLint(s))
}

/*
Block and wait for a sync object to become signaled

Parameters
    sync - The sync object whose status to wait on.
    flags - A bitfield controlling the command flushing behavior. flags may be SyncFlushCommandsBit.
    timeout - The timeout, specified in nanoseconds, for which the implementation should wait for sync to become signaled.
*/
func ClientWaitSync(sync GLsync, flags GLbitfield, timeout uint64) GLenum {
	return GLenum(C.glClientWaitSync(C.GLsync(sync), C.GLbitfield(flags), C.GLuint64(timeout)))
}

/*
Enable and disable writing of frame buffer color components.

Parameters
    red, green, blue, alpha - Specify whether red, green, blue, and alpha can or cannot be written into the frame buffer. The initial values are all true, indicating that the color components can be written.
*/
func ColorMask(red, green, blue, alpha bool) {
	C.glColorMask(boolToGLBool(red), boolToGLBool(green), boolToGLBool(blue), boolToGLBool(alpha))
}

/*
Compiles a shader object

Parameters
    shader - Specifies the shader object to be compiled.
*/
func CompileShader(shader uint32) {
	C.glCompileShader(C.GLuint(shader))
}

/*
Specify a one-dimensional texture image in a compressed format

Parameters
    target - Specifies the target texture. Must be Texture1d or ProxyTexture1d.
    level - Specifies the level-of-detail number. Level 0 is the base image level. Level n is the nth mipmap reduction image.
    internalformat - Specifies the format of the compressed image data stored at address data.
    width - Specifies the width of the texture image. All implementations support texture images that are at least 64 texels wide. The height of the 1D texture image is 1.
    border - This value must be 0.
    data - Specifies a pointer to the compressed image data in memory.
*/
func CompressedTexImage1D(target GLenum, level int, internalFormat GLenum, width, border int, data interface{}) error {
	_, size, _, ptr, err := sliceToGLData(data)
	if err != nil {
		return err
	}

	C.glCompressedTexImage1D(C.GLenum(target), C.GLint(level), C.GLenum(internalFormat), C.GLsizei(width), C.GLint(border), C.GLsizei(size), ptr)

	return nil
}

/*
Specify a two-dimensional texture image in a compressed format

Parameters
    target - Specifies the target texture. Must be Texture2d, ProxyTexture2d, Texture1dArray, ProxyTexture1dArray, TextureCubeMapPositiveX, TextureCubeMapNegativeX, TextureCubeMapPositiveY, TextureCubeMapNegativeY, TextureCubeMapPositiveZ, TextureCubeMapNegativeZ, or ProxyTextureCubeMap.
    level - Specifies the level-of-detail number. Level 0 is the base image level. Level n is the nth mipmap reduction image.
    internalformat - Specifies the format of the compressed image data stored at address data.
    width - Specifies the width of the texture image. All implementations support 2D texture images that are at least 64 texels wide and cube-mapped texture images that are at least 16 texels wide.
    height - Specifies the height of the texture image. All implementations support 2D texture images that are at least 64 texels high and cube-mapped texture images that are at least 16 texels high.
    border - This value must be 0.
    data - Specifies a pointer to the compressed image data in memory.
*/
func CompressedTexImage2D(target GLenum, level int, internalFormat GLenum, width, height, border int, data interface{}) error {
	_, size, _, ptr, err := sliceToGLData(data)
	if err != nil {
		return err
	}

	C.glCompressedTexImage2D(C.GLenum(target), C.GLint(level), C.GLenum(internalFormat), C.GLsizei(width), C.GLsizei(height), C.GLint(border), C.GLsizei(size), ptr)

	return nil
}

/*
Specify a three-dimensional texture image in a compressed format

Parameters
    target - Specifies the target texture. Must be Texture3d, ProxyTexture3d, Texture2dArray or ProxyTexture2dArray.
    level - Specifies the level-of-detail number. Level 0 is the base image level. Level n is the nth mipmap reduction image.
    internalformat - Specifies the format of the compressed image data stored at address data.
    width - Specifies the width of the texture image. All implementations support 3D texture images that are at least 16 texels wide.
    height - Specifies the height of the texture image. All implementations support 3D texture images that are at least 16 texels high.
    depth - Specifies the depth of the texture image. All implementations support 3D texture images that are at least 16 texels deep.
    border - This value must be 0.
    data - Specifies a pointer to the compressed image data in memory.
*/
func CompressedTexImage3D(target GLenum, level int, internalFormat GLenum, width, height, depth, border int, data interface{}) error {
	_, size, _, ptr, err := sliceToGLData(data)
	if err != nil {
		return err
	}

	C.glCompressedTexImage3D(C.GLenum(target), C.GLint(level), C.GLenum(internalFormat), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLint(border), C.GLsizei(size), ptr)

	return nil
}

/*
Specify a one-dimensional texture subimage in a compressed format

Parameters
    target - Specifies the target texture. Must be Texture1d.
    level - Specifies the level-of-detail number. Level 0 is the base image level. Level n is the nth mipmap reduction image.
    xoffset - Specifies a texel offset in the x direction within the texture array.
    width - Specifies the width of the texture image. All implementations support texture images that are at least 64 texels wide. The height of the 1D texture image is 1.
    format - Specifies the format of the compressed image data stored at address data.
    border - This value must be 0.
    data - Specifies a pointer to the compressed image data in memory.
*/
func CompressedTexSubImage1D(target GLenum, level, xoffset, width int, format GLenum, data interface{}) error {
	_, size, _, ptr, err := sliceToGLData(data)
	if err != nil {
		return err
	}

	C.glCompressedTexSubImage1D(C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLsizei(width), C.GLenum(format), C.GLsizei(size), ptr)

	return nil
}

/*
Specify a two-dimensional texture subimage in a compressed format

Parameters
    target - Specifies the target texture. Must be Texture2d, TextureCubeMapPositiveX, TextureCubeMapNegativeX, TextureCubeMapPositiveY, TextureCubeMapNegativeY, TextureCubeMapPositiveZ, or TextureCubeMapNegativeZ.
    level - Specifies the level-of-detail number. Level 0 is the base image level. Level n is the nth mipmap reduction image.
    xoffset - Specifies a texel offset in the x direction within the texture array.
    yoffset - Specifies a texel offset in the y direction within the texture array.
    width - Specifies the width of the texture subimage.
    height - Specifies the height of the texture subimage.
    format - Specifies the format of the compressed image data stored at address data.
    data - Specifies a pointer to the compressed image data in memory.
*/
func CompressedTexSubImage2D(target GLenum, level, xoffset, yoffset, width, height int, format GLenum, data interface{}) error {
	_, size, _, ptr, err := sliceToGLData(data)
	if err != nil {
		return err
	}

	C.glCompressedTexSubImage2D(C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLsizei(size), ptr)

	return nil
}

/*
Specify a three-dimensional texture subimage in a compressed format

Parameters
    target - Specifies the target texture. Must be Texture3d.
    level - Specifies the level-of-detail number. Level 0 is the base image level. Level n is the nth mipmap reduction image.
    xoffset - Specifies a texel offset in the x direction within the texture array.
    yoffset - Specifies a texel offset in the y direction within the texture array.
    zoffset - Specifies a texel offset in the z direction within the texture array.
    width - Specifies the width of the texture subimage.
    height - Specifies the height of the texture subimage.
    depth - Specifies the depth of the texture subimage.
    format - Specifies the format of the compressed image data stored at address data.
    data - Specifies a pointer to the compressed image data in memory.
*/
func CompressedTexSubImage3D(target GLenum, level, xoffset, yoffset, zoffset, width, height, depth int, format GLenum, data interface{}) error {
	_, size, _, ptr, err := sliceToGLData(data)
	if err != nil {
		return err
	}

	C.glCompressedTexSubImage3D(C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(zoffset), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLenum(format), C.GLsizei(size), ptr)

	return nil
}

/*
Copy part of the data store of a buffer object to the data store of another buffer object

Parameters
    readTarget - Specifies the target from whose data store data should be read.
    writeTarget - Specifies the target to whose data store data should be written.
    readOffset - Specifies the offset, in basic machine units, within the data store of readTarget from which data should be read.
    writeOffset - Specifies the offset, in basic machine units, within the data store of writeTarget to which data should be written.
    size - Specifies the size, in basic machine units, of the data to be copied from readTarget to writeTarget.
*/
func CopyBufferSubData(readTarget, writeTarget GLenum, readOffset, writeOffset, size int) {
	C.glCopyBufferSubData(C.GLenum(readTarget), C.GLenum(writeTarget), C.GLintptr(readOffset), C.GLintptr(writeOffset), C.GLsizeiptr(size))
}

/*
Copy pixels into a 1D texture image.

Parameters
    target - Specifies the target texture. Must be Texture1d.
    level - Specifies the level-of-detail number. Level 0 is the base image level. Level n is the nth mipmap reduction image.
    internalFormat - Specifies the internal format of the texture. Must be one of the following symbolic constants: CompressedRed, CompressedRg, CompressedRgb, CompressedRgba. CompressedSrgb, CompressedSrgbAlpha. DepthComponent, DepthComponent16, DepthComponent24, DepthComponent32, Red, Rg, Rgb, R3G3B2, Rgb4, Rgb5, Rgb8, Rgb10, Rgb12, Rgb16, Rgba, Rgba2, Rgba4, Rgb5A1, Rgba8, Rgb10A2, Rgba12, Rgba16, Srgb, Srgb8, SrgbAlpha, or Srgb8Alpha8.
    x, y - Specify the window coordinates of the left corner of the row of pixels to be copied.
    width - Specifies the width of the texture image. The height of the texture image is 1.
    border - Must be 0.
*/
func CopyTexImage1D(target GLenum, level int, internalFormat GLenum, x, y, width, border int) {
	C.glCopyTexImage1D(C.GLenum(target), C.GLint(level), C.GLenum(internalFormat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLint(border))
}

/*
Copy pixels into a 2D texture image.

Parameters
    target - Specifies the target texture. Must be Texture2d, TextureCubeMapPositiveX, TextureCubeMapNegativeX, TextureCubeMapPositiveY, TextureCubeMapNegativeY, TextureCubeMapPositiveZ, or TextureCubeMapNegativeZ.
    level - Specifies the level-of-detail number. Level 0 is the base image level. Level n is the nth mipmap reduction image.
    internalFormat - Specifies the internal format of the texture. Must be one of the following symbolic constants: CompressedRed, CompressedRg, CompressedRgb, CompressedRgba. CompressedSrgb, CompressedSrgbAlpha. DepthComponent, DepthComponent16, DepthComponent24, DepthComponent32, Red, Rg, Rgb, R3G3B2, Rgb4, Rgb5, Rgb8, Rgb10, Rgb12, Rgb16, Rgba, Rgba2, Rgba4, Rgb5A1, Rgba8, Rgb10A2, Rgba12, Rgba16, Srgb, Srgb8, SrgbAlpha, or Srgb8Alpha8.
    x, y - Specify the window coordinates of the left corner of the row of pixels to be copied.
    width - Specifies the width of the texture image.
    height - Specifies the width of the texture image.
    border - Must be 0.
*/
func CopyTexImage2D(target GLenum, level int, internalFormat GLenum, x, y, width, height, border int) {
	C.glCopyTexImage2D(C.GLenum(target), C.GLint(level), C.GLenum(internalFormat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLint(border))
}

/*
Copy a one-dimensional texture subimage

Parameters
    target - Must be Texture1d.
    level - Specifies the level-of-detail number. Level 0 is the base image level. Level n is the nth mipmap reduction image.
    xoffset - Specifies the texel offset within the texture array.
    x, y - Specify the window coordinates of the left corner of the row of pixels to be copied.
    width - Specifies the width of the texture subimage
*/
func CopyTexSubImage1D(target GLenum, level, xoffset, x, y, width int) {
	C.glCopyTexSubImage1D(C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(x), C.GLint(y), C.GLsizei(width))
}

/*
Copy a two-dimensional texture subimage

Parameters
    target - Specifies the target texture. Must be Texture2d, TextureCubeMapPositiveX, TextureCubeMapNegativeX, TextureCubeMapPositiveY, TextureCubeMapNegativeY, TextureCubeMapPositiveZ, or TextureCubeMapNegativeZ.
    level - Specifies the level-of-detail number. Level 0 is the base image level. Level n is the nth mipmap reduction image.
    xoffset - Specifies a texel offset in the x direction within the texture array.
    yoffset - Specifies a texel offset in the y direction within the texture array.
    x, y - Specify the window coordinates of the left corner of the row of pixels to be copied.
    width - Specifies the width of the texture subimage
    height - Specifies the height of the texture subimage
*/
func CopyTexSubImage2D(target GLenum, level, xoffset, yoffset, x, y, width, height int) {
	C.glCopyTexSubImage2D(C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}

/*
Copy a three-dimensional texture subimage

Parameters
    target - Specifies the target texture. Must be Texture3d or Texture2dArray.
    level - Specifies the level-of-detail number. Level 0 is the base image level. Level n is the nth mipmap reduction image.
    xoffset - Specifies a texel offset in the x direction within the texture array.
    yoffset - Specifies a texel offset in the y direction within the texture array.
    zoffset - Specifies a texel offset in the z direction within the texture array.
    x, y - Specify the window coordinates of the left corner of the row of pixels to be copied.
    width - Specifies the width of the texture subimage
    height - Specifies the height of the texture subimage
*/
func CopyTexSubImage3D(target GLenum, level, xoffset, yoffset, zoffset, x, y, width, height int) {
	C.glCopyTexSubImage3D(C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(zoffset), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}

/*
Creates a program object
*/
func CreateProgram() uint32 {
	return uint32(C.glCreateProgram())
}

/*
Creates a shader object

Parameters
    shaderType - Specifies the type of shader to be created. Must be one of VertexShader, GeometryShader or FragmentShader.
*/
func CreateShader(shaderType GLenum) uint32 {
	return uint32(C.glCreateShader(C.GLenum(shaderType)))
}

/*
Specify whether front- or back-facing facets can be culled.

Parameters
    mode - Specifies whether front- or back-facing facets are candidates for culling. Symbolic constants Front, Back, and FrontAndBack are accepted. The initial value is Back.
*/
func CullFace(mode GLenum) {
	C.glCullFace(C.GLenum(mode))
}

/*
Delete named buffer objects

Parameters
    buffers - Specifies one or several buffer objects to be deleted.
*/
func DeleteBuffers(buffers ...uint32) {
	count := len(buffers)
	C.glDeleteBuffers(C.GLsizei(count), (*C.GLuint)(&buffers[0]))
}

/*
Delete framebuffer objects

Parameters
    framebuffers - Specifies one or several framebuffer objects to be deleted.
*/
func DeleteFramebuffers(framebuffers ...uint32) {
	count := len(framebuffers)
	C.glDeleteFramebuffers(C.GLsizei(count), (*C.GLuint)(&framebuffers[0]))
}

/*
Deletes a program object

Parameters
    program - Specifies the program object to be deleted.
*/
func DeleteProgram(program uint32) {
	C.glDeleteProgram(C.GLuint(program))
}

/*
Delete named query objects

Parameters
    ids - Specifies one or several query objects to be deleted.
*/
func DeleteQeueries(ids ...uint32) {
	count := len(ids)
	C.glDeleteQueries(C.GLsizei(count), (*C.GLuint)(&ids[0]))
}

/*
Delete renderbuffer objects

Parameters
    renderbuffers - Specifies one or several renderbuffer objects to be deleted.
*/
func DeleteRenderbuffers(renderbuffers ...uint32) {
	count := len(renderbuffers)
	C.glDeleteRenderbuffers(C.GLsizei(count), (*C.GLuint)(&renderbuffers[0]))
}

/*
Deletes a shader object

Parameters
    shader - Specifies the shader object to be deleted.
*/
func DeleteShader(shader uint32) {
	C.glDeleteShader(C.GLuint(shader))
}

/*
Deletes a sync object

Parameters
    sync - Specifies the sync object to be deleted.
*/
func DeleteSync(sync GLsync) {
	C.glDeleteSync(C.GLsync(sync))
}

/*
Delete named textures

Parameters
    textures - Specifies one or several textures to be deleted.
*/
func DeleteTextures(textures ...uint32) {
	count := len(textures)
	C.glDeleteTextures(C.GLsizei(count), (*C.GLuint)(&textures[0]))
}

/*
Delete vertex array objects

Parameters
    arrays - Specifies one or several vertex arrays to be deleted.
*/
func DeleteVertexArrays(arrays ...uint32) {
	count := len(arrays)
	C.glDeleteVertexArrays(C.GLsizei(count), (*C.GLuint)(&arrays[0]))
}

/*
Specify the value used for depth buffer comparisons

Parameters
    constFunc - Specifies the depth comparison function. Symbolic constants Never, Less, Equal, LEqual, Greater, NotEqual, GEqual, and Always are accepted. The initial value is Less.
*/
func DepthFunc(constFunc GLenum) {
	C.glDepthFunc(C.GLenum(constFunc))
}

/*
Enable or disable writing into the depth buffer

Parameters
    flag - Specifies whether the depth buffer is enabled for writing. If flag is false, depth buffer writing is disabled. Otherwise, it is enabled. Initially, depth buffer writing is enabled.
*/
func DepthMask(flag bool) {
	C.glDepthMask(boolToGLBool(flag))
}

/*
Specify mapping of depth values from normalized device coordinates to window coordinates

Parameters
    nearVal - Specifies the mapping of the near clipping plane to window coordinates. The initial value is 0.
    farVal - Specifies the mapping of the far clipping plane to window coordinates. The initial value is 1.
*/
func DepthRange(nearVal, farVal float32) {
	C.glDepthRange(C.GLclampd(nearVal), C.GLclampd(farVal))
}

/*
Detaches a shader object from a program object to which it is attached

Parameters
    program - Specifies the program object from which to detach the shader object.
    shader - Specifies the shader object to be detached.
*/
func DetachShader(program, shader uint32) {
	C.glDetachShader(C.GLuint(program), C.GLuint(shader))
}

/*
Disable server-side GL capabilities

Parameters
    capability - Specifies a symbolic constant indicating a GL capability.
*/
func Disable(capability GLenum) {
	C.glDisable(C.GLenum(capability))
}

/*
Disable a generic vertex attribute array

Parameters
    index - Specifies the index of the generic vertex attribute to be disabled.
*/
func DisableVertexAttribArray(index uint32) {
	C.glDisableVertexAttribArray(C.GLuint(index))
}

/*
Disable server-side GL capabilities

Parameters
    capability - Specifies a symbolic constant indicating a GL capability.
    index - Specifies the index of the switch to disable.
*/
func Disablei(capability GLenum, index uint32) {
	C.glDisablei(C.GLenum(capability), C.GLuint(index))
}

/*
Render primitives from array data

Parameters
    mode - Specifies what kind of primitives to render. Symbolic constants Points, LineStrip, LineLoop, Lines, LineStripAdjacency, LinesAdjacency, TriangleStrip, TriangleFan, Triangles, TriangleStripAdjacency and TrianglesAdjacency are accepted.
    first - Specifies the starting index in the enabled arrays.
    count - Specifies the number of indices to be rendered.
*/
func DrawArrays(mode GLenum, first, count int) {
	C.glDrawArrays(C.GLenum(mode), C.GLint(first), C.GLsizei(count))
}

/*
Draw multiple instances of a range of elements

Parameters
    mode - Specifies what kind of primitives to render. Symbolic constants Points, LineStrip, LineLoop, Lines, LineStripAdjacency, LinesAdjacency, TriangleStrip, TriangleFan, Triangles, TriangleStripAdjacency and TrianglesAdjacency are accepted.
    first - Specifies the starting index in the enabled arrays.
    count - Specifies the number of indices to be rendered.
    primcount - Specifies the number of instances of the specified range of indices to be rendered.
*/
func DrawArraysInstanced(mode GLenum, first, count, primcount int) {
	C.glDrawArraysInstanced(C.GLenum(mode), C.GLint(first), C.GLsizei(count), C.GLsizei(primcount))
}

/*
Specify which color buffers are to be drawn into

Parameters
    mode - Specifies up to four color buffers to be drawn into. Symbolic constants None, FrontLeft, FrontRight, BackLeft, BackRight, Front, Back, Left, Right, and FrontAndBack are accepted. The initial value is Front for single-buffered contexts, and Back for double-buffered contexts.
*/
func DrawBuffer(mode GLenum) {
	C.glDrawBuffer(C.GLenum(mode))
}

/*
Specifies a list of color buffers to be drawn into

Parameters
    bufs - Specifies an array of symbolic constants specifying the buffers into which fragment colors or data values will be written.
*/
func DrawBuffers(bufs []uint32) {
	count := len(bufs)
	C.glDrawBuffers(C.GLsizei(count), (*C.GLenum)(&bufs[0]))
}

/*
Render primitives from array data

Parameters
    mode - Specifies what kind of primitives to render. Symbolic constants Points, LineStrip, LineLoop, Lines, LineStripAdjacency, LinesAdjacency, TriangleStrip, TriangleFan, Triangles, TriangleStripAdjacency and TrianglesAdjacency are accepted.
    num - Amount of indices
    sliceType - The type of data in the buffer
*/
func DrawElementsBuffer(mode GLenum, num int, sliceType GLenum) {
	C.glDrawElements(C.GLenum(mode), C.GLsizei(num*glTypeSize(sliceType)), C.GLenum(sliceType), nil)
}

/*
Render primitives from array data

Parameters
    mode - Specifies what kind of primitives to render. Symbolic constants Points, LineStrip, LineLoop, Lines, LineStripAdjacency, LinesAdjacency, TriangleStrip, TriangleFan, Triangles, TriangleStripAdjacency and TrianglesAdjacency are accepted.
    data - Specifies a slice containing the indices.
*/
func DrawElements(mode GLenum, data interface{}) error {
	_, size, sliceType, ptr, err := sliceToGLData(data)
	if err != nil {
		return err
	}

	C.glDrawElements(C.GLenum(mode), size, sliceType, ptr)
	return nil
}

/*
Render primitives from array data with a per-element offset

Parameters
    mode - Specifies what kind of primitives to render. Symbolic constants Points, LineStrip, LineLoop, Lines, LineStripAdjacency, LinesAdjacency, TriangleStrip, TriangleFan, Triangles, TriangleStripAdjacency and TrianglesAdjacency are accepted.
    data - Specifies a slice containing the indices.
    basevertex - Specifies a constant that should be added to each element of indices when chosing elements from the enabled vertex arrays.
*/
func DrawElementsBaseVertex(mode GLenum, data interface{}, basevertex int) error {
	_, size, sliceType, ptr, err := sliceToGLData(data)
	if err != nil {
		return err
	}

	C.glDrawElementsBaseVertex(C.GLenum(mode), size, sliceType, ptr, C.GLint(basevertex))
	return nil
}

/*
Draw multiple instances of a set of elements

Parameters
    mode - Specifies what kind of primitives to render. Symbolic constants Points, LineStrip, LineLoop, Lines, LineStripAdjacency, LinesAdjacency, TriangleStrip, TriangleFan, Triangles, TriangleStripAdjacency and TrianglesAdjacency are accepted.
    data - Specifies a slice containing the indices.
    primcount - Specifies the number of instances of the specified range of indices to be rendered.
*/
func DrawElementsInstanced(mode GLenum, data interface{}, primcount int) error {
	_, size, sliceType, ptr, err := sliceToGLData(data)
	if err != nil {
		return err
	}

	C.glDrawElementsInstanced(C.GLenum(mode), size, sliceType, ptr, C.GLsizei(primcount))
	return nil
}

/*
Render multiple instances of a set of primitives from array data with a per-element offset

Parameters
    mode - Specifies what kind of primitives to render. Symbolic constants Points, LineStrip, LineLoop, Lines, LineStripAdjacency, LinesAdjacency, TriangleStrip, TriangleFan, Triangles, TriangleStripAdjacency and TrianglesAdjacency are accepted.
    data - Specifies a slice containing the indices.
    primcount - Specifies the number of instances of the specified range of indices to be rendered.
    basevertex - Specifies a constant that should be added to each element of indices when chosing elements from the enabled vertex arrays.
*/
func DrawElementsInstancedBaseVertex(mode GLenum, data interface{}, primcount, basevertex int) error {
	_, size, sliceType, ptr, err := sliceToGLData(data)
	if err != nil {
		return err
	}

	C.glDrawElementsInstancedBaseVertex(C.GLenum(mode), size, sliceType, ptr, C.GLsizei(primcount), C.GLint(basevertex))
	return nil
}

/*
Render primitives from array data

Parameters
    mode - Specifies what kind of primitives to render. Symbolic constants Points, LineStrip, LineLoop, Lines, LineStripAdjacency, LinesAdjacency, TriangleStrip, TriangleFan, Triangles, TriangleStripAdjacency and TrianglesAdjacency are accepted.
    start - Specifies the minimum array index contained in indices.
    end - Specifies the maximum array index contained in indices.
    count - Specifies the number of elements to be rendered.
    data - Specifies a slice containing the indices.
*/
func DrawRangeElements(mode GLenum, start, end uint32, count int, data interface{}) error {
	_, _, sliceType, ptr, err := sliceToGLData(data)
	if err != nil {
		return err
	}

	C.glDrawRangeElements(C.GLenum(mode), C.GLuint(start), C.GLuint(end), C.GLsizei(count), sliceType, ptr)
	return nil
}

/*
Render primitives from array data with a per-element offset

Parameters
    mode - Specifies what kind of primitives to render. Symbolic constants Points, LineStrip, LineLoop, Lines, LineStripAdjacency, LinesAdjacency, TriangleStrip, TriangleFan, Triangles, TriangleStripAdjacency and TrianglesAdjacency are accepted.
    start - Specifies the minimum array index contained in indices.
    end - Specifies the maximum array index contained in indices.
    count - Specifies the number of elements to be rendered.
    data - Specifies a slice containing the indices.
    basevertex - Specifies a constant that should be added to each element of indices when chosing elements from the enabled vertex arrays.
*/
func DrawRangeElementsBaseVertex(mode GLenum, start, end uint32, count int, data interface{}, basevertex int) error {
	_, _, sliceType, ptr, err := sliceToGLData(data)
	if err != nil {
		return err
	}

	C.glDrawRangeElementsBaseVertex(C.GLenum(mode), C.GLuint(start), C.GLuint(end), C.GLsizei(count), sliceType, ptr, C.GLint(basevertex))
	return nil
}

/*
Enable server-side GL capabilities

Parameters
    capability - Specifies a symbolic constant indicating a GL capability.
*/
func Enable(capability GLenum) {
	C.glEnable(C.GLenum(capability))
}

/*
Enable a generic vertex attribute array

Parameters
    index - Specifies the index of the generic vertex attribute to be enabled.
*/
func EnableVertexAttribArray(index uint32) {
	C.glEnableVertexAttribArray(C.GLuint(index))
}

/*
Enable server-side GL capabilities

Parameters
    capability - Specifies a symbolic constant indicating a GL capability.
    index - Specifies the index of the switch to enable.
*/
func Enablei(capability GLenum, index uint32) {
	C.glEnablei(C.GLenum(capability), C.GLuint(index))
}

/*
End conditional rendering
*/
func EndConditionalRender() {
	C.glEndConditionalRender()
}

/*
End delimiting the boundaries of a query object

Parameters
    target - Specifies the target type of query object to be concluded. The symbolic constant must be one of SamplesPassed, AnySamplesPassed, PrimitivesGenerated, TransformFeedbackPrimitivesWritten, or TimeElapsed.
*/
func EndQuery(target GLenum) {
	C.glEndQuery(C.GLenum(target))
}

/*
End transform feedback operation
*/
func EndTransformFeedback() {
	C.glEndTransformFeedback()
}

/*
Create a new sync object and insert it into the GL command stream

Parameters
    condition - Specifies the condition that must be met to set the sync object's state to signaled. condition must be SyncGpuCommandsComplete.
    flags - Specifies a bitwise combination of flags controlling the behavior of the sync object. No flags are presently defined for this operation and flags will be ignored.
*/
func FenceSync(condition GLenum, flags GLbitfield) GLsync {
	return GLsync(C.glFenceSync(C.GLenum(condition), 0))
}

/*
Block until all GL execution is complete
*/
func Finish() {
	C.glFinish()
}

/*
Force execution of GL commands in finite time
*/
func Flush() {
	C.glFlush()
}

/*
Indicate modifications to a range of a mapped buffer

Parameters
    target - Specifies the target of the flush operation. target must be ArrayBuffer, CopyReadBuffer, CopyWriteBuffer, ElementArrayBuffer, PixelPackBuffer, PixelUnpackBuffer, TextureBuffer, TransformFeedbackBuffer, or UniformBuffer.
    offset - Specifies the start of the buffer subrange, in basic machine units.
    length - Specifies the length of the buffer subrange, in basic machine units.
*/
func FlushMappedBufferRange(target GLenum, offset, length uint32) {
	C.glFlushMappedBufferRange(C.GLenum(target), C.GLintptr(offset), C.GLsizeiptr(length))
}

/*
Attach a renderbuffer as a logical buffer to the currently bound framebuffer object

Parameters
    target - Specifies the framebuffer target. target must be DrawFramebuffer, ReadFramebuffer, or Framebuffer. Framebuffer is equivalent to DrawFramebuffer.
    attachment - Specifies the attachment point of the framebuffer.
    rednerbuffertarget - Specifies the renderbuffer target and must be Renderbuffer.
    renderbuffer - Specifies the name of an existing renderbuffer object of type renderbuffertarget to attach.
*/
func FramebufferRenderbuffer(target, attachment, renderbuffertarget GLenum, renderbuffer uint32) {
	C.glFramebufferRenderbuffer(C.GLenum(target), C.GLenum(attachment), C.GLenum(renderbuffertarget), C.GLuint(renderbuffer))
}

/*
Attach a level of a texture object as a logical buffer to the currently bound framebuffer object

Parameters
    target - Specifies the framebuffer target. target must be DrawFramebuffer, ReadFramebuffer, or Framebuffer. Framebuffer is equivalent to DrawFramebuffer.
    attachment - Specifies the attachment point of the framebuffer. attachment must be ColorAttachmenti, DepthAttachment, StencilAttachment or DepthStencilAttachment.
    texture - Specifies the texture object to attach to the framebuffer attachment point named by attachment.
    level - Specifies the mipmap level of texture to attach.
*/
func FramebufferTexture(target, attachment GLenum, texture uint32, level int) {
	C.glFramebufferTexture(C.GLenum(target), C.GLenum(attachment), C.GLuint(texture), C.GLint(level))
}

/*
Attach a level of a texture object as a logical buffer to the currently bound framebuffer object

Parameters
    target - Specifies the framebuffer target. target must be DrawFramebuffer, ReadFramebuffer, or Framebuffer. Framebuffer is equivalent to DrawFramebuffer.
    attachment - Specifies the attachment point of the framebuffer. attachment must be ColorAttachmenti, DepthAttachment, StencilAttachment or DepthStencilAttachment.
    textarget - For FramebufferTexture1d, FramebufferTexture2d and FramebufferTexture3d, specifies what type of texture is expected in the texture parameter, or for cube map textures, which face is to be attached.
    texture - Specifies the texture object to attach to the framebuffer attachment point named by attachment.
    level - Specifies the mipmap level of texture to attach.
*/
func FramebufferTexture1d(target, attachment, textarget GLenum, texture uint32, level int) {
	C.glFramebufferTexture1D(C.GLenum(target), C.GLenum(attachment), C.GLenum(textarget), C.GLuint(texture), C.GLint(level))
}

/*
Attach a level of a texture object as a logical buffer to the currently bound framebuffer object

Parameters
    target - Specifies the framebuffer target. target must be DrawFramebuffer, ReadFramebuffer, or Framebuffer. Framebuffer is equivalent to DrawFramebuffer.
    attachment - Specifies the attachment point of the framebuffer. attachment must be ColorAttachmenti, DepthAttachment, StencilAttachment or DepthStencilAttachment.
    textarget - For FramebufferTexture1d, FramebufferTexture2d and FramebufferTexture3d, specifies what type of texture is expected in the texture parameter, or for cube map textures, which face is to be attached.
    texture - Specifies the texture object to attach to the framebuffer attachment point named by attachment.
    level - Specifies the mipmap level of texture to attach.
*/
func FramebufferTexture2d(target, attachment, textarget GLenum, texture uint32, level int) {
	C.glFramebufferTexture2D(C.GLenum(target), C.GLenum(attachment), C.GLenum(textarget), C.GLuint(texture), C.GLint(level))
}

/*
Attach a level of a texture object as a logical buffer to the currently bound framebuffer object

Parameters
    target - Specifies the framebuffer target. target must be DrawFramebuffer, ReadFramebuffer, or Framebuffer. Framebuffer is equivalent to DrawFramebuffer.
    attachment - Specifies the attachment point of the framebuffer. attachment must be ColorAttachmenti, DepthAttachment, StencilAttachment or DepthStencilAttachment.
    textarget - For FramebufferTexture1d, FramebufferTexture2d and FramebufferTexture3d, specifies what type of texture is expected in the texture parameter, or for cube map textures, which face is to be attached.
    texture - Specifies the texture object to attach to the framebuffer attachment point named by attachment.
    level - Specifies the mipmap level of texture to attach.
    layer - Specifies the layer of texture to attach.
*/
func FramebufferTexture3d(target, attachment, textarget GLenum, texture uint32, level, layer int) {
	C.glFramebufferTexture3D(C.GLenum(target), C.GLenum(attachment), C.GLenum(textarget), C.GLuint(texture), C.GLint(level), C.GLint(layer))
}

/*
Attach a single layer of a texture to a framebuffer

Parameters
    target - Specifies the framebuffer target. target must be DrawFramebuffer, ReadFramebuffer, or Framebuffer. Framebuffer is equivalent to DrawFramebuffer.
    attachment - Specifies the attachment point of the framebuffer. attachment must be ColorAttachmenti, DepthAttachment, StencilAttachment or DepthStencilAttachment.
    texture - Specifies the texture object to attach to the framebuffer attachment point named by attachment.
    level - Specifies the mipmap level of texture to attach.
    layer - Specifies the layer of texture to attach.
*/
func FramebufferTextureLayer(target, attachment GLenum, texture uint32, level, layer int) {
	C.glFramebufferTextureLayer(C.GLenum(target), C.GLenum(attachment), C.GLuint(texture), C.GLint(level), C.GLint(layer))
}

/*
Define front- and back-facing polygons

Parameters
    mode - Specifies the orientation of front-facing polygons. CW and CCW are accepted. The initial value is CCW.
*/
func FrontFace(mode GLenum) {
	C.glFrontFace(C.GLenum(mode))
}

/*
Generate buffer object names

Parameters
    num - Specifies the number of buffer object names to be generated.
*/
func GenBuffers(num int) []uint32 {
	cBuffers := make([]C.GLuint, num)
	buffers := make([]uint32, num)

	C.glGenBuffers(C.GLsizei(num), (*C.GLuint)(&cBuffers[0]))

	for key, _ := range cBuffers {
		buffers[key] = uint32(cBuffers[key])
	}

	return buffers
}

/*
Generate framebuffer object names

Parameters
    num - Specifies the number of framebuffer object names to be generated.
*/
func GenFramebuffers(num int) []uint32 {
	cBuffers := make([]C.GLuint, num)
	buffers := make([]uint32, num)

	C.glGenFramebuffers(C.GLsizei(num), (*C.GLuint)(&cBuffers[0]))

	for key, _ := range cBuffers {
		buffers[key] = uint32(cBuffers[key])
	}

	return buffers
}

/*
Generate query object names

Parameters
    num - Specifies the number of query object names to be generated.
*/
func GenQueries(num int) []uint32 {
	cQueries := make([]C.GLuint, num)
	queries := make([]uint32, num)

	C.glGenQueries(C.GLsizei(num), (*C.GLuint)(&cQueries[0]))

	for key, _ := range cQueries {
		queries[key] = uint32(cQueries[key])
	}

	return queries
}

/*
Generate renderbuffer object names

Parameters
    num - Specifies the number of renderbuffer object names to be generated.
*/
func GenRenderbuffers(num int) []uint32 {
	cBuffers := make([]C.GLuint, num)
	buffers := make([]uint32, num)

	C.glGenRenderbuffers(C.GLsizei(num), (*C.GLuint)(&cBuffers[0]))

	for key, _ := range cBuffers {
		buffers[key] = uint32(cBuffers[key])
	}

	return buffers
}

/*
Generate texture names

Parameters
    num - Specifies the number of texture names to be generated.
*/
func GenTextures(num int) []uint32 {
	cTextures := make([]C.GLuint, num)
	textures := make([]uint32, num)

	C.glGenTextures(C.GLsizei(num), (*C.GLuint)(&cTextures[0]))

	for key, _ := range cTextures {
		textures[key] = uint32(cTextures[key])
	}

	return textures
}

/*
Generate vertex array object names

Parameters
    num - Specifies the number of vertex array object names to be generated.
*/
func GenVertexArrays(num int) []uint32 {
	cVertexArrays := make([]C.GLuint, num)
	vertexArrays := make([]uint32, num)

	C.glGenVertexArrays(C.GLsizei(num), (*C.GLuint)(&cVertexArrays[0]))

	for key, _ := range cVertexArrays {
		vertexArrays[key] = uint32(cVertexArrays[key])
	}

	return vertexArrays
}

/*
Generate mipmaps for a specified texture target

Parameters
    target - Specifies the target to which the texture whose mimaps to generate is bound. target must be Texture1d, Texture2d, Texture3d, Texture1dArray, Texture2dArray or TextureCubeMap.
*/
func GenerateMipmap(target GLenum) {
	C.glGenerateMipmap(C.GLenum(target))
}

/*
Return the value or values of a selected parameter

Parameters
    pname - Specifies the parameter value to be returned.
*/
func GetBoolean(pname GLenum) []bool {
	cValues := make([]C.GLboolean, 1, 32)
	values := make([]bool, 1, 32)

	C.glGetBooleanv(C.GLenum(pname), (*C.GLboolean)(&cValues[0]))

	for key, _ := range cValues {
		values[key] = glBoolToBool(cValues[key])
	}

	return values
}

/*
Return the value or values of a selected parameter

Parameters
    pname - Specifies the parameter value to be returned.
*/
func GetDouble(pname GLenum) []float64 {
	cValues := make([]C.GLdouble, 1, 32)
	values := make([]float64, 1, 32)

	C.glGetDoublev(C.GLenum(pname), (*C.GLdouble)(&cValues[0]))

	for key, _ := range cValues {
		values[key] = float64(cValues[key])
	}

	return values
}

/*
Return the value or values of a selected parameter

Parameters
    pname - Specifies the parameter value to be returned.
*/
func GetFloat(pname GLenum) []float32 {
	cValues := make([]C.GLfloat, 1, 32)
	values := make([]float32, 1, 32)

	C.glGetFloatv(C.GLenum(pname), (*C.GLfloat)(&cValues[0]))

	for key, _ := range cValues {
		values[key] = float32(cValues[key])
	}

	return values
}

/*
Return the value or values of a selected parameter

Parameters
    pname - Specifies the parameter value to be returned.
*/
func GetInteger(pname GLenum) []int {
	cValues := make([]C.GLint, 1, 32)
	values := make([]int, 1, 32)

	C.glGetIntegerv(C.GLenum(pname), (*C.GLint)(&cValues[0]))

	for key, _ := range cValues {
		values[key] = int(cValues[key])
	}

	return values
}

/*
Return the value or values of a selected parameter

Parameters
    pname - Specifies the parameter value to be returned.
*/
func GetInteger64(pname GLenum) []int64 {
	cValues := make([]C.GLint64, 1, 32)
	values := make([]int64, 1, 32)

	C.glGetInteger64v(C.GLenum(pname), (*C.GLint64)(&cValues[0]))

	for key, _ := range cValues {
		values[key] = int64(cValues[key])
	}

	return values
}

/*
Return the value or values of a selected parameter

Parameters
    pname - Specifies the parameter value to be returned.
    index - Specifies the index of the particular element being queried.
*/
func GetBooleani(pname GLenum, index uint32) []bool {
	cValues := make([]C.GLboolean, 1, 32)
	values := make([]bool, 1, 32)

	C.glGetBooleani_v(C.GLenum(pname), C.GLuint(index), (*C.GLboolean)(&cValues[0]))

	for key, _ := range cValues {
		values[key] = glBoolToBool(cValues[key])
	}

	return values
}

/*
Return the value or values of a selected parameter

Parameters
    pname - Specifies the parameter value to be returned.
    index - Specifies the index of the particular element being queried.
*/
func GetIntegeri(pname GLenum, index uint32) []int {
	cValues := make([]C.GLint, 1, 32)
	values := make([]int, 1, 32)

	C.glGetIntegeri_v(C.GLenum(pname), C.GLuint(index), (*C.GLint)(&cValues[0]))

	for key, _ := range cValues {
		values[key] = int(cValues[key])
	}

	return values
}

/*
Return the value or values of a selected parameter

Parameters
    pname - Specifies the parameter value to be returned.
    index - Specifies the index of the particular element being queried.
*/
func GetInteger64i(pname GLenum, index uint32) []int64 {
	cValues := make([]C.GLint64, 1, 32)
	values := make([]int64, 1, 32)

	C.glGetInteger64i_v(C.GLenum(pname), C.GLuint(index), (*C.GLint64)(&cValues[0]))

	for key, _ := range cValues {
		values[key] = int64(cValues[key])
	}

	return values
}

/*
Returns information about an active attribute variable for the specified program object

Parameters
    program - Specifies the program object to be queried.
    index - Specifies the index of the attribute variable to be queried.
*/
func GetActiveAttrib(program, index uint32) (int, GLenum, string) {
	var cStr [256]C.GLchar
	var size C.GLint
	var enumType C.GLenum

	C.glGetActiveAttrib(C.GLuint(program), C.GLuint(index), 256, nil, &size, &enumType, &cStr[0])

	return int(size), GLenum(enumType), C.GoString((*C.char)(&cStr[0]))
}

/*
Returns information about an active uniform variable for the specified program object

Parameters
    program - Specifies the program object to be queried.
    index - Specifies the index of the uniform variable to be queried.
*/
func GetActiveUniform(program, index uint32) (int, GLenum, string) {
	var cStr [256]C.GLchar
	var size C.GLint
	var enumType C.GLenum

	C.glGetActiveUniform(C.GLuint(program), C.GLuint(index), 256, nil, &size, &enumType, &cStr[0])

	return int(size), GLenum(enumType), C.GoString((*C.char)(&cStr[0]))
}

/*
Retrieve the name of an active uniform block

Parameters
    program - Specifies the program containing the active uniform index uniformIndex.
    uniformBlockIndex - Specifies the index of the active uniform whose name to query.
*/
func GetActiveUniformName(program, uniformBlockIndex uint32) string {
	var cStr [256]C.GLchar
	var size C.GLsizei

	C.glGetActiveUniformName(C.GLuint(program), C.GLuint(uniformBlockIndex), 256, &size, &cStr[0])

	return C.GoString((*C.char)(&cStr[0]))
}

/*
Query the name of an active uniform

Parameters
    program - Specifies the name of a program containing the uniform block.
    uniformBlockIndex - Specifies the index of the uniform block within program.
*/
func GetActiveUniformBlockName(program, uniformBlockIndex uint32) string {
	var cStr [256]C.GLchar
	var size C.GLsizei

	C.glGetActiveUniformBlockName(C.GLuint(program), C.GLuint(uniformBlockIndex), 256, &size, &cStr[0])

	return C.GoString((*C.char)(&cStr[0]))
}

/*
Returns information about several active uniform variables for the specified program object

Parameters
    program - Specifies the program object to be queried.
    uniformIndices - Specifies the address of an array of uniformCount integers containing the indices of uniforms within program whose parameter pname should be queried.
    pname - Specifies the property of each uniform in uniformIndices that should be returned.
*/
func GetActiveUniforms(program uint32, uniformIndices []uint32, pname GLenum) []int {
	count := len(uniformIndices)
	cParams := make([]C.GLint, count)
	params := make([]int, count)

	C.glGetActiveUniformsiv(C.GLuint(program), C.GLsizei(count), (*C.GLuint)(&uniformIndices[0]), C.GLenum(pname), (*C.GLint)(&cParams[0]))

	for key, _ := range cParams {
		params[key] = int(cParams[key])
	}

	return params
}

/*
Returns the handles of the shader objects attached to a program object

Parameters
    program - Specifies the program object to be queried.
*/
func GetAttachedShaders(program uint32) []uint32 {
	var count C.GLsizei
	cShaders := make([]C.GLuint, 64)
	shaders := make([]uint32, 64)

	C.glGetAttachedShaders(C.GLuint(program), 64, &count, (*C.GLuint)(&shaders[0]))

	for key, _ := range cShaders {
		shaders[key] = uint32(cShaders[key])
	}

	return shaders
}

/*
Returns the location of an attribute variable

Parameters
    program - Specifies the program object to be queried.
    name - Points to a null terminated string containing the name of the attribute variable whose location is to be queried.
*/
func GetAttribLocation(program uint32, name string) uint32 {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return uint32(C.glGetAttribLocation(C.GLuint(program), (*C.GLchar)(cName)))
}

/*
Return parameters of a buffer object

Parameters
    target - Specifies the target buffer object. The symbolic constant must be ArrayBuffer, CopyReadBuffer, CopyWriteBuffer, ElementArrayBuffer, PixelPackBuffer, PixelUnpackBuffer, TextureBuffer, TransformFeedbackBuffer, or UniformBuffer.
    value - Specifies the symbolic name of a buffer object parameter. Accepted values are BufferAccess, BufferMapped, BufferSize, or BufferUsage.
*/
func GetBufferParameter(target, value GLenum) int {
	var data C.GLint

	C.glGetBufferParameteriv(C.GLenum(target), C.GLenum(value), &data)

	return int(data)
}

/*
Return parameters of a buffer object

Parameters
    target - Specifies the target buffer object. The symbolic constant must be ArrayBuffer, CopyReadBuffer, CopyWriteBuffer, ElementArrayBuffer, PixelPackBuffer, PixelUnpackBuffer, TextureBuffer, TransformFeedbackBuffer, or UniformBuffer.
    value - Specifies the symbolic name of a buffer object parameter. Accepted values are BufferAccess, BufferMapped, BufferSize, or BufferUsage.
*/
func GetBufferParameter64(target, value GLenum) int64 {
	var data C.GLint64

	C.glGetBufferParameteri64v(C.GLenum(target), C.GLenum(value), &data)

	return int64(data)
}

/*
Return the pointer to a mapped buffer object's data store

Parameters
    target - Specifies the target buffer object. The symbolic constant must be ArrayBuffer, CopyReadBuffer, CopyWriteBuffer, ElementArrayBuffer, PixelPackBuffer, PixelUnpackBuffer, TextureBuffer, TransformFeedbackBuffer, or UniformBuffer.
    pname - Specifies the pointer to be returned. The symbolic constant must be BufferMapPointer.
*/
func GetBufferPointer(target, pname GLenum) unsafe.Pointer {
	var params unsafe.Pointer

	C.glGetBufferPointerv(C.GLenum(target), C.GLenum(pname), &params)

	return params
}

/*
Returns a subset of a buffer object's data store

Parameters
    target - Specifies the target buffer object. The symbolic constant must be ArrayBuffer, CopyReadBuffer, CopyWriteBuffer, ElementArrayBuffer, PixelPackBuffer, PixelUnpackBuffer, TextureBuffer, TransformFeedbackBuffer, or UniformBuffer.
    offset - Specifies the offset into the buffer object's data store from which data will be returned, measured in bytes.
    size - Specifies the size in bytes of the data store region being returned.
*/
func GetBufferSubData(target GLenum, offset, size int) []byte {
	data := make([]byte, size)

	C.glGetBufferSubData(C.GLenum(target), C.GLintptr(offset), C.GLsizeiptr(size), unsafe.Pointer(&data[0]))

	return data
}

/*
Return a compressed texture image

Parameters
    target - Specifies which texture is to be obtained. Texture1d, Texture2d, Texture3d, TextureCubeMapPositiveX, TextureCubeMapNegativeX, TextureCubeMapPositiveY, TextureCubeMapNegativeY, TextureCubeMapPositiveZ, and TextureCubeMapNegativeZ are accepted.
    lod - Specifies the level-of-detail number of the desired image. Level 0 is the base image level. Level n is the nth mipmap reduction image.
*/
func GetCompressedTexImage(target GLenum, lod int) []byte {
	img := make([]byte, C.GL_TEXTURE_COMPRESSED_IMAGE_SIZE)

	C.glGetCompressedTexImage(C.GLenum(target), C.GLint(lod), unsafe.Pointer(&img[0]))

	return img
}

/*
Return error information
*/
func GetError() GLenum {
	return GLenum(C.glGetError())
}

/*
Query the bindings of color numbers to user-defined varying out variables

Parameters
    program - The name of the program containing varying out variable whose binding to query.
    name - The name of the user-defined varying out variable whose index to query.
*/
func GetFragDataLocation(program uint32, name string) int {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return int(C.glGetFragDataLocation(C.GLuint(program), (*C.GLchar)(cName)))
}

/*
Retrieve information about attachments of a bound framebuffer object

Parameters
    target - Specifies the target of the query operation.
    attachment - Specifies the attachment within target
    pname - Specifies the parameter of attachment to query.
*/
func GetFramebufferAttachmentParameter(target, attachment, pname GLenum) int {
	var params C.GLint

	C.glGetFramebufferAttachmentParameteriv(C.GLenum(target), C.GLenum(attachment), C.GLenum(pname), &params)

	return int(params)
}

/*
Retrieve the location of a sample

Parameters
    pname - Specifies the sample parameter name. pname must be SamplePosition.
    index - Specifies the index of the sample whose position to query.
*/
func GetMultisample(pname GLenum, index uint32) (float32, float32) {
	val := [2]C.GLfloat{}

	C.glGetMultisamplefv(C.GLenum(pname), C.GLuint(index), (*C.GLfloat)(unsafe.Pointer(&val[0])))

	return float32(val[0]), float32(val[1])
}

/*
Returns a parameter from a program object

Parameters
    program - Specifies the program object to be queried.
    pname - Specifies the object parameter. Accepted symbolic names are DeleteStatus, LinkStatus, ValidateStatus, InfoLogLength, AttachedShaders, ActiveAttributes, ActiveAttributeMaxLength, ActiveUniforms, ActiveUniformBlocks, ActiveUniformBlockMaxNameLength, ActiveUniformMaxLength, TransformFeedbackBufferMode, TransformFeedbackVaryings, TransformFeedbackVaryingMaxLength, GeometryVerticesOut, GeometryInputType, and GeometryOutputType.
*/
func GetProgram(program uint32, pname GLenum) int {
	var params C.GLint

	C.glGetProgramiv(C.GLuint(program), C.GLenum(pname), &params)

	return int(params)
}

/*
Returns the information log for a program object

Parameter
    program - Specifies the program object whose information log is to be queried.
*/
func GetProgramInfoLog(program uint32) string {
	var infoLog [256]C.GLchar

	C.glGetProgramInfoLog(C.GLuint(program), 256, nil, &infoLog[0])

	return C.GoString((*C.char)(unsafe.Pointer(&infoLog[0])))
}

/*
Return parameters of a query object

Parameters
    id - Specifies the name of a query object.
    pname - Specifies the symbolic name of a query object parameter. Accepted values are QueryResult or QueryResultAvailable.
*/
func GetQueryObject(id uint32, pname GLenum) int {
	var params C.GLint

	C.glGetQueryObjectiv(C.GLuint(id), C.GLenum(pname), &params)

	return int(params)
}

/*
Return parameters of a query object

Parameters
    id - Specifies the name of a query object.
    pname - Specifies the symbolic name of a query object parameter. Accepted values are QueryResult or QueryResultAvailable.
*/
func GetQueryObjectu(id uint32, pname GLenum) uint32 {
	var params C.GLuint

	C.glGetQueryObjectuiv(C.GLuint(id), C.GLenum(pname), &params)

	return uint32(params)
}

/*
Return parameters of a query object target

Parameters
    target - Specifies a query object target. Must be SamplesPassed, AnySamplesPassed, PrimitivesGenerated, TransformFeedbackPrimitivesWritten, TimeElapsed, or Timestamp.
    pname - Specifies the symbolic name of a query object target parameter. Accepted values are CurrentQuery or QueryCounterBits.
*/
func GetQuery(target, pname GLenum) int {
	var params C.GLint

	C.glGetQueryiv(C.GLenum(target), C.GLenum(pname), &params)

	return int(params)
}

/*
Retrieve information about a bound renderbuffer object

Parameters
    target - Specifies the target of the query operation. target must be Renderbuffer.
    pname - Specifies the parameter whose value to retrieve from the renderbuffer bound to target.
*/
func GetRenderbufferParameter(target, pname GLenum) int {
	var params C.GLint

	C.glGetRenderbufferParameteriv(C.GLenum(target), C.GLenum(pname), &params)

	return int(params)
}

/*
Returns a parameter from a shader object

Parameters
    shader - Specifies the shader object to be queried.
    pname - Specifies the object parameter. Accepted symbolic names are ShaderType, DeleteStatus, CompileStatus, InfoLogLength, ShaderSourceLength.
*/
func GetShader(shader uint32, pname GLenum) int {
	var params C.GLint

	C.glGetShaderiv(C.GLuint(shader), C.GLenum(pname), &params)

	return int(params)
}

/*
Returns the information log for a shader object

Parameter
    shader - Specifies the shader object whose information log is to be queried.
*/
func GetShaderInfoLog(shader uint32) string {
	var infoLog [256]C.GLchar

	C.glGetShaderInfoLog(C.GLuint(shader), 256, nil, &infoLog[0])

	return C.GoString((*C.char)(unsafe.Pointer(&infoLog[0])))
}

/*
Returns the information log for a shader object

Parameter
    shader - Specifies the shader object whose information log is to be queried.
*/
func GetShaderSource(shader uint32) string {
	var source [1024]C.GLchar

	C.glGetShaderSource(C.GLuint(shader), 1024, nil, &source[0])

	return C.GoString((*C.char)(unsafe.Pointer(&source[0])))
}

/*
Return a string describing the current GL connection

Parameters
    name - Specifies a symbolic constant, one of Vendor, Renderer, Version, or ShadingLanguageVersion.
*/
func GetString(name GLenum) string {
	return C.GoString((*C.char)(unsafe.Pointer(C.glGetString(C.GLenum(name)))))
}

/*
Return a string describing the current GL connection

Parameters
    name - Specifies a symbolic constant, one of Vendor, Renderer, Version, ShadingLanguageVersion, or Extensions.
    index - Specifies the index of the string to return.
*/
func GetStringi(name GLenum, index uint32) string {
	return C.GoString((*C.char)(unsafe.Pointer(C.glGetStringi(C.GLenum(name), C.GLuint(index)))))
}

/*
Query the properties of a sync object

Parameters
    sync - Specifies the sync object whose properties to query.
    pname - Specifies the parameter whose value to retrieve from the sync object specified in sync.
*/
func GetSync(sync GLsync, pname GLenum) []int {
	values := make([]int, 32)

	C.glGetSynciv(C.GLsync(sync), C.GLenum(pname), 32, nil, (*C.GLint)(unsafe.Pointer(&values[0])))

	return values
}

/*
Return a texture image

Parameters
    target - Specifies which texture is to be obtained. Texture1d, Texture2d, Texture3d, Texture1dArray, Texture2dArray, TextureRectangle, TextureCubeMapPositiveX, TextureCubeMapNegativeX, TextureCubeMapPositiveY, TextureCubeMapNegativeY, TextureCubeMapPositiveZ, and TextureCubeMapNegativeZ are accepted.
    level - Specifies the level-of-detail number of the desired image. Level 0 is the base image level. Level n is the nth mipmap reduction image.
    format - Specifies a pixel format for the returned data. The supported formats are StencilIndex, DepthComponent, DepthStencil, Red, Green, Blue, Rg, Rgb, Rgba, Bgr, Bgra, RedInteger, GreenInteger, BlueInteger, RgInteger, RgbInteger, RgbaInteger, BgrInteger, BgraInteger.
    enumType - Specifies a pixel type for the returned data. The supported types are UnsignedByte, Byte, UnsignedShort, Short, UnsignedInt, Int, HalfFloat, Float, UnsignedByte332, UnsignedByte233Rev, UnsignedShort565, UnsignedShort565Rev, UnsignedShort4444, UnsignedShort4444Rev, UnsignedShort5551, UnsignedShort1555Rev, UnsignedInt8888, UnsignedInt8888Rev, UnsignedInt1010102, UnsignedInt2101010Rev, UnsignedInt248, UnsignedInt10f11f11fRev, UnsignedInt5999Rev, and Float32UnsignedInt248Rev.
*/
func GetTexImage(target GLenum, level int, format, enumType GLenum) unsafe.Pointer {
	var img unsafe.Pointer

	C.glGetTexImage(C.GLenum(target), C.GLint(level), C.GLenum(format), C.GLenum(enumType), img)

	return img
}

/*
Return texture parameter values for a specific level of detail

Parameters
    target - Specifies the symbolic name of the target texture, one of Texture1d, Texture2d, Texture3d, Texture1dArray, Texture2dArray, TextureRectangle, Texture2dMultisample, Texture2dMultisampleArray, TextureCubeMapPositiveX, TextureCubeMapNegativeX, TextureCubeMapPositiveY, TextureCubeMapNegativeY, TextureCubeMapPositiveZ, TextureCubeMapNegativeZ, ProxyTexture1d, ProxyTexture2d, ProxyTexture3d, ProxyTexture1dArray, ProxyTexture2dArray, ProxyTextureRectangle, ProxyTexture2dMultisample, ProxyTexture2dMultisampleArray, ProxyTextureCubeMap, or TextureBuffer.
    level - Specifies the level-of-detail number of the desired image. Level 0 is the base image level. Level n is the nth mipmap reduction image.
    pname - Specifies the symbolic name of a texture parameter. TextureWidth, TextureHeight, TextureDepth, TextureInternalFormat, TextureRedSize, TextureGreenSize, TextureBlueSize, TextureAlphaSize, TextureDepthSize, TextureCompressed, and TextureCompressedImageSize are accepted.
*/
func GetTexLevelParameterf(target GLenum, level int, pname GLenum) float32 {
	var params C.GLfloat

	C.glGetTexLevelParameterfv(C.GLenum(target), C.GLint(level), C.GLenum(pname), &params)

	return float32(params)
}

/*
Return texture parameter values for a specific level of detail

Parameters
    target - Specifies the symbolic name of the target texture, one of Texture1d, Texture2d, Texture3d, Texture1dArray, Texture2dArray, TextureRectangle, Texture2dMultisample, Texture2dMultisampleArray, TextureCubeMapPositiveX, TextureCubeMapNegativeX, TextureCubeMapPositiveY, TextureCubeMapNegativeY, TextureCubeMapPositiveZ, TextureCubeMapNegativeZ, ProxyTexture1d, ProxyTexture2d, ProxyTexture3d, ProxyTexture1dArray, ProxyTexture2dArray, ProxyTextureRectangle, ProxyTexture2dMultisample, ProxyTexture2dMultisampleArray, ProxyTextureCubeMap, or TextureBuffer.
    level - Specifies the level-of-detail number of the desired image. Level 0 is the base image level. Level n is the nth mipmap reduction image.
    pname - Specifies the symbolic name of a texture parameter. TextureWidth, TextureHeight, TextureDepth, TextureInternalFormat, TextureRedSize, TextureGreenSize, TextureBlueSize, TextureAlphaSize, TextureDepthSize, TextureCompressed, and TextureCompressedImageSize are accepted.
*/
func GetTexLevelParameteri(target GLenum, level int, pname GLenum) int {
	var params C.GLint

	C.glGetTexLevelParameteriv(C.GLenum(target), C.GLint(level), C.GLenum(pname), &params)

	return int(params)
}

/*
Return texture parameter values

Parameters
    target - Specifies the symbolic name of the target texture. Texture1d, Texture2d, Texture1dArray, Texture2dArray, Texture3d, TextureRectangle, and TextureCubeMap are accepted.
    pname - Specifies the symbolic name of a texture parameter. TextureBaseLevel, TextureBorderColor, TextureCompareMode, TextureCompareFunc, TextureLodBias, TextureMagFilter, TextureMaxLevel, TextureMaxLod, TextureMinFilter, TextureMinLod, TextureSwizzleR, TextureSwizzleG, TextureSwizzleB, TextureSwizzleA, TextureSwizzleRgba, TextureWrapS, TextureWrapT, and TextureWrapR are accepted.
*/
func GetTexParameterf(target, pname GLenum) []float32 {
	params := make([]float32, 4)

	C.glGetTexParameterfv(C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(&params[0])))

	return params
}

/*
Return texture parameter values

Parameters
    target - Specifies the symbolic name of the target texture. Texture1d, Texture2d, Texture1dArray, Texture2dArray, Texture3d, TextureRectangle, and TextureCubeMap are accepted.
    pname - Specifies the symbolic name of a texture parameter. TextureBaseLevel, TextureBorderColor, TextureCompareMode, TextureCompareFunc, TextureLodBias, TextureMagFilter, TextureMaxLevel, TextureMaxLod, TextureMinFilter, TextureMinLod, TextureSwizzleR, TextureSwizzleG, TextureSwizzleB, TextureSwizzleA, TextureSwizzleRgba, TextureWrapS, TextureWrapT, and TextureWrapR are accepted.
*/
func GetTexParameteri(target, pname GLenum) []int {
	params := make([]int, 4)

	C.glGetTexParameteriv(C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(&params[0])))

	return params
}

/*
Return texture parameter values

Parameters
    target - Specifies the symbolic name of the target texture. Texture1d, Texture2d, Texture1dArray, Texture2dArray, Texture3d, TextureRectangle, and TextureCubeMap are accepted.
    pname - Specifies the symbolic name of a texture parameter. TextureBaseLevel, TextureBorderColor, TextureCompareMode, TextureCompareFunc, TextureLodBias, TextureMagFilter, TextureMaxLevel, TextureMaxLod, TextureMinFilter, TextureMinLod, TextureSwizzleR, TextureSwizzleG, TextureSwizzleB, TextureSwizzleA, TextureSwizzleRgba, TextureWrapS, TextureWrapT, and TextureWrapR are accepted.
*/
func GetTexParameterIi(target, pname GLenum) []int {
	params := make([]int, 4)

	C.glGetTexParameterIiv(C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(&params[0])))

	return params
}

/*
Return texture parameter values

Parameters
    target - Specifies the symbolic name of the target texture. Texture1d, Texture2d, Texture1dArray, Texture2dArray, Texture3d, TextureRectangle, and TextureCubeMap are accepted.
    pname - Specifies the symbolic name of a texture parameter. TextureBaseLevel, TextureBorderColor, TextureCompareMode, TextureCompareFunc, TextureLodBias, TextureMagFilter, TextureMaxLevel, TextureMaxLod, TextureMinFilter, TextureMinLod, TextureSwizzleR, TextureSwizzleG, TextureSwizzleB, TextureSwizzleA, TextureSwizzleRgba, TextureWrapS, TextureWrapT, and TextureWrapR are accepted.
*/
func GetTexParameterIui(target, pname GLenum) []uint32 {
	params := make([]uint32, 4)

	C.glGetTexParameterIuiv(C.GLenum(target), C.GLenum(pname), (*C.GLuint)(unsafe.Pointer(&params[0])))

	return params
}

/*
Retrieve information about varying variables selected for transform feedback

Parameters
    program - The name of the target program object.
    index - The index of the varying variable whose information to retrieve.
*/
func GetTransformFeedbackVarying(program, index uint32) (int, GLenum, string) {
	var cStr [256]C.GLchar
	var size C.GLsizei
	var enumType C.GLenum

	C.glGetTransformFeedbackVarying(C.GLuint(program), C.GLuint(index), 256, nil, &size, &enumType, &cStr[0])

	return int(size), GLenum(enumType), C.GoString((*C.char)(&cStr[0]))
}

/*
Returns the value of a uniform variable

Parameters
    program - Specifies the program object to be queried.
    location - Specifies the location of the uniform variable to be queried.
*/
func GetUniformf(program uint32, location int) []float32 {
	params := make([]float32, 16)

	C.glGetUniformfv(C.GLuint(program), C.GLint(location), (*C.GLfloat)(unsafe.Pointer(&params[0])))

	return params
}

/*
Returns the value of a uniform variable

Parameters
    program - Specifies the program object to be queried.
    location - Specifies the location of the uniform variable to be queried.
*/
func GetUniformi(program uint32, location int) []int {
	params := make([]int, 16)

	C.glGetUniformiv(C.GLuint(program), C.GLint(location), (*C.GLint)(unsafe.Pointer(&params[0])))

	return params
}

/*
Returns the value of a uniform variable

Parameters
    program - Specifies the program object to be queried.
    location - Specifies the location of the uniform variable to be queried.
*/
func GetUniformui(program uint32, location int) []uint32 {
	params := make([]uint32, 16)

	C.glGetUniformuiv(C.GLuint(program), C.GLint(location), (*C.GLuint)(unsafe.Pointer(&params[0])))

	return params
}

/*
Retrieve the index of a named uniform block

Parameters
    program - Specifies the name of a program containing the uniform block.
    uniformBlockName - Specifies the name of the uniform block whose index to retrieve.
*/
func GetUniformBlockIndex(program uint32, uniformBlockName string) uint32 {
	cName := C.CString(uniformBlockName)
	defer C.free(unsafe.Pointer(cName))

	return uint32(C.glGetUniformBlockIndex(C.GLuint(program), (*C.GLchar)(cName)))
}

/*
Retrieve the index of a named uniform block

Parameters
    program - Specifies the name of a program containing uniforms whose indices to query.
    uniformCount - Specifies the number of uniforms whose indices to query.
    uniformNames - Specifies the address of an array of pointers to buffers containing the names of the queried uniforms.
*/
func GetUniformIndices(program uint32, uniformCount int, uniformNames []string) []uint32 {
	count := len(uniformNames)
	cNames := make([]*C.char, count)
	uniformIndices := make([]uint32, count)

	for key, _ := range uniformNames {
		cNames[key] = C.CString(uniformNames[key])
	}

	C.glGetUniformIndices(C.GLuint(program), C.GLsizei(uniformCount), (**C.GLchar)(unsafe.Pointer(&cNames[0])), (*C.GLuint)(unsafe.Pointer(&uniformIndices[0])))

	for key, _ := range uniformNames {
		C.free(unsafe.Pointer(cNames[key]))
	}

	return uniformIndices
}

/*
Returns the location of a uniform variable

Parameters
    program - Specifies the name of a program containing the uniform block.
    name - Specifies the name of the uniform variable whose location is to be queried.
*/
func GetUniformLocation(program uint32, name string) int {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return int(C.glGetUniformLocation(C.GLuint(program), (*C.GLchar)(cName)))
}

/*
Return a generic vertex attribute parameter

Parameters
    index - Specifies the generic vertex attribute parameter to be queried.
    pname - Specifies the symbolic name of the vertex attribute parameter to be queried. Accepted values are VertexAttribArrayBufferBinding, VertexAttribArrayEnabled, VertexAttribArraySize, VertexAttribArrayStride, VertexAttribArrayType, VertexAttribArrayNormalized, VertexAttribArrayInteger, VertexAttribArrayDivisor, or CurrentVertexAttrib.
*/
func GetVertexAttribd(index uint32, pname GLenum) []float64 {
	params := make([]float64, 4)

	C.glGetVertexAttribdv(C.GLuint(index), C.GLenum(pname), (*C.GLdouble)(unsafe.Pointer(&params[0])))

	return params
}

/*
Return a generic vertex attribute parameter

Parameters
    index - Specifies the generic vertex attribute parameter to be queried.
    pname - Specifies the symbolic name of the vertex attribute parameter to be queried. Accepted values are VertexAttribArrayBufferBinding, VertexAttribArrayEnabled, VertexAttribArraySize, VertexAttribArrayStride, VertexAttribArrayType, VertexAttribArrayNormalized, VertexAttribArrayInteger, VertexAttribArrayDivisor, or CurrentVertexAttrib.
*/
func GetVertexAttribf(index uint32, pname GLenum) []float32 {
	params := make([]float32, 4)

	C.glGetVertexAttribfv(C.GLuint(index), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(&params[0])))

	return params
}

/*
Return a generic vertex attribute parameter

Parameters
    index - Specifies the generic vertex attribute parameter to be queried.
    pname - Specifies the symbolic name of the vertex attribute parameter to be queried. Accepted values are VertexAttribArrayBufferBinding, VertexAttribArrayEnabled, VertexAttribArraySize, VertexAttribArrayStride, VertexAttribArrayType, VertexAttribArrayNormalized, VertexAttribArrayInteger, VertexAttribArrayDivisor, or CurrentVertexAttrib.
*/
func GetVertexAttribi(index uint32, pname GLenum) []int {
	params := make([]int, 4)

	C.glGetVertexAttribiv(C.GLuint(index), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(&params[0])))

	return params
}

/*
Return a generic vertex attribute parameter

Parameters
    index - Specifies the generic vertex attribute parameter to be queried.
    pname - Specifies the symbolic name of the vertex attribute parameter to be queried. Accepted values are VertexAttribArrayBufferBinding, VertexAttribArrayEnabled, VertexAttribArraySize, VertexAttribArrayStride, VertexAttribArrayType, VertexAttribArrayNormalized, VertexAttribArrayInteger, VertexAttribArrayDivisor, or CurrentVertexAttrib.
*/
func GetVertexAttribIi(index uint32, pname GLenum) []int {
	params := make([]int, 4)

	C.glGetVertexAttribIiv(C.GLuint(index), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(&params[0])))

	return params
}

/*
Return a generic vertex attribute parameter

Parameters
    index - Specifies the generic vertex attribute parameter to be queried.
    pname - Specifies the symbolic name of the vertex attribute parameter to be queried. Accepted values are VertexAttribArrayBufferBinding, VertexAttribArrayEnabled, VertexAttribArraySize, VertexAttribArrayStride, VertexAttribArrayType, VertexAttribArrayNormalized, VertexAttribArrayInteger, VertexAttribArrayDivisor, or CurrentVertexAttrib.
*/
func GetVertexAttribIui(index uint32, pname GLenum) []uint32 {
	params := make([]uint32, 4)

	C.glGetVertexAttribIuiv(C.GLuint(index), C.GLenum(pname), (*C.GLuint)(unsafe.Pointer(&params[0])))

	return params
}

/*
Return the address of the specified generic vertex attribute pointer

Parameters
    index - Specifies the generic vertex attribute parameter to be returned.
    pname - Specifies the symbolic name of the generic vertex attribute parameter to be returned. Must be VertexAttribArrayPointer.
*/
func GetVertexAttribPointer(index uint32, pname GLenum) unsafe.Pointer {
	var pointer unsafe.Pointer

	C.glGetVertexAttribPointerv(C.GLuint(index), C.GLenum(pname), &pointer)

	return pointer
}

/*
Specify implementation-specific hints

Parameters
    target - Specifies a symbolic constant indicating the behavior to be controlled. LineSmoothHint, PolygonSmoothHint, TextureCompressionHint, and FragmentShaderDerivativeHint are accepted.
    mode - Specifies a symbolic constant indicating the desired behavior. Fastest, Nicest, and DontCare are accepted.
*/
func Hint(target, mode GLenum) {
	C.glHint(C.GLenum(target), C.GLenum(mode))
}

/*
Determine if a name corresponds to a buffer object

Parameters
    buffer - Specifies a value that may be the name of a buffer object.
*/
func IsBuffer(buffer uint32) bool {
	return glBoolToBool(C.glIsBuffer(C.GLuint(buffer)))
}

/*
Test whether a capability is enabled

Parameters
    capacity - Specifies a symbolic constant indicating a GL capability.
*/
func IsEnabled(capacity GLenum) bool {
	return glBoolToBool(C.glIsEnabled(C.GLenum(capacity)))
}

/*
Test whether a capability is enabled

Parameters
    capacity - Specifies a symbolic constant indicating a GL capability.
    index - Specifies the index of the capability.
*/
func IsEnabledi(capacity GLenum, index uint32) bool {
	return glBoolToBool(C.glIsEnabledi(C.GLenum(capacity), C.GLuint(index)))
}

/*
Determine if a name corresponds to a framebuffer object

Parameters
    framebuffer - Specifies a value that may be the name of a framebuffer object.
*/
func IsFramebuffer(framebuffer uint32) bool {
	return glBoolToBool(C.glIsFramebuffer(C.GLuint(framebuffer)))
}

/*
Determine if a name corresponds to a program object

Parameters
    program - Specifies a value that may be the name of a program object.
*/
func IsProgram(program uint32) bool {
	return glBoolToBool(C.glIsProgram(C.GLuint(program)))
}

/*
Determine if a name corresponds to a query object

Parameters
    id - Specifies a value that may be the name of a query object.
*/
func IsQuery(id uint32) bool {
	return glBoolToBool(C.glIsQuery(C.GLuint(id)))
}

/*
Determine if a name corresponds to a renderbuffer object

Parameters
    renderbuffer - Specifies a value that may be the name of a renderbuffer object.
*/
func IsRenderbuffer(renderbuffer uint32) bool {
	return glBoolToBool(C.glIsRenderbuffer(C.GLuint(renderbuffer)))
}

/*
Determine if a name corresponds to a shader object

Parameters
    shader - Specifies a value that may be the name of a shader object.
*/
func IsShader(shader uint32) bool {
	return glBoolToBool(C.glIsShader(C.GLuint(shader)))
}

/*
Determine if a name corresponds to a sync object

Parameters
    sync - Specifies a value that may be the name of a sync object.
*/
func IsSync(sync GLsync) bool {
	return glBoolToBool(C.glIsSync(C.GLsync(sync)))
}

/*
Determine if a name corresponds to a texture object

Parameters
    texture - Specifies a value that may be the name of a texture object.
*/
func IsTexture(texture uint32) bool {
	return glBoolToBool(C.glIsTexture(C.GLuint(texture)))
}

/*
Determine if a name corresponds to a vertex array object

Parameters
    array - Specifies a value that may be the name of a vertex array object.
*/
func IsVertexArray(array uint32) bool {
	return glBoolToBool(C.glIsVertexArray(C.GLuint(array)))
}

/*
Specify the width of rasterized lines

Parameters
    width - Specifies the width of rasterized lines. The initial value is 1.
*/
func LineWidth(width float32) {
	C.glLineWidth(C.GLfloat(width))
}

/*
Links a program object

Parameters
    program - Specifies the handle of the program object to be linked.
*/
func LinkProgram(program uint32) {
	C.glLinkProgram(C.GLuint(program))
}

/*
Specify a logical pixel operation for rendering

Parameters
    opcode - Specifies a symbolic constant that selects a logical operation. The following symbols are accepted: Clear, Set, Copy, CopyInverted, Noop, Invert, And, Nand, Or, Nor, Xor, Equiv, AndReverse, AndInverted, OrReverse, and OrInverted. The initial value is Copy.
*/
func LogicOp(opcode GLenum) {
	C.glLogicOp(C.GLenum(opcode))
}

/*
Map a buffer object's data store

Parameters
    target - Specifies the target buffer object being mapped. The symbolic constant must be ArrayBuffer, CopyReadBuffer, CopyWriteBuffer, ElementArrayBuffer, PixelPackBuffer, PixelUnpackBuffer, TextureBuffer, TransformFeedbackBuffer or UniformBuffer.
    access - Specifies the access policy, indicating whether it will be possible to read from, write to, or both read from and write to the buffer object's mapped data store. The symbolic constant must be ReadOnly, WriteOnly, or ReadWrite.
*/
func MapBuffer(target, access GLenum) unsafe.Pointer {
	return C.glMapBuffer(C.GLenum(target), C.GLenum(access))
}

/*
Map a section of a buffer object's data store

Parameters
    target - Specifies the target buffer object being mapped. The symbolic constant must be ArrayBuffer, CopyReadBuffer, CopyWriteBuffer, ElementArrayBuffer, PixelPackBuffer, PixelUnpackBuffer, TextureBuffer, TransformFeedbackBuffer or UniformBuffer.
    offset - Specifies a the starting offset within the buffer of the range to be mapped.
    length - Specifies a length of the range to be mapped.
    access - Specifies the access policy, indicating whether it will be possible to read from, write to, or both read from and write to the buffer object's mapped data store. The symbolic constant must be ReadOnly, WriteOnly, or ReadWrite.
*/
func MapBufferRange(target GLenum, offset, length int, access GLbitfield) unsafe.Pointer {
	return C.glMapBufferRange(C.GLenum(target), C.GLintptr(offset), C.GLsizeiptr(length), C.GLbitfield(access))
}

/*
Render multiple sets of primitives from array data

Parameters
    mode - Specifies what kind of primitives to render. Symbolic constants Points, LineStrip, LineLoop, Lines, LineStripAdjacency, LinesAdjacency, TriangleStrip, TriangleFan, Triangles, TriangleStripAdjacency and TrianglesAdjacency are accepted.
    first - An array of starting indices in the enabled arrays.
    count - An array of the number of indices to be rendered.
*/
func MultiDrawArrays(mode GLenum, first, count []int) {
	C.glMultiDrawArrays(C.GLenum(mode), (*C.GLint)(unsafe.Pointer(&first[0])), (*C.GLsizei)(unsafe.Pointer(&count[0])), C.GLsizei(len(first)))
}

/*
Render multiple sets of primitives by specifying indices of array data elements

Parameters
    mode - Specifies what kind of primitives to render. Symbolic constants Points, LineStrip, LineLoop, Lines, LineStripAdjacency, LinesAdjacency, TriangleStrip, TriangleFan, Triangles, TriangleStripAdjacency and TrianglesAdjacency are accepted.
    count - An array of the elements counts.
    indices - An array of where the indices are stored.
*/
func MultiDrawElements(mode GLenum, count []int, indices []interface{}) error {
	_, size, sliceType, _, err := sliceToGLData(indices[0])
	if err != nil {
		return err
	}
	ptr := unsafe.Pointer(&indices[0])

	C.glMultiDrawElements(C.GLenum(mode), (*C.GLsizei)(unsafe.Pointer(&count[0])), sliceType, &ptr, size)

	return nil
}

/*
Render multiple sets of primitives by specifying indices of array data elements and an index to apply to each index

Parameters
    mode - Specifies what kind of primitives to render. Symbolic constants Points, LineStrip, LineLoop, Lines, LineStripAdjacency, LinesAdjacency, TriangleStrip, TriangleFan, Triangles, TriangleStripAdjacency and TrianglesAdjacency are accepted.
    count - An array of the elements counts.
    indices - An array of where the indices are stored.
    basevertex - An array of where the base vertices are stored.
*/
func MultiDrawElementsBaseVertex(mode GLenum, count []int, indices []interface{}, basevertex []int) error {
	_, size, sliceType, _, err := sliceToGLData(indices[0])
	if err != nil {
		return err
	}
	ptr := unsafe.Pointer(&indices[0])

	C.glMultiDrawElementsBaseVertex(C.GLenum(mode), (*C.GLsizei)(unsafe.Pointer(&count[0])), sliceType, &ptr, size, (*C.GLint)(unsafe.Pointer(&basevertex[0])))

	return nil
}

/*
Set pixel storage modes

Parameters
    pname - Specifies the symbolic name of the parameter to be set. Six values affect the packing of pixel data into memory: PackSwapBytes, PackLsbFirst, PackRowLength, PackImageHeight, PackSkipPixels, PackSkipRows, PackSkipImages, and PackAlignment. Six more affect the unpacking of pixel data from memory: UnpackSwapBytes, UnpackLsbFirst, UnpackRowLength, UnpackImageHeight, UnpackSkipPixels, UnpackSkipRows, UnpackSkipImages, and UnpackAlignment.
    param - Specifies the value that pname is set to.
*/
func PixelStoref(pname GLenum, param float32) {
	C.glPixelStoref(C.GLenum(pname), C.GLfloat(param))
}

/*
Set pixel storage modes

Parameters
    pname - Specifies the symbolic name of the parameter to be set. Six values affect the packing of pixel data into memory: PackSwapBytes, PackLsbFirst, PackRowLength, PackImageHeight, PackSkipPixels, PackSkipRows, PackSkipImages, and PackAlignment. Six more affect the unpacking of pixel data from memory: UnpackSwapBytes, UnpackLsbFirst, UnpackRowLength, UnpackImageHeight, UnpackSkipPixels, UnpackSkipRows, UnpackSkipImages, and UnpackAlignment.
    param - Specifies the value that pname is set to.
*/
func PixelStorei(pname GLenum, param int) {
	C.glPixelStorei(C.GLenum(pname), C.GLint(param))
}

/*
Specify point parameters

Parameters
    pname - Specifies a single-valued point parameter. PointFadeThresholdSize, and PointSpriteCoordOrigin are accepted.
    param - Specifies the value that pname is set to.
*/
func PointParameterf(pname GLenum, param float32) {
	C.glPointParameterf(C.GLenum(pname), C.GLfloat(param))
}

/*
Specify point parameters

Parameters
    pname - Specifies a single-valued point parameter. PointFadeThresholdSize, and PointSpriteCoordOrigin are accepted.
    param - Specifies the value that pname is set to.
*/
func PointParameteri(pname GLenum, param int) {
	C.glPointParameteri(C.GLenum(pname), C.GLint(param))
}

/*
Specify the diameter of rasterized points

Parameters
    size - Specifies the diameter of rasterized points. The initial value is 1.
*/
func PointSize(size float32) {
	C.glPointSize(C.GLfloat(size))
}

/*
Select a polygon rasterization mode

Parameters
    face - Specifies the polygons that mode applies to. Must be FrontAndBack for front- and back-facing polygons.
    mode - Specifies how polygons will be rasterized. Accepted values are Point, Line, and Fill. The initial value is Fill for both front- and back-facing polygons.
*/
func PolygonMode(face, mode GLenum) {
	C.glPolygonMode(C.GLenum(face), C.GLenum(mode))
}

/*
Set the scale and units used to calculate depth values

Parameters
    factor - Specifies a scale factor that is used to create a variable depth offset for each polygon. The initial value is 0.
    units - Is multiplied by an implementation-specific value to create a constant depth offset. The initial value is 0.
*/
func PolygonOffset(factor, units float32) {
	C.glPolygonOffset(C.GLfloat(factor), C.GLfloat(units))
}

/*
Specify the primitive restart index

Parameters
    index - Specifies the value to be interpreted as the primitive restart index.
*/
func PrimitiveRestartIndex(index uint32) {
	C.glPrimitiveRestartIndex(C.GLuint(index))
}

/*
Specifiy the vertex to be used as the source of data for flat shaded varyings

Parameters
    provokeMode - Specifies the vertex to be used as the source of data for flat shaded varyings.
*/
func ProvokingVertex(provokeMode GLenum) {
	C.glProvokingVertex(C.GLenum(provokeMode))
}

/*
Select a color buffer source for pixels

Parameters
    mode - Specifies a color buffer. Accepted values are FrontLeft, FrontRight, BackLeft, BackRight, Front, Back, Left, and the constants ColorAttachmenti.
*/
func ReadBuffer(mode GLenum) {
	C.glReadBuffer(C.GLenum(mode))
}

/*
Read a block of pixels from the frame buffer

Parameters
    x, y - Specify the window coordinates of the first pixel that is read from the frame buffer. This location is the lower left corner of a rectangular block of pixels.
    width, height - Specify the dimensions of the pixel rectangle. width and height of one correspond to a single pixel.
    format - Specifies the format of the pixel data. The following symbolic values are accepted: StencilIndex, DepthComponent, DepthStencil, Red, Green, Blue, Rgb, Bgr, Rgba, and Bgra.
    enumType - Specifies the data type of the pixel data. Must be one of UnsignedByte, Byte, UnsignedShort, Short, UnsignedInt, Int, HalfFloat, Float, UnsignedByte332, UnsignedByte233Rev, UnsignedShort565, UnsignedShort565Rev, UnsignedShort4444, UnsignedShort4444Rev, UnsignedShort5551, UnsignedShort1555Rev, UnsignedInt8888, UnsignedInt8888Rev, UnsignedInt1010102, UnsignedInt2101010Rev, UnsignedInt248, UnsignedInt10f11f11fRev, UnsignedInt5999Rev, or Float32UnsignedInt248Rev.
*/
func ReadPixels(x, y, width, height int, format, enumType GLenum) unsafe.Pointer {
	var data unsafe.Pointer

	C.glReadPixels(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(enumType), data)

	return data
}

/*
Establish data storage, format and dimensions of a renderbuffer object's image

Parameters
    target - Specifies a binding to which the target of the allocation and must be Renderbuffer.
    internalFormat - Specifies the internal format to use for the renderbuffer object's image.
    width - Specifies the width of the renderbuffer, in pixels.
    height - Specifies the height of the renderbuffer, in pixels.
*/
func RenderbufferStorage(target, internalFormat GLenum, width, height int) {
	C.glRenderbufferStorage(C.GLenum(target), C.GLenum(internalFormat), C.GLsizei(width), C.GLsizei(height))
}

/*
Establish data storage, format, dimensions and sample count of a renderbuffer object's image

Parameters
    target - Specifies a binding to which the target of the allocation and must be Renderbuffer.
    samples - Specifies the number of samples to be used for the renderbuffer object's storage.
    internalFormat - Specifies the internal format to use for the renderbuffer object's image.
    width - Specifies the width of the renderbuffer, in pixels.
    height - Specifies the height of the renderbuffer, in pixels.
*/
func RenderbufferStorageMultisample(target GLenum, samples int, internalFormat GLenum, width, height int) {
	C.glRenderbufferStorageMultisample(C.GLenum(target), C.GLsizei(samples), C.GLenum(internalFormat), C.GLsizei(width), C.GLsizei(height))
}

/*
Specify multisample coverage parameters

Parameters
    value - Specify a single floating-point sample coverage value. The value is clamped to the range 0 1 . The initial value is 1.0.
    invert - Specify a single boolean value representing if the coverage masks should be inverted. The initial value is false.
*/
func SampleCoverage(value float32, invert bool) {
	C.glSampleCoverage(C.GLclampf(value), boolToGLBool(invert))
}

/*
Set the value of a sub-word of the sample mask

Parameters
    maskNumber - Specifies which 32-bit sub-word of the sample mask to update.
    mask - Specifies the new value of the mask sub-word.
*/
func SampleMask(maskNumber uint32, mask GLbitfield) {
	C.glSampleMaski(C.GLuint(maskNumber), C.GLbitfield(mask))
}

/*
Define the scissor box

Parameters
    x, y - Specify the lower left corner of the scissor box. Initially (0, 0).
    width, height - Specify the width and height of the scissor box. When a GL context is first attached to a window, width and height are set to the dimensions of that window.
*/
func Scissor(x, y, width, height int) {
	C.glScissor(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}

/*
Replaces the source code in a shader object

Parameters
    shader - Specifies the handle of the shader object whose source code is to be replaced.
    source - Specifies one or several strings containing the source code of the shader.
*/
func ShaderSource(shader uint32, source ...string) {
	count := len(source)
	str := make([]*C.char, count)

	for key, _ := range source {
		str[key] = C.CString(source[key])
	}

	C.glShaderSource(C.GLuint(shader), C.GLsizei(count), (**C.GLchar)(unsafe.Pointer(&str[0])), nil)

	for key, _ := range str {
		C.free(unsafe.Pointer(str[key]))
	}
}

/*
Set front and back function and reference value for stencil testing

Parameters
    fun - Specifies the test function. Eight symbolic constants are valid: Never, Less, Lequal, Greater, Gequal, Equal, Notequal, and Always. The initial value is Always.
    ref - Specifies the reference value for the stencil test. ref is clamped to the range 0 2 n - 1 , where n is the number of bitplanes in the stencil buffer. The initial value is 0.
    mask - Specifies a mask that is ANDed with both the reference value and the stored stencil value when the test is done. The initial value is all 1's.
*/
func StencilFunc(fun GLenum, ref int, mask uint32) {
	C.glStencilFunc(C.GLenum(fun), C.GLint(ref), C.GLuint(mask))
}

/*
Set front and/or back function and reference value for stencil testing

Parameters
    face - Specifies whether front and/or back stencil state is updated. Three symbolic constants are valid: Front, Back, and FrontAndBack.
    fun - Specifies the test function. Eight symbolic constants are valid: Never, Less, Lequal, Greater, Gequal, Equal, Notequal, and Always. The initial value is Always.
    ref - Specifies the reference value for the stencil test. ref is clamped to the range 0 2 n - 1 , where n is the number of bitplanes in the stencil buffer. The initial value is 0.
    mask - Specifies a mask that is ANDed with both the reference value and the stored stencil value when the test is done. The initial value is all 1's.
*/
func StencilFuncSeparate(face, fun GLenum, ref int, mask uint32) {
	C.glStencilFuncSeparate(C.GLenum(face), C.GLenum(fun), C.GLint(ref), C.GLuint(mask))
}

/*
Control the front and back writing of individual bits in the stencil planes

Parameters
    mask - Specifies a bit mask to enable and disable writing of individual bits in the stencil planes. Initially, the mask is all 1's.
*/
func StencilMask(mask uint32) {
	C.glStencilMask(C.GLuint(mask))
}

/*
Control the front and/or back writing of individual bits in the stencil planes

Parameters
    face - Specifies whether the front and/or back stencil writemask is updated. Three symbolic constants are valid: Front, Back, and FrontAndBack.
    mask - Specifies a bit mask to enable and disable writing of individual bits in the stencil planes. Initially, the mask is all 1's.
*/
func StencilMaskSeparate(face GLenum, mask uint32) {
	C.glStencilMaskSeparate(C.GLenum(face), C.GLuint(mask))
}

/*
Set front and back stencil test actions

Parameters
    sfail - Specifies the action to take when the stencil test fails. Eight symbolic constants are accepted: Keep, Zero, Replace, Incr, IncrWrap, Decr, DecrWrap, and Invert. The initial value is Keep.
    dpfail - Specifies the stencil action when the stencil test passes, but the depth test fails. dpfail accepts the same symbolic constants as sfail. The initial value is Keep.
    dppass - Specifies the stencil action when both the stencil test and the depth test pass, or when the stencil test passes and either there is no depth buffer or depth testing is not enabled. dppass accepts the same symbolic constants as sfail. The initial value is Keep.
*/
func StencilOp(sfail, dpfail, dppass GLenum) {
	C.glStencilOp(C.GLenum(sfail), C.GLenum(dpfail), C.GLenum(dppass))
}

/*
Set front and/or back stencil test actions

Parameters
    face - Specifies whether front and/or back stencil state is updated. Three symbolic constants are valid: Front, Back, and FrontAndBack.
    sfail - Specifies the action to take when the stencil test fails. Eight symbolic constants are accepted: Keep, Zero, Replace, Incr, IncrWrap, Decr, DecrWrap, and Invert. The initial value is Keep.
    dpfail - Specifies the stencil action when the stencil test passes, but the depth test fails. dpfail accepts the same symbolic constants as sfail. The initial value is Keep.
    dppass - Specifies the stencil action when both the stencil test and the depth test pass, or when the stencil test passes and either there is no depth buffer or depth testing is not enabled. dppass accepts the same symbolic constants as sfail. The initial value is Keep.
*/
func StencilOpSeparate(face, sfail, dpfail, dppass GLenum) {
	C.glStencilOpSeparate(C.GLenum(face), C.GLenum(sfail), C.GLenum(dpfail), C.GLenum(dppass))
}

/*
Attach the storage for a buffer object to the active buffer texture

Parameters
    target - Specifies the target of the operation and must be TextureBuffer.
    internalFormat - Specifies the internal format of the data in the store belonging to buffer.
    buffer - Specifies the name of the buffer object whose storage to attach to the active buffer texture.
*/
func TexBuffer(target, internalFormat GLenum, buffer uint32) {
	C.glTexBuffer(C.GLenum(target), C.GLenum(internalFormat), C.GLuint(buffer))
}

/*
Specify a one-dimensional texture image

Parameters
    target - Specifies the target texture. Must be Texture1d or ProxyTexture1d.
    level - Specifies the level-of-detail number. Level 0 is the base image level. Level n is the nth mipmap reduction image.
    internalFormat - Specifies the number of color components in the texture. Must be one of the following symbolic constants: CompressedRed, CompressedRg, CompressedRgb, CompressedRgba, CompressedSrgb, CompressedSrgbAlpha, DepthComponent, DepthComponent16, DepthComponent24, DepthComponent32, R3G3B2, Red, Rg, Rgb, Rgb4, Rgb5, Rgb8, Rgb10, Rgb12, Rgb16, Rgba, Rgba2, Rgba4, Rgb5A1, Rgba8, Rgb10A2, Rgba12, Rgba16, Srgb, Srgb8, SrgbAlpha, or Srgb8Alpha8.
    width - Specifies the width of the texture image. All implementations support texture images that are at least 1024 texels wide. The height of the 1D texture image is 1.
    border - This value must be 0.
    format - Specifies the format of the pixel data. The following symbolic values are accepted: Red, Rg, Rgb, Bgr, Rgba, and Bgra.
    data - Specifies a slice with the image data in memory.
*/
func TexImage1D(target GLenum, level int, internalFormat GLenum, width int, border int, format GLenum, data interface{}) error {
	_, _, enumType, ptr, err := sliceToGLData(data)
	if err != nil {
		return err
	}

	C.glTexImage1D(C.GLenum(target), C.GLint(level), C.GLint(internalFormat), C.GLsizei(width), C.GLint(border), C.GLenum(format), enumType, ptr)

	return nil
}

/*
Specify a two-dimensional texture image

Parameters
    target - Specifies the target texture. Must be Texture2d, ProxyTexture2d, Texture1dArray, ProxyTexture1dArray, TextureRectangle, ProxyTextureRectangle, TextureCubeMapPositiveX, TextureCubeMapNegativeX, TextureCubeMapPositiveY, TextureCubeMapNegativeY, TextureCubeMapPositiveZ, TextureCubeMapNegativeZ, or ProxyTextureCubeMap.
    level - Specifies the level-of-detail number. Level 0 is the base image level. Level n is the nth mipmap reduction image. If target is TextureRectangle or ProxyTextureRectangle, level must be 0.
    internalFormat - Specifies the number of color components in the texture. Must be one of the following symbolic constants: CompressedRed, CompressedRg, CompressedRgb, CompressedRgba, CompressedSrgb, CompressedSrgbAlpha, DepthComponent, DepthComponent16, DepthComponent24, DepthComponent32, R3G3B2, Red, Rg, Rgb, Rgb4, Rgb5, Rgb8, Rgb10, Rgb12, Rgb16, Rgba, Rgba2, Rgba4, Rgb5A1, Rgba8, Rgb10A2, Rgba12, Rgba16, Srgb, Srgb8, SrgbAlpha, or Srgb8Alpha8.
    width - Specifies the width of the texture image. All implementations support texture images that are at least 1024 texels wide.
    height - Specifies the height of the texture image, or the number of layers in a texture array, in the case of the Texture1dArray and ProxyTexture1dArray targets. All implementations support 2D texture images that are at least 1024 texels high, and texture arrays that are at least 256 layers deep.
    border - This value must be 0.
    format - Specifies the format of the pixel data. The following symbolic values are accepted: Red, Rg, Rgb, Bgr, Rgba, and Bgra.
    data - Specifies a slice with the image data in memory.
*/
func TexImage2D(target GLenum, level int, internalFormat GLenum, width, height int, border int, format GLenum, data interface{}) error {
	_, _, enumType, ptr, err := sliceToGLData(data)
	if err != nil {
		return err
	}

	C.glTexImage2D(C.GLenum(target), C.GLint(level), C.GLint(internalFormat), C.GLsizei(width), C.GLsizei(height), C.GLint(border), C.GLenum(format), enumType, ptr)

	return nil
}

/*
Establish the data storage, format, dimensions, and number of samples of a multisample texture's image

Parameters
    target - Specifies the target of the operation. target must be Texture2dMultisample or ProxyTexture2dMultisample.
    samples - The number of samples in the multisample texture's image.
    internalFormat - The internal format to be used to store the multisample texture's image. internalformat must specify a color-renderable, depth-renderable, or stencil-renderable format.
    width - The width of the multisample texture's image, in texels.
    height - The height of the multisample texture's image, in texels.
    fixedsamplelocations - Specifies whether the image will use identical sample locations and the same number of samples for all texels in the image, and the sample locations will not depend on the internal format or size of the image.
*/
func TexImage2DMultisample(target GLenum, samples int, internalFormat GLenum, width, height int, fixedsamplelocations bool) {
	C.glTexImage2DMultisample(C.GLenum(target), C.GLsizei(samples), C.GLenum(internalFormat), C.GLsizei(width), C.GLsizei(height), boolToGLBool(fixedsamplelocations))
}

/*
Specify a three-dimensional texture image

Parameters
    target - Specifies the target texture. Must be one of Texture3d, ProxyTexture3d, Texture2dArray or ProxyTexture2dArray.
    level - Specifies the level-of-detail number. Level 0 is the base image level. Level n is the nth mipmap reduction image. If target is TextureRectangle or ProxyTextureRectangle, level must be 0.
    internalFormat - Specifies the number of color components in the texture. Must be one of the following symbolic constants: Rgba32f, Rgba32i, Rgba32ui, Rgba16, Rgba16f, Rgba16i, Rgba16ui, Rgba8, Rgba8ui, Srgb8Alpha8, Rgb10A2, Rgb10A2ui, R11fG11fB10f, Rg32f, Rg32i, Rg32ui, Rg16, Rg16f, Rgb16i, Rgb16ui, Rg8, Rg8i, Rg8ui, R32f, R32i, R32ui, R16f, R16i, R16ui, R8, R8i, R8ui, Rgba16Snorm, Rgba8Snorm, Rgb32f, Rgb32i, Rgb32ui, Rgb16Snorm, Rgb16f, Rgb16i, Rgb16ui, Rgb16, Rgb8Snorm, Rgb8, Rgb8i, Rgb8ui, Srgb8, Rgb9E5, Rg16Snorm, Rg8Snorm, CompressedRgRgtc2, CompressedSignedRgRgtc2, R16Snorm, R8Snorm, CompressedRedRgtc1, CompressedSignedRedRgtc1, DepthComponent32f, DepthComponent24, DepthComponent16, Depth32fStencil8, Depth24Stencil8.
    width - Specifies the width of the texture image. All implementations support 3D texture images that are at least 16 texels wide.
    height - Specifies the height of the texture image. All implementations support 3D texture images that are at least 256 texels high.
    depth - Specifies the depth of the texture image, or the number of layers in a texture array. All implementations support 3D texture images that are at least 256 texels deep, and texture arrays that are at least 256 layers deep.
    border - This value must be 0.
    format - Specifies the format of the pixel data. The following symbolic values are accepted: Red, Rg, Rgb, Bgr, Rgba, and Bgra.
    data - Specifies a slice with the image data in memory.
*/
func TexImage3D(target GLenum, level int, internalFormat GLenum, width, height, depth int, border int, format GLenum, data interface{}) error {
	_, _, enumType, ptr, err := sliceToGLData(data)
	if err != nil {
		return err
	}

	C.glTexImage3D(C.GLenum(target), C.GLint(level), C.GLint(internalFormat), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLint(border), C.GLenum(format), enumType, ptr)

	return nil
}

/*
Establish the data storage, format, dimensions, and number of samples of a multisample texture's image

Parameters
    target - Specifies the target of the operation. target must be Texture2dMultisample or ProxyTexture2dMultisample.
    samples - The number of samples in the multisample texture's image.
    internalFormat - The internal format to be used to store the multisample texture's image. internalformat must specify a color-renderable, depth-renderable, or stencil-renderable format.
    width - The width of the multisample texture's image, in texels.
    height - The height of the multisample texture's image, in texels.
    depth - The depth of the multisample texture's image, in texels.
    fixedsamplelocations - Specifies whether the image will use identical sample locations and the same number of samples for all texels in the image, and the sample locations will not depend on the internal format or size of the image.
*/
func TexImage3DMultisample(target GLenum, samples int, internalFormat GLenum, width, height, depth int, fixedsamplelocations bool) {
	C.glTexImage3DMultisample(C.GLenum(target), C.GLsizei(samples), C.GLenum(internalFormat), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), boolToGLBool(fixedsamplelocations))
}

/*
Set texture parameters

Parameters
    target - Specifies the target texture, which must be either Texture1d, Texture2d, Texture3d, Texture1dArray, Texture2dArray, TextureRectangle, or TextureCubeMap.
    pname - Specifies the symbolic name of a single-valued texture parameter. pname can be one of the following: TextureBaseLevel, TextureCompareFunc, TextureCompareMode, TextureLodBias, TextureMinFilter, TextureMagFilter, TextureMinLod, TextureMaxLod, TextureMaxLevel, TextureSwizzleR, TextureSwizzleG, TextureSwizzleB, TextureSwizzleA, TextureWrapS, TextureWrapT, or TextureWrapR.
    param - Specifies the value of pname.
*/
func TexParameterf(target, pname GLenum, param float32) {
	C.glTexParameterf(C.GLenum(target), C.GLenum(pname), C.GLfloat(param))
}

/*
Set texture parameters

Parameters
    target - Specifies the target texture, which must be either Texture1d, Texture2d, Texture3d, Texture1dArray, Texture2dArray, TextureRectangle, or TextureCubeMap.
    pname - Specifies the symbolic name of a single-valued texture parameter. pname can be one of the following: TextureBaseLevel, TextureCompareFunc, TextureCompareMode, TextureLodBias, TextureMinFilter, TextureMagFilter, TextureMinLod, TextureMaxLod, TextureMaxLevel, TextureSwizzleR, TextureSwizzleG, TextureSwizzleB, TextureSwizzleA, TextureWrapS, TextureWrapT, or TextureWrapR.
    param - Specifies the value of pname.
*/
func TexParameteri(target, pname GLenum, param int) {
	C.glTexParameteri(C.GLenum(target), C.GLenum(pname), C.GLint(param))
}

/*
Set texture parameters

Parameters
    target - Specifies the target texture, which must be either Texture1d, Texture2d, Texture3d, Texture1dArray, Texture2dArray, TextureRectangle, or TextureCubeMap.
    pname - Specifies the symbolic name of a single-valued texture parameter. pname can be one of the following: TextureBaseLevel, TextureCompareFunc, TextureCompareMode, TextureLodBias, TextureMinFilter, TextureMagFilter, TextureMinLod, TextureMaxLod, TextureMaxLevel, TextureSwizzleR, TextureSwizzleG, TextureSwizzleB, TextureSwizzleA, TextureWrapS, TextureWrapT, or TextureWrapR.
    param - Specifies a slice where the value or values of pname are stored.
*/
func TexParametersf(target, pname GLenum, param []float32) {
	C.glTexParameterfv(C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(&param[0])))
}

/*
Set texture parameters

Parameters
    target - Specifies the target texture, which must be either Texture1d, Texture2d, Texture3d, Texture1dArray, Texture2dArray, TextureRectangle, or TextureCubeMap.
    pname - Specifies the symbolic name of a single-valued texture parameter. pname can be one of the following: TextureBaseLevel, TextureCompareFunc, TextureCompareMode, TextureLodBias, TextureMinFilter, TextureMagFilter, TextureMinLod, TextureMaxLod, TextureMaxLevel, TextureSwizzleR, TextureSwizzleG, TextureSwizzleB, TextureSwizzleA, TextureWrapS, TextureWrapT, or TextureWrapR.
    param - Specifies a slice where the value or values of pname are stored.
*/
func TexParametersi(target, pname GLenum, param []int) {
	C.glTexParameteriv(C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(&param[0])))
}

/*
Set texture parameters

Parameters
    target - Specifies the target texture, which must be either Texture1d, Texture2d, Texture3d, Texture1dArray, Texture2dArray, TextureRectangle, or TextureCubeMap.
    pname - Specifies the symbolic name of a single-valued texture parameter. pname can be one of the following: TextureBaseLevel, TextureCompareFunc, TextureCompareMode, TextureLodBias, TextureMinFilter, TextureMagFilter, TextureMinLod, TextureMaxLod, TextureMaxLevel, TextureSwizzleR, TextureSwizzleG, TextureSwizzleB, TextureSwizzleA, TextureWrapS, TextureWrapT, or TextureWrapR.
    param - Specifies a slice where the value or values of pname are stored.
*/
func TexParametersIi(target, pname GLenum, param []int) {
	C.glTexParameterIiv(C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(&param[0])))
}

/*
Set texture parameters

Parameters
    target - Specifies the target texture, which must be either Texture1d, Texture2d, Texture3d, Texture1dArray, Texture2dArray, TextureRectangle, or TextureCubeMap.
    pname - Specifies the symbolic name of a single-valued texture parameter. pname can be one of the following: TextureBaseLevel, TextureCompareFunc, TextureCompareMode, TextureLodBias, TextureMinFilter, TextureMagFilter, TextureMinLod, TextureMaxLod, TextureMaxLevel, TextureSwizzleR, TextureSwizzleG, TextureSwizzleB, TextureSwizzleA, TextureWrapS, TextureWrapT, or TextureWrapR.
    param - Specifies a slice where the value or values of pname are stored.
*/
func TexParametersIui(target, pname GLenum, param []uint32) {
	C.glTexParameterIuiv(C.GLenum(target), C.GLenum(pname), (*C.GLuint)(unsafe.Pointer(&param[0])))
}

/*
Specify a one-dimensional texture subimage

Parameters
    target - Specifies the target texture. Must be Texture1d.
    level - Specifies the level-of-detail number. Level 0 is the base image level. Level n is the nth mipmap reduction image.
    xoffset - Specifies a texel offset in the x direction within the texture array.
    width - Specifies the width of the texture subimage.
    format - Specifies the format of the pixel data. The following symbolic values are accepted: Red, Rg, Rgb, Bgr, Rgba, and Bgra.
    data - Specifies a slice with the image data in memory.
*/
func TexSubImage1D(target GLenum, level, xoffset, width int, format GLenum, data interface{}) error {
	_, _, enumType, ptr, err := sliceToGLData(data)
	if err != nil {
		return err
	}

	C.glTexSubImage1D(C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLsizei(width), C.GLenum(format), enumType, ptr)

	return nil
}

/*
Specify a two-dimensional texture subimage

Parameters
    target - Specifies the target texture. Must be Texture2d, TextureCubeMapPositiveX, TextureCubeMapNegativeX, TextureCubeMapPositiveY, TextureCubeMapNegativeY, TextureCubeMapPositiveZ, TextureCubeMapNegativeZ, or Texture1dArray.
    level - Specifies the level-of-detail number. Level 0 is the base image level. Level n is the nth mipmap reduction image.
    xoffset - Specifies a texel offset in the x direction within the texture array.
    yoffset - Specifies a texel offset in the y direction within the texture array.
    width - Specifies the width of the texture subimage.
    height - Specifies the height of the texture subimage.
    format - Specifies the format of the pixel data. The following symbolic values are accepted: Red, Rg, Rgb, Bgr, Rgba, and Bgra.
    data - Specifies a slice with the image data in memory.
*/
func TexSubImage2D(target GLenum, level, xoffset, yoffset, width, height int, format GLenum, data interface{}) error {
	_, _, enumType, ptr, err := sliceToGLData(data)
	if err != nil {
		return err
	}

	C.glTexSubImage2D(C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), enumType, ptr)

	return nil
}

/*
Specify a three-dimensional texture subimage

Parameters
    target - Specifies the target texture. Must be Texture3d or Texture2dArray.
    level - Specifies the level-of-detail number. Level 0 is the base image level. Level n is the nth mipmap reduction image.
    xoffset - Specifies a texel offset in the x direction within the texture array.
    yoffset - Specifies a texel offset in the y direction within the texture array.
    zoffset - Specifies a texel offset in the z direction within the texture array.
    width - Specifies the width of the texture subimage.
    height - Specifies the height of the texture subimage.
    depth - Specifies the depth of the texture subimage.
    format - Specifies the format of the pixel data. The following symbolic values are accepted: Red, Rg, Rgb, Bgr, Rgba, and Bgra.
    data - Specifies a slice with the image data in memory.
*/
func TexSubImage3D(target GLenum, level, xoffset, yoffset, zoffset, width, height, depth int, format GLenum, data interface{}) error {
	_, _, enumType, ptr, err := sliceToGLData(data)
	if err != nil {
		return err
	}

	C.glTexSubImage3D(C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(zoffset), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLenum(format), enumType, ptr)

	return nil
}

/*
Specify values to record in transform feedback buffers

Parameters
    program - The name of the target program object.
    varyings - A slice of strings specifying the names of the varying variables to use for transform feedback.
    bufferMode - Identifies the mode used to capture the varying variables when transform feedback is active. bufferMode must be InterleavedAttribs or SeparateAttribs.
*/
func TransformFeedbackVaryings(program uint, varyings []string, bufferMode GLenum) {
	count := len(varyings)
	str := make([]*C.char, count)

	for key, _ := range varyings {
		str[key] = C.CString(varyings[key])
	}

	C.glTransformFeedbackVaryings(C.GLuint(program), C.GLsizei(count), (**C.GLchar)(unsafe.Pointer(&str[0])), C.GLenum(bufferMode))

	for key, _ := range str {
		C.free(unsafe.Pointer(str[key]))
	}
}

/*
Specify the value of a or several uniform variable(s) for the current program object

Parameters
    location - Specifies the location of the uniform variable to be modified.
    v0, v1, v2, v3 - Specifies the new values to be used for the specified uniform variable.
*/
func Uniform1f(location int, v0 float32) {
	C.glUniform1f(C.GLint(location), C.GLfloat(v0))
}

/*
Specify the value of a or several uniform variable(s) for the current program object

Parameters
    location - Specifies the location of the uniform variable to be modified.
    v0, v1, v2, v3 - Specifies the new values to be used for the specified uniform variable.
*/
func Uniform2f(location int, v0, v1 float32) {
	C.glUniform2f(C.GLint(location), C.GLfloat(v0), C.GLfloat(v1))
}

/*
Specify the value of a or several uniform variable(s) for the current program object

Parameters
    location - Specifies the location of the uniform variable to be modified.
    v0, v1, v2, v3 - Specifies the new values to be used for the specified uniform variable.
*/
func Uniform3f(location int, v0, v1, v2 float32) {
	C.glUniform3f(C.GLint(location), C.GLfloat(v0), C.GLfloat(v1), C.GLfloat(v2))
}

/*
Specify the value of a or several uniform variable(s) for the current program object

Parameters
    location - Specifies the location of the uniform variable to be modified.
    v0, v1, v2, v3 - Specifies the new values to be used for the specified uniform variable.
*/
func Uniform4f(location int, v0, v1, v2, v3 float32) {
	C.glUniform4f(C.GLint(location), C.GLfloat(v0), C.GLfloat(v1), C.GLfloat(v2), C.GLfloat(v3))
}

/*
Specify the value of a or several uniform variable(s) for the current program object

Parameters
    location - Specifies the location of the uniform variable to be modified.
    v0, v1, v2, v3 - Specifies the new values to be used for the specified uniform variable.
*/
func Uniform1i(location int, v0 int) {
	C.glUniform1i(C.GLint(location), C.GLint(v0))
}

/*
Specify the value of a or several uniform variable(s) for the current program object

Parameters
    location - Specifies the location of the uniform variable to be modified.
    v0, v1, v2, v3 - Specifies the new values to be used for the specified uniform variable.
*/
func Uniform2i(location int, v0, v1 int) {
	C.glUniform2i(C.GLint(location), C.GLint(v0), C.GLint(v1))
}

/*
Specify the value of a or several uniform variable(s) for the current program object

Parameters
    location - Specifies the location of the uniform variable to be modified.
    v0, v1, v2, v3 - Specifies the new values to be used for the specified uniform variable.
*/
func Uniform3i(location int, v0, v1, v2 int) {
	C.glUniform3i(C.GLint(location), C.GLint(v0), C.GLint(v1), C.GLint(v2))
}

/*
Specify the value of a or several uniform variable(s) for the current program object

Parameters
    location - Specifies the location of the uniform variable to be modified.
    v0, v1, v2, v3 - Specifies the new values to be used for the specified uniform variable.
*/
func Uniform4i(location int, v0, v1, v2, v3 int) {
	C.glUniform4i(C.GLint(location), C.GLint(v0), C.GLint(v1), C.GLint(v2), C.GLint(v3))
}

/*
Specify the value of a or several uniform variable(s) for the current program object

Parameters
    location - Specifies the location of the uniform variable to be modified.
    v0, v1, v2, v3 - Specifies the new values to be used for the specified uniform variable.
*/
func Uniform1ui(location int, v0 uint32) {
	C.glUniform1ui(C.GLint(location), C.GLuint(v0))
}

/*
Specify the value of a or several uniform variable(s) for the current program object

Parameters
    location - Specifies the location of the uniform variable to be modified.
    v0, v1, v2, v3 - Specifies the new values to be used for the specified uniform variable.
*/
func Uniform2ui(location int, v0, v1 uint32) {
	C.glUniform2ui(C.GLint(location), C.GLuint(v0), C.GLuint(v1))
}

/*
Specify the value of a or several uniform variable(s) for the current program object

Parameters
    location - Specifies the location of the uniform variable to be modified.
    v0, v1, v2, v3 - Specifies the new values to be used for the specified uniform variable.
*/
func Uniform3ui(location int, v0, v1, v2 uint32) {
	C.glUniform3ui(C.GLint(location), C.GLuint(v0), C.GLuint(v1), C.GLuint(v2))
}

/*
Specify the value of a or several uniform variable(s) for the current program object

Parameters
    location - Specifies the location of the uniform variable to be modified.
    v0, v1, v2, v3 - Specifies the new values to be used for the specified uniform variable.
*/
func Uniform4ui(location int, v0, v1, v2, v3 uint32) {
	C.glUniform4ui(C.GLint(location), C.GLuint(v0), C.GLuint(v1), C.GLuint(v2), C.GLuint(v3))
}

/*
Specify the value of a or several uniform variable(s) for the current program object

Parameters
    location - Specifies the location of the uniform variable to be modified.
    count - Specifies the number of elements that are to be modified. This should be 1 if the targeted uniform variable is not an array, and 1 or more if it is an array.
    values - Specifies a slice of count values that will be used to update the specified uniform variable.
*/
func Uniforms1f(location, count int, values []float32) {
	C.glUniform1fv(C.GLint(location), C.GLsizei(count), (*C.GLfloat)(unsafe.Pointer(&values[0])))
}

/*
Specify the value of a or several uniform variable(s) for the current program object

Parameters
    location - Specifies the location of the uniform variable to be modified.
    count - Specifies the number of elements that are to be modified. This should be 1 if the targeted uniform variable is not an array, and 1 or more if it is an array.
    values - Specifies a slice of count values that will be used to update the specified uniform variable.
*/
func Uniforms2f(location, count int, values []float32) {
	C.glUniform2fv(C.GLint(location), C.GLsizei(count), (*C.GLfloat)(unsafe.Pointer(&values[0])))
}

/*
Specify the value of a or several uniform variable(s) for the current program object

Parameters
    location - Specifies the location of the uniform variable to be modified.
    count - Specifies the number of elements that are to be modified. This should be 1 if the targeted uniform variable is not an array, and 1 or more if it is an array.
    values - Specifies a slice of count values that will be used to update the specified uniform variable.
*/
func Uniforms3f(location, count int, values []float32) {
	C.glUniform3fv(C.GLint(location), C.GLsizei(count), (*C.GLfloat)(unsafe.Pointer(&values[0])))
}

/*
Specify the value of a or several uniform variable(s) for the current program object

Parameters
    location - Specifies the location of the uniform variable to be modified.
    count - Specifies the number of elements that are to be modified. This should be 1 if the targeted uniform variable is not an array, and 1 or more if it is an array.
    values - Specifies a slice of count values that will be used to update the specified uniform variable.
*/
func Uniforms4f(location, count int, values []float32) {
	C.glUniform4fv(C.GLint(location), C.GLsizei(count), (*C.GLfloat)(unsafe.Pointer(&values[0])))
}

/*
Specify the value of a or several uniform variable(s) for the current program object

Parameters
    location - Specifies the location of the uniform variable to be modified.
    count - Specifies the number of elements that are to be modified. This should be 1 if the targeted uniform variable is not an array, and 1 or more if it is an array.
    values - Specifies a slice of count values that will be used to update the specified uniform variable.
*/
func Uniforms1i(location, count int, values []int) {
	C.glUniform1iv(C.GLint(location), C.GLsizei(count), (*C.GLint)(unsafe.Pointer(&values[0])))
}

/*
Specify the value of a or several uniform variable(s) for the current program object

Parameters
    location - Specifies the location of the uniform variable to be modified.
    count - Specifies the number of elements that are to be modified. This should be 1 if the targeted uniform variable is not an array, and 1 or more if it is an array.
    values - Specifies a slice of count values that will be used to update the specified uniform variable.
*/
func Uniforms2i(location, count int, values []int) {
	C.glUniform2iv(C.GLint(location), C.GLsizei(count), (*C.GLint)(unsafe.Pointer(&values[0])))
}

/*
Specify the value of a or several uniform variable(s) for the current program object

Parameters
    location - Specifies the location of the uniform variable to be modified.
    count - Specifies the number of elements that are to be modified. This should be 1 if the targeted uniform variable is not an array, and 1 or more if it is an array.
    values - Specifies a slice of count values that will be used to update the specified uniform variable.
*/
func Uniforms3i(location, count int, values []int) {
	C.glUniform3iv(C.GLint(location), C.GLsizei(count), (*C.GLint)(unsafe.Pointer(&values[0])))
}

/*
Specify the value of a or several uniform variable(s) for the current program object

Parameters
    location - Specifies the location of the uniform variable to be modified.
    count - Specifies the number of elements that are to be modified. This should be 1 if the targeted uniform variable is not an array, and 1 or more if it is an array.
    values - Specifies a slice of count values that will be used to update the specified uniform variable.
*/
func Uniforms4i(location, count int, values []int) {
	C.glUniform4iv(C.GLint(location), C.GLsizei(count), (*C.GLint)(unsafe.Pointer(&values[0])))
}

/*
Specify the value of a or several uniform variable(s) for the current program object

Parameters
    location - Specifies the location of the uniform variable to be modified.
    count - Specifies the number of elements that are to be modified. This should be 1 if the targeted uniform variable is not an array, and 1 or more if it is an array.
    values - Specifies a slice of count values that will be used to update the specified uniform variable.
*/
func Uniforms1ui(location, count int, values []uint32) {
	C.glUniform1uiv(C.GLint(location), C.GLsizei(count), (*C.GLuint)(unsafe.Pointer(&values[0])))
}

/*
Specify the value of a or several uniform variable(s) for the current program object

Parameters
    location - Specifies the location of the uniform variable to be modified.
    count - Specifies the number of elements that are to be modified. This should be 1 if the targeted uniform variable is not an array, and 1 or more if it is an array.
    values - Specifies a slice of count values that will be used to update the specified uniform variable.
*/
func Uniforms2ui(location, count int, values []uint32) {
	C.glUniform2uiv(C.GLint(location), C.GLsizei(count), (*C.GLuint)(unsafe.Pointer(&values[0])))
}

/*
Specify the value of a or several uniform variable(s) for the current program object

Parameters
    location - Specifies the location of the uniform variable to be modified.
    count - Specifies the number of elements that are to be modified. This should be 1 if the targeted uniform variable is not an array, and 1 or more if it is an array.
    values - Specifies a slice of count values that will be used to update the specified uniform variable.
*/
func Uniforms3ui(location, count int, values []uint32) {
	C.glUniform3uiv(C.GLint(location), C.GLsizei(count), (*C.GLuint)(unsafe.Pointer(&values[0])))
}

/*
Specify the value of a or several uniform variable(s) for the current program object

Parameters
    location - Specifies the location of the uniform variable to be modified.
    count - Specifies the number of elements that are to be modified. This should be 1 if the targeted uniform variable is not an array, and 1 or more if it is an array.
    values - Specifies a slice of count values that will be used to update the specified uniform variable.
*/
func Uniforms4ui(location, count int, values []uint32) {
	C.glUniform4uiv(C.GLint(location), C.GLsizei(count), (*C.GLuint)(unsafe.Pointer(&values[0])))
}

/*
Specify the value of a or several uniform variable(s) for the current program object

Parameters
    location - Specifies the location of the uniform variable to be modified.
    count - Specifies the number of matrices that are to be modified. This should be 1 if the targeted uniform variable is not an array of matrices, and 1 or more if it is an array of matrices.
    transpose - Specifies whether to transpose the matrix as the values are loaded into the uniform variable.
    values - Specifies a slice of count values that will be used to update the specified uniform variable.
*/
func UniformMatrix2f(location, count int, transpose bool, values []float32) {
	C.glUniformMatrix2fv(C.GLint(location), C.GLsizei(count), boolToGLBool(transpose), (*C.GLfloat)(unsafe.Pointer(&values[0])))
}

/*
Specify the value of a or several uniform variable(s) for the current program object

Parameters
    location - Specifies the location of the uniform variable to be modified.
    count - Specifies the number of matrices that are to be modified. This should be 1 if the targeted uniform variable is not an array of matrices, and 1 or more if it is an array of matrices.
    transpose - Specifies whether to transpose the matrix as the values are loaded into the uniform variable.
    values - Specifies a slice of count values that will be used to update the specified uniform variable.
*/
func UniformMatrix3f(location, count int, transpose bool, values []float32) {
	C.glUniformMatrix3fv(C.GLint(location), C.GLsizei(count), boolToGLBool(transpose), (*C.GLfloat)(unsafe.Pointer(&values[0])))
}

/*
Specify the value of a or several uniform variable(s) for the current program object

Parameters
    location - Specifies the location of the uniform variable to be modified.
    count - Specifies the number of matrices that are to be modified. This should be 1 if the targeted uniform variable is not an array of matrices, and 1 or more if it is an array of matrices.
    transpose - Specifies whether to transpose the matrix as the values are loaded into the uniform variable.
    values - Specifies a slice of count values that will be used to update the specified uniform variable.
*/
func UniformMatrix4f(location, count int, transpose bool, values []float32) {
	C.glUniformMatrix4fv(C.GLint(location), C.GLsizei(count), boolToGLBool(transpose), (*C.GLfloat)(unsafe.Pointer(&values[0])))
}

/*
Specify the value of a or several uniform variable(s) for the current program object

Parameters
    location - Specifies the location of the uniform variable to be modified.
    count - Specifies the number of matrices that are to be modified. This should be 1 if the targeted uniform variable is not an array of matrices, and 1 or more if it is an array of matrices.
    transpose - Specifies whether to transpose the matrix as the values are loaded into the uniform variable.
    values - Specifies a slice of count values that will be used to update the specified uniform variable.
*/
func UniformMatrix2x3f(location, count int, transpose bool, values []float32) {
	C.glUniformMatrix2x3fv(C.GLint(location), C.GLsizei(count), boolToGLBool(transpose), (*C.GLfloat)(unsafe.Pointer(&values[0])))
}

/*
Specify the value of a or several uniform variable(s) for the current program object

Parameters
    location - Specifies the location of the uniform variable to be modified.
    count - Specifies the number of matrices that are to be modified. This should be 1 if the targeted uniform variable is not an array of matrices, and 1 or more if it is an array of matrices.
    transpose - Specifies whether to transpose the matrix as the values are loaded into the uniform variable.
    values - Specifies a slice of count values that will be used to update the specified uniform variable.
*/
func UniformMatrix3x2f(location, count int, transpose bool, values []float32) {
	C.glUniformMatrix3x2fv(C.GLint(location), C.GLsizei(count), boolToGLBool(transpose), (*C.GLfloat)(unsafe.Pointer(&values[0])))
}

/*
Specify the value of a or several uniform variable(s) for the current program object

Parameters
    location - Specifies the location of the uniform variable to be modified.
    count - Specifies the number of matrices that are to be modified. This should be 1 if the targeted uniform variable is not an array of matrices, and 1 or more if it is an array of matrices.
    transpose - Specifies whether to transpose the matrix as the values are loaded into the uniform variable.
    values - Specifies a slice of count values that will be used to update the specified uniform variable.
*/
func UniformMatrix2x4f(location, count int, transpose bool, values []float32) {
	C.glUniformMatrix2x4fv(C.GLint(location), C.GLsizei(count), boolToGLBool(transpose), (*C.GLfloat)(unsafe.Pointer(&values[0])))
}

/*
Specify the value of a or several uniform variable(s) for the current program object

Parameters
    location - Specifies the location of the uniform variable to be modified.
    count - Specifies the number of matrices that are to be modified. This should be 1 if the targeted uniform variable is not an array of matrices, and 1 or more if it is an array of matrices.
    transpose - Specifies whether to transpose the matrix as the values are loaded into the uniform variable.
    values - Specifies a slice of count values that will be used to update the specified uniform variable.
*/
func UniformMatrix4x2f(location, count int, transpose bool, values []float32) {
	C.glUniformMatrix4x2fv(C.GLint(location), C.GLsizei(count), boolToGLBool(transpose), (*C.GLfloat)(unsafe.Pointer(&values[0])))
}

/*
Specify the value of a or several uniform variable(s) for the current program object

Parameters
    location - Specifies the location of the uniform variable to be modified.
    count - Specifies the number of matrices that are to be modified. This should be 1 if the targeted uniform variable is not an array of matrices, and 1 or more if it is an array of matrices.
    transpose - Specifies whether to transpose the matrix as the values are loaded into the uniform variable.
    values - Specifies a slice of count values that will be used to update the specified uniform variable.
*/
func UniformMatrix3x4f(location, count int, transpose bool, values []float32) {
	C.glUniformMatrix3x4fv(C.GLint(location), C.GLsizei(count), boolToGLBool(transpose), (*C.GLfloat)(unsafe.Pointer(&values[0])))
}

/*
Specify the value of a or several uniform variable(s) for the current program object

Parameters
    location - Specifies the location of the uniform variable to be modified.
    count - Specifies the number of matrices that are to be modified. This should be 1 if the targeted uniform variable is not an array of matrices, and 1 or more if it is an array of matrices.
    transpose - Specifies whether to transpose the matrix as the values are loaded into the uniform variable.
    values - Specifies a slice of count values that will be used to update the specified uniform variable.
*/
func UniformMatrix4x3f(location, count int, transpose bool, values []float32) {
	C.glUniformMatrix4x3fv(C.GLint(location), C.GLsizei(count), boolToGLBool(transpose), (*C.GLfloat)(unsafe.Pointer(&values[0])))
}

/*
Assign a binding point to an active uniform block

Parameters
    program - The name of a program object containing the active uniform block whose binding to assign.
    uniformBlockIndex - The index of the active uniform block within program whose binding to assign.
    uniformBlockBinding - Specifies the binding point to which to bind the uniform block with index uniformBlockIndex within program.
*/
func UniformBlockBinding(program, uniformBlockIndex, uniformBlockBinding uint32) {
	C.glUniformBlockBinding(C.GLuint(program), C.GLuint(uniformBlockIndex), C.GLuint(uniformBlockBinding))
}

/*
Unmap a buffer object's data store

Parameters
    target - Specifies the target buffer object being unmapped. The symbolic constant must be ArrayBuffer, CopyReadBuffer, CopyWriteBuffer, ElementArrayBuffer, PixelPackBuffer, PixelUnpackBuffer, TextureBuffer, TransformFeedbackBuffer or UniformBuffer.
*/
func UnmapBuffer(target GLenum) bool {
	return glBoolToBool(C.glUnmapBuffer(C.GLenum(target)))
}

/*
Installs a program object as part of current rendering state

Parameters
    program - Specifies the handle of the program object whose executables are to be used as part of current rendering state.
*/
func UseProgram(program uint32) {
	C.glUseProgram(C.GLuint(program))
}

/*
Validates a program object

Parameters
    program - Specifies the handle of the program object to be validated.
*/
func ValidateProgram(program uint32) {
	C.glValidateProgram(C.GLuint(program))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v0 - Specifies the new value to be used for the specified vertex attribute.
*/
func VertexAttrib1f(index uint32, v0 float32) {
	C.glVertexAttrib1f(C.GLuint(index), C.GLfloat(v0))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v0 - Specifies the new value to be used for the specified vertex attribute.
*/
func VertexAttrib1s(index uint32, v0 float32) {
	C.glVertexAttrib1s(C.GLuint(index), C.GLshort(v0))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v0 - Specifies the new value to be used for the specified vertex attribute.
*/
func VertexAttrib1d(index uint32, v0 float64) {
	C.glVertexAttrib1d(C.GLuint(index), C.GLdouble(v0))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v0 - Specifies the new value to be used for the specified vertex attribute.
*/
func VertexAttribI1i(index uint32, v0 int) {
	C.glVertexAttribI1i(C.GLuint(index), C.GLint(v0))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v0 - Specifies the new value to be used for the specified vertex attribute.
*/
func VertexAttribI1ui(index uint32, v0 uint32) {
	C.glVertexAttribI1ui(C.GLuint(index), C.GLuint(v0))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v0, v1 - Specifies the new values to be used for the specified vertex attribute.
*/
func VertexAttrib2f(index uint32, v0, v1 float32) {
	C.glVertexAttrib2f(C.GLuint(index), C.GLfloat(v0), C.GLfloat(v1))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v0, v1 - Specifies the new values to be used for the specified vertex attribute.
*/
func VertexAttrib2s(index uint32, v0, v1 float32) {
	C.glVertexAttrib2s(C.GLuint(index), C.GLshort(v0), C.GLshort(v1))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v0, v1 - Specifies the new values to be used for the specified vertex attribute.
*/
func VertexAttrib2d(index uint32, v0, v1 float64) {
	C.glVertexAttrib2d(C.GLuint(index), C.GLdouble(v0), C.GLdouble(v1))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v0, v1 - Specifies the new values to be used for the specified vertex attribute.
*/
func VertexAttribI2i(index uint32, v0, v1 int) {
	C.glVertexAttribI2i(C.GLuint(index), C.GLint(v0), C.GLint(v1))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v0, v1 - Specifies the new values to be used for the specified vertex attribute.
*/
func VertexAttribI2ui(index uint32, v0, v1 uint32) {
	C.glVertexAttribI2ui(C.GLuint(index), C.GLuint(v0), C.GLuint(v1))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v0, v1, v2 - Specifies the new values to be used for the specified vertex attribute.
*/
func VertexAttrib3f(index uint32, v0, v1, v2 float32) {
	C.glVertexAttrib3f(C.GLuint(index), C.GLfloat(v0), C.GLfloat(v1), C.GLfloat(v2))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v0, v1, v2 - Specifies the new values to be used for the specified vertex attribute.
*/
func VertexAttrib3s(index uint32, v0, v1, v2 float32) {
	C.glVertexAttrib3s(C.GLuint(index), C.GLshort(v0), C.GLshort(v1), C.GLshort(v2))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v0, v1, v2 - Specifies the new values to be used for the specified vertex attribute.
*/
func VertexAttrib3d(index uint32, v0, v1, v2 float64) {
	C.glVertexAttrib3d(C.GLuint(index), C.GLdouble(v0), C.GLdouble(v1), C.GLdouble(v2))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v0, v1, v2 - Specifies the new values to be used for the specified vertex attribute.
*/
func VertexAttribI3i(index uint32, v0, v1, v2 int) {
	C.glVertexAttribI3i(C.GLuint(index), C.GLint(v0), C.GLint(v1), C.GLint(v2))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v0, v1, v2 - Specifies the new values to be used for the specified vertex attribute.
*/
func VertexAttribI3ui(index uint32, v0, v1, v2 uint32) {
	C.glVertexAttribI3ui(C.GLuint(index), C.GLuint(v0), C.GLuint(v1), C.GLuint(v2))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v0, v1, v2, v3 - Specifies the new values to be used for the specified vertex attribute.
*/
func VertexAttrib4f(index uint32, v0, v1, v2, v3 float32) {
	C.glVertexAttrib4f(C.GLuint(index), C.GLfloat(v0), C.GLfloat(v1), C.GLfloat(v2), C.GLfloat(v3))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v0, v1, v2, v3 - Specifies the new values to be used for the specified vertex attribute.
*/
func VertexAttrib4s(index uint32, v0, v1, v2, v3 float32) {
	C.glVertexAttrib4s(C.GLuint(index), C.GLshort(v0), C.GLshort(v1), C.GLshort(v2), C.GLshort(v3))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v0, v1, v2, v3 - Specifies the new values to be used for the specified vertex attribute.
*/
func VertexAttrib4d(index uint32, v0, v1, v2, v3 float64) {
	C.glVertexAttrib4d(C.GLuint(index), C.GLdouble(v0), C.GLdouble(v1), C.GLdouble(v2), C.GLdouble(v3))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v0, v1, v2, v3 - Specifies the new values to be used for the specified vertex attribute.
*/
func VertexAttrib4Nub(index uint32, v0, v1, v2, v3 byte) {
	C.glVertexAttrib4Nub(C.GLuint(index), C.GLubyte(v0), C.GLubyte(v1), C.GLubyte(v2), C.GLubyte(v3))
}

/*q
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v0, v1, v2, v3 - Specifies the new values to be used for the specified vertex attribute.
*/
func VertexAttribI4i(index uint32, v0, v1, v2, v3 int) {
	C.glVertexAttribI4i(C.GLuint(index), C.GLint(v0), C.GLint(v1), C.GLint(v2), C.GLint(v3))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v0, v1, v2, v3 - Specifies the new values to be used for the specified vertex attribute.
*/
func VertexAttribI4ui(index uint32, v0, v1, v2, v3 uint32) {
	C.glVertexAttribI4ui(C.GLuint(index), C.GLuint(v0), C.GLuint(v1), C.GLuint(v2), C.GLuint(v3))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v - Specifies a slice of values to be used for the generic vertex attribute.
*/
func VertexAttribs1f(index uint32, v []float32) {
	C.glVertexAttrib1fv(C.GLuint(index), (*C.GLfloat)(unsafe.Pointer(&v[0])))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v - Specifies a slice of values to be used for the generic vertex attribute.
*/
func VertexAttribs1s(index uint32, v []float32) {
	C.glVertexAttrib1sv(C.GLuint(index), (*C.GLshort)(unsafe.Pointer(&v[0])))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v - Specifies a slice of values to be used for the generic vertex attribute.
*/
func VertexAttribs1d(index uint32, v []float64) {
	C.glVertexAttrib1dv(C.GLuint(index), (*C.GLdouble)(unsafe.Pointer(&v[0])))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v - Specifies a slice of values to be used for the generic vertex attribute.
*/
func VertexAttribsI1i(index uint32, v []int) {
	C.glVertexAttribI1iv(C.GLuint(index), (*C.GLint)(unsafe.Pointer(&v[0])))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v - Specifies a slice of values to be used for the generic vertex attribute.
*/
func VertexAttribsI1ui(index uint32, v []uint32) {
	C.glVertexAttribI1uiv(C.GLuint(index), (*C.GLuint)(unsafe.Pointer(&v[0])))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v - Specifies a slice of values to be used for the generic vertex attribute.
*/
func VertexAttribs2f(index uint32, v []float32) {
	C.glVertexAttrib2fv(C.GLuint(index), (*C.GLfloat)(unsafe.Pointer(&v[0])))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v - Specifies a slice of values to be used for the generic vertex attribute.
*/
func VertexAttribs2s(index uint32, v []float32) {
	C.glVertexAttrib2sv(C.GLuint(index), (*C.GLshort)(unsafe.Pointer(&v[0])))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v - Specifies a slice of values to be used for the generic vertex attribute.
*/
func VertexAttribs2d(index uint32, v []float64) {
	C.glVertexAttrib2dv(C.GLuint(index), (*C.GLdouble)(unsafe.Pointer(&v[0])))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v - Specifies a slice of values to be used for the generic vertex attribute.
*/
func VertexAttribsI2i(index uint32, v []int) {
	C.glVertexAttribI2iv(C.GLuint(index), (*C.GLint)(unsafe.Pointer(&v[0])))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v - Specifies a slice of values to be used for the generic vertex attribute.
*/
func VertexAttribsI2ui(index uint32, v []uint32) {
	C.glVertexAttribI2uiv(C.GLuint(index), (*C.GLuint)(unsafe.Pointer(&v[0])))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v - Specifies a slice of values to be used for the generic vertex attribute.
*/
func VertexAttribs3f(index uint32, v []float32) {
	C.glVertexAttrib3fv(C.GLuint(index), (*C.GLfloat)(unsafe.Pointer(&v[0])))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v - Specifies a slice of values to be used for the generic vertex attribute.
*/
func VertexAttribs3s(index uint32, v []float32) {
	C.glVertexAttrib3sv(C.GLuint(index), (*C.GLshort)(unsafe.Pointer(&v[0])))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v - Specifies a slice of values to be used for the generic vertex attribute.
*/
func VertexAttribs3d(index uint32, v []float64) {
	C.glVertexAttrib3dv(C.GLuint(index), (*C.GLdouble)(unsafe.Pointer(&v[0])))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v - Specifies a slice of values to be used for the generic vertex attribute.
*/
func VertexAttribsI3i(index uint32, v []int) {
	C.glVertexAttribI3iv(C.GLuint(index), (*C.GLint)(unsafe.Pointer(&v[0])))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v - Specifies a slice of values to be used for the generic vertex attribute.
*/
func VertexAttribsI3ui(index uint32, v []uint32) {
	C.glVertexAttribI3uiv(C.GLuint(index), (*C.GLuint)(unsafe.Pointer(&v[0])))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v - Specifies a slice of values to be used for the generic vertex attribute.
*/
func VertexAttribs4f(index uint32, v []float32) {
	C.glVertexAttrib4fv(C.GLuint(index), (*C.GLfloat)(unsafe.Pointer(&v[0])))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v - Specifies a slice of values to be used for the generic vertex attribute.
*/
func VertexAttribs4s(index uint32, v []float32) {
	C.glVertexAttrib4sv(C.GLuint(index), (*C.GLshort)(unsafe.Pointer(&v[0])))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v - Specifies a slice of values to be used for the generic vertex attribute.
*/
func VertexAttribs4d(index uint32, v []float64) {
	C.glVertexAttrib4dv(C.GLuint(index), (*C.GLdouble)(unsafe.Pointer(&v[0])))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v - Specifies a slice of values to be used for the generic vertex attribute.
*/
func VertexAttribs4i(index uint32, v []int) {
	C.glVertexAttrib4iv(C.GLuint(index), (*C.GLint)(unsafe.Pointer(&v[0])))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v - Specifies a slice of values to be used for the generic vertex attribute.
*/
func VertexAttribs4b(index uint32, v []byte) {
	C.glVertexAttrib4bv(C.GLuint(index), (*C.GLbyte)(unsafe.Pointer(&v[0])))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v - Specifies a slice of values to be used for the generic vertex attribute.
*/
func VertexAttribs4ub(index uint32, v []byte) {
	C.glVertexAttrib4ubv(C.GLuint(index), (*C.GLubyte)(unsafe.Pointer(&v[0])))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v - Specifies a slice of values to be used for the generic vertex attribute.
*/
func VertexAttribs4us(index uint32, v []float32) {
	C.glVertexAttrib4usv(C.GLuint(index), (*C.GLushort)(unsafe.Pointer(&v[0])))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v - Specifies a slice of values to be used for the generic vertex attribute.
*/
func VertexAttribs4ui(index uint32, v []uint32) {
	C.glVertexAttrib4uiv(C.GLuint(index), (*C.GLuint)(unsafe.Pointer(&v[0])))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v - Specifies a slice of values to be used for the generic vertex attribute.
*/
func VertexAttribs4Nb(index uint32, v []byte) {
	C.glVertexAttrib4Nbv(C.GLuint(index), (*C.GLbyte)(unsafe.Pointer(&v[0])))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v - Specifies a slice of values to be used for the generic vertex attribute.
*/
func VertexAttribs4Ns(index uint32, v []float32) {
	C.glVertexAttrib4Nsv(C.GLuint(index), (*C.GLshort)(unsafe.Pointer(&v[0])))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v - Specifies a slice of values to be used for the generic vertex attribute.
*/
func VertexAttribs4Ni(index uint32, v []int) {
	C.glVertexAttrib4Niv(C.GLuint(index), (*C.GLint)(unsafe.Pointer(&v[0])))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v - Specifies a slice of values to be used for the generic vertex attribute.
*/
func VertexAttribs4Nub(index uint32, v []byte) {
	C.glVertexAttrib4Nubv(C.GLuint(index), (*C.GLubyte)(unsafe.Pointer(&v[0])))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v - Specifies a slice of values to be used for the generic vertex attribute.
*/
func VertexAttribs4Nus(index uint32, v []float32) {
	C.glVertexAttrib4Nusv(C.GLuint(index), (*C.GLushort)(unsafe.Pointer(&v[0])))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v - Specifies a slice of values to be used for the generic vertex attribute.
*/
func VertexAttribs4Nui(index uint32, v []uint32) {
	C.glVertexAttrib4Nuiv(C.GLuint(index), (*C.GLuint)(unsafe.Pointer(&v[0])))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v - Specifies a slice of values to be used for the generic vertex attribute.
*/
func VertexAttribsI4b(index uint32, v []byte) {
	C.glVertexAttribI4bv(C.GLuint(index), (*C.GLbyte)(unsafe.Pointer(&v[0])))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v - Specifies a slice of values to be used for the generic vertex attribute.
*/
func VertexAttribsI4ub(index uint32, v []byte) {
	C.glVertexAttribI4ubv(C.GLuint(index), (*C.GLubyte)(unsafe.Pointer(&v[0])))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v - Specifies a slice of values to be used for the generic vertex attribute.
*/
func VertexAttribsI4s(index uint32, v []float32) {
	C.glVertexAttribI4sv(C.GLuint(index), (*C.GLshort)(unsafe.Pointer(&v[0])))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v - Specifies a slice of values to be used for the generic vertex attribute.
*/
func VertexAttribsI4us(index uint32, v []float32) {
	C.glVertexAttribI4usv(C.GLuint(index), (*C.GLushort)(unsafe.Pointer(&v[0])))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v - Specifies a slice of values to be used for the generic vertex attribute.
*/
func VertexAttribsI4i(index uint32, v []int) {
	C.glVertexAttribI4iv(C.GLuint(index), (*C.GLint)(unsafe.Pointer(&v[0])))
}

/*
Specifies the value of a generic vertex attribute

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    v - Specifies a slice of values to be used for the generic vertex attribute.
*/
func VertexAttribsI4ui(index uint32, v []uint32) {
	C.glVertexAttribI4uiv(C.GLuint(index), (*C.GLuint)(unsafe.Pointer(&v[0])))
}

/*
Define an array of generic vertex attribute data

Parameters
    index - Specifies the index of the generic vertex attribute to be modified.
    size - Specifies the number of components per generic vertex attribute. Must be 1, 2, 3, 4. Additionally, the symbolic constant Bgra is accepted by glVertexAttribPointer. The initial value is 4.
    enumType - Specifies the data type of each component in the array. The symbolic constants Byte, UnsignedByte, Short, UnsignedShort, Int, and UnsignedInt are accepted by both functions. Additionally HalfFloat, Float, Double, Int2101010Rev, and UnsignedInt2101010Rev are accepted by glVertexAttribPointer. The initial value is Float.
    normalized - For glVertexAttribPointer, specifies whether fixed-point data values should be normalized (true) or converted directly as fixed-point values (false) when they are accessed.
    stride - Specifies the byte offset between consecutive generic vertex attributes. If stride is 0, the generic vertex attributes are understood to be tightly packed in the array. The initial value is 0.
    offset - Specifies a offset of the first component of the first generic vertex attribute in the array in the data store of the buffer currently bound to the ArrayBuffer target. The initial value is 0.
*/
func VertexAttribPointer(index uint32, size int, enumType GLenum, normalized bool, stride int, offset int) {
	C.glVertexAttribPointer(C.GLuint(index), C.GLint(size), C.GLenum(enumType), boolToGLBool(normalized), C.GLsizei(stride), unsafe.Pointer(uintptr(offset)))
}

/*
Set the viewport

Parameters
    x, y - Specify the lower left corner of the viewport rectangle, in pixels. The initial value is (0, 0).
    width, height - Specify the width and height of the viewport. When a GL context is first attached to a window, width and height are set to the dimensions of that window.
*/
func Viewport(x, y, width, height int) {
	C.glViewport(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}

/*
Instruct the GL server to block until the specified sync object becomes signaled

Parameters
    sync - Specifies the sync object whose status to wait on.
    flags - A bitfield controlling the command flushing behavior. flags may be zero.
    timeout - Specifies the timeout that the server should wait before continuing. timeout must be TimeoutIgnored.
*/
func WaitSync(sync GLsync, flags GLbitfield, timeout uint64) {
	C.glWaitSync(C.GLsync(sync), C.GLbitfield(flags), C.GLuint64(timeout))
}
