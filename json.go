package gres

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/go-chi/render"
)

// ResponseJSONRender will simply JSON the response and add the specific headers
func ResponseJSONRender(w http.ResponseWriter, r *http.Request, e interface{}) {
	// Render JSON
	eJSON, errJSON := JSONMarshal(e)
	if errJSON != nil {
		errRend := render.Render(w, r, ErrRender(errJSON))
		if errRend != nil {
			logrus.Errorf("ResponseJSONRender %s", errRend)
		}
		return
	}

	_, errW := fmt.Fprintf(w, string(eJSON))
	if errW != nil {
		errRend := render.Render(w, r, ErrRender(errW))
		if errRend != nil {
			logrus.Errorf("ResponseJSONRender %s", errRend)
		}
		return
	}
}

// JSONMarshal marshals json and ignoreing HTML chars
func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}
