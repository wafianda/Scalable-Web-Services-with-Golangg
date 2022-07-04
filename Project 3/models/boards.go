package models

type boardManager struct{}

var BoardManager = boardManager{}

func (k *boardManager) Create(board *Board) error {
	return DB.Create(&board).Error
}

func (k *boardManager) Update(board *Board, id uint64) error {
	return DB.Model(Board{}).
		Where("id = ?", id).
		Updates(board).
		Error
}

func (k *boardManager) GetById(id uint64) (*Board, error) {
	var board Board
	return &board, DB.First(&board, "id = ?", id).Error
}

func (k *boardManager) GetAll() (*[]Board, error) {
	var boards []Board
	return &boards, DB.Preload("Tickets").Find(&boards).Error
}

func (k *boardManager) DeleteById(id uint64) error {
	return DB.Unscoped().Delete(&Board{}, id).Error
}
