//go:build cgo

package dbtests

import (
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"os"
	"testing"

	"github.com/datrine/basic_crud_with_auth/internal/db"
	sharedexports "github.com/datrine/basic_crud_with_auth/internal/shared_exports"
	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/assert"
	_ "github.com/tursodatabase/go-libsql"
)

func TestCreateUser(t *testing.T) {
	st := &db.SqlLiteSQLDB{
		DB: setupUserTableForTest(),
	}
	hash := md5.New()
	by := hash.Sum([]byte("password"))
	err := st.CreateUser(context.TODO(), &sharedexports.CreateUser{
		Email:        "test@testing.co",
		LastName:     "alabi",
		FirstName:    "temitope",
		PasswordHash: string(hex.EncodeToString(by)),
	})
	assert.NoError(t, err)
	t.Cleanup(func() {
		cleanUserTableAfterTest(st.DB)
		removeLocalDB()
		fmt.Println("Cleanup done")
	})
}

func TestUpdateUser(t *testing.T) {
	email := "test@email.co"
	d := setupUserTableForTest()
	fillUserTableWithDummyForTest(d)
	sqlDB := &db.SqlLiteSQLDB{
		DB: d,
	}
	t.Run("normal update", func(t1 *testing.T) {
		defer func() {
			t1.Cleanup(func() {
				d.ExecContext(context.Background(), "DELETE FROM users WHERE email = @email ", sql.Named("email", email))
				cleanUserTableAfterTest(d)

				fmt.Println("Cleanup done")
			})
		}()
		err := sqlDB.UpdateUser(context.Background(), email, &sharedexports.UpdateUser{
			FirstName: "johnny",
			LastName:  "joe",
		})
		/*
			d.Ping()
			_, err := d.ExecContext(context.Background(), "UPDATE users SET first_name = @first_name WHERE email = @email",
				[]any{sql.Named("first_name", "johnny"), sql.Named("email", "test@email.co")}...)

		*/
		assert.NoError(t1, err)

		userFromDb, err := sqlDB.QueryUserByEmail(context.Background(), email)
		assert.NoError(t1, err)
		assert.Equal(t1, "johnny", userFromDb.FirstName)
		assert.Equal(t1, "joe", userFromDb.LastName)
	})
	func() {

		t.Cleanup(func() {
			removeLocalDB()
			fmt.Println("All Cleanup done")
		})
	}()
}

func setupUserTableForTest() *sql.DB {
	dbName := "file:./local.db"
	d, err := goose.OpenDBWithDriver("turso", dbName)
	if err != nil {
		panic(err.Error())
	}
	err = goose.Up(d, "../migrations/sqlite")
	if err != nil {
		panic(err.Error())
	}
	return d
}

func fillUserTableWithDummyForTest(d *sql.DB) {
	if err := d.Ping(); err != nil {
		panic(err.Error())
	}
	hash := md5.New()
	by := hash.Sum([]byte("password"))
	password_hash := hex.EncodeToString(by)
	result, err := d.ExecContext(context.Background(), `
	INSERT INTO users (email,last_name, first_name,password_hash) 
	VALUES (@email, @last_name, @first_name, @password_hash ) `,
		sql.Named("email", "test@email.co"),
		sql.Named("last_name", "alabi"),
		sql.Named("first_name", "temitope"),
		sql.Named("password_hash", password_hash))
	if err != nil {
		panic(err.Error())
	}
	affected, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	if affected != 1 {
		//panic(fmt.Errorf("Number of rows affected: %d", affected))
	}
	fmt.Println("Number of rows affected: ", affected)
}

func cleanUserTableAfterTest(d *sql.DB) {
	err := goose.Down(d, "../migrations/sqlite")
	if err != nil {
		panic(err.Error())
	}
	tx, err := d.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	tx.ExecContext(context.TODO(), "DROP TABLE IF EXISTS users")
	err = tx.Commit()
	if err != nil {
		panic(err.Error())
	}
	err = d.Close()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Cleanup done")
}

func removeLocalDB() {
	err := os.Remove("./local.db")
	if err != nil {
		panic(err.Error())
	}
}
