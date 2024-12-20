package handlers

import (
	"net/http"

	"projectGo/bot"
)

// HandleDelete удаляет каналы и перемещает пользователей в исходную комнату
func HandleDelete(w http.ResponseWriter) {
	for name, channelID := range bot.TeamChannelIDs {
		users, _ := bot.GetUsersInChannel(channelID)
		for _, user := range users {
			// Перемещаем пользователя в исходную комнату
			roomForDistributionID := bot.RoomForDistributionID
			botSession.MoveUserToChannel(user, roomForDistributionID)
		}
		botSession.DeleteChannel(channelID)
		delete(bot.TeamChannelIDs, name)
	}
	w.Write([]byte("Каналы удалены, пользователи перемещены в комнату 'Для распределения'."))
}
