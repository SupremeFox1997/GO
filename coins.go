package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func handleUnmute(w http.ResponseWriter) {
	// Получаем список пользователей из канала
	users, err := getUsersInChannel(RoomForDistributionID)
	if err != nil {
		http.Error(w, "Ошибка получения пользователей: "+err.Error(), http.StatusInternalServerError)
		return
	}

	for _, user := range users {
		// Получаем данные о пользователе
		member, err := botSession.GuildMember(ServerID, user)
		if err != nil {
			// Если не удалось получить информацию о пользователе, пропускаем
			fmt.Printf("Не удалось получить данные для пользователя %s: %v\n", user, err)
			continue
		}

		// Проверяем текущее состояние мьюта пользователя
		if !member.Mute {
			// Если пользователь уже не замьючен, пропускаем
			continue
		}

		// Снимаем мьют
		err = botSession.GuildMemberMute(ServerID, user, false)
		if err != nil {
			// Логируем ошибку для отладки, но не прерываем выполнение
			fmt.Printf("Ошибка снятия мьюта для пользователя %s: %v\n", user, err)
			continue
		}

		// Логируем успешное изменение статуса
		fmt.Printf("Пользователь %s размьючен.\n", user)
	}

	// Возвращаем успешный ответ
	w.Write([]byte("Пользователи успешно размьючены."))
}