package api

import (
	JobsAPI "controller/api/Jobs"
)

func (s *Server) Routes() {
	router := s.Engine
	groupStream := router.Group("/faas")
	groupV1 := groupStream.Group("/v1")

	var jobsController JobsAPI.Controller
	jobsController.Routes(groupV1, s.Messaging)
}
