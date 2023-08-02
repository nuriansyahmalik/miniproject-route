package handlers

import (
	"github.com/evermos/boilerplate-go/internal/domain/menu"
	"github.com/evermos/boilerplate-go/shared/failure"
	"github.com/evermos/boilerplate-go/transport/http/middleware"
	"github.com/evermos/boilerplate-go/transport/http/response"
	"github.com/go-chi/chi"
	"github.com/gofrs/uuid"
	"net/http"
)

type MenuHandler struct {
	MenuService    menu.MenuService
	AuthMiddleware *middleware.Authentication
}

func ProvideMenuHandler(menuService menu.MenuService, authMiddleware *middleware.Authentication) MenuHandler {
	return MenuHandler{MenuService: menuService,
		AuthMiddleware: authMiddleware}
}

func (h *MenuHandler) Router(r chi.Router) {
	r.Route("/menu", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(h.AuthMiddleware.AuthMiddleware)
			r.Get("/{id}", h.ResolveMenuByID)
			r.Get("/", h.ResolveMenuByQuery)
		})
	})
}

func (h *MenuHandler) ResolveMenuByID(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")
	id, err := uuid.FromString(idString)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}
	menu, err := h.MenuService.ResolveByID(id)
	if err != nil {
		response.WithError(w, err)
		return
	}
	w.Header().Set("ID", menu.ID)
	response.WithJSON(w, http.StatusOK, menu)
}

func (h *MenuHandler) ResolveMenuByQuery(w http.ResponseWriter, r *http.Request) {
	idS := r.URL.Query().Get("id")
	if idS == "" {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}
	id, err := uuid.FromString(idS)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}
	menu, err := h.MenuService.ResolveByID(id)
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, menu)
}
