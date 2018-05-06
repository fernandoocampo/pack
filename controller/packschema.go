package controller

import (
	"github.com/fernandoocampo/pack/model"
	"github.com/graphql-go/graphql"
)

// typeInterface is the type of the pack.
var typeInterface = graphql.NewObject(graphql.ObjectConfig{
	Name:        "PackType",
	Description: "The type of a pack",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.Int,
			Description: "The id of the pack type.",
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: "Name of the pack type.",
		},
	},
})

// inputType contains info for pack type argument
var inputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "inputType",
		Fields: graphql.InputObjectConfigFieldMap{
			"id": &graphql.InputObjectFieldConfig{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "The id of the pack type",
			},
			"name": &graphql.InputObjectFieldConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The name of the pack type",
			},
		},
	},
)

// termInterface is the validity of the pack.
var termInterface = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Term",
	Description: "The validity of a pack",
	Fields: graphql.Fields{
		"unit_id": &graphql.Field{
			Type:        graphql.Int,
			Description: "The id of the unit.",
		},
		"unit": &graphql.Field{
			Type:        graphql.String,
			Description: "The name of units. e.g. days, weeks, etc.",
		},
		"amount": &graphql.Field{
			Type:        graphql.Int,
			Description: "number of units.",
		},
	},
})

// resourceInterface is a resource of a pack.
var resourceInterface = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Resource",
	Description: "The resource of a pack",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.Int,
			Description: "The id of the resource.",
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: "The name of the resource. e.g. internet, whatsapp, etc.",
		},
		"units": &graphql.Field{
			Type:        graphql.String,
			Description: "the units name used for this resource.",
		},
		"amount": &graphql.Field{
			Type:        graphql.Int,
			Description: "number of units.",
		},
		"isfree": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "indicates if the resource is free or not.",
		},
	},
})

// inputTerm contains info for the validity of the pack
var inputTerm = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "inputTerm",
		Fields: graphql.InputObjectConfigFieldMap{
			"unit_id": &graphql.InputObjectFieldConfig{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "The id of the unit used for the pack validity",
			},
			"unit": &graphql.InputObjectFieldConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The name of the unit used for the pack validity",
			},
			"amount": &graphql.InputObjectFieldConfig{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "The amount of the unit used for the pack validity",
			},
		},
	},
)

// mnoInterface is the network operator owner of the pack.
var mnoInterface = graphql.NewObject(graphql.ObjectConfig{
	Name:        "MnoPack",
	Description: "The network operator owner of the pack.",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.Int,
			Description: "The id of the network operator.",
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: "Name of the network operator.",
		},
	},
})

// inputMno contains info for the mno owner of the pack
var inputMno = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "inputMno",
		Fields: graphql.InputObjectConfigFieldMap{
			"id": &graphql.InputObjectFieldConfig{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "The id of the mobile network operator",
			},
			"name": &graphql.InputObjectFieldConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The name of the mobile network operator",
			},
		},
	},
)

// ccyInterface is the currency used for the pack price.
var ccyInterface = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Currency",
	Description: "The currency used for the pack price.",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.Int,
			Description: "The id of the currency.",
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: "Name of the currency.",
		},
	},
})

// inputCcy contains info for the mno owner of the pack
var inputCcy = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "inputCurrency",
		Fields: graphql.InputObjectConfigFieldMap{
			"id": &graphql.InputObjectFieldConfig{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "The id of the currency used for the price",
			},
			"name": &graphql.InputObjectFieldConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The name of the currency used for the price",
			},
		},
	},
)

// resultType contains data for operations result.
var resultType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Result",
	Fields: graphql.Fields{
		"code": &graphql.Field{
			Type: graphql.String,
		},
		"success": &graphql.Field{
			Type: graphql.Boolean,
		},
		"msg": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// resourceType contains data for resource of a pack.
var resourceType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "Resource",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type:        graphql.Int,
			Description: "id of the kind of resource",
		},
		"name": &graphql.InputObjectFieldConfig{
			Type:        graphql.String,
			Description: "name of the resource: e.g. data, voice, sms, etc.",
		},
		"units": &graphql.InputObjectFieldConfig{
			Type:        graphql.String,
			Description: "unit of measurement of the resource.",
		},
		"amount": &graphql.InputObjectFieldConfig{
			Type:        graphql.Float,
			Description: "amount of units of measurement of the resource.",
		},
		"isfree": &graphql.InputObjectFieldConfig{
			Type:        graphql.Boolean,
			Description: "indicates if this resource is free.",
		},
	},
})

// Pack information
var packType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Pack",
	Description: "A pack with basic information",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.String,
			Description: "The id of the pack.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// required to return just the hex representation of objectid
				pack := p.Source.(*model.Pack)
				if pack == nil {
					return nil, nil
				}
				idhex := pack.ID.Hex()
				return idhex, nil
			},
		},
		"packcode": &graphql.Field{
			Type:        graphql.String,
			Description: "pack code.",
		},
		"prodid": &graphql.Field{
			Type:        graphql.String,
			Description: "pack internal id.",
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: "Name of the pack.",
		},
		"desc": &graphql.Field{
			Type:        graphql.String,
			Description: "Description of the pack.",
		},
		"imgurl": &graphql.Field{
			Type:        graphql.String,
			Description: "Image url.",
		},
		"kwds": &graphql.Field{
			Type:        graphql.String,
			Description: "Key words for searching.",
		},
		"price": &graphql.Field{
			Type:        graphql.Int,
			Description: "Pack sales price.",
		},
		"ownerid": &graphql.Field{
			Type:        graphql.Int,
			Description: "id of the owner of the pack.",
		},
		"type": &graphql.Field{
			Type:        typeInterface,
			Description: "Pack type e.g. Internet, sms, etc.",
		},
		"mno": &graphql.Field{
			Type:        mnoInterface,
			Description: "Mobile network operator name.",
		},
		"term": &graphql.Field{
			Type:        termInterface,
			Description: "Pack validity.",
		},
		"currency": &graphql.Field{
			Type:        ccyInterface,
			Description: "Currency of the pack's price.",
		},
		"state": &graphql.Field{
			Type:        graphql.Int,
			Description: "state of the pack. 1. active, 2. deactive",
		},
		"resources": &graphql.Field{
			Type:        graphql.NewList(resourceInterface),
			Description: "a list of resources that this pack provides",
		},
	},
})

// root query for packs
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "packQuery",
	Fields: graphql.Fields{
		"byID": &graphql.Field{
			Type:        packType,
			Description: "query a pack by its id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return getByID(params)
			},
		},
		"byCode": &graphql.Field{
			Type:        packType,
			Description: "query a pack by its code",
			Args: graphql.FieldConfigArgument{
				"packcode": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return getByPackCode(params)
			},
		},
		"idByCode": &graphql.Field{
			Type:        graphql.String,
			Description: "query a pack id by its code",
			Args: graphql.FieldConfigArgument{
				"packcode": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return getIDByCode(params)
			},
		},
		"byProductID": &graphql.Field{
			Type:        packType,
			Description: "query a pack by its internal product id",
			Args: graphql.FieldConfigArgument{
				"productid": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return getByProductID(params)
			},
		},
		"byKeys": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "query a pack by mno id and (productid or packcode)",
			Args: graphql.FieldConfigArgument{
				"mnoid": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"packcode": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"productid": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return getByKeys(params)
			},
		},
	},
})

// packMutation root mutation schema for User, here we specify the app capabilities.
// Note that
// - Here we specify the capabilities.
// - Here we set the arguments in each capability.
// - Here we specify the function that executes this capability.
// - It is separate of queries
var packMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "PackMutation",
	Fields: graphql.Fields{
		/*
			create a pack.
		*/
		"create": &graphql.Field{
			Type:        resultType, // the return type for this field
			Description: "creates a new Pack",
			Args: graphql.FieldConfigArgument{
				"prodid": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"packcode": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"desc": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"imgurl": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"kwds": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"price": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"ownerid": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"type": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(inputType),
				},
				"mno": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(inputMno),
				},
				"term": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(inputTerm),
				},
				"currency": &graphql.ArgumentConfig{
					Type: inputCcy,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return create(params)
			},
		},
		/*
			change state of a pack
		*/
		"changeState": &graphql.Field{
			Type:        resultType, // the return type for this field
			Description: "updates a pack state",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"state": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return changeState(params)
			},
		},
		/*
			change product id of a pack
		*/
		"changeProductID": &graphql.Field{
			Type:        resultType, // the return type for this field
			Description: "updates a pack internal product id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"mnoid": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"productid": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return changeProductID(params)
			},
		},
		/*
			change pack code
		*/
		"changePackCode": &graphql.Field{
			Type:        resultType, // the return type for this field
			Description: "updates a pack code",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"mnoid": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"packcode": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return changePackCode(params)
			},
		},
		/*
			change pack name
		*/
		"changeName": &graphql.Field{
			Type:        resultType, // the return type for this field
			Description: "updates a pack name",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"newname": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return changeName(params)
			},
		},
		/*
			change pack description
		*/
		"changeDescription": &graphql.Field{
			Type:        resultType, // the return type for this field
			Description: "updates a pack description",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"newdesc": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return changeDesc(params)
			},
		},
		/*
			change pack image url
		*/
		"changeImageUrl": &graphql.Field{
			Type:        resultType, // the return type for this field
			Description: "updates a pack image url",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"newimgurl": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return changeImg(params)
			},
		},
		/*
			change pack key words
		*/
		"changeKeywords": &graphql.Field{
			Type:        resultType, // the return type for this field
			Description: "updates a pack search keywords",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"newkeywords": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return changeKeyword(params)
			},
		},
		/*
			change pack price
		*/
		"changePrice": &graphql.Field{
			Type:        resultType, // the return type for this field
			Description: "updates a pack price",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"newprice": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return changePrice(params)
			},
		},
		/*
			change pack type
		*/
		"changeType": &graphql.Field{
			Type:        resultType, // the return type for this field
			Description: "updates a pack type",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"newtype": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(inputType),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return changePackType(params)
			},
		},
		/*
			change pack mobile network operator owner
		*/
		"changeMno": &graphql.Field{
			Type:        resultType, // the return type for this field
			Description: "updates a pack mobile network operator owner",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"productid": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"packcode": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"newmno": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(inputMno),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return changeMNO(params)
			},
		},
		/*
			change pack validity
		*/
		"changeValidity": &graphql.Field{
			Type:        resultType, // the return type for this field
			Description: "updates a pack validity duration",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"newterm": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(inputTerm),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return changeValidity(params)
			},
		},
		/*
			change pack's price currency.
		*/
		"changeCurrency": &graphql.Field{
			Type:        resultType, // the return type for this field
			Description: "updates a pack's price currency",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"newcurrency": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(inputCcy),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return changeCurrency(params)
			},
		},
		/*
			update resources of the given pack.
		*/
		"replaceResources": &graphql.Field{
			Type:        resultType, // the return type for this field
			Description: "updates pack's resources",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.String),
					Description: "Id of the pack",
				},
				"newresources": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.NewList(resourceType)),
					Description: "The new list of resources for this pack",
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return replaceResources(params)
			},
		},
		/*
			move pack stock to increase and reduce inventory
		*/
		"moveStock": &graphql.Field{
			Type:        resultType, // the return type for this field
			Description: "updates the stock of a pack, send a negative value to reduce.",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"amount": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return moveStock(params)
			},
		},
		/*
			delete a pack
		*/
		"delete": &graphql.Field{
			Type:        resultType, // the return type for this field
			Description: "deletes a given pack",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return delete(params)
			},
		},
		/*
			delete a pack resources.
		*/
		"deletePackResources": &graphql.Field{
			Type:        resultType, // the return type for this field
			Description: "deletes the resources of a given pack",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return deletePackResources(params)
			},
		},
	},
})
