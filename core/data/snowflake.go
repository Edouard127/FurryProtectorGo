package data

import "strconv"

type Snowflake uint64

func (s Snowflake) String() string {
	return strconv.FormatUint(uint64(s), 10)
}
