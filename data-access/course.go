package db

import (
	"context"
	"fmt"
	"gocourseCRUD/cockroach"
	"gocourseCRUD/exceptions"
	"gocourseCRUD/model"
	"log"
	"strings"

	"github.com/georgysavva/scany/v2/pgxscan"
)

const TABLE_NAME = "courses"

func AddCourse(ctx context.Context, name string, platform string, price string) (string, error) {
	query := "INSERT INTO " + TABLE_NAME + " (name, platform, price) VALUES ($1, $2, $3) RETURNING id"

	result, err := cockroach.GetConnectionPool(ctx).Query(ctx, query, name, platform, price)

	if err != nil {
		log.Println(err)
		return "", exceptions.UnknownError("", "")
	}
	result.Next()

	var id string
	result.Scan(&id)

	return id, nil
}

func GetCourseByName(ctx context.Context, name string) ([]model.Course, error) {
	query := "SELECT * FROM " + TABLE_NAME + " WHERE name = $1"

	rows, err := cockroach.GetConnectionPool(ctx).Query(ctx, query, name)

	if err != nil {
		fmt.Println(err)
		return nil, exceptions.UnknownError("", "")
	}

	var response []model.Course
	if err := pgxscan.ScanAll(&response, rows); err != nil {
		log.Println(err)
		return nil, exceptions.UnknownError("", "")
	}
	return response, nil
}

func GetCourse(ctx context.Context, id string) ([]model.Course, error) {
	query := "SELECT * FROM " + TABLE_NAME + " WHERE id = $1"

	rows, err := cockroach.GetConnectionPool(ctx).Query(ctx, query, id)

	if err != nil {
		fmt.Println(err)
		return nil, exceptions.UnknownError("", "")
	}

	var response []model.Course

	if err := pgxscan.ScanAll(&response, rows); err != nil {
		log.Println(err)
		return nil, exceptions.UnknownError("", "")
	}
	return response, nil
}

func GetAllCouseList(ctx context.Context) ([]model.Course, error) {
	query := "SELECT * FROM " + TABLE_NAME

	rows, err := cockroach.GetConnectionPool(ctx).Query(ctx, query)

	if err != nil {
		fmt.Println(err)
		return nil, exceptions.UnknownError("", "")
	}

	var response []model.Course

	if err := pgxscan.ScanAll(&response, rows); err != nil {
		log.Println(err)
		return nil, exceptions.UnknownError("", "")
	}
	return response, nil
}

func DeleteCourse(ctx context.Context, id string) error {
	query := "DELETE FROM " + TABLE_NAME + " WHERE id = $1"

	_, err := cockroach.GetConnectionPool(ctx).Exec(ctx, query, id)

	if err != nil {
		log.Println(err)
		return exceptions.UnknownError("", "")
	}
	return nil
}

func UpdateCourse(ctx context.Context, id string, courseDetail model.Course) error {
	count := 1
	var prepareStatementValue []string
	var value []interface{}
	if courseDetail.Name != "" {
		prepareStatementValue = append(prepareStatementValue, "name = $"+fmt.Sprint(count))
		value = append(value, courseDetail.Name)
		count++
	}
	if courseDetail.Platform != "" {
		prepareStatementValue = append(prepareStatementValue, "platform = $"+fmt.Sprint(count))
		value = append(value, courseDetail.Platform)
		count++
	}
	if courseDetail.Price != "" {
		prepareStatementValue = append(prepareStatementValue, "price = $"+fmt.Sprint(count))
		value = append(value, courseDetail.Price)
		count++
	}

	value = append(value, id)

	updateQuery := "UPDATE " + TABLE_NAME + " SET " + strings.Join(prepareStatementValue, ", ") + " WHERE id = $" + fmt.Sprint(count)

	_, err := cockroach.GetConnectionPool(ctx).Exec(ctx, updateQuery, value...)

	if err != nil {
		log.Println(err)
		return exceptions.UnknownError("", "")
	}

	return nil

}
