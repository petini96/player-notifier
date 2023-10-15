package player

type Player struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	Number        string   `json:"number"`
	Nick          string   `json:"nick"`
	Serie         string   `json:"serie"`
	PresenceLevel int      `json:"presence_level"`
	Position      []string `json:"position"`
	IsOlder       bool     `json:"is_older"`
}
