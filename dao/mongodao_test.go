package dao_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/fernandoocampo/pack/dao"
	"github.com/fernandoocampo/pack/model"
)

// TestGetByID search a pack with the given id
// and return it.
func TestGetByID(t *testing.T) {
	// GIVEN data of a pack created
	dao.SetDBname("amphora")
	dao.SetMongoAddrs([]string{"localhost:27017"})
	dao.SetTimeout(60)

	dao.InitMgoSession()
	defer dao.CloseMgoSession()

	mongodao := new(dao.MongoDAO)

	strcode := "wh13"
	newpack := newPackData(strcode, "Whatsapp weekend 13", "13")

	err1 := mongodao.Create(newpack)

	if err1 != nil {
		t.Fatalf("Expected err1 to be nil but it was: %s", err1)
	}

	packbycode, err2 := mongodao.GetByPackCode(strcode)

	if err2 != nil {
		t.Fatalf("Expected err2 to be nil but it was: %s", err2)
	}

	// fmt.Printf("bson.ID: %+v <-> %s\n", packdata.ID, packdata.ID)

	// WHEN we need to search a pack by id
	packbyid, err3 := mongodao.GetByID(packbycode.ID.Hex())

	// THEN we check that pack was query by id
	if err3 != nil {
		t.Fatalf("Expected err3 to be nil but it was: %s", err3)
	}

	if packbyid.Name == "" {
		t.Fatalf("Expected pack.Name to be not nil but it was nil")
	}

	if packbyid.Packcode != strcode {
		t.Fatalf("Expected pack.code equals to be: %s, but it was: %s\n", strcode, packbyid.Packcode)
	}

}

// TestGetIDByCode searches a pack with the given
// pack code and return it.
func TestGetIDByCode(t *testing.T) {
	// GIVEN data of a pack created
	dao.SetDBname("amphora")
	dao.SetMongoAddrs([]string{"localhost:27017"})
	dao.SetTimeout(60)

	dao.InitMgoSession()
	defer dao.CloseMgoSession()

	mongodao := new(dao.MongoDAO)

	strcode := "wh14"
	newpack := newPackData(strcode, "Whatsapp weekend 14", "14")

	err1 := mongodao.Create(newpack)

	if err1 != nil {
		t.Fatalf("Expected err1 to be nil but it was: %s", err1)
	}

	// WHEN we need to query its ID by pack code
	packid, err2 := mongodao.GetIDByCode(strcode)

	// THEN we check that the pack exists and we get the id
	if err2 != nil {
		t.Fatalf("Expected err2 to be nil but it was: %s", err2)
	}

	if packid == "" {
		t.Fatalf("Expected pack id to be not nil but it was nil")
	}
}

// TestGetByPackCode search a pack with the given shortcode
// and return it.
func TestGetByPackCode(t *testing.T) {
	// GIVEN data of a pack created
	dao.SetDBname("amphora")
	dao.SetMongoAddrs([]string{"localhost:27017"})
	dao.SetTimeout(60)

	dao.InitMgoSession()
	defer dao.CloseMgoSession()

	mongodao := new(dao.MongoDAO)

	strcode := "wh15"
	newpack := newPackData(strcode, "Whatsapp weekend 15", "15")

	err1 := mongodao.Create(newpack)

	if err1 != nil {
		t.Fatalf("Expected err1 to be nil but it was: %s", err1)
	}

	// WHEN we need to query a pack by its pack code
	packbycode, err2 := mongodao.GetByPackCode(strcode)

	// THEN we check that there is a pack with the the given pack code
	if err2 != nil {
		t.Fatalf("Expected err2 to be nil but it was: %s", err2)
	}

	if packbycode.Name == "" {
		t.Fatalf("Expected pack name to be not nil but it was nil")
	}
}

// TestGetByProductId search a pack with the given mno internal id
// and return it.
func TestGetByProductID(t *testing.T) {
	// GIVEN data of a pack created
	dao.SetDBname("amphora")
	dao.SetMongoAddrs([]string{"localhost:27017"})
	dao.SetTimeout(60)

	dao.InitMgoSession()
	defer dao.CloseMgoSession()

	mongodao := new(dao.MongoDAO)

	strprodid := "16"
	newpack := newPackData("wkd16", "Whatsapp weekend 16", "16")

	err1 := mongodao.Create(newpack)

	if err1 != nil {
		t.Fatalf("Expected err1 to be nil but it was: %s", err1)
	}

	// WHEN we need to query a pack by its product id
	packbyprodid, err2 := mongodao.GetByProductID(strprodid)

	// THEN we check that there is a pack with the the given pack code
	if err2 != nil {
		t.Fatalf("Expected err2 to be nil but it was: %s", err2)
	}

	if packbyprodid.Name == "" {
		t.Fatalf("Expected pack name to be not nil but it was nil")
	}
}

// TestCreate inserts a new Pack in the system. Returns
// true if the Pack is created
func TestCreate(t *testing.T) {
	// GIVEN data to create a new pack
	dao.SetDBname("amphora")
	dao.SetMongoAddrs([]string{"localhost:27017"})
	dao.SetTimeout(60)

	dao.InitMgoSession()
	defer dao.CloseMgoSession()

	mongodao := new(dao.MongoDAO)

	newpack := newPackData("wh12", "Whatsapp weekend", "12")

	// WHEN we create a pack
	err := mongodao.Create(newpack)

	// THEN we check if the pack creation was correct.
	if err != nil {
		t.Fatalf("Expected err to be nil but it was: %s", err)
	}

	pack, err := mongodao.GetByPackCode("wh12")

	if err != nil {
		t.Fatalf("Expected err to be nil but it was: %s", err)
	}

	if pack.Name == "" {
		t.Fatalf("Expected pack.Name to be not nil but it was nil")
	}
}

// TestChangeState verify the changes the state of a pack.
func TestChangeState(t *testing.T) {
	// GIVEN data of a pack created
	dao.SetDBname("amphora")
	dao.SetMongoAddrs([]string{"localhost:27017"})
	dao.SetTimeout(60)

	dao.InitMgoSession()
	defer dao.CloseMgoSession()

	mongodao := new(dao.MongoDAO)

	strcode := "wh17"
	newpack := newPackData(strcode, "Whatsapp weekend 17", "17")

	err1 := mongodao.Create(newpack)

	if err1 != nil {
		t.Fatalf("Expected err1 to be nil but it was: %s", err1)
	}

	packbycode, err2 := mongodao.GetByPackCode(strcode)

	if err2 != nil {
		t.Fatalf("Expected err2 to be nil but it was: %s", err2)
	}

	// WHEN we need to change the state of a pack
	err3 := mongodao.ChangeState(packbycode.ID.Hex(), model.Active)

	// THEN we check that we can modify the state of a pack
	if err3 != nil {
		t.Fatalf("Expected err3 to be nil but it was: %s", err3)
	}

	packbycode2, err4 := mongodao.GetByPackCode(strcode)

	if err4 != nil {
		t.Fatalf("Expected err4 to be nil but it was: %s", err4)
	}

	if packbycode2.State != model.Active {
		t.Fatalf("Expected pack state to be %d but it was %d\n", model.Active, packbycode2.State)
	}
}

// TestChangeProductID verify the changes the mno internal product id.
func TestChangeProductID(t *testing.T) {
	// GIVEN data of a pack created
	dao.SetDBname("amphora")
	dao.SetMongoAddrs([]string{"localhost:27017"})
	dao.SetTimeout(60)

	dao.InitMgoSession()
	defer dao.CloseMgoSession()

	mongodao := new(dao.MongoDAO)

	strcode := "wh18"
	newpack := newPackData(strcode, "Whatsapp weekend 18", "18")

	err1 := mongodao.Create(newpack)

	if err1 != nil {
		t.Fatalf("Expected err1 to be nil but it was: %s", err1)
	}

	packbycode, err2 := mongodao.GetByPackCode(strcode)

	if err2 != nil {
		t.Fatalf("Expected err2 to be nil but it was: %s", err2)
	}

	// WHEN we need to change the product id of a pack
	newprodid := "18_1"
	err3 := mongodao.ChangeProductID(packbycode.ID.Hex(), newprodid)

	// THEN we check that we can modify the state of a pack
	if err3 != nil {
		t.Fatalf("Expected err3 to be nil but it was: %s", err3)
	}

	packbycode2, err4 := mongodao.GetByPackCode(strcode)

	if err4 != nil {
		t.Fatalf("Expected err4 to be nil but it was: %s", err4)
	}

	if packbycode2.ProdID != newprodid {
		t.Fatalf("Expected pack prodid to be %s but it was %s\n", newprodid, packbycode2.ProdID)
	}
}

// TestChangePackCode verify the changes the pack short code
func TestChangePackCode(t *testing.T) {
	// GIVEN data of a pack created
	dao.SetDBname("amphora")
	dao.SetMongoAddrs([]string{"localhost:27017"})
	dao.SetTimeout(60)

	dao.InitMgoSession()
	defer dao.CloseMgoSession()

	mongodao := new(dao.MongoDAO)

	strcode := "wh19"
	newpack := newPackData(strcode, "Whatsapp weekend 19", "19")

	err1 := mongodao.Create(newpack)

	if err1 != nil {
		t.Fatalf("Expected err1 to be nil but it was: %s", err1)
	}

	packbycode, err2 := mongodao.GetByPackCode(strcode)

	if err2 != nil {
		t.Fatalf("Expected err2 to be nil but it was: %s", err2)
	}

	// WHEN we need to change the product id of a pack
	newpackcode := "wh19_1"
	err3 := mongodao.ChangePackCode(packbycode.ID.Hex(), newpackcode)

	// THEN we check that we can modify the state of a pack
	if err3 != nil {
		t.Fatalf("Expected err3 to be nil but it was: %s", err3)
	}

	packbycode2, err4 := mongodao.GetByPackCode(newpackcode)

	if err4 != nil {
		t.Fatalf("Expected err4 to be nil but it was: %s", err4)
	}

	if packbycode2.Packcode != newpackcode {
		t.Fatalf("Expected pack prodid to be %s but it was %s\n", newpackcode, packbycode2.Packcode)
	}
}

// TestChangeName verify the changes the name of an existent pack.
func TestChangeName(t *testing.T) {
	// GIVEN data of a pack created
	dao.SetDBname("amphora")
	dao.SetMongoAddrs([]string{"localhost:27017"})
	dao.SetTimeout(60)

	dao.InitMgoSession()
	defer dao.CloseMgoSession()

	mongodao := new(dao.MongoDAO)

	strcode := "wh20"
	strname := "Whatsapp weekend 20"
	newpack := newPackData(strcode, strname, "20")

	err1 := mongodao.Create(newpack)

	if err1 != nil {
		t.Fatalf("Expected err1 to be nil but it was: %s", err1)
	}

	packbycode, err2 := mongodao.GetByPackCode(strcode)

	if err2 != nil {
		t.Fatalf("Expected err2 to be nil but it was: %s", err2)
	}

	// WHEN we need to change the product id of a pack
	newname := "Whatsapp weekend 20-1"
	err3 := mongodao.ChangeName(packbycode.ID.Hex(), newname)

	// THEN we check that we can modify the state of a pack
	if err3 != nil {
		t.Fatalf("Expected err3 to be nil but it was: %s", err3)
	}

	packbycode2, err4 := mongodao.GetByPackCode(strcode)

	if err4 != nil {
		t.Fatalf("Expected err4 to be nil but it was: %s", err4)
	}

	if packbycode2.Name == strname {
		t.Fatalf("Expected pack name to be %s but it was %s\n", newname, strname)
	}
}

// TestChangeDesc verify the changes the description of an existent pack.
func TestChangeDesc(t *testing.T) {
	// GIVEN data of a pack created
	dao.SetDBname("amphora")
	dao.SetMongoAddrs([]string{"localhost:27017"})
	dao.SetTimeout(60)

	dao.InitMgoSession()
	defer dao.CloseMgoSession()

	mongodao := new(dao.MongoDAO)

	strcode := "wh21"
	strname := "Whatsapp weekend 21"
	strdesc := "Whatsapp para el fin de semana, para que hables con tus amigos todo el dia."
	newpack := newPackData(strcode, strname, "21")

	err1 := mongodao.Create(newpack)

	if err1 != nil {
		t.Fatalf("Expected err1 to be nil but it was: %s", err1)
	}

	packbycode, err2 := mongodao.GetByPackCode(strcode)

	if err2 != nil {
		t.Fatalf("Expected err2 to be nil but it was: %s", err2)
	}

	// WHEN we need to change the product id of a pack
	newdesc := "Whatsapp para que hables con tus amigos todo el dia."
	err3 := mongodao.ChangeDesc(packbycode.ID.Hex(), newdesc)

	// THEN we check that we can modify the state of a pack
	if err3 != nil {
		t.Fatalf("Expected err3 to be nil but it was: %s", err3)
	}

	packbycode2, err4 := mongodao.GetByPackCode(strcode)

	if err4 != nil {
		t.Fatalf("Expected err4 to be nil but it was: %s", err4)
	}

	if packbycode2.Desc == strdesc {
		t.Fatalf("Expected pack desc to be %s but it was %s\n", newdesc, strdesc)
	}
}

// TestChangeImg verify the changes the image url of the given pack.
func TestChangeImg(t *testing.T) {
	// GIVEN data of a pack created
	dao.SetDBname("amphora")
	dao.SetMongoAddrs([]string{"localhost:27017"})
	dao.SetTimeout(60)

	dao.InitMgoSession()
	defer dao.CloseMgoSession()

	mongodao := new(dao.MongoDAO)

	strcode := "wh22"
	strname := "Whatsapp weekend 22"
	strimg := "/appdata/img/whatsappwknd.png"
	newpack := newPackData(strcode, strname, "22")

	err1 := mongodao.Create(newpack)

	if err1 != nil {
		t.Fatalf("Expected err1 to be nil but it was: %s", err1)
	}

	packbycode, err2 := mongodao.GetByPackCode(strcode)

	if err2 != nil {
		t.Fatalf("Expected err2 to be nil but it was: %s", err2)
	}

	// WHEN we need to change the product id of a pack
	newimg := "/appdata/img/whatsappwknd2.png"
	err3 := mongodao.ChangeImg(packbycode.ID.Hex(), newimg)

	// THEN we check that we can modify the state of a pack
	if err3 != nil {
		t.Fatalf("Expected err3 to be nil but it was: %s", err3)
	}

	packbycode2, err4 := mongodao.GetByPackCode(strcode)

	if err4 != nil {
		t.Fatalf("Expected err4 to be nil but it was: %s", err4)
	}

	if packbycode2.Img == strimg {
		t.Fatalf("Expected pack image url to be %s but it was %s\n", newimg, strimg)
	}
}

// TestChangeKeyword verify the changes the key words of the given pack.
func TestChangekeyword(t *testing.T) {
	// GIVEN data of a pack created
	dao.SetDBname("amphora")
	dao.SetMongoAddrs([]string{"localhost:27017"})
	dao.SetTimeout(60)

	dao.InitMgoSession()
	defer dao.CloseMgoSession()

	mongodao := new(dao.MongoDAO)

	strcode := "wh23"
	strname := "Whatsapp weekend 23"
	strkwd := "internet whatsapp dia fin semana"
	newpack := newPackData(strcode, strname, "23")

	err1 := mongodao.Create(newpack)

	if err1 != nil {
		t.Fatalf("Expected err1 to be nil but it was: %s", err1)
	}

	packbycode, err2 := mongodao.GetByPackCode(strcode)

	if err2 != nil {
		t.Fatalf("Expected err2 to be nil but it was: %s", err2)
	}

	// WHEN we need to change the product id of a pack
	newkwd := "internet whatsapp dia fin semana futuro"
	err3 := mongodao.ChangeKeyword(packbycode.ID.Hex(), newkwd)

	// THEN we check that we can modify the state of a pack
	if err3 != nil {
		t.Fatalf("Expected err3 to be nil but it was: %s", err3)
	}

	packbycode2, err4 := mongodao.GetByPackCode(strcode)

	if err4 != nil {
		t.Fatalf("Expected err4 to be nil but it was: %s", err4)
	}

	if packbycode2.Kwds == strkwd {
		t.Fatalf("Expected pack image url to be %s but it was %s\n", newkwd, strkwd)
	}
}

// TestChangePrice verify the changes the price of an existent pack.
func TestChangePrice(t *testing.T) {
	// GIVEN data of a pack created
	dao.SetDBname("amphora")
	dao.SetMongoAddrs([]string{"localhost:27017"})
	dao.SetTimeout(60)

	dao.InitMgoSession()
	defer dao.CloseMgoSession()

	mongodao := new(dao.MongoDAO)

	strcode := "wh24"
	strname := "Whatsapp weekend 24"
	strprice := 2500
	newpack := newPackData(strcode, strname, "24")

	err1 := mongodao.Create(newpack)

	if err1 != nil {
		t.Fatalf("Expected err1 to be nil but it was: %s", err1)
	}

	packbycode, err2 := mongodao.GetByPackCode(strcode)

	if err2 != nil {
		t.Fatalf("Expected err2 to be nil but it was: %s", err2)
	}

	// WHEN we need to change the product id of a pack
	newprice := 3500
	err3 := mongodao.ChangePrice(packbycode.ID.Hex(), newprice)

	// THEN we check that we can modify the state of a pack
	if err3 != nil {
		t.Fatalf("Expected err3 to be nil but it was: %s", err3)
	}

	packbycode2, err4 := mongodao.GetByPackCode(strcode)

	if err4 != nil {
		t.Fatalf("Expected err4 to be nil but it was: %s", err4)
	}

	if packbycode2.Price == strprice {
		t.Fatalf("Expected pack price to be %d but it was %d\n", newprice, strprice)
	}
}

// TestChangePackType verify the changes the pack type of an existent pack
func TestChangePackType(t *testing.T) {
	// GIVEN data of a pack created
	dao.SetDBname("amphora")
	dao.SetMongoAddrs([]string{"localhost:27017"})
	dao.SetTimeout(60)

	dao.InitMgoSession()
	defer dao.CloseMgoSession()

	mongodao := new(dao.MongoDAO)

	strcode := "wh25"
	strname := "Whatsapp weekend 25"
	complex := model.Type{
		ID:   1,
		Name: "App",
	}
	newpack := newPackData(strcode, strname, "25")

	err1 := mongodao.Create(newpack)

	if err1 != nil {
		t.Fatalf("Expected err1 to be nil but it was: %s", err1)
	}

	packbycode, err2 := mongodao.GetByPackCode(strcode)

	if err2 != nil {
		t.Fatalf("Expected err2 to be nil but it was: %s", err2)
	}

	// WHEN we need to change the type of a pack
	newcomplex := model.Type{
		ID:   2,
		Name: "App2",
	}
	err3 := mongodao.ChangePackType(packbycode.ID.Hex(), &newcomplex)

	// THEN we check that we can modify the state of a pack
	if err3 != nil {
		t.Fatalf("Expected err3 to be nil but it was: %s", err3)
	}

	packbycode2, err4 := mongodao.GetByPackCode(strcode)

	if err4 != nil {
		t.Fatalf("Expected err4 to be nil but it was: %s", err4)
	}

	if packbycode2.Packtype == nil {
		t.Fatalf("Expected pack to be not nill and it was nil")
	}

	if packbycode2.Packtype.ID == complex.ID {
		t.Fatalf("Expected pack type with ID %d but it was %d\n", newcomplex.ID, packbycode2.Packtype.ID)
	}

	if packbycode2.Packtype.Name == complex.Name {
		t.Fatalf("Expected pack type with Name %s but it was %s\n", newcomplex.Name, packbycode2.Packtype.Name)
	}
}

// TestChangeMNO verify the changes the Mobile Network Operator owner of the pack.
func TestChangeMNO(t *testing.T) {
	// GIVEN data of a pack created
	dao.SetDBname("amphora")
	dao.SetMongoAddrs([]string{"localhost:27017"})
	dao.SetTimeout(60)

	dao.InitMgoSession()
	defer dao.CloseMgoSession()

	mongodao := new(dao.MongoDAO)

	strcode := "wh26"
	strname := "Whatsapp weekend 26"
	complex := model.Mno{
		ID:   2,
		Name: "Claro",
	}
	newpack := newPackData(strcode, strname, "26")

	err1 := mongodao.Create(newpack)

	if err1 != nil {
		t.Fatalf("Expected err1 to be nil but it was: %s", err1)
	}

	packbycode, err2 := mongodao.GetByPackCode(strcode)

	if err2 != nil {
		t.Fatalf("Expected err2 to be nil but it was: %s", err2)
	}

	// WHEN we need to change the type of a pack
	newcomplex := model.Mno{
		ID:   5,
		Name: "Uff",
	}
	err3 := mongodao.ChangeMNO(packbycode.ID.Hex(), &newcomplex)

	// THEN we check that we can modify the state of a pack
	if err3 != nil {
		t.Fatalf("Expected err3 to be nil but it was: %s", err3)
	}

	packbycode2, err4 := mongodao.GetByPackCode(strcode)

	if err4 != nil {
		t.Fatalf("Expected err4 to be nil but it was: %s", err4)
	}

	if packbycode2.Mno == nil {
		t.Fatalf("Expected pack mno to be not nill and it was nil")
	}

	if packbycode2.Mno.ID == complex.ID {
		t.Fatalf("Expected pack mno with ID %d but it was %d\n", newcomplex.ID, packbycode2.Mno.ID)
	}

	if packbycode2.Mno.Name == complex.Name {
		t.Fatalf("Expected pack mno with Name %s but it was %s\n", newcomplex.Name, packbycode2.Mno.Name)
	}
}

// TestChangeValidity verify the changes the validity of the given pack.
func TestChangeValidity(t *testing.T) {
	// GIVEN data of a pack created
	dao.SetDBname("amphora")
	dao.SetMongoAddrs([]string{"localhost:27017"})
	dao.SetTimeout(60)

	dao.InitMgoSession()
	defer dao.CloseMgoSession()

	mongodao := new(dao.MongoDAO)

	strcode := "wh27"
	strname := "Whatsapp weekend 27"
	complex := model.Term{
		UnitID: 4,
		Unit:   "dia",
		Amount: 2,
	}
	newpack := newPackData(strcode, strname, "27")

	err1 := mongodao.Create(newpack)

	if err1 != nil {
		t.Fatalf("Expected err1 to be nil but it was: %s", err1)
	}

	packbycode, err2 := mongodao.GetByPackCode(strcode)

	if err2 != nil {
		t.Fatalf("Expected err2 to be nil but it was: %s", err2)
	}

	// WHEN we need to change the type of a pack
	newcomplex := model.Term{
		UnitID: 5,
		Unit:   "week",
		Amount: 3,
	}
	err3 := mongodao.ChangeValidity(packbycode.ID.Hex(), &newcomplex)

	// THEN we check that we can modify the state of a pack
	if err3 != nil {
		t.Fatalf("Expected err3 to be nil but it was: %s", err3)
	}

	packbycode2, err4 := mongodao.GetByPackCode(strcode)

	if err4 != nil {
		t.Fatalf("Expected err4 to be nil but it was: %s", err4)
	}

	if packbycode2.Term == nil {
		t.Fatalf("Expected pack mno to be not nill and it was nil")
	}

	if packbycode2.Term.UnitID == complex.UnitID {
		t.Fatalf("Expected pack term with unitID %d but it was %d\n", newcomplex.UnitID, packbycode2.Term.UnitID)
	}

	if packbycode2.Term.Unit == complex.Unit {
		t.Fatalf("Expected pack term unit %s but it was %s\n", newcomplex.Unit, packbycode2.Term.Unit)
	}

	if packbycode2.Term.Amount == complex.Amount {
		t.Fatalf("Expected pack term with amount %d but it was %d\n", newcomplex.Amount, packbycode2.Term.Amount)
	}
}

// TestChangeCurrency verify the changes the currency of price of the given pack.
func TestChangeCurrency(t *testing.T) {
	// GIVEN data of a pack created
	dao.SetDBname("amphora")
	dao.SetMongoAddrs([]string{"localhost:27017"})
	dao.SetTimeout(60)

	dao.InitMgoSession()
	defer dao.CloseMgoSession()

	mongodao := new(dao.MongoDAO)

	strcode := "wh28"
	strname := "Whatsapp weekend 28"
	complex := model.Currency{
		ID:   3,
		Name: "cop",
	}
	newpack := newPackData(strcode, strname, "28")

	err1 := mongodao.Create(newpack)

	if err1 != nil {
		t.Fatalf("Expected err1 to be nil but it was: %s", err1)
	}

	packbycode, err2 := mongodao.GetByPackCode(strcode)

	if err2 != nil {
		t.Fatalf("Expected err2 to be nil but it was: %s", err2)
	}

	// WHEN we need to change the type of a pack
	newcomplex := model.Currency{
		ID:   2,
		Name: "pe",
	}
	err3 := mongodao.ChangeCurrency(packbycode.ID.Hex(), &newcomplex)

	// THEN we check that we can modify the state of a pack
	if err3 != nil {
		t.Fatalf("Expected err3 to be nil but it was: %s", err3)
	}

	packbycode2, err4 := mongodao.GetByPackCode(strcode)

	if err4 != nil {
		t.Fatalf("Expected err4 to be nil but it was: %s", err4)
	}

	if packbycode2.Ccy == nil {
		t.Fatalf("Expected pack mno to be not nill and it was nil")
	}

	if packbycode2.Ccy.ID == complex.ID {
		t.Fatalf("Expected pack currency with ID %d but it was %d\n", newcomplex.ID, packbycode2.Ccy.ID)
	}

	if packbycode2.Ccy.Name == complex.Name {
		t.Fatalf("Expected pack Currency with Name %s but it was %s\n", newcomplex.Name, packbycode2.Ccy.Name)
	}
}

// TestDelete removes an existent Pack and returns true if the pack can be deleted.
func TestDelete(t *testing.T) {
	// GIVEN data to create a new pack
	dao.SetDBname("amphora")
	dao.SetMongoAddrs([]string{"localhost:27017"})
	dao.SetTimeout(60)

	dao.InitMgoSession()
	defer dao.CloseMgoSession()

	mongodao := new(dao.MongoDAO)

	newpack := newPackData("wh30", "Whatsapp weekend 30", "30")

	err := mongodao.Create(newpack)

	if err != nil {
		t.Fatalf("Expected err to be nil but it was: %s", err)
	}

	pack, err2 := mongodao.GetByPackCode("wh30")

	if err2 != nil {
		t.Fatalf("Expected err to be nil but it was: %s", err2)
	}

	// WHEN we delete  a pack
	err3 := mongodao.Delete(pack.ID.Hex())
	// THEN we check if the pack deletion was correct.
	if err3 != nil {
		t.Fatalf("Expected err to be nil but it was: %s", err3)
	}

	pack2, err2 := mongodao.GetByPackCode("wh30")

	if err2 != nil {
		t.Fatalf("Expected err to be nil but it was: %s", err2)
	}

	if pack2 != nil {
		t.Fatalf("Expected packed was deleted but we have a result %+v\n", pack2)
	}
}

// TestIsThereThisPack verify that works
func TestIsThereThisPack(t *testing.T) {
	// GIVEN data to check if a pack exists
	dao.SetDBname("amphora")
	dao.SetMongoAddrs([]string{"localhost:27017"})
	dao.SetTimeout(60)

	dao.InitMgoSession()
	defer dao.CloseMgoSession()

	mongodao := new(dao.MongoDAO)

	packcode := "wh31"
	productid := "31"
	newpack := newPackData(packcode, "Whatsapp weekend 31", productid)

	err := mongodao.Create(newpack)

	if err != nil {
		t.Fatalf("Expected err to be nil but it was: %s", err)
	}

	// pack 1
	packexists := model.NewPackExists(newpack)
	// pack 2
	packexists2 := new(model.PackExists)
	packexists2.MnoID = 2
	packexists2.Packcode = packcode
	packexists2.ProdID = "131"
	// pack 3
	packexists3 := new(model.PackExists)
	packexists3.MnoID = 2
	packexists3.Packcode = "wh131"
	packexists3.ProdID = productid
	// pack 4
	packexists4 := new(model.PackExists)
	packexists4.MnoID = 100
	packexists4.Packcode = packcode
	packexists4.ProdID = productid
	// pack 5
	packexists5 := new(model.PackExists)
	packexists5.MnoID = 2
	packexists5.Packcode = "wh131"
	packexists5.ProdID = "131"

	// WHEN we check if a pack exists with code, productid for a mno pack owner.
	result, err2 := mongodao.IsThereThisPack(packexists)
	result2, err3 := mongodao.IsThereThisPack(packexists2)
	result3, err4 := mongodao.IsThereThisPack(packexists3)
	result4, err5 := mongodao.IsThereThisPack(packexists4)
	result5, err6 := mongodao.IsThereThisPack(packexists5)

	// THEN we check if results are ok
	if err2 != nil {
		t.Fatalf("Expected err to be nil but it was: %s", err2)
	}
	if !result {
		t.Fatalf("Expected result to be true but it was false")
	}

	if err3 != nil {
		t.Fatalf("Expected err to be nil but it was: %s", err2)
	}
	if !result2 {
		t.Fatalf("Expected result to be true but it was false")
	}
	if err4 != nil {
		t.Fatalf("Expected err to be nil but it was: %s", err2)
	}
	if !result3 {
		t.Fatalf("Expected result to be true but it was false")
	}
	if err5 != nil {
		t.Fatalf("Expected err to be nil but it was: %s", err2)
	}
	if result4 {
		t.Fatalf("Expected result to be false but it was true")
	}
	if err6 != nil {
		t.Fatalf("Expected err to be nil but it was: %s", err2)
	}
	if result5 {
		t.Fatalf("Expected result to be false but it was true")
	}

}

// TestChangeStock tests if the movement of stock of a pack is good or not.
func TestChangeStock(t *testing.T) {
	// GIVEN data of a pack created
	dao.SetDBname("amphora")
	dao.SetMongoAddrs([]string{"localhost:27017"})
	dao.SetTimeout(60)

	dao.InitMgoSession()
	defer dao.CloseMgoSession()

	mongodao := new(dao.MongoDAO)

	strcode := "wh32"
	strname := "Whatsapp weekend 32"
	newpack := newPackData(strcode, strname, "32")

	err1 := mongodao.Create(newpack)

	if err1 != nil {
		t.Fatalf("Expected err1 to be nil but it was: %s", err1)
	}

	packbycode, err2 := mongodao.GetByPackCode(strcode)

	if err2 != nil {
		t.Fatalf("Expected err2 to be nil but it was: %s", err2)
	}

	// WHEN we need to change the stock of a pack
	err3 := mongodao.ChangeStock(packbycode.ID.Hex(), 2)
	err4 := mongodao.ChangeStock(packbycode.ID.Hex(), -1)

	// THEN we check that we can modify the state of a pack
	if err3 != nil {
		t.Fatalf("Expected err3 to be nil but it was: %s", err3)
	}
	if err4 != nil {
		t.Fatalf("Expected err3 to be nil but it was: %s", err4)
	}

	packbycode2, err5 := mongodao.GetByPackCode(strcode)

	if err5 != nil {
		t.Fatalf("Expected err4 to be nil but it was: %s", err5)
	}

	if packbycode2.Stock != 1 {
		t.Fatalf("Expected pack stock to be 1 and it was %d", packbycode2.Stock)
	}
}

func newPackData(code string, name string, prodid string) *model.Pack {
	expresult := model.Pack{
		ID:       "",
		ProdID:   prodid,
		Packcode: code,
		Name:     name,
		Desc:     "Whatsapp para el fin de semana, para que hables con tus amigos todo el dia.",
		Img:      "/appdata/img/whatsappwknd.png",
		Kwds:     "internet whatsapp dia fin semana",
		Price:    2500,
		Ownerid:  1,
		Created:  time.Time{},
		Updated:  time.Time{},
		Packtype: &model.Type{
			ID:   1,
			Name: "App",
		},
		Mno: &model.Mno{
			ID:   2,
			Name: "Claro",
		},
		Term: &model.Term{
			UnitID: 4,
			Unit:   "dia",
			Amount: 2,
		},
		Ccy: &model.Currency{
			ID:   3,
			Name: "cop",
		},
		State: model.Inactive,
	}
	return &expresult
}

func TestMongoDAO_UpdateResources(t *testing.T) {
	// create a pack first
	// GIVEN data of a pack created
	dao.SetDBname("amphora")
	dao.SetMongoAddrs([]string{"localhost:27017"})
	dao.SetTimeout(60)

	dao.InitMgoSession()
	defer dao.CloseMgoSession()

	mongodao := new(dao.MongoDAO)

	strcode := "wh33"
	strname := "Whatsapp weekend 33"
	newpack := newPackData(strcode, strname, "33")

	err1 := mongodao.Create(newpack)

	if err1 != nil {
		t.Fatalf("Expected err1 to be nil but it was: %s", err1)
	}

	packid, err2 := mongodao.GetIDByCode(strcode)

	if err2 != nil {
		t.Fatalf("Expected err2 to be nil but it was: %s", err2)
	}

	// WHEN we need to add or remove resources.
	type args struct {
		id           string
		newresources []model.Resource
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "add resources", args: args{id: packid, newresources: buildResources(0, 3)}, wantErr: false},
		{name: "replace resources", args: args{id: packid, newresources: buildResources(5, 2)}, wantErr: false},
		{name: "empty resources", args: args{id: packid, newresources: []model.Resource{}}, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := mongodao.UpdateResources(tt.args.id, tt.args.newresources); (err != nil) != tt.wantErr {
				t.Errorf("MongoDAO.UpdateResources() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func buildResources(initnumber, amount int) []model.Resource {
	s := make([]model.Resource, amount)
	for i := initnumber; i < (initnumber + amount); i++ {
		res := model.Resource{
			ID:     int16(i),
			Name:   fmt.Sprintf("res%d", i),
			Units:  fmt.Sprintf("u%d", i),
			Amount: 2.5,
			Isfree: false,
		}
		s[i-initnumber] = res
	}
	return s
}
