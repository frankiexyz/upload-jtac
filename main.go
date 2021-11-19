package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/pkg/sftp"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	user := "jtac"
	pass := "anonymous"
	remote := "sftp.juniper.net"
	port := ":22"
	var caseid = flag.String("caseid", "", "Input Your caseid.")
	var filePath = flag.String("file", "", "Input Your file like /tmp/log.tar.gz")
	flag.Parse()
	remotePath := "/pub/incoming/" + *caseid
	hostKey := getHostKey(remote)
	if *caseid == "" || *filePath == "" {
		log.Fatal("Please provide caseid and filePath")
		os.Exit(1)
	}

	log.Info("case id:" + *caseid)
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		HostKeyCallback: ssh.FixedHostKey(hostKey),
	}

	conn, err := ssh.Dial("tcp", remote+port, config)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	log.Info("Trying to connect to ", remote)
	client, err := sftp.NewClient(conn)
	if err != nil {
		log.Fatal("fail to connect to sftp", err)
	}
	defer client.Close()

	_, err = client.ReadDir(remotePath)
	if err != nil {
		log.Info("Trying to create folder ", remotePath)
		err = client.Mkdir(remotePath)
		if err != nil {
			log.Fatal("failed to create dir", err)
		}
	}
	file := strings.Split(*filePath, "/")
	remoteFilePath := remotePath + "/" + file[len(file)-1]
	log.Info("trying to create file on sftp ", remoteFilePath)
	dstFile, err := client.Create(remoteFilePath)
	if err != nil {
		log.Fatal("fail to create file ", err)
	}
	defer dstFile.Close()

	srcFile, err := os.Open(*filePath)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("trying to upload file to sftp " + remotePath)
	bytes, err := io.Copy(dstFile, srcFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d bytes copied\n", bytes)
}

func getHostKey(host string) ssh.PublicKey {
	file, err := os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hostKey ssh.PublicKey
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		if len(fields) != 3 {
			continue
		}
		if strings.Contains(fields[0], host) {
			var err error
			hostKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
			if err != nil {
				log.Fatalf("error parsing %q: %v", fields[2], err)
			}
			break
		}
	}

	if hostKey == nil {
		log.Fatalf("no hostkey found for %s", host)
	}

	return hostKey
}
