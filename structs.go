package plashikigo

// ! Main types

// Episode is struct for anime episode
type Episode struct {
	Authors []Author `json:"authors,omitempty"`
	Players []string `json:"players,omitempty"`
}

// Author provides authors of anime episode
type Author struct {
	Kind         string        `json:"kind,omitempty"`
	Lang         string        `json:"lang,omitempty"`
	MetaTag      string        `json:"meta_tag,omitempty"`
	Name         string        `json:"name,omitempty"`
	OptName      string        `json:"opt_name,omitempty"`
	Translations []Translation `json:"translations,omitempty"`
}

// Translation provides translations by author
type Translation struct {
	ID       int    `json:"id,omitempty"`
	Quality  string `json:"quality,omitempty"`
	Author   string `json:"author,omitempty"`
	Uploader string `json:"uploader,omitempty"`
	Name     string `json:"name,omitempty"`
	URL      string `json:"url,omitempty"`
}

// ! /anime/v2/:id query

// AnimeIDResponse provides response for /anime/v2/:id request
type AnimeIDResponse struct {
	OK        bool            `json:"ok,omitempty"`
	Reason    string          `json:"reason,omitempty"`
	Result    map[int]Episode `json:"result,omitempty"`
	ServeTime float32         `json:"serve_time,omitempty"`
}

// ! /anime/v2/:id/episodes/:episode query

// AnimeEpisodeResponse provides response for /anime/v2/:id/episodes/:episode request
type AnimeEpisodeResponse struct {
	OK        bool            `json:"ok,omitempty"`
	Reason    string          `json:"reason,omitempty"`
	Result    map[int]Episode `json:"result,omitempty"`
	ServeTime float32         `json:"serve_time,omitempty"`
}

// ! /anime/v2/:id/episodes query

// AnimeEpisodesResponse provides response for /anime/v2/:id/episodes request
type AnimeEpisodesResponse struct {
	OK        bool    `json:"ok,omitempty"`
	Reason    string  `json:"reason,omitempty"`
	Result    []int   `json:"result,omitempty"`
	ServeTime float32 `json:"serve_time,omitempty"`
}

// ! /anime/query query

// AnimeQueryResponse provides response for /anime/query request
type AnimeQueryResponse struct {
	OK        bool          `json:"ok,omitempty"`
	Reason    string        `json:"reason,omitempty"`
	Result    []Translation `json:"result,omitempty"`
	ServeTime float32       `json:"serve_time,omitempty"`
}
