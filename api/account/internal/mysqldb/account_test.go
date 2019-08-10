package mysqldb

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMysqlDB_InsertAccount(t *testing.T) {
	m, err := NewMysqlDB("hatlonely:keaiduo1@/hads?charset=utf8&parseTime=True&loc=Local")
	Convey("test mysqldb insert account", t, func() {
		So(err, ShouldBeNil)
		So(m, ShouldNotBeNil)

		m.db.Where("email='hatlonely1@foxmail.com'").
			Or("phone='13112345678'").Delete(&Account{})

		Convey("insert account", func() {
			birthday, _ := time.Parse("2006-01-02", "1992-01-01")
			ok, err := m.InsertAccount(&Account{
				Email:      "hatlonely1@foxmail.com",
				Phone:      "13112345678",
				Password:   "123456",
				FirstName:  "孙",
				SecondName: "悟空",
				Birthday:   birthday,
				Gender:     1,
			})
			So(err, ShouldBeNil)
			So(ok, ShouldBeTrue)

			account, err := m.SelectAccountByPhoneOrEmail("hatlonely1@foxmail.com")
			So(err, ShouldBeNil)
			So(account.Email, ShouldEqual, "hatlonely1@foxmail.com")
			So(account.Phone, ShouldEqual, "13112345678")
			So(account.Password, ShouldEqual, "123456")
			So(account.FirstName, ShouldEqual, "孙")
			So(account.SecondName, ShouldEqual, "悟空")
			So(account.Birthday, ShouldEqual, birthday)
			So(account.Gender, ShouldEqual, 1)

			Convey("insert dup email", func() {
				ok, err := m.InsertAccount(&Account{
					Email:      "hatlonely1@foxmail.com",
					Phone:      "13812345678",
					Password:   "123456",
					FirstName:  "孙",
					SecondName: "悟空",
					Birthday:   birthday,
					Gender:     1,
				})
				So(err, ShouldNotBeNil)
				So(ok, ShouldBeFalse)
			})

			Convey("insert dup phone", func() {
				ok, err := m.InsertAccount(&Account{
					Email:      "hatlonely2@foxmail.com",
					Phone:      "13112345678",
					Password:   "123456",
					FirstName:  "孙",
					SecondName: "悟空",
					Birthday:   birthday,
					Gender:     1,
				})
				So(err, ShouldNotBeNil)
				So(ok, ShouldBeFalse)
			})
		})
	})
}

func TestMysqlDB_SelectAccountByUsernameOrTelephoneOrEmail(t *testing.T) {
	m, err := NewMysqlDB("hatlonely:keaiduo1@/hads?charset=utf8&parseTime=True&loc=Local")
	Convey("test mysqldb select account by username or phone or email", t, func() {
		So(err, ShouldBeNil)
		So(m, ShouldNotBeNil)

		m.db.Where("email='hatlonely1@foxmail.com'").
			Or("phone='13112345678'").Delete(&Account{})

		Convey("select account use empty key", func() {
			account, err := m.SelectAccountByPhoneOrEmail("")
			So(account, ShouldBeNil)
			So(err, ShouldNotBeNil)
		})

		Convey("select account use nonexists key", func() {
			account, err := m.SelectAccountByPhoneOrEmail("hatlonely1")
			So(account, ShouldBeNil)
			So(err, ShouldBeNil)
		})

		birthday, _ := time.Parse("2006-01-02", "1992-01-01")
		ok, err := m.InsertAccount(&Account{
			Email:      "hatlonely1@foxmail.com",
			Phone:      "13112345678",
			Password:   "123456",
			FirstName:  "孙",
			SecondName: "悟空",
			Birthday:   birthday,
			Gender:     1,
		})
		So(err, ShouldBeNil)
		So(ok, ShouldBeTrue)

		Convey("select account by phone", func() {
			account, err := m.SelectAccountByPhoneOrEmail("hatlonely1@foxmail.com")
			So(err, ShouldBeNil)
			So(account.Email, ShouldEqual, "hatlonely1@foxmail.com")
			So(account.Phone, ShouldEqual, "13112345678")
			So(account.Password, ShouldEqual, "123456")
			So(account.FirstName, ShouldEqual, "孙")
			So(account.SecondName, ShouldEqual, "悟空")
			So(account.Birthday, ShouldEqual, birthday)
			So(account.Gender, ShouldEqual, 1)
		})

		Convey("select account by email", func() {
			account, err := m.SelectAccountByPhoneOrEmail("13112345678")
			So(err, ShouldBeNil)
			So(account.Email, ShouldEqual, "hatlonely1@foxmail.com")
			So(account.Phone, ShouldEqual, "13112345678")
			So(account.Password, ShouldEqual, "123456")
			So(account.FirstName, ShouldEqual, "孙")
			So(account.SecondName, ShouldEqual, "悟空")
			So(account.Birthday, ShouldEqual, birthday)
			So(account.Gender, ShouldEqual, 1)
		})
	})
}
