package repositories

import (
	"../common"
	"../datamodels"
	"database/sql"
	"strconv"
)

type IProductRepository interface {
	//连接数据库
	Conn() error
	Insert(product *datamodels.Product) (int64, error)
	Delete(int64) bool
	Update(product *datamodels.Product) error
	SelectByKey(int64) (*datamodels.Product, error)
	SelectAll() ([]*datamodels.Product, error)
}

func NewProductRepositoryImp(table string, db *sql.DB) IProductRepository {
	return &ProductManager{
		Table:     table,
		MySqlConn: db,
	}
}

type ProductManager struct {
	Table     string
	MySqlConn *sql.DB
}

func (p *ProductManager) Conn() (err error) {
	if p.MySqlConn != nil {
		p.MySqlConn, err = common.NewMysqlConn()
	}
	if p.Table == "" {
		p.Table = "product"
	}
	return
}

func (p *ProductManager) Insert(product *datamodels.Product) (id int64, err error) {
	//判断连接是否存在
	if err = p.Conn(); err != nil {
		return 0, err
	}
	//准备sql
	sql := "INSERT ? SET productName=?,productNum=?,productImage=?,productUrl=?"
	if stmt, err := p.MySqlConn.Prepare(sql); err != nil {
		return 0, err
	} else if result, err := stmt.Exec(p.Table, product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl); err != nil {
		return 0, err
	} else {
		defer stmt.Close()
		return result.LastInsertId()
	}

}

func (p *ProductManager) Delete(productId int64) bool {
	//判断连接是否存在
	if err := p.Conn(); err != nil {
		return false
	}
	sql := "DELETE FROM ? where ID=?"
	stmt, err := p.MySqlConn.Prepare(sql)
	if err != nil {
		return false
	}
	if _, err := stmt.Exec(p.Table, productId); err != nil {
		return false
	}
	return true
}

func (p *ProductManager) Update(product *datamodels.Product) (err error) {
	//判断连接是否存在
	if err := p.Conn(); err != nil {
		return err
	}
	sql := "UPDATE ? SET productName=?,productNum=?,productImage=?,productUrl=? WHERE ID=?"
	if stmt, err := p.MySqlConn.Prepare(sql); err != nil {
		return err
	} else if _, err := stmt.Exec(p.Table, product.ProductName, product.ProductNum, product.ProductImage,
		product.ProductUrl, strconv.FormatInt(product.ID, 10)); err != nil {
		return err
	} else {
		defer stmt.Close()
		return nil
	}

}

func (p *ProductManager) SelectByKey(productId int64) (productResult *datamodels.Product, err error) {
	//判断连接是否存在
	if err := p.Conn(); err != nil {
		return nil, err
	}
	sql := "SELECT ID,productName,productNum,productImage.productUrl FROM ? WHERE ID=?"
	stmt, err := p.MySqlConn.Prepare(sql)
	if err != nil {
		return nil, err
	}
	/*row := stmt.QueryRow(p.Table, strconv.FormatInt(productId, 10))
	return readRow(row)*/
	rows, err := stmt.Query(p.Table, strconv.FormatInt(productId, 10))
	rowMap := common.GetResultRow(rows)
	common.DataToStructByTagSql(rowMap, &productResult)
	return
}

func (p *ProductManager) SelectAll() (products []*datamodels.Product, err error) {
	//判断连接是否存在
	if err := p.Conn(); err != nil {
		return nil, err
	}
	sql := "SELECT ID,productName,productNum,productImage.productUrl FROM ?"
	stmt, err := p.MySqlConn.Prepare(sql)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(p.Table)
	if err != nil {
		return nil, err
	}
	rowsMap := common.GetResultRows(rows)
	for _, v := range rowsMap {
		var product *datamodels.Product
		common.DataToStructByTagSql(v, &product)
		products = append(products, product)
	}
	return products, nil
}
