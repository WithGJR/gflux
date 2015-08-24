package main

import (
	"fmt"
	"os"
)

type rollbackType int

const (
	beforeCreatedFiles rollbackType = iota
	beforeCreatedDirectories
)

type Generator struct {
	appConfig
	commander
	errorReporter
	createdDirectories []string
	createdFiles       []string
}

func (this *Generator) rootPath() string {
	return this.GetAppName() + "/"
}

func (this *Generator) createDirectory(dirName string) error {
	this.createdDirectories = append(this.createdDirectories, dirName)
	if err := os.Mkdir(dirName, 0751); err != nil && os.IsExist(err) != true {
		return err
	}
	return nil
}

func (this *Generator) createRootDirectory() error {
	return this.createDirectory(this.rootPath())
}

func (this *Generator) CreateDirectories(message string, dirNames []string) {
	if this.err != nil || this.createRootDirectory() != nil {
		return
	}
	fmt.Println(message)

	for i := 0; i < len(dirNames); i++ {
		if err := this.createDirectory(this.rootPath() + dirNames[i]); err != nil {
			this.rollback(beforeCreatedDirectories, i-1)
			this.log(err)
			return
		}
	}
}

func (this *Generator) rollback(rollbackTyp rollbackType, index int) {
	var removeFunc func()

	switch rollbackTyp {
	case beforeCreatedFiles:
		removeFunc = func() {
			os.Remove(this.createdFiles[index])
		}
	case beforeCreatedDirectories:
		removeFunc = func() {
			os.Remove(this.createdDirectories[index])
		}
	}

	for i := index; i > 0; i-- {
		removeFunc()
	}
}

func (this *Generator) writeFile(fileName string, content string) error {
	err := func() error {
		file, err := os.Create(fileName)
		if err != nil {
			return err
		}

		if _, err := file.WriteString(content); err != nil {
			return err
		}

		this.createdFiles = append(this.createdFiles, fileName)
		return nil
	}()

	if err != nil {
		return err
	}
	return nil
}

func (this *Generator) WriteFiles(message string, fileInfos [][2]string) {
	if this.err != nil {
		return
	}
	fmt.Println(message)

	for i := 0; i < len(fileInfos); i++ {
		fileName, fileContent := fileInfos[i][0], fileInfos[i][1]
		if err := this.writeFile(this.rootPath()+fileName, fileContent); err != nil {
			this.rollback(beforeCreatedFiles, i-1)
			this.log(err)
			return
		}
	}
}

func (this *Generator) Run(message string, commandName string, arg ...string) {
	if this.err != nil {
		return
	}
	fmt.Println(message)

	if err := this.run(commandName, arg...); err != nil {
		this.log(err)
	}
}
