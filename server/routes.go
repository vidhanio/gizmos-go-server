package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *GizmoServer) GetGizmos(w http.ResponseWriter, r *http.Request) {
	gizmos, err := s.db.GetGizmos()
	if err != nil {
		log.Error().
			Err(err).
			Msg("Error getting gizmos.")

		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(NewGizmosResponse("Error getting gizmos.", nil))

		return
	}

	json.NewEncoder(w).Encode(NewGizmosResponse("Gizmos retrieved.", gizmos))
}

func (s *GizmoServer) GetGizmo(w http.ResponseWriter, r *http.Request) {
	resource, err := strconv.Atoi(chi.URLParam(r, "resource"))
	if err != nil {
		log.Error().
			Err(err).
			Msg("Error converting resource to int.")

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(NewGizmoResponse("Error converting resource to int.", nil))

		return
	}

	gizmo, err := s.db.GetGizmo(resource)
	if err == mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusNotFound)

		json.NewEncoder(w).Encode(NewGizmoResponse("Gizmo not found.", nil))

		return
	} else if err != nil {
		log.Error().
			Err(err).
			Msg("Error getting gizmo.")

		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(NewGizmoResponse("Error getting gizmo.", nil))

		return
	}

	w.WriteHeader(http.StatusFound)

	json.NewEncoder(w).Encode(NewGizmoResponse("Gizmo retrieved.", gizmo))
}

func (s *GizmoServer) PostGizmo(w http.ResponseWriter, r *http.Request) {
	gizmo := &Gizmo{}

	err := json.NewDecoder(r.Body).Decode(gizmo)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Error decoding gizmo.")

		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(NewGizmoResponse("Error decoding gizmo.", nil))

		return
	}

	err = s.db.InsertGizmo(gizmo)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Error inserting gizmo.")

		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(NewGizmoResponse("Error inserting gizmo.", nil))

		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(NewGizmoResponse("Gizmo created.", gizmo))
}

func (s *GizmoServer) PutGizmo(w http.ResponseWriter, r *http.Request) {
	resource, err := strconv.Atoi(chi.URLParam(r, "resource"))
	if err != nil {
		log.Error().
			Err(err).
			Msg("Error converting resource to int.")

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(NewGizmoResponse("Error converting resource to int.", nil))

		return
	}

	gizmo := &Gizmo{}

	err = json.NewDecoder(r.Body).Decode(gizmo)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Error decoding gizmo.")

		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(NewGizmoResponse("Error decoding gizmo.", nil))

		return
	}

	err = s.db.UpdateGizmo(resource, gizmo)
	if err == mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusNotFound)

		json.NewEncoder(w).Encode(NewGizmoResponse("Gizmo not found.", nil))

		return
	} else if err != nil {
		log.Error().
			Err(err).
			Msg("Error updating gizmo.")

		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(NewGizmoResponse("Error updating gizmo.", nil))

		return
	}

	json.NewEncoder(w).Encode(NewGizmoResponse("Gizmo updated.", nil))
}

func (s *GizmoServer) DeleteGizmo(w http.ResponseWriter, r *http.Request) {
	resource, err := strconv.Atoi(chi.URLParam(r, "resource"))
	if err != nil {
		log.Error().
			Err(err).
			Msg("Error converting resource to int.")

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(NewGizmoResponse("Error converting resource to int.", nil))

		return
	}

	err = s.db.DeleteGizmo(resource)
	if err == mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusNotFound)

		json.NewEncoder(w).Encode(NewGizmoResponse("Gizmo not found.", nil))

		return
	} else if err != nil {
		log.Error().
			Err(err).
			Msg("Error deleting gizmo.")

		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(NewGizmoResponse("Error deleting gizmo.", nil))

		return
	}

	json.NewEncoder(w).Encode(NewGizmoResponse("Gizmo deleted.", nil))
}
