// Shows different functionality that can be achieved with a map
package main

import "fmt"

func states() {
	states := map[string]string{"tn": "Tennessee", "mi": "Michigan", "mt": "Montana", "al": "Alabama", "wa": "Washington"}
	fmt.Println(states)
	fmt.Println(len(states))
	fmt.Println(states["tn"], states["mt"], states["wa"])

}

func stateCodes() {
	stateCodes := map[string]int{"tn": 1, "mi": 2, "mt": 3, "al": 3, "wa": 5}
	value, found := stateCodes["wa"]
	fmt.Println(found, value)
	value, found = stateCodes["ca"]
	fmt.Println(found, value)

}

func stateAdd() {
	states := map[string]string{"tn": "Tennessee", "mi": "Michigan", "mt": "Montana", "al": "Alabama", "wa": "Washington"}
	states["ca"] = "California"
	fmt.Println(states)

}

func stateUpdate() {
	states := map[string]string{"tn": "Tennessee", "mi": "Michigan", "mt": "Montana", "al": "Alabama", "wa": "Washington"}
	states["wa"] = "Seattle, WA"
	fmt.Println(states)

}

func stateDelete() {
	states := map[string]string{"tn": "Tennessee", "mi": "Michigan", "mt": "Montana", "al": "Alabama", "wa": "Washington"}
	fmt.Println(states)
	delete(states, "al")
	fmt.Println(states)
}

func stateLoop() {
	states := map[string]string{"tn": "Tennessee", "mi": "Michigan", "mt": "Montana", "al": "Alabama", "wa": "Washington"}
	for key, value := range states {
		fmt.Println(key, ":", value)
	}
}

func stateTruncate() {
	states := map[string]string{"tn": "Tennessee", "mi": "Michigan", "mt": "Montana", "al": "Alabama", "wa": "Washington"}
	for key := range states {
		delete(states, key)
	}
	fmt.Println(states)
}

func stateZeroOut() {
	states := map[string]string{"tn": "Tennessee", "mi": "Michigan", "mt": "Montana", "al": "Alabama", "wa": "Washington"}
	states = make(map[string]string)
	fmt.Println(states)
}

func main() {
	states()
	stateCodes()
	stateAdd()
	stateUpdate()
	stateDelete()
	stateLoop()
	stateTruncate()
	stateZeroOut()

}
