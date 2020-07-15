package util

import (
	"fmt"
	"sort"
)

type Stage int

const (
	StageInit   = Stage(iota)
	StageNormal = Stage(iota)
	StageFinish = Stage(iota)
)

type Module struct {
	Stage    Stage
	Name     string
	Handler  func()
	Position string
}

func (stage Stage) RegisterModule(Name string, handler func()) Module {
	var position string
	lines := GetStackLines(8, 9)
	if len(lines) == 1 {
		position = fmt.Sprintf("%s\r\n", lines[0])
	}
	m := Module{Name: Name, Stage: stage, Handler: handler, Position: position}
	Modules = append(Modules, m)
	return m
}
func (m *Module) Load() {
	DebugPrintln("Herb-go util debug: Init module " + m.Name)
	if m.Position != "" {
		DebugPrint(m.Position)
	}
	m.Handler()
}

type modulelist []Module

var unloaders []func()

func OnUnloadModules(f func()) {
	unloaders = append(unloaders, f)
}

func UnloadModules() {
	for i := len(unloaders) - 1; i >= 0; i-- {
		unloaders[i]()
	}
}
func (m modulelist) Len() int {
	return len(m)
}
func (m modulelist) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m modulelist) Less(i, j int) bool {
	if m[i].Stage != m[j].Stage {
		return m[i].Stage < m[j].Stage
	}
	return m[i].Name < m[j].Name
}

var Modules = modulelist{}

func CleanModules() {
	Modules = modulelist{}
}
func RegisterModule(Name string, handler func()) Module {
	var position string
	lines := GetStackLines(8, 9)
	if len(lines) == 1 {
		position = fmt.Sprintf("%s\r\n", lines[0])
	}
	m := Module{Name: Name, Stage: StageNormal, Handler: handler, Position: position}
	Modules = append(Modules, m)
	return m
}

func InitModulesOrderByName(enabledModules ...string) {
	unloaders = []func(){}
	CleanWarnings()
	MustLoadRegisteredFolders()
	sort.Sort(Modules)
NextModule:
	for _, v := range Modules {
		if len(enabledModules) > 0 {
			for _, m := range enabledModules {
				if v.Name == m {
					v.Load()
					continue NextModule
				}
			}
		} else {
			v.Load()
		}
	}
	if Debug || ForceDebug {
		SetWarning("Util", "Debug mode enabled.")
	}
	if HasWarning() {
		output := "Warning:\r\n"
		for k, v := range warnings {
			output = output + "  " + k + ":\r\n"
			for _, wv := range v {
				output = output + "    " + wv + "\r\n"
			}
		}
		ErrorLogger(output)
	}
}
