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

// NewBattleHandler cria uma nova instância de BattleHandler com o serviço de batalha fornecido
func NewBattleHandler(battleService *service.BattleService) *BattleHandler {
    return &BattleHandler{BattleService: battleService}
}

// CreateBattle lida com a criação de uma batalha entre um jogador e um inimigo
func (bh *BattleHandler) CreateBattle(w http.ResponseWriter, r *http.Request) {
    // Estrutura para receber a requisição JSON
    var request struct {
        Player string `json:"Player"`
        Enemy  string `json:"Enemy"`
    }

    // Decodifica o corpo da requisição
    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Validação básica dos campos da requisição
    if request.Player == "" || request.Enemy == "" {
        http.Error(w, "Player and Enemy fields are required", http.StatusBadRequest)
        return
    }

    // Chama o serviço para criar a batalha
    battle, result, err := bh.BattleService.CreateBattle(request.Player, request.Enemy)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Estrutura de resposta com os dados da batalha e resultado
    response := struct {
        Battle *entity.Battle `json:"battle"`
        Result string         `json:"result"`
    }{
        Battle: battle,
        Result: result,
    }

    // Define o tipo de conteúdo da resposta como JSON
    w.Header().Set("Content-Type", "application/json")

    // Codifica e envia a resposta
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
        return
    }
}
