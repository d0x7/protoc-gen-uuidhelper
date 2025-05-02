package main
var version = "v0.0.1+dirty"
func (b backend) Version() string {
  return version
}
