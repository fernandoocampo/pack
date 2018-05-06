// dao package for mongo functions
package dao

import (
	"errors"
	"fmt"
	"time"

	"github.com/fernandoocampo/pack/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// mongoDB is the mongo database name
const mongoDB = "amphora"

// mongoColl is the mongo collection name
const mongoColl = "packs"

// MongoDAO struct for mongo connection
type MongoDAO struct {
}

// GetByID implements *IPackDAO.GetByID using mongo implementation.
// id goes on hex representation. e.g. 59dc3017e78aab3ad5821c85 .
func (m *MongoDAO) GetByID(id string) (*model.Pack, error) {
	if id == "" || &id == nil {
		return nil, errors.New("Invalid pack id")
	}
	idval := bson.ObjectIdHex(id)
	// make a connection to mongo database
	sessionCopy := newMgoSession()
	defer sessionCopy.Close()

	// References the mongo collection
	c := sessionCopy.DB(mongoDB).C(mongoColl)

	result := model.Pack{}
	err := c.Find(bson.M{"_id": idval}).One(&result)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, nil
		}
		errmsg := "An error finding an pack by id - mongodao"
		log.Errorf("%s : %v\n", errmsg, err)
		return nil, fmt.Errorf(errmsg, err)
	}

	return &result, nil
}

// GetIDByCode implements *IPackDAO.GetIDByCode.
func (m *MongoDAO) GetIDByCode(packcode string) (string, error) {
	if packcode == "" || &packcode == nil {
		return "", errors.New("Invalid pack code")
	}
	// make a connection to mongo database
	sessionCopy := newMgoSession()
	defer sessionCopy.Close()
	// References the mongo collection
	c := sessionCopy.DB(mongoDB).C(mongoColl)

	var result struct {
		ID bson.ObjectId `bson:"_id"`
	}
	err := c.Find(bson.M{"packcode": packcode}).Select(bson.M{"_id": 1}).One(&result)
	if err != nil {
		if err == mgo.ErrNotFound {
			return "", nil
		}
		errmsg := "An error finding a pack by packcode - mongodao"
		log.Errorf("%s : %v\n", errmsg, err)
		return "", fmt.Errorf(errmsg, err)
	}

	return result.ID.Hex(), nil
}

// GetByPackCode implements *IPackDAO.GetByPackCode.
func (m *MongoDAO) GetByPackCode(packcode string) (*model.Pack, error) {
	if packcode == "" || &packcode == nil {
		return nil, errors.New("Invalid pack code")
	}
	// make a connection to mongo database
	sessionCopy := newMgoSession()
	defer sessionCopy.Close()
	// References the mongo collection
	c := sessionCopy.DB(mongoDB).C(mongoColl)

	result := model.Pack{}
	err := c.Find(bson.M{"packcode": packcode}).One(&result)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, nil
		}
		errmsg := "An error finding a pack by pack code - mongodao"
		log.Errorf("%s : %v\n", errmsg, err)
		return nil, fmt.Errorf(errmsg, err)
	}

	return &result, nil
}

// GetByProductID implements *IPackDAO.GetByProductId.
func (m *MongoDAO) GetByProductID(productid string) (*model.Pack, error) {
	if productid == "" || &productid == nil {
		return nil, errors.New("Invalid pack internal product id")
	}
	// make a connection to mongo database
	sessionCopy := newMgoSession()
	defer sessionCopy.Close()
	// References the mongo collection
	c := sessionCopy.DB(mongoDB).C(mongoColl)
	result := model.Pack{}
	err := c.Find(bson.M{"prodid": productid}).One(&result)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, nil
		}
		errmsg := "An error finding a pack by prodid - mongodao"
		log.Errorf("%s : %v\n", errmsg, err)
		return nil, fmt.Errorf(errmsg, err)
	}

	return &result, nil
}

// IsThereThisPack implements *IPackDAO.IsThereThisPack.
func (m *MongoDAO) IsThereThisPack(pack *model.PackExists) (bool, error) {
	if pack == nil || (pack.Packcode == "" && pack.ProdID == "") {
		return false, errors.New("Invalid pack data for check if pack exists")
	}
	// make a connection to mongo database
	sessionCopy := newMgoSession()
	defer sessionCopy.Close()
	// References the mongo collection
	c := sessionCopy.DB(mongoDB).C(mongoColl)

	var orfilter bson.M
	if pack.Packcode != "" && pack.ProdID != "" {
		orfilter = bson.M{"$or": []bson.M{bson.M{"prodid": pack.ProdID}, bson.M{"packcode": pack.Packcode}}}
	} else {
		if pack.Packcode != "" {
			orfilter = bson.M{"packcode": pack.Packcode}
		} else {
			orfilter = bson.M{"prodid": pack.ProdID}
		}
	}
	andfilter := bson.M{"$and": []bson.M{bson.M{"mno.id": pack.MnoID}, orfilter}}

	var result struct {
		ID bson.ObjectId `bson:"_id"`
	}
	err := c.Find(andfilter).Select(bson.M{"_id": 1}).One(&result)

	if err != nil {
		if err == mgo.ErrNotFound {
			return false, nil
		}
		errmsg := "An error on IsThereThisPack - mongodao"
		log.Errorf("%s : %v\n", errmsg, err)
		return false, fmt.Errorf(errmsg, err)
	}

	return true, nil
}

// Create implements *IPackDAO.Create.
func (m *MongoDAO) Create(packdata *model.Pack) error {
	if &packdata == nil {
		return errors.New("Invalid pack data")
	}
	// make a connection to mongo database
	sessionCopy := newMgoSession()
	defer sessionCopy.Close()
	// References the mongo collection
	c := sessionCopy.DB(mongoDB).C(mongoColl)
	err := c.Insert(&packdata)

	if err != nil {
		errmsg := "An error on pack creation function - mongodao"
		log.Errorf("%s : %v\n", errmsg, err)
		return fmt.Errorf(errmsg, err)
	}

	return nil
}

// ChangeState implements *IPackDAO.ChangeState.
func (m *MongoDAO) ChangeState(id string, newstate model.PackState) error {
	if id == "" {
		return errors.New("Invalid pack id data")
	}
	// create update json map
	change := bson.M{"$set": bson.M{"state": newstate, "updated": time.Now()}}
	err := updateDataByID(id, change)

	if err != nil {
		errmsg := "An error updating a pack state - mongodao"
		log.Errorf("%s : %v\n", errmsg, err)
		return fmt.Errorf(errmsg, err)
	}
	return nil
}

// ChangeProductID implements *IPackDAO.ChangeProductID.
func (m *MongoDAO) ChangeProductID(id string, newprodid string) error {
	if id == "" || newprodid == "" {
		return errors.New("Invalid pack id and product id data")
	}
	// create update json map
	change := bson.M{"$set": bson.M{"prodid": newprodid, "updated": time.Now()}}
	err := updateDataByID(id, change)

	if err != nil {
		errmsg := "An error updating a pack prodid - mongodao"
		log.Errorf("%s : %v\n", errmsg, err)
		return fmt.Errorf(errmsg, err)
	}
	return nil
}

// ChangePackCode implements *IPackDAO.ChangePackCode.
func (m *MongoDAO) ChangePackCode(id string, newpackcode string) error {
	if id == "" || newpackcode == "" {
		return errors.New("Invalid pack id and pack code data")
	}
	// create update json map
	change := bson.M{"$set": bson.M{"packcode": newpackcode, "updated": time.Now()}}
	err := updateDataByID(id, change)

	if err != nil {
		errmsg := "An error updating a packcode - mongodao"
		log.Errorf("%s : %v\n", errmsg, err)
		return fmt.Errorf(errmsg, err)
	}
	return nil
}

// ChangeName implements *IPackDAO.ChangeName.
func (m *MongoDAO) ChangeName(id string, newname string) error {
	if id == "" || newname == "" {
		return errors.New("Invalid pack id and pack name data")
	}
	// create update json map
	change := bson.M{"$set": bson.M{"name": newname, "updated": time.Now()}}
	err := updateDataByID(id, change)

	if err != nil {
		errmsg := "An error updating a pack name - mongodao"
		log.Errorf("%s : %v\n", errmsg, err)
		return fmt.Errorf(errmsg, err)
	}
	return nil
}

// ChangeDesc implements *IPackDAO.ChangeDesc.
func (m *MongoDAO) ChangeDesc(id string, newdesc string) error {
	if id == "" || newdesc == "" {
		return errors.New("Invalid pack id and pack code data")
	}
	// create update json map
	change := bson.M{"$set": bson.M{"desc": newdesc, "updated": time.Now()}}
	err := updateDataByID(id, change)

	if err != nil {
		errmsg := "An error updating a pack description - mongodao"
		log.Errorf("%s : %v\n", errmsg, err)
		return fmt.Errorf(errmsg, err)
	}
	return nil
}

// ChangeImg implements *IPackDAO.ChangeImg.
func (m *MongoDAO) ChangeImg(id string, newimgurl string) error {
	if id == "" || newimgurl == "" {
		return errors.New("Invalid pack id and pack image url data")
	}
	// create update json map
	change := bson.M{"$set": bson.M{"imgurl": newimgurl, "updated": time.Now()}}
	err := updateDataByID(id, change)

	if err != nil {
		errmsg := "An error updating a pack image url - mongodao"
		log.Errorf("%s : %v\n", errmsg, err)
		return fmt.Errorf(errmsg, err)
	}
	return nil
}

// ChangeKeyword implements *IPackDAO.ChangeKeyword.
func (m *MongoDAO) ChangeKeyword(id string, newkeyword string) error {
	if id == "" || newkeyword == "" {
		return errors.New("Invalid pack id and pack keyword data")
	}
	// create update json map
	change := bson.M{"$set": bson.M{"kwds": newkeyword, "updated": time.Now()}}
	err := updateDataByID(id, change)

	if err != nil {
		errmsg := "An error updating a pack keyword - mongodao"
		log.Errorf("%s : %v\n", errmsg, err)
		return fmt.Errorf(errmsg, err)
	}
	return nil
}

// ChangePrice implements *IPackDAO.ChangePrice.
func (m *MongoDAO) ChangePrice(id string, newprice int) error {
	if id == "" || newprice < 0 {
		return errors.New("Invalid pack id and pack price data")
	}
	// create update json map
	change := bson.M{"$set": bson.M{"price": newprice, "updated": time.Now()}}
	err := updateDataByID(id, change)

	if err != nil {
		errmsg := "An error updating a pack type - mongodao"
		log.Errorf("%s : %v\n", errmsg, err)
		return fmt.Errorf(errmsg, err)
	}
	return nil
}

// ChangePackType implements *IPackDAO.ChangePackType.
func (m *MongoDAO) ChangePackType(id string, newtype *model.Type) error {
	if id == "" || newtype == nil || newtype.ID < 1 || newtype.Name == "" {
		return errors.New("Invalid pack id and pack type data")
	}
	// create update json map
	change := bson.M{"$set": bson.M{"type.id": newtype.ID,
		"type.name": newtype.Name, "updated": time.Now()}}
	err := updateDataByID(id, change)

	if err != nil {
		errmsg := "An error updating a pack mno - mongodao"
		log.Errorf("%s : %v\n", errmsg, err)
		return fmt.Errorf(errmsg, err)
	}
	return nil
}

// ChangeMNO implements *IPackDAO.ChangeMNO.
func (m *MongoDAO) ChangeMNO(id string, newmno *model.Mno) error {
	if id == "" || newmno == nil || newmno.ID < 1 || newmno.Name == "" {
		return errors.New("Invalid pack id and pack mno data")
	}
	// create update json map
	change := bson.M{"$set": bson.M{"mno.id": newmno.ID,
		"mno.name": newmno.Name, "updated": time.Now()}}
	err := updateDataByID(id, change)

	if err != nil {
		errmsg := "An error updating a pack mno - mongodao"
		log.Errorf("%s : %v\n", errmsg, err)
		return fmt.Errorf(errmsg, err)
	}
	return nil
}

// ChangeValidity implements *IPackDAO.ChangeValidity.
func (m *MongoDAO) ChangeValidity(id string, newterm *model.Term) error {
	if id == "" || newterm == nil || newterm.UnitID < 1 || newterm.Unit == "" {
		return errors.New("Invalid pack id and pack validity data")
	}
	// create update json map
	change := bson.M{"$set": bson.M{"term.unit_id": newterm.UnitID,
		"term.unit": newterm.Unit, "term.amount": newterm.Amount,
		"updated": time.Now()}}
	err := updateDataByID(id, change)

	if err != nil {
		errmsg := "An error updating a pack validity - mongodao"
		log.Errorf("%s : %v\n", errmsg, err)
		return fmt.Errorf(errmsg, err)
	}
	return nil
}

// ChangeCurrency implements *IPackDAO.ChangeCurrency.
func (m *MongoDAO) ChangeCurrency(id string, newccy *model.Currency) error {
	if id == "" || newccy == nil || newccy.ID < 1 || newccy.Name == "" {
		return errors.New("Invalid pack id and pack price currency data")
	}
	// create update json map
	change := bson.M{"$set": bson.M{"currency.id": newccy.ID,
		"currency.name": newccy.Name, "updated": time.Now()}}
	err := updateDataByID(id, change)

	if err != nil {
		errmsg := "An error updating a pack currency - mongodao"
		log.Errorf("%s : %v\n", errmsg, err)
		return fmt.Errorf(errmsg, err)
	}
	return nil
}

// ChangeStock implements IPackDAO.ChangeStock using mongo driver.
func (m *MongoDAO) ChangeStock(id string, amount int) error {
	if id == "" || amount == 0 {
		return nil
	}

	// increase or descrease update json map
	change := bson.M{"$inc": bson.M{"stock": amount}}
	err := updateDataByID(id, change)

	if err != nil {
		errmsg := "An error changing pack: %s stock: %d - mongodao: %s\n"
		log.Errorf(errmsg, id, amount, err.Error())
		return fmt.Errorf(errmsg, err)
	}

	return nil
}

// UpdateResources implements IPackDAO.UpdateResources.
func (m *MongoDAO) UpdateResources(id string, newresources []model.Resource) error {
	if id == "" {
		return errors.New("Invalid pack id")
	}

	var resources []model.Resource

	if newresources == nil {
		resources = []model.Resource{}
	} else {
		resources = newresources
	}
	// create update json map
	change := bson.M{"$set": bson.M{"resources": resources, "updated": time.Now()}}
	err := updateDataByID(id, change)

	if err != nil {
		errmsg := fmt.Sprintf("An error updating a pack resources - mongodao: %v", err)
		log.Errorf(errmsg)
		return err
	}
	return nil
}

// Delete implements *IPackDAO.Delete.
func (m *MongoDAO) Delete(id string) error {
	if id == "" {
		return errors.New("pack id is mandatory")
	}
	// make a connection to mongo database
	sessionCopy := newMgoSession()
	defer sessionCopy.Close()
	// References the mongo collection
	c := sessionCopy.DB(mongoDB).C(mongoColl)

	bsonid := bson.ObjectIdHex(id)
	err := c.Remove(bson.M{"_id": bsonid})

	if err != nil {
		errmsg := "An error deleting a pack - mongodao"
		log.Errorf("%s : %v\n", errmsg, err)
		return fmt.Errorf(errmsg, err)
	}

	return nil
}

// updateDataById update pack with the given parameter map
// that contains the data to update. Returns error if something
// goes wrong. id must be hex representation.
func updateDataByID(id string, change interface{}) error {

	// make a connection to mongo database
	sessionCopy := newMgoSession()
	defer sessionCopy.Close()

	// query the user in the database
	c := sessionCopy.DB(mongoDB).C(mongoColl)

	//convert the ID to bson.ObjectID
	bsonid := bson.ObjectIdHex(id)

	//Here the filter is formed
	colQuerier := bson.M{"_id": bsonid}

	//Here the update is perform
	err := c.Update(colQuerier, change)

	return err

}
