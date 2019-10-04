package types

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// ID さくらのクラウド上のリソースのIDを示す
//
// APIリクエスト/レスポンスに文字列/数値が混在するためここで吸収する
type ID int64

// UnmarshalJSON implements unmarshal from both of JSON number and JSON string
func (i *ID) UnmarshalJSON(b []byte) error {
	s := string(b)
	if s == "" || s == "null" {
		return nil
	}
	var n json.Number
	if err := json.Unmarshal(b, &n); err != nil {
		return err
	}
	id, err := n.Int64()
	if err != nil {
		return err
	}
	*i = ID(id)
	return nil
}

// IsEmpty return true if ID is empty or zero value
func (i ID) IsEmpty() bool {
	return i.Int64() == 0
}

// String returns the literal text of the number.
func (i ID) String() string {
	if i.IsEmpty() {
		return ""
	}
	return fmt.Sprintf("%d", i)
}

// Int64 returns the number as an int64.
func (i ID) Int64() int64 {
	return int64(i)
}

// Int64ID creates new ID from int64
func Int64ID(id int64) ID {
	return ID(id)
}

// StringID creates new ID from string
func StringID(id string) ID {
	intID, _ := strconv.ParseInt(id, 10, 64)
	return ID(intID)
}
