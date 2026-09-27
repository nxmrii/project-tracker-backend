package task

import (
	"context"
	"fmt"
	"log"
)

// هذه الدالة تقوم بإبطال ذاكرة التخزين المؤقت لتقدم المشروع في Redis عند تعديل المهام المتعلقة بالمشروع.
func (s *Service) invalidateProjectProgressCache(
	ctx context.Context,
	projectID int64,
) {

	if s.cache == nil {
		return
	}

	key := fmt.Sprintf(
		"project:%d:progress",
		projectID,
	)

	err := s.cache.Delete(
		ctx,
		key,
	)

	if err != nil {
		log.Println(
			"Cache invalidation error:",
			err,
		)

		return
	}

	log.Println(
		"CACHE INVALIDATED:",
		key,
	)
}
