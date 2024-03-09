package main

type STVResponse struct {
	ID         string        `json:"id"`
	Name       string        `json:"name"`
	Flags      int           `json:"flags"`
	Tags       []interface{} `json:"tags"`
	Immutable  bool          `json:"immutable"`
	Privileged bool          `json:"privileged"`
	Emotes     []STVEmote   `json:"emotes"`
	EmoteCount int           `json:"emote_count"`
	Capacity   int           `json:"capacity"`
	Owner      STVOwner      `json:"owner"`
}
type STVStyle struct {
}
type STVOwner struct {
	ID          string   `json:"id"`
	Username    string   `json:"username"`
	DisplayName string   `json:"display_name"`
	AvatarURL   string   `json:"avatar_url"`
	Style       STVStyle `json:"style"`
	Roles       []string `json:"roles"`
}
type STVFile struct {
	Name       string `json:"name"`
	StaticName string `json:"static_name"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	FrameCount int    `json:"frame_count"`
	Size       int    `json:"size"`
	Format     string `json:"format"`
}
type STVHost struct {
	URL   string     `json:"url"`
	Files []STVFile `json:"files"`
}
type STVData struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Flags     int      `json:"flags"`
	Lifecycle int      `json:"lifecycle"`
	State     []string `json:"state"`
	Listed    bool     `json:"listed"`
	Animated  bool     `json:"animated"`
	Owner     STVOwner `json:"owner"`
	Host      STVHost  `json:"host"`
}
type STVEmote struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Flags     int         `json:"flags"`
	Timestamp int64       `json:"timestamp"`
	ActorID   interface{} `json:"actor_id"`
	Data      STVData     `json:"data,omitempty"`
}
