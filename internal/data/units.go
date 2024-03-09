package data

import (
	"fmt"
	"strconv"
	"strings"
	"errors"
)

var ErrInvalidRuntimeFormat = errors.New("invalid format")

type Walking int64
type Hydrate int32
type Sleep int32

func (r Walking) MarshalJSON() ([]byte, error){
	jsonValue := fmt.Sprintf("%d steps", r)
	quotedJSONValue := strconv.Quote(jsonValue)
	return []byte(quotedJSONValue), nil
}

func (r *Walking) UnmarshalJSON(jsonValue []byte) error{
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	parts := strings.Split(unquotedJSONValue, " ")
	if len(parts) != 2 || parts[1] != "steps" {
		return ErrInvalidRuntimeFormat
	}

	i, err:= strconv.ParseInt(parts[0], 10, 32)
	if err != nil{
		return ErrInvalidRuntimeFormat
	}
	*r = Walking(i)
	return nil
}

func (r Hydrate) MarshalJSON() ([]byte, error){
	jsonValue := fmt.Sprintf("%d gallons", r)
	quotedJSONValue := strconv.Quote(jsonValue)
	return []byte(quotedJSONValue), nil
}

func (r *Hydrate) UnmarshalJSON(jsonValue []byte) error{
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	parts := strings.Split(unquotedJSONValue, " ")
	if len(parts) != 2 || parts[1] != "gallons" {
		return ErrInvalidRuntimeFormat
	}

	i, err:= strconv.ParseInt(parts[0], 10, 32)
	if err != nil{
		return ErrInvalidRuntimeFormat
	}
	*r = Hydrate(i)
	return nil
}

func (r Sleep) MarshalJSON() ([]byte, error){
	jsonValue := fmt.Sprintf("%d hours", r)
	quotedJSONValue := strconv.Quote(jsonValue)
	return []byte(quotedJSONValue), nil
}

func (r *Sleep) UnmarshalJSON(jsonValue []byte) error{
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	parts := strings.Split(unquotedJSONValue, " ")
	if len(parts) != 2 || parts[1] != "hours" {
		return ErrInvalidRuntimeFormat
	}

	i, err:= strconv.ParseInt(parts[0], 10, 32)
	if err != nil{
		return ErrInvalidRuntimeFormat
	}
	*r = Sleep(i)
	return nil
}