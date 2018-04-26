package pkg

// InvalidBlockError is raised when we find an invalid block
type InvalidBlockError struct {
	message string
	block   Block
}

// NewInvalidBlockError returns an error that formats as the given text.
func NewInvalidBlockError(message string, block *Block) *InvalidBlockError {
	return &InvalidBlockError{
		message: message,
	}
}
func (e *InvalidBlockError) Error() string {
	return e.message
}

// BlockNotFoundError is raised when we can't find an indes in the block
type BlockNotFoundError struct {
	message string
	block   Block
}

// NewBlockNotFoundError returns an error that formats as the given text.
func NewBlockNotFoundError(message string, index int) *BlockNotFoundError {
	return &BlockNotFoundError{
		message: message,
	}
}
func (e *BlockNotFoundError) Error() string {
	return e.message
}

// EnvNotFoundError is raised when we can't find an indes in the block
type EnvNotFoundError struct {
	message string
}

// NewEnvNotFoundError returns an error that formats as the given text.
func NewEnvNotFoundError(message string) *EnvNotFoundError {
	return &EnvNotFoundError{
		message: message,
	}
}
func (e *EnvNotFoundError) Error() string {
	return e.message
}
