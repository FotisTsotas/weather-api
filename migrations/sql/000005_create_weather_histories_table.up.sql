CREATE TABLE IF NOT EXISTS weather_histories (
    id SERIAL PRIMARY KEY,
    city_id INT NOT NULL REFERENCES cities(id) ON DELETE CASCADE,
    temperature DOUBLE NOT NULL,
    humidity DOUBLE NOT NULL,
    wind_speed DOUBLE NOT NULL,
    recorded_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX (recorded_at)
);