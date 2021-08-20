package main

import "net/http"

func test(w http.ResponseWriter, r *http.Request) {

	podsList := []Project{
		{ID: "1", PodName: "V1"},
		{ID: "2", PodName: "V2"},
		{ID: "3", PodName: "V3"},
		{ID: "4", PodName: "V4"},
		{ID: "5", PodName: "V5"},
	}

	podsProject := map[string]Project{}

	// for _, v := range podsList {
	// 	fmt.Println("Pod ID:", v.ID, "PodName:", v.PodName)
	// 	podsProject[v.ID] = v
	// }

	// for m, n := range podsProject {
	// 	n.PodName = "UpdatedData for " + n.ID
	// 	podsProject[m] = n
	// 	fmt.Println("Data key:", m, "Value:", n.PodName)
	// }

	for _, v := range podsList {
		podsProject[v.ID] = v
	}

	for id, pod := range podsProject {
		w.Write([]byte("Pod ID: " + id + " Value: " + pod.PodName))
		w.Write([]byte("\n"))
	}

}
