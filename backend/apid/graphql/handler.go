package graphql

import (
	v2 "github.com/sensu/sensu-go/api/core/v2"
	"github.com/sensu/sensu-go/backend/apid/graphql/globalid"
	"github.com/sensu/sensu-go/backend/apid/graphql/schema"
	"github.com/sensu/sensu-go/graphql"
	"github.com/sensu/sensu-go/types"
	"github.com/sensu/sensu-go/util/strings"
)

var _ schema.HandlerFieldResolvers = (*handlerImpl)(nil)
var _ schema.HandlerSocketFieldResolvers = (*handlerSocketImpl)(nil)

//
// Implement HandlerFieldResolvers
//

type handlerImpl struct {
	schema.HandlerAliases
	factory ClientFactory
}

// ID implements response to request for 'id' field.
func (*handlerImpl) ID(p graphql.ResolveParams) (string, error) {
	return globalid.HandlerTranslator.EncodeToString(p.Source), nil
}

// Mutator implements response to request for 'mutator' field.
func (r *handlerImpl) Mutator(p graphql.ResolveParams) (interface{}, error) {
	src := p.Source.(*types.Handler)
	if src.Mutator == "" {
		return nil, nil
	}

	ctx := types.SetContextFromResource(p.Context, src)
	client := r.factory.NewWithContext(ctx)
	res, err := client.FetchMutator(src.Mutator)

	return handleFetchResult(res, err)
}

// Handlers implements response to request for 'handlers' field.
func (r *handlerImpl) Handlers(p graphql.ResolveParams) (interface{}, error) {
	src := p.Source.(*types.Handler)
	results, err := loadHandlers(p.Context, src.Namespace)
	records := filterHandlers(results, func(obj *types.Handler) bool {
		return strings.FoundInArray(obj.Name, src.Handlers)
	})
	return records, err
}

// IsTypeOf is used to determine if a given value is associated with the type
func (*handlerImpl) IsTypeOf(s interface{}, p graphql.IsTypeOfParams) bool {
	_, ok := s.(*types.Handler)
	return ok
}

// ToJSON implements response to request for 'toJSON' field.
func (*handlerImpl) ToJSON(p graphql.ResolveParams) (interface{}, error) {
	return types.WrapResource(p.Source.(v2.Resource)), nil
}

//
// Implement HandlerSocketFieldResolvers
//

type handlerSocketImpl struct{}

// Host implements response to request for 'host' field.
func (*handlerSocketImpl) Host(p graphql.ResolveParams) (string, error) {
	socket := p.Source.(*types.HandlerSocket)
	return socket.Host, nil
}

// Port implements response to request for 'port' field.
func (*handlerSocketImpl) Port(p graphql.ResolveParams) (int, error) {
	socket := p.Source.(*types.HandlerSocket)
	return int(socket.Port), nil
}

// IsTypeOf is used to determine if a given value is associated with the type
func (*handlerSocketImpl) IsTypeOf(s interface{}, p graphql.IsTypeOfParams) bool {
	_, ok := s.(*types.HandlerSocket)
	return ok
}
