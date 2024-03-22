package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ledisdb/ledisdb/ledis"
)

type Command struct {
	Text string

	query string
}

func (cmd *Command) Execute(ledisInfo *LedisInfo) error {
	text := strings.ToLower(cmd.Text)
	if text == "keys" {
		return cmd.runKeys(ledisInfo)
	}

	cmd.parseQuery()

	if strings.HasPrefix(text, "hgetall") {
		return cmd.runHGetAll(ledisInfo)
	}
	if strings.HasPrefix(text, "hget") {
		return cmd.runHGet(ledisInfo)
	}
	if strings.HasPrefix(text, "smembers") {
		return cmd.runSmembers(ledisInfo)
	}
	if strings.HasPrefix(text, "get") {
		return cmd.runGet(ledisInfo)
	}
	if strings.HasPrefix(text, "llen") {
		return cmd.runLlen(ledisInfo)
	}
	if strings.HasPrefix(text, "lrange") {
		return cmd.runLrange(ledisInfo)
	}

	return fmt.Errorf("unknown command: %s\n", cmd.Text)
}

func (cmd *Command) parseQuery() {
	parts := strings.SplitN(cmd.Text, " ", 2)
	if len(parts) == 2 {
		cmd.query = parts[1]
	}
}

func (cmd *Command) runKeys(ledisInfo *LedisInfo) error {
	keyTypes := []ledis.DataType{
		ledis.KV, ledis.SET, ledis.LIST, ledis.HASH, ledis.ZSET,
	}
	for _, keyType := range keyTypes {
		keys, err := ledisInfo.GetKeyList(keyType)
		if err != nil {
			return err
		}

		fmt.Printf("======keys of type %v\n", keyType)
		for _, key := range keys {
			fmt.Printf("%s\n", key)
		}
	}

	return nil
}

func (cmd *Command) runHGetAll(ledisInfo *LedisInfo) error {
	vals, err := ledisInfo.Db.HGetAll([]byte(cmd.query))
	if err != nil {
		return err
	}
	for _, val := range vals {
		fmt.Printf("=======field:\n")
		fmt.Printf("%s:%s\n", val.Field, val.Value)
	}
	return nil
}

func (cmd *Command) runHGet(ledisInfo *LedisInfo) error {
	return errors.New("not implemented")
}

func (cmd *Command) runSmembers(ledisInfo *LedisInfo) error {
	vals, err := ledisInfo.Db.SMembers([]byte(cmd.query))
	if err != nil {
		return err
	}
	for _, val := range vals {
		fmt.Printf("%s\n", val)
	}
	return nil
}

func (cmd *Command) runGet(ledisInfo *LedisInfo) error {
	val, err := ledisInfo.Db.Get([]byte(cmd.query))
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", val)
	return nil
}

func (cmd *Command) runLlen(ledisInfo *LedisInfo) error {
	val, err := ledisInfo.Db.LLen([]byte(cmd.query))
	if err != nil {
		return err
	}
	fmt.Printf("%d\n", val)
	return nil
}

func (cmd *Command) runLrange(ledisInfo *LedisInfo) error {
	start := int32(0)
	end := int32(0)

	parts := strings.Split(cmd.query, " ")
	if len(parts) == 3 {
		s, err := strconv.Atoi(parts[1])
		if err != nil {
			return err
		}
		start = int32(s)
		e, err := strconv.Atoi(parts[2])
		if err != nil {
			return err
		}
		end = int32(e)
	} else {
		return fmt.Errorf("command should have 3 parts (these are single space delimited)")
	}

	vals, err := ledisInfo.Db.LRange([]byte(parts[0]), start, end)
	if err != nil {
		return err
	}

	for _, val := range vals {
		fmt.Printf("%s\n", val)
	}

	return nil
}
