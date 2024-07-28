package entity

import "time"

type (
	DeleteUserReq struct {
		UserID       string `json:"user_id" bson:"user_id"`
		IsHardDelete bool   `json:"is_hard_delete" bson:"is_hard_delete"`
	}

	GetAllUserReq struct {
		Field   string `json:"field" bson:"field"`
		Value   string `json:"value" bson:"value"`
		Page    string `json:"page" bson:"page"`
		Limit   string `json:"limit" bson:"limit"`
		StartAt string `json:"start_at" bson:"start_at"`
		EndAt   string `json:"end_at" bson:"end_at"`
	}

	GetAllUserRes struct {
		Users []User `json:"users" bson:"users"`
		Count int32  `json:"count" bson:"count"`
	}

	GetUserReq struct {
		Field string `json:"field" bson:"field"`
		Value string `json:"value" bson:"value"`
	}

	UpdateUserReq struct {
		FirstName string `json:"first_name" bson:"first_name"`
		LastName  string `json:"last_name" bson:"last_name"`
	}

	UpdatePasswordReq struct {
		Password    string `json:"password" bson:"password"`
		NewPassword string `json:"new_password" bson:"new_password"`
	}

	UpdateEmailReq struct {
		Password string `json:"password" bson:"password"`
		NewEmail string `json:"new_email" bson:"new_email"`
	}

	LoginReq struct {
		Email    string `json:"email" bson:"email"`
		Password string `json:"password" bson:"password"`
	}

	LoginRes struct {
		Token string `json:"token" bson:"token"`
	}

	StatusUser struct {
		Successfully bool `json:"successfully" bson:"successfully"`
	}

	Profile struct {
		FirstName string    `json:"first_name" bson:"first_name"`
		CreatedAt time.Time `json:"created_at" bson:"created_at"`
		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`
		Address   string    `json:"address" bson:"address"`
	}

	User struct {
		FirstName string  `json:"first_name" bson:"first_name"`
		LastName  string  `json:"last_name" bson:"last_name"`
		Email     string  `json:"email" bson:"email"`
		Password  string  `json:"password" bson:"password"`
		ID        string  `json:"id" bson:"id"`
		Profile   Profile `json:"profile" bson:"profile"`
	}

	CreateUserReq struct {
		FirstName string `json:"first_name" bson:"first_name"`
		LastName  string `json:"last_name" bson:"last_name"`
		Email     string `json:"email" bson:"email"`
		Password  string `json:"password" bson:"password"`
		Address   string `json:"address" bson:"address"`
	}
)
