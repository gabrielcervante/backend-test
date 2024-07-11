package repo

import (
	"database/sql"
)

func Scan[T any](rows *sql.Rows, inputErr error, entityFields func(breeds *T) []any) (output []T, err error) {
	var entity T

	if inputErr != nil {
		return nil, inputErr
	}

	for rows.Next() {
		if err = rows.Scan(entityFields(&entity)...); err != nil {
			return
		}

		output = append(output, entity)
	}

	return
}
