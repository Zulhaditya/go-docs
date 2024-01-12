package belajar_golang_viper

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var config *viper.Viper = viper.New()

func TestViper(t *testing.T) {
	assert.NotNil(t, config)
}

func TestJSON(t *testing.T) {
	config := viper.New()
	config.SetConfigName("config") // nama file config
	config.SetConfigType("json")   // extention filenya
	config.AddConfigPath(".")

	// membaca config
	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "belajar-golang-viper", config.GetString("app.name"))
	assert.Equal(t, "Zulhaditya", config.GetString("app.author"))
	assert.Equal(t, "localhost", config.GetString("database.host"))
	assert.Equal(t, 3306, config.GetInt("database.port"))
	assert.Equal(t, true, config.GetBool("database.show_sql"))
}

func TestYAML(t *testing.T) {
	config := viper.New()
	config.SetConfigFile("config.yaml")
	config.AddConfigPath(".")

	// membaca config
	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "belajar-golang-viper", config.GetString("app.name"))
	assert.Equal(t, "Zulhaditya", config.GetString("app.author"))
	assert.Equal(t, "localhost", config.GetString("database.host"))
	assert.Equal(t, 3306, config.GetInt("database.port"))
	assert.Equal(t, true, config.GetBool("database.show_sql"))
}

func TestENV(t *testing.T) {
	config := viper.New()
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")

	// membaca config
	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "belajar-golang-viper", config.GetString("APP_NAME"))
	assert.Equal(t, "Zulhaditya", config.GetString("APP_AUTHOR"))
	assert.Equal(t, 3306, config.GetInt("DATABASE_PORT"))
	assert.Equal(t, true, config.GetBool("DATABASE_SHOW_SQL"))
}

func TestENVLocal(t *testing.T) {
	config := viper.New()
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")
	config.AutomaticEnv() // membaca env local

	// membaca config
	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "belajar-golang-viper", config.GetString("APP_NAME"))
	assert.Equal(t, "Zulhaditya", config.GetString("APP_AUTHOR"))
	assert.Equal(t, 3306, config.GetInt("DATABASE_PORT"))
	assert.Equal(t, true, config.GetBool("DATABASE_SHOW_SQL"))

	// export FROM_ENV=Hello, perintah di terminal untuk menambahkan env local
	assert.Equal(t, "Hello", config.GetString("FROM_ENV"))
}
