package image

import (
	"testing"

	"github.com/dnephin/dobi/config"
	"github.com/gotestyourself/gotestyourself/assert"
	is "github.com/gotestyourself/gotestyourself/assert/cmp"
)

func TestForEachTag(t *testing.T) {
	task := Task{
		config: &config.ImageConfig{
			Image: "imagename",
			Tags:  []string{"plain", "repo:tag"},
		},
	}

	expected := []string{"imagename:plain", "repo:tag"}
	tags := []string{}

	eachFunc := func(imageTag string) error {
		tags = append(tags, imageTag)
		return nil
	}

	err := task.ForEachTag(nil, eachFunc)
	assert.Check(t, is.NilError(err))
	assert.Check(t, is.Compare(expected, tags))
}
