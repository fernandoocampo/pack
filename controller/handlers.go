package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fernandoocampo/pack/model"
	"github.com/fernandoocampo/pack/service"
	"github.com/graphql-go/graphql"
)

// healthservice is the reference to the health service object.
var healthservice service.IHealthService

// packService references the IPackService
var packService service.IPackService

// GetByID implements *IPackService.GetByID using mongo implementation.
func getByID(params graphql.ResolveParams) (interface{}, error) {
	packid, _ := params.Args["id"].(string)
	result, err := packService.FindByID(packid)
	return result, err
}

// GetByPackCode implements *IPackService.GetByPackCode.
func getByPackCode(params graphql.ResolveParams) (interface{}, error) {
	packcode, _ := params.Args["packcode"].(string)
	return packService.GetByPackCode(packcode)
}

// GetByProductID implements *IPackService.GetByProductId.
func getByProductID(params graphql.ResolveParams) (interface{}, error) {
	prodid, _ := params.Args["productid"].(string)
	return packService.GetByProductID(prodid)
}

// GetIDByCode implements *IPackDAO.GetIDByCode.
func getIDByCode(params graphql.ResolveParams) (interface{}, error) {
	packcode, _ := params.Args["packcode"].(string)
	return packService.GetIDByCode(packcode)
}

// getByKeys implements *IPackDAO.GetIDByCode.
func getByKeys(params graphql.ResolveParams) (interface{}, error) {
	keys := model.NewPackExistsFromMap(params.Args)
	return packService.IsThereThisPack(keys)
}

// Create implements *IPackService.Create.
func create(params graphql.ResolveParams) (interface{}, error) {
	pack := model.NewPack(params.Args)

	err := packService.Create(pack)
	if err != nil {
		return model.NewKOResult("-1", err.Error()), nil
	}
	return model.NewOKResult("10"), nil
}

// changeState implements *IPackService.changeState.
func changeState(params graphql.ResolveParams) (interface{}, error) {
	id, _ := params.Args["id"].(string)
	state, _ := params.Args["state"].(int)

	newstate := model.NewPackState(state)

	err := packService.ChangeState(id, newstate)

	if err != nil {
		return model.NewKOResult("-1", err.Error()), nil
	}
	return model.NewOKResult("10"), nil
}

// changeProductID implements *IPackService.changeProductID.
func changeProductID(params graphql.ResolveParams) (interface{}, error) {
	id, _ := params.Args["id"].(string)
	mnoid, _ := params.Args["mnoid"].(int)
	prodid, _ := params.Args["productid"].(string)

	err := packService.ChangeProductID(id, int8(mnoid), prodid)

	if err != nil {
		return model.NewKOResult("-1", err.Error()), nil
	}
	return model.NewOKResult("10"), nil
}

// changePackCode implements *IPackService.changePackCode.
func changePackCode(params graphql.ResolveParams) (interface{}, error) {
	id, _ := params.Args["id"].(string)
	mnoid, _ := params.Args["mnoid"].(int)
	packcode, _ := params.Args["packcode"].(string)

	err := packService.ChangePackCode(id, int8(mnoid), packcode)

	if err != nil {
		return model.NewKOResult("-1", err.Error()), nil
	}
	return model.NewOKResult("10"), nil
}

// changeName implements *IPackService.changeName.
func changeName(params graphql.ResolveParams) (interface{}, error) {
	id, _ := params.Args["id"].(string)
	name, _ := params.Args["newname"].(string)

	err := packService.ChangeName(id, name)

	if err != nil {
		return model.NewKOResult("-1", err.Error()), nil
	}
	return model.NewOKResult("10"), nil
}

// changeDesc implements *IPackService.changeDesc.
func changeDesc(params graphql.ResolveParams) (interface{}, error) {
	id, _ := params.Args["id"].(string)
	desc, _ := params.Args["newdesc"].(string)

	err := packService.ChangeDesc(id, desc)

	if err != nil {
		return model.NewKOResult("-1", err.Error()), nil
	}
	return model.NewOKResult("10"), nil
}

// changeImg implements *IPackService.changeImg.
func changeImg(params graphql.ResolveParams) (interface{}, error) {
	id, _ := params.Args["id"].(string)
	img, _ := params.Args["newimgurl"].(string)

	err := packService.ChangeImg(id, img)

	if err != nil {
		return model.NewKOResult("-1", err.Error()), nil
	}
	return model.NewOKResult("10"), nil
}

// changeKeyword implements *IPackService.changeKeyword.
func changeKeyword(params graphql.ResolveParams) (interface{}, error) {
	id, _ := params.Args["id"].(string)
	kwds, _ := params.Args["newkeywords"].(string)

	err := packService.ChangeKeyword(id, kwds)

	if err != nil {
		return model.NewKOResult("-1", err.Error()), nil
	}
	return model.NewOKResult("10"), nil
}

// changePrice implements *IPackService.changePrice.
func changePrice(params graphql.ResolveParams) (interface{}, error) {
	id, _ := params.Args["id"].(string)
	price, _ := params.Args["newprice"].(int)

	err := packService.ChangePrice(id, price)

	if err != nil {
		return model.NewKOResult("-1", err.Error()), nil
	}
	return model.NewOKResult("10"), nil
}

// changePackType implements *IPackService.changePackType.
func changePackType(params graphql.ResolveParams) (interface{}, error) {
	id, _ := params.Args["id"].(string)
	ptype, _ := params.Args["newtype"].(map[string]interface{})

	newtype := model.NewType(ptype)

	err := packService.ChangePackType(id, newtype)

	if err != nil {
		return model.NewKOResult("-1", err.Error()), nil
	}
	return model.NewOKResult("10"), nil
}

// changeMNO implements *IPackService.changeMNO.
func changeMNO(params graphql.ResolveParams) (interface{}, error) {
	id, _ := params.Args["id"].(string)
	pmno, _ := params.Args["newmno"].(map[string]interface{})
	prodid, _ := params.Args["productid"].(string)
	packcode, _ := params.Args["packcode"].(string)

	newmno := model.NewMno(pmno)

	err := packService.ChangeMNO(id, prodid, packcode, newmno)

	if err != nil {
		return model.NewKOResult("-1", err.Error()), nil
	}
	return model.NewOKResult("10"), nil
}

// changeValidity implements *IPackService.changeValidity.
func changeValidity(params graphql.ResolveParams) (interface{}, error) {
	id, _ := params.Args["id"].(string)
	pterm, _ := params.Args["newterm"].(map[string]interface{})

	newterm := model.NewTerm(pterm)

	err := packService.ChangeValidity(id, newterm)

	if err != nil {
		return model.NewKOResult("-1", err.Error()), nil
	}
	return model.NewOKResult("10"), nil
}

// changeCurrency implements *IPackService.changeCurrency.
func changeCurrency(params graphql.ResolveParams) (interface{}, error) {
	id, _ := params.Args["id"].(string)
	ccy, _ := params.Args["newcurrency"].(map[string]interface{})

	newccy := model.NewCurrency(ccy)

	err := packService.ChangeCurrency(id, newccy)

	if err != nil {
		return model.NewKOResult("-1", err.Error()), nil
	}
	return model.NewOKResult("10"), nil
}

// Delete delete a pack
func delete(params graphql.ResolveParams) (interface{}, error) {
	id, _ := params.Args["id"].(string)
	err := packService.Delete(id)

	if err != nil {
		return model.NewKOResult("-1", err.Error()), nil
	}
	return model.NewOKResult("10"), nil
}

// replaceResources replaces the resources configured for a pack
func replaceResources(params graphql.ResolveParams) (interface{}, error) {
	id, _ := params.Args["id"].(string)
	newresources, _ := params.Args["newresources"]

	// get the array parameter for resources.
	resources := model.NewResourcesFromInterface(newresources)

	err := packService.UpdateResources(id, resources)

	if err != nil {
		return model.NewKOResult("-1", err.Error()), nil
	}

	return model.NewOKResult("10"), nil
}

// Delete delete resources of a pack
func deletePackResources(params graphql.ResolveParams) (interface{}, error) {
	id, _ := params.Args["id"].(string)
	err := packService.DeleteResources(id)

	if err != nil {
		return model.NewKOResult("-1", err.Error()), nil
	}

	return model.NewOKResult("10"), nil
}

// moveStock change the stock of a pack.
func moveStock(params graphql.ResolveParams) (interface{}, error) {
	id, _ := params.Args["id"].(string)
	amount, _ := params.Args["amount"].(int)
	err := packService.MoveStock(id, amount)

	if err != nil {
		return model.NewKOResult("-1", err.Error()), nil
	}
	return model.NewOKResult("10"), nil
}

// SetService  set the pack service for this business logic.
func SetService(service service.IPackService) {
	packService = service
}

// Health check the status of the resources of this service
func Health(w http.ResponseWriter, r *http.Request) {
	// close the body buffer at the end of the function
	defer r.Body.Close()

	health := healthservice.Health()
	// If all is ok, marshal into JSON, write headers and content
	respondWithJSON(w, http.StatusOK, health)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(response)))
	w.WriteHeader(code)
	w.Write(response)
}

// SetHealthService sets the health service for this handler
func SetHealthService(health service.IHealthService) {
	healthservice = health
}
