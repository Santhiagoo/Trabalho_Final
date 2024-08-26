package handler

import (
	"encoding/json"
	"net/http"
	//"strings"
	"errors"

	"TRABALHO_FINAL/internal/entity"
	"TRABALHO_FINAL/internal/service"
)

type PlayerHandler struct {
	PlayerService *service.PlayerService
}

func NewPlayerHandler(playerService *service.PlayerService) *PlayerHandler {
	return &PlayerHandler{PlayerService: playerService}
}

// Função auxiliar para lidar com erros de resposta

func (ph *PlayerHandler) AddPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var player entity.Player
	if err := json.NewDecoder(r.Body).Decode(&player); err != nil {
		respondWithError(w, err)
		return
	}

	result, err := ph.PlayerService.AddPlayer(player.Nickname, player.Life, player.Attack, player.Defesa)
	if err != nil {
		respondWithError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (ph *PlayerHandler) LoadPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	players, err := ph.PlayerService.LoadPlayers()
	if err != nil {
		respondWithError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(players)
}

func (ph *PlayerHandler) DeletePlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")
	if id == "" {
		respondWithError(w, errors.New("invalid ID"))
		return
	}

	if err := ph.PlayerService.DeletePlayer(id); err != nil {
		respondWithError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (ph *PlayerHandler) LoadPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")
	if id == "" {
		respondWithError(w, errors.New("invalid ID"))
		return
	}

	player, err := ph.PlayerService.LoadPlayer(id)
	if err != nil {
		respondWithError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(player)
}

func (ph *PlayerHandler) SavePlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")
	if id == "" {
		respondWithError(w, errors.New("invalid ID"))
		return
	}

	var player entity.Player
	if err := json.NewDecoder(r.Body).Decode(&player); err != nil {
		respondWithError(w, err)
		return
	}

	result, err := ph.PlayerService.SavePlayer(id, player.Nickname, player.Life, player.Attack, player.Defesa)
	if err != nil {
		respondWithError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
