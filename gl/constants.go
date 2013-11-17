package gl

// #include "common.h"
import "C"

type GLConstant int

const (
	// Accum
	AccumC GLConstant = C.GL_ACCUM
	Load              = C.GL_LOAD
	Add               = C.GL_ADD
	Mult              = C.GL_MULT
	Return            = C.GL_RETURN

	// Begin
	Points        = C.GL_POINTS
	Lines         = C.GL_LINES
	LineStrip     = C.GL_LINE_STRIP
	LineLoop      = C.GL_LINE_LOOP
	Triangles     = C.GL_TRIANGLES
	TriangleStrip = C.GL_TRIANGLE_STRIP
	TriangleFan   = C.GL_TRIANGLE_FAN
	Quads         = C.GL_QUADS
	QuadStrip     = C.GL_QUAD_STRIP
	Polygon       = C.GL_POLYGON

	// BindTexture
	Texture1D = C.GL_TEXTURE_1D
	Texture2D = C.GL_TEXTURE_2D

	// BlendFunc
	Zero             = C.GL_ZERO
	One              = C.GL_ONE
	DstColor         = C.GL_DST_COLOR
	OneMinusDstColor = C.GL_ONE_MINUS_DST_COLOR
	SrcAlpha         = C.GL_SRC_ALPHA
	OneMinusSrcAlpha = C.GL_ONE_MINUS_SRC_ALPHA
	DstAlpha         = C.GL_DST_ALPHA
	OneMinusDstAlpha = C.GL_ONE_MINUS_DST_ALPHA
	SrcAlphaSaturate = C.GL_SRC_ALPHA_SATURATE

	// DepthFunc, AlphaFunc
	Never    = C.GL_NEVER
	Less     = C.GL_LESS
	Equal    = C.GL_EQUAL
	LEqual   = C.GL_LEQUAL
	Greater  = C.GL_GREATER
	NotEqual = C.GL_NOTEQUAL
	GEqual   = C.GL_EQUAL
	Always   = C.GL_ALWAYS

	// Clear
	ColorBufferBit   = C.GL_COLOR_BUFFER_BIT
	DepthBufferBit   = C.GL_DEPTH_BUFFER_BIT
	StencilBufferBit = C.GL_STENCIL_BUFFER_BIT

	// ClipPlane
	ClipPlane0    = C.GL_CLIP_PLANE0
	ClipPlane1    = C.GL_CLIP_PLANE1
	ClipPlane2    = C.GL_CLIP_PLANE2
	ClipPlane3    = C.GL_CLIP_PLANE3
	ClipPlane4    = C.GL_CLIP_PLANE4
	ClipPlane5    = C.GL_CLIP_PLANE5
	MaxClipPlanes = C.GL_MAX_CLIP_PLANES

	// ColorMaterial, CullFace
	Front             = C.GL_FRONT
	Back              = C.GL_BACK
	FrontAndBack      = C.GL_FRONT_AND_BACK
	Emission          = C.GL_EMISSION
	Ambient           = C.GL_AMBIENT
	Diffuse           = C.GL_DIFFUSE
	Specular          = C.GL_SPECULAR
	AmbientAndDiffuse = C.GL_AMBIENT_AND_DIFFUSE

	// CopyPixels
	Color   = C.GL_COLOR
	Depth   = C.GL_DEPTH
	Stencil = C.GL_STENCIL

	// CopyTexImage
	Alpha              = C.GL_ALPHA
	Alpha4             = C.GL_ALPHA4
	Alpha8             = C.GL_ALPHA8
	Alpha12            = C.GL_ALPHA12
	Alpha16            = C.GL_ALPHA16
	Luminance          = C.GL_LUMINANCE
	Luminance4         = C.GL_LUMINANCE4
	Luminance8         = C.GL_LUMINANCE8
	Luminance12        = C.GL_LUMINANCE12
	Luminance16        = C.GL_LUMINANCE16
	LuminanceAlpha     = C.GL_LUMINANCE_ALPHA
	Luminance4Alpha4   = C.GL_LUMINANCE4_ALPHA4
	Luminance6Alpha2   = C.GL_LUMINANCE6_ALPHA2
	Luminance8Alpha8   = C.GL_LUMINANCE8_ALPHA8
	Luminance12Alpha4  = C.GL_LUMINANCE12_ALPHA4
	Luminance12Alpha12 = C.GL_LUMINANCE12_ALPHA12
	Luminance16Alpha16 = C.GL_LUMINANCE16_ALPHA16
	Intensity          = C.GL_INTENSITY
	Intensity4         = C.GL_INTENSITY4
	Intensity8         = C.GL_INTENSITY8
	Intensity12        = C.GL_INTENSITY12
	Intensity16        = C.GL_INTENSITY16
	Rgb                = C.GL_RGB
	R3G3B2             = C.GL_R3_G3_B2
	Rgb4               = C.GL_RGB4
	Rgb5               = C.GL_RGB5
	Rgb8               = C.GL_RGB8
	Rgb10              = C.GL_RGB10
	Rgb12              = C.GL_RGB12
	Rgb16              = C.GL_RGB16
	Rgba               = C.GL_RGBA
	Rgba2              = C.GL_RGBA2
	Rgba4              = C.GL_RGBA4
	Rgb5A1             = C.GL_RGB5_A1
	Rgba8              = C.GL_RGBA8
	Rgb10A2            = C.GL_RGB10_A2
	Rgba12             = C.GL_RGBA12
	Rgba16             = C.GL_RGBA16

	// Enable
	Blend = C.GL_BLEND
	//ClipDistance                      = GL_CLIP_DISTANCE
	ColorLogicOp           = C.GL_COLOR_LOGIC_OP
	CullFaceC              = C.GL_CULL_FACE
	DepthClamp             = C.GL_DEPTH_CLAMP
	DepthTest              = C.GL_DEPTH_TEST
	Dither                 = C.GL_DITHER
	FrameBufferSRGB        = C.GL_FRAMEBUFFER_SRGB
	LineSmooth             = C.GL_LINE_SMOOTH
	Multisample            = C.GL_MULTISAMPLE
	PolygonOffsetFill      = C.GL_POLYGON_OFFSET_FILL
	PolygonOffsetLine      = C.GL_POLYGON_OFFSET_LINE
	PolygonOffsetPoint     = C.GL_POLYGON_OFFSET_POINT
	PolygonSmooth          = C.GL_POLYGON_SMOOTH
	PrimitiveRestart       = C.GL_PRIMITIVE_RESTART
	RasterizerDiscard      = C.GL_RASTERIZER_DISCARD
	SampleAlphaToCoverage  = C.GL_SAMPLE_ALPHA_TO_COVERAGE
	SampleAlphaToOne       = C.GL_SAMPLE_ALPHA_TO_ONE
	SampleCoverage         = C.GL_SAMPLE_COVERAGE
	SampleShading          = C.GL_SAMPLE_SHADING
	SampleMask             = C.GL_SAMPLE_MASK
	ScissorTest            = C.GL_SCISSOR_TEST
	StencilTest            = C.GL_STENCIL_TEST
	TextureCubeMapSeamless = C.GL_TEXTURE_CUBE_MAP_SEAMLESS
	ProgramPointSIze       = C.GL_PROGRAM_POINT_SIZE
)
