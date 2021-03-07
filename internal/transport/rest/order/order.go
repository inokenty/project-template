package order

import (
	"encoding/json"
	"net/http"
	"project-template/internal/iface"
	"project-template/internal/repo/model"
	"project-template/internal/transport/rest"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
)

type client struct {
	service iface.OrderService
}

func New(s iface.OrderService) iface.OrderHandler {
	return &client{service: s}
}

func (c *client) AddRoutes(router *chi.Mux) {
	router.Post("/orders", c.CreateOrder)
	router.Get("/orders/{id}", c.GetOrderByID)
	router.Get("/orders", c.ListOrders)
}

func (c *client) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var args CreateOrderArgs

	if err := json.NewDecoder(r.Body).Decode(&args); err != nil {
		rest.ReplyError(w, http.StatusBadRequest, errors.New("json.Decode"))

		return
	}

	createdAt := time.Now()

	id, err := c.service.CreateOrder(&model.Order{
		CreatedAt: &createdAt,
		UserID:    args.UserID,
		Sum:       args.Sum,
	})
	if err != nil {
		rest.ReplyError(w, http.StatusInternalServerError, errors.Wrap(err, "service.CreateOrder"))

		return
	}

	rest.ReplySuccess(w, http.StatusOK, CreateOrderReply{
		ID: id,
	})
}

func (c *client) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		rest.ReplyError(w, http.StatusBadRequest, errors.Wrap(err, "strconv.ParseUint"))

		return
	}

	order, err := c.service.GetOrderByID(uint(id))
	if err != nil {
		rest.ReplyError(w, http.StatusInternalServerError, errors.Wrap(err, "service.GetOrderByID"))

		return
	}

	rest.ReplySuccess(w, http.StatusOK, &GetOrderByIDReply{
		Order: &Order{
			ID:        order.ID,
			CreatedAt: order.CreatedAt,
			Sum:       order.Sum,
			UserID:    order.UserID,
		},
	})
}

func (c *client) ListOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := c.service.ListOrders()
	if err != nil {
		rest.ReplyError(w, http.StatusInternalServerError, errors.Wrap(err, "service.ListOrders"))

		return
	}

	items := make([]*Order, 0, len(orders))

	for _, order := range orders {
		items = append(items, &Order{
			ID:        order.ID,
			CreatedAt: order.CreatedAt,
			Sum:       order.Sum,
			UserID:    order.UserID,
		})
	}

	rest.ReplySuccess(w, http.StatusOK, ListOrdersReply{
		Items: items,
	})
}
