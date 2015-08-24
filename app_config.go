package main

type appConfig struct {
	name   string
	author string
}

func (this *appConfig) SetAppName(appName string) {
	this.name = appName
}

func (this *appConfig) GetAppName() string {
	return this.name
}

func (this *appConfig) SetAuthorName(authorName string) {
	this.author = authorName
}

func (this *appConfig) GetAuthorName() string {
	return this.author
}
