package graphql

import (
	"testing"

	client "github.com/sensu/sensu-go/backend/apid/graphql/mockclient"
	"github.com/sensu/sensu-go/backend/apid/graphql/schema"
	"github.com/sensu/sensu-go/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestQueryTypeEventField(t *testing.T) {
	client, factory := client.NewClientFactory()
	impl := queryImpl{factory: factory}

	event := types.FixtureEvent("a", "b")
	args := schema.QueryEventFieldResolverArgs{Namespace: "ns", Entity: "a", Check: "b"}
	params := schema.QueryEventFieldResolverParams{Args: args}

	// Success
	client.On("FetchEvent", event.Entity.Name, event.Check.Name).Return(event, nil).Once()
	res, err := impl.Event(params)
	require.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestQueryTypeEventFilterField(t *testing.T) {
	client, factory := client.NewClientFactory()
	impl := queryImpl{factory: factory}

	filter := types.FixtureEventFilter("a")
	args := schema.QueryEventFilterFieldResolverArgs{Namespace: "ns", Name: "a"}
	params := schema.QueryEventFilterFieldResolverParams{Args: args}

	// Success
	client.On("FetchFilter", filter.Name).Return(filter, nil).Once()
	res, err := impl.EventFilter(params)
	require.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestQueryTypeNamespaceField(t *testing.T) {
	client, factory := client.NewClientFactory()
	impl := queryImpl{factory: factory}

	nsp := types.FixtureNamespace("sensu")
	params := schema.QueryNamespaceFieldResolverParams{}
	params.Args.Name = nsp.Name

	// Success
	client.On("FetchNamespace", nsp.Name).Return(nsp, nil).Once()
	res, err := impl.Namespace(params)
	require.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestQueryTypeEntityField(t *testing.T) {
	client, factory := client.NewClientFactory()
	impl := queryImpl{factory: factory}

	entity := types.FixtureEntity("a")
	params := schema.QueryEntityFieldResolverParams{}
	params.Args.Namespace = entity.Namespace
	params.Args.Name = entity.Name

	// Sucess
	client.On("FetchEntity", entity.Name).Return(entity, nil).Once()
	res, err := impl.Entity(params)
	require.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestQueryTypeCheckField(t *testing.T) {
	client, factory := client.NewClientFactory()
	impl := queryImpl{factory: factory}

	check := types.FixtureCheckConfig("a")
	params := schema.QueryCheckFieldResolverParams{}
	params.Args.Namespace = check.Namespace
	params.Args.Name = check.Name

	// Sucess
	client.On("FetchCheck", check.Name).Return(check, nil).Once()
	res, err := impl.Check(params)
	require.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestQueryTypeHandlerField(t *testing.T) {
	client, factory := client.NewClientFactory()
	impl := queryImpl{factory: factory}

	handler := types.FixtureHandler("a")
	params := schema.QueryHandlerFieldResolverParams{}
	params.Args.Namespace = handler.Namespace
	params.Args.Name = handler.Name

	// Success
	client.On("FetchHandler", handler.Name).Return(handler, nil).Once()
	res, err := impl.Handler(params)
	require.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestQueryTypeMuatorField(t *testing.T) {
	client, factory := client.NewClientFactory()
	impl := queryImpl{factory: factory}

	mutator := types.FixtureMutator("a")
	params := schema.QueryMutatorFieldResolverParams{}
	params.Args.Namespace = mutator.Namespace
	params.Args.Name = mutator.Name

	// Success
	client.On("FetchMutator", mutator.Name).Return(mutator, nil).Once()
	res, err := impl.Mutator(params)
	require.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestQueryTypeSuggestField(t *testing.T) {
	client, factory := client.NewClientFactory()
	impl := queryImpl{factory: factory}

	params := schema.QuerySuggestFieldResolverParams{}
	params.Args.Namespace = "default"
	params.Args.Ref = "core/v2/check_config/subscriptions"

	// Success
	client.On("List", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
	res, err := impl.Suggest(params)
	require.NoError(t, err)
	assert.NotEmpty(t, res)
}
