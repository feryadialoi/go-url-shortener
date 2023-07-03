package model

type URLShortenEntity struct {
	ID        int64
	ShortPath string
	RealURL   string
}

type CreateURLShortenRequest struct {
	ShortPath string `json:"shortPath"`
	RealURL   string `json:"realURL"`
}

type URLShortenResponse struct {
	ID        int64  `json:"id"`
	ShortPath string `json:"shortPath"`
	RealURL   string `json:"realURL"`
}
