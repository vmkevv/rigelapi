package handlers

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"github.com/vmkevv/rigelapi/ent"
)

/* handler expected error */

type ClientErr struct {
	Status  int
	Message string
}

func (c ClientErr) Error() string {
	return c.Message
}
func NewClientErr(code int, msg string) ClientErr {
	return ClientErr{code, msg}
}

/* Data included in the token */

type AppClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

/* response of any sync request */

type SyncIDResp struct {
	LastSyncID string `json:"last_sync_id"`
}

/* request of any sync request */

type DBOperation = string

const (
	Insert DBOperation = "INSERT"
	Update             = "UPDATE"
	Delete             = "DELETE"
)

type SyncReqBase struct {
	ID       string      `json:"id"`
	DateTime int64       `json:"date_time"`
	Type     DBOperation `json:"type"`
}

/* rollback a transaction */

func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
