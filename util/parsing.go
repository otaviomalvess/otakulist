package util

import (
	"encoding/json"
	"io"
	"log"
	"otavio-alves/OtakuList/configs"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
)

// StringConv .. Converts string into given bit size
func StringConv(str string, bitSize int) (val int64, err error) {

	// Converts string into given bit size
	if val, err = strconv.ParseInt(str, 10, bitSize); err != nil {
		log.Printf(configs.WARN_PARSING_STR, err)
	}

	return
}

// BodyToModel .. Parses request body into entity struct
func BodyToModel(body io.Reader, entity interface{}) (err error) {

	// Parses requests body into entity struct
	if err = json.NewDecoder(body).Decode(&entity); err != nil {
		log.Printf(configs.WARN_PARSING_JSON, err)
	}

	return
}

// MapToBSON ..
func MapToBSON(m map[string][]string) (b bson.D) {

	for k, v := range m {
		for _, v2 := range v {
			if v2 != configs.EMPTY_STR {
				b = append(b, bson.E{Key: k, Value: v2})
			}
		}
	}

	return
}

// Marshal ..
type Marshal func(interface{}) ([]byte, error)

// StructToSON .. Parses given struct into a JSON or a BSON bytes slice
func StructToSON(entity interface{}, m Marshal) (bytes []byte, err error) {

	// Parses entity struct into a JSON or a BSON bytes slice
	if bytes, err = m(entity); err != nil {
		log.Printf(configs.WARN_PARSING_SON, err)
	}

	return
}
