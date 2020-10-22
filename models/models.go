package models

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"io"
	"strconv"
	"time"
)

// Объявим базовый тип int для ID
func MarshalTimestamp(t time.Time) graphql.Marshaler {
	timestamp := t.Unix() * 1000

	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.FormatInt(timestamp, 10))
	})
}

func UnmarshalTimestamp(v interface{}) (time.Time, error) {
	if tmpStr, ok := v.(int); ok {
		return time.Unix(int64(tmpStr), 0), nil
	}
	return time.Time{}, fmt.Errorf("Timestamp error ")
}
