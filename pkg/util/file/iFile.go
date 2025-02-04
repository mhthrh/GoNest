package file

type IFile interface {
	Read() ([]byte, error)
	Write([]byte) error
}
