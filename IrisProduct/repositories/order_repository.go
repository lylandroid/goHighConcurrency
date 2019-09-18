package repositories

import (
	"../datamodels"
	"../common"
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
)

type IOrderRepository interface {
	Conn() error
	Insert(order *datamodels.Order) (int64, error)
	Delete(int64) bool
	Update(order *datamodels.Order) error
	SelectByKey(int64) (*datamodels.Order, error)
	SelectAll() ([]*datamodels.Order, error)
	SelectAllWithInfo() (map[int]map[string]string, error)
}

func NewOrderManagerRepository(table string, sql *sql.DB) IOrderRepository {
	return &OrderManagerRepository{
		table:     table,
		mySqlConn: sql,
	}

}

type OrderManagerRepository struct {
	table     string
	mySqlConn *sql.DB
	gromDb    *gorm.DB
}

func (r *OrderManagerRepository) ConnGorm() (err error) {
	if r.gromDb == nil {
		conn, err := common.NewMySqlGormConn()
		if err != nil {
			return err
		}
		r.gromDb = conn
		if r.table == "" {
			r.table = "order"
		}
	}
	return nil
}

func (r *OrderManagerRepository) Conn() error {
	if r.mySqlConn == nil {
		sqlConn, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		r.mySqlConn = sqlConn
		if r.table == "" {
			r.table = "order"
		}
	}
	return nil
}

func (r *OrderManagerRepository) Insert(order *datamodels.Order) (productId int64, err error) {
	if err = r.Conn(); err != nil {
		return 0, err
	}
	sql := fmt.Sprintf(`INSERT %s set userID=?,productID=?,orderStatus=?`, r.table)
	result, err := r.mySqlConn.Exec(sql, order.UserId, order.ProductId, order.OrderStatus)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (r *OrderManagerRepository) Delete(orderId int64) (isOK bool) {
	if err := r.Conn(); err != nil {
		return false
	}
	sql := fmt.Sprintf(`DELETE FROM %s where ID=?`, r.table)
	_, err := r.mySqlConn.Exec(sql, strconv.FormatInt(orderId, 10))
	return err == nil
}

func (r *OrderManagerRepository) Update(order *datamodels.Order) (err error) {
	if err := r.Conn(); err != nil {
		return err
	}
	sql := fmt.Sprintf(`UPDATE %s set userID=?,productID=?,orderStatus=? where ID=?`, r.table)
	stmt, err := r.mySqlConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.UserId, order.ProductId, order.OrderStatus, strconv.FormatInt(order.ID, 10))
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderManagerRepository) SelectByKey(id int64) (order *datamodels.Order, err error) {
	if err := r.Conn(); err != nil {
		return nil, err
	}
	sql := fmt.Sprintf(`select * from %s where ID=?`, r.table)
	rows, err := r.mySqlConn.Query(sql, strconv.FormatInt(id, 10))
	if err != nil {
		return nil, err
	}
	rowMap := common.GetResultRow(rows)
	order = &datamodels.Order{}
	common.DataToStructByTagSql(rowMap, order)
	return order, nil

}

func (r *OrderManagerRepository) SelectAll() (orders []*datamodels.Order, err error) {
	if err := r.ConnGorm(); err != nil {
		return nil, err
	}
	orders = []*datamodels.Order{}
	r.gromDb.Find(&orders)
	return
}

/*func (r *OrderManagerRepository) SelectAll() (orders []*datamodels.Order, err error) {
	if err := r.Conn(); err != nil {
		return nil, err
	}
	sql := fmt.Sprintf(`select * from %s`, r.table)
	rows, err := r.mySqlConn.Query(sql)
	if err != nil {
		return nil, err
	}
	rowsMap := common.GetResultRows(rows)
	if len(rowsMap) == 0 {
		return orders, nil
	}
	for _, v := range rowsMap {
		order := &datamodels.Order{}
		common.DataToStructByTagSql(v, order)
		orders = append(orders, order)
	}
	return orders, nil
}*/

func (r *OrderManagerRepository) SelectAllWithInfo() (orderMap map[int]map[string]string, err error) {
	if err := r.Conn(); err != nil {
		return nil, err
	}
	sql := `SELECT o.ID,p.productName,o.orderStatus FROM imooc.order as o left join product as p on o.productID=p.ID`
	rows, err := r.mySqlConn.Query(sql)
	if err != nil {
		return nil, err
	}
	return common.GetResultRows(rows), nil
}
/*func (r *OrderManagerRepository) SelectAllWithInfo() (orderMap map[int]map[string]string, err error) {
	if err := r.ConnGorm(); err != nil {
		return nil, err
	}
	//sql := `SELECT o.ID,p.productName,o.orderStatus FROM imooc.order as o left join product as p on o.productID=p.ID`
	//rows, err := r.mySqlConn.Query(sql)
	orderMap = map[int]map[string]string{}
	orders := []datamodels.Order{}
	/*find := r.gromDb.Table("order").
		Select("order.ID,product.productName,order.orderStatus").
		Joins("left join product on order.productID = product.ID").Find(&orders)
		r.gromDb.Find(&orders)
	find := r.gromDb.Table("order").Find(&orders)
	fmt.Println("orders: ", orders, find.Error)
	return orderMap, nil
}*/
