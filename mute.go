package handlers

import (
	"fmt"
	"net/http"
	"project/bot"
	"strings"
)

// HandleMute мутирует пользователей в канале
func HandleMute(w http.ResponseWriter) {
	// Получаем список пользователей из канала
	users, err := bot.GetUsersInChannel(bot.RoomForDistributionID)
	if err != nil {
		http.Error(w, "Ошибка получения пользователей: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Разделяем исключаемые роли по запятой
	roleExclusions := strings.Split(bot.RoleExclusions, ",")

	for _, user := range users {
		// Получаем данные о пользователе
		member, err := botSession.GuildMember(bot.ServerID, user)
		if err != nil {
			// Если не удалось получить информацию о пользователе, пропускаем
			continue
		}

		// Проверяем, есть ли у пользователя роли из списка исключений
		mute := true // Изначально предполагаем, что нужно замьютить
		for _, role := range member.Roles {
			if contains(roleExclusions, role) {
				mute = false // Если роль найдена, пользователь не должен быть замьючен
				break
			}
		}

		// Применяем новый статус мьюта
		err = botSession.GuildMemberMute(bot.ServerID, user, mute)
		if err != nil {
			// Логируем ошибку для отладки, но не прерываем выполнение
			fmt.Printf("Ошибка изменения состояния мьюта для пользователя %s: %v\n", user, err)
		}
	}

	// Возвращаем успешный ответ
	w.Write([]byte("Пользователи успешно замьючены/размьючены."))
}
