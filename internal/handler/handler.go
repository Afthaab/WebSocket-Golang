package handler

import (
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode(),
)

func Home(w http.ResponseWriter, r *http.Request) {
	err := RenderPage(w, "home.jet", nil)
	if err != nil {
		log.Println("Error : ", err)
	}
}

//tmpl is the templete to render and data is the data needed to pass to the template
func RenderPage(w http.ResponseWriter, tmpl string, data jet.VarMap) error {
	view, err := views.GetTemplate(tmpl)
	if err != nil {
		log.Println("Error : ", err)
		return err
	}
	err = view.Execute(w, data, nil) // the third variable is the context which has been ignored in this course
	if err != nil {
		log.Println("Error : ", err)
		return err
	}
	return nil
}
