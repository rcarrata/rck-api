package api

import "net/http"

type Test struct {
	ID      string `json:"id"`
	PodName string `json:"value"`
}

func test(w http.ResponseWriter, r *http.Request) {

	podsList := []Test{
		{ID: "1", PodName: "V1"},
		{ID: "2", PodName: "V2"},
		{ID: "3", PodName: "V3"},
		{ID: "4", PodName: "V4"},
		{ID: "5", PodName: "V5"},
	}

	podsTest := map[string]Test{}

	// for _, v := range podsList {
	// 	fmt.Println("Pod ID:", v.ID, "PodName:", v.PodName)
	// 	podsTest[v.ID] = v
	// }

	// for m, n := range podsTest {
	// 	n.PodName = "UpdatedData for " + n.ID
	// 	podsTest[m] = n
	// 	fmt.Println("Data key:", m, "Value:", n.PodName)
	// }

	for _, v := range podsList {
		podsTest[v.ID] = v
	}

	for id, pod := range podsTest {
		w.Write([]byte("Pod ID: " + id + " Value: " + pod.PodName))
		w.Write([]byte("\n"))
	}

}
