package presentation

import (
	"encoding/json"
	"net/http"

	"github.com/thinhlh/go-web-101/internal/core/dto"
)

type OrderController struct{}

func NewOrderController() OrderController {
	return OrderController{}
}

func (c OrderController) Ping(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(dto.NewSuccessResponse("pong"))
}

func (c OrderController) GetOrders(w http.ResponseWriter, r *http.Request) {

}
