package utils

import (
	"net/url"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
)

func GetPagination(query url.Values) (uint64, uint64) {
	limitStr := query.Get("limit")
	pageStr := query.Get("page")

	limit, err := strconv.ParseUint(limitStr, 10, 64)
	if err != nil {
		limit = 32
	}

	page, err := strconv.ParseUint(pageStr, 10, 64)
	if err != nil {
		page = 0
	}

	return limit, page
}

func GetFilter(query url.Values, parsers map[string]func(val string) (string, any)) bson.M {
	filter := bson.M{}
	for field, fun := range parsers {
		val := query.Get(field)
		if val != "" {
			f, v := fun(val)
			filter[f] = v
		}
	}
	return filter
}
