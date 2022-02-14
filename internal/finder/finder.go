package finder

import (
	"fmt"
	"sort"

	"github.com/HunterGooD/indexter/internal/models"
)

type Finder struct {
	File models.FileIndexer
}

func (f *Finder) search(wordSearch string) {
	index := sort.Search(len(f.File.File), func(i int) bool {
		return f.File.File[i].Value == wordSearch
	})
	fmt.Print(index)
}
