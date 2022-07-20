package randomvalues

import (
	"fmt"
	faker "github.com/bxcodec/faker/v3"
	"github.com/jxskiss/base62"
	"math/rand"
	"time"
)

const n = 1000000

func init() {
	rand.Seed(time.Now().Unix())
}

func RandomUrlId() string {
	dst := base62.FormatUint(rand.Uint64())
	return string(dst)
}

func RandomOriginalUrl() string {
	return faker.URL()
}

func RandomUserId() string {
	return faker.Email()
}

func InvalidUrlId() string {
	return fmt.Sprintf("%s%s", faker.Word(), "$·!")
}

func InvalidOriginalUrl() string {
	return faker.Word()
}

func InvalidUserId() string {
	return faker.Word()
}
