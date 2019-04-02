package search

import (
	"github.com/gobwas/glob"
	"github.com/hashicorp/consul/api"
	"github.com/pkg/errors"
	"regexp"
)

type ResultPair struct {
	Key   string
	Value string
}

type ConsulSearch struct {
	consulClient *api.Client
}

func (cs *ConsulSearch) SearchGlob(query string, path string) ([]ResultPair, error) {

	foundPairs := make([]ResultPair, 0)
	pattern := glob.MustCompile(query)

	pairs, err := cs.getKVPairs(path)
	if err != nil {
		return nil, errors.Wrap(err, "Could not get K/V pairs in glob search")
	}

	for _, element := range pairs {
		if pattern.Match(string(element.Value)) {
			pair := ResultPair{
				Key:   string(element.Key),
				Value: string(element.Value),
			}
			foundPairs = append(foundPairs, pair)
		}
	}

	return foundPairs, nil
}

func (cs *ConsulSearch) SearchRegex(query string, path string) ([]ResultPair, error) {
	foundPairs := make([]ResultPair, 0)
	pattern := regexp.MustCompile(query)

	pairs, err := cs.getKVPairs(path)
	if err != nil {
		return nil, errors.Wrap(err, "Could not get K/V pairs in regex search")
	}

	for _, element := range pairs {
		if pattern.Match(element.Value) {
			pair := ResultPair{
				Key:   string(element.Key),
				Value: string(element.Value),
			}
			foundPairs = append(foundPairs, pair)
		}
	}

	return foundPairs, nil
}

func (cs *ConsulSearch) getKVPairs(path string) (api.KVPairs, error) {
	kv := cs.consulClient.KV()
	pairs, _, err := kv.List(path, nil)

	return pairs, err
}

func NewConsulSearch(consulClient *api.Client) *ConsulSearch {
	return &ConsulSearch{
		consulClient: consulClient,
	}
}
