// Code generated from jsonrpc schema by rpcgen v2.5.0; DO NOT EDIT.

package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/vmkteam/appkit"
	"github.com/vmkteam/zenrpc/v2"
)

const name = "logshub"

var (
	// Always import time package. Generated models can contain time.Time fields.
	_ time.Time
)

type Client struct {
	rpcClient *rpcClient

	Logs *svcLogs
}

func NewClient(endpoint string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: time.Second * 30}
	}
	c := &Client{
		rpcClient: newRPCClient(endpoint, httpClient),
	}

	c.Logs = newClientLogs(c.rpcClient)

	return c
}

type Log struct {
	Error_code int    `json:"error_code"`
	Message    string `json:"message"`
	// Params     `json:"params"`
	Tg_user_id int    `json:"tg_user_id"`
	Timestamp  string `json:"timestamp"`
	Type       string `json:"type"`
}

type LogsService struct {
	Logs    []Log           `json:"logs"`
	Service ServiceResponse `json:"service"`
}

type ModelsAnimation struct {
	Duration       int              `json:"duration"`
	File_id        string           `json:"file_id"`
	File_name      string           `json:"file_name"`
	File_size      int              `json:"file_size"`
	File_unique_id string           `json:"file_unique_id"`
	Height         int              `json:"height"`
	Mime_type      string           `json:"mime_type"`
	Thumbnail      *ModelsPhotoSize `json:"thumbnail,omitempty"`
	Width          int              `json:"width"`
}

type ModelsAudio struct {
	Duration       int              `json:"duration"`
	File_id        string           `json:"file_id"`
	File_name      string           `json:"file_name"`
	File_size      int              `json:"file_size"`
	File_unique_id string           `json:"file_unique_id"`
	Mime_type      string           `json:"mime_type"`
	Performer      string           `json:"performer"`
	Thumbnail      *ModelsPhotoSize `json:"thumbnail,omitempty"`
	Title          string           `json:"title"`
}

type ModelsBackgroundFill struct {
	FreeformGradient *ModelsBackgroundFillFreeformGradient `json:"FreeformGradient,omitempty"`
	Gradient         *ModelsBackgroundFillGradient         `json:"Gradient,omitempty"`
	Solid            *ModelsBackgroundFillSolid            `json:"Solid,omitempty"`
	Type             string                                `json:"Type"`
}

type ModelsBackgroundFillFreeformGradient struct {
	Colors []int  `json:"colors"`
	Type   string `json:"type"`
}

type ModelsBackgroundFillGradient struct {
	Bottom_color   int    `json:"bottom_color"`
	Rotation_angle int    `json:"rotation_angle"`
	Top_color      int    `json:"top_color"`
	Type           string `json:"type"`
}

type ModelsBackgroundFillSolid struct {
	Color int    `json:"color"`
	Type  string `json:"type"`
}

type ModelsBackgroundType struct {
	Fill      *ModelsBackgroundTypeFill      `json:"Fill,omitempty"`
	Pattern   *ModelsBackgroundTypePattern   `json:"Pattern,omitempty"`
	Theme     *ModelsBackgroundTypeChatTheme `json:"Theme,omitempty"`
	Type      string                         `json:"Type"`
	Wallpaper *ModelsBackgroundTypeWallpaper `json:"Wallpaper,omitempty"`
}

type ModelsBackgroundTypeChatTheme struct {
	Theme_name string `json:"theme_name"`
	Type       string `json:"type"`
}

type ModelsBackgroundTypeFill struct {
	Dark_theme_dimming int                  `json:"dark_theme_dimming"`
	Fill               ModelsBackgroundFill `json:"fill"`
	Type               string               `json:"type"`
}

type ModelsBackgroundTypePattern struct {
	Document    ModelsDocument       `json:"document"`
	Fill        ModelsBackgroundFill `json:"fill"`
	Intensity   int                  `json:"intensity"`
	Is_inverted bool                 `json:"is_inverted"`
	Is_moving   bool                 `json:"is_moving"`
	Type        string               `json:"type"`
}

type ModelsBackgroundTypeWallpaper struct {
	Dark_theme_dimming int            `json:"dark_theme_dimming"`
	Document           ModelsDocument `json:"document"`
	Is_blurred         bool           `json:"is_blurred"`
	Is_moving          bool           `json:"is_moving"`
	Type               string         `json:"type"`
}

type ModelsBusinessBotRights struct {
	Can_change_gift_settings       bool `json:"can_change_gift_settings"`
	Can_convert_gifts_to_stars     bool `json:"can_convert_gifts_to_stars"`
	Can_delete_all_messages        bool `json:"can_delete_all_messages"`
	Can_delete_outgoing_messages   bool `json:"can_delete_outgoing_messages"`
	Can_edit_bio                   bool `json:"can_edit_bio"`
	Can_edit_name                  bool `json:"can_edit_name"`
	Can_edit_profile_photo         bool `json:"can_edit_profile_photo"`
	Can_edit_username              bool `json:"can_edit_username"`
	Can_manage_stories             bool `json:"can_manage_stories"`
	Can_read_messages              bool `json:"can_read_messages"`
	Can_reply                      bool `json:"can_reply"`
	Can_transfer_and_upgrade_gifts bool `json:"can_transfer_and_upgrade_gifts"`
	Can_transfer_stars             bool `json:"can_transfer_stars"`
	Can_view_gifts_and_stars       bool `json:"can_view_gifts_and_stars"`
}

type ModelsBusinessConnection struct {
	Date         int                      `json:"date"`
	ID           string                   `json:"id"`
	Is_enabled   bool                     `json:"is_enabled"`
	Rights       *ModelsBusinessBotRights `json:"rights,omitempty"`
	User         ModelsUser               `json:"user"`
	User_chat_id int                      `json:"user_chat_id"`
}

type ModelsBusinessMessagesDeleted struct {
	Business_connection_id string     `json:"business_connection_id"`
	Chat                   ModelsChat `json:"chat"`
	Message_ids            []int      `json:"message_ids"`
}

type ModelsCallbackGame struct {
	Chat_id              int  `json:"chat_id"`
	Disable_edit_message bool `json:"disable_edit_message"`
	Force                bool `json:"force"`
	Inline_message_id    int  `json:"inline_message_id"`
	Message_id           int  `json:"message_id"`
	Score                int  `json:"score"`
	User_id              int  `json:"user_id"`
}

type ModelsCallbackQuery struct {
	Chat_instance     string                         `json:"chat_instance"`
	Data              string                         `json:"data"`
	From              ModelsUser                     `json:"from"`
	Game_short_name   string                         `json:"game_short_name"`
	ID                string                         `json:"id"`
	Inline_message_id string                         `json:"inline_message_id"`
	Message           ModelsMaybeInaccessibleMessage `json:"message"`
}

type ModelsChat struct {
	First_name         string `json:"first_name"`
	ID                 int    `json:"id"`
	Is_direct_messages bool   `json:"is_direct_messages"`
	Is_forum           bool   `json:"is_forum"`
	Last_name          string `json:"last_name"`
	Title              string `json:"title"`
	Type               string `json:"type"`
	Username           string `json:"username"`
}

type ModelsChatBackground struct {
	Type ModelsBackgroundType `json:"type"`
}

type ModelsChatBoost struct {
	Add_date        int                   `json:"add_date"`
	Boost_id        string                `json:"boost_id"`
	Expiration_date int                   `json:"expiration_date"`
	Source          ModelsChatBoostSource `json:"source"`
}

type ModelsChatBoostAdded struct {
	Boost_count int `json:"boost_count"`
}

type ModelsChatBoostRemoved struct {
	Boost_id    string                `json:"boost_id"`
	Chat        ModelsChat            `json:"chat"`
	Remove_date int                   `json:"remove_date"`
	Source      ModelsChatBoostSource `json:"source"`
}

type ModelsChatBoostSource struct {
	ChatBoostSourceGiftCode *ModelsChatBoostSourceGiftCode `json:"ChatBoostSourceGiftCode,omitempty"`
	ChatBoostSourceGiveaway *ModelsChatBoostSourceGiveaway `json:"ChatBoostSourceGiveaway,omitempty"`
	ChatBoostSourcePremium  *ModelsChatBoostSourcePremium  `json:"ChatBoostSourcePremium,omitempty"`
	Source                  string                         `json:"Source"`
}

type ModelsChatBoostSourceGiftCode struct {
	// always “gift_code”
	Source string     `json:"source"`
	User   ModelsUser `json:"user"`
}

type ModelsChatBoostSourceGiveaway struct {
	Giveaway_message_id int  `json:"giveaway_message_id"`
	Is_unclaimed        bool `json:"is_unclaimed"`
	Prize_star_count    int  `json:"prize_star_count"`
	// always “giveaway”
	Source string     `json:"source"`
	User   ModelsUser `json:"user"`
}

type ModelsChatBoostSourcePremium struct {
	// always “premium”
	Source string     `json:"source"`
	User   ModelsUser `json:"user"`
}

type ModelsChatBoostUpdated struct {
	Boost ModelsChatBoost `json:"boost"`
	Chat  ModelsChat      `json:"chat"`
}

type ModelsChatInviteLink struct {
	Creates_join_request       bool       `json:"creates_join_request"`
	Creator                    ModelsUser `json:"creator"`
	Expire_date                int        `json:"expire_date"`
	Invite_link                string     `json:"invite_link"`
	Is_primary                 bool       `json:"is_primary"`
	Is_revoked                 bool       `json:"is_revoked"`
	Member_limit               int        `json:"member_limit"`
	Name                       string     `json:"name"`
	Pending_join_request_count int        `json:"pending_join_request_count"`
}

type ModelsChatJoinRequest struct {
	Bio          string                `json:"bio"`
	Chat         ModelsChat            `json:"chat"`
	Date         int                   `json:"date"`
	From         ModelsUser            `json:"from"`
	Invite_link  *ModelsChatInviteLink `json:"invite_link,omitempty"`
	User_chat_id int                   `json:"user_chat_id"`
}

type ModelsChatMember struct {
	Administrator *ModelsChatMemberAdministrator `json:"Administrator,omitempty"`
	Banned        *ModelsChatMemberBanned        `json:"Banned,omitempty"`
	Left          *ModelsChatMemberLeft          `json:"Left,omitempty"`
	Member        *ModelsChatMemberMember        `json:"Member,omitempty"`
	Owner         *ModelsChatMemberOwner         `json:"Owner,omitempty"`
	Restricted    *ModelsChatMemberRestricted    `json:"Restricted,omitempty"`
	Type          string                         `json:"Type"`
}

type ModelsChatMemberAdministrator struct {
	Can_be_edited              bool   `json:"can_be_edited"`
	Can_change_info            bool   `json:"can_change_info"`
	Can_delete_messages        bool   `json:"can_delete_messages"`
	Can_delete_stories         bool   `json:"can_delete_stories"`
	Can_edit_messages          bool   `json:"can_edit_messages"`
	Can_edit_stories           bool   `json:"can_edit_stories"`
	Can_invite_users           bool   `json:"can_invite_users"`
	Can_manage_chat            bool   `json:"can_manage_chat"`
	Can_manage_direct_messages bool   `json:"can_manage_direct_messages"`
	Can_manage_topics          bool   `json:"can_manage_topics"`
	Can_manage_video_chats     bool   `json:"can_manage_video_chats"`
	Can_pin_messages           bool   `json:"can_pin_messages"`
	Can_post_messages          bool   `json:"can_post_messages"`
	Can_post_stories           bool   `json:"can_post_stories"`
	Can_promote_members        bool   `json:"can_promote_members"`
	Can_restrict_members       bool   `json:"can_restrict_members"`
	Custom_title               string `json:"custom_title"`
	Is_anonymous               bool   `json:"is_anonymous"`
	// The member's status in the chat, always “administrator”
	Status string     `json:"status"`
	User   ModelsUser `json:"user"`
}

type ModelsChatMemberBanned struct {
	// The member's status in the chat, always “kicked”
	Status     string      `json:"status"`
	Until_date int         `json:"until_date"`
	User       *ModelsUser `json:"user,omitempty"`
}

type ModelsChatMemberLeft struct {
	// The member's status in the chat, always “left”
	Status string      `json:"status"`
	User   *ModelsUser `json:"user,omitempty"`
}

type ModelsChatMemberMember struct {
	// The member's status in the chat, always “member”
	Status     string      `json:"status"`
	Until_date int         `json:"until_date"`
	User       *ModelsUser `json:"user,omitempty"`
}

type ModelsChatMemberOwner struct {
	Custom_title string `json:"custom_title"`
	Is_anonymous bool   `json:"is_anonymous"`
	// The member's status in the chat, always “creator”
	Status string      `json:"status"`
	User   *ModelsUser `json:"user,omitempty"`
}

type ModelsChatMemberRestricted struct {
	Can_add_web_page_previews bool `json:"can_add_web_page_previews"`
	Can_change_info           bool `json:"can_change_info"`
	Can_invite_users          bool `json:"can_invite_users"`
	Can_manage_topics         bool `json:"can_manage_topics"`
	Can_pin_messages          bool `json:"can_pin_messages"`
	Can_send_audios           bool `json:"can_send_audios"`
	Can_send_documents        bool `json:"can_send_documents"`
	Can_send_messages         bool `json:"can_send_messages"`
	Can_send_other_messages   bool `json:"can_send_other_messages"`
	Can_send_photos           bool `json:"can_send_photos"`
	Can_send_polls            bool `json:"can_send_polls"`
	Can_send_video_notes      bool `json:"can_send_video_notes"`
	Can_send_videos           bool `json:"can_send_videos"`
	Can_send_voice_notes      bool `json:"can_send_voice_notes"`
	Is_member                 bool `json:"is_member"`
	// The member's status in the chat, always “restricted”
	Status     string      `json:"status"`
	Until_date int         `json:"until_date"`
	User       *ModelsUser `json:"user,omitempty"`
}

type ModelsChatMemberUpdated struct {
	Chat                        ModelsChat            `json:"chat"`
	Date                        int                   `json:"date"`
	From                        ModelsUser            `json:"from"`
	Invite_link                 *ModelsChatInviteLink `json:"invite_link,omitempty"`
	New_chat_member             ModelsChatMember      `json:"new_chat_member"`
	Old_chat_member             ModelsChatMember      `json:"old_chat_member"`
	Via_chat_folder_invite_link bool                  `json:"via_chat_folder_invite_link"`
	Via_join_request            bool                  `json:"via_join_request"`
}

type ModelsChatShared struct {
	Chat_id    int               `json:"chat_id"`
	Photo      []ModelsPhotoSize `json:"photo"`
	Request_id int               `json:"request_id"`
	Title      string            `json:"title"`
	Username   string            `json:"username"`
}

type ModelsChecklist struct {
	Others_can_add_tasks          bool                  `json:"others_can_add_tasks"`
	Others_can_mark_tasks_as_done bool                  `json:"others_can_mark_tasks_as_done"`
	Tasks                         []ModelsChecklistTask `json:"tasks"`
	Title                         string                `json:"title"`
	Title_entities                []ModelsMessageEntity `json:"title_entities"`
}

type ModelsChecklistTask struct {
	Completed_by_user *ModelsUser           `json:"completed_by_user,omitempty"`
	Completion_date   int                   `json:"completion_date"`
	ID                int                   `json:"id"`
	Text              string                `json:"text"`
	Text_entities     []ModelsMessageEntity `json:"text_entities"`
}

type ModelsChecklistTasksAdded struct {
	Checklist_message *ModelsMessage        `json:"checklist_message,omitempty"`
	Tasks             []ModelsChecklistTask `json:"tasks"`
}

type ModelsChecklistTasksDone struct {
	Checklist_message           *ModelsMessage `json:"checklist_message,omitempty"`
	Marked_as_done_task_ids     []int          `json:"marked_as_done_task_ids"`
	Marked_as_not_done_task_ids []int          `json:"marked_as_not_done_task_ids"`
}

type ModelsChosenInlineResult struct {
	From              ModelsUser      `json:"from"`
	Inline_message_id string          `json:"inline_message_id"`
	Location          *ModelsLocation `json:"location,omitempty"`
	Query             string          `json:"query"`
	Result_id         string          `json:"result_id"`
}

type ModelsContact struct {
	First_name   string `json:"first_name"`
	Last_name    string `json:"last_name"`
	Phone_number string `json:"phone_number"`
	User_id      int    `json:"user_id"`
	Vcard        string `json:"vcard"`
}

type ModelsCopyTextButton struct {
	Text string `json:"text"`
}

type ModelsDice struct {
	Emoji string `json:"emoji"`
	Value int    `json:"value"`
}

type ModelsDirectMessagePriceChanged struct {
	Are_direct_messages_enabled bool `json:"are_direct_messages_enabled"`
	Direct_message_star_count   int  `json:"direct_message_star_count"`
}

type ModelsDirectMessagesTopic struct {
	Topic_id int         `json:"topic_id"`
	User     *ModelsUser `json:"user,omitempty"`
}

type ModelsDocument struct {
	File_id        string           `json:"file_id"`
	File_name      string           `json:"file_name"`
	File_size      int              `json:"file_size"`
	File_unique_id string           `json:"file_unique_id"`
	Mime_type      string           `json:"mime_type"`
	Thumbnail      *ModelsPhotoSize `json:"thumbnail,omitempty"`
}

type ModelsEncryptedCredentials struct {
	Data   string `json:"data"`
	Hash   string `json:"hash"`
	Secret string `json:"secret"`
}

type ModelsEncryptedPassportElement struct {
	Data         string               `json:"data"`
	Email        string               `json:"email"`
	Files        []ModelsPassportFile `json:"files"`
	Front_side   *ModelsPassportFile  `json:"front_side,omitempty"`
	Hash         string               `json:"hash"`
	Phone_number string               `json:"phone_number"`
	Reverse_side *ModelsPassportFile  `json:"reverse_side,omitempty"`
	Selfie       *ModelsPassportFile  `json:"selfie,omitempty"`
	Translation  []ModelsPassportFile `json:"translation"`
	Type         string               `json:"type"`
}

type ModelsExternalReplyInfo struct {
	Animation            *ModelsAnimation          `json:"animation,omitempty"`
	Audio                *ModelsAudio              `json:"audio,omitempty"`
	Chat                 *ModelsChat               `json:"chat,omitempty"`
	Checklist            *ModelsChecklist          `json:"checklist,omitempty"`
	Contact              *ModelsContact            `json:"contact,omitempty"`
	Dice                 *ModelsDice               `json:"dice,omitempty"`
	Document             *ModelsDocument           `json:"document,omitempty"`
	Game                 *ModelsGame               `json:"game,omitempty"`
	Giveaway             *ModelsGiveaway           `json:"giveaway,omitempty"`
	Giveaway_winners     *ModelsGiveawayWinners    `json:"giveaway_winners,omitempty"`
	Has_media_spoiler    bool                      `json:"has_media_spoiler"`
	Invoice              *ModelsInvoice            `json:"invoice,omitempty"`
	Link_preview_options *ModelsLinkPreviewOptions `json:"link_preview_options,omitempty"`
	Location             *ModelsLocation           `json:"location,omitempty"`
	Message_id           int                       `json:"message_id"`
	Origin               ModelsMessageOrigin       `json:"origin"`
	Paid_media           *ModelsPaidMediaInfo      `json:"paid_media,omitempty"`
	Photo                []ModelsPhotoSize         `json:"photo"`
	Poll                 *ModelsPoll               `json:"poll,omitempty"`
	Sticker              *ModelsSticker            `json:"sticker,omitempty"`
	Story                *ModelsStory              `json:"story,omitempty"`
	Venue                *ModelsVenue              `json:"venue,omitempty"`
	Video                *ModelsVideo              `json:"video,omitempty"`
	Video_note           *ModelsVideoNote          `json:"video_note,omitempty"`
	Voice                *ModelsVoice              `json:"voice,omitempty"`
}

type ModelsFile struct {
	File_id        string `json:"file_id"`
	File_path      string `json:"file_path"`
	File_size      int    `json:"file_size"`
	File_unique_id string `json:"file_unique_id"`
}

type ModelsForumTopicClosed struct {
}

type ModelsForumTopicCreated struct {
	Icon_color           int    `json:"icon_color"`
	Icon_custom_emoji_id string `json:"icon_custom_emoji_id"`
	Name                 string `json:"name"`
}

type ModelsForumTopicEdited struct {
	Icon_custom_emoji_id string `json:"icon_custom_emoji_id"`
	Name                 string `json:"name"`
}

type ModelsForumTopicReopened struct {
}

type ModelsGame struct {
	Animation     *ModelsAnimation      `json:"animation,omitempty"`
	Description   string                `json:"description"`
	Photo         []ModelsPhotoSize     `json:"photo"`
	Text          string                `json:"text"`
	Text_entities []ModelsMessageEntity `json:"text_entities"`
	Title         string                `json:"title"`
}

type ModelsGeneralForumTopicHidden struct {
}

type ModelsGeneralForumTopicUnhidden struct {
}

type ModelsGift struct {
	ID                 string        `json:"id"`
	Publisher_chat     *ModelsChat   `json:"publisher_chat,omitempty"`
	Remaining_count    int           `json:"remaining_count"`
	Star_count         int           `json:"star_count"`
	Sticker            ModelsSticker `json:"sticker"`
	Total_count        int           `json:"total_count"`
	Upgrade_star_count int           `json:"upgrade_star_count"`
}

type ModelsGiftInfo struct {
	Can_be_upgraded            bool                  `json:"can_be_upgraded"`
	Convert_star_count         int                   `json:"convert_star_count"`
	Entities                   []ModelsMessageEntity `json:"entities"`
	Gift                       ModelsGift            `json:"gift"`
	Is_private                 bool                  `json:"is_private"`
	Owned_gift_id              string                `json:"owned_gift_id"`
	Prepaid_upgrade_star_count int                   `json:"prepaid_upgrade_star_count"`
	Text                       string                `json:"text"`
}

type ModelsGiveaway struct {
	Chats                            []ModelsChat `json:"chats"`
	Country_codes                    []string     `json:"country_codes"`
	Has_public_winners               bool         `json:"has_public_winners"`
	Only_new_members                 bool         `json:"only_new_members"`
	Premium_subscription_month_count int          `json:"premium_subscription_month_count"`
	Prize_description                string       `json:"prize_description"`
	Prize_star_count                 int          `json:"prize_star_count"`
	Winner_count                     int          `json:"winner_count"`
	Winners_selection_date           int          `json:"winners_selection_date"`
}

type ModelsGiveawayCompleted struct {
	Giveaway_message      *ModelsMessage `json:"giveaway_message,omitempty"`
	Is_star_giveaway      bool           `json:"is_star_giveaway"`
	Unclaimed_prize_count int            `json:"unclaimed_prize_count"`
	Winner_count          int            `json:"winner_count"`
}

type ModelsGiveawayCreated struct {
	Prize_star_count int `json:"prize_star_count"`
}

type ModelsGiveawayWinners struct {
	Additional_chat_count            int          `json:"additional_chat_count"`
	Chat                             ModelsChat   `json:"chat"`
	Giveaway_message_id              int          `json:"giveaway_message_id"`
	Only_new_members                 bool         `json:"only_new_members"`
	Premium_subscription_month_count int          `json:"premium_subscription_month_count"`
	Prize_description                string       `json:"prize_description"`
	Prize_star_count                 int          `json:"prize_star_count"`
	Unclaimed_prize_count            int          `json:"unclaimed_prize_count"`
	Was_refunded                     bool         `json:"was_refunded"`
	Winner_count                     int          `json:"winner_count"`
	Winners                          []ModelsUser `json:"winners"`
	Winners_selection_date           int          `json:"winners_selection_date"`
}

type ModelsInaccessibleMessage struct {
	Chat       ModelsChat `json:"chat"`
	Date       int        `json:"date"`
	Message_id int        `json:"message_id"`
}

type ModelsInlineKeyboardButton struct {
	Callback_data                    string                             `json:"callback_data"`
	Callback_game                    *ModelsCallbackGame                `json:"callback_game,omitempty"`
	Copy_text                        ModelsCopyTextButton               `json:"copy_text"`
	Login_url                        *ModelsLoginURL                    `json:"login_url,omitempty"`
	Pay                              bool                               `json:"pay"`
	Switch_inline_query              string                             `json:"switch_inline_query"`
	Switch_inline_query_chosen_chat  *ModelsSwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`
	Switch_inline_query_current_chat string                             `json:"switch_inline_query_current_chat"`
	Text                             string                             `json:"text"`
	Url                              string                             `json:"url"`
	Web_app                          *ModelsWebAppInfo                  `json:"web_app,omitempty"`
}

type ModelsInlineKeyboardMarkup struct {
	Inline_keyboard []ModelsInlineKeyboardButton `json:"inline_keyboard"`
}

type ModelsInlineQuery struct {
	Chat_type string          `json:"chat_type"`
	From      *ModelsUser     `json:"from,omitempty"`
	ID        string          `json:"id"`
	Location  *ModelsLocation `json:"location,omitempty"`
	Offset    string          `json:"offset"`
	Query     string          `json:"query"`
}

type ModelsInvoice struct {
	Currency        string `json:"currency"`
	Description     string `json:"description"`
	Start_parameter string `json:"start_parameter"`
	Title           string `json:"title"`
	Total_amount    int    `json:"total_amount"`
}

type ModelsLinkPreviewOptions struct {
	Is_disabled        *bool   `json:"is_disabled,omitempty"`
	Prefer_large_media *bool   `json:"prefer_large_media,omitempty"`
	Prefer_small_media *bool   `json:"prefer_small_media,omitempty"`
	Show_above_text    *bool   `json:"show_above_text,omitempty"`
	Url                *string `json:"url,omitempty"`
}

type ModelsLocation struct {
	Heading                int     `json:"heading"`
	Horizontal_accuracy    float64 `json:"horizontal_accuracy"`
	Latitude               float64 `json:"latitude"`
	Live_period            int     `json:"live_period"`
	Longitude              float64 `json:"longitude"`
	Proximity_alert_radius int     `json:"proximity_alert_radius"`
}

type ModelsLoginURL struct {
	Bot_username         string `json:"bot_username"`
	Forward_text         string `json:"forward_text"`
	Request_write_access bool   `json:"request_write_access"`
	Url                  string `json:"url"`
}

type ModelsMaskPosition struct {
	Point   string  `json:"point"`
	Scale   float64 `json:"scale"`
	X_shift float64 `json:"x_shift"`
	Y_shift float64 `json:"y_shift"`
}

type ModelsMaybeInaccessibleMessage struct {
	InaccessibleMessage *ModelsInaccessibleMessage `json:"InaccessibleMessage,omitempty"`
	Message             *ModelsMessage             `json:"Message,omitempty"`
	Type                int                        `json:"Type"`
}

type ModelsMessage struct {
	Animation                         *ModelsAnimation                     `json:"animation,omitempty"`
	Audio                             *ModelsAudio                         `json:"audio,omitempty"`
	Author_signature                  string                               `json:"author_signature"`
	Boost_added                       *ModelsChatBoostAdded                `json:"boost_added,omitempty"`
	Business_connection_id            string                               `json:"business_connection_id"`
	Caption                           string                               `json:"caption"`
	Caption_entities                  []ModelsMessageEntity                `json:"caption_entities"`
	Channel_chat_created              bool                                 `json:"channel_chat_created"`
	Chat                              ModelsChat                           `json:"chat"`
	Chat_background_set               *ModelsChatBackground                `json:"chat_background_set,omitempty"`
	Chat_shared                       *ModelsChatShared                    `json:"chat_shared,omitempty"`
	Checklist                         *ModelsChecklist                     `json:"checklist,omitempty"`
	Checklist_tasks_added             *ModelsChecklistTasksAdded           `json:"checklist_tasks_added,omitempty"`
	Checklist_tasks_done              *ModelsChecklistTasksDone            `json:"checklist_tasks_done,omitempty"`
	Connected_website                 string                               `json:"connected_website"`
	Contact                           *ModelsContact                       `json:"contact,omitempty"`
	Date                              int                                  `json:"date"`
	Delete_chat_photo                 bool                                 `json:"delete_chat_photo"`
	Dice                              *ModelsDice                          `json:"dice,omitempty"`
	Direct_message_price_changed      *ModelsDirectMessagePriceChanged     `json:"direct_message_price_changed,omitempty"`
	Direct_messages_topic             *ModelsDirectMessagesTopic           `json:"direct_messages_topic,omitempty"`
	Document                          *ModelsDocument                      `json:"document,omitempty"`
	Edit_date                         int                                  `json:"edit_date"`
	Effect_id                         string                               `json:"effect_id"`
	Entities                          []ModelsMessageEntity                `json:"entities"`
	External_reply                    *ModelsExternalReplyInfo             `json:"external_reply,omitempty"`
	Forum_topic_closed                *ModelsForumTopicClosed              `json:"forum_topic_closed,omitempty"`
	Forum_topic_created               *ModelsForumTopicCreated             `json:"forum_topic_created,omitempty"`
	Forum_topic_edited                *ModelsForumTopicEdited              `json:"forum_topic_edited,omitempty"`
	Forum_topic_reopened              *ModelsForumTopicReopened            `json:"forum_topic_reopened,omitempty"`
	Forward_origin                    *ModelsMessageOrigin                 `json:"forward_origin,omitempty"`
	From                              *ModelsUser                          `json:"from,omitempty"`
	Game                              *ModelsGame                          `json:"game,omitempty"`
	General_forum_topic_hidden        *ModelsGeneralForumTopicHidden       `json:"general_forum_topic_hidden,omitempty"`
	General_forum_topic_unhidden      *ModelsGeneralForumTopicUnhidden     `json:"general_forum_topic_unhidden,omitempty"`
	Gift                              *ModelsGiftInfo                      `json:"gift,omitempty"`
	Giveaway                          *ModelsGiveaway                      `json:"giveaway,omitempty"`
	Giveaway_completed                *ModelsGiveawayCompleted             `json:"giveaway_completed,omitempty"`
	Giveaway_created                  *ModelsGiveawayCreated               `json:"giveaway_created,omitempty"`
	Giveaway_winners                  *ModelsGiveawayWinners               `json:"giveaway_winners,omitempty"`
	Group_chat_created                bool                                 `json:"group_chat_created"`
	Has_media_spoiler                 bool                                 `json:"has_media_spoiler"`
	Has_protected_content             bool                                 `json:"has_protected_content"`
	Invoice                           *ModelsInvoice                       `json:"invoice,omitempty"`
	Is_automatic_forward              bool                                 `json:"is_automatic_forward"`
	Is_from_offline                   bool                                 `json:"is_from_offline"`
	Is_paid_post                      bool                                 `json:"is_paid_post"`
	Is_topic_message                  bool                                 `json:"is_topic_message"`
	Left_chat_member                  *ModelsUser                          `json:"left_chat_member,omitempty"`
	Link_preview_options              *ModelsLinkPreviewOptions            `json:"link_preview_options,omitempty"`
	Location                          *ModelsLocation                      `json:"location,omitempty"`
	Media_group_id                    string                               `json:"media_group_id"`
	Message_auto_delete_timer_changed *ModelsMessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
	Message_id                        int                                  `json:"message_id"`
	Message_thread_id                 int                                  `json:"message_thread_id"`
	Migrate_from_chat_id              int                                  `json:"migrate_from_chat_id"`
	Migrate_to_chat_id                int                                  `json:"migrate_to_chat_id"`
	New_chat_members                  []ModelsUser                         `json:"new_chat_members"`
	New_chat_photo                    []ModelsPhotoSize                    `json:"new_chat_photo"`
	New_chat_title                    string                               `json:"new_chat_title"`
	Paid_media                        *ModelsPaidMediaInfo                 `json:"paid_media,omitempty"`
	Paid_message_price_changed        *ModelsPaidMessagePriceChanged       `json:"paid_message_price_changed,omitempty"`
	Paid_star_count                   int                                  `json:"paid_star_count"`
	Passport_data                     *ModelsPassportData                  `json:"passport_data,omitempty"`
	Photo                             []ModelsPhotoSize                    `json:"photo"`
	Pinned_message                    *ModelsMaybeInaccessibleMessage      `json:"pinned_message,omitempty"`
	Poll                              *ModelsPoll                          `json:"poll,omitempty"`
	Proximity_alert_triggered         *ModelsProximityAlertTriggered       `json:"proximity_alert_triggered,omitempty"`
	Quote                             *ModelsTextQuote                     `json:"quote,omitempty"`
	Refunded_payment                  *ModelsRefundedPayment               `json:"refunded_payment,omitempty"`
	Reply_markup                      *ModelsInlineKeyboardMarkup          `json:"reply_markup,omitempty"`
	Reply_to_checklist_task_id        int                                  `json:"reply_to_checklist_task_id"`
	Reply_to_message                  *ModelsMessage                       `json:"reply_to_message,omitempty"`
	Reply_to_store                    *ModelsStory                         `json:"reply_to_store,omitempty"`
	Sender_boost_count                int                                  `json:"sender_boost_count"`
	Sender_business_bot               *ModelsUser                          `json:"sender_business_bot,omitempty"`
	Sender_chat                       *ModelsChat                          `json:"sender_chat,omitempty"`
	Show_caption_above_media          bool                                 `json:"show_caption_above_media"`
	Sticker                           *ModelsSticker                       `json:"sticker,omitempty"`
	Story                             *ModelsStory                         `json:"story,omitempty"`
	Successful_payment                *ModelsSuccessfulPayment             `json:"successful_payment,omitempty"`
	Suggested_post_approval_failed    *ModelsSuggestedPostApprovalFailed   `json:"suggested_post_approval_failed,omitempty"`
	Suggested_post_approved           *ModelsSuggestedPostApproved         `json:"suggested_post_approved,omitempty"`
	Suggested_post_declined           *ModelsSuggestedPostDeclined         `json:"suggested_post_declined,omitempty"`
	Suggested_post_info               *ModelsSuggestedPostInfo             `json:"suggested_post_info,omitempty"`
	Suggested_post_paid               *ModelsSuggestedPostPaid             `json:"suggested_post_paid,omitempty"`
	Suggested_post_refunded           *ModelsSuggestedPostRefunded         `json:"suggested_post_refunded,omitempty"`
	Supergroup_chat_created           bool                                 `json:"supergroup_chat_created"`
	Text                              string                               `json:"text"`
	Unique_gift                       *ModelsUniqueGiftInfo                `json:"unique_gift,omitempty"`
	Users_shared                      *ModelsUsersShared                   `json:"users_shared,omitempty"`
	Venue                             *ModelsVenue                         `json:"venue,omitempty"`
	Via_bot                           *ModelsUser                          `json:"via_bot,omitempty"`
	Video                             *ModelsVideo                         `json:"video,omitempty"`
	Video_note                        *ModelsVideoNote                     `json:"video_note,omitempty"`
	Voice                             *ModelsVoice                         `json:"voice,omitempty"`
	Voice_chat_ended                  *ModelsVoiceChatEnded                `json:"voice_chat_ended,omitempty"`
	Voice_chat_participants_invited   *ModelsVoiceChatParticipantsInvited  `json:"voice_chat_participants_invited,omitempty"`
	Voice_chat_scheduled              *ModelsVoiceChatScheduled            `json:"voice_chat_scheduled,omitempty"`
	Voice_chat_started                *ModelsVoiceChatStarted              `json:"voice_chat_started,omitempty"`
	Web_app_data                      *ModelsWebAppData                    `json:"web_app_data,omitempty"`
	Write_access_allowed              *ModelsWriteAccessAllowed            `json:"write_access_allowed,omitempty"`
}

type ModelsMessageAutoDeleteTimerChanged struct {
	Message_auto_delete_time int `json:"message_auto_delete_time"`
}

type ModelsMessageEntity struct {
	Custom_emoji_id string      `json:"custom_emoji_id"`
	Language        string      `json:"language"`
	Length          int         `json:"length"`
	Offset          int         `json:"offset"`
	Type            string      `json:"type"`
	Url             string      `json:"url"`
	User            *ModelsUser `json:"user,omitempty"`
}

type ModelsMessageOrigin struct {
	MessageOriginChannel    *ModelsMessageOriginChannel    `json:"MessageOriginChannel,omitempty"`
	MessageOriginChat       *ModelsMessageOriginChat       `json:"MessageOriginChat,omitempty"`
	MessageOriginHiddenUser *ModelsMessageOriginHiddenUser `json:"MessageOriginHiddenUser,omitempty"`
	MessageOriginUser       *ModelsMessageOriginUser       `json:"MessageOriginUser,omitempty"`
	Type                    string                         `json:"Type"`
}

type ModelsMessageOriginChannel struct {
	Author_signature *string    `json:"author_signature,omitempty"`
	Chat             ModelsChat `json:"chat"`
	Date             int        `json:"date"`
	Message_id       int        `json:"message_id"`
	// always “channel”
	Type string `json:"type"`
}

type ModelsMessageOriginChat struct {
	Author_signature *string    `json:"author_signature,omitempty"`
	Date             int        `json:"date"`
	Sender_chat      ModelsChat `json:"sender_chat"`
	// always “chat”
	Type string `json:"type"`
}

type ModelsMessageOriginHiddenUser struct {
	Date             int    `json:"date"`
	Sender_user_name string `json:"sender_user_name"`
	// always “hidden_user”
	Type string `json:"type"`
}

type ModelsMessageOriginUser struct {
	Date        int        `json:"date"`
	Sender_user ModelsUser `json:"sender_user"`
	// always “user”
	Type string `json:"type"`
}

type ModelsMessageReactionCountUpdated struct {
	Chat       ModelsChat            `json:"chat"`
	Date       int                   `json:"date"`
	Message_id int                   `json:"message_id"`
	Reactions  []ModelsReactionCount `json:"reactions"`
}

type ModelsMessageReactionUpdated struct {
	Actor_chat   *ModelsChat          `json:"actor_chat,omitempty"`
	Chat         ModelsChat           `json:"chat"`
	Date         int                  `json:"date"`
	Message_id   int                  `json:"message_id"`
	New_reaction []ModelsReactionType `json:"new_reaction"`
	Old_reaction []ModelsReactionType `json:"old_reaction"`
	User         *ModelsUser          `json:"user,omitempty"`
}

type ModelsOrderInfo struct {
	Email            string                 `json:"email"`
	Name             string                 `json:"name"`
	Phone_number     string                 `json:"phone_number"`
	Shipping_address *ModelsShippingAddress `json:"shipping_address,omitempty"`
}

type ModelsPaidMedia struct {
	Photo   *ModelsPaidMediaPhoto   `json:"Photo,omitempty"`
	Preview *ModelsPaidMediaPreview `json:"Preview,omitempty"`
	Type    string                  `json:"Type"`
	Video   *ModelsPaidMediaVideo   `json:"Video,omitempty"`
}

type ModelsPaidMediaInfo struct {
	Paid_media []ModelsPaidMedia `json:"paid_media"`
	Star_count int               `json:"star_count"`
}

type ModelsPaidMediaPhoto struct {
	Type  string            `json:"Type"`
	Photo []ModelsPhotoSize `json:"photo"`
}

type ModelsPaidMediaPreview struct {
	Type     string `json:"Type"`
	Duration int    `json:"duration"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
}

type ModelsPaidMediaPurchased struct {
	From               ModelsUser `json:"from"`
	Paid_media_payload string     `json:"paid_media_payload"`
}

type ModelsPaidMediaVideo struct {
	Type  string      `json:"Type"`
	Video ModelsVideo `json:"video"`
}

type ModelsPaidMessagePriceChanged struct {
	Paid_message_star_count int `json:"paid_message_star_count"`
}

type ModelsPassportData struct {
	Credentials ModelsEncryptedCredentials       `json:"credentials"`
	Data        []ModelsEncryptedPassportElement `json:"data"`
}

type ModelsPassportFile struct {
	File_date      int    `json:"file_date"`
	File_id        string `json:"file_id"`
	File_size      int    `json:"file_size"`
	File_unique_id string `json:"file_unique_id"`
}

type ModelsPhotoSize struct {
	File_id        string `json:"file_id"`
	File_size      int    `json:"file_size"`
	File_unique_id string `json:"file_unique_id"`
	Height         int    `json:"height"`
	Width          int    `json:"width"`
}

type ModelsPoll struct {
	Allows_multiple_answers bool                  `json:"allows_multiple_answers"`
	Close_date              int                   `json:"close_date"`
	Correct_option_id       int                   `json:"correct_option_id"`
	Explanation             string                `json:"explanation"`
	Explanation_entities    []ModelsMessageEntity `json:"explanation_entities"`
	ID                      string                `json:"id"`
	Is_anonymous            bool                  `json:"is_anonymous"`
	Is_closed               bool                  `json:"is_closed"`
	Open_period             int                   `json:"open_period"`
	Options                 []ModelsPollOption    `json:"options"`
	Question                string                `json:"question"`
	Question_entities       []ModelsMessageEntity `json:"question_entities"`
	Total_voter_count       int                   `json:"total_voter_count"`
	Type                    string                `json:"type"`
}

type ModelsPollAnswer struct {
	Option_ids []int       `json:"option_ids"`
	Poll_id    string      `json:"poll_id"`
	User       *ModelsUser `json:"user,omitempty"`
	Voter_chat *ModelsChat `json:"voter_chat,omitempty"`
}

type ModelsPollOption struct {
	Text          string                `json:"text"`
	Text_entities []ModelsMessageEntity `json:"text_entities"`
	Voter_count   int                   `json:"voter_count"`
}

type ModelsPreCheckoutQuery struct {
	Currency           string           `json:"currency"`
	From               *ModelsUser      `json:"from,omitempty"`
	ID                 string           `json:"id"`
	Invoice_payload    string           `json:"invoice_payload"`
	Order_info         *ModelsOrderInfo `json:"order_info,omitempty"`
	Shipping_option_id string           `json:"shipping_option_id"`
	Total_amount       int              `json:"total_amount"`
}

type ModelsProximityAlertTriggered struct {
	Distance int        `json:"distance"`
	Traveler ModelsUser `json:"traveler"`
	Watcher  ModelsUser `json:"watcher"`
}

type ModelsReactionCount struct {
	Total_count int                `json:"total_count"`
	Type        ModelsReactionType `json:"type"`
}

type ModelsReactionType struct {
	ReactionTypeCustomEmoji *ModelsReactionTypeCustomEmoji `json:"ReactionTypeCustomEmoji,omitempty"`
	ReactionTypeEmoji       *ModelsReactionTypeEmoji       `json:"ReactionTypeEmoji,omitempty"`
	ReactionTypePaid        *ModelsReactionTypePaid        `json:"ReactionTypePaid,omitempty"`
	Type                    string                         `json:"Type"`
}

type ModelsReactionTypeCustomEmoji struct {
	Custom_emoji_id string `json:"custom_emoji_id"`
	Type            string `json:"type"`
}

type ModelsReactionTypeEmoji struct {
	Emoji string `json:"emoji"`
	Type  string `json:"type"`
}

type ModelsReactionTypePaid struct {
	Type string `json:"type"`
}

type ModelsRefundedPayment struct {
	Currency                   string `json:"currency"`
	Invoice_payload            string `json:"invoice_payload"`
	Provider_payment_charge_id string `json:"provider_payment_charge_id"`
	Telegram_payment_charge_id string `json:"telegram_payment_charge_id"`
	Total_amount               int    `json:"total_amount"`
}

type ModelsSharedUser struct {
	First_name string            `json:"first_name"`
	Last_name  string            `json:"last_name"`
	Photo      []ModelsPhotoSize `json:"photo"`
	User_id    int               `json:"user_id"`
	Username   string            `json:"username"`
}

type ModelsShippingAddress struct {
	City         string `json:"city"`
	Country_code string `json:"country_code"`
	Post_code    string `json:"post_code"`
	State        string `json:"state"`
	Street_line1 string `json:"street_line1"`
	Street_line2 string `json:"street_line2"`
}

type ModelsShippingQuery struct {
	From             *ModelsUser           `json:"from,omitempty"`
	ID               string                `json:"id"`
	Invoice_payload  string                `json:"invoice_payload"`
	Shipping_address ModelsShippingAddress `json:"shipping_address"`
}

type ModelsStarAmount struct {
	Amount          int `json:"amount"`
	Nanostar_amount int `json:"nanostar_amount"`
}

type ModelsSticker struct {
	Custom_emoji_id   string              `json:"custom_emoji_id"`
	Emoji             string              `json:"emoji"`
	File_id           string              `json:"file_id"`
	File_size         int                 `json:"file_size"`
	File_unique_id    string              `json:"file_unique_id"`
	Height            int                 `json:"height"`
	Is_animated       bool                `json:"is_animated"`
	Is_video          bool                `json:"is_video"`
	Mask_position     *ModelsMaskPosition `json:"mask_position,omitempty"`
	Needs_repainting  bool                `json:"needs_repainting"`
	Premium_animation *ModelsFile         `json:"premium_animation,omitempty"`
	Set_name          string              `json:"set_name"`
	Thumbnail         *ModelsPhotoSize    `json:"thumbnail,omitempty"`
	Type              string              `json:"type"`
	Width             int                 `json:"width"`
}

type ModelsStory struct {
	Chat ModelsChat `json:"chat"`
	ID   int        `json:"id"`
}

type ModelsSuccessfulPayment struct {
	Currency                     string           `json:"currency"`
	Invoice_payload              string           `json:"invoice_payload"`
	Is_first_recurring           bool             `json:"is_first_recurring"`
	Is_recurring                 bool             `json:"is_recurring"`
	Order_info                   *ModelsOrderInfo `json:"order_info,omitempty"`
	Provider_payment_charge_id   string           `json:"provider_payment_charge_id"`
	Shipping_option_id           string           `json:"shipping_option_id"`
	Subscription_expiration_date int              `json:"subscription_expiration_date"`
	Telegram_payment_charge_id   string           `json:"telegram_payment_charge_id"`
	Total_amount                 int              `json:"total_amount"`
}

type ModelsSuggestedPostApprovalFailed struct {
	Price                  *ModelsSuggestedPostPrice `json:"price,omitempty"`
	Suggested_post_message *ModelsMessage            `json:"suggested_post_message,omitempty"`
}

type ModelsSuggestedPostApproved struct {
	Price                  *ModelsSuggestedPostPrice `json:"price,omitempty"`
	Send_date              int                       `json:"send_date"`
	Suggested_post_message *ModelsMessage            `json:"suggested_post_message,omitempty"`
}

type ModelsSuggestedPostDeclined struct {
	Comment                string         `json:"comment"`
	Suggested_post_message *ModelsMessage `json:"suggested_post_message,omitempty"`
}

type ModelsSuggestedPostInfo struct {
	Price     *ModelsSuggestedPostPrice `json:"price,omitempty"`
	Send_date int                       `json:"send_date"`
	State     string                    `json:"state"`
}

type ModelsSuggestedPostPaid struct {
	Amount                 int               `json:"amount"`
	Currency               string            `json:"currency"`
	Star_amount            *ModelsStarAmount `json:"star_amount,omitempty"`
	Suggested_post_message *ModelsMessage    `json:"suggested_post_message,omitempty"`
}

type ModelsSuggestedPostPrice struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}

type ModelsSuggestedPostRefunded struct {
	Reason                 string         `json:"reason"`
	Suggested_post_message *ModelsMessage `json:"suggested_post_message,omitempty"`
}

type ModelsSwitchInlineQueryChosenChat struct {
	Allow_bot_chats     bool   `json:"allow_bot_chats"`
	Allow_channel_chats bool   `json:"allow_channel_chats"`
	Allow_group_chats   bool   `json:"allow_group_chats"`
	Allow_user_chats    bool   `json:"allow_user_chats"`
	Query               string `json:"query"`
}

type ModelsTextQuote struct {
	Entities  []ModelsMessageEntity `json:"entities"`
	Is_manual bool                  `json:"is_manual"`
	Position  int                   `json:"position"`
	Text      string                `json:"text"`
}

type ModelsUniqueGift struct {
	Backdrop       ModelsUniqueGiftBackdrop `json:"backdrop"`
	Base_name      string                   `json:"base_name"`
	Model          ModelsUniqueGiftModel    `json:"model"`
	Name           string                   `json:"name"`
	Number         int                      `json:"number"`
	Publisher_chat *ModelsChat              `json:"publisher_chat,omitempty"`
	Symbol         ModelsUniqueGiftSymbol   `json:"symbol"`
}

type ModelsUniqueGiftBackdrop struct {
	Colors           ModelsUniqueGiftBackdropColors `json:"colors"`
	Name             string                         `json:"name"`
	Rarity_per_mille int                            `json:"rarity_per_mille"`
}

type ModelsUniqueGiftBackdropColors struct {
	Center_color int `json:"center_color"`
	Edge_color   int `json:"edge_color"`
	Symbol_color int `json:"symbol_color"`
	Text_color   int `json:"text_color"`
}

type ModelsUniqueGiftInfo struct {
	Gift                   ModelsUniqueGift `json:"gift"`
	Last_resale_star_count int              `json:"last_resale_star_count"`
	Origin                 string           `json:"origin"`
	Owned_gift_id          string           `json:"owned_gift_id"`
	Transfer_star_count    int              `json:"transfer_star_count"`
}

type ModelsUniqueGiftModel struct {
	Name             string        `json:"name"`
	Rarity_per_mille int           `json:"rarity_per_mille"`
	Sticker          ModelsSticker `json:"sticker"`
}

type ModelsUniqueGiftSymbol struct {
	Name             string        `json:"name"`
	Rarity_per_mille int           `json:"rarity_per_mille"`
	Sticker          ModelsSticker `json:"sticker"`
}

type ModelsUpdate struct {
	Business_connection       *ModelsBusinessConnection          `json:"business_connection,omitempty"`
	Business_message          *ModelsMessage                     `json:"business_message,omitempty"`
	Callback_query            *ModelsCallbackQuery               `json:"callback_query,omitempty"`
	Channel_post              *ModelsMessage                     `json:"channel_post,omitempty"`
	Chat_boost                *ModelsChatBoostUpdated            `json:"chat_boost,omitempty"`
	Chat_join_request         *ModelsChatJoinRequest             `json:"chat_join_request,omitempty"`
	Chat_member               *ModelsChatMemberUpdated           `json:"chat_member,omitempty"`
	Chosen_inline_result      *ModelsChosenInlineResult          `json:"chosen_inline_result,omitempty"`
	Deleted_business_messages *ModelsBusinessMessagesDeleted     `json:"deleted_business_messages,omitempty"`
	Edited_business_message   *ModelsMessage                     `json:"edited_business_message,omitempty"`
	Edited_channel_post       *ModelsMessage                     `json:"edited_channel_post,omitempty"`
	Edited_message            *ModelsMessage                     `json:"edited_message,omitempty"`
	Inline_query              *ModelsInlineQuery                 `json:"inline_query,omitempty"`
	Message                   *ModelsMessage                     `json:"message,omitempty"`
	Message_reaction          *ModelsMessageReactionUpdated      `json:"message_reaction,omitempty"`
	Message_reaction_count    *ModelsMessageReactionCountUpdated `json:"message_reaction_count,omitempty"`
	My_chat_member            *ModelsChatMemberUpdated           `json:"my_chat_member,omitempty"`
	Poll                      *ModelsPoll                        `json:"poll,omitempty"`
	Poll_answer               *ModelsPollAnswer                  `json:"poll_answer,omitempty"`
	Pre_checkout_query        *ModelsPreCheckoutQuery            `json:"pre_checkout_query,omitempty"`
	Purchased_paid_media      *ModelsPaidMediaPurchased          `json:"purchased_paid_media,omitempty"`
	Removed_chat_boost        *ModelsChatBoostRemoved            `json:"removed_chat_boost,omitempty"`
	Shipping_query            *ModelsShippingQuery               `json:"shipping_query,omitempty"`
	Update_id                 int                                `json:"update_id"`
}

type ModelsUser struct {
	Added_to_attachment_menu    bool   `json:"added_to_attachment_menu"`
	Can_connect_to_business     bool   `json:"can_connect_to_business"`
	Can_join_groups             bool   `json:"can_join_groups"`
	Can_read_all_group_messages bool   `json:"can_read_all_group_messages"`
	First_name                  string `json:"first_name"`
	ID                          int    `json:"id"`
	Is_bot                      bool   `json:"is_bot"`
	Is_premium                  bool   `json:"is_premium"`
	Language_code               string `json:"language_code"`
	Last_name                   string `json:"last_name"`
	Support_inline_queries      bool   `json:"support_inline_queries"`
	Username                    string `json:"username"`
}

type ModelsUsersShared struct {
	Request_id int                `json:"request_id"`
	Users      []ModelsSharedUser `json:"users"`
}

type ModelsVenue struct {
	Address           string         `json:"address"`
	Foursquare_id     string         `json:"foursquare_id"`
	Foursquare_type   string         `json:"foursquare_type"`
	Google_place_id   string         `json:"google_place_id"`
	Google_place_type string         `json:"google_place_type"`
	Location          ModelsLocation `json:"location"`
	Title             string         `json:"title"`
}

type ModelsVideo struct {
	Cover           []ModelsPhotoSize `json:"cover"`
	Duration        int               `json:"duration"`
	File_id         string            `json:"file_id"`
	File_name       string            `json:"file_name"`
	File_size       int               `json:"file_size"`
	File_unique_id  string            `json:"file_unique_id"`
	Height          int               `json:"height"`
	Mime_type       string            `json:"mime_type"`
	Start_timestamp int               `json:"start_timestamp"`
	Thumbnail       *ModelsPhotoSize  `json:"thumbnail,omitempty"`
	Width           int               `json:"width"`
}

type ModelsVideoNote struct {
	Duration       int              `json:"duration"`
	File_id        string           `json:"file_id"`
	File_size      int              `json:"file_size"`
	File_unique_id string           `json:"file_unique_id"`
	Length         int              `json:"length"`
	Thumbnail      *ModelsPhotoSize `json:"thumbnail,omitempty"`
}

type ModelsVoice struct {
	Duration       int    `json:"duration"`
	File_id        string `json:"file_id"`
	File_size      int    `json:"file_size"`
	File_unique_id string `json:"file_unique_id"`
	Mime_type      string `json:"mime_type"`
}

type ModelsVoiceChatEnded struct {
	Duration int `json:"duration"`
}

type ModelsVoiceChatParticipantsInvited struct {
	Users []ModelsUser `json:"users"`
}

type ModelsVoiceChatScheduled struct {
	Start_date int `json:"start_date"`
}

type ModelsVoiceChatStarted struct {
	Duration int `json:"duration"`
}

type ModelsWebAppData struct {
	Button_text string `json:"button_text"`
	Data        string `json:"data"`
}

type ModelsWebAppInfo struct {
	Url string `json:"url"`
}

type ModelsWriteAccessAllowed struct {
	From_attachment_menu bool   `json:"from_attachment_menu"`
	From_request         bool   `json:"from_request"`
	Web_app_name         string `json:"web_app_name"`
}

type ServiceResponse struct {
	ID   string `json:"ID"`
	Name string `json:"Name"`
}

type svcLogs struct {
	client *rpcClient
}

func newClientLogs(client *rpcClient) *svcLogs {
	return &svcLogs{
		client: client,
	}
}

func (c *svcLogs) AddTelegramLog(ctx context.Context, update ModelsUpdate) (err error) {
	_req := struct {
		Update ModelsUpdate
	}{
		Update: update,
	}

	err = c.client.call(ctx, "logs.AddTelegramLog", _req, nil)

	return
}

func (c *svcLogs) Get(ctx context.Context) (res []ServiceResponse, err error) {
	_req := struct {
	}{}

	err = c.client.call(ctx, "logs.Get", _req, &res)

	return
}

func (c *svcLogs) GetLogsByServiceID(ctx context.Context, serviceID int) (res LogsService, err error) {
	_req := struct {
		ServiceID int
	}{
		ServiceID: serviceID,
	}

	err = c.client.call(ctx, "logs.GetLogsByServiceID", _req, &res)

	return
}

type rpcClient struct {
	endpoint string
	cl       *http.Client

	requestID uint64
}

func newRPCClient(endpoint string, httpClient *http.Client) *rpcClient {
	return &rpcClient{
		endpoint: endpoint,
		cl:       httpClient,
	}
}

func (rc *rpcClient) call(ctx context.Context, methodName string, request, result interface{}) error {
	// encode params
	bts, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("encode params: %w", err)
	}

	requestID := atomic.AddUint64(&rc.requestID, 1)
	requestIDBts := json.RawMessage(strconv.Itoa(int(requestID)))

	req := zenrpc.Request{
		Version: zenrpc.Version,
		ID:      &requestIDBts,
		Method:  methodName,
		Params:  bts,
	}

	ctx = appkit.NewCallerNameContext(ctx, name)

	res, err := rc.Exec(ctx, req)
	if err != nil {
		return err
	}

	if res == nil {
		return nil
	}

	if res.Error != nil {
		return res.Error
	}

	if res.Result == nil {
		return nil
	}

	if result == nil {
		return nil
	}

	return json.Unmarshal(*res.Result, result)
}

// Exec makes http request to jsonrpc endpoint and returns json rpc response.
func (rc *rpcClient) Exec(ctx context.Context, rpcReq zenrpc.Request) (*zenrpc.Response, error) {
	if appkit.NotificationFromContext(ctx) {
		rpcReq.ID = nil
	}

	c, err := json.Marshal(rpcReq)
	if err != nil {
		return nil, fmt.Errorf("json marshal call failed: %w", err)
	}

	buf := bytes.NewReader(c)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, rc.endpoint, buf)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	req.Header.Add("Content-Type", "application/json")
	appkit.SetXRequestIDFromCtx(ctx, req)

	// Do request
	resp, err := rc.cl.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return nil, fmt.Errorf("make request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad response (%d)", resp.StatusCode)
	}

	var zresp zenrpc.Response
	if rpcReq.ID == nil {
		return &zresp, nil
	}

	bb, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("response body (%s) read failed: %w", bb, err)
	}

	if err = json.Unmarshal(bb, &zresp); err != nil {
		return nil, fmt.Errorf("json decode failed (%s): %w", bb, err)
	}

	return &zresp, nil
}
