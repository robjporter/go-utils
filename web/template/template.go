package template

import (
	"github.com/flosch/pongo2"
)

func init() {

}

func TemplateToString(templatestring string, values map[string]interface{}) string {
	tpl, err := pongo2.FromString(templatestring)
	if err != nil {
		panic(err)
	}

	pon := pongo2.Context{}
	for k, v := range values {
		pon[k] = v
	}

	out, err := tpl.Execute(pon)

	if err != nil {
		panic(err)
	}
	return out
}
