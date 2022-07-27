package domain

import (
	"errors"
	"fmt"
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/log"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo/randomvalues"
	"github.com/jxskiss/base62"
	"sync"
)

type RangeKey struct {
	start  uint64
	end    uint64
	offset uint64
	mutex  sync.Mutex
}

func NewRangeKey(start, end uint64) *RangeKey {
	return &RangeKey{
		start:  start,
		end:    end,
		offset: start,
		mutex:  sync.Mutex{},
	}
}

func (rk *RangeKey) GetKey() (string, error) {
	auxOffset, err := rk.newOffset()
	if err != nil {
		return "", err
	}
	dst := base62.FormatUint(auxOffset)
	str := string(dst)
	key := fmt.Sprintf("%s%s", str, rk.subString())
	return key, nil
}
func (rk *RangeKey) newOffset() (uint64, error) {
	rk.mutex.Lock()
	defer rk.mutex.Unlock()
	if rk.offset > rk.end {
		return uint64(0), errors.New("key out of range")
	}
	auxOffset := rk.offset
	log.Debug("assigned key: %d", auxOffset)
	rk.offset = rk.offset + 1

	return auxOffset, nil
}

func (rk *RangeKey) subString() string {
	str := randomvalues.RandomUrlId()
	if len(str) >= 2 {
		str = str[0:1]
	}
	return str
}
