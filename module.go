package util

import (
	"fmt"
	"log"
	"sort"
)

type Module struct {
	Name     string
	Handler  func()
	Position string
}

func (m *Module) Load() {
	if Debug || ForceDebug {
		fmt.Println("Herb-go util debug: Init module " + m.Name)
		if m.Position != "" {
			fmt.Print(m.Position)
		}
	}
	m.Handler()
}

var Debug = false

func DebugPrintln(args ...interface{}) {
	if Debug {
		fmt.Println(args...)
	}
}

type modulelist []Module

var unloaders []func()

func OnUnloadModules(f func()) {
	unloaders = append(unloaders, f)
}

func UnloadModules() {
	for _, v := range unloaders {
		v()
	}
}
func (m modulelist) Len() int {
	return len(m)
}
func (m modulelist) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m modulelist) Less(i, j int) bool {
	return m[i].Name < m[j].Name
}

var Modules = modulelist{}

func RegisterModule(Name string, handler func()) Module {
	var position string
	lines := GetStackLines(8, 9)
	if len(lines) == 1 {
		position = fmt.Sprintf("%s\r\n", lines[0])
	}
	m := Module{Name: Name, Handler: handler, Position: position}
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
		}
		v.Load()
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
		log.Print(output)
	}
}
