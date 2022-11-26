package config

import (
	"os"
	"path/filepath"
)

var (
	prog *program
	app  *application
)

type program struct {
	Path    string
	AbsPath string
	Dir     string
	File    string
}
type application struct {
	// 应用名称
	Name string
	// 应用版本
	Version string
	// 应用调试
	Debug bool
	// 基础路径
	BasePath string
}

func init() {
	initApp()
}

func initProgram() {
	abspath, err := filepath.Abs(os.Args[0])
	if err != nil {
		panic(err)
	}
	dir, file := filepath.Split(abspath)
	prog = &program{
		Path:    os.Args[0],
		AbsPath: abspath,
		Dir:     dir[:len(dir)-1],
		File:    file,
	}
}

func initApp() {
	initProgram()
	app = &application{}
	app.BasePath = prog.Dir
}

func App() *application {
	return app
}
