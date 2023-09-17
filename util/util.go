package util

import (
	"net/http"
	"otavio-alves/OtakuList/model"
)

// QueryToContentStruct .. Transform query map into content struct
func QueryToContentStruct(query map[string][]string, content model.FilterableContent) model.FilterableContent {

	//if val, ok := query["name"]; ok {
	//		content.Name = val[0]
	//	}
	if val, ok := query["type"]; ok {
		content.Type = val[0]
	}
	//content.Status = query["status"][0]
	//content.Producer = query["producer"][0]
	//content.Rating = query["rating"][0]
	//content.Genres

	return content
}

// WriteResponse .. Writes request response
func WriteResponse(w http.ResponseWriter, bytes []byte) {

	// Writes header in wire format
	_, err := w.Write(bytes)

	// Checks if any error occurs when writing header
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
