package special

import "specialteam/config"

func (st *Specialteam) CreateSpecialteam() error {
	if err := config.DB.Create(st).Error; err != nil {
		return err
	}
	return nil
}

func (st *Specialteam) UpdateSpecialteam(id int) error {
	if err := config.DB.Model(&Specialteam{}).Where("id = ?", id).Updates(st).Error; err != nil {
		return err
	}
	return nil
}

func (st *Specialteam) DeleteSpecialteam() error {
	if err := config.DB.Delete(st).Error; err != nil {
		return err
	}
	return nil
}

func GetOneById(id int) (Specialteam, error) {
	var st Specialteam
	result := config.DB.Where("id = ?", id).First(&st)
	return st, result.Error
}

func GetAll(keywords string) ([]Specialteam, error) {
	var st []Specialteam
	result := config.DB.Where("email LIKE ? OR name LIKE ?", "%"+keywords+"%", "%"+keywords+"%").Find(&st)

	return st, result.Error
}
