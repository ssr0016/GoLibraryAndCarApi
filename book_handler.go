package rest

import (
	"encoding/json"
	"main/pkg/api/errors"
	"main/pkg/api/response"
	"main/pkg/routing"
	book "main/pkg/tsm/library"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) NewBookHanlder(r *routing.Router) {
	groubBO := r.Group("/api/bo/book")

	groubBO.GET("/", s.search)
	groubBO.POST("/", s.create)
	groubBO.GET("/detail/:id", s.get)
	groubBO.PUT("/id/Update", s.update)

}

func (s *Server) search(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		query       book.SearchBookQuery
		queryValues = r.URL.Query()
	)

	if len(queryValues[""]) != 1 {
		if err := decoder.Decode(&query, queryValues); err != nil {
			response.Error(http.StatusBadRequest, errors.ErrBadRequest).WriteTo(w)
			return
		}
	}

	result, err := s.Dependencies.LibrarySvc.Search(r.Context(), &query)
	if err != nil {
		response.Error(http.StatusInternalServerError, err).WriteTo(w)
		return
	}

	response.JSON(http.StatusOK, result).WriteTo(w)

}

func (s *Server) create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var cmd book.CreateBookCommand

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

	err = s.Dependencies.LibrarySvc.Create(r.Context(), &cmd)
	if err != nil {
		response.Error(http.StatusInternalServerError, err).WriteTo(w)
		return
	}

	response.Success("Book created successfully").WriteTo(w)

}

func (s *Server) get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ID, err := strconv.ParseInt(p.ByName("id"), 10, 64)
	if err != nil {
		response.Error(http.StatusBadRequest, errors.ErrBadRequest).WriteTo(w)
		return
	}

	if ID <= 0 {
		response.Error(http.StatusBadRequest, book.ErrAuthorIDInvalid).WriteTo(w)
	}
	result, err := s.Dependencies.LibrarySvc.Get(r.Context(), ID)
	if err != nil {
		response.Error(http.StatusInternalServerError, err).WriteTo(w)
		return
	}

	response.JSON(http.StatusOK, result).WriteTo(w)

}

func (s *Server) update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var cmd book.UpdateBookCommand

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

	err = s.Dependencies.LibrarySvc.Update(r.Context(), &cmd)
	if err != nil {
		response.Error(http.StatusInternalServerError, err).WriteTo(w)
		return
	}

	response.Success("Book updated successfully").WriteTo(w)

}
