# Simple AI With Golang

simple use

1. set config go-ai, ex:
``` go
ai := goai.New(goai.Config{
	IsLower:       true,
	IsBinary:      false,
	Smoot:         true,
	NormalizeType: goai.EuclideanSumSquare,
})
```

2. learn data:
``` go
ai.Learn(data) // data is map[string][]string
```

3. set text you wont to predict, and predict
``` go
dataTest := []string{
	"pak tiket ke jakarta ada?",
}

predicted, err := ai.Predict(dataTest)
if err != nil {
	//handle error
}
```

