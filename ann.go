package ann

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

// nNet menyimpan semua informasi yang mendefinisikan jaringan syaraf yang sudah dilatih. 
type nNet struct {
	config 	nNetConf
	wHide	*mat.Dense
	bHide	*mat.Dense
	wOut	*mat.Dense
	bOut	*mat.Dense
}

//nNetConf mendefinisikan arsitektur dan parameter pembelajaran dari jaringan syaraf yang kita miliki.
type nNetConf struct {
	inputSyaraf 	int
	outputSyaraf 	int
	hiddenSyaraf 	int
	jumEpoch		int
	rateLearning	float64
}

func main() {

}