package project

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	projectDTO "project-tracker-backend/internal/dtos/project"

	"github.com/redis/go-redis/v9"
)

func (s *Service) GetProgress(
	ctx context.Context,
	projectID int64,
) (*projectDTO.ProgressResponse, error) {

	if projectID <= 0 {
		return nil, errors.New("invalid project id")
	}

	if s.taskRepo == nil {
		return nil, errors.New(
			"task repository is not configured",
		)
	}

	cacheKey := fmt.Sprintf(
		"project:%d:progress",
		projectID,
	)

	// 1. ابحث في Redis أولًا
	if s.cache != nil {

		cachedData, err :=
			s.cache.Get(ctx, cacheKey)

		if err == nil {

			var cachedProgress projectDTO.ProgressResponse

			if err := json.Unmarshal(
				[]byte(cachedData),
				&cachedProgress,
			); err == nil {

				log.Println(
					"CACHE HIT:",
					cacheKey,
				)

				return &cachedProgress, nil
			}
		}

		if err != nil &&
			err != redis.Nil {

			log.Println(
				"Redis get error:",
				err,
			)
		}

		log.Println(
			"CACHE MISS:",
			cacheKey,
		)
	}

	// 2. لم نجدها في Redis
	// نذهب إلى PostgreSQL
	totalTasks, completedTasks, err :=
		s.taskRepo.GetProjectTaskCounts(
			ctx,
			projectID,
		)

	if err != nil {
		return nil, err
	}

	progress := 0.0

	if totalTasks > 0 {
		progress =
			float64(completedTasks) /
				float64(totalTasks) *
				100
	}

	result := &projectDTO.ProgressResponse{
		ProjectID:      projectID,
		TotalTasks:     totalTasks,
		CompletedTasks: completedTasks,
		Progress:       progress,
	}

	// 3. خزّن النتيجة في Redis
	if s.cache != nil {

		data, err := json.Marshal(result)

		if err == nil {

			err = s.cache.Set(
				ctx,
				cacheKey,
				string(data),
				// يعيش لمدة خمس دقائق فقطاـCache تعني ان
				5*time.Minute,
			)

			if err != nil {
				log.Println(
					"Redis set error:",
					err,
				)
			}
		}
	}

	return result, nil
}
