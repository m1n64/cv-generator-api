package utils

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func ParseStartAndEndDate(startDate string, endDate *string) (time.Time, *time.Time, error) {
	startDateParsed, err := time.Parse(time.DateOnly, startDate)
	if err != nil {
		return time.Time{}, nil, status.Error(codes.InvalidArgument, "invalid start_date format")
	}

	var endDateParsed *time.Time
	if endDate != nil && *endDate != "" {
		parsedEndDate, err := time.Parse(time.DateOnly, *endDate)
		if err != nil {
			return time.Time{}, nil, status.Error(codes.InvalidArgument, "invalid end_date format")
		}
		endDateParsed = &parsedEndDate
	}

	return startDateParsed, endDateParsed, nil
}

func PtrToTimeString(t *time.Time) *string {
	if t == nil {
		return nil
	}
	str := t.String()
	return &str
}

func RandStringBytesRmndr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
