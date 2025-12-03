package middleware

import "github.com/go-telegram/bot/models"

func ToModelsUpdate(update models.Update) ModelsUpdate {
	result := ModelsUpdate{
		Update_id: int(update.ID),
	}

	if update.Message != nil {
		msg := toModelsMessage(update.Message)
		result.Message = &msg
	}

	if update.EditedMessage != nil {
		msg := toModelsMessage(update.EditedMessage)
		result.Edited_message = &msg
	}

	if update.ChannelPost != nil {
		msg := toModelsMessage(update.ChannelPost)
		result.Channel_post = &msg
	}

	if update.EditedChannelPost != nil {
		msg := toModelsMessage(update.EditedChannelPost)
		result.Edited_channel_post = &msg
	}

	return result
}

func toModelsMessage(msg *models.Message) ModelsMessage {
	result := ModelsMessage{
		Message_id: msg.ID,
		Date:       msg.Date,
		Text:       msg.Text,
		Caption:    msg.Caption,
	}

	result.Chat = toModelsChat(&msg.Chat)

	if msg.From != nil {
		user := toModelsUser(msg.From)
		result.From = &user
	}

	return result
}

func toModelsChat(chat *models.Chat) ModelsChat {
	return ModelsChat{
		ID:         int(chat.ID),
		Type:       string(chat.Type),
		Title:      chat.Title,
		Username:   chat.Username,
		First_name: chat.FirstName,
		Last_name:  chat.LastName,
		Is_forum:   chat.IsForum,
	}
}

func toModelsUser(user *models.User) ModelsUser {
	return ModelsUser{
		ID:                          int(user.ID),
		Is_bot:                      user.IsBot,
		First_name:                  user.FirstName,
		Last_name:                   user.LastName,
		Username:                    user.Username,
		Language_code:               user.LanguageCode,
		Is_premium:                  user.IsPremium,
		Added_to_attachment_menu:    user.AddedToAttachmentMenu,
		Can_join_groups:             user.CanJoinGroups,
		Can_read_all_group_messages: user.CanReadAllGroupMessages,
		Support_inline_queries:      user.SupportInlineQueries,
	}
}
