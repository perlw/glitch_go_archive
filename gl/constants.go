package gl

// #include "common.h"
import "C"

type GLenum int

const (
	// ActiveTexture
	Texture0                     GLenum = C.GL_TEXTURE0
	Texture1                            = C.GL_TEXTURE1
	Texture2                            = C.GL_TEXTURE2
	Texture3                            = C.GL_TEXTURE3
	Texture4                            = C.GL_TEXTURE4
	Texture5                            = C.GL_TEXTURE5
	Texture6                            = C.GL_TEXTURE6
	Texture7                            = C.GL_TEXTURE7
	Texture8                            = C.GL_TEXTURE8
	Texture9                            = C.GL_TEXTURE9
	Texture10                           = C.GL_TEXTURE10
	Texture11                           = C.GL_TEXTURE11
	Texture12                           = C.GL_TEXTURE12
	Texture13                           = C.GL_TEXTURE13
	Texture14                           = C.GL_TEXTURE14
	Texture15                           = C.GL_TEXTURE15
	Texture16                           = C.GL_TEXTURE16
	Texture17                           = C.GL_TEXTURE17
	Texture18                           = C.GL_TEXTURE18
	Texture19                           = C.GL_TEXTURE19
	Texture20                           = C.GL_TEXTURE20
	Texture21                           = C.GL_TEXTURE21
	Texture22                           = C.GL_TEXTURE22
	Texture23                           = C.GL_TEXTURE23
	Texture24                           = C.GL_TEXTURE24
	Texture25                           = C.GL_TEXTURE25
	Texture26                           = C.GL_TEXTURE26
	Texture27                           = C.GL_TEXTURE27
	Texture28                           = C.GL_TEXTURE28
	Texture29                           = C.GL_TEXTURE29
	Texture30                           = C.GL_TEXTURE30
	Texture31                           = C.GL_TEXTURE31
	MaxCombinedTextureImageUnits        = C.GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS

	// BeginConditionalRender
	QueryByRegionWait   = C.GL_QUERY_BY_REGION_WAIT
	QueryByRegionNoWait = C.GL_QUERY_BY_REGION_NO_WAIT

	// BeginQuery
	SamplesPassed                      = C.GL_SAMPLES_PASSED
	AnySamplesPassed                   = C.GL_ANY_SAMPLES_PASSED
	PrimiticesGenerated                = C.GL_PRIMITIVES_GENERATED
	TransformFeedbackPrimitivesWritten = C.GL_TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN
	TimeElapsed                        = C.GL_TIME_ELAPSED

	// BeginTransformFeedback
	Points                 = C.GL_POINTS
	Lines                  = C.GL_LINES
	LineLoop               = C.GL_LINE_LOOP
	LineStrip              = C.GL_LINE_STRIP
	LinesAdjacency         = C.GL_LINES_ADJACENCY
	LineStripAdjacency     = C.GL_LINE_STRIP_ADJACENCY
	Triangles              = C.GL_TRIANGLES
	TriangleStrip          = C.GL_TRIANGLE_STRIP
	TriangleFan            = C.GL_TRIANGLE_FAN
	TrianglesAdjacency     = C.GL_TRIANGLES_ADJACENCY
	TriangleStripAdjacency = C.GL_TRIANGLE_STRIP_ADJACENCY

	// BindBuffer, BufferData
	ArrayBuffer        = C.GL_ARRAY_BUFFER
	CopyReadBuffer     = C.GL_COPY_READ_BUFFER
	CopyWriteBuffer    = C.GL_COPY_WRITE_BUFFER
	ElementArrayBuffer = C.GL_ELEMENT_ARRAY_BUFFER
	PixelPackBuffer    = C.GL_PIXEL_PACK_BUFFER
	PixelUnpackBuffer  = C.GL_PIXEL_UNPACK_BUFFER
	// BindBufferBase, BindBufferRange
	TransformFeedbackBuffer = C.GL_TRANSFORM_FEEDBACK_BUFFER
	UniformBuffer           = C.GL_UNIFORM_BUFFER

	// BindFramebuffer
	DrawFramebuffer = C.GL_DRAW_FRAMEBUFFER
	ReadFramebuffer = C.GL_READ_FRAMEBUFFER
	Framebuffer     = C.GL_FRAMEBUFFER

	// BindTexture
	Texture1d                 = C.GL_TEXTURE_1D
	Texture2d                 = C.GL_TEXTURE_2D
	Texture3d                 = C.GL_TEXTURE_3D
	Texture1dArray            = C.GL_TEXTURE_1D_ARRAY
	Texture2dArray            = C.GL_TEXTURE_2D_ARRAY
	TextureRectangle          = C.GL_TEXTURE_RECTANGLE
	TextureCubeMap            = C.GL_TEXTURE_CUBE_MAP
	TextureBuffer             = C.GL_TEXTURE_BUFFER
	Texture2dMultisample      = C.GL_TEXTURE_2D_MULTISAMPLE
	Texture2dMultisampleArray = C.GL_TEXTURE_2D_MULTISAMPLE_ARRAY

	// BlendEquation, BlendEquationSeparate
	FuncAdd             = C.GL_FUNC_ADD
	FuncSubtract        = C.GL_FUNC_SUBTRACT
	FuncReverseSubtract = C.GL_FUNC_REVERSE_SUBTRACT
	Min                 = C.GL_MIN
	Max                 = C.GL_MAX

	// BlendFunc
	Zero                  = C.GL_ZERO
	One                   = C.GL_ONE
	SrcColor              = C.GL_SRC_COLOR
	OneMinusSrcColor      = C.GL_ONE_MINUS_SRC_COLOR
	DstColor              = C.GL_DST_COLOR
	OneMinusDstColor      = C.GL_ONE_MINUS_DST_COLOR
	SrcAlpha              = C.GL_SRC_ALPHA
	OneMinusSrcAlpha      = C.GL_ONE_MINUS_SRC_ALPHA
	DstAlpha              = C.GL_DST_ALPHA
	OneMinusDstAlpha      = C.GL_ONE_MINUS_DST_ALPHA
	ConstantColor         = C.GL_CONSTANT_COLOR
	OneMinusConstantColor = C.GL_ONE_MINUS_CONSTANT_COLOR
	ConstantAlpha         = C.GL_CONSTANT_ALPHA
	OneMinusConstantAlpha = C.GL_ONE_MINUS_CONSTANT_ALPHA
	SrcAlphaSaturate      = C.GL_SRC_ALPHA_SATURATE
	Src1Color             = C.GL_SRC1_COLOR
	OneMinusSrc1Color     = C.GL_ONE_MINUS_SRC1_COLOR
	Src1Alpha             = C.GL_SRC1_ALPHA
	OneMinusSrc1Alpha     = C.GL_ONE_MINUS_SRC1_ALPHA

	// BlitFramebuffer
	Nearest = C.GL_NEAREST
	Linear  = C.GL_LINEAR

	// BufferData
	StreamDraw  = C.GL_STREAM_DRAW
	StreamRead  = C.GL_STREAM_READ
	StreamCopy  = C.GL_STREAM_COPY
	StaticDraw  = C.GL_STATIC_DRAW
	StaticRead  = C.GL_STATIC_READ
	StaticCopy  = C.GL_STATIC_COPY
	DynamicDraw = C.GL_DYNAMIC_DRAW
	DynamicRead = C.GL_DYNAMIC_READ
	DynamicCopy = C.GL_DYNAMIC_COPY
)

type GLbitfield int

const (
	ColorBufferBit   GLbitfield = C.GL_COLOR_BUFFER_BIT
	DepthBufferBit              = C.GL_DEPTH_BUFFER_BIT
	StencilBufferBit            = C.GL_STENCIL_BUFFER_BIT
)

type GLConstant int

const (
	// DepthFunc, AlphaFunc
	Never    = C.GL_NEVER
	Less     = C.GL_LESS
	Equal    = C.GL_EQUAL
	LEqual   = C.GL_LEQUAL
	Greater  = C.GL_GREATER
	NotEqual = C.GL_NOTEQUAL
	GEqual   = C.GL_EQUAL
	Always   = C.GL_ALWAYS

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
