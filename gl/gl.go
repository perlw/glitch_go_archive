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
Operate on the accumulation buffer.

Parameters
    op - Specifies the accumulation buffer operation. Symbolic constants Accum, Load, Add, Mult, and Return are accepted.
    value - Specifies a floating-point value used in the accumulation buffer operation. op determines how value is used.
*/
func Accum(op GLConstant, value float32) {
	C.glAccum(C.GLenum(op), C.GLfloat(value))
}

/*
Specify the alpha test function.

Parameters
    constFunc - Specifies the alpha comparison function. Symbolic constants Never, Less, Equal, LEqual, Greater, NotEqual, GEqual, and Always are accepted. The initial value is Always.
    ref - Specifies the reference value that incoming alpha values are compared to. This value is clamped to the range 0 through 1, where 0 represents the lowest possible alpha value and 1 the highest possible value. The initial reference value is 0.
*/
func AlphaFunc(constFunc GLConstant, ref float32) {
	C.glAlphaFunc(C.GLenum(constFunc), C.GLclampf(ref))
}

/*
Determine if textures are loaded in texture memory.

Parameters
    textures - A slice of texture ids to be queried.

Returns
    A slice of bools representing if each texture is resident.
*/
func AreTexturesResident(textures []uint32) ([]bool, error) {
	size := len(textures)
	if size == 0 {
		return nil, errors.New("gl: Empty textures")
	}
	residences := make([]bool, size)

	if ret := C.glAreTexturesResident(C.GLsizei(size), (*C.GLuint)(unsafe.Pointer(&textures[0])), (*C.GLboolean)(unsafe.Pointer(&residences[0]))); ret == C.GL_TRUE {
		for t := range residences {
			residences[t] = true
		}
	}
	return residences, nil
}

/*
Render a vertex using the specified vertex array element.

Parameters
    index - Specifies an index into the enabled vertex data arrays.
*/
func ArrayElement(index int) {
	C.glArrayElement(C.GLint(index))
}

/*
Delimit the vertices of a primitive or a group of like primitives.

Parameters
    mode - Specifies the primitive or primitives that will be created from vertices presented between Begin and the subsequent End. Ten symbolic constants are accepted: Points, Lines, LineStrip, LineLoop, Triangles, TriangleStrip, TriangleFan, Quads, QuadStrip, and Polygon.
*/
func Begin(mode GLConstant) {
	C.glBegin(C.GLenum(mode))
}

/*
Bind a named texture to a texture target.

Parameters
    target - Specifies the target to which the texture is bound. Must be either Texture1D or Texture2D. Initially, both Texture1D and Texture2D are bound to texture 0.
    texture - Specifies the name of a texture.
*/
func BindTexture(target GLConstant, texture uint32) {
	C.glBindTexture(C.GLenum(target), C.GLuint(texture))
}

/*
Draw a bitmap.

Parameters
    width, height - Specify the pixel width and height of the bitmap image.
    xorig, yorig - Specify the location of the origin in the bitmap image. The origin is measured from the lower left corner of the bitmap, with right and up being the positive axes.
    xmove, ymove - Specify the x and y offsets to be added to the current raster position after the bitmap is drawn.
    bitmap - A slice of raw bitmap data.
*/
func Bitmap(width, height int, xorig, yorig, xmove, ymove float32, bitmap []byte) error {
	size := len(bitmap)
	if size == 0 {
		return errors.New("gl: Empty bitmap")
	}

	C.glBitmap(C.GLsizei(width), C.GLsizei(height), C.GLfloat(xorig), C.GLfloat(yorig), C.GLfloat(xmove), C.GLfloat(ymove), (*C.GLubyte)(unsafe.Pointer(&bitmap[0])))

	return nil
}

/*
Set the blend color.

Parameters
    red, green, blue, alpha - Specify the components of BlendColorExt.
*/
func BlendColorEXT(red, green, blue, alpha float32) {
	C.glBlendColorEXT(C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha))
}

/*
Specify pixel arithmetic.

Parameters
    sfactor - Specifies how the red, green, blue, and alpha source blending factors are computed. Nine symbolic constants are accepted: Zero, One, DstColor, OneMinusDstColor, SrcAlpha, OneMinusSrcAlpha, DstAlpha, OneMinusDstAlpha, and SrcAlphaSaturate. The initial value is One.
    dfactor - Specifies how the red, green, blue, and alpha destination blending factors are computed. Eight symbolic constants are accepted: Zero, One, OneMinusDstColor, SrcAlpha, OneMinusSrcAlpha, DstAlpha, OneMinusDstAlpha, and SrcAlphaSaturate. The initial value is Zero.
*/
func BlendFunc(sfactor, dfactor GLConstant) {
	C.glBlendFunc(C.GLenum(sfactor), C.GLenum(dfactor))
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

func ColorPointer() {

}

func CopyPixels() {

}

func CopyTexImage1D() {

}

func CopyTexImage2D() {

}

func CopyTexSubImage1D() {

}

func CopyTexSubImage2D() {

}

func CullFace() {

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
