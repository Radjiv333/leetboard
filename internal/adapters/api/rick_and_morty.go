package api

type RickAndMortyResponse struct {
	Info    RickAndMortyInfo         `json:"info"`
	Results []RickAndMortyCharacters `json:"results"`
}

type RickAndMortyCharacters struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image"`
}

type RickAndMortyInfo struct {
	Count   int    `json:"count"`
	Pages   int    `json:"pages"`
	NextURL *string `json:"next"`
}
