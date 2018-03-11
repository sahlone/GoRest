package model

type SearchCriteria struct {
	SearchByID         string      `json:"searchById,omitempty"`
	SearchByName       string      `json:"searchByName,omitempty"`
	SearchByPrep       Preparation `json:"searchByPrep,omitempty"`
	SearchByDifficulty Difficulty  `json:"searchByDifficulty,omitempty"`
	SearchByIsVeg      string      `json:"isVeg,,omitempty"`
	SearchByArchive    string      `json:"isArchive,omitempty"`
	Start              int         `json:"start,omitempty"`
	Limit              int         `json:"limit,omitempty"`
}
