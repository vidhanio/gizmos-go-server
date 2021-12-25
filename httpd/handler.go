package httpd

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/vidhanio/gizmos-go-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetGizmos(c *mongo.Collection, w http.ResponseWriter, _ *http.Request) {
	// Initialize response

	w.WriteHeader(http.StatusOK)
	gizmosResponse := models.NewGizmosResponse("Gizmos retrieved.", []models.Gizmo{})

	// Initialize context

	ctx, cancel := getCtx()
	defer cancel()

	// Initialize cursor

	cursor, err := c.Find(ctx, bson.D{},
		&options.FindOptions{
			Sort: bson.D{
				{Key: "resource", Value: 1},
			},
		},
	)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Error finding gizmos.")

		w.WriteHeader(http.StatusInternalServerError)

		gizmosResponse.Message = "Error reading request body."

		response, err := json.Marshal(*gizmosResponse)
		if err != nil {
			log.Error().
				Err(err).
				Msg("Error marshalling error response.")

			return
		}

		_, err = w.Write(response)
		if err != nil {
			log.Error().
				Err(err).
				Msg("Error writing response.")

			return
		}

		return
	}

	// Write documents to gizmos

	err = cursor.All(ctx, &gizmosResponse.Gizmos)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Error decoding gizmos.")
	}

	// Marshal gizmos to JSON

	response, err := json.Marshal(*gizmosResponse)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Error marshalling gizmos.")
	}

	// Write response to response writer

	_, err = w.Write(response)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Error writing response.")
	}

	log.Debug().Str("method", "GET").Str("endpoint", "/gizmos").Msg("Gizmos retrieved.")
}

func GetGizmo(c *mongo.Collection, w http.ResponseWriter, r *http.Request) {
	// Initialize response

	w.WriteHeader(http.StatusOK)
	gizmoResponse := models.NewGizmoResponse("Gizmo retrieved.", models.Gizmo{})

	// Initialize context

	ctx, cancel := getCtx()
	defer cancel()

	// Get resource from parameters

	params := mux.Vars(r)
	resource, err := strconv.Atoi(params["resource"])
	if err != nil {
		log.Error().
			Err(err).
			Msg("Error converting resource to int.")

		w.WriteHeader(http.StatusBadRequest)

		gizmoResponse.Message = "Error converting resource to int."

		response, err := json.Marshal(*gizmoResponse)
		if err != nil {
			log.Error().
				Err(err).
				Msg("Error marshalling error response.")

			return
		}

		_, err = w.Write(response)
		if err != nil {
			log.Error().
				Err(err).
				Msg("Error writing response.")

			return
		}

		return
	}

	// Get gizmo from database

	err = c.FindOne(ctx, bson.D{{Key: "resource", Value: resource}}).Decode(&gizmoResponse.Gizmo)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Error finding gizmo.")

		w.WriteHeader(http.StatusInternalServerError)

		gizmoResponse.Message = "Error reading request body."

		response, err := json.Marshal(*gizmoResponse)
		if err != nil {
			log.Error().
				Err(err).
				Msg("Error marshalling error response.")

			return
		}

		_, err = w.Write(response)
		if err != nil {
			log.Error().
				Err(err).
				Msg("Error writing response.")

			return
		}

		return
	}

	// Marshal gizmos to JSON

	response, err := json.Marshal(*gizmoResponse)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Error marshalling gizmo.")
	}

	// Write response to response writer

	_, err = w.Write(response)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Error writing response.")
	}

	log.Debug().Str("method", "GET").Str("endpoint", "/gizmo").Msg("Gizmo retrieved.")
}

func CreateGizmo(c *mongo.Collection, w http.ResponseWriter, r *http.Request) {
	// Initialize response

	w.WriteHeader(http.StatusCreated)
	gizmoResponse := models.NewGizmoResponse("Gizmo created.", models.Gizmo{})

	// Initialize context

	ctx, cancel := getCtx()
	defer cancel()

	// Initialize decoder

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Error reading request body.")

		w.WriteHeader(http.StatusBadRequest)

		gizmoResponse.Message = "Error reading request body."

		response, err := json.Marshal(*gizmoResponse)
		if err != nil {
			log.Error().
				Err(err).
				Msg("Error marshalling error response.")

			return
		}

		_, err = w.Write(response)
		if err != nil {
			log.Error().
				Err(err).
				Msg("Error writing response.")

			return
		}

		return
	}

	// Unmarshal request body to gizmo

	err = json.Unmarshal(body, &gizmoResponse.Gizmo)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Error unmarshalling gizmo.")

		gizmoResponse.Message = "Error unmarshalling gizmo."

		w.WriteHeader(http.StatusInternalServerError)
	}

	// Insert gizmo

	result, err := c.InsertOne(ctx, &gizmoResponse.Gizmo)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Error inserting gizmo.")

		gizmoResponse.Message = "Error inserting gizmo."

		w.WriteHeader(http.StatusInternalServerError)
	}

	// Marshal gizmo to JSON

	gizmoResponse.Gizmo.ID = result.InsertedID.(primitive.ObjectID)

	response, err := json.Marshal(*gizmoResponse)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Error marshalling gizmo.")

		return
	}

	// Write response to response writer

	_, err = w.Write(response)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Error writing response.")
	}

	log.Debug().Str("method", "POST").Str("endpoint", "/add-gizmo").Msg("Gizmo created.")
}

func UpdateGizmo(c *mongo.Collection, w http.ResponseWriter, r *http.Request) {
	// Initialize response

	w.WriteHeader(http.StatusCreated)
	gizmoResponse := models.NewGizmoResponse("Gizmo updated.", models.Gizmo{})

	// Initialize context

	ctx, cancel := getCtx()
	defer cancel()

	// Get resource from parameters

	params := mux.Vars(r)
	resource, err := strconv.Atoi(params["resource"])
	if err != nil {
		log.Error().
			Err(err).
			Msg("Error converting resource to int.")

		w.WriteHeader(http.StatusBadRequest)

		gizmoResponse.Message = "Error converting resource to int."

		response, err := json.Marshal(*gizmoResponse)
		if err != nil {
			log.Error().
				Err(err).
				Msg("Error marshalling error response.")

			return
		}

		_, err = w.Write(response)
		if err != nil {
			log.Error().
				Err(err).
				Msg("Error writing response.")

			return
		}

		return
	}

	// Initialize decoder

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Error reading request body.")

		w.WriteHeader(http.StatusBadRequest)

		gizmoResponse.Message = "Error reading request body."

		response, err := json.Marshal(*gizmoResponse)
		if err != nil {
			log.Error().
				Err(err).
				Msg("Error marshalling error response.")

			return
		}

		_, err = w.Write(response)
		if err != nil {
			log.Error().
				Err(err).
				Msg("Error writing response.")

			return
		}

		return
	}

	// Unmarshal request body to gizmo

	err = json.Unmarshal(body, &gizmoResponse.Gizmo)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Error unmarshalling gizmo.")

		gizmoResponse.Message = "Error unmarshalling gizmo."

		w.WriteHeader(http.StatusInternalServerError)
	}

	// Insert gizmo

	err = c.FindOneAndUpdate(ctx, bson.D{{Key: "resource", Value: resource}}, &gizmoResponse.Gizmo).Decode(&gizmoResponse.Gizmo)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Error updating gizmo.")

		gizmoResponse.Message = "Error updating gizmo."

		w.WriteHeader(http.StatusInternalServerError)
	}

	// Marshal gizmo to JSON

	response, err := json.Marshal(*gizmoResponse)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Error marshalling gizmo.")

		return
	}

	// Write response to response writer

	_, err = w.Write(response)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Error writing response.")
	}

	log.Debug().Str("method", "PUT").Str("endpoint", "/update-gizmo").Msg("Gizmo updated.")
}

func DeleteGizmo(c *mongo.Collection, w http.ResponseWriter, r *http.Request) {
	// Initialize response

	w.WriteHeader(http.StatusCreated)
	gizmoResponse := models.NewGizmoResponse("Gizmo deleted.", models.Gizmo{})

	// Initialize context

	ctx, cancel := getCtx()
	defer cancel()

	// Get resource from parameters

	params := mux.Vars(r)
	resource, err := strconv.Atoi(params["resource"])
	if err != nil {
		log.Error().
			Err(err).
			Msg("Error converting resource to int.")

		w.WriteHeader(http.StatusBadRequest)

		gizmoResponse.Message = "Error converting resource to int."

		response, err := json.Marshal(*gizmoResponse)
		if err != nil {
			log.Error().
				Err(err).
				Msg("Error marshalling error response.")

			return
		}

		_, err = w.Write(response)
		if err != nil {
			log.Error().
				Err(err).
				Msg("Error writing response.")

			return
		}

		return
	}

	// Delete gizmo

	err = c.FindOneAndDelete(ctx, bson.D{{Key: "resource", Value: resource}}).Decode(&gizmoResponse.Gizmo)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Error deleting gizmo.")

		gizmoResponse.Message = "Error deleting gizmo."

		w.WriteHeader(http.StatusInternalServerError)
	}

	// Marshal gizmo to JSON

	response, err := json.Marshal(*gizmoResponse)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Error marshalling gizmo.")

		return
	}

	// Write response to response writer

	_, err = w.Write(response)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Error writing response.")
	}

	log.Debug().Str("method", "DELETE").Str("endpoint", "/delete-gizmo").Msg("Gizmo deleted.")
}

func getCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}
