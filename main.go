package main

import (
	"bytes"
	"log"
	"os"

	"github.com/jackpal/bencode-go"
	"github.com/nohajc/go-qbittorrent/qbt"
	"github.com/spf13/cast"
)

var (
	host = "http://192.168.1.2:8080"
	user = "admin"
	pass = "admin"
)

func main() {
	qb := qbt.NewClient(host)
	err := qb.Login(qbt.LoginOptions{
		Username: user,
		Password: pass,
	})
	if err != nil {
		panic(err)
	}

	torrents, err := qb.Torrents(nil)
	if err != nil {
		panic(err)
	}

	for _, torrent := range torrents {
		// 检查种子是否存在
		if !PathExists("./BT_backup/" + torrent.Hash + ".torrent") {
			continue
		}
		// 读取种子
		rt, err := ReadTorrent(torrent.Hash)
		if err != nil {
			log.Fatalln(torrent.Hash, " 读取失败: ", err)
		}
		// 写入种子
		err = WriteTorrent(torrent.Hash, torrent.Tracker, rt)
		if err != nil {
			log.Fatalln(torrent.Hash, " 写入失败: ", err)
		}
	}
}

func ReadTorrent(hash string) (map[string]interface{}, error) {
	r, err := os.Open("./BT_backup/" + hash + ".torrent")
	if err != nil {
		return nil, err
	}
	data, err := bencode.Decode(r)
	r.Close()
	if err != nil {
		return nil, err
	}

	return cast.ToStringMap(data), nil
}

func WriteTorrent(hash, tracker string, torrentMap map[string]interface{}) error {
	torrentMap["announce"] = tracker
	var buf bytes.Buffer
	err := bencode.Marshal(&buf, torrentMap)
	if err != nil {
		return err
	}
	err = os.WriteFile("./export/"+hash+".torrent", buf.Bytes(), 0666)
	if err != nil {
		return err
	}

	return nil
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return false
}
