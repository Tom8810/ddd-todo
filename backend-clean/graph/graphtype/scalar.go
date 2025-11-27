package graphtype

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalDate(t time.Time) graphql.Marshaler {
	if t.IsZero() {
		return graphql.Null
	}

	return graphql.WriterFunc(func(w io.Writer) {
		_, err := io.WriteString(w, strconv.Quote(t.Format(time.DateOnly)))
		if err != nil {
			fmt.Printf("error: %v", err)
		}
	})
}

func UnmarshalDate(v interface{}) (time.Time, error) {
	str, ok := v.(string)
	if !ok {
		return time.Time{}, fmt.Errorf("dates must be strings")
	}

	t, err := time.Parse(time.DateOnly, str)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date format: %w", err)
	}

	return t, nil
}

func MarshalDateTime(t time.Time) graphql.Marshaler {
	if t.IsZero() {
		return graphql.Null
	}

	return graphql.WriterFunc(func(w io.Writer) {
		_, err := io.WriteString(w, strconv.Quote(t.Format(time.RFC3339)))
		if err != nil {
			fmt.Printf("error: %v", err)
		}
	})
}

func UnmarshalDateTime(v interface{}) (time.Time, error) {
	str, ok := v.(string)
	if !ok {
		return time.Time{}, fmt.Errorf("datetime must be strings")
	}

	// Try RFC3339 format first (ISO 8601)
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		// Fallback to other common datetime formats
		formats := []string{
			"2006-01-02T15:04:05Z07:00",
			"2006-01-02 15:04:05",
			"2006-01-02T15:04:05",
		}
		
		for _, format := range formats {
			if t, err = time.Parse(format, str); err == nil {
				break
			}
		}
		
		if err != nil {
			return time.Time{}, fmt.Errorf("invalid datetime format: %w", err)
		}
	}

	return t, nil
}
