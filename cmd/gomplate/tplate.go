//+build foo

package foo

// import (
// 	"io"
// 	"text/template"
// )

// // tplate - models a tplate file...
// type tplate struct {
// 	name     string
// 	target   io.Writer
// 	contents string
// }

// func (t *tplate) toGoTemplate(g *Gomplate) (*template.Template, error) {
// 	tmpl := template.New(t.name)
// 	tmpl.Option("missingkey=error")
// 	tmpl.Funcs(g.funcMap)
// 	tmpl.Delims(g.leftDelim, g.rightDelim)
// 	return tmpl.Parse(t.contents)
// }

// // loadContents - reads the template in _once_ if it hasn't yet been read. Uses the name!
// func (t *tplate) loadContents() (err error) {
// 	if t.contents == "" {
// 		t.contents, err = readInput(t.name)
// 	}
// 	return err
// }

// func (t *tplate) addTarget(outFile string) error {
// 	if t.target == nil {
// 		target, err := openOutFile(outFile)
// 		if err != nil {
// 			return err
// 		}
// 		addCleanupHook(func() {
// 			// nolint: errcheck
// 			target.Close()
// 		})
// 		t.target = target
// 	}
// 	return nil
// }
