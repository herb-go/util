package app

import (
	"flag"
)

type Module interface {
	ID() string
	Cmd() string
	Help(a *Application) string
	Desc(a *Application) string
	Group(a *Application) string
	FlagSet() *flag.FlagSet
	Exec(a *Application, args []string) error
}

type BasicModule struct {
	flagset *flag.FlagSet
}

func (m *BasicModule) ID() string {
	return ""
}

func (m *BasicModule) Cmd() string {
	return ""
}

func (m *BasicModule) Help(a *Application) string {
	return ""
}

func (m *BasicModule) Desc(a *Application) string {
	return ""
}

func (m *BasicModule) Group(a *Application) string {
	return ""
}
func (m *BasicModule) FlagSet() *flag.FlagSet {
	if m.flagset == nil {
		m.flagset = flag.NewFlagSet("", flag.ContinueOnError)
	}
	return m.flagset
}
func (m *BasicModule) Exec(a *Application, args []string) error {
	return nil
}

type Modules []Module

func (modules *Modules) Add(m Module) {
	cmd := m.Cmd()
	for k := range *modules {
		if cmd == (*modules)[k].Cmd() {
			(*modules)[k] = m
			return
		}
	}
	*modules = append(*modules, m)

}

func (modules *Modules) Get(cmd string) Module {
	for k := range *modules {
		if (*modules)[k].Cmd() == cmd {
			return (*modules)[k]
		}
	}
	return nil

}
func NewModules() *Modules {
	return &Modules{}
}

var RegisteredModules = NewModules()

func Register(m Module) {
	RegisteredModules.Add(m)
}
