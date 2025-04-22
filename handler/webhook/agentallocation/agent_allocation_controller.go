package agentallocation

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Manage handles agent allocation requests.
func (co *Controller) Manage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req Request

	// Parse JSON body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("error binding request, err: %v", err)
		http.Error(w, `{"message":"error binding request"}`, http.StatusBadRequest)
		return
	}

	// Push to Redis
	if err := co.redis.LPush(ctx, "agent-allocator-queue",
		fmt.Sprintf("%s:%s", req.AppId, req.RoomId)).Err(); err != nil {
		log.Printf("Failed to queue request: %v", err)
		http.Error(w, `{"error":"Failed to queue request"}`, http.StatusInternalServerError)
		return
	}

	// Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	resp := co.transformToResponse(ctx, "Success save agent allocator into queue")
	resp.Status = http.StatusCreated

	json.NewEncoder(w).Encode(resp)
}
