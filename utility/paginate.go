package utility

import (
	"fmt"
	"log"
)

func GetPaginateURL(paths []string, page *int, limit *int, numberRecords int) (string, string, int) {
	var tempPage int = *page
	var tempLimit int = *limit

	if tempPage < 0 {
		tempPage = 1
	}

	if tempLimit < 0 {
		tempLimit = 10
	} else if tempLimit > 25 {
		tempLimit = 25
	}

	totalPages := 0
	if totalPages = numberRecords / tempLimit; numberRecords%tempLimit != 0 {
		totalPages += 1
	}

	if tempPage > totalPages {
		tempPage = totalPages
	}

	nextPage := ""
	prevPage := ""
	if len(paths) == 1 {
		// api/:collections?page=[number]&limit=[number]
		nextPage = fmt.Sprintf("api/%s?page=%d&limit=%d", paths[0], tempPage+1, tempLimit)
		prevPage = fmt.Sprintf("api/%s?page=%d&limit=%d", paths[0], tempPage-1, tempLimit)
	} else if len(paths) == 3 {
		// api/:parent/:parent-id/:child?page=[number]&limit=[number]
		nextPage = fmt.Sprintf("api/%s/%s/%s?page=%d&limit=%d", paths[0], paths[1], paths[2], tempPage+1, tempLimit)
		prevPage = fmt.Sprintf("api/%s/%s/%s?page=%d&limit=%d", paths[0], paths[1], paths[2], tempPage-1, tempLimit)
	}

	if (tempPage + 1) > totalPages {
		nextPage = ""
		tempPage = totalPages
	} else if (tempPage - 1) < 1 {
		nextPage = ""
		tempPage = 1
	}

	if tempPage >= 1 && tempLimit >= numberRecords {
		tempPage = 1
		tempLimit = numberRecords
		prevPage = ""
	}

	*page = tempPage
	*limit = tempLimit
	log.Printf("page: %d\n", *page)
	log.Printf("tempPage: %d\n", tempPage)
	return nextPage, prevPage, totalPages
}
