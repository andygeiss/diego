# DIEGO - A Data Driven Inference Engine written in GO 

[![Go Report Card](https://goreportcard.com/badge/github.com/andygeiss/diego)](https://goreportcard.com/report/github.com/andygeiss/diego)

DIEGO provides a framework to link practical user interfaces to domain specific knowledge of experts.

An expert system usually consists of three parts:
* a knowledge base (facts and rules),
* an inference engine (forward-chaining).

The inference engine uses a top-down method to take facts as they become available and
apply rules to draw conclusions.

##### Table of Contents

- [Installation](README.md#installation)
    * [From Source](README.md#from-source)
- [Usage](README.md#usage)

## Installation

### From Source

    go get -u github.com/andygeiss/diego

## Usage

The following code ...
* specifies the survey questions and options for the explanation repository and
* specifies the facts and rules/conditions for the inference engine:

```go
import (
	expRepos "github.com/andygeiss/diego/pkg/explanation/repositories"
	"github.com/andygeiss/diego/pkg/inference/engines"
	infRepos "github.com/andygeiss/diego/pkg/inference/repositories"
	"github.com/andygeiss/diego/pkg/survey/services"
)

func main() {
    // Initialize the repositories by using JSON-files.
    expRepo := expRepos.NewFileRepository("../../../testdata/explanation.json")
    infRepo := infRepos.NewFileRepository("../../../testdata/inference.json")
    // Configure the engine to use a specific survey by name.
    engine := engines.NewDefaultEngine("SURVEY NAME", infRepo)
    // Combine the explanation repository and inference engine.
    service := services.NewDefaultService(expRepo, engine)
    // Get the facts by setting a specific condition.
    facts, err := service.GetResultsByFacts([]string{"#Q1 = 1"})
    ...    
}
```

See the [testdata](https://github.com/andygeiss/diego/tree/master/testdata) directory for examples.
