package shaders

import (
	"errors"
	"github.com/doxxan/glitch/gl"
	"io/ioutil"
	"path/filepath"
)

type Shader struct {
	id uint32
}

func NewShader(shType ShaderType) (*Shader, error) {
	var glType gl.GLenum

	switch shType {
	case Vertex:
		glType = gl.VertexShaderConst
	case Fragment:
		glType = gl.FragmentShaderConst
	default:
		return nil, errors.New("Unknown shader extension")
	}

	var id uint32
	if id = gl.CreateShader(glType); id <= 0 {
		return nil, errors.New("Could not create shader")
	}

	s := &Shader{
		id: id,
	}

	return s, nil
}

func NewShaderFromFilePath(filePath string) (*Shader, error) {
	var shType ShaderType

	ext := filepath.Ext(filePath)
	switch ext {
	case ".vert":
		shType = Vertex
	case ".frag":
		shType = Fragment
	default:
		return nil, errors.New("Unknown shader extension " + ext)
	}

	shaderSource, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return NewShaderFromString(string(shaderSource), shType)
}

func NewShaderFromString(source string, shType ShaderType) (*Shader, error) {
	s, err := NewShader(shType)
	if err != nil {
		return nil, err
	}

	s.SetSource(source)
	if err := s.Compile(); err != nil {
		return nil, err
	}

	return s, nil
}

func (s Shader) SetSource(source string) {
	gl.ShaderSource(s.id, string(source))
}

func (s Shader) Compile() error {
	gl.CompileShader(s.id)
	if ok := gl.GetShader(s.id, gl.CompileStatusConst); !gl.IsTrue(ok) {
		return errors.New("Shader compilation FAIL\n" + gl.GetShaderInfoLog(s.id))
	}
	return nil
}

func (s *Shader) Delete() {
	gl.DeleteShader(s.id)
	s.id = 0
}
