package main

import (
	"fmt"
	"log"
	"math"
	"runtime"
	"strings"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	vertexShaderSource = `
		#version 410
		in vec3 vp;
		void main() {
			gl_Position = vec4(vp, 1.0);
		}
	` + "\x00"

	fragmentShaderHourSource = `
	#version 410 core
	out vec4 color;
	void main()
	{
		color = vec4(1.0f, 0.0f, 0.0f, 1.0f);
	}
	` + "\x00"

	fragmentShaderMinuteSource = `
	#version 410 core
	out vec4 color;
	void main()
	{
		color = vec4(0.0f, 0.0f, 1.0f, 1.0f);
	}
	` + "\x00"

	fragmentShaderSecondSource = `
	#version 410 core
	out vec4 color;
	void main()
	{
		color = vec4(1.0f, 1.0f, 0.0f, 1.0f);
	}
	` + "\x00"

	fragmentShaderCardinalsSource = `
	#version 410 core
	out vec4 color;
	void main()
	{
		color = vec4(1.0f, 1.0f, 1.0f, 1.0f);
	}	` + "\x00"
)

var shaderProgramMinute uint32
var shaderProgramSecond uint32
var shaderProgramHour uint32
var shaderProgramCardinals uint32

func main() {

	runtime.LockOSThread()

	window := initGlfw()
	defer glfw.Terminate()
	initOpenGL()

	for !window.ShouldClose() {

		hour := float64(time.Now().Hour())
		minute := float64(time.Now().Minute())
		second := float64(time.Now().Second())
		millis := float64(time.Now().Nanosecond() / 1000000)

		//gl.ClearColor(0.0, 0.0, 0.0, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		drawBezel()
		drawCardinals()
		drawHourHand(hour, minute, second)
		drawMinuteHand(minute, second, millis)
		drawSecondHand(second, millis)
		drawBoss()
		glfw.PollEvents()
		window.SwapBuffers()
		time.Sleep(time.Millisecond * 50)
	}
}

func drawBezel() {
	for theta := 0.0; theta < 2*math.Pi; theta += math.Pi / 30 {
		gl.UseProgram(shaderProgramCardinals)
		draw(initVertices(theta, 0.78, 0.01, 0.8))
	}
}

func drawCardinals() {
	for theta := 0.0; theta < 2*math.Pi; theta += math.Pi / 6 {
		gl.UseProgram(shaderProgramCardinals)
		draw(initVertices(theta, 0.8, 0.01, 0.7))
	}
}

func drawBoss() {
	for theta := 0.0; theta < 2*math.Pi; theta += math.Pi / 12 {
		gl.UseProgram(shaderProgramSecond)
		draw(initVertices(theta, 0.05, 0.01, 0.0))
	}
}
func drawSecondHand(second, millis float64) {
	theta := 2.0 * math.Pi * (second + millis/1000) / 60.0
	gl.UseProgram(shaderProgramSecond)
	draw(initVertices(theta, 0.8, 0.005, -0.1))
}

func drawHourHand(hour, minute, second float64) {
	theta := 2.0 * math.Pi * ((hour + (minute+second/60)/60.0) / 12.0)
	gl.UseProgram(shaderProgramHour)
	draw(initVertices(theta, 0.6, 0.02, -0.1))
}

func drawMinuteHand(minute, second, millis float64) {

	theta := 2.0 * math.Pi * (minute + (second+millis/1000)/60.0) / 60.0
	gl.UseProgram(shaderProgramMinute)
	draw(initVertices(theta, 0.7, 0.01, -0.1))
}

func initVertices(theta, r, w, y float64) []float32 {
	cosTheta := math.Sin(theta)
	sinTheta := math.Cos(theta)

	vertices := []float32{
		float32(r*cosTheta - w*sinTheta),
		float32(r*sinTheta + w*cosTheta),
		float32(r*cosTheta + w*sinTheta),
		float32(r*sinTheta - w*cosTheta),
		float32(y*cosTheta + w*sinTheta),
		float32(y*sinTheta - w*cosTheta),
		float32(y*cosTheta - w*sinTheta),
		float32(y*sinTheta + w*cosTheta),
	}

	return vertices
}

// Draw :
func draw(vertices []float32) {
	var vao, vbo, ebo uint32
	elements := []uint32{
		0, 1, 2,
		2, 3, 0}

	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(vertices), gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.GenBuffers(1, &ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, 4*len(elements), gl.Ptr(elements), gl.STATIC_DRAW)

	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 0, nil)
	gl.BindVertexArray(vao)

	gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)
	gl.DeleteBuffers(1, &ebo)
	gl.DeleteBuffers(1, &vbo)
	gl.DeleteBuffers(1, &vao)
}

// initGlfw initializes glfw and returns a Window to use.
func initGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.Samples, 10)

	window, err := glfw.CreateWindow(300, 300, "Clock", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}

// initOpenGL initializes OpenGL and returns an intiialized program.
func initOpenGL() {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	gl.Enable(gl.MULTISAMPLE)
	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	fragmentShaderMinute, err := compileShader(fragmentShaderMinuteSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	fragmentShaderHour, err := compileShader(fragmentShaderHourSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	fragmentShaderSecond, err := compileShader(fragmentShaderSecondSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	fragmentShaderCardinlas, err := compileShader(fragmentShaderCardinalsSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}
	shaderProgramMinute = linkShaders([]uint32{vertexShader, fragmentShaderMinute})
	shaderProgramSecond = linkShaders([]uint32{vertexShader, fragmentShaderSecond})
	shaderProgramHour = linkShaders([]uint32{vertexShader, fragmentShaderHour})
	shaderProgramCardinals = linkShaders([]uint32{vertexShader, fragmentShaderCardinlas})
}
func linkShaders(shaders []uint32) uint32 {
	program := gl.CreateProgram()
	for _, shader := range shaders {
		gl.AttachShader(program, shader)
	}
	gl.LinkProgram(program)

	return program
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
