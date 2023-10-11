package cfg

type SpotifyScope struct {
	Value  string   `json:"value"`
	Name   string   `json:"name"`
	Scopes []string `json:"scopes"`
}

const (
	SPOTIFY_API_URL     = "https://accounts.spotify.com/api/token"
	AUTH_SCOPES_MAPPING = `
	[
		{
			"value": "default",
			"name": "Read & modify playback.",
			"scopes": [
				"user-read-playback-state",
				"user-modify-playback-state",
				"user-read-recently-played"
			]
		},
		{
			"value": "user-read",
			"name": "Read user library, followed artists/users, and top artists/tracks.",
			"scopes": [
				"user-top-read",
				"user-library-read",
				"user-follow-read"
			]
		},
		{
			"value": "user-modify",
			"name": "Modify user library and followed artists/users.",
			"scopes": [
				"user-library-modify",
				"user-follow-modify"
			]
		}
	]
	`
)
