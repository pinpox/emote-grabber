package main

import (
	"time"
)

type BTTVResponse struct {
	ID           string             `json:"id"`
	Name         string             `json:"name"`
	DisplayName  string             `json:"displayName"`
	ProviderID   string             `json:"providerId"`
	Bots         []interface{}      `json:"bots"`
	Emotes       []BTTVEmote       `json:"channelEmotes"`
	SharedEmotes []BTTVSharedEmote `json:"sharedEmotes"`
}
type BTTVEmote struct {
	ID             string    `json:"id"`
	Code           string    `json:"code"`
	ImageType      string    `json:"imageType"`
	Animated       bool      `json:"animated"`
	UserID         string    `json:"userId"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Global         bool      `json:"global"`
	Live           bool      `json:"live"`
	Sharing        bool      `json:"sharing"`
	ApprovalStatus string    `json:"approvalStatus"`
}
type BTTVUser struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	ProviderID  string `json:"providerId"`
}
type BTTVSharedEmote struct {
	ID             string    `json:"id"`
	Code           string    `json:"code"`
	ImageType      string    `json:"imageType"`
	Animated       bool      `json:"animated"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Global         bool      `json:"global"`
	Live           bool      `json:"live"`
	Sharing        bool      `json:"sharing"`
	ApprovalStatus string    `json:"approvalStatus"`
	User           BTTVUser  `json:"user"`
	Width          int       `json:"width,omitempty"`
	Height         int       `json:"height,omitempty"`
}
