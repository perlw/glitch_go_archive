package shaders

import (
	"errors"
	"github.com/doxxan/glitch/gl"
	"github.com/doxxan/glitch/math/matrix"
)

var currentShaderProgram uint32

type ShaderProgram struct {
	id       uint32
	uniforms map[string]int
}

func NewShaderProgram() (*ShaderProgram, error) {
	var id uint32
	if id = gl.CreateProgram(); id <= 0 {
		return nil, errors.New("Could not create shaderprogram")
	}

	sp := &ShaderProgram{
		id:       id,
		uniforms: make(map[string]int),
	}

	return sp, nil
}

func NewShaderProgramFromFilePath(filePath string) (*ShaderProgram, error) {
	vertexShader, err := NewShaderFromFilePath(filePath + ".vert")
	if err != nil {
		return nil, err
	}
	defer vertexShader.Delete()

	fragmentShader, err := NewShaderFromFilePath(filePath + ".frag")
	if err != nil {
		return nil, err
	}
	defer fragmentShader.Delete()

	shaderProgram, err := NewShaderProgram()
	if err != nil {
		return nil, err
	}

	shaderProgram.AttachShader(vertexShader)
	defer shaderProgram.DetachShader(vertexShader)
	shaderProgram.AttachShader(fragmentShader)
	defer shaderProgram.DetachShader(fragmentShader)

	if err := shaderProgram.Link(); err != nil {
		return nil, err
	}

	return shaderProgram, nil
}

func (sp ShaderProgram) AttachShader(shader *Shader) {
	gl.AttachShader(sp.id, shader.id)
}

func (sp ShaderProgram) DetachShader(shader *Shader) {
	gl.DetachShader(sp.id, shader.id)
}

func (sp ShaderProgram) Link() error {
	gl.LinkProgram(sp.id)
	if ok := gl.GetProgram(sp.id, gl.LinkStatusConst); !gl.IsTrue(ok) {
		return errors.New("Shader program linking FAIL\n" + gl.GetProgramInfoLog(sp.id))
	}
	return nil
}

func (sp ShaderProgram) Use() {
	if sp.id == currentShaderProgram {
		return
	}
	currentShaderProgram = sp.id
	gl.UseProgram(sp.id)
}

func (sp ShaderProgram) GetInternalId() uint32 {
	return sp.id
}

func (sp *ShaderProgram) GetUniformLocation(name string) (int, error) {
	var id int
	id, ok := sp.uniforms[name]
	if !ok {
		if id = gl.GetUniformLocation(sp.id, name); id < 0 {
			return -1, errors.New("Could not find uniform location for " + name)
		}
		sp.uniforms[name] = id
	}
	return id, nil
}

func (sp *ShaderProgram) SetUniformMatrix(location int, matrix matrix.Matrix) {
	gl.UniformMatrix4f(location, 1, false, matrix)
}

func (sp *ShaderProgram) SetUniformMatrixByName(name string, matrix matrix.Matrix) error {
	id, err := sp.GetUniformLocation(name)
	if err != nil {
		return err
	}
	gl.UniformMatrix4f(id, 1, false, matrix)
	return nil
}
