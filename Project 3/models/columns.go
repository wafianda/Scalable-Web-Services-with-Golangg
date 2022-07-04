package models

type columnManager struct{}

var ColumnManager = columnManager{}

func (c *columnManager) Create(column *Column) error {
	return DB.Create(&column).Error
}

func (c *columnManager) Update(column *Column, id uint64) error {
	return DB.Model(Column{}).
		Where("id = ?", id).
		Updates(column).
		Error
}

func (c *columnManager) GetById(id uint64) (*Column, error) {
	var column Column
	return &column, DB.First(&column, id).Error
}

func (c *columnManager) GetAll() (*[]Column, error) {
	var columns []Column
	return &columns, DB.Find(&columns).Error
}

func (c *columnManager) DeleteById(id uint) error {
	return DB.Unscoped().Delete(&Column{}, id).Error
}
