package zookeeper

import (
	"errors"
	"github.com/creepzed/url-shortener-service/app/key-generation-service/domain"
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/log"
	"github.com/samuel/go-zookeeper/zk"
	"strconv"
	"strings"
	"time"
)

const (
	keyQuantity = 10000
)

type zookeeperRepository struct {
	zkNodes []string
	path    string
}

func NewZookeeperRepository(path string, zkNodes ...string) *zookeeperRepository {
	return &zookeeperRepository{
		zkNodes: zkNodes,
		path:    path,
	}
}

func (zkr zookeeperRepository) GetRange() (*domain.RangeKey, error) {
	conn, _, err := zk.Connect(zkr.zkNodes, time.Second)
	conn.SetLogger(log.Logger())
	if err != nil {
		return &domain.RangeKey{}, err
	}
	defer conn.Close()

	data := []byte("")
	acl := zk.WorldACL(0)

	if b, _, _ := conn.Exists("/kgs"); !b {
		zNodeAux, err := conn.Create("/kgs", data, 0, acl)
		if err != nil {
			return &domain.RangeKey{}, err
		}
		log.Debug("create zNode father: %s", zNodeAux)
	}

	zNode, err := conn.CreateProtectedEphemeralSequential("/kgs", data, acl)
	if err != nil {
		return &domain.RangeKey{}, err
	}
	log.Debug("create zNode protected ephemeral sequential: %s", zNode)

	startRange, endRange, err := zkr.getRange(zNode, "kgs")
	if err != nil {
		return &domain.RangeKey{}, err
	}
	rangeKey := domain.NewRangeKey(startRange, endRange)

	return rangeKey, nil
}

func (zkr zookeeperRepository) getRange(zNode string, key string) (uint64, uint64, error) {

	split := strings.Split(zNode, "-")

	if len(split) != 2 {
		return uint64(0), uint64(0), errors.New("invalid kgs range")
	}
	rangeData := split[1]
	str := strings.ReplaceAll(rangeData, key, "")
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, 0, errors.New("invalid kgs range")
	}

	start := uint64(num*keyQuantity) + 1
	end := uint64((num+1)*keyQuantity) - 1

	return start, end, nil
}
