package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	ent "main/internal/auth/domain/entities"
	"main/internal/auth/domain/ports"
	cm "main/pkg/utils/common"
)

// postgresqlRepo Struct
type postgresqlRepo struct {
	db *pgxpool.Pool
}

// NewPostgresqlRepository Auth Domain postgresql repository constructor
func NewPostgresqlRepository(db *pgxpool.Pool) ports.IPostgresqlRepository {
	return &postgresqlRepo{db: db}
}

// Register user Registration function
func (r *postgresqlRepo) Register(ctx context.Context, req ent.RegisterReq) (record int64) {
	record = 1
	var query string
	var errDb error
	var auth_id int64

	if req.UserType == 1 {
		query = `INSERT INTO human_resources.auth (src, company_name, user_title, user_name, user_pass) VALUES ($1,$2,$3,$4,$5) RETURNING auth_id`
		errDb = r.db.QueryRow(ctx, query, req.Src, req.CompanyName, req.UserTitle, req.UserName, req.UserPass).Scan(&auth_id)
	} else {
		query = `INSERT INTO candidate.auth (src, user_title, user_name, user_pass) VALUES ($1,$2,$3,$4) RETURNING auth_id`
		errDb = r.db.QueryRow(ctx, query, req.Src, req.UserTitle, req.UserName, req.UserPass).Scan(&auth_id)
	}

	if errDb != nil && cm.CheckStringIfContains(errDb.Error(), "duplicate key value violates") == false {
		fmt.Printf("Register DB error : %s \n", errDb.Error())
		record = -1
	} else if errDb != nil && cm.CheckStringIfContains(errDb.Error(), "duplicate key value violates") == true {
		fmt.Printf("Duplicate error on Register : %s", errDb.Error())
		record = -3
	} else {
		r.SaveContact(ctx, auth_id, req)
	}

	return
}

// Login Login with user_name and password param
func (r *postgresqlRepo) Login(ctx context.Context, req ent.LoginReq) (record int64, auth ent.Auth) {
	record = 1
	var query string
	var errDb error

	if req.UserType == 1 {
		query = `SELECT auth_id, lang, parent_id, user_type, acc_type, COALESCE(company_name,''), user_title, user_name, is_demo, unique_id, status FROM human_resources.auth WHERE user_name=$1 AND user_pass=$2`
		errDb = r.db.QueryRow(ctx, query, req.UserName, req.UserPass).Scan(&auth.AuthId, &auth.Lang, &auth.ParentId, &auth.UserType, &auth.AccountType, &auth.CompanyName, &auth.UserTitle, &auth.UserName, &auth.IsDemo, &auth.UniqueId, &auth.Status)
	} else {
		query = `SELECT auth_id, lang, manager_id, user_type, user_title, user_name, unique_id, status FROM candidate.auth WHERE user_name=$1 AND user_pass=$2`
		errDb = r.db.QueryRow(ctx, query, req.UserName, req.UserPass).Scan(&auth.AuthId, &auth.Lang, &auth.ManagerId, &auth.UserType, &auth.UserTitle, &auth.UserName, &auth.UniqueId, &auth.Status)
	}

	if errDb != nil && cm.CheckStringIfContains(errDb.Error(), "no rows in result set") == false {
		fmt.Printf("Login Err : %s \n", errDb.Error())
		record = -1
	} else if errDb != nil && cm.CheckStringIfContains(errDb.Error(), "no rows in result set") == true {
		record = 0
	}

	return
}

// Login Login with user_name and password param
func (r *postgresqlRepo) SaveContact(ctx context.Context, auth_id int64, req ent.RegisterReq) (record int64) {
	record = 1
	var query string
	var errDb error

	if req.UserType == 1 {
		query = `INSERT INTO human_resources.auth_contact(auth_id, type_of, description, definition, is_default) VALUES ($1,$2,$3,$4,$5)`
		_, errDb = r.db.Exec(ctx, query, auth_id, 1, "Mobile phone number on register", req.UserPhone, true)
	} else {
		query = `INSERT INTO candidate.auth_contact(auth_id, type_of, description, definition, is_default) VALUES ($1,$2,$3,$4,$5)`
		_, errDb = r.db.Exec(ctx, query, auth_id, 1, "Mobile phone number on register", req.UserPhone, true)
	}

	if errDb != nil && cm.CheckStringIfContains(errDb.Error(), "no rows in result set") == false {
		fmt.Printf("Save Contact Error : %s \n", errDb.Error())
		record = -1
	} else if errDb != nil && cm.CheckStringIfContains(errDb.Error(), "no rows in result set") == true {
		record = 0
	}

	return
}
