package request

import (
	"database/sql"

	"github.com/GolangNorhtwindRestApi/helper"
)

type Repository interface {
	GetRequestsOrders(params *getRequestsRequest) ([]*RequestOrder, error)
	GetTotalRequestsOrders() (int64, error)
	GetRequestById(params *getRequestBYIDRequest) (*RequestOrder, error)
	InsertRequestOrder(params *addRequestOrderRequest) (int64, error)
	UpdateRequestOrder(params *updateRequestOrderRequest) (int64, error)
	CancelRequestOrder(params *cancelRequestOrderRequest) (int64, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db}
}

func (repo *repository) GetRequestsOrder(params *getRequestsRequest) ([]*RequestOrder, error) {
	const sql = `SELECT
	re.id,
	re.id_customer as id_customer,
	ce.fantasy_name as customer,
	re.id_supplier as id_supplier,
	su.fantasy_name as supplier,
	re.date_request,
	re.time_request,
	re.total_quantity,
	re.total_price,
	st.description as status
FROM 	goolivery_provider_customer.orders_requests re
INNER JOIN  goolivery_provider_global.status st ON re.id_status = st.id
INNER JOIN  goolivery_provider_customer.customers ce ON re.id_customer = ce.id
INNER JOIN  goolivery_provider_supplier.suppliers su ON re.id_supplier = su.id
WHERE re.cancel = false
LIMIT ? OFFSET ?
`
	results, err := repo.db.Query(sql, params.Limit, params.Offset)
	helper.Catch(err)
	var requestsOrders []*RequestOrder
	for results.Next() {
		reqOrder := &RequestOrder{}
		err = results.Scan(
			&reqOrder.Id,
			&reqOrder.IdCustomer,
			&reqOrder.Customer,
			&reqOrder.IdSupplier,
			&reqOrder.Supplier,
			&reqOrder.DateRequest,
			&reqOrder.TimeRequest,
			&reqOrder.TotalQuantity,
			&reqOrder.TotalPrice,
			&reqOrder.IdStatus,
			&reqOrder.DescriptStatus)
		helper.Catch(err)

		requestsOrders = append(requestsOrders, reqOrder)
	}
	return requestsOrders, nil
}

func (repo *repository) GetTotalRequestsOrders() (int64, error) {
	const sql = "SELECT COUNT(*) FROM orders_requests"
	var total int64
	row := repo.db.QueryRow(sql)
	err := row.Scan(&total)
	helper.Catch(err)
	return total, nil
}

func (repo *repository) GetRequestsOrdersById(params *getRequestBYIDRequest) (*RequestOrder, error) {
	const sql = `SELECT re.id, 
	re.id_customer,
	ce.fantasy_name as customer,
	re.id_supplier as id_supplier,
	su.fantasy_name as supplier,
    re.date_request,
	re.time_request,
	re.total_quantity,
	re.total_price,
	re.tot_price_widhout_tax,
	re.tax,
	re.other_tax,
	re.country,
	re.doc_state,
	re.email,
	re.celphone,
	re.description,
	re.dt_reg,
	re.id_user,
	re.cancel,
	re.id_cancel,
	re.delivery_supplier,
	re.id_status,
	st.description as status,
	re.id_event,
	ev.description as event
			FROM 	goolivery_provider_customer.orders_requests re
			INNER JOIN  goolivery_provider_global.status st ON re.id_status = st.id
			INNER JOIN  goolivery_provider_customer.customers ce ON re.id_customer = ce.id
			INNER JOIN  goolivery_provider_supplier.suppliers su ON re.id_supplier = su.id
			INNER JOIN goolivery_provider_global.events ev ON re.id_event = ev.id
			WHERE re.cancel = 'false' and
				  re.id= ? and
				  re.country= ?`
	row := repo.db.QueryRow(sql, params.RequestID, params.Country)
	requestOrder := &RequestOrder{}
	err := row.Scan(&requestOrder.Id,
		&requestOrder.IdCustomer,
		&requestOrder.Customer,
		&requestOrder.IdSupplier,
		&requestOrder.Supplier,
		&requestOrder.DateRequest,
		&requestOrder.TimeRequest,
		&requestOrder.TotalQuantity,
		&requestOrder.TotalPrice,
		&requestOrder.TotPriceWidhoutTax,
		&requestOrder.Tax,
		&requestOrder.Other_tax,
		&requestOrder.Country,
		&requestOrder.DocState,
		&requestOrder.Email,
		&requestOrder.Cellphone,
		&requestOrder.Description,
		&requestOrder.DtReg,
		&requestOrder.IdUser,
		&requestOrder.Cancel,
		&requestOrder.IdCancel,
		&requestOrder.DeliverySupplier,
		&requestOrder.IdStatus,
		&requestOrder.DescriptStatus,
		&requestOrder.Id_event,
		&requestOrder.Event)
	helper.Catch(err)

	return requestOrder, nil
}

func (repo *repository) InsertResquestOrder(params *addRequestOrderRequest) (int64, error) {
	const sql = `INSERT INTO goolivery_provider_customer.orders_requests(
		id_supplier,
		id_supplier,
		total_quantity,
		total_price,
		tot_price_widhout_tax,
		tax,
		other_tax,
		country,
		doc_state,
		email,
		celphone,
		description,
		id_user,
		delivery_supplier,
		id_status,
		id_event)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 1, 1)`

	result, err := repo.db.Exec(sql,
		params.IdCustomer,
		params.IdSupplier,
		params.TotalQuantity,
		params.TotalPrice,
		params.TotPriceWidhoutTax,
		params.Tax,
		params.Other_tax,
		params.Country,
		params.DocState,
		params.Email,
		params.Cellphone,
		params.Description,
		params.IdUser,
		params.DeliverySupplier,
		params.IdStatus,
		params.idEvent)

	helper.Catch(err)
	id, _ := result.LastInsertId()
	return id, nil
}

func (repo *repository) UpdateRequestOrder(params *updateRequestOrderRequest) (int64, error) {
	const sql = `UPDATE goolivery_provider_customer.orders_requests
	SET 
		id_customer=?,
		id_supplier=?,
		total_quantity=?,
		total_price=?,
		tot_price_widhout_tax=?,
		tax=?,
		other_tax=?,
		country=?,
		doc_state=?,
		email=?,
		celphone=?,
		description=?,
		delivery_supplier=?,
		id_status=?,
		id_event=?
	WHERE id=?`
	result, err := repo.db.Exec(sql,
		params.IdCustomer,
		params.IdSupplier,
		params.TotalQuantity,
		params.TotalPrice,
		params.TotPriceWidhoutTax,
		params.Tax,
		params.Other_tax,
		params.Country,
		params.DocState,
		params.Email,
		params.Cellphone,
		params.Description,
		params.DeliverySupplier,
		params.IdStatus,
		params.idEvent,
		params.Id)

	helper.Catch(err)
	id, _ := result.LastInsertId()
	return id, nil
}

func (repo *repository) CancelRequestOrder(params *cancelRequestOrderRequest) (int64, error) {
	const sql = `UPDATE goolivery_provider_customer.orders_requests
	SET cancel='Y',
		id_cancel=?,
		id_status=?,
		id_event=?
	WHERE id =?`
	result, err := repo.db.Exec(sql, params.Id)
	helper.Catch(err)
	count, err := result.RowsAffected()
	helper.Catch(err)
	return count, nil
}
