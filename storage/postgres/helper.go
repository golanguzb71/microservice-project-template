package postgres

import (
	"database/sql"
	"fmt"
	pb "github.com/golanguzb71/microservice-project-template/genproto/template_service"
	"github.com/golanguzb71/microservice-project-template/pkg/logger"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleDatabaseError(err error, log logger.Logger, message string) error {
	if err == nil {
		return nil
	}
	log.Error(message + ": " + err.Error())
	switch err {
	case sql.ErrNoRows:
		return status.Error(codes.NotFound, "This information does not exists.")
	case sql.ErrConnDone:
		return err
	case sql.ErrTxDone:
		return err
	}

	switch e := err.(type) {
	case *pq.Error:
		// Handle Postgres-specific errors
		switch e.Code.Name() {
		case "unique_violation":
			return status.Error(codes.AlreadyExists, "Already exists")
		case "foreign_key_violation":
			return status.Error(codes.InvalidArgument, "The elements related to this element should be deleted first")
		default:
			return err
		}
	default:
		// Handle all other errors
		return err
	}
}

func PrepareWhere(filters []*pb.Filters) (squirrel.And, bool) {
	res := squirrel.And{}
	search := squirrel.Or{}

	for _, e := range filters {
		switch e.Type {
		case "search":
			search = append(search, squirrel.ILike{e.Field: "%" + e.Value + "%"})
		case "=":
			res = append(res, squirrel.Eq{e.Field: e.Value})
		case "!=":
			res = append(res, squirrel.NotEq{e.Field: e.Value})
		case "<=":
			res = append(res, squirrel.LtOrEq{e.Field: e.Value})
		case "<":
			res = append(res, squirrel.Lt{e.Field: e.Value})
		case ">=":
			res = append(res, squirrel.GtOrEq{e.Field: e.Value})
		case ">":
			res = append(res, squirrel.Gt{e.Field: e.Value})
		}
	}
	if len(search) > 0 {
		res = append(res, search)
	}

	return res, len(res) > 0
}

func PrepareOrder(orders []*pb.SortBy) (string, bool) {
	res := []string{}
	for _, e := range orders {
		switch e.Type {
		case "desc", "asc":
			res = append(res, fmt.Sprintf("%s %s", e.Field, e.Type))
		}
	}

	return strings.Join(res, ", "), len(res) > 0
}
