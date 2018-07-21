package glb

import (
	"strings"
	"github.com/satori/go.uuid"
)

/**
	
 */

func UUID() string {
	return strings.Replace(uuid.Must(uuid.NewV4()).String(), "-", "", 4)
}

