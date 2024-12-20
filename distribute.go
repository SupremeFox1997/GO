package handlers

import (
	"math/rand"
	"net/http"
	"projectGo/bot"
	"time"
)

// HandleDistribute распределяет пользователей по каналам
func HandleDistribute(w http.ResponseWriter) {
	// Получаем список пользователей в канале для распределения
	users, err := bot.GetUsersInChannel(bot.RoomForDistributionID)
	if err != nil {
		http.Error(w, "Ошибка получения пользователей: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Создаем каналы, если они еще не были созданы
	for _, name := range bot.TeamChannels {
		if _, exists := bot.TeamChannelIDs[name]; !exists {
			channel, err := botSession.CreateChannel(name, bot.CategoryID)
			if err != nil {
				http.Error(w, "Ошибка создания канала: "+err.Error(), http.StatusInternalServerError)
				return
			}
			bot.TeamChannelIDs[name] = channel.ID
		}
	}

	// Перемешиваем пользователей для случайного распределения
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(users), func(i, j int) {
		users[i], users[j] = users[j], users[i]
	})

	// Распределяем пользователей по каналам
	for i, user := range users {
		// Выбираем канал на основе индекса
		targetChannel := bot.TeamChannels[i%len(bot.TeamChannels)]
		channelID := bot.TeamChannelIDs[targetChannel]

		// Перемещаем пользователя в канал
		botSession.MoveUserToChannel(user, channelID)
	}

	w.Write([]byte("Пользователи распределены по каналам."))
}
