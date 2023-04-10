package domain

import "errors"

type Movie struct{
	Id int `json:"id"`
	Title string `json:"title"`
	Year int `json:"year"`
	Genre string `json:"genre"`
	Rating float64 `json:"rating"`
 
}

func (mv *Movie) Validate() (err error)  {
	if mv.Title == "" || mv.Genre == "" {
		err = errors.New("title and genre are required")
		return
	}

	if mv.Year < 1888 {
		err = errors.New("year must bwe greater than 1888")
		return
	}
	if mv.Rating < 0 || mv.Rating > 10 {
		err = errors.New("rating must be between 0 and 10")
		return
	}
	return
}