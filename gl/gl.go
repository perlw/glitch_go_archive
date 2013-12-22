package gl

// #cgo linux   LDFLAGS: -lGLEW -lGL
// #cgo windows LDFLAGS: -lglew32 -lopengl32
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
Bind a user-defined varying out variable to a fragment shader color number

Parameters
    program - The name of the program containing varying out variable whose binding to modify.
    colorNumber - The color number to bind the user-defined varying out variable to.
    name - The name of the user-defined varying out variable whose binding to modify.
*/
func BindFragDataLocation(program, colorNumber uint32, name string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	C.glBindFragDataLocation(C.GLuint(program), C.GLuint(colorNumber), (*C.GLchar)(cName))
}

/*
Bind a user-defined varying out variable to a fragment shader color number and index

Parameters
    program - The name of the program containing varying out variable whose binding to modify.
    colorNumber - The color number to bind the user-defined varying out variable to.
    index - The index of the color input to bind the user-defined varying out variable to.
    name - The name of the user-defined varying out variable whose binding to modify.
*/
func BindFragDataLocationIndexed(program, colorNumber, index uint32, name string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	C.glBindFragDataLocationIndexed(C.GLuint(program), C.GLuint(colorNumber), C.GLuint(index), (*C.GLchar)(cName))
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
    renderbuffer - Specifies the name of the renderbuffer object to bind.
*/
func BindRenderbuffer(renderbuffer uint32) {
	C.glBindRenderbuffer(C.GL_RENDERBUFFER, C.GLuint(renderbuffer))
}

/*
Bind a named sampler to a texturing target

Parameters
    unit - Specifies the index of the texture unit to which the sampler is bound.
    sampler - Specifies the name of a sampler.
*/
func BindSampler(unit, sampler uint32) {
	C.glBindSampler(C.GLuint(unit), C.GLuint(sampler))
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
    clamp - Specifies whether to apply color clamping. clamp must be True or False.
*/
func ClampColor(clamp GLenum) {
	C.glClampColor(C.GL_CLAMP_READ_COLOR, C.GLenum(clamp))
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
    timeout - The timeout, specified in nanoseconds, for which the implementation should wait for sync to become signaled.
*/
func ClientWaitSync(sync GLsync, timeout uint64) GLenum {
	return GLenum(C.glClientWaitSync(C.GLsync(sync), C.GL_SYNC_FLUSH_COMMANDS_BIT, C.GLuint64(timeout)))
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
func CullFace(mode GLConstant) {
	C.glCullFace(C.GLenum(mode))
}

// <-------- THIS FAR --------->

/*
Specify the value used for depth buffer comparisons.

Parameters
    constFunc - Specifies the depth comparison function. Symbolic constants Never, Less, Equal, LEqual, Greater, NotEqual, GEqual, and Always are accepted. The initial value is Less.
*/
func DepthFunc(constFunc GLConstant) {
	C.glDepthFunc(C.GLenum(constFunc))
}

/*
Enable server-side GL capabilities.

Parameters
    capability - Specifies a symbolic constant indicating a GL capability.
*/
func Enable(capability GLConstant) {
	C.glEnable(C.GLenum(capability))
}

/*
Delimit the vertices of a primitive or a group of like primitives.
*/
func End() {
	C.glEnd()
}

/*
Disable server-side GL capabilities.

Parameters
    capability - Specifies a symbolic constant indicating a GL capability.
*/
func Disable(capability GLConstant) {
	C.glDisable(C.GLenum(capability))
}

/*
Set the viewport.

Parameters
    x, y - Specify the lower left corner of the viewport rectangle, in pixels. The initial value is (0, 0).
    width, height - Specify the width and height of the viewport. When a GL context is first attached to a window, width and height are set to the dimensions of that window.
*/
func Viewport(x, y, width, height int) {
	C.glViewport(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}
