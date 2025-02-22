package serverinfo

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"strings"
	"time"
)

type ServerInfo struct {
	Address     string
	Port        int
	Host        string
	Map         string
	Players     int
	Waves       int
	GameVersion int
	VerType     string
	Gamemode    string
	Limit       int
	Desc        string
	Latency     int
}

func (si *ServerInfo) Update() error {
	var err error

	start := time.Now()
	conn, err := net.Dial("udp4", net.JoinHostPort(si.Address, fmt.Sprintf("%d", si.Port)))
	if err != nil {
		return err
	}
	defer conn.Close()

	err = conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		return fmt.Errorf("failed to set read deadline: %w", err)
	}

	_, err = conn.Write([]byte{0xFE, 0x01})
	if err != nil {
		return fmt.Errorf("failed to send ping: %w", err)
	}

	buffer := make([]byte, 1024)
	offset, err := conn.Read(buffer)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	si.Latency = int(time.Since(start).Milliseconds())

	buffer = buffer[:offset]

	si.Host, offset = readString(buffer)
	buffer = buffer[offset:]

	si.Map, offset = readString(buffer)
	buffer = buffer[offset:]

	si.Players, offset = readNumber(buffer)
	buffer = buffer[offset:]

	si.Waves, offset = readNumber(buffer)
	buffer = buffer[offset:]

	si.GameVersion, offset = readNumber(buffer)
	buffer = buffer[offset:]

	si.VerType, offset = readString(buffer)
	buffer = buffer[offset:]

	si.Gamemode, offset = readGamemode(buffer)
	buffer = buffer[offset:]

	si.Limit, offset = readNumber(buffer)
	buffer = buffer[offset:]

	si.Desc, offset = readString(buffer)
	buffer = buffer[offset:]

	return nil
}

var gamemodes = map[int]string{
	0: "survival",
	1: "sandbox",
	2: "attack",
	3: "pvp",
	4: "editor",
}

func readGamemode(buffer []byte) (string, int) {
	gamemodeIndex := int(buffer[0])
	gamemode, found := gamemodes[gamemodeIndex]
	if found {
		return gamemode, 1
	}
	return "unknown", 1
}

func readString(buffer []byte) (string, int) {
	length := int(buffer[0]) + 1
	text := string(buffer[1:length])
	text = strings.ReplaceAll(text, "\r", " ")
	text = strings.ReplaceAll(text, "\n", " ")
	return text, length
}

func readNumber(buffer []byte) (int, int) {
	var number int32
	bytesReader := bytes.NewReader(buffer[:4])
	err := binary.Read(bytesReader, binary.BigEndian, &number)
	if err != nil {
		return -1, 4
	}
	return int(number), 4
}

func GetServerInfo(address string, port int) (*ServerInfo, error) {
	si := &ServerInfo{Address: address, Port: port}
	if err := si.Update(); err != nil {
		return &ServerInfo{}, err
	}
	return si, nil
}
