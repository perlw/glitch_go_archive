package gl

// #cgo linux   LDFLAGS: -lGLEW -lGL
// #cgo windows LDFLAGS: -lglew32 -lopengl32
// #include "common.h"
import "C"

import (
	"errors"
	"unsafe"
)

func Init() {
	C.glewExperimental = 1
	C.glewInit()
}

/*
Select active texture unit

Parameters
    texture - Specifies which texture unit to make active. The number of texture units is implementation dependent, but must be at least 48. texture must be one of TEXTUREi, where i ranges from 0 (MAX_COMBINED_TEXTURE_IMAGE_UNITS - 1). The initial value is TEXTURE0.
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
    target - Specifies the target type of query object established between BeginQuery and the subsequent EndQuery. The symbolic constant must be one of SAMPLES_PASSED, ANY_SAMPLES_PASSED, PRIMITIVES_GENERATED, TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN, or TIME_ELAPSED.
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
    target - Specifies the target to which the buffer object is bound. The symbolic constant must be ARRAY_BUFFER, COPY_READ_BUFFER, COPY_WRITE_BUFFER, ELEMENT_ARRAY_BUFFER, PIXEL_PACK_BUFFER, PIXEL_UNPACK_BUFFER, TEXTURE_BUFFER, TRANSFORM_FEEDBACK_BUFFER, or UNIFORM_BUFFER.
    buffer - Specifies the name of a buffer object.
*/
func BindBuffer(target GLenum, buffer uint32) {
	C.glBindBuffer(C.GLenum(target), C.GLuint(buffer))
}

/*
Bind a buffer object to an indexed buffer target

Parameters
    target - Specify the target of the bind operation. target must be either TRANSFORM_FEEDBACK_BUFFER or UNIFORM_BUFFER.
    index - Specify the index of the binding point within the array specified by target.
    buffer - The name of a buffer object to bind to the specified binding point.
*/
func BindBufferBase(target GLenum, index, buffer uint32) {
	C.glBindBufferBase(C.GLenum(target), C.GLuint(index), C.GLuint(buffer))
}

/*
Bind a range within a buffer object to an indexed buffer target

Parameters
    target - Specify the target of the bind operation. target must be either TRANSFORM_FEEDBACK_BUFFER or UNIFORM_BUFFER.
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
	C.glBindFragDataLocationIndexed(C.GLuint(program), C.GLuint(colorNumber), C.GLuint(index), cName)
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
    target - Specifies the target to which the texture is bound. Must be either TEXTURE_1D, TEXTURE_2D, TEXTURE_3D, or TEXTURE_1D_ARRAY, TEXTURE_2D_ARRAY, TEXTURE_RECTANGLE, TEXTURE_CUBE_MAP, TEXTURE_BUFFER, TEXTURE_2D_MULTISAMPLE or TEXTURE_2D_MULTISAMPLE_ARRAY.
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
    mode - Specifies how source and destination colors are combined. It must be FUNC_ADD, FUNC_SUBTRACT, FUNC_REVERSE_SUBTRACT, MIN, MAX.
*/
func BlendEquation(mode GLenum) {
	C.glBlendEquation(C.GLenum(mode))
}

/*
Set the RGB blend equation and the alpha blend equation separately

Parameters
    modeRGB - Specifies the RGB blend equation, how the red, green, and blue components of the source and destination colors are combined. It must be FUNC_ADD, FUNC_SUBTRACT, FUNC_REVERSE_SUBTRACT, MIN, MAX.
    modeAlpha - Specifies the alpha blend equation, how the alpha component of the source and destination colors are combined. It must be FUNC_ADD, FUNC_SUBTRACT, FUNC_REVERSE_SUBTRACT, MIN, MAX.
*/
func BlendEquationSeparate(modeRGB, modeAlpha GLenum) {
	C.glBlendEquationSeparate(C.GLenum(modeRGB), C.GLenum(modeAlpha))
}

/*
Specify pixel arithmetic.

Parameters
    sfactor - Specifies how the red, green, blue, and alpha source blending factors are computed. The initial value is ONE.
    dfactor - Specifies how the red, green, blue, and alpha destination blending factors are computed. The following symbolic constants are accepted: ZERO, ONE, SRC_COLOR, ONE_MINUS_SRC_COLOR, DST_COLOR, ONE_MINUS_DST_COLOR, SRC_ALPHA, ONE_MINUS_SRC_ALPHA, DST_ALPHA, ONE_MINUS_DST_ALPHA. CONSTANT_COLOR, ONE_MINUS_CONSTANT_COLOR, CONSTANT_ALPHA, and ONE_MINUS_CONSTANT_ALPHA. The initial value is ZERO.
*/
func BlendFunc(sfactor, dfactor GLenum) {
	C.glBlendFunc(C.GLenum(sfactor), C.GLenum(dfactor))
}

/*
Specify pixel arithmetic for RGB and alpha components separately

Parameters
    srcRGB - Specifies how the red, green, and blue blending factors are computed. The initial value is ONE.
    dstRGB - Specifies how the red, green, and blue destination blending factors are computed. The initial value is ZERO.
    srcAlpha - Specified how the alpha source blending factor is computed. The initial value is ONE.
    dstAlpha - Specified how the alpha destination blending factor is computed. The initial value is ZERO.
*/
func BlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha GLenum) {
	C.glBlendFuncSeparate(C.GLenum(srcRGB), C.GLenum(dstRGB), C.GLenum(srcAlpha), C.GLenum(dstAlpha))
}

/*
Copy a block of pixels from the read framebuffer to the draw framebuffer

Parameters
    srcX0, srcY0, srcX1, srcY1 - Specify the bounds of the source rectangle within the read buffer of the read framebuffer.
    dstX0, dstY0, dstX1, dstY1 - Specify the bounds of the destination rectangle within the write buffer of the write framebuffer.
    mask - The bitwise OR of the flags indicating which buffers are to be copied. The allowed flags are COLOR_BUFFER_BIT, DEPTH_BUFFER_BIT and STENCIL_BUFFER_BIT.
    filter - Specifies the interpolation to be applied if the image is stretched. Must be NEAREST or LINEAR.
*/
func BlitFramebuffer(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1 int, mask GLbitfield, filter GLenum) {
	C.glBlitFramebuffer(C.GLint(srcX0), C.GLint(srcY0), C.GLint(srcX1), C.GLint(srcY1), C.GLint(dstX0), C.GLint(dstY0), C.GLint(dstX1), C.GLint(dstY1), C.GLbitfield(mask), C.GLenum(filter))
}

/*
Creates and initializes a buffer object's data store

Parameters
    target - Specifies the target buffer object. The symbolic constant must be ARRAY_BUFFER, COPY_READ_BUFFER, COPY_WRITE_BUFFER, ELEMENT_ARRAY_BUFFER, PIXEL_PACK_BUFFER, PIXEL_UNPACK_BUFFER, TEXTURE_BUFFER, TRANSFORM_FEEDBACK_BUFFER, or UNIFORM_BUFFER.
    data - Specifies a slice of the data that will be copied into the data store for initialization, or nil if not data is to be copied.
    usage - Specifies the expected usage pattern of the data store. The symbolic constant must be STREAM_DRAW, STREAM_READ, STREAM_COPY, STATIC_DRAW, STATIC_READ, STATIC_COPY, DYNAMIC_DRAW, DYNAMIC_READ, or DYNAMIC_COPY.
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
    target - Specifies the target buffer object. The symbolic constant must be ARRAY_BUFFER, COPY_READ_BUFFER, COPY_WRITE_BUFFER, ELEMENT_ARRAY_BUFFER, PIXEL_PACK_BUFFER, PIXEL_UNPACK_BUFFER, TEXTURE_BUFFER, TRANSFORM_FEEDBACK_BUFFER, or UNIFORM_BUFFER.
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
Execute a display list.

Parameters
    list - Specifies the integer name of the display list to be executed.
*/
func CallList(list uint32) {
	C.glCallList(C.GLuint(list))
}

/*
Execute a list of display lists.

Parameters
    lists - Specifies a slice of name offsets in the display list
*/
func CallLists(lists []uint32) error {
	num, _, enum, ptr, err := sliceToGLData(lists)
	if err != nil {
		return err
	}

	C.glCallLists(num, enum, ptr)
	return nil
}

/*
Clear buffers to preset values.

Parameters
    mask - Bitwise OR of masks that indicate the buffers to be cleared. The four masks are ColorBufferBit, DepthBufferBit, AccumBufferBit, and StencilBufferBit.
*/
func Clear(mask GLConstant) {
	C.glClear(C.GLbitfield(mask))
}

/*
Specify clear values for the accumulation buffer.

Parameters
    red, green, blue, alpha - Specify the red, green, blue, and alpha values used when the accumulation buffers are cleared. The initial values are all 0.
*/
func ClearAccum(red, green, blue, alpha float32) {
	C.glClearAccum(C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue), C.GLfloat(alpha))
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
Specify the clear value for the color index buffers.

Parameters
    c - Specifies the index used when the color index buffers are cleared. The initial value is 0.
*/
func ClearIndex(c float32) {
	C.glClearIndex(C.GLfloat(c))
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
Specify a plane against which all geometry is clipped.

Parameters
    plane - Specifies which clipping plane is being positioned. Symbolic names of the form ClipPlanei, where i is an integer between 0 and MaxClipPlanes - 1, are accepted.
    equation - Specifies a slice of floating-point values. These values are interpreted as a plane equation.
*/
func ClipPlane(plane GLConstant, equation []float64) {
	C.glClipPlane(C.GLenum(plane), (*C.GLdouble)(&equation[0]))
}

/*
Set the current color.

Parameters
    red, green, blue - The red, green, and blue channels.
*/
func Color3b(red, green, blue byte) {
	C.glColor3b(C.GLbyte(red), C.GLbyte(green), C.GLbyte(blue))
}

/*
Set the current color.

Parameters
    red, green, blue - The red, green, and blue channels.
*/
func Color3d(red, green, blue float64) {
	C.glColor3d(C.GLdouble(red), C.GLdouble(green), C.GLdouble(blue))
}

/*
Set the current color.

Parameters
    red, green, blue - The red, green, and blue channels.
*/
func Color3f(red, green, blue float32) {
	C.glColor3f(C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue))
}

/*
Set the current color.

Parameters
    red, green, blue - The red, green, and blue channels.
*/
func Color3i(red, green, blue int32) {
	C.glColor3i(C.GLint(red), C.GLint(green), C.GLint(blue))
}

/*
Set the current color.

Parameters
    red, green, blue - The red, green, and blue channels.
*/
func Color3s(red, green, blue int16) {
	C.glColor3s(C.GLshort(red), C.GLshort(green), C.GLshort(blue))
}

/*
Set the current color.

Parameters
    red, green, blue - The red, green, and blue channels.
*/
func Color3ub(red, green, blue uint8) {
	C.glColor3ub(C.GLubyte(red), C.GLubyte(green), C.GLubyte(blue))
}

/*
Set the current color.

Parameters
    red, green, blue - The red, green, and blue channels.
*/
func Color3ui(red, green, blue uint32) {
	C.glColor3ui(C.GLuint(red), C.GLuint(green), C.GLuint(blue))
}

/*
Set the current color.

Parameters
    red, green, blue - The red, green, and blue channels.
*/
func Color3us(red, green, blue uint16) {
	C.glColor3us(C.GLushort(red), C.GLushort(green), C.GLushort(blue))
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
Cause a material color to track the current color.

Parameters
    face - Specifies whether front, back, or both front and back material parameters should track the current color. Accepted values are Front, Back, and FrontAndBack. The initial value is FrontAndBack.
    mode - Specifies which of several material parameters track the current color. Accepted values are Emission, Ambient, Diffuse, Specular, and AmbientAndDiffuse. The initial value is AmbientAndDiffuse.
*/
func ColorMaterial(face, mode GLConstant) {
	C.glColorMaterial(C.GLenum(face), C.GLenum(mode))
}

/*
Define an array of colors.

Parameters
    size - Specifies the number of components per color. Must be 3 or 4.
    stride - Specifies the byte offset between consecutive colors. If stride is 0, (the initial value), the colors are understood to be tightly packed in the slice.
    colors - A slice containing the color values.
*/
func ColorPointer(size, stride int, colors interface{}) error {
	if size != 3 || size != 4 {
		return errors.New("gl: size must be 3 or 4")
	}

	_, _, enum, ptr, err := sliceToGLData(colors)
	if err != nil {
		return err
	}

	C.glColorPointer(C.GLint(size), enum, C.GLsizei(stride), ptr)

	return nil
}

/*
Copy pixels in the frame buffer.

Parameters
    x, y - Specify the window coordinates of the lower left corner of the rectangular region of pixels to be copied.
    width, height - Specify the dimensions of the rectangular region of pixels to be copied. Both must be nonnegative.
    typeConstant - Specifies whether color values, depth values, or stencil values are to be copied. Symbolic constants Color, Depth, and Stencil are accepted.
*/
func CopyPixels(x, y, width, height int, typeConstant GLConstant) {
	C.glCopyPixels(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLenum(typeConstant))
}

/*
Copy pixels into a 1D texture image.

Parameters
    level - Specifies the level-of-detail number. Level 0 is the base image level. Level n is the nth mipmap reduction image.
    internalFormat - Specifies the internal format of the texture. Must be one of the following symbolic constants: Alpha, Alpha4, Alpha8, Alpha12, Alpha16, Luminance, Luminance4, Luminance8, Luminance12, Luminance16, LuminanceAlpha, Luminance4Alpha4, Luminance6Alpha2, Luminance8Alpha8, Luminance12Alpha4, Luminance12Alpha12, Luminance16Alpha16, Intensity, Intensity4, Intensity8, Intensity12, Intensity16, Rgb, R3G3B2, Rgb4, Rgb5, Rgb8, Rgb10, Rgb12, Rgb16, Rgba, Rgba2, Rgba4, Rgb5A1, Rgba8, Rgb10A2, Rgba12, or Rgba16.
    x, y - Specify the window coordinates of the left corner of the row of pixels to be copied.
    width - Specifies the width of the texture image. Must be 0 or 2n + 2xborder for some integer n. The height of the texture image is 1.
    border - Specifies the width of the border. Must be either 0 or 1.
*/
func CopyTexImage1D(level int, internalFormat GLConstant, x, y, width, border int) {
	C.glCopyTexImage1D(C.GL_TEXTURE_1D, C.GLint(level), C.GLenum(internalFormat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLint(border))
}

/*
Copy pixels into a 2D texture image.

Parameters
    level - Specifies the level-of-detail number. Level 0 is the base image level. Level n is the nth mipmap reduction image.
    internalFormat - Specifies the internal format of the texture. Must be one of the following symbolic constants: Alpha, Alpha4, Alpha8, Alpha12, Alpha16, Luminance, Luminance4, Luminance8, Luminance12, Luminance16, LuminanceAlpha, Luminance4Alpha4, Luminance6Alpha2, Luminance8Alpha8, Luminance12Alpha4, Luminance12Alpha12, Luminance16Alpha16, Intensity, Intensity4, Intensity8, Intensity12, Intensity16, Rgb, R3G3B2, Rgb4, Rgb5, Rgb8, Rgb10, Rgb12, Rgb16, Rgba, Rgba2, Rgba4, Rgb5A1, Rgba8, Rgb10A2, Rgba12, or Rgba16.
    x, y - Specify the window coordinates of the left corner of the row of pixels to be copied.
    width, height - Specifies the width and height of the texture image. Must be 0 or 2n + 2xborder for some integer n. The height of the texture image is 1.
    border - Specifies the width of the border. Must be either 0 or 1.
*/
func CopyTexImage2D(level int, internalFormat GLConstant, x, y, width, height, border int) {
	C.glCopyTexImage2D(C.GL_TEXTURE_2D, C.GLint(level), C.GLenum(internalFormat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLint(border))
}

/*
Copy a one-dimensional texture subimage.

Parameters
    level - Specifies the level-of-detail number. Level 0 is the base image level. Level n is the nth mipmap reduction image.
    xoffset - Specifies the texel offset within the texture array.
    x, y - Specify the window coordinates of the left corner of the row of pixels to be copied.
    width - Specifies the width of the texture subimage
*/
func CopyTexSubImage1D(level, xoffset, x, y, width int) {
	C.glCopyTexSubImage1D(C.GL_TEXTURE_1D, C.GLint(level), C.GLint(xoffset), C.GLint(x), C.GLint(y), C.GLsizei(width))
}

/*
Copy a two-dimensional texture subimage.

Parameters
    level - Specifies the level-of-detail number. Level 0 is the base image level. Level n is the nth mipmap reduction image.
    xoffset, yoffset - Specifies the texel offset within the texture array.
    x, y - Specify the window coordinates of the left corner of the row of pixels to be copied.
    width, height - Specifies the width and height of the texture subimage
*/
func CopyTexSubImage2D(level, xoffset, yoffset, x, y, width, height int) {
	C.glCopyTexSubImage2D(C.GL_TEXTURE_2D, C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}

/*
Specify whether front- or back-facing facets can be culled.

Parameters
    mode - Specifies whether front- or back-facing facets are candidates for culling. Symbolic constants Front, Back, and FrontAndBack are accepted. The initial value is Back.
*/
func CullFace(mode GLConstant) {
	C.glCullFace(C.GLenum(mode))
}

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
