package model

import (
	"reflect"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// PackState defines pack states
type PackState int8

// User states
const (
	Inactive PackState = 0
	Active   PackState = 1
)

// Mno contains the data of the Mobile Network Operator owner of the pack
type Mno struct {
	ID   int8   `json:"id" bson:"id"`     // id of the operator in the db
	Name string `json:"name" bson:"name"` // pack name
}

// Term contains information about the duration of pack usage.
type Term struct {
	UnitID int8   `json:"unit_id" bson:"unit_id"` // id of the unit of measurement
	Unit   string `json:"unit" bson:"unit"`       // unit name. ej. day, week.
	Amount int    `json:"amount" bson:"amount"`   // unit amount.
}

// Type contains the data of the type of the pack
type Type struct {
	ID   int8   `json:"id" bson:"id"`     // id of the type in the db
	Name string `json:"name" bson:"name"` // pack type name
}

// Currency of the price of the pack.
type Currency struct {
	ID   int8   `json:"id" bson:"id"`     // id of the currency
	Name string `json:"name" bson:"name"` // name of the currency.
}

// Resource contains the data of a pack resource.
type Resource struct {
	ID     int16   `json:"id" bson:"id"`
	Name   string  `json:"name" bson:"name"`
	Units  string  `json:"units" bson:"units"`
	Amount float32 `json:"amount" bson:"amount"`
	Isfree bool    `json:"isfree" bson:"isfree"`
}

// Pack contains the regarding to packs for admin purpose.
type Pack struct {
	ID        bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"` // id of the pack in the db
	ProdID    string        `json:"prodid" bson:"prodid"`              // internal mobile network provider package id
	Packcode  string        `json:"packcode" bson:"packcode"`          // pack code
	Name      string        `json:"name" bson:"name"`                  // pack name
	Desc      string        `json:"desc" bson:"desc"`                  // pack description
	Img       string        `json:"imgurl,omitempty" bson:"imgurl"`    // Icon image url for the pack
	Kwds      string        `json:"kwds" bson:"kwds"`                  // keywords for the pack searching
	Price     int           `json:"price" bson:"price"`                // price for the pack
	Stock     int           `json:"stock" bson:"stock"`                // pack stock
	Ownerid   int           `json:"ownerid"`                           // the company owner of the pack for resale
	Created   time.Time     `json:"created,omitempty" bson:"created"`
	Updated   time.Time     `json:"updated,omitempty" bson:"updated"`
	Packtype  *Type         `json:"type" bson:"type"`                               // pack type
	Mno       *Mno          `json:"mno" bson:"mno"`                                 // Mobile Network operator owner of the pack
	Term      *Term         `json:"term" bson:"term"`                               // Duration of the pack
	Ccy       *Currency     `json:"currency" bson:"currency"`                       // Currency of the price of the pack
	State     PackState     `json:"state,omitempty" bson:"state"`                   // state of the pack register
	Resources []Resource    `json:"resources,omitempty" bson:"resources,omitempty"` // resources that the pack contains
}

// PackExists contains pack data to check if the pack exists.
type PackExists struct {
	MnoID    int8   // Mno id owner of the pack
	ProdID   string // Product id of the pack
	Packcode string // Code of the pack
}

// Result contains result message
type Result struct {
	Code    string `json:"code"`
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}

// NewPackExists creates an instance of *PackExists from
// the given Pack.
func NewPackExists(pack *Pack) *PackExists {
	if pack == nil {
		return nil
	}
	packexists := new(PackExists)
	packexists.ProdID = pack.ProdID
	packexists.Packcode = pack.Packcode
	if pack.Mno != nil {
		packexists.MnoID = pack.Mno.ID
	}
	return packexists
}

// NewOKResult creates an instance of *Result with
// .Done = true and the values given in the param.
func NewOKResult(code string) *Result {
	resultst := new(Result)
	resultst.Success = true
	resultst.Code = code
	return resultst
}

// NewKOResult creates an instance of *Result with
// .Done = false and the values given in the param.
func NewKOResult(code string, msg string) *Result {
	resultst := new(Result)
	resultst.Success = false
	resultst.Code = code
	resultst.Msg = msg
	return resultst
}

// NewPackExistsFromMap creates a PackExists struct with the given parameters
func NewPackExistsFromMap(params map[string]interface{}) *PackExists {
	newkeys := new(PackExists)
	newkeys.MnoID = int8(params["mnoid"].(int))
	if _, ok := params["productid"]; ok {
		newkeys.ProdID = params["productid"].(string)
	}
	if _, ok := params["packcode"]; ok {
		newkeys.Packcode = params["packcode"].(string)
	}
	return newkeys
}

// NewPack creates a Pack struct with the given parameters
func NewPack(params map[string]interface{}) *Pack {
	newpack := new(Pack)
	newpack.ProdID = params["prodid"].(string)
	newpack.Packcode = params["packcode"].(string)
	newpack.Name = params["name"].(string)
	newpack.Desc = params["desc"].(string)
	newpack.Img = params["imgurl"].(string)
	newpack.Kwds = params["kwds"].(string)
	newpack.Price = params["price"].(int)
	newpack.Stock = 0
	newpack.Ownerid = params["ownerid"].(int)
	newpack.Packtype = NewType(params["type"].(map[string]interface{}))
	newpack.Mno = NewMno(params["mno"].(map[string]interface{}))
	newpack.Term = NewTerm(params["term"].(map[string]interface{}))
	newpack.Ccy = NewCurrency(params["currency"].(map[string]interface{}))
	if _, ok := params["resources"]; ok {
		newpack.Resources = NewResourcesFromInterface(params["resources"])
	}
	newpack.ID = ""
	newpack.State = Inactive
	newpack.Created = time.Time{}
	newpack.Updated = time.Time{}
	return newpack
}

// NewResourcesFromInterface from a given interface that must be a slice
// it get the items from interface using reflection. If the given parameter
// is not a slice an empty array is returned.
func NewResourcesFromInterface(params interface{}) []Resource {
	var resources []Resource
	// get the array parameter for resources.
	switch reflect.TypeOf(params).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(params)
		resources = make([]Resource, s.Len())

		for i := 0; i < s.Len(); i++ {
			valo := s.Index(i)
			resource := NewResource(valo.Interface().(map[string]interface{}))
			resources[i] = *resource
		}
	default:
		resources = []Resource{}
	}
	return resources
}

// NewResources creates an array of new Resource struct from a map.
// if param is null it returns an empty array.
func NewResources(params []map[string]interface{}) []Resource {
	if params != nil {
		array := make([]Resource, len(params))
		for i := 0; i < len(params); i++ {
			array[i] = *NewResource(params[i])
		}
		return array
	}
	return []Resource{}
}

// NewResource creates a new Resource struct from a map.
func NewResource(params map[string]interface{}) *Resource {
	new := new(Resource)
	new.ID = int16(params["id"].(int))
	new.Name = params["name"].(string)
	new.Units = params["units"].(string)
	new.Amount = float32(params["amount"].(float64))
	new.Isfree = params["isfree"].(bool)
	return new
}

// NewMno creates a new mno struct from a map
func NewMno(params map[string]interface{}) *Mno {
	new := new(Mno)
	new.ID = int8(params["id"].(int))
	new.Name = params["name"].(string)
	return new
}

// NewTerm creates a new term struct from a map
func NewTerm(params map[string]interface{}) *Term {
	new := new(Term)
	new.UnitID = int8(params["unit_id"].(int))
	new.Unit = params["unit"].(string)
	new.Amount = params["amount"].(int)
	return new
}

// NewType creates a new type struct from a map
func NewType(params map[string]interface{}) *Type {
	new := new(Type)
	new.ID = int8(params["id"].(int))
	new.Name = params["name"].(string)
	return new
}

// NewCurrency creates a new currency struct from a map
func NewCurrency(params map[string]interface{}) *Currency {
	new := new(Currency)
	new.ID = int8(params["id"].(int))
	new.Name = params["name"].(string)
	return new
}

// NewPackState from a int8 value we return a packstate.
func NewPackState(state int) PackState {
	var result PackState
	newstate := PackState(state)
	switch newstate {
	case Active:
		result = Active
	case Inactive:
		result = Inactive
	default:
		result = Active
	}

	return result
}
