package gl

// #include "common.h"
import "C"

type GLenum int
type GLbitfield int

const (
	// ActiveTexture
	Texture0Const                     GLenum = C.GL_TEXTURE0
	Texture1Const                            = C.GL_TEXTURE1
	Texture2Const                            = C.GL_TEXTURE2
	Texture3Const                            = C.GL_TEXTURE3
	Texture4Const                            = C.GL_TEXTURE4
	Texture5Const                            = C.GL_TEXTURE5
	Texture6Const                            = C.GL_TEXTURE6
	Texture7Const                            = C.GL_TEXTURE7
	Texture8Const                            = C.GL_TEXTURE8
	Texture9Const                            = C.GL_TEXTURE9
	Texture10Const                           = C.GL_TEXTURE10
	Texture11Const                           = C.GL_TEXTURE11
	Texture12Const                           = C.GL_TEXTURE12
	Texture13Const                           = C.GL_TEXTURE13
	Texture14Const                           = C.GL_TEXTURE14
	Texture15Const                           = C.GL_TEXTURE15
	Texture16Const                           = C.GL_TEXTURE16
	Texture17Const                           = C.GL_TEXTURE17
	Texture18Const                           = C.GL_TEXTURE18
	Texture19Const                           = C.GL_TEXTURE19
	Texture20Const                           = C.GL_TEXTURE20
	Texture21Const                           = C.GL_TEXTURE21
	Texture22Const                           = C.GL_TEXTURE22
	Texture23Const                           = C.GL_TEXTURE23
	Texture24Const                           = C.GL_TEXTURE24
	Texture25Const                           = C.GL_TEXTURE25
	Texture26Const                           = C.GL_TEXTURE26
	Texture27Const                           = C.GL_TEXTURE27
	Texture28Const                           = C.GL_TEXTURE28
	Texture29Const                           = C.GL_TEXTURE29
	Texture30Const                           = C.GL_TEXTURE30
	Texture31Const                           = C.GL_TEXTURE31
	MaxCombinedTextureImageUnitsConst        = C.GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS

	// BeginConditionalRender
	QueryByRegionWaitConst   = C.GL_QUERY_BY_REGION_WAIT
	QueryByRegionNoWaitConst = C.GL_QUERY_BY_REGION_NO_WAIT

	// BeginQuery
	SamplesPassedConst                      = C.GL_SAMPLES_PASSED
	AnySamplesPassedConst                   = C.GL_ANY_SAMPLES_PASSED
	PrimiticesGeneratedConst                = C.GL_PRIMITIVES_GENERATED
	TransformFeedbackPrimitivesWrittenConst = C.GL_TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN
	TimeElapsedConst                        = C.GL_TIME_ELAPSED

	// BeginTransformFeedback, DrawArrays, DrawArraysInstanced
	PointsConst                 = C.GL_POINTS
	LinesConst                  = C.GL_LINES
	LineLoopConst               = C.GL_LINE_LOOP
	LineStripConst              = C.GL_LINE_STRIP
	LinesAdjacencyConst         = C.GL_LINES_ADJACENCY
	LineStripAdjacencyConst     = C.GL_LINE_STRIP_ADJACENCY
	TrianglesConst              = C.GL_TRIANGLES
	TriangleStripConst          = C.GL_TRIANGLE_STRIP
	TriangleFanConst            = C.GL_TRIANGLE_FAN
	TrianglesAdjacencyConst     = C.GL_TRIANGLES_ADJACENCY
	TriangleStripAdjacencyConst = C.GL_TRIANGLE_STRIP_ADJACENCY

	// BindBuffer, BufferData, FlushMappedBufferRange
	ArrayBufferConst        = C.GL_ARRAY_BUFFER
	CopyReadBufferConst     = C.GL_COPY_READ_BUFFER
	CopyWriteBufferConst    = C.GL_COPY_WRITE_BUFFER
	ElementArrayBufferConst = C.GL_ELEMENT_ARRAY_BUFFER
	PixelPackBufferConst    = C.GL_PIXEL_PACK_BUFFER
	PixelUnpackBufferConst  = C.GL_PIXEL_UNPACK_BUFFER

	// BindBufferBase, BindBufferRange, FlushMappedBufferRange
	TransformFeedbackBufferConst = C.GL_TRANSFORM_FEEDBACK_BUFFER
	UniformBufferConst           = C.GL_UNIFORM_BUFFER

	// BindFramebuffer
	DrawFramebufferConst = C.GL_DRAW_FRAMEBUFFER
	ReadFramebufferConst = C.GL_READ_FRAMEBUFFER
	FramebufferConst     = C.GL_FRAMEBUFFER

	// BindRenderbuffer
	RenderbufferConst = C.GL_RENDERBUFFER

	// BindTexture, CompressedTexImage
	Texture1dConst                 = C.GL_TEXTURE_1D
	Texture2dConst                 = C.GL_TEXTURE_2D
	Texture3dConst                 = C.GL_TEXTURE_3D
	ProxyTexture1dConst            = C.GL_PROXY_TEXTURE_1D
	ProxyTexture2dConst            = C.GL_PROXY_TEXTURE_2D
	ProxyTexture3dConst            = C.GL_PROXY_TEXTURE_3D
	Texture1dArrayConst            = C.GL_TEXTURE_1D_ARRAY
	Texture2dArrayConst            = C.GL_TEXTURE_2D_ARRAY
	TextureRectangleConst          = C.GL_TEXTURE_RECTANGLE
	TextureCubeMapConst            = C.GL_TEXTURE_CUBE_MAP
	TextureBufferConst             = C.GL_TEXTURE_BUFFER
	Texture2dMultisampleConst      = C.GL_TEXTURE_2D_MULTISAMPLE
	Texture2dMultisampleArrayConst = C.GL_TEXTURE_2D_MULTISAMPLE_ARRAY
	ProxyTexture1dArrayConst       = C.GL_PROXY_TEXTURE_1D_ARRAY
	ProxyTexture2dArrayConst       = C.GL_PROXY_TEXTURE_2D_ARRAY
	TextureCubeMapPositiveXConst   = C.GL_TEXTURE_CUBE_MAP_POSITIVE_X
	TextureCubeMapNegativeXConst   = C.GL_TEXTURE_CUBE_MAP_NEGATIVE_X
	TextureCubeMapPositiveYConst   = C.GL_TEXTURE_CUBE_MAP_POSITIVE_Y
	TextureCubeMapNegativeYConst   = C.GL_TEXTURE_CUBE_MAP_NEGATIVE_Y
	TextureCubeMapPositiveZConst   = C.GL_TEXTURE_CUBE_MAP_POSITIVE_Z
	TextureCubeMapNegativeZConst   = C.GL_TEXTURE_CUBE_MAP_NEGATIVE_Z
	ProxyTextureCubeMapConst       = C.GL_PROXY_TEXTURE_CUBE_MAP

	// BlendEquation, BlendEquationSeparate
	FuncAddConst             = C.GL_FUNC_ADD
	FuncSubtractConst        = C.GL_FUNC_SUBTRACT
	FuncReverseSubtractConst = C.GL_FUNC_REVERSE_SUBTRACT
	MinConst                 = C.GL_MIN
	MaxConst                 = C.GL_MAX

	// BlendFunc
	ZeroConst                  = C.GL_ZERO
	OneConst                   = C.GL_ONE
	SrcColorConst              = C.GL_SRC_COLOR
	OneMinusSrcColorConst      = C.GL_ONE_MINUS_SRC_COLOR
	DstColorConst              = C.GL_DST_COLOR
	OneMinusDstColorConst      = C.GL_ONE_MINUS_DST_COLOR
	SrcAlphaConst              = C.GL_SRC_ALPHA
	OneMinusSrcAlphaConst      = C.GL_ONE_MINUS_SRC_ALPHA
	DstAlphaConst              = C.GL_DST_ALPHA
	OneMinusDstAlphaConst      = C.GL_ONE_MINUS_DST_ALPHA
	ConstantColorConst         = C.GL_CONSTANT_COLOR
	OneMinusConstantColorConst = C.GL_ONE_MINUS_CONSTANT_COLOR
	ConstantAlphaConst         = C.GL_CONSTANT_ALPHA
	OneMinusConstantAlphaConst = C.GL_ONE_MINUS_CONSTANT_ALPHA
	SrcAlphaSaturateConst      = C.GL_SRC_ALPHA_SATURATE
	Src1ColorConst             = C.GL_SRC1_COLOR
	OneMinusSrc1ColorConst     = C.GL_ONE_MINUS_SRC1_COLOR
	Src1AlphaConst             = C.GL_SRC1_ALPHA
	OneMinusSrc1AlphaConst     = C.GL_ONE_MINUS_SRC1_ALPHA

	// BlitFramebuffer
	NearestConst = C.GL_NEAREST
	LinearConst  = C.GL_LINEAR

	// BufferData
	StreamDrawConst  = C.GL_STREAM_DRAW
	StreamReadConst  = C.GL_STREAM_READ
	StreamCopyConst  = C.GL_STREAM_COPY
	StaticDrawConst  = C.GL_STATIC_DRAW
	StaticReadConst  = C.GL_STATIC_READ
	StaticCopyConst  = C.GL_STATIC_COPY
	DynamicDrawConst = C.GL_DYNAMIC_DRAW
	DynamicReadConst = C.GL_DYNAMIC_READ
	DynamicCopyConst = C.GL_DYNAMIC_COPY

	// CheckFramebufferStatus
	FramebufferUndefinedConst                   = C.GL_FRAMEBUFFER_UNDEFINED
	FramebufferIncompleteAttachmentConst        = C.GL_FRAMEBUFFER_INCOMPLETE_ATTACHMENT
	FramebufferIncompleteMissingAttachmentConst = C.GL_FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT
	FramebufferIncompleteDrawBufferConst        = C.GL_FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER
	FramebufferIncompleteReadBufferConst        = C.GL_FRAMEBUFFER_INCOMPLETE_READ_BUFFER
	FramebufferUnsupportedConst                 = C.GL_FRAMEBUFFER_UNSUPPORTED
	FramebufferIncompleteMultisampleConst       = C.GL_FRAMEBUFFER_INCOMPLETE_MULTISAMPLE
	FramebufferIncompleteLayerTargetsConst      = C.GL_FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS

	// ClampColor
	ClampReadColorConst = C.GL_CLAMP_READ_COLOR

	// CopyTexImage
	CompressedRedConst       = C.GL_COMPRESSED_RED
	CompressedRgConst        = C.GL_COMPRESSED_RG
	CompressedRgbConst       = C.GL_COMPRESSED_RGB
	CompressedRgbaConst      = C.GL_COMPRESSED_RGBA
	CompressedSrgbConst      = C.GL_COMPRESSED_SRGB
	CompressedSrgbAlphaConst = C.GL_COMPRESSED_SRGB_ALPHA
	DepthComponentConst      = C.GL_DEPTH_COMPONENT
	DepthComponent16Const    = C.GL_DEPTH_COMPONENT16
	DepthComponent24Const    = C.GL_DEPTH_COMPONENT24
	DepthComponent32Const    = C.GL_DEPTH_COMPONENT32
	RedConst                 = C.GL_RED
	RgConst                  = C.GL_RG
	RgbConst                 = C.GL_RGB
	R3G3B2Const              = C.GL_R3_G3_B2
	Rgb4Const                = C.GL_RGB4
	Rgb5Const                = C.GL_RGB5
	Rgb8Const                = C.GL_RGB8
	Rgb10Const               = C.GL_RGB10
	Rgb12Const               = C.GL_RGB12
	Rgb16Const               = C.GL_RGB16
	RgbaConst                = C.GL_RGBA
	Rgba2Const               = C.GL_RGBA2
	Rgba4Const               = C.GL_RGBA4
	Rgb5A1Const              = C.GL_RGB5_A1
	Rgba8Const               = C.GL_RGBA8
	Rgb10A2Const             = C.GL_RGB10_A2
	Rgba12Const              = C.GL_RGBA12
	Rgba16Const              = C.GL_RGBA16
	SrgbConst                = C.GL_SRGB
	Srgb8Const               = C.GL_SRGB8
	SrgbAlphaConst           = C.GL_SRGB_ALPHA
	Srgb8Alpha8Const         = C.GL_SRGB8_ALPHA8

	// CreateShader
	VertexShaderConst   = C.GL_VERTEX_SHADER
	GeometryShaderConst = C.GL_GEOMETRY_SHADER
	FragmentShaderConst = C.GL_FRAGMENT_SHADER

	// CullFace, DrawBuffer
	FrontConst        = C.GL_FRONT
	BackConst         = C.GL_BACK
	FrontAndBackConst = C.GL_FRONT_AND_BACK

	// DrawBuffer
	NoneConst       = C.GL_NONE
	FrontLeftConst  = C.GL_FRONT_LEFT
	FrontRightConst = C.GL_FRONT_RIGHT
	BackLeftConst   = C.GL_BACK_LEFT
	BackRightConst  = C.GL_BACK_RIGHT
	LeftConst       = C.GL_LEFT
	RightConst      = C.GL_RIGHT

	// DepthFunc
	NeverConst    = C.GL_NEVER
	LessConst     = C.GL_LESS
	EqualConst    = C.GL_EQUAL
	LEqualConst   = C.GL_LEQUAL
	GreaterConst  = C.GL_GREATER
	NotEqualConst = C.GL_NOTEQUAL
	GEqualConst   = C.GL_EQUAL
	AlwaysConst   = C.GL_ALWAYS

	// Disable, Enable
	BlendConst                  = C.GL_BLEND
	ClipDistance0Const          = C.GL_CLIP_DISTANCE0
	ClipDistance1Const          = C.GL_CLIP_DISTANCE1
	ClipDistance2Const          = C.GL_CLIP_DISTANCE2
	ClipDistance3Const          = C.GL_CLIP_DISTANCE3
	ClipDistance4Const          = C.GL_CLIP_DISTANCE4
	ClipDistance5Const          = C.GL_CLIP_DISTANCE5
	ColorLogicOpConst           = C.GL_COLOR_LOGIC_OP
	CullFaceConst               = C.GL_CULL_FACE
	DepthClampConst             = C.GL_DEPTH_CLAMP
	DepthTestConst              = C.GL_DEPTH_TEST
	DitherConst                 = C.GL_DITHER
	FramebufferSrgbConst        = C.GL_FRAMEBUFFER_SRGB
	LineSmoothConst             = C.GL_LINE_SMOOTH
	MultisampleConst            = C.GL_MULTISAMPLE
	PolygonOffsetFillConst      = C.GL_POLYGON_OFFSET_FILL
	PolygonOffsetLineConst      = C.GL_POLYGON_OFFSET_LINE
	PolygonOffsetPointConst     = C.GL_POLYGON_OFFSET_POINT
	PolygonSmoothConst          = C.GL_POLYGON_SMOOTH
	PrimitiveRestartConst       = C.GL_PRIMITIVE_RESTART
	SampleAlphaToCoverageConst  = C.GL_SAMPLE_ALPHA_TO_COVERAGE
	SampleAlphaToOneConst       = C.GL_SAMPLE_ALPHA_TO_ONE
	SampleCoverageConst         = C.GL_SAMPLE_COVERAGE
	ScissorTestConst            = C.GL_SCISSOR_TEST
	StencilTestConst            = C.GL_STENCIL_TEST
	TextureCubeMapSeamlessConst = C.GL_TEXTURE_CUBE_MAP_SEAMLESS
	ProgramPointSizeConst       = C.GL_PROGRAM_POINT_SIZE

	// FenceSync
	SyncGpuCommandsCompleteConst = C.GL_SYNC_GPU_COMMANDS_COMPLETE

	// glFramebufferTexture, glFramebufferTextureLayer
	ColorAttachment0Const       = C.GL_COLOR_ATTACHMENT0
	ColorAttachment1Const       = C.GL_COLOR_ATTACHMENT1
	ColorAttachment2Const       = C.GL_COLOR_ATTACHMENT2
	ColorAttachment3Const       = C.GL_COLOR_ATTACHMENT3
	ColorAttachment4Const       = C.GL_COLOR_ATTACHMENT4
	ColorAttachment5Const       = C.GL_COLOR_ATTACHMENT5
	ColorAttachment6Const       = C.GL_COLOR_ATTACHMENT6
	ColorAttachment7Const       = C.GL_COLOR_ATTACHMENT7
	ColorAttachment8Const       = C.GL_COLOR_ATTACHMENT8
	ColorAttachment9Const       = C.GL_COLOR_ATTACHMENT9
	ColorAttachment10Const      = C.GL_COLOR_ATTACHMENT10
	ColorAttachment11Const      = C.GL_COLOR_ATTACHMENT11
	ColorAttachment12Const      = C.GL_COLOR_ATTACHMENT12
	ColorAttachment13Const      = C.GL_COLOR_ATTACHMENT13
	ColorAttachment14Const      = C.GL_COLOR_ATTACHMENT14
	ColorAttachment15Const      = C.GL_COLOR_ATTACHMENT15
	DepthAttachmentConst        = C.GL_DEPTH_ATTACHMENT
	StencilAttachmentConst      = C.GL_STENCIL_ATTACHMENT
	DepthStencilAttachmentConst = C.GL_DEPTH_STENCIL_ATTACHMENT

	// FrontFace
	CWConst  = C.GL_CW
	CCWConst = C.GL_CCW

	// Everything else
	ActiveTextureConst                        = C.GL_ACTIVE_TEXTURE
	AliasedLineWidthRangeConst                = C.GL_ALIASED_LINE_WIDTH_RANGE
	SmoothLineWidthRangeConst                 = C.GL_SMOOTH_LINE_WIDTH_RANGE
	SmoothLineWidthGranularityConst           = C.GL_SMOOTH_LINE_WIDTH_GRANULARITY
	ArrayBufferBindingConst                   = C.GL_ARRAY_BUFFER_BINDING
	BlendColorConst                           = C.GL_BLEND_COLOR
	BlendDstAlphaConst                        = C.GL_BLEND_DST_ALPHA
	BlendDstRgbConst                          = C.GL_BLEND_DST_RGB
	BlendEquationRgbConst                     = C.GL_BLEND_EQUATION_RGB
	BlendEquationAlphaConst                   = C.GL_BLEND_EQUATION_ALPHA
	BlendSrcAlphaConst                        = C.GL_BLEND_SRC_ALPHA
	BlendSrcRgbConst                          = C.GL_BLEND_SRC_RGB
	ColorClearValueConst                      = C.GL_COLOR_CLEAR_VALUE
	ColorWritemaskConst                       = C.GL_COLOR_WRITEMASK
	CompressedTextureFormatsConst             = C.GL_COMPRESSED_TEXTURE_FORMATS
	CurrentProgramConst                       = C.GL_CURRENT_PROGRAM
	DepthClearValueConst                      = C.GL_DEPTH_CLEAR_VALUE
	DepthFuncConst                            = C.GL_DEPTH_FUNC
	DepthRangeConst                           = C.GL_DEPTH_RANGE
	DepthWritemaskConst                       = C.GL_DEPTH_WRITEMASK
	DoublebufferConst                         = C.GL_DOUBLEBUFFER
	DrawBufferConst                           = C.GL_DRAW_BUFFER
	DrawBuffer0Const                          = C.GL_DRAW_BUFFER0
	DrawBuffer1Const                          = C.GL_DRAW_BUFFER1
	DrawBuffer2Const                          = C.GL_DRAW_BUFFER2
	DrawBuffer3Const                          = C.GL_DRAW_BUFFER3
	DrawBuffer4Const                          = C.GL_DRAW_BUFFER4
	DrawBuffer5Const                          = C.GL_DRAW_BUFFER5
	DrawBuffer6Const                          = C.GL_DRAW_BUFFER6
	DrawBuffer7Const                          = C.GL_DRAW_BUFFER7
	DrawBuffer8Const                          = C.GL_DRAW_BUFFER8
	DrawBuffer9Const                          = C.GL_DRAW_BUFFER9
	DrawBuffer10Const                         = C.GL_DRAW_BUFFER10
	DrawBuffer11Const                         = C.GL_DRAW_BUFFER11
	DrawBuffer12Const                         = C.GL_DRAW_BUFFER12
	DrawBuffer13Const                         = C.GL_DRAW_BUFFER13
	DrawBuffer14Const                         = C.GL_DRAW_BUFFER14
	DrawBuffer15Const                         = C.GL_DRAW_BUFFER15
	DrawFramebufferBindingConst               = C.GL_DRAW_FRAMEBUFFER_BINDING
	ReadFramebufferBindingConst               = C.GL_READ_FRAMEBUFFER_BINDING
	ElementArrayBufferBindingConst            = C.GL_ELEMENT_ARRAY_BUFFER_BINDING
	RenderbufferBindingConst                  = C.GL_RENDERBUFFER_BINDING
	FragmentShaderDerivativeHintConst         = C.GL_FRAGMENT_SHADER_DERIVATIVE_HINT
	LineSmoothHintConst                       = C.GL_LINE_SMOOTH_HINT
	LineWidthConst                            = C.GL_LINE_WIDTH
	LogicOpModeConst                          = C.GL_LOGIC_OP_MODE
	Max3dTextureSizeConst                     = C.GL_MAX_3D_TEXTURE_SIZE
	MaxClipDistancesConst                     = C.GL_MAX_CLIP_DISTANCES
	MaxCombinedFragmentUniformComponentsConst = C.GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS
	MaxCombinedVertexUniformComponentsConst   = C.GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS
	MaxCombinedGeometryUniformComponentsConst = C.GL_MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS
	MaxVaryingComponentsConst                 = C.GL_MAX_VARYING_COMPONENTS
	MaxCombinedUniformBlocksConst             = C.GL_MAX_COMBINED_UNIFORM_BLOCKS
	MaxCubeMapTextureSizeConst                = C.GL_MAX_CUBE_MAP_TEXTURE_SIZE
	MaxDrawBuffersConst                       = C.GL_MAX_DRAW_BUFFERS
	MaxDualSourceDrawBuffersConst             = C.GL_MAX_DUAL_SOURCE_DRAW_BUFFERS
	MaxElementsIndicesConst                   = C.GL_MAX_ELEMENTS_INDICES
	MaxElementsVerticesConst                  = C.GL_MAX_ELEMENTS_VERTICES
	MaxFragmentUniformComponentsConst         = C.GL_MAX_FRAGMENT_UNIFORM_COMPONENTS
	MaxFragmentUniformBlocksConst             = C.GL_MAX_FRAGMENT_UNIFORM_BLOCKS
	MaxFragmentInputComponentsConst           = C.GL_MAX_FRAGMENT_INPUT_COMPONENTS
	MinProgramTexelOffsetConst                = C.GL_MIN_PROGRAM_TEXEL_OFFSET
	MaxProgramTexelOffsetConst                = C.GL_MAX_PROGRAM_TEXEL_OFFSET
	MaxRectangleTextureSizeConst              = C.GL_MAX_RECTANGLE_TEXTURE_SIZE
	MaxTextureImageUnitsConst                 = C.GL_MAX_TEXTURE_IMAGE_UNITS
	MaxTextureLodBiasConst                    = C.GL_MAX_TEXTURE_LOD_BIAS
	MaxTextureSizeConst                       = C.GL_MAX_TEXTURE_SIZE
	MaxRenderbufferSizeConst                  = C.GL_MAX_RENDERBUFFER_SIZE
	MaxArrayTextureLayersConst                = C.GL_MAX_ARRAY_TEXTURE_LAYERS
	MaxTextureBufferSizeConst                 = C.GL_MAX_TEXTURE_BUFFER_SIZE
	MaxUniformBlockSizeConst                  = C.GL_MAX_UNIFORM_BLOCK_SIZE
	MaxVaryingFloatsConst                     = C.GL_MAX_VARYING_FLOATS
	MaxVertexAttribsConst                     = C.GL_MAX_VERTEX_ATTRIBS
	MaxVertexTextureImageUnitsConst           = C.GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS
	MaxGeometryTextureImageUnitsConst         = C.GL_MAX_GEOMETRY_TEXTURE_IMAGE_UNITS
	MaxVertexUniformComponentsConst           = C.GL_MAX_VERTEX_UNIFORM_COMPONENTS
	MaxVertexOutputComponentsConst            = C.GL_MAX_VERTEX_OUTPUT_COMPONENTS
	MaxGeometryUniformComponentsConst         = C.GL_MAX_GEOMETRY_UNIFORM_COMPONENTS
	MaxSampleMaskWordsConst                   = C.GL_MAX_SAMPLE_MASK_WORDS
	MaxColorTextureSamplesConst               = C.GL_MAX_COLOR_TEXTURE_SAMPLES
	MaxDepthTextureSamplesConst               = C.GL_MAX_DEPTH_TEXTURE_SAMPLES
	MaxIntegerSamplesConst                    = C.GL_MAX_INTEGER_SAMPLES
	MaxServerWaitTimeoutConst                 = C.GL_MAX_SERVER_WAIT_TIMEOUT
	MaxUniformBufferBindingsConst             = C.GL_MAX_UNIFORM_BUFFER_BINDINGS
	UniformBufferOffsetAlignmentConst         = C.GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT
	MaxVertexUniformBlocksConst               = C.GL_MAX_VERTEX_UNIFORM_BLOCKS
	MaxGeometryUniformBlocksConst             = C.GL_MAX_GEOMETRY_UNIFORM_BLOCKS
	MaxGeometryInputComponentsConst           = C.GL_MAX_GEOMETRY_INPUT_COMPONENTS
	MaxGeometryOutputComponentsConst          = C.GL_MAX_GEOMETRY_OUTPUT_COMPONENTS
	MaxViewportDimsConst                      = C.GL_MAX_VIEWPORT_DIMS
	NumCompressedTextureFormatsConst          = C.GL_NUM_COMPRESSED_TEXTURE_FORMATS
	PackAlignmentConst                        = C.GL_PACK_ALIGNMENT
	PackImageHeightConst                      = C.GL_PACK_IMAGE_HEIGHT
	PackLsbFirstConst                         = C.GL_PACK_LSB_FIRST
	PackRowLengthConst                        = C.GL_PACK_ROW_LENGTH
	PackSkipImagesConst                       = C.GL_PACK_SKIP_IMAGES
	PackSkipPixelsConst                       = C.GL_PACK_SKIP_PIXELS
	PackSkipRowsConst                         = C.GL_PACK_SKIP_ROWS
	PackSwapBytesConst                        = C.GL_PACK_SWAP_BYTES
	PixelPackBufferBindingConst               = C.GL_PIXEL_PACK_BUFFER_BINDING
	PixelUnpackBufferBindingConst             = C.GL_PIXEL_UNPACK_BUFFER_BINDING
	PointFadeThresholdSizeConst               = C.GL_POINT_FADE_THRESHOLD_SIZE
	PrimitiveRestartIndexConst                = C.GL_PRIMITIVE_RESTART_INDEX
	ProvokingVertexConst                      = C.GL_PROVOKING_VERTEX
	PointSizeConst                            = C.GL_POINT_SIZE
	PointSizeGranularityConst                 = C.GL_POINT_SIZE_GRANULARITY
	PointSizeRangeConst                       = C.GL_POINT_SIZE_RANGE
	PolygonOffsetFactorConst                  = C.GL_POLYGON_OFFSET_FACTOR
	PolygonOffsetUnitsConst                   = C.GL_POLYGON_OFFSET_UNITS
	PolygonSmoothHintConst                    = C.GL_POLYGON_SMOOTH_HINT
	ReadBufferConst                           = C.GL_READ_BUFFER
	SampleBuffersConst                        = C.GL_SAMPLE_BUFFERS
	SampleCoverageValueConst                  = C.GL_SAMPLE_COVERAGE_VALUE
	SampleCoverageInvertConst                 = C.GL_SAMPLE_COVERAGE_INVERT
	SamplerBindingConst                       = C.GL_SAMPLER_BINDING
	SamplesConst                              = C.GL_SAMPLES
	ScissorBoxConst                           = C.GL_SCISSOR_BOX
	StencilBackFailConst                      = C.GL_STENCIL_BACK_FAIL
	StencilBackFuncConst                      = C.GL_STENCIL_BACK_FUNC
	StencilBackPassDepthFailConst             = C.GL_STENCIL_BACK_PASS_DEPTH_FAIL
	StencilBackPassDepthPassConst             = C.GL_STENCIL_BACK_PASS_DEPTH_PASS
	StencilBackRefConst                       = C.GL_STENCIL_BACK_REF
	StencilBackValueMaskConst                 = C.GL_STENCIL_BACK_VALUE_MASK
	StencilBackWritemaskConst                 = C.GL_STENCIL_BACK_WRITEMASK
	StencilClearValueConst                    = C.GL_STENCIL_CLEAR_VALUE
	StencilFailConst                          = C.GL_STENCIL_FAIL
	StencilFuncConst                          = C.GL_STENCIL_FUNC
	StencilPassDepthFailConst                 = C.GL_STENCIL_PASS_DEPTH_FAIL
	StencilPassDepthPassConst                 = C.GL_STENCIL_PASS_DEPTH_PASS
	StencilRefConst                           = C.GL_STENCIL_REF
	StencilValueMaskConst                     = C.GL_STENCIL_VALUE_MASK
	StencilWritemaskConst                     = C.GL_STENCIL_WRITEMASK
	StereoConst                               = C.GL_STEREO
	SubpixelBitsConst                         = C.GL_SUBPIXEL_BITS
	TextureBinding1dConst                     = C.GL_TEXTURE_BINDING_1D
	TextureBinding1dArrayConst                = C.GL_TEXTURE_BINDING_1D_ARRAY
	TextureBinding2dConst                     = C.GL_TEXTURE_BINDING_2D
	TextureBinding2dArrayConst                = C.GL_TEXTURE_BINDING_2D_ARRAY
	TextureBinding2dMultisampleConst          = C.GL_TEXTURE_BINDING_2D_MULTISAMPLE
	TextureBinding2dMultisampleArrayConst     = C.GL_TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY
	TextureBinding3dConst                     = C.GL_TEXTURE_BINDING_3D
	TextureBindingBufferConst                 = C.GL_TEXTURE_BINDING_BUFFER
	TextureBindingCubeMapConst                = C.GL_TEXTURE_BINDING_CUBE_MAP
	TextureBindingRectangleConst              = C.GL_TEXTURE_BINDING_RECTANGLE
	TextureCompressionHintConst               = C.GL_TEXTURE_COMPRESSION_HINT
	TimestampConst                            = C.GL_TIMESTAMP
	TransformFeedbackBufferBindingConst       = C.GL_TRANSFORM_FEEDBACK_BUFFER_BINDING
	TransformFeedbackBufferStartConst         = C.GL_TRANSFORM_FEEDBACK_BUFFER_START
	TransformFeedbackBufferSizeConst          = C.GL_TRANSFORM_FEEDBACK_BUFFER_SIZE
	UniformBufferBindingConst                 = C.GL_UNIFORM_BUFFER_BINDING
	UniformBufferStartConst                   = C.GL_UNIFORM_BUFFER_START
	UniformBufferSizeConst                    = C.GL_UNIFORM_BUFFER_SIZE
	UnpackAlignmentConst                      = C.GL_UNPACK_ALIGNMENT
	UnpackImageHeightConst                    = C.GL_UNPACK_IMAGE_HEIGHT
	UnpackLsbFirstConst                       = C.GL_UNPACK_LSB_FIRST
	UnpackRowLengthConst                      = C.GL_UNPACK_ROW_LENGTH
	UnpackSkipImagesConst                     = C.GL_UNPACK_SKIP_IMAGES
	UnpackSkipPixelsConst                     = C.GL_UNPACK_SKIP_PIXELS
	UnpackSkipRowsConst                       = C.GL_UNPACK_SKIP_ROWS
	UnpackSwapBytesConst                      = C.GL_UNPACK_SWAP_BYTES
	NumExtensionsConst                        = C.GL_NUM_EXTENSIONS
	MajorVersionConst                         = C.GL_MAJOR_VERSION
	MinorVersionConst                         = C.GL_MINOR_VERSION
	ContextFlagsConst                         = C.GL_CONTEXT_FLAGS
	ViewportConst                             = C.GL_VIEWPORT

	// GetActiveAttrib, GetActiveUniform, etc
	FloatConst                                = C.GL_FLOAT
	FloatVec2Const                            = C.GL_FLOAT_VEC2
	FloatVec3Const                            = C.GL_FLOAT_VEC3
	FloatVec4Const                            = C.GL_FLOAT_VEC4
	IntConst                                  = C.GL_INT
	IntVec2Const                              = C.GL_INT_VEC2
	IntVec3Const                              = C.GL_INT_VEC3
	IntVec4Const                              = C.GL_INT_VEC4
	UnsignedIntConst                          = C.GL_UNSIGNED_INT
	UnsignedIntVec2Const                      = C.GL_UNSIGNED_INT_VEC2
	UnsignedIntVec3Const                      = C.GL_UNSIGNED_INT_VEC3
	UnsignedIntVec4Const                      = C.GL_UNSIGNED_INT_VEC4
	BoolConst                                 = C.GL_BOOL
	BoolVec2Const                             = C.GL_BOOL_VEC2
	BoolVec3Const                             = C.GL_BOOL_VEC3
	BoolVec4Const                             = C.GL_BOOL_VEC4
	FloatMat2Const                            = C.GL_FLOAT_MAT2
	FloatMat3Const                            = C.GL_FLOAT_MAT3
	FloatMat4Const                            = C.GL_FLOAT_MAT4
	FLOATMAT2x3Const                          = C.GL_FLOAT_MAT2x3
	FLOATMAT2x4Const                          = C.GL_FLOAT_MAT2x4
	FLOATMAT3x2Const                          = C.GL_FLOAT_MAT3x2
	FLOATMAT3x4Const                          = C.GL_FLOAT_MAT3x4
	FLOATMAT4x2Const                          = C.GL_FLOAT_MAT4x2
	FLOATMAT4x3Const                          = C.GL_FLOAT_MAT4x3
	Sampler1dConst                            = C.GL_SAMPLER_1D
	Sampler2dConst                            = C.GL_SAMPLER_2D
	Sampler3dConst                            = C.GL_SAMPLER_3D
	SamplerCubeConst                          = C.GL_SAMPLER_CUBE
	Sampler1dShadowConst                      = C.GL_SAMPLER_1D_SHADOW
	Sampler2dShadowConst                      = C.GL_SAMPLER_2D_SHADOW
	Sampler1dArrayConst                       = C.GL_SAMPLER_1D_ARRAY
	Sampler2dArrayConst                       = C.GL_SAMPLER_2D_ARRAY
	Sampler1dArrayShadowConst                 = C.GL_SAMPLER_1D_ARRAY_SHADOW
	Sampler2dArrayShadowConst                 = C.GL_SAMPLER_2D_ARRAY_SHADOW
	Sampler2dMultisampleConst                 = C.GL_SAMPLER_2D_MULTISAMPLE
	Sampler2dMultisampleArrayConst            = C.GL_SAMPLER_2D_MULTISAMPLE_ARRAY
	SamplerCubeShadowConst                    = C.GL_SAMPLER_CUBE_SHADOW
	SamplerBufferConst                        = C.GL_SAMPLER_BUFFER
	Sampler2dRectConst                        = C.GL_SAMPLER_2D_RECT
	Sampler2dRectShadowConst                  = C.GL_SAMPLER_2D_RECT_SHADOW
	IntSampler1dConst                         = C.GL_INT_SAMPLER_1D
	IntSampler2dConst                         = C.GL_INT_SAMPLER_2D
	IntSampler3dConst                         = C.GL_INT_SAMPLER_3D
	IntSamplerCubeConst                       = C.GL_INT_SAMPLER_CUBE
	IntSampler1dArrayConst                    = C.GL_INT_SAMPLER_1D_ARRAY
	IntSampler2dArrayConst                    = C.GL_INT_SAMPLER_2D_ARRAY
	IntSampler2dMultisampleConst              = C.GL_INT_SAMPLER_2D_MULTISAMPLE
	IntSampler2dMultisampleArrayConst         = C.GL_INT_SAMPLER_2D_MULTISAMPLE_ARRAY
	IntSamplerBufferConst                     = C.GL_INT_SAMPLER_BUFFER
	IntSampler2dRectConst                     = C.GL_INT_SAMPLER_2D_RECT
	UnsignedIntSampler1dConst                 = C.GL_UNSIGNED_INT_SAMPLER_1D
	UnsignedIntSampler2dConst                 = C.GL_UNSIGNED_INT_SAMPLER_2D
	UnsignedIntSampler3dConst                 = C.GL_UNSIGNED_INT_SAMPLER_3D
	UnsignedIntSamplerCubeConst               = C.GL_UNSIGNED_INT_SAMPLER_CUBE
	UnsignedIntSampler1dArrayConst            = C.GL_UNSIGNED_INT_SAMPLER_1D_ARRAY
	UnsignedIntSampler2dArrayConst            = C.GL_UNSIGNED_INT_SAMPLER_2D_ARRAY
	UnsignedIntSampler2dMultisampleConst      = C.GL_UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE
	UnsignedIntSampler2dMultisampleArrayConst = C.GL_UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE_ARRAY
	UnsignedIntSamplerBufferConst             = C.GL_UNSIGNED_INT_SAMPLER_BUFFER
	UnsignedIntSampler2dRectConst             = C.GL_UNSIGNED_INT_SAMPLER_2D_RECT

	// GetActiveUniformBlock, etc
	UniformBlockBindingConst                    = C.GL_UNIFORM_BLOCK_BINDING
	UniformBlockDataSizeConst                   = C.GL_UNIFORM_BLOCK_DATA_SIZE
	UniformBlockNameLengthConst                 = C.GL_UNIFORM_BLOCK_NAME_LENGTH
	UniformBlockActiveUniformsConst             = C.GL_UNIFORM_BLOCK_ACTIVE_UNIFORMS
	UniformBlockActiveUniformIndicesConst       = C.GL_UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES
	UniformBlockReferencedByVertexShaderConst   = C.GL_UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER
	UniformBlockReferencedByGeometryShaderConst = C.GL_UNIFORM_BLOCK_REFERENCED_BY_GEOMETRY_SHADER
	UniformBlockReferencedByFragmentShaderConst = C.GL_UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER

	// GetBufferParameter, GetBufferPointer, etc
	BufferAccessConst = C.GL_BUFFER_ACCESS
	BufferMappedConst = C.GL_BUFFER_MAPPED
	BufferSizeConst   = C.GL_BUFFER_SIZE
	BufferUsageConst  = C.GL_BUFFER_USAGE

	// GetError
	NoErrorConst                     = C.GL_NO_ERROR
	InvalidEnumConst                 = C.GL_INVALID_ENUM
	InvalidValueConst                = C.GL_INVALID_VALUE
	InvalidOperationConst            = C.GL_INVALID_OPERATION
	InvalidFramebufferOperationConst = C.GL_INVALID_FRAMEBUFFER_OPERATION
	OutOfMemoryConst                 = C.GL_OUT_OF_MEMORY

	// GetMultisample
	SamplePositionConst = C.GL_SAMPLE_POSITION

	// GetProgram
	DeleteStatusConst                      = C.GL_DELETE_STATUS
	LinkStatusConst                        = C.GL_LINK_STATUS
	ValidateStatusConst                    = C.GL_VALIDATE_STATUS
	InfoLogLengthConst                     = C.GL_INFO_LOG_LENGTH
	AttachedShadersConst                   = C.GL_ATTACHED_SHADERS
	ActiveAttributesConst                  = C.GL_ACTIVE_ATTRIBUTES
	ActiveAttributeMaxLengthConst          = C.GL_ACTIVE_ATTRIBUTE_MAX_LENGTH
	ActiveUniformsConst                    = C.GL_ACTIVE_UNIFORMS
	ActiveUniformBlocksConst               = C.GL_ACTIVE_UNIFORM_BLOCKS
	ActiveUniformBlockMaxNameLengthConst   = C.GL_ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH
	ActiveUniformMaxLengthConst            = C.GL_ACTIVE_UNIFORM_MAX_LENGTH
	TransformFeedbackBufferModeConst       = C.GL_TRANSFORM_FEEDBACK_BUFFER_MODE
	TransformFeedbackVaryingsConst         = C.GL_TRANSFORM_FEEDBACK_VARYINGS
	TransformFeedbackVaryingMaxLengthConst = C.GL_TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH
	GeometryVerticesOutConst               = C.GL_GEOMETRY_VERTICES_OUT
	GeometryInputTypeConst                 = C.GL_GEOMETRY_INPUT_TYPE
	GeometryOutputTypeConst                = C.GL_GEOMETRY_OUTPUT_TYPE

	// GetQueryObject
	QueryResultConst          = C.GL_QUERY_RESULT
	QueryResultAvailableConst = C.GL_QUERY_RESULT_AVAILABLE

	// GetQuery
	PrimitivesGeneratedConst = C.GL_PRIMITIVES_GENERATED
	CurrentQueryConst        = C.GL_CURRENT_QUERY
	QueryCounterBitsConst    = C.GL_QUERY_COUNTER_BITS

	// GetSamplerParameter
	TextureMagFilterConst   = C.GL_TEXTURE_MAG_FILTER
	TextureMinFilterConst   = C.GL_TEXTURE_MIN_FILTER
	TextureMinLodConst      = C.GL_TEXTURE_MIN_LOD
	TextureMaxLodConst      = C.GL_TEXTURE_MAX_LOD
	TextureLodBiasConst     = C.GL_TEXTURE_LOD_BIAS
	TextureWrapSConst       = C.GL_TEXTURE_WRAP_S
	TextureWrapTConst       = C.GL_TEXTURE_WRAP_T
	TextureWrapRConst       = C.GL_TEXTURE_WRAP_R
	TextureBorderColorConst = C.GL_TEXTURE_BORDER_COLOR
	TextureCompareModeConst = C.GL_TEXTURE_COMPARE_MODE
	TextureCompareFuncConst = C.GL_TEXTURE_COMPARE_FUNC

	// GetShader
	ShaderTypeConst         = C.GL_SHADER_TYPE
	CompileStatusConst      = C.GL_COMPILE_STATUS
	ShaderSourceLengthConst = C.GL_SHADER_SOURCE_LENGTH

	// GetString
	VendorConst                 = C.GL_VENDOR
	RendererConst               = C.GL_RENDERER
	VersionConst                = C.GL_VERSION
	ShadingLanguageVersionConst = C.GL_SHADING_LANGUAGE_VERSION
	ExtensionsConst             = C.GL_EXTENSIONS

	// GetTexImage
	StencilIndexConst             = C.GL_STENCIL_INDEX
	DepthStencilConst             = C.GL_DEPTH_STENCIL
	GreenConst                    = C.GL_GREEN
	BlueConst                     = C.GL_BLUE
	BgrConst                      = C.GL_BGR
	BgraConst                     = C.GL_BGRA
	RedIntegerConst               = C.GL_RED_INTEGER
	GreenIntegerConst             = C.GL_GREEN_INTEGER
	BlueIntegerConst              = C.GL_BLUE_INTEGER
	RgIntegerConst                = C.GL_RG_INTEGER
	RgbIntegerConst               = C.GL_RGB_INTEGER
	RgbaIntegerConst              = C.GL_RGBA_INTEGER
	BgrIntegerConst               = C.GL_BGR_INTEGER
	BgraIntegerConst              = C.GL_BGRA_INTEGER
	UnsignedByteConst             = C.GL_UNSIGNED_BYTE
	ByteConst                     = C.GL_BYTE
	UnsignedShortConst            = C.GL_UNSIGNED_SHORT
	ShortConst                    = C.GL_SHORT
	HalfFloatConst                = C.GL_HALF_FLOAT
	UnsignedByte332Const          = C.GL_UNSIGNED_BYTE_3_3_2
	UnsignedByte233RevConst       = C.GL_UNSIGNED_BYTE_2_3_3_REV
	UnsignedShort565Const         = C.GL_UNSIGNED_SHORT_5_6_5
	UnsignedShort565RevConst      = C.GL_UNSIGNED_SHORT_5_6_5_REV
	UnsignedShort4444Const        = C.GL_UNSIGNED_SHORT_4_4_4_4
	UnsignedShort4444RevConst     = C.GL_UNSIGNED_SHORT_4_4_4_4_REV
	UnsignedShort5551Const        = C.GL_UNSIGNED_SHORT_5_5_5_1
	UnsignedShort1555RevConst     = C.GL_UNSIGNED_SHORT_1_5_5_5_REV
	UnsignedInt8888Const          = C.GL_UNSIGNED_INT_8_8_8_8
	UnsignedInt8888RevConst       = C.GL_UNSIGNED_INT_8_8_8_8_REV
	UnsignedInt1010102Const       = C.GL_UNSIGNED_INT_10_10_10_2
	UnsignedInt2101010RevConst    = C.GL_UNSIGNED_INT_2_10_10_10_REV
	UnsignedInt248Const           = C.GL_UNSIGNED_INT_24_8
	UnsignedInt10f11f11fRevConst  = C.GL_UNSIGNED_INT_10F_11F_11F_REV
	UnsignedInt5999RevConst       = C.GL_UNSIGNED_INT_5_9_9_9_REV
	Float32UnsignedInt248RevConst = C.GL_FLOAT_32_UNSIGNED_INT_24_8_REV

	// GetTexLevelParameter
	TextureWidthConst               = C.GL_TEXTURE_WIDTH
	TextureHeightConst              = C.GL_TEXTURE_HEIGHT
	TextureDepthConst               = C.GL_TEXTURE_DEPTH
	TextureInternalFormatConst      = C.GL_TEXTURE_INTERNAL_FORMAT
	TextureRedSizeConst             = C.GL_TEXTURE_RED_SIZE
	TextureGreenSizeConst           = C.GL_TEXTURE_GREEN_SIZE
	TextureBlueSizeConst            = C.GL_TEXTURE_BLUE_SIZE
	TextureAlphaSizeConst           = C.GL_TEXTURE_ALPHA_SIZE
	TextureDepthSizeConst           = C.GL_TEXTURE_DEPTH_SIZE
	TextureCompressedConst          = C.GL_TEXTURE_COMPRESSED
	TextureCompressedImageSizeConst = C.GL_TEXTURE_COMPRESSED_IMAGE_SIZE

	// GetTexParameter
	TextureBaseLevelConst   = C.GL_TEXTURE_BASE_LEVEL
	TextureMaxLevelConst    = C.GL_TEXTURE_MAX_LEVEL
	TextureSwizzleRConst    = C.GL_TEXTURE_SWIZZLE_R
	TextureSwizzleGConst    = C.GL_TEXTURE_SWIZZLE_G
	TextureSwizzleBConst    = C.GL_TEXTURE_SWIZZLE_B
	TextureSwizzleAConst    = C.GL_TEXTURE_SWIZZLE_A
	TextureSwizzleRgbaConst = C.GL_TEXTURE_SWIZZLE_RGBA

	// GetVertexAttrib
	VertexAttribArrayBufferBindingConst = C.GL_VERTEX_ATTRIB_ARRAY_BUFFER_BINDING
	VertexAttribArrayEnabledConst       = C.GL_VERTEX_ATTRIB_ARRAY_ENABLED
	VertexAttribArraySizeConst          = C.GL_VERTEX_ATTRIB_ARRAY_SIZE
	VertexAttribArrayStrideConst        = C.GL_VERTEX_ATTRIB_ARRAY_STRIDE
	VertexAttribArrayTypeConst          = C.GL_VERTEX_ATTRIB_ARRAY_TYPE
	VertexAttribArrayNormalizedConst    = C.GL_VERTEX_ATTRIB_ARRAY_NORMALIZED
	VertexAttribArrayIntegerConst       = C.GL_VERTEX_ATTRIB_ARRAY_INTEGER
	VertexAttribArrayDivisorConst       = C.GL_VERTEX_ATTRIB_ARRAY_DIVISOR
	CurrentVertexAttribConst            = C.GL_CURRENT_VERTEX_ATTRIB

	// GetVertexAttribPointer
	VertexAttribArrayPointerConst = C.GL_VERTEX_ATTRIB_ARRAY_POINTER

	// LogicOp
	ClearConst        = C.GL_CLEAR
	SetConst          = C.GL_SET
	CopyConst         = C.GL_COPY
	CopyInvertedConst = C.GL_COPY_INVERTED
	NoopConst         = C.GL_NOOP
	InvertConst       = C.GL_INVERT
	AndConst          = C.GL_AND
	NandConst         = C.GL_NAND
	OrConst           = C.GL_OR
	NorConst          = C.GL_NOR
	XorConst          = C.GL_XOR
	EquivConst        = C.GL_EQUIV
	AndReverseConst   = C.GL_AND_REVERSE
	AndInvertedConst  = C.GL_AND_INVERTED
	OrReverseConst    = C.GL_OR_REVERSE
	OrInvertedConst   = C.GL_OR_INVERTED

	// MapBuffer
	ReadOnlyConst  = C.GL_READ_ONLY
	WriteOnlyConst = C.GL_WRITE_ONLY
	ReadWriteConst = C.GL_READ_WRITE

	// PointParameter
	PointSpriteCoordOriginConst = C.GL_POINT_SPRITE_COORD_ORIGIN

	// StencilOp, StencilOpSeparate
	KeepConst     = C.GL_KEEP
	ReplaceConst  = C.GL_REPLACE
	IncrConst     = C.GL_INCR
	IncrWrapConst = C.GL_INCR_WRAP
	DecrConst     = C.GL_DECR
	DecrWrapConst = C.GL_DECR_WRAP

	// TexBuffer
	R8Const       = C.GL_R8
	R16Const      = C.GL_R16
	R16fConst     = C.GL_R16F
	R32fConst     = C.GL_R32F
	R8iConst      = C.GL_R8I
	R16iConst     = C.GL_R16I
	R32iConst     = C.GL_R32I
	R8uiConst     = C.GL_R8UI
	R16uiConst    = C.GL_R16UI
	R32uiConst    = C.GL_R32UI
	Rg8Const      = C.GL_RG8
	Rg16Const     = C.GL_RG16
	Rg16fConst    = C.GL_RG16F
	Rg32fConst    = C.GL_RG32F
	Rg8iConst     = C.GL_RG8I
	Rg16iConst    = C.GL_RG16I
	Rg32iConst    = C.GL_RG32I
	Rg8uiConst    = C.GL_RG8UI
	Rg16uiConst   = C.GL_RG16UI
	Rg32uiConst   = C.GL_RG32UI
	Rgba16fConst  = C.GL_RGBA16F
	Rgba32fConst  = C.GL_RGBA32F
	Rgba8iConst   = C.GL_RGBA8I
	Rgba16iConst  = C.GL_RGBA16I
	Rgba32iConst  = C.GL_RGBA32I
	Rgba8uiConst  = C.GL_RGBA8UI
	Rgba16uiConst = C.GL_RGBA16UI
	Rgba32uiConst = C.GL_RGBA32UI

	// TransformFeedbackVaryings
	InterleavedAttribsConst = C.GL_INTERLEAVED_ATTRIBS
	SeparateAttribsConst    = C.GL_SEPARATE_ATTRIBS
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
