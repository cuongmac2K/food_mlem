package acount

type LoginBody struct {
	UserName string `json:"user_name" bson:"user_name"`
	PassWord string `json:"pass_word" bson:"pass_word"`
}
type RegisterBody struct {
	Role     string `bson:"role" json:"role"`
	UserName string `json:"user_name" bson:"user_name"`
	PassWord string `json:"pass_word" bson:"pass_word"`
}
