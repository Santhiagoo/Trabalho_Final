package handler

import (
	"encoding/json"
	"net/http"
	"strings"
	"errors"

	"TRABALHO_FINAL/internal/entity"
	"TRABALHO_FINAL/internal/service"
)

type EnemyHandler struct {
	EnemyService *service.EnemyService
}

func NewEnemyHandler(enemyService *service.EnemyService) *EnemyHandler {
	return &EnemyHandler{EnemyService: enemyService}
}

// Função para responder com um erro formatado
func respondWithError(w http.ResponseWriter, err error) {
	if strings.Contains(err.Error(), "internal server error") {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(entity.ErrorResponse{Message: err.Error()})
}

func (eh *EnemyHandler) AddEnemy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var enemy entity.Enemy
	if err := json.NewDecoder(r.Body).Decode(&enemy); err != nil {
		respondWithError(w, errors.New("invalid request body"))
		return
	}

	result, err := eh.EnemyService.AddEnemy(enemy.Nickname, enemy.Life, enemy.Attack, enemy.Defesa)
	if err != nil {
		respondWithError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (eh *EnemyHandler) LoadEnemies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	enemies, err := eh.EnemyService.LoadEnemies()
	if err != nil {
		respondWithError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(enemies)
}

func (eh *EnemyHandler) DeleteEnemy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")
	if id == "" {
		respondWithError(w, errors.New("invalid ID"))
		return
	}

	if err := eh.EnemyService.DeleteEnemy(id); err != nil {
		respondWithError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (eh *EnemyHandler) LoadEnemy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")
	if id == "" {
		respondWithError(w, errors.New("invalid ID"))
		return
	}

	enemy, err := eh.EnemyService.LoadEnemy(id)
	if err != nil {
		respondWithError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(enemy)
}

func (eh *EnemyHandler) SaveEnemy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")
	if id == "" {
		respondWithError(w, errors.New("invalid ID"))
		return
	}

	var enemy entity.Enemy
	if err := json.NewDecoder(r.Body).Decode(&enemy); err != nil {
		respondWithError(w, errors.New("invalid request body"))
		return
	}

	result, err := eh.EnemyService.SaveEnemy(id, enemy.Nickname, enemy.Life, enemy.Attack, enemy.Defesa)
	if err != nil {
		respondWithError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
