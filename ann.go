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
	jumEpoch	int
	rateLearning	float64
}

func main() {
	//membuka file dataset train`
	data, err := os.Open("train.csv")
	
	//arsitektur jaringan dan parameter pembelajaran
	config := nNetconf{
		inNeurons: 4,
		ouNeurons: 3,
		hiNeurons: 3,
		nuEpochs: 5000,
		leRate: 0.3,
	}
	
	//train jaringan syaraf
	net := newNet(config)
	if err := net.train(inputs, labels); err !nil {
		log.Fatal(err)
	}
	
	//

}
