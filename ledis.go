package main

import (
	"os"
	"path/filepath"

	"github.com/siddontang/ledisdb/config"
	"github.com/siddontang/ledisdb/ledis"
)

type LedisInfo struct {
	Name     string
	Path     string
	DbFolder string
	Conn     *ledis.Ledis
	Db       *ledis.DB
}

func getLedisInfo(dir string) (*LedisInfo, error) {
	db := &LedisInfo{Name: "db", DbFolder: "db"}
	if err := db.init(dir); err != nil {
		return nil, err
	}

	return db, nil
}

func (l *LedisInfo) init(databasePath string) error {
	fullPath := filepath.Join(databasePath, l.DbFolder)

	if err := os.MkdirAll(fullPath, 0700); err != nil {
		return err
	}

	l.Path = fullPath
	cfg := &config.Config{
		DataDir: fullPath,
	}

	conn, err := ledis.Open(cfg)
	if err != nil {
		return err
	}
	l.Conn = conn

	db, err := conn.Select(0)
	if err != nil {
		return err
	}
	l.Db = db

	return nil
}

func (l *LedisInfo) GetKeyList(dataType ledis.DataType) ([]string, error) {
	batchSize := 10
	var lastKey []byte
	result := []string{}

	lastKey = nil
	for true {
		keyBatch, err := l.Db.Scan(dataType, lastKey, batchSize, false, "")
		if err != nil {
			return nil, err
		}

		for _, k := range keyBatch {
			lastKey = k
			result = append(result, string(k))
		}

		if len(keyBatch) < batchSize {
			break
		}
	}

	return result, nil
}
