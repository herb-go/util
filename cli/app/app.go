package app

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"text/template"

	"github.com/herb-go/util"
)

var HelpModuleCmd = "help"

type ApplicationConfig struct {
	Name          string
	Cmd           string
	Version       string
	IntroTemplate string
}
type Application struct {
	Config        *ApplicationConfig
	Args          []string
	Env           AppEnv
	Cwd           string
	Modules       *Modules
	Stdout        io.Writer
	Stdin         io.Reader
	FlagDefaults  string
	HelpModuleCmd string
}

func (a *Application) ShowIntro() error {
	t, err := template.New("intro").Parse(a.Config.IntroTemplate)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer([]byte{})
	err = t.Execute(buf, a)
	if err != nil {
		return err
	}
	_, err = a.Println(buf.String())
	return err
}
func (a *Application) Print(args ...interface{}) (n int, err error) {
	return fmt.Fprint(a.Stdout, args...)
}

func (a *Application) Println(args ...interface{}) (n int, err error) {
	return fmt.Fprintln(a.Stdout, args...)
}

func (a *Application) Printf(format string, args ...interface{}) (n int, err error) {
	return fmt.Fprintf(a.Stdout, format, args...)
}

func (a *Application) Getenv(key string) string {
	return a.Env.Getenv(key)
}
func (a *Application) appendFlagDefaults(flgs *flag.FlagSet) {
	output := bytes.NewBuffer(nil)
	flgs.SetOutput(output)
	flgs.PrintDefaults()
	a.FlagDefaults += output.String()
}

func (a *Application) PrintModuleHelp(m Module) {
	a.FlagDefaults = ""
	a.ShowIntro()
	a.Print(m.Help(a))
	a.appendFlagDefaults(m.FlagSet())
	if a.FlagDefaults != "" {
		a.Println("Options:")
		a.Println(a.FlagDefaults)
	}
}
func (a *Application) Setenv(key string, value string) error {
	return a.Env.Setenv(key, value)
}
func (a *Application) Run() {
	var cmd string
	var args []string
	var err error
	if len(a.Args) < 2 {
		err = a.ShowIntro()
		if err == nil {
			a.Printf("Type \"%s %s\" to get help.\n", a.Config.Cmd, HelpModuleCmd)
		}
	} else {
		cmd = a.Args[1]
		args = a.Args[2:]
		m := a.Modules.Get(cmd)
		if m == nil {
			m = a.Modules.Get(HelpModuleCmd)
		}
		err = m.Exec(a, args)
	}
	if err != nil {
		util.Println("Error: " + err.Error())
	}
}

func NewApplication(config *ApplicationConfig) *Application {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return &Application{
		Config:        config,
		Cwd:           cwd,
		Modules:       NewModules(),
		Stdout:        os.Stdout,
		Stdin:         os.Stdin,
		Env:           OsEnv,
		HelpModuleCmd: HelpModuleCmd,
	}
}

func NewApplicationConfig() *ApplicationConfig {
	return &ApplicationConfig{}
}
func init() {
}
