package helper

func GetNextPage(nPage int, perPage int, nProducts int) int {
	if (nPage+1)*perPage >= nProducts {
		return -1
	}
	return nPage + 1
}
