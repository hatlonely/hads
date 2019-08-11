package mysqldb

import (
	"fmt"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Account struct {
	ID        int       `gorm:"type:bigint(20) auto_increment;primary_key" json:"id"`
	Email     string    `gorm:"type:varchar(64);not null;unique_index:email_idx" json:"email"`
	Phone     string    `gorm:"type:varchar(64);not null;unique_index:phone_idx" json:"phone"`
	FirstName string    `gorm:"type:varchar(32);not null" json:"firstName"`
	LastName  string    `gorm:"type:varchar(32);not null" json:"lastName"`
	Password  string    `gorm:"type:varchar(32);not null" json:"password"`
	Birthday  time.Time `gorm:"type:timestamp;not null" json:"birthday"`
	Gender    int       `gorm:"type:int(1);not null" json:"gender"`
	Role      int       `gorm:"type:bigint(20) default 0;not null" json:"role"`
}

func (m *MysqlDB) SelectAccountByPhoneOrEmail(key string) (*Account, error) {
	account := &Account{}
	if key == "" {
		return nil, fmt.Errorf("account key is null")
	}
	if err := m.db.Where("phone=?", key).Or("email=?", key).First(account).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return account, nil
}

func (m *MysqlDB) InsertAccount(account *Account) (bool, error) {
	if account.Email == "" || account.Phone == "" {
		return false, fmt.Errorf("email or phone are is null, account [%#v]", account)
	}

	accountDB := &Account{}
	var conditions []string
	if account.ID != 0 {
		conditions = append(conditions, fmt.Sprintf("id=%v", account.ID))
	}
	if account.Phone != "" {
		conditions = append(conditions, fmt.Sprintf("phone='%v'", account.Phone))
	}
	if account.Email != "" {
		conditions = append(conditions, fmt.Sprintf("email='%v'", account.Email))
	}
	err := m.db.Where(strings.Join(conditions, " OR ")).First(accountDB).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if err != gorm.ErrRecordNotFound {
		if accountDB.ID == account.ID {
			return false, fmt.Errorf("accountID [%v] is already exists", accountDB.ID)
		}
		if accountDB.Phone == account.Phone {
			return false, fmt.Errorf("phone [%v] is already exists", accountDB.Phone)
		}
		if accountDB.Email == account.Email {
			return false, fmt.Errorf("email [%v] is already exists", accountDB.Email)
		}
	}

	if err := m.db.Create(account).Error; err != nil {
		return false, err
	}

	return true, nil
}
