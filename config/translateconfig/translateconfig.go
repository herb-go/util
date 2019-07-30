package translateconfig

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"

	"github.com/herb-go/herb/translate"
)

//Config translate config
type Config struct {
	Language  string
	Avaliable []string
}

// Apply apply translate config
func (c *Config) Apply(path string) error {
	translate.Lang = c.Language
	for _, v := range c.Avaliable {
		langpath := filepath.Join(path, v)
		s, err := os.Stat(langpath)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return err
		}
		if !s.IsDir() {
			continue
		}
		filepath.Walk(langpath, func(path string, f os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !f.IsDir() {
				if strings.HasSuffix(path, ".toml") {
					m := translate.NewMessages()
					base := filepath.Base(path)
					modulename := base[:len(base)-5]
					data, err := ioutil.ReadFile(path)
					if err != nil {
						return err
					}
					err = toml.Unmarshal(data, m)
					if err != nil {
						return err
					}
					translate.DefaultTranslations.SetMessages(v, modulename, m)
				}
			}
			return nil
		})
	}
	return nil
}
