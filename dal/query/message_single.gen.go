// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/tdfxlyh/go-gin-api/dal/models"
)

func newMessageSingle(db *gorm.DB, opts ...gen.DOOption) messageSingle {
	_messageSingle := messageSingle{}

	_messageSingle.messageSingleDo.UseDB(db, opts...)
	_messageSingle.messageSingleDo.UseModel(&models.MessageSingle{})

	tableName := _messageSingle.messageSingleDo.TableName()
	_messageSingle.ALL = field.NewAsterisk(tableName)
	_messageSingle.ID = field.NewInt64(tableName, "id")
	_messageSingle.SenderUserID = field.NewInt64(tableName, "sender_user_id")
	_messageSingle.ReceiverUserID = field.NewInt64(tableName, "receiver_user_id")
	_messageSingle.MessageType = field.NewInt64(tableName, "message_type")
	_messageSingle.Content = field.NewString(tableName, "content")
	_messageSingle.Extra = field.NewString(tableName, "extra")
	_messageSingle.ReadStatusInfo = field.NewInt32(tableName, "read_status_info")
	_messageSingle.SenderStatusInfo = field.NewInt32(tableName, "sender_status_info")
	_messageSingle.ReceiverStatusInfo = field.NewInt32(tableName, "receiver_status_info")
	_messageSingle.Withdraw = field.NewInt32(tableName, "withdraw")
	_messageSingle.CreateTime = field.NewTime(tableName, "create_time")
	_messageSingle.ModifyTime = field.NewTime(tableName, "modify_time")
	_messageSingle.Status = field.NewInt32(tableName, "status")

	_messageSingle.fillFieldMap()

	return _messageSingle
}

type messageSingle struct {
	messageSingleDo messageSingleDo

	ALL                field.Asterisk
	ID                 field.Int64  // 主键id
	SenderUserID       field.Int64  // 消息发送方
	ReceiverUserID     field.Int64  // 消息接收方
	MessageType        field.Int64  // 消息类型 1:文本 2:图片 3:音频 4:视频 5:文件
	Content            field.String // 消息内容
	Extra              field.String // 扩展信息
	ReadStatusInfo     field.Int32  // 判断接收者 0:未读 1:已读
	SenderStatusInfo   field.Int32  // 发送方信息状态 0:正常 1:删除
	ReceiverStatusInfo field.Int32  // 接收方信息状态 0:正常 1:删除
	Withdraw           field.Int32  // 是否撤回 0:正常 1:撤回
	CreateTime         field.Time   // 创建时间
	ModifyTime         field.Time   // 修改时间
	Status             field.Int32  // 0存在，1删除

	fieldMap map[string]field.Expr
}

func (m messageSingle) Table(newTableName string) *messageSingle {
	m.messageSingleDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m messageSingle) As(alias string) *messageSingle {
	m.messageSingleDo.DO = *(m.messageSingleDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *messageSingle) updateTableName(table string) *messageSingle {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewInt64(table, "id")
	m.SenderUserID = field.NewInt64(table, "sender_user_id")
	m.ReceiverUserID = field.NewInt64(table, "receiver_user_id")
	m.MessageType = field.NewInt64(table, "message_type")
	m.Content = field.NewString(table, "content")
	m.Extra = field.NewString(table, "extra")
	m.ReadStatusInfo = field.NewInt32(table, "read_status_info")
	m.SenderStatusInfo = field.NewInt32(table, "sender_status_info")
	m.ReceiverStatusInfo = field.NewInt32(table, "receiver_status_info")
	m.Withdraw = field.NewInt32(table, "withdraw")
	m.CreateTime = field.NewTime(table, "create_time")
	m.ModifyTime = field.NewTime(table, "modify_time")
	m.Status = field.NewInt32(table, "status")

	m.fillFieldMap()

	return m
}

func (m *messageSingle) WithContext(ctx context.Context) *messageSingleDo {
	return m.messageSingleDo.WithContext(ctx)
}

func (m messageSingle) TableName() string { return m.messageSingleDo.TableName() }

func (m messageSingle) Alias() string { return m.messageSingleDo.Alias() }

func (m *messageSingle) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *messageSingle) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 13)
	m.fieldMap["id"] = m.ID
	m.fieldMap["sender_user_id"] = m.SenderUserID
	m.fieldMap["receiver_user_id"] = m.ReceiverUserID
	m.fieldMap["message_type"] = m.MessageType
	m.fieldMap["content"] = m.Content
	m.fieldMap["extra"] = m.Extra
	m.fieldMap["read_status_info"] = m.ReadStatusInfo
	m.fieldMap["sender_status_info"] = m.SenderStatusInfo
	m.fieldMap["receiver_status_info"] = m.ReceiverStatusInfo
	m.fieldMap["withdraw"] = m.Withdraw
	m.fieldMap["create_time"] = m.CreateTime
	m.fieldMap["modify_time"] = m.ModifyTime
	m.fieldMap["status"] = m.Status
}

func (m messageSingle) clone(db *gorm.DB) messageSingle {
	m.messageSingleDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m messageSingle) replaceDB(db *gorm.DB) messageSingle {
	m.messageSingleDo.ReplaceDB(db)
	return m
}

type messageSingleDo struct{ gen.DO }

func (m messageSingleDo) Debug() *messageSingleDo {
	return m.withDO(m.DO.Debug())
}

func (m messageSingleDo) WithContext(ctx context.Context) *messageSingleDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m messageSingleDo) ReadDB() *messageSingleDo {
	return m.Clauses(dbresolver.Read)
}

func (m messageSingleDo) WriteDB() *messageSingleDo {
	return m.Clauses(dbresolver.Write)
}

func (m messageSingleDo) Session(config *gorm.Session) *messageSingleDo {
	return m.withDO(m.DO.Session(config))
}

func (m messageSingleDo) Clauses(conds ...clause.Expression) *messageSingleDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m messageSingleDo) Returning(value interface{}, columns ...string) *messageSingleDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m messageSingleDo) Not(conds ...gen.Condition) *messageSingleDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m messageSingleDo) Or(conds ...gen.Condition) *messageSingleDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m messageSingleDo) Select(conds ...field.Expr) *messageSingleDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m messageSingleDo) Where(conds ...gen.Condition) *messageSingleDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m messageSingleDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *messageSingleDo {
	return m.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (m messageSingleDo) Order(conds ...field.Expr) *messageSingleDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m messageSingleDo) Distinct(cols ...field.Expr) *messageSingleDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m messageSingleDo) Omit(cols ...field.Expr) *messageSingleDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m messageSingleDo) Join(table schema.Tabler, on ...field.Expr) *messageSingleDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m messageSingleDo) LeftJoin(table schema.Tabler, on ...field.Expr) *messageSingleDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m messageSingleDo) RightJoin(table schema.Tabler, on ...field.Expr) *messageSingleDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m messageSingleDo) Group(cols ...field.Expr) *messageSingleDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m messageSingleDo) Having(conds ...gen.Condition) *messageSingleDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m messageSingleDo) Limit(limit int) *messageSingleDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m messageSingleDo) Offset(offset int) *messageSingleDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m messageSingleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *messageSingleDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m messageSingleDo) Unscoped() *messageSingleDo {
	return m.withDO(m.DO.Unscoped())
}

func (m messageSingleDo) Create(values ...*models.MessageSingle) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m messageSingleDo) CreateInBatches(values []*models.MessageSingle, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m messageSingleDo) Save(values ...*models.MessageSingle) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m messageSingleDo) First() (*models.MessageSingle, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.MessageSingle), nil
	}
}

func (m messageSingleDo) Take() (*models.MessageSingle, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.MessageSingle), nil
	}
}

func (m messageSingleDo) Last() (*models.MessageSingle, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.MessageSingle), nil
	}
}

func (m messageSingleDo) Find() ([]*models.MessageSingle, error) {
	result, err := m.DO.Find()
	return result.([]*models.MessageSingle), err
}

func (m messageSingleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.MessageSingle, err error) {
	buf := make([]*models.MessageSingle, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m messageSingleDo) FindInBatches(result *[]*models.MessageSingle, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m messageSingleDo) Attrs(attrs ...field.AssignExpr) *messageSingleDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m messageSingleDo) Assign(attrs ...field.AssignExpr) *messageSingleDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m messageSingleDo) Joins(fields ...field.RelationField) *messageSingleDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m messageSingleDo) Preload(fields ...field.RelationField) *messageSingleDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m messageSingleDo) FirstOrInit() (*models.MessageSingle, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.MessageSingle), nil
	}
}

func (m messageSingleDo) FirstOrCreate() (*models.MessageSingle, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.MessageSingle), nil
	}
}

func (m messageSingleDo) FindByPage(offset int, limit int) (result []*models.MessageSingle, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m messageSingleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m messageSingleDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m messageSingleDo) Delete(models ...*models.MessageSingle) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *messageSingleDo) withDO(do gen.Dao) *messageSingleDo {
	m.DO = *do.(*gen.DO)
	return m
}
