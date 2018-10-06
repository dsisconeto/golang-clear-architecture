package helper

import (
    "encoding/json"
    "github.com/dsisconeto/arquitetura/core/common"
    "net/http"
)

func RequestJson(r *http.Request, data interface{}) error {
    jsonDecode := json.NewDecoder(r.Body)
    err := jsonDecode.Decode(&data)
    return err
}

func ResponseJson(w http.ResponseWriter, data interface{}) {
    responseJson, err := json.Marshal(data)
    if err != nil {
        ResponseInternalServeError(w)
        return
    }
    w.Write(responseJson)
}

func ResponseErrorJson(w http.ResponseWriter, error error) {
    response := struct {
        Message string `json:"message"`
    }{}
    switch error.(type) {
    case common.IDomainError:
        w.WriteHeader(http.StatusUnprocessableEntity)
        response.Message = error.Error()
        ResponseJson(w, response)

    case common.EntityNotFoundError:
        w.WriteHeader(http.StatusNotFound)
        response.Message = error.Error()
        ResponseJson(w, response)
    default:
        ResponseInternalServeError(w)
    }
}

func ResponseBadRequest(w http.ResponseWriter) {
    w.WriteHeader(http.StatusBadRequest)
}
func ResponseInternalServeError(w http.ResponseWriter) {
    w.WriteHeader(http.StatusInternalServerError)
}
