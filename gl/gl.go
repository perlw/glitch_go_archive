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
Clear buffers to preset values.

Parameters
    mask - Bitwise OR of masks that indicate the buffers to be cleared. The four masks are ColorBufferBit, DepthBufferBit, AccumBufferBit, and StencilBufferBit.
*/
func Clear(mask GLConstant) {
	C.glClear(C.GLbitfield(mask))
}

/*
Specify clear values for the color buffers.

Parameters
    red, green, blue, alpha - Specify the red, green, blue, and alpha values used when the color buffers are cleared. The initial values are all 0.
*/
func ClearColor(r, g, b, a float32) {
	C.glClearColor(C.GLclampf(r), C.GLclampf(g), C.GLclampf(b), C.GLclampf(a))
}

/*
Specify the clear value for the depth buffer.

Parameters
    depth Specifies the depth value used when the depth buffer is cleared. The initial value is 1.
*/
func ClearDepth(depth float32) {
	C.glClearDepth(C.GLclampd(depth))
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
