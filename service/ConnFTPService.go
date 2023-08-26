package service

import (
	"FILEMANAGESYS/entity"
	"github.com/jlaffaye/ftp"
)

// 连接FTP服务器并返回客户端
func ConnectFTP(config entity.FTPConfig) (*ftp.ServerConn, error) {
	// 创建FTP客户端
	client, err := ftp.Dial(config.Host + ":" + config.Port)
	if err != nil {
		return nil, err
	}

	// 登录FTP服务器
	err = client.Login(config.User, config.Pass)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// 获取文件信息
func GetFileInfo(client *ftp.ServerConn, path string) ([]*ftp.Entry, error) {
	// 确定文件路径
	err := client.ChangeDir(path)
	if err != nil {
		return nil, err
	}

	// 当前路径下的所有文件
	files, err := client.List("")
	if err != nil {
		return nil, err
	}

	return files, nil
}
