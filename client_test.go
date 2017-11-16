package unleash_test

import (
	"testing"

	unleash "github.com/Unleash/unleash-client-go"
	"github.com/Unleash/unleash-client-go/context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type RolloutByGroupIDStrategy struct {
}

func (s *RolloutByGroupIDStrategy) Name() string {
	return "rolloutByGroupId"
}

func (s *RolloutByGroupIDStrategy) IsEnabled(map[string]interface{}, *context.Context) bool {
	return true
}

func TestGetEnabledFeatures(t *testing.T) {
	// given
	client, err := unleash.NewClient(
		unleash.WithAppName("fabric8"),
		unleash.WithUrl("http://localhost:4242/api/"),
		unleash.WithStrategies(&RolloutByGroupIDStrategy{}),
	)
	require.Nil(t, err)
	// when
	ctx := &context.Context{
		Properties: map[string]string{
			"groupID": "BETA",
		},
	}
	features := client.GetEnabledFeatures(ctx)
	// then
	assert.NotEmpty(t, features)

}
