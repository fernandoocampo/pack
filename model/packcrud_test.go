package model

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

// TestNewPack tests instances a Pack struct
func TestNewPack(t *testing.T) {
	expresult := createExpPack()
	// GIVEN data to build a pack
	params := createGivenParam()

	// When we need to get a pack struct
	result := NewPack(params)

	// THEN system returns a good Pack
	if result == nil {
		t.Fatalf("Expected a Pack struct with data but got nil")
	}

	validatePackField(expresult, result, t)

}

// Test if we convert a json value to Pack struct works well
func TestJsonNewPack(t *testing.T) {
	expresult := createExpPack()
	// GIVEN data to build a pack
	param := createJSONParam()
	var result *Pack
	err := json.Unmarshal(param, &result)

	if err != nil {
		t.Fatalf("Expected to work but got err: %s", err)
	}

	validatePackField(expresult, result, t)
}

// TestNewOKResult tests instances a Result struct
func TestNewOKResult(t *testing.T) {
	expresult := Result{
		Success: true,
		Code:    "10",
	}
	// GIVEN data to build a result.
	code := "10"
	// WHEN we need to encapsulate the data to return result data.
	result := NewOKResult(code)
	// THEN system returns a good LoginResult
	if result == nil {
		t.Fatalf("Expected a Result struct with data but got nil")
	}
	if result.Success != expresult.Success {
		t.Fatalf("Expected Result#Success %t but got %t", expresult.Success, result.Success)
	}

	if result.Msg != expresult.Msg {
		t.Fatalf("Expected Result#Msg %s but got %s", expresult.Msg, result.Msg)
	}

	if result.Code != expresult.Code {
		t.Fatalf("Expected Result#Code %s but got %s", expresult.Code, result.Code)
	}
}

// TestNewPackExists tests instances a PackExists struct
func TestNewPackExists(t *testing.T) {
	// GIVEN data for create a new pack exists
	newpack := createExpPack()
	exppackexists := createExpPackExists()
	// WHEN we need to create a pack exists
	result := NewPackExists(newpack)

	// THEN we check if the creation is ok.
	if result == nil {
		t.Fatalf("Expected a Result struct with data but got nil")
	}

	if result.MnoID != exppackexists.MnoID {
		t.Fatalf("Expected a Result.MnoID to be %d but got %d", exppackexists.MnoID, result.MnoID)
	}

	if result.Packcode != exppackexists.Packcode {
		t.Fatalf("Expected a Result.Packcode to be %s but got %s", exppackexists.Packcode, result.Packcode)
	}

	if result.ProdID != exppackexists.ProdID {
		t.Fatalf("Expected a Result.ProdID to be %s but got %s", exppackexists.ProdID, result.ProdID)
	}
}

// TestNewPackExistsFromMap verify function that converts a map to PackExists struct.
func TestNewPackExistsFromMap(t *testing.T) {
	// GIVEN data for create a new pack exists
	params := make(map[string]interface{})
	params["mnoid"] = 2
	params["productid"] = "3"
	params["packcode"] = "wh12"
	exppackexists := createExpPackExists()
	// WHEN we need to create a pack exists
	result := NewPackExistsFromMap(params)

	// THEN we check if the creation is ok.
	if result == nil {
		t.Fatalf("Expected a Result struct with data but got nil")
	}

	if result.MnoID != exppackexists.MnoID {
		t.Fatalf("Expected a Result.MnoID to be %d but got %d", exppackexists.MnoID, result.MnoID)
	}

	if result.Packcode != exppackexists.Packcode {
		t.Fatalf("Expected a Result.Packcode to be %s but got %s", exppackexists.Packcode, result.Packcode)
	}

	if result.ProdID != exppackexists.ProdID {
		t.Fatalf("Expected a Result.ProdID to be %s but got %s", exppackexists.ProdID, result.ProdID)
	}
}

func validatePackField(expected interface{}, result interface{}, t *testing.T) {
	val := reflect.ValueOf(result).Elem()
	expval := reflect.ValueOf(expected).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		fieldName := val.Type().Field(i).Name
		expValueField := expval.FieldByName(fieldName)
		valueData := valueField.Interface()
		expValueData := expValueField.Interface()
		switch fieldName {
		case "Packtype":
			validatePackField(valueData, expValueData, t)
		case "Mno":
			validatePackField(valueData, expValueData, t)
		case "Term":
			validatePackField(valueData, expValueData, t)
		case "Ccy":
			validatePackField(valueData, expValueData, t)
		default:
			if valueData != expValueData {
				t.Fatalf("Expected %s %s but got %s\n", fieldName, expValueData, expValueData)
			}
		}

	}
}

func createJSONParam() []byte {
	tag := `{"prodid":"3","packcode":"wh12","name":"Whatsapp weekend",
		"desc":"Whatsapp para el fin de semana, para que hables con tus amigos todo el dia.",
		"imgurl":"/appdata/img/whatsappwknd.png","kwds":"internet whatsapp dia fin semana",
		"price":2500,"type":{"id":1,"name":"App"},"mno":{"id":2,"name":"Claro"},
		"term":{"unit_id":4,"unit":"dia","amount":2},"currency":{"id":3,"name":"cop"}}`
	bytetag := []byte(tag)
	return bytetag
}

func createGivenParam() map[string]interface{} {
	mnoparams := make(map[string]interface{})
	mnoparams["id"] = 2
	mnoparams["name"] = "Claro"
	typeparams := make(map[string]interface{})
	typeparams["id"] = 1
	typeparams["name"] = "App"
	termparams := make(map[string]interface{})
	termparams["unit_id"] = 4
	termparams["unit"] = "dia"
	termparams["amount"] = 2
	ccyparams := make(map[string]interface{})
	ccyparams["id"] = 3
	ccyparams["name"] = "cop"
	params := make(map[string]interface{})
	params["prodid"] = "3"
	params["packcode"] = "wh12"
	params["name"] = "Whatsapp weekend"
	params["desc"] = "Whatsapp para el fin de semana, para que hables con tus amigos todo el dia."
	params["imgurl"] = "/appdata/img/whatsappwknd.png"
	params["kwds"] = "internet whatsapp dia fin semana"
	params["price"] = 2500
	params["type"] = typeparams
	params["mno"] = mnoparams
	params["term"] = termparams
	params["currency"] = ccyparams
	return params
}

func createExpPackExists() *PackExists {
	exppackexists := PackExists{
		MnoID:    2,
		Packcode: "wh12",
		ProdID:   "3",
	}
	return &exppackexists
}

func createExpPack() *Pack {
	expresult := Pack{
		ID:       "",
		ProdID:   "3",
		Packcode: "wh12",
		Name:     "Whatsapp weekend",
		Desc:     "Whatsapp para el fin de semana, para que hables con tus amigos todo el dia.",
		Img:      "/appdata/img/whatsappwknd.png",
		Kwds:     "internet whatsapp dia fin semana",
		Price:    2500,
		Created:  time.Time{},
		Updated:  time.Time{},
		Packtype: &Type{
			ID:   1,
			Name: "App",
		},
		Mno: &Mno{
			ID:   2,
			Name: "Claro",
		},
		Term: &Term{
			UnitID: 4,
			Unit:   "dia",
			Amount: 2,
		},
		Ccy: &Currency{
			ID:   3,
			Name: "cop",
		},
		State: Inactive,
	}
	return &expresult
}
