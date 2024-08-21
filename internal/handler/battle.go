package handler

import (
    "encoding/json"
    "net/http"

    "TRABALHO_FINAL/internal/entity"
    "TRABALHO_FINAL/internal/service"
)

type BattleHandler struct {
    BattleService *service.BattleService
}

func NewBattleHandler(battleService *service.BattleService) *BattleHandler {
    return &BattleHandler{BattleService: battleService}
}

func (bh *BattleHandler) CreateBattle(w http.ResponseWriter, r *http.Request) {
    var request struct {
        Player string `json:"Player"`
        Enemy  string `json:"Enemy"`
    }

    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    battle, result, err := bh.BattleService.CreateBattle(request.Player, request.Enemy)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    response := struct {
        Battle *entity.Battle `json:"battle"`
        Result string         `json:"result"`
    }{
        Battle: battle,
        Result: result,
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
    }
}
