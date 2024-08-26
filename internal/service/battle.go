package service

import (
	"errors"
	"fmt"
	"strconv"

	"TRABALHO_FINAL/internal/entity"
	"TRABALHO_FINAL/internal/repository"
)

type BattleService struct {
	PlayerRepository repository.PlayerRepository
	EnemyRepository  repository.EnemyRepository
	BattleRepository repository.BattleRepository
}

func NewBattleService(playerRepo repository.PlayerRepository, enemyRepo repository.EnemyRepository, battleRepo repository.BattleRepository) *BattleService {
	return &BattleService{
		PlayerRepository: playerRepo,
		EnemyRepository:  enemyRepo,
		BattleRepository: battleRepo,
	}
}

func (bs *BattleService) CreateBattle(playerNickname, enemyNickname string) (*entity.Battle, string, error) {
	player, err := bs.PlayerRepository.LoadPlayerByNickname(playerNickname)
	if err != nil || player == nil {
		return nil, "", errors.New("jogador não encontrado")
	}

	enemy, err := bs.EnemyRepository.LoadEnemyByNickname(enemyNickname)
	if err != nil || enemy == nil {
		return nil, "", errors.New("inimigo não encontrado")
	}

	// Verifica e solicita valor da vida do jogador
	if player.Life <= 0 {
		fmt.Println("A vida do jogador está zerada. Por favor, insira um valor para a vida do jogador:")
		var lifeInput string
		fmt.Scanln(&lifeInput)
		lifeValue, err := strconv.Atoi(lifeInput)
		if err != nil || lifeValue <= 0 {
			return nil, "", errors.New("valor inválido para a vida do jogador")
		}
		player.Life = lifeValue
	}

	// Verifica e solicita valor da vida do inimigo
	if enemy.Life <= 0 {
		fmt.Println("A vida do inimigo está zerada. Por favor, insira um valor para a vida do inimigo:")
		var lifeInput string
		fmt.Scanln(&lifeInput)
		lifeValue, err := strconv.Atoi(lifeInput)
		if err != nil || lifeValue <= 0 {
			return nil, "", errors.New("valor inválido para a vida do inimigo")
		}
		enemy.Life = lifeValue
	}

	battle := entity.NewBattle(player.ID, enemy.ID, player.Nickname, enemy.Nickname)
	dice := battle.DiceThrown

	var result string

	if dice <= 3 {
		// Inimigo ataca o jogador
		damage := enemy.Attack - player.Defesa
		if damage < 0 {
			damage = 0
		}
		player.Life -= damage
		if player.Life < 0 {
			player.Life = 0
		}

		// Implementação da funcionalidade de contra-ataque
		if damage >= 8 { // Alterado para 8
			counterDamage := int(float64(damage) * 0.3) // 30% do dano recebido
			enemy.Life -= counterDamage
			if enemy.Life < 0 {
				enemy.Life = 0
			}
			result += "Contra-ataque ativado! Dano causado ao inimigo: " + strconv.Itoa(counterDamage) + "\n"
		}

		if err := bs.PlayerRepository.SavePlayer(player.ID, player); err != nil {
			return nil, "", errors.New("falha ao atualizar a vida do jogador")
		}

		if err := bs.EnemyRepository.SaveEnemy(enemy.ID, enemy); err != nil {
			return nil, "", errors.New("falha ao atualizar a vida do inimigo")
		}

		// Exibir quem atacou, resultados do ataque, status do jogador e do inimigo
		result += "Inimigo atacou! Dano final causado: " + strconv.Itoa(damage) + " --- " +
			"Atributos do Jogador: Vida: " + strconv.Itoa(player.Life) +
			", Defesa: " + strconv.Itoa(player.Defesa) +
			", Ataque: " + strconv.Itoa(player.Attack) + " --- " +
			"Atributos do Inimigo: Vida: " + strconv.Itoa(enemy.Life) +
			", Defesa: " + strconv.Itoa(enemy.Defesa) +
			", Ataque: " + strconv.Itoa(enemy.Attack) + "\n"

	} else {
		// Jogador ataca o inimigo
		damage := player.Attack - enemy.Defesa
		if damage < 0 {
			damage = 0
		}
		enemy.Life -= damage
		if enemy.Life < 0 {
			enemy.Life = 0
		}

		// Implementação da funcionalidade de contra-ataque
		if damage >= 8 { // Alterado para 8
			counterDamage := int(float64(damage) * 0.3) // 30% do dano recebido
			player.Life -= counterDamage
			if player.Life < 0 {
				player.Life = 0
			}
			result += "Contra-ataque ativado! Dano causado ao jogador: " + strconv.Itoa(counterDamage) + "\n"
		}

		if err := bs.PlayerRepository.SavePlayer(player.ID, player); err != nil {
			return nil, "", errors.New("falha ao atualizar a vida do jogador")
		}

		if err := bs.EnemyRepository.SaveEnemy(enemy.ID, enemy); err != nil {
			return nil, "", errors.New("falha ao atualizar a vida do inimigo")
		}

		// Exibir quem atacou, resultados do ataque, status do jogador e do inimigo
		result += "Jogador atacou! Dano final causado: " + strconv.Itoa(damage) + " --- " +
			"Atributos do Jogador: Vida: " + strconv.Itoa(player.Life) +
			", Defesa: " + strconv.Itoa(player.Defesa) +
			", Ataque: " + strconv.Itoa(player.Attack) + " --- " +
			"Atributos do Inimigo: Vida: " + strconv.Itoa(enemy.Life) +
			", Defesa: " + strconv.Itoa(enemy.Defesa) +
			", Ataque: " + strconv.Itoa(enemy.Attack) + "\n"
	}

	// Verificar se alguém venceu a batalha
	if player.Life == 0 {
		battle.Result = "Inimigo venceu"
		result = "Inimigo venceu a batalha --- " + result
	} else if enemy.Life == 0 {
		battle.Result = "Jogador venceu"
		result = "Jogador venceu a batalha --- " + result
	} else {
		battle.Result = "A batalha continua --- " + result
	}

	if _, err := bs.BattleRepository.AddBattle(battle); err != nil {
		return nil, "", err
	}

	return battle, result, nil
}
