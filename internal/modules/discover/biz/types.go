package biz

import (
	cms "thmanyah/internal/modules/cms/biz"
)

type SearchResults struct {
	Categories []*cms.Category `json:"categories"`
	Programs   []*cms.Program  `json:"programs"`
	Episodes   []*cms.Episode  `json:"episodes"`
}
