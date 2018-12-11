package crawler

import (
	"github.com/MwlLj/go-crawler/common"
)

type CCrawler struct {
}

func New() common.ICrawler {
	craw := CCrawler{}
	return &craw
}
