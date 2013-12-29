package gl

// #include "common.h"
import "C"

type GLenum int
type GLbitfield int

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

	// BeginTransformFeedback, DrawArrays, DrawArraysInstanced
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

	// BindBuffer, BufferData, FlushMappedBufferRange
	ArrayBuffer        = C.GL_ARRAY_BUFFER
	CopyReadBuffer     = C.GL_COPY_READ_BUFFER
	CopyWriteBuffer    = C.GL_COPY_WRITE_BUFFER
	ElementArrayBuffer = C.GL_ELEMENT_ARRAY_BUFFER
	PixelPackBuffer    = C.GL_PIXEL_PACK_BUFFER
	PixelUnpackBuffer  = C.GL_PIXEL_UNPACK_BUFFER

	// BindBufferBase, BindBufferRange, FlushMappedBufferRange
	TransformFeedbackBuffer = C.GL_TRANSFORM_FEEDBACK_BUFFER
	UniformBuffer           = C.GL_UNIFORM_BUFFER

	// BindFramebuffer
	DrawFramebuffer = C.GL_DRAW_FRAMEBUFFER
	ReadFramebuffer = C.GL_READ_FRAMEBUFFER
	Framebuffer     = C.GL_FRAMEBUFFER

	// BindRenderbuffer
	Renderbuffer = C.GL_RENDERBUFFER

	// BindTexture, CompressedTexImage
	Texture1d                 = C.GL_TEXTURE_1D
	Texture2d                 = C.GL_TEXTURE_2D
	Texture3d                 = C.GL_TEXTURE_3D
	ProxyTexture1d            = C.GL_PROXY_TEXTURE_1D
	ProxyTexture2d            = C.GL_PROXY_TEXTURE_2D
	ProxyTexture3d            = C.GL_PROXY_TEXTURE_3D
	Texture1dArray            = C.GL_TEXTURE_1D_ARRAY
	Texture2dArray            = C.GL_TEXTURE_2D_ARRAY
	TextureRectangle          = C.GL_TEXTURE_RECTANGLE
	TextureCubeMap            = C.GL_TEXTURE_CUBE_MAP
	TextureBuffer             = C.GL_TEXTURE_BUFFER
	Texture2dMultisample      = C.GL_TEXTURE_2D_MULTISAMPLE
	Texture2dMultisampleArray = C.GL_TEXTURE_2D_MULTISAMPLE_ARRAY
	ProxyTexture1dArray       = C.GL_PROXY_TEXTURE_1D_ARRAY
	ProxyTexture2dArray       = C.GL_PROXY_TEXTURE_2D_ARRAY
	TextureCubeMapPositiveX   = C.GL_TEXTURE_CUBE_MAP_POSITIVE_X
	TextureCubeMapNegativeX   = C.GL_TEXTURE_CUBE_MAP_NEGATIVE_X
	TextureCubeMapPositiveY   = C.GL_TEXTURE_CUBE_MAP_POSITIVE_Y
	TextureCubeMapNegativeY   = C.GL_TEXTURE_CUBE_MAP_NEGATIVE_Y
	TextureCubeMapPositiveZ   = C.GL_TEXTURE_CUBE_MAP_POSITIVE_Z
	TextureCubeMapNegativeZ   = C.GL_TEXTURE_CUBE_MAP_NEGATIVE_Z
	ProxyTextureCubeMap       = C.GL_PROXY_TEXTURE_CUBE_MAP

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

	// CheckFramebufferStatus
	FramebufferUndefined                   = C.GL_FRAMEBUFFER_UNDEFINED
	FramebufferIncompleteAttachment        = C.GL_FRAMEBUFFER_INCOMPLETE_ATTACHMENT
	FramebufferIncompleteMissingAttachment = C.GL_FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT
	FramebufferIncompleteDrawBuffer        = C.GL_FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER
	FramebufferIncompleteReadBuffer        = C.GL_FRAMEBUFFER_INCOMPLETE_READ_BUFFER
	FramebufferUnsupported                 = C.GL_FRAMEBUFFER_UNSUPPORTED
	FramebufferIncompleteMultisample       = C.GL_FRAMEBUFFER_INCOMPLETE_MULTISAMPLE
	FramebufferIncompleteLayerTargets      = C.GL_FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS

	// ClampColor
	ClampReadColor = C.GL_CLAMP_READ_COLOR

	// CopyTexImage
	CompressedRed       = C.GL_COMPRESSED_RED
	CompressedRg        = C.GL_COMPRESSED_RG
	CompressedRgb       = C.GL_COMPRESSED_RGB
	CompressedRgba      = C.GL_COMPRESSED_RGBA
	CompressedSrgb      = C.GL_COMPRESSED_SRGB
	CompressedSrgbAlpha = C.GL_COMPRESSED_SRGB_ALPHA
	DepthComponent      = C.GL_DEPTH_COMPONENT
	DepthComponent16    = C.GL_DEPTH_COMPONENT16
	DepthComponent24    = C.GL_DEPTH_COMPONENT24
	DepthComponent32    = C.GL_DEPTH_COMPONENT32
	Red                 = C.GL_RED
	Rg                  = C.GL_RG
	Rgb                 = C.GL_RGB
	R3G3B2              = C.GL_R3_G3_B2
	Rgb4                = C.GL_RGB4
	Rgb5                = C.GL_RGB5
	Rgb8                = C.GL_RGB8
	Rgb10               = C.GL_RGB10
	Rgb12               = C.GL_RGB12
	Rgb16               = C.GL_RGB16
	Rgba                = C.GL_RGBA
	Rgba2               = C.GL_RGBA2
	Rgba4               = C.GL_RGBA4
	Rgb5A1              = C.GL_RGB5_A1
	Rgba8               = C.GL_RGBA8
	Rgb10A2             = C.GL_RGB10_A2
	Rgba12              = C.GL_RGBA12
	Rgba16              = C.GL_RGBA16
	Srgb                = C.GL_SRGB
	Srgb8               = C.GL_SRGB8
	SrgbAlpha           = C.GL_SRGB_ALPHA
	Srgb8Alpha8         = C.GL_SRGB8_ALPHA8

	// CreateShader
	VertexShader   = C.GL_VERTEX_SHADER
	GeometryShader = C.GL_GEOMETRY_SHADER
	FragmentShader = C.GL_FRAGMENT_SHADER

	// CullFace, DrawBuffer
	Front        = C.GL_FRONT
	Back         = C.GL_BACK
	FrontAndBack = C.GL_FRONT_AND_BACK

	// DrawBuffer
	None       = C.GL_NONE
	FrontLeft  = C.GL_FRONT_LEFT
	FrontRight = C.GL_FRONT_RIGHT
	BackLeft   = C.GL_BACK_LEFT
	BackRight  = C.GL_BACK_RIGHT
	Left       = C.GL_LEFT
	Right      = C.GL_RIGHT

	// DepthFunc
	Never    = C.GL_NEVER
	Less     = C.GL_LESS
	Equal    = C.GL_EQUAL
	LEqual   = C.GL_LEQUAL
	Greater  = C.GL_GREATER
	NotEqual = C.GL_NOTEQUAL
	GEqual   = C.GL_EQUAL
	Always   = C.GL_ALWAYS

	// Disable, Enable
	Blend                  = C.GL_BLEND
	ClipDistance0          = C.GL_CLIP_DISTANCE0
	ClipDistance1          = C.GL_CLIP_DISTANCE1
	ClipDistance2          = C.GL_CLIP_DISTANCE2
	ClipDistance3          = C.GL_CLIP_DISTANCE3
	ClipDistance4          = C.GL_CLIP_DISTANCE4
	ClipDistance5          = C.GL_CLIP_DISTANCE5
	ColorLogicOp           = C.GL_COLOR_LOGIC_OP
	CullFaceConst          = C.GL_CULL_FACE
	DepthClamp             = C.GL_DEPTH_CLAMP
	DepthTest              = C.GL_DEPTH_TEST
	Dither                 = C.GL_DITHER
	FramebufferSrgb        = C.GL_FRAMEBUFFER_SRGB
	LineSmooth             = C.GL_LINE_SMOOTH
	Multisample            = C.GL_MULTISAMPLE
	PolygonOffsetFill      = C.GL_POLYGON_OFFSET_FILL
	PolygonOffsetLine      = C.GL_POLYGON_OFFSET_LINE
	PolygonOffsetPoint     = C.GL_POLYGON_OFFSET_POINT
	PolygonSmooth          = C.GL_POLYGON_SMOOTH
	PrimitiveRestart       = C.GL_PRIMITIVE_RESTART
	SampleAlphaToCoverage  = C.GL_SAMPLE_ALPHA_TO_COVERAGE
	SampleAlphaToOne       = C.GL_SAMPLE_ALPHA_TO_ONE
	SampleCoverage         = C.GL_SAMPLE_COVERAGE
	ScissorTest            = C.GL_SCISSOR_TEST
	StencilTest            = C.GL_STENCIL_TEST
	TextureCubeMapSeamless = C.GL_TEXTURE_CUBE_MAP_SEAMLESS
	ProgramPointSize       = C.GL_PROGRAM_POINT_SIZE

	// FenceSync
	SyncGpuCommandsComplete = C.GL_SYNC_GPU_COMMANDS_COMPLETE

	// glFramebufferTexture, glFramebufferTextureLayer
	ColorAttachment0       = C.GL_COLOR_ATTACHMENT0
	ColorAttachment1       = C.GL_COLOR_ATTACHMENT1
	ColorAttachment2       = C.GL_COLOR_ATTACHMENT2
	ColorAttachment3       = C.GL_COLOR_ATTACHMENT3
	ColorAttachment4       = C.GL_COLOR_ATTACHMENT4
	ColorAttachment5       = C.GL_COLOR_ATTACHMENT5
	ColorAttachment6       = C.GL_COLOR_ATTACHMENT6
	ColorAttachment7       = C.GL_COLOR_ATTACHMENT7
	ColorAttachment8       = C.GL_COLOR_ATTACHMENT8
	ColorAttachment9       = C.GL_COLOR_ATTACHMENT9
	ColorAttachment10      = C.GL_COLOR_ATTACHMENT10
	ColorAttachment11      = C.GL_COLOR_ATTACHMENT11
	ColorAttachment12      = C.GL_COLOR_ATTACHMENT12
	ColorAttachment13      = C.GL_COLOR_ATTACHMENT13
	ColorAttachment14      = C.GL_COLOR_ATTACHMENT14
	ColorAttachment15      = C.GL_COLOR_ATTACHMENT15
	DepthAttachment        = C.GL_DEPTH_ATTACHMENT
	StencilAttachment      = C.GL_STENCIL_ATTACHMENT
	DepthStencilAttachment = C.GL_DEPTH_STENCIL_ATTACHMENT

	// FrontFace
	CW  = C.GL_CW
	CCW = C.GL_CCW

	// Everything else
	ActiveTextureConst                   = C.GL_ACTIVE_TEXTURE
	AliasedLineWidthRange                = C.GL_ALIASED_LINE_WIDTH_RANGE
	SmoothLineWidthRange                 = C.GL_SMOOTH_LINE_WIDTH_RANGE
	SmoothLineWidthGranularity           = C.GL_SMOOTH_LINE_WIDTH_GRANULARITY
	ArrayBufferBinding                   = C.GL_ARRAY_BUFFER_BINDING
	BlendColorConst                      = C.GL_BLEND_COLOR
	BlendDstAlpha                        = C.GL_BLEND_DST_ALPHA
	BlendDstRgb                          = C.GL_BLEND_DST_RGB
	BlendEquationRgb                     = C.GL_BLEND_EQUATION_RGB
	BlendEquationAlpha                   = C.GL_BLEND_EQUATION_ALPHA
	BlendSrcAlpha                        = C.GL_BLEND_SRC_ALPHA
	BlendSrcRgb                          = C.GL_BLEND_SRC_RGB
	ColorClearValue                      = C.GL_COLOR_CLEAR_VALUE
	ColorWritemask                       = C.GL_COLOR_WRITEMASK
	CompressedTextureFormats             = C.GL_COMPRESSED_TEXTURE_FORMATS
	CurrentProgram                       = C.GL_CURRENT_PROGRAM
	DepthClearValue                      = C.GL_DEPTH_CLEAR_VALUE
	DepthFuncConst                       = C.GL_DEPTH_FUNC
	DepthRangeConst                      = C.GL_DEPTH_RANGE
	DepthWritemask                       = C.GL_DEPTH_WRITEMASK
	Doublebuffer                         = C.GL_DOUBLEBUFFER
	DrawBufferConst                      = C.GL_DRAW_BUFFER
	DrawBuffer0                          = C.GL_DRAW_BUFFER0
	DrawBuffer1                          = C.GL_DRAW_BUFFER1
	DrawBuffer2                          = C.GL_DRAW_BUFFER2
	DrawBuffer3                          = C.GL_DRAW_BUFFER3
	DrawBuffer4                          = C.GL_DRAW_BUFFER4
	DrawBuffer5                          = C.GL_DRAW_BUFFER5
	DrawBuffer6                          = C.GL_DRAW_BUFFER6
	DrawBuffer7                          = C.GL_DRAW_BUFFER7
	DrawBuffer8                          = C.GL_DRAW_BUFFER8
	DrawBuffer9                          = C.GL_DRAW_BUFFER9
	DrawBuffer10                         = C.GL_DRAW_BUFFER10
	DrawBuffer11                         = C.GL_DRAW_BUFFER11
	DrawBuffer12                         = C.GL_DRAW_BUFFER12
	DrawBuffer13                         = C.GL_DRAW_BUFFER13
	DrawBuffer14                         = C.GL_DRAW_BUFFER14
	DrawBuffer15                         = C.GL_DRAW_BUFFER15
	DrawFramebufferBinding               = C.GL_DRAW_FRAMEBUFFER_BINDING
	ReadFramebufferBinding               = C.GL_READ_FRAMEBUFFER_BINDING
	ElementArrayBufferBinding            = C.GL_ELEMENT_ARRAY_BUFFER_BINDING
	RenderbufferBinding                  = C.GL_RENDERBUFFER_BINDING
	FragmentShaderDerivativeHint         = C.GL_FRAGMENT_SHADER_DERIVATIVE_HINT
	LineSmoothHint                       = C.GL_LINE_SMOOTH_HINT
	LineWidthConst                       = C.GL_LINE_WIDTH
	LogicOpMode                          = C.GL_LOGIC_OP_MODE
	Max3dTextureSize                     = C.GL_MAX_3D_TEXTURE_SIZE
	MaxClipDistances                     = C.GL_MAX_CLIP_DISTANCES
	MaxCombinedFragmentUniformComponents = C.GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS
	MaxCombinedVertexUniformComponents   = C.GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS
	MaxCombinedGeometryUniformComponents = C.GL_MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS
	MaxVaryingComponents                 = C.GL_MAX_VARYING_COMPONENTS
	MaxCombinedUniformBlocks             = C.GL_MAX_COMBINED_UNIFORM_BLOCKS
	MaxCubeMapTextureSize                = C.GL_MAX_CUBE_MAP_TEXTURE_SIZE
	MaxDrawBuffers                       = C.GL_MAX_DRAW_BUFFERS
	MaxDualSourceDrawBuffers             = C.GL_MAX_DUAL_SOURCE_DRAW_BUFFERS
	MaxElementsIndices                   = C.GL_MAX_ELEMENTS_INDICES
	MaxElementsVertices                  = C.GL_MAX_ELEMENTS_VERTICES
	MaxFragmentUniformComponents         = C.GL_MAX_FRAGMENT_UNIFORM_COMPONENTS
	MaxFragmentUniformBlocks             = C.GL_MAX_FRAGMENT_UNIFORM_BLOCKS
	MaxFragmentInputComponents           = C.GL_MAX_FRAGMENT_INPUT_COMPONENTS
	MinProgramTexelOffset                = C.GL_MIN_PROGRAM_TEXEL_OFFSET
	MaxProgramTexelOffset                = C.GL_MAX_PROGRAM_TEXEL_OFFSET
	MaxRectangleTextureSize              = C.GL_MAX_RECTANGLE_TEXTURE_SIZE
	MaxTextureImageUnits                 = C.GL_MAX_TEXTURE_IMAGE_UNITS
	MaxTextureLodBias                    = C.GL_MAX_TEXTURE_LOD_BIAS
	MaxTextureSize                       = C.GL_MAX_TEXTURE_SIZE
	MaxRenderbufferSize                  = C.GL_MAX_RENDERBUFFER_SIZE
	MaxArrayTextureLayers                = C.GL_MAX_ARRAY_TEXTURE_LAYERS
	MaxTextureBufferSize                 = C.GL_MAX_TEXTURE_BUFFER_SIZE
	MaxUniformBlockSize                  = C.GL_MAX_UNIFORM_BLOCK_SIZE
	MaxVaryingFloats                     = C.GL_MAX_VARYING_FLOATS
	MaxVertexAttribs                     = C.GL_MAX_VERTEX_ATTRIBS
	MaxVertexTextureImageUnits           = C.GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS
	MaxGeometryTextureImageUnits         = C.GL_MAX_GEOMETRY_TEXTURE_IMAGE_UNITS
	MaxVertexUniformComponents           = C.GL_MAX_VERTEX_UNIFORM_COMPONENTS
	MaxVertexOutputComponents            = C.GL_MAX_VERTEX_OUTPUT_COMPONENTS
	MaxGeometryUniformComponents         = C.GL_MAX_GEOMETRY_UNIFORM_COMPONENTS
	MaxSampleMaskWords                   = C.GL_MAX_SAMPLE_MASK_WORDS
	MaxColorTextureSamples               = C.GL_MAX_COLOR_TEXTURE_SAMPLES
	MaxDepthTextureSamples               = C.GL_MAX_DEPTH_TEXTURE_SAMPLES
	MaxIntegerSamples                    = C.GL_MAX_INTEGER_SAMPLES
	MaxServerWaitTimeout                 = C.GL_MAX_SERVER_WAIT_TIMEOUT
	MaxUniformBufferBindings             = C.GL_MAX_UNIFORM_BUFFER_BINDINGS
	UniformBufferOffsetAlignment         = C.GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT
	MaxVertexUniformBlocks               = C.GL_MAX_VERTEX_UNIFORM_BLOCKS
	MaxGeometryUniformBlocks             = C.GL_MAX_GEOMETRY_UNIFORM_BLOCKS
	MaxGeometryInputComponents           = C.GL_MAX_GEOMETRY_INPUT_COMPONENTS
	MaxGeometryOutputComponents          = C.GL_MAX_GEOMETRY_OUTPUT_COMPONENTS
	MaxViewportDims                      = C.GL_MAX_VIEWPORT_DIMS
	NumCompressedTextureFormats          = C.GL_NUM_COMPRESSED_TEXTURE_FORMATS
	PackAlignment                        = C.GL_PACK_ALIGNMENT
	PackImageHeight                      = C.GL_PACK_IMAGE_HEIGHT
	PackLsbFirst                         = C.GL_PACK_LSB_FIRST
	PackRowLength                        = C.GL_PACK_ROW_LENGTH
	PackSkipImages                       = C.GL_PACK_SKIP_IMAGES
	PackSkipPixels                       = C.GL_PACK_SKIP_PIXELS
	PackSkipRows                         = C.GL_PACK_SKIP_ROWS
	PackSwapBytes                        = C.GL_PACK_SWAP_BYTES
	PixelPackBufferBinding               = C.GL_PIXEL_PACK_BUFFER_BINDING
	PixelUnpackBufferBinding             = C.GL_PIXEL_UNPACK_BUFFER_BINDING
	PointFadeThresholdSize               = C.GL_POINT_FADE_THRESHOLD_SIZE
	PrimitiveRestartIndex                = C.GL_PRIMITIVE_RESTART_INDEX
	ProvokingVertex                      = C.GL_PROVOKING_VERTEX
	PointSize                            = C.GL_POINT_SIZE
	PointSizeGranularity                 = C.GL_POINT_SIZE_GRANULARITY
	PointSizeRange                       = C.GL_POINT_SIZE_RANGE
	PolygonOffsetFactor                  = C.GL_POLYGON_OFFSET_FACTOR
	PolygonOffsetUnits                   = C.GL_POLYGON_OFFSET_UNITS
	PolygonSmoothHint                    = C.GL_POLYGON_SMOOTH_HINT
	ReadBuffer                           = C.GL_READ_BUFFER
	SampleBuffers                        = C.GL_SAMPLE_BUFFERS
	SampleCoverageValue                  = C.GL_SAMPLE_COVERAGE_VALUE
	SampleCoverageInvert                 = C.GL_SAMPLE_COVERAGE_INVERT
	SamplerBinding                       = C.GL_SAMPLER_BINDING
	Samples                              = C.GL_SAMPLES
	ScissorBox                           = C.GL_SCISSOR_BOX
	StencilBackFail                      = C.GL_STENCIL_BACK_FAIL
	StencilBackFunc                      = C.GL_STENCIL_BACK_FUNC
	StencilBackPassDepthFail             = C.GL_STENCIL_BACK_PASS_DEPTH_FAIL
	StencilBackPassDepthPass             = C.GL_STENCIL_BACK_PASS_DEPTH_PASS
	StencilBackRef                       = C.GL_STENCIL_BACK_REF
	StencilBackValueMask                 = C.GL_STENCIL_BACK_VALUE_MASK
	StencilBackWritemask                 = C.GL_STENCIL_BACK_WRITEMASK
	StencilClearValue                    = C.GL_STENCIL_CLEAR_VALUE
	StencilFail                          = C.GL_STENCIL_FAIL
	StencilFunc                          = C.GL_STENCIL_FUNC
	StencilPassDepthFail                 = C.GL_STENCIL_PASS_DEPTH_FAIL
	StencilPassDepthPass                 = C.GL_STENCIL_PASS_DEPTH_PASS
	StencilRef                           = C.GL_STENCIL_REF
	StencilValueMask                     = C.GL_STENCIL_VALUE_MASK
	StencilWritemask                     = C.GL_STENCIL_WRITEMASK
	Stereo                               = C.GL_STEREO
	SubpixelBits                         = C.GL_SUBPIXEL_BITS
	TextureBinding1d                     = C.GL_TEXTURE_BINDING_1D
	TextureBinding1dArray                = C.GL_TEXTURE_BINDING_1D_ARRAY
	TextureBinding2d                     = C.GL_TEXTURE_BINDING_2D
	TextureBinding2dArray                = C.GL_TEXTURE_BINDING_2D_ARRAY
	TextureBinding2dMultisample          = C.GL_TEXTURE_BINDING_2D_MULTISAMPLE
	TextureBinding2dMultisampleArray     = C.GL_TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY
	TextureBinding3d                     = C.GL_TEXTURE_BINDING_3D
	TextureBindingBuffer                 = C.GL_TEXTURE_BINDING_BUFFER
	TextureBindingCubeMap                = C.GL_TEXTURE_BINDING_CUBE_MAP
	TextureBindingRectangle              = C.GL_TEXTURE_BINDING_RECTANGLE
	TextureCompressionHint               = C.GL_TEXTURE_COMPRESSION_HINT
	Timestamp                            = C.GL_TIMESTAMP
	TransformFeedbackBufferBinding       = C.GL_TRANSFORM_FEEDBACK_BUFFER_BINDING
	TransformFeedbackBufferStart         = C.GL_TRANSFORM_FEEDBACK_BUFFER_START
	TransformFeedbackBufferSize          = C.GL_TRANSFORM_FEEDBACK_BUFFER_SIZE
	UniformBufferBinding                 = C.GL_UNIFORM_BUFFER_BINDING
	UniformBufferStart                   = C.GL_UNIFORM_BUFFER_START
	UniformBufferSize                    = C.GL_UNIFORM_BUFFER_SIZE
	UnpackAlignment                      = C.GL_UNPACK_ALIGNMENT
	UnpackImageHeight                    = C.GL_UNPACK_IMAGE_HEIGHT
	UnpackLsbFirst                       = C.GL_UNPACK_LSB_FIRST
	UnpackRowLength                      = C.GL_UNPACK_ROW_LENGTH
	UnpackSkipImages                     = C.GL_UNPACK_SKIP_IMAGES
	UnpackSkipPixels                     = C.GL_UNPACK_SKIP_PIXELS
	UnpackSkipRows                       = C.GL_UNPACK_SKIP_ROWS
	UnpackSwapBytes                      = C.GL_UNPACK_SWAP_BYTES
	NumExtensions                        = C.GL_NUM_EXTENSIONS
	MajorVersion                         = C.GL_MAJOR_VERSION
	MinorVersion                         = C.GL_MINOR_VERSION
	ContextFlags                         = C.GL_CONTEXT_FLAGS
	ViewportConst                        = C.GL_VIEWPORT

	// GetActiveAttrib, GetActiveUniform, etc
	Float                                = C.GL_FLOAT
	FloatVec2                            = C.GL_FLOAT_VEC2
	FloatVec3                            = C.GL_FLOAT_VEC3
	FloatVec4                            = C.GL_FLOAT_VEC4
	Int                                  = C.GL_INT
	IntVec2                              = C.GL_INT_VEC2
	IntVec3                              = C.GL_INT_VEC3
	IntVec4                              = C.GL_INT_VEC4
	UnsignedInt                          = C.GL_UNSIGNED_INT
	UnsignedIntVec2                      = C.GL_UNSIGNED_INT_VEC2
	UnsignedIntVec3                      = C.GL_UNSIGNED_INT_VEC3
	UnsignedIntVec4                      = C.GL_UNSIGNED_INT_VEC4
	Bool                                 = C.GL_BOOL
	BoolVec2                             = C.GL_BOOL_VEC2
	BoolVec3                             = C.GL_BOOL_VEC3
	BoolVec4                             = C.GL_BOOL_VEC4
	FloatMat2                            = C.GL_FLOAT_MAT2
	FloatMat3                            = C.GL_FLOAT_MAT3
	FloatMat4                            = C.GL_FLOAT_MAT4
	FLOATMAT2x3                          = C.GL_FLOAT_MAT2x3
	FLOATMAT2x4                          = C.GL_FLOAT_MAT2x4
	FLOATMAT3x2                          = C.GL_FLOAT_MAT3x2
	FLOATMAT3x4                          = C.GL_FLOAT_MAT3x4
	FLOATMAT4x2                          = C.GL_FLOAT_MAT4x2
	FLOATMAT4x3                          = C.GL_FLOAT_MAT4x3
	Sampler1d                            = C.GL_SAMPLER_1D
	Sampler2d                            = C.GL_SAMPLER_2D
	Sampler3d                            = C.GL_SAMPLER_3D
	SamplerCube                          = C.GL_SAMPLER_CUBE
	Sampler1dShadow                      = C.GL_SAMPLER_1D_SHADOW
	Sampler2dShadow                      = C.GL_SAMPLER_2D_SHADOW
	Sampler1dArray                       = C.GL_SAMPLER_1D_ARRAY
	Sampler2dArray                       = C.GL_SAMPLER_2D_ARRAY
	Sampler1dArrayShadow                 = C.GL_SAMPLER_1D_ARRAY_SHADOW
	Sampler2dArrayShadow                 = C.GL_SAMPLER_2D_ARRAY_SHADOW
	Sampler2dMultisample                 = C.GL_SAMPLER_2D_MULTISAMPLE
	Sampler2dMultisampleArray            = C.GL_SAMPLER_2D_MULTISAMPLE_ARRAY
	SamplerCubeShadow                    = C.GL_SAMPLER_CUBE_SHADOW
	SamplerBuffer                        = C.GL_SAMPLER_BUFFER
	Sampler2dRect                        = C.GL_SAMPLER_2D_RECT
	Sampler2dRectShadow                  = C.GL_SAMPLER_2D_RECT_SHADOW
	IntSampler1d                         = C.GL_INT_SAMPLER_1D
	IntSampler2d                         = C.GL_INT_SAMPLER_2D
	IntSampler3d                         = C.GL_INT_SAMPLER_3D
	IntSamplerCube                       = C.GL_INT_SAMPLER_CUBE
	IntSampler1dArray                    = C.GL_INT_SAMPLER_1D_ARRAY
	IntSampler2dArray                    = C.GL_INT_SAMPLER_2D_ARRAY
	IntSampler2dMultisample              = C.GL_INT_SAMPLER_2D_MULTISAMPLE
	IntSampler2dMultisampleArray         = C.GL_INT_SAMPLER_2D_MULTISAMPLE_ARRAY
	IntSamplerBuffer                     = C.GL_INT_SAMPLER_BUFFER
	IntSampler2dRect                     = C.GL_INT_SAMPLER_2D_RECT
	UnsignedIntSampler1d                 = C.GL_UNSIGNED_INT_SAMPLER_1D
	UnsignedIntSampler2d                 = C.GL_UNSIGNED_INT_SAMPLER_2D
	UnsignedIntSampler3d                 = C.GL_UNSIGNED_INT_SAMPLER_3D
	UnsignedIntSamplerCube               = C.GL_UNSIGNED_INT_SAMPLER_CUBE
	UnsignedIntSampler1dArray            = C.GL_UNSIGNED_INT_SAMPLER_1D_ARRAY
	UnsignedIntSampler2dArray            = C.GL_UNSIGNED_INT_SAMPLER_2D_ARRAY
	UnsignedIntSampler2dMultisample      = C.GL_UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE
	UnsignedIntSampler2dMultisampleArray = C.GL_UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE_ARRAY
	UnsignedIntSamplerBuffer             = C.GL_UNSIGNED_INT_SAMPLER_BUFFER
	UnsignedIntSampler2dRect             = C.GL_UNSIGNED_INT_SAMPLER_2D_RECT

	// GetActiveUniformBlock, etc
	UniformBlockBinding                    = C.GL_UNIFORM_BLOCK_BINDING
	UniformBlockDataSize                   = C.GL_UNIFORM_BLOCK_DATA_SIZE
	UniformBlockNameLength                 = C.GL_UNIFORM_BLOCK_NAME_LENGTH
	UniformBlockActiveUniforms             = C.GL_UNIFORM_BLOCK_ACTIVE_UNIFORMS
	UniformBlockActiveUniformIndices       = C.GL_UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES
	UniformBlockReferencedByVertexShader   = C.GL_UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER
	UniformBlockReferencedByGeometryShader = C.GL_UNIFORM_BLOCK_REFERENCED_BY_GEOMETRY_SHADER
	UniformBlockReferencedByFragmentShader = C.GL_UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER

	// GetBufferParameter, GetBufferPointer, etc
	BufferAccess = C.GL_BUFFER_ACCESS
	BufferMapped = C.GL_BUFFER_MAPPED
	BufferSize   = C.GL_BUFFER_SIZE
	BufferUsage  = C.GL_BUFFER_USAGE

	// GetError
	NoError                     = C.GL_NO_ERROR
	InvalidEnum                 = C.GL_INVALID_ENUM
	InvalidValue                = C.GL_INVALID_VALUE
	InvalidOperation            = C.GL_INVALID_OPERATION
	InvalidFramebufferOperation = C.GL_INVALID_FRAMEBUFFER_OPERATION
	OutOfMemory                 = C.GL_OUT_OF_MEMORY

	// GetMultisample
	SamplePosition = C.GL_SAMPLE_POSITION

	// GetProgram
	DeleteStatus                      = C.GL_DELETE_STATUS
	LinkStatus                        = C.GL_LINK_STATUS
	ValidateStatus                    = C.GL_VALIDATE_STATUS
	InfoLogLength                     = C.GL_INFO_LOG_LENGTH
	AttachedShaders                   = C.GL_ATTACHED_SHADERS
	ActiveAttributes                  = C.GL_ACTIVE_ATTRIBUTES
	ActiveAttributeMaxLength          = C.GL_ACTIVE_ATTRIBUTE_MAX_LENGTH
	ActiveUniforms                    = C.GL_ACTIVE_UNIFORMS
	ActiveUniformBlocks               = C.GL_ACTIVE_UNIFORM_BLOCKS
	ActiveUniformBlockMaxNameLength   = C.GL_ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH
	ActiveUniformMaxLength            = C.GL_ACTIVE_UNIFORM_MAX_LENGTH
	TransformFeedbackBufferMode       = C.GL_TRANSFORM_FEEDBACK_BUFFER_MODE
	TransformFeedbackVaryings         = C.GL_TRANSFORM_FEEDBACK_VARYINGS
	TransformFeedbackVaryingMaxLength = C.GL_TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH
	GeometryVerticesOut               = C.GL_GEOMETRY_VERTICES_OUT
	GeometryInputType                 = C.GL_GEOMETRY_INPUT_TYPE
	GeometryOutputType                = C.GL_GEOMETRY_OUTPUT_TYPE

	// GetQueryObject
	QueryResult          = C.GL_QUERY_RESULT
	QueryResultAvailable = C.GL_QUERY_RESULT_AVAILABLE

	// GetQuery
	PrimitivesGenerated = C.GL_PRIMITIVES_GENERATED
	CurrentQuery        = C.GL_CURRENT_QUERY
	QueryCounterBits    = C.GL_QUERY_COUNTER_BITS

	// GetSamplerParameter
	TextureMagFilter   = C.GL_TEXTURE_MAG_FILTER
	TextureMinFilter   = C.GL_TEXTURE_MIN_FILTER
	TextureMinLod      = C.GL_TEXTURE_MIN_LOD
	TextureMaxLod      = C.GL_TEXTURE_MAX_LOD
	TextureLodBias     = C.GL_TEXTURE_LOD_BIAS
	TextureWrapS       = C.GL_TEXTURE_WRAP_S
	TextureWrapT       = C.GL_TEXTURE_WRAP_T
	TextureWrapR       = C.GL_TEXTURE_WRAP_R
	TextureBorderColor = C.GL_TEXTURE_BORDER_COLOR
	TextureCompareMode = C.GL_TEXTURE_COMPARE_MODE
	TextureCompareFunc = C.GL_TEXTURE_COMPARE_FUNC

	// GetShader
	ShaderType         = C.GL_SHADER_TYPE
	CompileStatus      = C.GL_COMPILE_STATUS
	ShaderSourceLength = C.GL_SHADER_SOURCE_LENGTH

	// GetString
	Vendor                 = C.GL_VENDOR
	Renderer               = C.GL_RENDERER
	Version                = C.GL_VERSION
	ShadingLanguageVersion = C.GL_SHADING_LANGUAGE_VERSION
	Extensions             = C.GL_EXTENSIONS

	// GetTexImage
	StencilIndex             = C.GL_STENCIL_INDEX
	DepthStencil             = C.GL_DEPTH_STENCIL
	Green                    = C.GL_GREEN
	Blue                     = C.GL_BLUE
	Bgr                      = C.GL_BGR
	Bgra                     = C.GL_BGRA
	RedInteger               = C.GL_RED_INTEGER
	GreenInteger             = C.GL_GREEN_INTEGER
	BlueInteger              = C.GL_BLUE_INTEGER
	RgInteger                = C.GL_RG_INTEGER
	RgbInteger               = C.GL_RGB_INTEGER
	RgbaInteger              = C.GL_RGBA_INTEGER
	BgrInteger               = C.GL_BGR_INTEGER
	BgraInteger              = C.GL_BGRA_INTEGER
	UnsignedByte             = C.GL_UNSIGNED_BYTE
	Byte                     = C.GL_BYTE
	UnsignedShort            = C.GL_UNSIGNED_SHORT
	Short                    = C.GL_SHORT
	HalfFloat                = C.GL_HALF_FLOAT
	UnsignedByte332          = C.GL_UNSIGNED_BYTE_3_3_2
	UnsignedByte233Rev       = C.GL_UNSIGNED_BYTE_2_3_3_REV
	UnsignedShort565         = C.GL_UNSIGNED_SHORT_5_6_5
	UnsignedShort565Rev      = C.GL_UNSIGNED_SHORT_5_6_5_REV
	UnsignedShort4444        = C.GL_UNSIGNED_SHORT_4_4_4_4
	UnsignedShort4444Rev     = C.GL_UNSIGNED_SHORT_4_4_4_4_REV
	UnsignedShort5551        = C.GL_UNSIGNED_SHORT_5_5_5_1
	UnsignedShort1555Rev     = C.GL_UNSIGNED_SHORT_1_5_5_5_REV
	UnsignedInt8888          = C.GL_UNSIGNED_INT_8_8_8_8
	UnsignedInt8888Rev       = C.GL_UNSIGNED_INT_8_8_8_8_REV
	UnsignedInt1010102       = C.GL_UNSIGNED_INT_10_10_10_2
	UnsignedInt2101010Rev    = C.GL_UNSIGNED_INT_2_10_10_10_REV
	UnsignedInt248           = C.GL_UNSIGNED_INT_24_8
	UnsignedInt10f11f11fRev  = C.GL_UNSIGNED_INT_10F_11F_11F_REV
	UnsignedInt5999Rev       = C.GL_UNSIGNED_INT_5_9_9_9_REV
	Float32UnsignedInt248Rev = C.GL_FLOAT_32_UNSIGNED_INT_24_8_REV

	// GetTexLevelParameter
	TextureWidth               = C.GL_TEXTURE_WIDTH
	TextureHeight              = C.GL_TEXTURE_HEIGHT
	TextureDepth               = C.GL_TEXTURE_DEPTH
	TextureInternalFormat      = C.GL_TEXTURE_INTERNAL_FORMAT
	TextureRedSize             = C.GL_TEXTURE_RED_SIZE
	TextureGreenSize           = C.GL_TEXTURE_GREEN_SIZE
	TextureBlueSize            = C.GL_TEXTURE_BLUE_SIZE
	TextureAlphaSize           = C.GL_TEXTURE_ALPHA_SIZE
	TextureDepthSize           = C.GL_TEXTURE_DEPTH_SIZE
	TextureCompressed          = C.GL_TEXTURE_COMPRESSED
	TextureCompressedImageSize = C.GL_TEXTURE_COMPRESSED_IMAGE_SIZE

	// GetTexParameter
	TextureBaseLevel   = C.GL_TEXTURE_BASE_LEVEL
	TextureMaxLevel    = C.GL_TEXTURE_MAX_LEVEL
	TextureSwizzleR    = C.GL_TEXTURE_SWIZZLE_R
	TextureSwizzleG    = C.GL_TEXTURE_SWIZZLE_G
	TextureSwizzleB    = C.GL_TEXTURE_SWIZZLE_B
	TextureSwizzleA    = C.GL_TEXTURE_SWIZZLE_A
	TextureSwizzleRgba = C.GL_TEXTURE_SWIZZLE_RGBA

	// GetVertexAttrib
	VertexAttribArrayBufferBinding = C.GL_VERTEX_ATTRIB_ARRAY_BUFFER_BINDING
	VertexAttribArrayEnabled       = C.GL_VERTEX_ATTRIB_ARRAY_ENABLED
	VertexAttribArraySize          = C.GL_VERTEX_ATTRIB_ARRAY_SIZE
	VertexAttribArrayStride        = C.GL_VERTEX_ATTRIB_ARRAY_STRIDE
	VertexAttribArrayType          = C.GL_VERTEX_ATTRIB_ARRAY_TYPE
	VertexAttribArrayNormalized    = C.GL_VERTEX_ATTRIB_ARRAY_NORMALIZED
	VertexAttribArrayInteger       = C.GL_VERTEX_ATTRIB_ARRAY_INTEGER
	VertexAttribArrayDivisor       = C.GL_VERTEX_ATTRIB_ARRAY_DIVISOR
	CurrentVertexAttrib            = C.GL_CURRENT_VERTEX_ATTRIB

	// GetVertexAttribPointer
	VertexAttribArrayPointer = C.GL_VERTEX_ATTRIB_ARRAY_POINTER

	// LogicOp
	ClearConst   = C.GL_CLEAR
	Set          = C.GL_SET
	Copy         = C.GL_COPY
	CopyInverted = C.GL_COPY_INVERTED
	Noop         = C.GL_NOOP
	Invert       = C.GL_INVERT
	And          = C.GL_AND
	Nand         = C.GL_NAND
	Or           = C.GL_OR
	Nor          = C.GL_NOR
	Xor          = C.GL_XOR
	Equiv        = C.GL_EQUIV
	AndReverse   = C.GL_AND_REVERSE
	AndInverted  = C.GL_AND_INVERTED
	OrReverse    = C.GL_OR_REVERSE
	OrInverted   = C.GL_OR_INVERTED

	// MapBuffer
	ReadOnly  = C.GL_READ_ONLY
	WriteOnly = C.GL_WRITE_ONLY
	ReadWrite = C.GL_READ_WRITE
)

const (
	ColorBufferBit         GLbitfield = C.GL_COLOR_BUFFER_BIT
	DepthBufferBit                    = C.GL_DEPTH_BUFFER_BIT
	StencilBufferBit                  = C.GL_STENCIL_BUFFER_BIT
	SyncFlushCommandsBit              = C.GL_SYNC_FLUSH_COMMANDS_BIT
	MapReadBit                        = C.GL_MAP_READ_BIT
	MapWriteBit                       = C.GL_MAP_WRITE_BIT
	MapInvalidateRangeBit             = C.GL_MAP_INVALIDATE_RANGE_BIT
	MapInvalidateBufferBit            = C.GL_MAP_INVALIDATE_BUFFER_BIT
	MapFlushExplicitBit               = C.GL_MAP_FLUSH_EXPLICIT_BIT
	MapUnsynchronizedBit              = C.GL_MAP_UNSYNCHRONIZED_BIT
)
