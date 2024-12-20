package handlers

import (
	"encoding/json"
	"net/http"

	"projectGo/bot"
)

// HandleCommand обрабатывает HTTP запросы с командами
func HandleCommand(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Только POST-запросы", http.StatusMethodNotAllowed)
		return
	}

	var data struct {
		Command string `json:"command"`
	}

	// Чтение и обработка команды
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Неверный формат JSON", http.StatusBadRequest)
		return
	}

	switch data.Command {
	case "распределить":
		bot.HandleDistribute(w)
	case "вернуть":
		bot.HandleReturn(w)
	case "перераспределить":
		bot.HandleRedistribute(w)
	case "удалить":
		bot.HandleDelete(w)
	case "мут":
		bot.HandleMute(w)
	case "размут":
		bot.HandleUnmute(w)
	default:
		http.Error(w, "Неизвестная команда", http.StatusBadRequest)
	}
}
