package bot

// GetUsersInChannel получает список пользователей, находящихся в канале
func GetUsersInChannel(channelID string) ([]string, error) {
	members, err := botSession.GuildMembers(ServerID, "", 1000)
	if err != nil {
		return nil, err
	}

	var users []string
	for _, member := range members {
		voiceState, err := botSession.State.VoiceState(ServerID, member.User.ID)
		if err != nil {
			continue
		}

		if voiceState != nil && voiceState.ChannelID == channelID {
			users = append(users, member.User.ID)
		}
	}
	return users, nil
}
