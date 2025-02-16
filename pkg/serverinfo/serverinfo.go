package serverinfo

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

type ServerInfo struct {
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

var Gamemodes = []string{"survival", "sandbox", "attack", "pvp", "editor"}

func readGamemode(buffer []byte) (string, int) {
	gamemodeIndex := int(buffer[0])
	if gamemodeIndex >= len(Gamemodes) || gamemodeIndex < 0 {
		return "unknown", 1
	}
	return Gamemodes[gamemodeIndex], 1
}

func readString(buffer []byte) (string, int) {
	length := int(buffer[0]) + 1
	text := string(buffer[1:length])
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
	start := time.Now()
	conn, err := net.Dial("udp4", fmt.Sprintf("%s:%d", address, port))
	if err != nil {
		return nil, fmt.Errorf("failed to dial server: %w", err)
	}
	defer conn.Close()

	pingBuffer := []byte{0xFE, 0x01}
	_, err = conn.Write(pingBuffer)
	if err != nil {
		return nil, fmt.Errorf("failed to send ping: %w", err)
	}

	if err := conn.SetReadDeadline(time.Now().Add(5 * time.Second)); err != nil {
		return nil, fmt.Errorf("failed to set read deadline: %w", err)
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	msg := buffer[:n]

	host, offset := readString(msg)
	msg = msg[offset:]

	mapName, offset := readString(msg)
	msg = msg[offset:]

	players, offset := readNumber(msg)
	msg = msg[offset:]

	waves, offset := readNumber(msg)
	msg = msg[offset:]

	gameVersion, offset := readNumber(msg)
	msg = msg[offset:]

	verType, offset := readString(msg)
	msg = msg[offset:]

	gamemode, offset := readGamemode(msg)
	msg = msg[offset:]

	limit, offset := readNumber(msg)
	msg = msg[offset:]

	desc, offset := readString(msg)
	msg = msg[offset:]

	// modeName
	readString(msg)

	duration := time.Since(start)
	latency := int(duration.Milliseconds())

	info := &ServerInfo{
		Host:        host,
		Map:         mapName,
		Players:     players,
		Waves:       waves,
		GameVersion: gameVersion,
		VerType:     verType,
		Gamemode:    gamemode,
		Limit:       limit,
		Desc:        desc,
		Latency:     latency,
	}

	return info, nil
}
