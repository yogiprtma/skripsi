package render

import (
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
)

func LoadTemplates(templatesDir string) multitemplate.Renderer {
	render := multitemplate.NewRenderer()
	render.AddFromFiles("session_new", templatesDir+"/session/session_new.html")
	render.AddFromFiles("session_employee_new", templatesDir+"/session/session_pegawai_new.html")
	render.AddFromFiles("check_document", templatesDir+"/check-document/check_document.html")
	render.AddFromFiles("document_not_found", templatesDir+"/check-document/document-not-found.html")
	render.AddFromFiles("document_found", templatesDir+"/check-document/document-found.html")

	layouts, err := filepath.Glob(templatesDir + "/layouts/*")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/includes/**/*")
	if err != nil {
		panic(err.Error())
	}

	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		render.AddFromFiles(filepath.Base(include), files...)
	}
	return render
}
