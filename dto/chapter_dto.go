package dto

import "golern/model"

type com struct {
}

type ChapterDto struct {
	ID       int           `json:"id"`
	Name     string        `json:"name"`
	Pid      int           `json:"Pid"`
	Did      int           `json:"did"`
	Children []*ChapterDto `json:"children"`
}

func ToChapterDto(chapters []model.Chapter) map[int]*ChapterDto {
	chapterRes := make(map[int]*ChapterDto)
	for _, n := range chapters {
		chapterRes[int(n.ID)] = &ChapterDto{
			ID:   int(n.ID),
			Name: n.Name,
			Pid:  n.Pid,
			Did:  n.Did,
		}
	}
	return chapterRes
}
