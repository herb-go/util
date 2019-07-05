package app

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

var Config = NewApplicationConfig()

func TestApp(t *testing.T) {
	App := NewApplication(Config)
	output := bytes.NewBuffer([]byte{})
	App.Stdout = output
	App.Print("123")
	bs, err := ioutil.ReadAll(output)
	if err != nil {
		t.Fatal(err)
	}
	if string(bs) != fmt.Sprint("123") {
		t.Fatal(string(bs))
	}
	App.Println("123")
	bs, err = ioutil.ReadAll(output)
	if err != nil {
		t.Fatal(err)
	}
	if string(bs) != fmt.Sprintln("123") {
		t.Fatal(string(bs))
	}
	App.Printf("%s", "123")
	bs, err = ioutil.ReadAll(output)
	if err != nil {
		t.Fatal(err)
	}
	if string(bs) != fmt.Sprintf("%s", "123") {
		t.Fatal(string(bs))
	}
}

func TestIntro(t *testing.T) {
	App := NewApplication(Config)
	output := bytes.NewBuffer([]byte{})
	App.Stdout = output
	err := App.ShowIntro()
	if err != nil {
		t.Fatal(err)
	}
	intro := output.String()
	if !strings.Contains(intro, Config.Name) || !strings.Contains(intro, Config.Version) || !strings.Contains(intro, Config.Cmd) || !strings.Contains(intro, App.HelpModuleCmd) {
		t.Fatal(intro)
	}
}

func TestAppRun(t *testing.T) {
	app := NewApplication(Config)
	app.Args = []string{}
	output := bytes.NewBuffer([]byte{})
	app.Stdout = output
	app.Run()
	intro := output.String()
	if !strings.Contains(intro, Config.Name) || !strings.Contains(intro, Config.Version) || !strings.Contains(intro, Config.Cmd) || !strings.Contains(intro, app.HelpModuleCmd) {
		t.Fatal(intro)
	}
	var testHelpModule = newTestModule(HelpModuleCmd)
	testHelpModule.output = "helpoutput"
	var testExecModule = newTestModule("exec")
	testExecModule.output = "execoutput"
	modules := NewModules()
	modules.Add(testHelpModule)
	modules.Add(testExecModule)
	app = NewApplication(Config)
	app.Args = []string{"app", "cmdNotExist"}
	app.Modules = modules
	output = bytes.NewBuffer([]byte{})
	app.Stdout = output
	app.Run()
	intro = output.String()
	if intro != fmt.Sprintln("helpoutput") {
		t.Fatal(intro)
	}
	app = NewApplication(Config)
	app.Args = []string{"app", "exec"}
	app.Modules = modules
	output = bytes.NewBuffer([]byte{})
	app.Stdout = output
	app.Run()
	intro = output.String()
	if intro != fmt.Sprintln("execoutput") {
		t.Fatal(intro)
	}
}

func TestFlagDefaults(t *testing.T) {
	app := NewApplication(nil)
	flgs := flag.NewFlagSet("", flag.ContinueOnError)
	flgs.String("test", "testdefault", "testinfo")
	app.appendFlagDefaults(flgs)
	if !strings.Contains(app.FlagDefaults, "testdefault") || !strings.Contains(app.FlagDefaults, "testinfo") {
		t.Fatal(app.FlagDefaults)
	}
}

func TestHelp(t *testing.T) {
	app := NewApplication(NewApplicationConfig())
	output := bytes.NewBuffer(nil)
	app.Stdout = output
	module := newTestModule("test")
	module.FlagSet().String("testfield", "", "testuseage")
	app.PrintModuleHelp(module)
	out := output.String()
	if !strings.Contains(out , "testuseage") {
		t.Fatal(out)
	}
}
func init() {
	Config.Name = "Herb-go cli tool"
	Config.Cmd = "herb-go"
	Config.Version = "0.1"
	Config.IntroTemplate = "{{.Config.Name}} Version {{.Config.Version}}\nCli tool to create herb-go app.\nType \"{{.Config.Cmd}} {{.HelpModuleCmd}}\" to get help."
}
