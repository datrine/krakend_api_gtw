package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/datrine/basic_crud_with_auth/config"
	sharedexports "github.com/datrine/basic_crud_with_auth/internal/shared_exports"
	"github.com/pressly/goose/v3"
	"github.com/tursodatabase/go-libsql"
)

var dir string
var connector *libsql.Connector
var db *sql.DB

type SqlLiteSQLDB struct {
	DB *sql.DB
}

func init() {
	//Connect()
}

func Migrate(db *sql.DB, dir string) {

	goose.Up(db, dir)
}

func Connect() {
	var err error
	dbName := config.GetLocalDBName()
	primaryUrl := config.GetTursoDatabaseURL()
	authToken := config.GetTursoAuthToken()
	dir, err = os.MkdirTemp("", "libsql-*")
	if err != nil {
		fmt.Println("Error creating temporary directory: ", err.Error())
		os.Exit(1)
	}
	dbPath := filepath.Join(dir, dbName)
	connector, err = libsql.NewEmbeddedReplicaConnector(dbPath, primaryUrl, libsql.WithAuthToken(authToken))
	if err != nil {
		fmt.Println("Error creating connector:", err)
		os.Exit(1)
	}

	db = sql.OpenDB(connector)
}

func SetupRemoteDbLocalReplicatorsConnection(dbName, primaryUrl, authToken string) (*sql.DB, error) {
	var err error
	dir, err = os.MkdirTemp("", "libsql-*")
	if err != nil {
		fmt.Println("Error creating temporary directory: ", err.Error())
		os.Exit(1)
	}
	dbPath := filepath.Join(dir, dbName)
	connector, err = libsql.NewEmbeddedReplicaConnector(dbPath, primaryUrl, libsql.WithAuthToken(authToken))
	if err != nil {
		fmt.Println("Error creating connector:", err)
		os.Exit(1)
	}
	db := sql.OpenDB(connector)
	Migrate(db, "./migrations/sqlite")
	return db, nil
}

func SetupLocalDbConnection(dbName string) (*sql.DB, error) {
	db, err := sql.Open("libsql", dbName)
	if err != nil {
		return nil, err
	}
	Migrate(db, "./migrations/sqlite")
	return db, nil
}

func SetupRemoteDbConnection(dbName, token string) (*sql.DB, error) {
	url := strings.Join([]string{config.GetTursoDatabaseURL() + "?token", token}, "=")

	db, err := sql.Open("libsql", url)
	if err != nil {
		return nil, err
	}
	Migrate(db, "./migrations/sqlite")
	return db, nil
}

func NewRemoteSqlLiteSQLDB() *SqlLiteSQLDB {
	primaryUrl := config.GetTursoDatabaseURL()
	authToken := config.GetTursoAuthToken()
	db, err := SetupRemoteDbLocalReplicatorsConnection("", primaryUrl, authToken)
	if err != nil {
		panic(err.Error())
	}
	return &SqlLiteSQLDB{
		DB: db,
	}
}

func NewLocalSqlLiteSQLDB(dbName string) *SqlLiteSQLDB {
	db, err := SetupLocalDbConnection(dbName)
	if err != nil {
		panic(err.Error())
	}
	return &SqlLiteSQLDB{
		DB: db,
	}
}

func Cleanup() {
	os.RemoveAll(dir)
	connector.Close()
	db.Close()
}

func (d *SqlLiteSQLDB) CreateUserTable(ctx context.Context) error {
	result, err := d.DB.ExecContext(ctx, `CREATE TABLE IF NOT EXIST users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email VARCHAR(50) UNIQUE NOT NULL , password_hash VARCHAR(500) NOT NULL, 
	last_name VARCHAR(50) NOT NULL, 
	first_name,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	created_at TIMESTAMP
	);
	CREATE TRIGGER IF NOT EXISTS update_on_trigger AFTER UPDATE ON users
	BEGIN
		UPDATE users SET created_at = datetime() WHERE email = OLD.email; 
	END;
	`)
	if err != nil {
		return err
	}
	rowId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	fmt.Println("%s", rowId)
	return nil
}

func (d *SqlLiteSQLDB) PragmaUser(ctx context.Context, userTocreate *sharedexports.CreateUser) error {
	rows, err := d.DB.Query("PRAGMA table_info('users')")
	if err != nil {
		return err
	}

	var uu []map[string]interface{}
	for {
		if rows.Next() {
			var res map[string]interface{}
			res = make(map[string]interface{})
			cid := ""
			name := ""
			type_n := ""
			notnull := ""
			var dflt_value interface{}
			pk := ""
			var dest []any = []any{&cid, &name, &type_n, &notnull, &dflt_value, &pk}
			err = rows.Scan(dest...)
			if err != nil {
				return err
			}
			uu = append(uu, res)
		} else {
			print(uu)
			break
		}
	}
	return nil
}

func (d *SqlLiteSQLDB) CreateUser(ctx context.Context, userTocreate *sharedexports.CreateUser) error {
	result, err := d.DB.ExecContext(ctx, `INSERT INTO users 
	    ( email,password_hash, last_name, first_name) 
	VALUES
	    (?,?,?,?)`,
		userTocreate.Email,
		userTocreate.PasswashHash,
		userTocreate.LastName,
		userTocreate.FirstName)
	if err != nil {
		return err
	}
	rowId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	fmt.Printf("%d\n", rowId)
	return nil
}

func (d *SqlLiteSQLDB) QueryUserByEmail(ctx context.Context, email string) (*sharedexports.User, error) {
	user := &sharedexports.User{}
	emailR := ""
	lastNameR := ""
	firstNameR := ""
	passwordHashR := ""
	var createdAtR time.Time
	var updatedAtR sql.NullTime
	rd := []interface{}{&emailR, &lastNameR, &firstNameR, &createdAtR, &updatedAtR, &passwordHashR}
	row := d.DB.QueryRowContext(ctx, "SELECT * FROM users WHERE email=@email", sql.Named("email", email))
	err := row.Scan(rd...)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("no user with email %s\n", email)
			return nil, nil
		}
		return nil, err
	}
	user.Email = email
	user.LastName = lastNameR
	user.FirstName = firstNameR
	user.PasswashHash = passwordHashR
	user.CreatedAt = createdAtR
	user.UpdatedAt = updatedAtR.Time
	return user, nil
}

func (d *SqlLiteSQLDB) UpdateUser(ctx context.Context, email string, updates *sharedexports.UpdateUser) error {
	// important to downcast each item as interface
	sqlNamedList := []any{}
	updateSql := ""
	if updates.FirstName != "" {
		if updateSql == "" {
			updateSql = "first_name = @first_name"
		} else {
			updateSql = strings.Join([]string{updateSql, "first_name = @first_name"}, ", ")
		}
		sqlNamedList = append(sqlNamedList, sql.Named("first_name", updates.FirstName))
	}
	if updates.LastName != "" {
		if updateSql == "" {
			updateSql = "last_name = @last_name"
		} else {
			updateSql = strings.Join([]string{updateSql, "last_name = @last_name"}, ", ")
		}
		sqlNamedList = append(sqlNamedList, sql.Named("last_name", updates.LastName))
	}
	sqlNamedList = append(sqlNamedList, sql.Named("email", email))
	result, err := d.DB.ExecContext(context.Background(), "UPDATE users SET "+updateSql+" WHERE email = @email;",
		sqlNamedList...)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Printf("Rows affected: %v\n", rowsAffected)
	return nil
}

func (d *SqlLiteSQLDB) DeleteUserByEmail(ctx context.Context, email string) (*sharedexports.User, error) {
	user := &sharedexports.User{}
	row, err := d.DB.ExecContext(ctx, "DELETE FROM users WHERE email=@email", sql.Named("email", email))
	if err != nil {
		return nil, err
	}
	rowsAffected, err := row.RowsAffected()
	if err != nil {
		return nil, err
	}
	fmt.Println("Rows affected: %v\n", rowsAffected)
	return user, nil
}
