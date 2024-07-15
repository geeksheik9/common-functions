package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// RespondWithError Utility function to convert an error message into a JSON response.
func RespondWithError(w http.ResponseWriter, code int, msg string) {
	// Clean up all quote marks for readability, marshal adds additional "\" escape char in the response JSON.
	RespondWithJSON(w, code, map[string]string{"error": strings.Replace(msg, `"`, ``, -1)})
}

// RespondNoContent Utility function to send a response without any content.
func RespondNoContent(w http.ResponseWriter, code int) {
	if w != nil {
		w.Header().Set("Content-Length", "0")
		w.WriteHeader(code)
	}
}

// RespondWithJSON Utility function to convert the payload into a JSON response.
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	if w != nil {
		response, err := json.Marshal(payload)
		if err != nil {
			log.Errorf("Error in RespondWithJSON marshal: %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(response)
	}
}

// GetJSONRequestBody function to return the request body as string
func GetJSONRequestBody(r *http.Request) (requestBodyJSON string) {
	defer r.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	requestBodyJSON = string(bodyBytes)

	return requestBodyJSON
}

// StringToObjectID takes a string and checks if it is a valid objectId hex if so it returns an objectID
func StringToObjectID(ID string) (primitive.ObjectID, error) {
	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	if objID.IsZero() {
		return objID, errors.New("StringToObjectID() in api.go failed, object ID returned as zero and is not valid")
	}

	return objID, nil
}

// CheckError checks the err message and returns a code based on the message.
func CheckError(err error) int {
	switch {
	case err == nil:
		return http.StatusOK
	case strings.Contains(err.Error(), "no documents in results"), strings.Contains(err.Error(), "out of bounds"), strings.Contains(err.Error(), "not found"):
		return http.StatusNotFound
	case strings.Contains(err.Error(), "duplicate key error"):
		return http.StatusConflict
	case strings.Contains(err.Error(), "Invalid request payload, unable to marshal into json, err: "):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
