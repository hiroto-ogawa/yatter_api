package timelines

import (
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/domain/object"
)

func (h *handler) GetPublic(w http.ResponseWriter, r *http.Request) {
	// maxIDStr := r.URL.Query().Get("max_id")
	// sinceIDStr := r.URL.Query().Get("since_id")
	// limitStr := r.URL.Query().Get("limit")

	// maxID, err := strconv.ParseInt(maxIDStr, 10, 64)
	// if err != nil {
	// 	http.Error(w, "Malformed parameter max_id", http.StatusBadRequest)
	// 	return
	// }

	// sinceID, err := strconv.ParseInt(sinceIDStr, 10, 64)
	// if err != nil {
	// 	http.Error(w, "Malformed parameter since_id", http.StatusBadRequest)
	// 	return
	// }

	// limit, err := strconv.ParseInt(limitStr, 10, 64)
	// if err != nil {
	// 	http.Error(w, "Malformed parameter limit", http.StatusBadRequest)
	// 	return
	// }

	timeline := new(object.Timeline)
	// timeline.MaxID = &maxID
	// timeline.SinceID = &sinceID
	// timeline.Limit = &limit

	statuses, err := h.tr.GetPublicTimeline(timeline)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(statuses); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
