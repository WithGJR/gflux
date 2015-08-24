package main

type errorReporter struct {
	err error
}

func (this *errorReporter) log(err error) {
	this.err = err
}

func (this *errorReporter) Err() error {
	return this.err
}
