package gl

// #include "common.h"
import "C"

type GLConstant int

const (
	Never            GLConstant = C.GL_NEVER
	Less                        = C.GL_LESS
	Equal                       = C.GL_EQUAL
	LEqual                      = C.GL_LEQUAL
	Greater                     = C.GL_GREATER
	NotEqual                    = C.GL_NOTEQUAL
	GEqual                      = C.GL_EQUAL
	Always                      = C.GL_ALWAYS
	ColorBufferBit              = C.GL_COLOR_BUFFER_BIT
	DepthBufferBit              = C.GL_DEPTH_BUFFER_BIT
	StencilBufferBit            = C.GL_STENCIL_BUFFER_BIT
	Blend                       = C.GL_BLEND
	//ClipDistance                      = GL_CLIP_DISTANCE
	ColorLogicOp           = C.GL_COLOR_LOGIC_OP
	CullFace               = C.GL_CULL_FACE
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
