package rest

import (
	"encoding/json"
	"main/pkg/api/errors"
	"main/pkg/api/response"
	"main/pkg/routing"
	"main/pkg/tsm/carbrand"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) NewCarBrandHandler(r *routing.Router) {
	groubBO := r.Group("/api/bo/carbrand")

	groubBO.GET("/", s.searchCarBrand)
	groubBO.POST("/", s.create)
}

func (s *Server) searchCarBrand(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		query       carbrand.SearchCarBrandQuery
		queryValues = r.URL.Query()
	)

	if len(queryValues[""]) != 1 {
		if err := decoder.Decode(&query, queryValues); err != nil {
			response.Error(http.StatusBadRequest, errors.ErrBadRequest).WriteTo(w)
			return
		}
	}

	result, err := s.Dependencies.CarBrandSvc.Search(r.Context(), &query)
	if err != nil {
		response.Error(http.StatusInternalServerError, err).WriteTo(w)
		return
	}

	response.JSON(http.StatusOK, result).WriteTo(w)
}

func (s *Server) created(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var cmd carbrand.CreateCarBrandCommand

	err := json.NewDecoder(r.Body).Decode(&cmd)
	if err != nil {
		response.Error(http.StatusBadRequest, errors.ErrBadRequest).WriteTo(w)
		return
	}

	err = cmd.Validate()
	if err != nil {
		response.Error(http.StatusBadRequest, err).WriteTo(w)
		return
	}

	err = s.Dependencies.CarBrandSvc.Create(r.Context(), &cmd)
	if err != nil {
		response.Error(http.StatusInternalServerError, err).WriteTo(w)
		return
	}

	response.Success("Car brand created").WriteTo(w)
}
