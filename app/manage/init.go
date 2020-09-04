package manage

import (
	"elastic-transfer/app/mq"
	"elastic-transfer/app/schema"
	"elastic-transfer/app/types"
	"errors"
	"github.com/elastic/go-elasticsearch/v8"
)

type ElasticManager struct {
	client *elasticsearch.Client
	mq     *mq.MessageQueue
	pipes  map[string]*types.PipeOption
	schema *schema.Schema
}

func NewElasticManager(config elasticsearch.Config, mq *mq.MessageQueue) (manager *ElasticManager, err error) {
	manager = new(ElasticManager)
	manager.client, err = elasticsearch.NewClient(config)
	manager.mq = mq
	if err != nil {
		return
	}
	manager.pipes = make(map[string]*types.PipeOption)
	manager.schema = schema.New()
	return
}

func (c *ElasticManager) empty(identity string) error {
	if c.pipes[identity] == nil {
		return errors.New("this identity does not exists")
	}
	return nil
}

func (c *ElasticManager) GetIdentityCollection() []string {
	var keys []string
	for key := range c.pipes {
		keys = append(keys, key)
	}
	return keys
}

func (c *ElasticManager) GetOption(identity string) (option *types.PipeOption, err error) {
	if err = c.empty(identity); err != nil {
		return
	}
	option = c.pipes[identity]
	return
}
