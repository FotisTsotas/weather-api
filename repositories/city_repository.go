package repositories

import (
	"database/sql"
)

type City struct {
	ID        int
	Name      string
	Latitude  float64
	Longitude float64
}

type CityRepository struct {
	db *sql.DB
}

func NewCityRepository(db *sql.DB) *CityRepository {
	return &CityRepository{db: db}
}

func (r *CityRepository) GetCities() ([]City, error) {
	rows, err := r.db.Query("SELECT id, name, latitude, longitude FROM cities")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cities []City
	for rows.Next() {
		var city City
		if err := rows.Scan(&city.ID, &city.Name, &city.Latitude, &city.Longitude); err != nil {
			return nil, err
		}
		cities = append(cities, city)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return cities, nil
}

func (r *CityRepository) GetCityByID(id int) (*City, error) {
	row := r.db.QueryRow("SELECT id, name, latitude, longitude FROM cities WHERE id = ?", id)
	var city City
	if err := row.Scan(&city.ID, &city.Name, &city.Latitude, &city.Longitude); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &city, nil
}

func (r *CityRepository) CreateCity(name string, latitude, longitude float64) (*City, error) {
	result, err := r.db.Exec("INSERT INTO cities (name, latitude, longitude) VALUES (?, ?, ?)", name, latitude, longitude)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &City{ID: int(id), Name: name, Latitude: latitude, Longitude: longitude}, nil
}

func (r *CityRepository) UpdateCity(id int, name string, latitude, longitude float64) (*City, error) {
	_, err := r.db.Exec("UPDATE cities SET name = ?, latitude = ?, longitude = ? WHERE id = ?", name, latitude, longitude, id)
	if err != nil {
		return nil, err
	}
	return &City{ID: id, Name: name, Latitude: latitude, Longitude: longitude}, nil
}

func (r *CityRepository) DeleteCity(id int) error {
	_, err := r.db.Exec("DELETE FROM cities WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
