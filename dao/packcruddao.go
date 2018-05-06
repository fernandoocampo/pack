package dao

import "github.com/fernandoocampo/pack/model"

// IPackDAO defines pack data access behavior for management purpose.
type IPackDAO interface {
	// GetByID search a pack with the given id
	// and return it.
	GetByID(id string) (*model.Pack, error)
	// GetIDByCode searches a pack with the given
	// pack code and return it.
	GetIDByCode(packcode string) (string, error)
	// GetByPackCode search a pack with the given shortcode
	// and return it.
	GetByPackCode(packcode string) (*model.Pack, error)
	// GetByProductID search a pack with the given mno internal id
	// and return it.
	GetByProductID(productid string) (*model.Pack, error)
	// IsThereThisPack search a pack with the given parameters, it checks
	// that combination between mnoid and pack code or mnoid and
	// product id or mnoid don't exist
	IsThereThisPack(pack *model.PackExists) (bool, error)
	// Create inserts a new Pack in the system. Returns
	// true if the Pack is created
	Create(packdata *model.Pack) error
	// ChangeState changes the state of a pack.
	ChangeState(id string, newstate model.PackState) error
	// ChangeProductID changes the mno internal product id.
	ChangeProductID(id string, newprodid string) error
	// ChangePackCode changes the pack short code
	ChangePackCode(id string, newpackcode string) error
	// ChangeName changes the name of an existent pack.
	ChangeName(id string, newname string) error
	// ChangeDesc changes the description of an existent pack.
	ChangeDesc(id string, newdesc string) error
	// ChangeImg changes the image url of the given pack.
	ChangeImg(id string, newimgurl string) error
	// ChangeKeyword changes the key words of the given pack.
	ChangeKeyword(id string, newkeyword string) error
	// ChangePrice changes the price of an existent pack.
	ChangePrice(id string, newprice int) error
	// ChangePackType changes the pack type of an existent pack
	ChangePackType(id string, newtype *model.Type) error
	// ChangeMNO changes the Mobile Network Operator owner of the pack.
	ChangeMNO(id string, newmno *model.Mno) error
	// ChangeValidity changes the validity of the given pack.
	ChangeValidity(id string, newterm *model.Term) error
	// ChangeCurrency changes the currency of price of the given pack.
	ChangeCurrency(id string, newccy *model.Currency) error
	// ChangeStock add or reduce stock to the given pack.
	ChangeStock(id string, amount int) error
	// UpdateResources replace the resources that we configured for a pack.
	// Send newresources empty if you want to remove all the resources.
	UpdateResources(id string, newresources []model.Resource) error
	// Delete removes an existent Pack and returns true if the pack can be deleted.
	Delete(id string) error
}
