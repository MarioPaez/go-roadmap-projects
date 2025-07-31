package model

type Event struct {
	Id        string     `json:"id"`
	Type      string     `json:"type"`
	Actor     Actor      `json:"actor"`
	Repo      Repository `json:"repo"`
	Payload   Payload    `json:"payload"`
	Public    bool       `json:"public"`
	CreatedAt string     `json:"created_at"`
}

type Actor struct {
	ID           int    `json:"id"`
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
	GravatarId   string `json:"gravatar_id"`
	Url          string `json:"url"`
	AvatarUrl    string `json:"avatar_url"`
}

type Repository struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Payload struct {
	RepositoryID int      `json:"repository_id"`
	Action       string   `json:"action"`
	PushID       int64    `json:"push_id"`
	Size         int      `json:"size"`
	DistinctSize int      `json:"distinct_size"`
	Ref          string   `json:"ref"`
	Head         string   `json:"head"`
	Before       string   `json:"before"`
	Commits      []Commit `json:"commits"`
}

type Commit struct {
	Sha      string `json:"sha"`
	Author   Author `json:"author"`
	Message  string `json:"message"`
	Distinct bool   `json:"distinct"`
	Url      string `json:"url"`
}

type Author struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
