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

		http.Error(w, "Error getting gizmos.", http.StatusInternalServerError)

		return
	}

	json.NewEncoder(w).Encode(NewGizmosResponse("Gizmos retrieved.", NewGizmosFromDBGizmos(gizmos)))
}

func (s *GizmoServer) GetGizmo(w http.ResponseWriter, r *http.Request) {
	resource, err := strconv.Atoi(chi.URLParam(r, "resource"))
	if err != nil {
		log.Error().
			Err(err).
			Msg("Error converting resource to int.")

		http.Error(w, "Error converting resource to int.", http.StatusBadRequest)

		return
	}

	gizmo, err := s.db.GetGizmo(resource)
	if err == mongo.ErrNoDocuments {
		http.Error(w, "Gizmo not found.", http.StatusNotFound)

		return
	} else if err != nil {
		log.Error().
			Err(err).
			Msg("Error getting gizmo.")

		http.Error(w, "Error getting gizmo.", http.StatusInternalServerError)

		return
	}

	json.NewEncoder(w).Encode(NewGizmoResponse("Gizmo retrieved.", NewGizmoFromDBGizmo(gizmo)))
}

func (s *GizmoServer) PostGizmo(w http.ResponseWriter, r *http.Request) {
	gizmo := &Gizmo{}

	err := json.NewDecoder(r.Body).Decode(gizmo)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Error decoding gizmo.")

		http.Error(w, "Error decoding gizmo.", http.StatusInternalServerError)

		return
	}

	err = s.db.InsertGizmo(NewDBGizmoFromGizmo(gizmo))
	if err != nil {
		log.Error().
			Err(err).
			Msg("Error inserting gizmo.")

		http.Error(w, "Error inserting gizmo.", http.StatusInternalServerError)

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

		http.Error(w, "Error converting resource to int.", http.StatusBadRequest)

		return
	}

	gizmo := &Gizmo{}

	err = json.NewDecoder(r.Body).Decode(gizmo)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Error decoding gizmo.")

		http.Error(w, "Error decoding gizmo.", http.StatusInternalServerError)

		return
	}

	err = s.db.UpdateGizmo(resource, NewDBGizmoFromGizmo(gizmo))
	if err == mongo.ErrNoDocuments {
		http.Error(w, "Gizmo not found.", http.StatusNotFound)

		return
	} else if err != nil {
		log.Error().
			Err(err).
			Msg("Error updating gizmo.")

		http.Error(w, "Error updating gizmo.", http.StatusInternalServerError)

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

		http.Error(w, "Error converting resource to int.", http.StatusBadRequest)

		return
	}

	err = s.db.DeleteGizmo(resource)
	if err == mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusNotFound)

		http.Error(w, "Gizmo not found.", http.StatusNotFound)

		return
	} else if err != nil {
		log.Error().
			Err(err).
			Msg("Error deleting gizmo.")

		http.Error(w, "Error deleting gizmo.", http.StatusInternalServerError)

		return
	}

	json.NewEncoder(w).Encode(NewGizmoResponse("Gizmo deleted.", nil))
}
