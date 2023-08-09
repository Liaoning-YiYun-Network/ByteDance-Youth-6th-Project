package util

import (
	"github.com/google/uuid"
	"golang.org/x/exp/rand"
	"strings"
	"time"
)

// UUID 生成UUID
func UUID() (string, error) {
	rand.Seed(uint64(time.Now().Unix()))
	u, err := uuid.NewRandom()
	return u.String(), err
}

// UUIDWithoutHyphen 生成不带横杠的UUID
func UUIDWithoutHyphen() (string, error) {
	rand.Seed(uint64(time.Now().Unix()))
	u, err := uuid.NewRandom()
	str := u.String()
	strings.ReplaceAll(str, "-", "")
	return str, err
}
