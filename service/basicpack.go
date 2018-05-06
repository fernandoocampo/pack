package service

import (
	"fmt"
	"time"

	"github.com/fernandoocampo/pack/dao"
	"github.com/fernandoocampo/pack/model"
)

// packDAO makes references to DAO
var packDAO dao.IPackDAO

// BasicPack implements the behaviour made in pack Services
type BasicPack struct {
}

// FindByID implements *IPackService.FindByID using mongo implementation.
func (m *BasicPack) FindByID(id string) (*model.Pack, error) {
	if id == "" {
		return &model.Pack{}, nil
	}

	return packDAO.GetByID(id)
}

// GetByPackCode implements *IPackService.GetByPackCode.
func (m *BasicPack) GetByPackCode(packcode string) (*model.Pack, error) {
	if packcode == "" {
		return &model.Pack{}, nil
	}

	return packDAO.GetByPackCode(packcode)
}

// GetByProductID implements *IPackService.GetByProductId.
func (m *BasicPack) GetByProductID(productid string) (*model.Pack, error) {
	if productid == "" {
		return &model.Pack{}, nil
	}

	return packDAO.GetByProductID(productid)
}

// GetIDByCode implements *IPackDAO.GetIDByCode.
func (m *BasicPack) GetIDByCode(packcode string) (string, error) {
	if packcode == "" {
		return "", nil
	}

	return packDAO.GetIDByCode(packcode)
}

// IsThereThisPack implements *IPackService.IsThereThisPack
func (m *BasicPack) IsThereThisPack(pack *model.PackExists) (bool, error) {
	if pack == nil || pack.MnoID < 1 || (pack.Packcode == "" && pack.ProdID == "") {
		return false, nil
	}
	return packDAO.IsThereThisPack(pack)
}

// Create implements *IPackService.Create.
func (m *BasicPack) Create(packdata *model.Pack) error {
	// check valid input data
	err0 := isValidPackToCreate(packdata)
	if err0 != nil {
		return err0
	}

	// Check that packcode and product id do not exist
	packexists := model.NewPackExists(packdata)
	result, err1 := packDAO.IsThereThisPack(packexists)
	if err1 != nil {
		return fmt.Errorf("06") // existing pack cannot be validated
	}
	if result {
		return fmt.Errorf("07") // There is a pack with mnoid and (pack code or product id)
	}

	packdata.State = model.Active
	packdata.Created = time.Now()
	packdata.Updated = time.Now()
	return packDAO.Create(packdata)
}

// ChangeState implements *IPackService.ChangeState.
func (m *BasicPack) ChangeState(id string, newstate model.PackState) error {
	if id == "" {
		return fmt.Errorf("08") // pack id for change state is empty
	}

	return packDAO.ChangeState(id, newstate)
}

// ChangeProductID implements *IPackService.ChangeProductID.
func (m *BasicPack) ChangeProductID(id string, mnoid int8, newprodid string) error {
	if id == "" || newprodid == "" {
		return fmt.Errorf("09") // pack id for change product id is empty
	}

	// check if the product id with the given mno id already exists
	packexists := new(model.PackExists)
	packexists.MnoID = mnoid
	packexists.ProdID = newprodid
	result, err1 := packDAO.IsThereThisPack(packexists)
	if err1 != nil {
		return fmt.Errorf("06") // existing pack cannot be validated
	}
	if result {
		return fmt.Errorf("10") // There is a pack with mnoid and given product id
	}

	return packDAO.ChangeProductID(id, newprodid)
}

// ChangePackCode implements *IPackService.ChangePackCode.
func (m *BasicPack) ChangePackCode(id string, mnoid int8, newpackcode string) error {
	if id == "" || newpackcode == "" {
		return fmt.Errorf("11") // pack id for change pack code or pack code is empty
	}

	// check if the product id with the given mno id already exists
	packexists := new(model.PackExists)
	packexists.MnoID = mnoid
	packexists.Packcode = newpackcode
	result, err1 := packDAO.IsThereThisPack(packexists)
	if err1 != nil {
		return fmt.Errorf("06") // existing pack cannot be validated
	}
	if result {
		return fmt.Errorf("12") // There is a pack with mnoid and given pack code
	}

	return packDAO.ChangePackCode(id, newpackcode)
}

// ChangeName implements *IPackService.ChangeName.
func (m *BasicPack) ChangeName(id string, newname string) error {
	if id == "" || newname == "" {
		return fmt.Errorf("13") // pack id or new name for change name is empty
	}

	return packDAO.ChangeName(id, newname)
}

// ChangeDesc implements *IPackService.ChangeDesc.
func (m *BasicPack) ChangeDesc(id string, newdesc string) error {
	if id == "" || newdesc == "" {
		return fmt.Errorf("14") // pack id or new desc for change desc is empty
	}

	return packDAO.ChangeDesc(id, newdesc)
}

// ChangeImg implements *IPackService.ChangeImg.
func (m *BasicPack) ChangeImg(id string, newimgurl string) error {
	if id == "" || newimgurl == "" {
		return fmt.Errorf("15") // pack id or new image for change image is empty
	}

	return packDAO.ChangeImg(id, newimgurl)
}

// ChangeKeyword implements *IPackService.ChangeKeyword.
func (m *BasicPack) ChangeKeyword(id string, newkeyword string) error {
	if id == "" || newkeyword == "" {
		return fmt.Errorf("16") // pack id or new key word for change key word is empty
	}

	return packDAO.ChangeKeyword(id, newkeyword)
}

// ChangePrice implements *IPackService.ChangePrice.
func (m *BasicPack) ChangePrice(id string, newprice int) error {
	if id == "" {
		return fmt.Errorf("17") // pack id or new price for change price is empty
	}

	return packDAO.ChangePrice(id, newprice)
}

// ChangePackType implements *IPackService.ChangePackType.
func (m *BasicPack) ChangePackType(id string, newtype *model.Type) error {
	if id == "" || newtype == nil || newtype.ID < 1 || newtype.Name == "" {
		return fmt.Errorf("18") // pack id or new type for change type is empty
	}

	return packDAO.ChangePackType(id, newtype)
}

// ChangeMNO implements *IPackService.ChangeMNO.
func (m *BasicPack) ChangeMNO(id string, prodid string, packcode string, newmno *model.Mno) error {
	if id == "" || prodid == "" || packcode == "" || newmno == nil ||
		newmno.ID < 1 || newmno.Name == "" {
		return fmt.Errorf("19") // pack id or new mno for change mno is empty
	}

	// check if the product id and pack code with the given mno id already exists
	packexists := new(model.PackExists)
	packexists.MnoID = newmno.ID
	packexists.ProdID = prodid
	packexists.Packcode = packcode

	result, err1 := packDAO.IsThereThisPack(packexists)

	if err1 != nil {
		return fmt.Errorf("06") // existing pack cannot be validated
	}
	if result {
		return fmt.Errorf("07") // There is a pack with mnoid and (pack code or product id)
	}

	return packDAO.ChangeMNO(id, newmno)
}

// ChangeValidity implements *IPackService.ChangeValidity.
func (m *BasicPack) ChangeValidity(id string, newterm *model.Term) error {
	if id == "" || newterm == nil || newterm.UnitID < 1 || newterm.Unit == "" {
		return fmt.Errorf("20") // pack id or new term for change validity is empty
	}

	return packDAO.ChangeValidity(id, newterm)
}

// ChangeCurrency implements *IPackService.ChangeCurrency.
func (m *BasicPack) ChangeCurrency(id string, newccy *model.Currency) error {
	if id == "" || newccy == nil || newccy.ID < 1 || newccy.Name == "" {
		return fmt.Errorf("21") // pack id or new currency for change currency is empty
	}

	return packDAO.ChangeCurrency(id, newccy)
}

// MoveStock implements *IPackService.MoveStock.
func (m *BasicPack) MoveStock(id string, amount int) error {
	if id == "" || amount == 0 {
		return fmt.Errorf("23") // invalid values
	}
	return packDAO.ChangeStock(id, amount)
}

// Delete implements *IPackService.Delete.
func (m *BasicPack) Delete(id string) error {
	if id == "" {
		return fmt.Errorf("22") // pack id for delete is empty
	}

	return packDAO.Delete(id)
}

// UpdateResources replace the resources that we configured for a pack.
func (m *BasicPack) UpdateResources(id string, newresources []model.Resource) error {
	if id == "" || newresources == nil {
		return fmt.Errorf("24") // pack id or resources for update resources are empty
	}

	return packDAO.UpdateResources(id, newresources)
}

// DeleteResources remove the resources that we configured for a pack.
func (m *BasicPack) DeleteResources(id string) error {
	if id == "" {
		return fmt.Errorf("25") // pack id for delete resources is empty
	}

	return packDAO.UpdateResources(id, []model.Resource{})
}

// isValidPackToCreate validates if model.Pack data is ok..
func isValidPackToCreate(pack *model.Pack) error {
	if pack == nil {
		return fmt.Errorf("00")
	}
	errbasic := validateBasicData(pack)
	if errbasic != nil {
		return errbasic
	}
	errcomplex := validateComplexData(pack)

	return errcomplex
}

func validateBasicData(pack *model.Pack) error {
	if pack.ProdID == "" || pack.Packcode == "" || pack.Name == "" {
		return fmt.Errorf("01")
	}
	if pack.Desc == "" || pack.Kwds == "" || pack.Img == "" {
		return fmt.Errorf("02")
	}
	return nil
}

func validateComplexData(pack *model.Pack) error {
	if pack.Packtype == nil || pack.Packtype.ID < 1 || pack.Packtype.Name == "" {
		return fmt.Errorf("03")
	}
	if pack.Mno == nil || pack.Mno.ID < 1 || pack.Mno.Name == "" {
		return fmt.Errorf("04")
	}
	if pack.Term == nil || pack.Term.UnitID < 1 || pack.Term.Unit == "" {
		return fmt.Errorf("05")
	}
	return nil
}

// SetPackDAO set the pack dao for this business logic.
func SetPackDAO(dao dao.IPackDAO) {
	packDAO = dao
}
