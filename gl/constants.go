package gl

// #include "common.h"
import "C"

type GLenum int

const (
	// ActiveTexture
	TEXTURE0  GLenum = C.GL_TEXTURE0
	TEXTURE1         = C.GL_TEXTURE1
	TEXTURE2         = C.GL_TEXTURE2
	TEXTURE3         = C.GL_TEXTURE3
	TEXTURE4         = C.GL_TEXTURE4
	TEXTURE5         = C.GL_TEXTURE5
	TEXTURE6         = C.GL_TEXTURE6
	TEXTURE7         = C.GL_TEXTURE7
	TEXTURE8         = C.GL_TEXTURE8
	TEXTURE9         = C.GL_TEXTURE9
	TEXTURE10        = C.GL_TEXTURE10
	TEXTURE11        = C.GL_TEXTURE11
	TEXTURE12        = C.GL_TEXTURE12
	TEXTURE13        = C.GL_TEXTURE13
	TEXTURE14        = C.GL_TEXTURE14
	TEXTURE15        = C.GL_TEXTURE15
	TEXTURE16        = C.GL_TEXTURE16
	TEXTURE17        = C.GL_TEXTURE17
	TEXTURE18        = C.GL_TEXTURE18
	TEXTURE19        = C.GL_TEXTURE19
	TEXTURE20        = C.GL_TEXTURE20
	TEXTURE21        = C.GL_TEXTURE21
	TEXTURE22        = C.GL_TEXTURE22
	TEXTURE23        = C.GL_TEXTURE23
	TEXTURE24        = C.GL_TEXTURE24
	TEXTURE25        = C.GL_TEXTURE25
	TEXTURE26        = C.GL_TEXTURE26
	TEXTURE27        = C.GL_TEXTURE27
	TEXTURE28        = C.GL_TEXTURE28
	TEXTURE29        = C.GL_TEXTURE29
	TEXTURE30        = C.GL_TEXTURE30
	TEXTURE31        = C.GL_TEXTURE31

	// BeginConditionalRender
	QUERY_BY_REGION_WAIT    = C.GL_QUERY_BY_REGION_WAIT
	QUERY_BY_REGION_NO_WAIT = C.GL_QUERY_BY_REGION_NO_WAIT

	// BeginQuery
	SAMPLES_PASSED                        = C.GL_SAMPLES_PASSED
	ANY_SAMPLES_PASSED                    = C.GL_ANY_SAMPLES_PASSED
	PRIMITICES_GENERATED                  = C.GL_PRIMITIVES_GENERATED
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN = C.GL_TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN
	TIME_ELAPSED                          = C.GL_TIME_ELAPSED

	// BeginTransformFeedback
	POINTS                   = C.GL_POINTS
	LINES                    = C.GL_LINES
	LINE_LOOP                = C.GL_LINE_LOOP
	LINE_STRIP               = C.GL_LINE_STRIP
	LINES_ADJACENCY          = C.GL_LINES_ADJACENCY
	LINE_STRIP_ADJACENCY     = C.GL_LINE_STRIP_ADJACENCY
	TRIANGLES                = C.GL_TRIANGLES
	TRIANGLE_STRIP           = C.GL_TRIANGLE_STRIP
	TRIANGLE_FAN             = C.GL_TRIANGLE_FAN
	TRIANGLES_ADJACENCY      = C.GL_TRIANGLES_ADJACENCY
	TRIANGLE_STRIP_ADJACENCY = C.GL_TRIANGLE_STRIP_ADJACENCY

	// BindBuffer, BufferData
	ARRAY_BUFFER         = C.GL_ARRAY_BUFFER
	COPY_READ_BUFFER     = C.GL_COPY_READ_BUFFER
	COPY_WRITE_BUFFER    = C.GL_COPY_WRITE_BUFFER
	ELEMENT_ARRAY_BUFFER = C.GL_ELEMENT_ARRAY_BUFFER
	PIXEL_PACK_BUFFER    = C.GL_PIXEL_PACK_BUFFER
	PIXEL_UNPACK_BUFFER  = C.GL_PIXEL_UNPACK_BUFFER
	// BindBufferBase, BindBufferRange
	TRANSFORM_FEEDBACK_BUFFER = C.GL_TRANSFORM_FEEDBACK_BUFFER
	UNIFORM_BUFFER            = C.GL_UNIFORM_BUFFER

	// BindFramebuffer
	DRAW_FRAMEBUFFER = C.GL_DRAW_FRAMEBUFFER
	READ_FRAMEBUFFER = C.GL_READ_FRAMEBUFFER
	FRAMEBUFFER      = C.GL_FRAMEBUFFER

	// BindTexture
	TEXTURE_1D                   = C.GL_TEXTURE_1D
	TEXTURE_2D                   = C.GL_TEXTURE_2D
	TEXTURE_3D                   = C.GL_TEXTURE_3D
	TEXTURE_1D_ARRAY             = C.GL_TEXTURE_1D_ARRAY
	TEXTURE_2D_ARRAY             = C.GL_TEXTURE_2D_ARRAY
	TEXTURE_RECTANGLE            = C.GL_TEXTURE_RECTANGLE
	TEXTURE_CUBE_MAP             = C.GL_TEXTURE_CUBE_MAP
	TEXTURE_BUFFER               = C.GL_TEXTURE_BUFFER
	TEXTURE_2D_MULTISAMPLE       = C.GL_TEXTURE_2D_MULTISAMPLE
	TEXTURE_2D_MULTISAMPLE_ARRAY = C.GL_TEXTURE_2D_MULTISAMPLE_ARRAY

	// BlendEquation, BlendEquationSeparate
	FUNC_ADD              = C.GL_FUNC_ADD
	FUNC_SUBTRACT         = C.GL_FUNC_SUBTRACT
	FUNC_REVERSE_SUBTRACT = C.GL_FUNC_REVERSE_SUBTRACT
	MIN                   = C.GL_MIN
	MAX                   = C.GL_MAX

	// BlendFunc
	ZERO                     = C.GL_ZERO
	ONE                      = C.GL_ONE
	SRC_COLOR                = C.GL_SRC_COLOR
	ONE_MINUS_SRC_COLOR      = C.GL_ONE_MINUS_SRC_COLOR
	DST_COLOR                = C.GL_DST_COLOR
	ONE_MINUS_DST_COLOR      = C.GL_ONE_MINUS_DST_COLOR
	SRC_ALPHA                = C.GL_SRC_ALPHA
	ONE_MINUS_SRC_ALPHA      = C.GL_ONE_MINUS_SRC_ALPHA
	DST_ALPHA                = C.GL_DST_ALPHA
	ONE_MINUS_DST_ALPHA      = C.GL_ONE_MINUS_DST_ALPHA
	CONSTANT_COLOR           = C.GL_CONSTANT_COLOR
	ONE_MINUS_CONSTANT_COLOR = C.GL_ONE_MINUS_CONSTANT_COLOR
	CONSTANT_ALPHA           = C.GL_CONSTANT_ALPHA
	ONE_MINUS_CONSTANT_ALPHA = C.GL_ONE_MINUS_CONSTANT_ALPHA
	SRC_ALPHA_SATURATE       = C.GL_SRC_ALPHA_SATURATE
	SRC1_COLOR               = C.GL_SRC1_COLOR
	ONE_MINUS_SRC1_COLOR     = C.GL_ONE_MINUS_SRC1_COLOR
	SRC1_ALPHA               = C.GL_SRC1_ALPHA
	ONE_MINUS_SRC1_ALPHA     = C.GL_ONE_MINUS_SRC1_ALPHA

	// BlitFramebuffer
	NEAREST = C.GL_NEAREST
	LINEAR  = C.GL_LINEAR

	// BufferData
	STREAM_DRAW  = C.GL_STREAM_DRAW
	STREAM_READ  = C.GL_STREAM_READ
	STREAM_COPY  = C.GL_STREAM_COPY
	STATIC_DRAW  = C.GL_STATIC_DRAW
	STATIC_READ  = C.GL_STATIC_READ
	STATIC_COPY  = C.GL_STATIC_COPY
	DYNAMIC_DRAW = C.GL_DYNAMIC_DRAW
	DYNAMIC_READ = C.GL_DYNAMIC_READ
	DYNAMIC_COPY = C.GL_DYNAMIC_COPY
)

type GLbitfield int

const (
	COLOR_BUFFER_BIT   GLbitfield = C.GL_COLOR_BUFFER_BIT
	DEPTH_BUFFER_BIT              = C.GL_DEPTH_BUFFER_BIT
	STENCIL_BUFFER_BIT            = C.GL_STENCIL_BUFFER_BIT
)

type GLConstant int

const (
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
