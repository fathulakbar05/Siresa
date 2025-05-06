package models

import "time"

type Reservation struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	RoomID    uint      `json:"room_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`

	User User `json:"user" gorm:"foreignKey:UserID"`
	Room Room `json:"room" gorm:"foreignKey:RoomID"`
}
