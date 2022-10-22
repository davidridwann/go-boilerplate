package loader

import (
	"github.com/davidridwann/wlb-test.git/pkg/env"
	"github.com/davidridwann/wlb-test.git/pkg/log"
	"gopkg.in/yaml.v2"
	"io"
)

type YAMLLoader[Dest any] struct {
}

func (Y YAMLLoader[Dest]) Load(reader io.Reader) (Dest, error) {
	var destination Dest

	log.Std().Infof("loading yaml [%s] configuration", env.Current())
	if err := yaml.NewDecoder(reader).Decode(&destination); err != nil {
		log.Err().Fatalln("cannot unmarshal yaml configuration from reader, err:", err)
	}

	return destination, nil
}
