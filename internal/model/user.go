package model

import (
	userv1 "github.com/Shchirov/road-map-api/gen/go/user"
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	FirstName  string    `gorm:"first_name,notnull"`
	SecondName string    `gorm:"second_name,notnull"`
	Surname    string    `gorm:"surname"`
	IsActive   bool      `gorm:"is_active,notnull,default:true"`
	Avatar     string    `gorm:"avatar"`
	PostId     int32     `gorm:"post_id,notnull,default:'now()'"`
	StreamId   int32     `gorm:"stream_id,notnull,default:'now()'"`
	Stream     Stream    `gorm:"foreignKey:StreamId;references:Id"`
	Post       Post      `gorm:"foreignKey:PostId;references:Id"`
	CreatedDt  time.Time `gorm:"created_dt,notnull,default:'now()'"`
}

func (user *User) ToResponse() *userv1.UserResponse {
	post := user.Post
	stream := user.Stream
	return &userv1.UserResponse{
		Id:         user.Id.String(),
		SecondName: user.SecondName,
		FirstName:  user.FirstName,
		Surname:    user.Surname,
		Avatar:     user.Avatar,
		Post: &userv1.PostResponse{
			Id:          post.Id,
			Code:        post.Code,
			Description: post.Description,
		},
		Stream: &userv1.StreamResponse{
			Id:          stream.Id,
			Code:        stream.Code,
			Description: stream.Description,
		},
	}
}

func NewUser(id uuid.UUID, firstName string, secondName string, surname string, isActive bool, avatar string, postId int32, streamId int32) *User {
	return &User{Id: id, FirstName: firstName, SecondName: secondName, Surname: surname, IsActive: isActive, Avatar: avatar, PostId: postId, StreamId: streamId}
}
