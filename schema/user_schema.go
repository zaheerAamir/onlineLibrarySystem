package schema

type UserDto struct {
	LAST_NAME  string `json:"last_name"`
	FIRST_NAME string `json:"first_name"`
	EMAIL      string `json:"email"`
	PASSWORD   string `json:"password"`
}

type UserSchema struct {
	ID            int    `json:"user_id"`
	LAST_NAME     string `json:"last_name"`
	FIRST_NAME    string `json:"first_name"`
	EMAIL         string `json:"email"`
	HASH_PASSWORD string `json:"hash_password"`
	SALT          string `json:"salt"`
}

type UserLoginDto struct {
	EMAIL    string `json:"email"`
	PASSWORD string `json:"password"`
}

type UserSucces struct {
	CODE          int    `json:"status_code"`
	TEXT          string `json:"text"`
	ACCES_TOKEN   string `json:"access_token"`
	REFRESH_TOKEN string `json:"refresh_toke"`
}

type RefreshTokenDTO struct {
	REFRESH_TOKEN string `json:"token"`
}

type AccessTokenSchema struct {
	AccessToken string `json:"access_token"`
}

type Logout struct {
	EMAIL string `json:"email"`
}
