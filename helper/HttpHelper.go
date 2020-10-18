package helper

import (
	"encoding/json"
	"fmt"
	"math/bits"
	"net/http"
	"strconv"
	"strings"
)

type CustomUint struct {
	Data   uint64
	IsNull bool
}

func GetError(err error) {
	if err != nil {
		panic(err)
	}
}

func PathId(r *http.Request, controllerPath string) CustomUint {
	if pathId := strings.ReplaceAll(r.URL.Path, controllerPath, ""); pathId != "" {
		parsedPathId, err := strconv.ParseUint(pathId, 10, bits.UintSize)
		if err != nil {
			panic(fmt.Errorf("ivalid pathid: %v", pathId))
		}
		return CustomUint{Data: parsedPathId}
	}
	return CustomUint{IsNull: true}
}

func GetBodyJson(r *http.Request, model interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&model)
	GetError(err)
}

func ResponseJson(w http.ResponseWriter, model interface{}, statusCode int) {
	res, err := json.Marshal(model)
	GetError(err)
	w.WriteHeader(statusCode)
	_, err = w.Write(res)
	GetError(err)
}
