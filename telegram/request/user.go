package request

type CreateUsersRequest struct {
	UserChatID   int64   `json:"user_chat_id"`
	UserName     *string `json:"user_name"`
	UserNickName string  `json:"user_nick_name"`
}
