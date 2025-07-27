package services

import (
	"context"
	"time"

	"github.com/amaxyza/shadro/models"
	"github.com/jackc/pgx/v5"
)

func CreateProgram(owner_id int, name, source string) (*models.Shader, error) {
	// Create shader entry in postgresql database
	var program_id int
	var created_at, updated_at time.Time

	err := pool.QueryRow(
		context.Background(),
		`INSERT INTO program (owner_id, name, glsl) VALUES (@ownerid, @programname, @glslsource) RETURNING id, created_at, updated_at`,
		pgx.NamedArgs{
			"ownerid":     owner_id,
			"programname": name,
			"glslsource":  source,
		},
	).Scan(&program_id, &created_at, &updated_at)

	if err != nil {
		return nil, err
	}

	// Return shader struct
	shader := models.NewShader(program_id, owner_id, name, source, created_at, updated_at)

	return &shader, nil
}

func GetProgram(program_id int) (models.Shader, error) {
	var id, owner_id int
	var name, source string
	var created_at, updated_at time.Time

	err := pool.QueryRow(
		context.Background(),
		`SELECT * FROM program WHERE id = @programid`,
		pgx.NamedArgs{
			"programid": program_id,
		},
	).Scan(&id, &owner_id, &name, &source, &created_at, &updated_at)

	if err != nil {
		return nil, err
	}

	shader := models.NewShader(id, owner_id, name, source, created_at, updated_at)

	return shader, nil
}

func GetAllUserPrograms(user_id int) ([]models.Shader, error) {
	rows, err := pool.Query(
		context.Background(),
		`SELECT * FROM program WHERE owner_id = @user_id`,
		pgx.NamedArgs{"user_id": user_id},
	)

	if err != nil {
		return nil, err
	}

	var programs_slice []models.Shader

	scans := []any{new(int), new(int), new(string), new(string), new(time.Time), new(time.Time)}
	_, err = pgx.ForEachRow(rows, scans, func() error {
		var id, owner_id int
		var name, source string
		var created_at, updated_at time.Time

		err := rows.Scan(&id, &owner_id, &name, &source, &created_at, &updated_at)
		if err != nil {
			return err
		}

		programs_slice = append(programs_slice, models.NewShader(id, owner_id, name, source, created_at, updated_at))

		return nil
	})

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return programs_slice, nil
}

func UpdateProgram(shaderprogram models.Shader, name, source string) error {
	_, err := pool.Exec(
		context.Background(),
		`UPDATE program SET name = @newname, glsl = @newsrc, updated_at = NOW() WHERE id = @pid`,
		pgx.NamedArgs{
			"newname": name,
			"newsrc":  source,
			"pid":     shaderprogram.GetID(),
		},
	)

	return err
}

func DeleteProgram(id int) error {
	_, err := pool.Exec(
		context.Background(),
		`DELETE FROM program WHERE id = @pid`,
		pgx.NamedArgs{
			"pid": id,
		},
	)

	return err
}
