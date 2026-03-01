package usercase

import (
	er "Board_of_issuses/internal/core"
	dn "Board_of_issuses/internal/core/domains"
	"context"
	"time"
)

func (s *Service) GetAllTasks(ctx context.Context, userId, deskID int) ([]dn.Task, error) {
	if err := s.accessVerificationToDeskForUser(ctx, userId, deskID); err != nil {
		return nil, err
	}

	repoTasks, err := s.repo.GetAllTasksFromOneDesk(ctx, deskID)
	if err != nil {
		return nil, err
	}

	servTasks := make([]dn.Task, len(repoTasks))

	for i := 0; i < len(repoTasks); i++ {
		repoTask := repoTasks[i]

		servTask := dn.Task{
			Id:          repoTask.Id,
			UserId:      repoTask.UserId,
			DeskId:      repoTask.DeskId,
			Name:        repoTask.Name,
			Description: repoTask.Description,
			Done:        repoTask.Done,
			Time:        repoTask.Time,
			Created_at:  repoTask.Created_at,
		}

		servTasks[i] = servTask
	}

	return servTasks, nil
}

func (s *Service) GetTasksWithParams(ctx context.Context, userId, deskID int, done bool) ([]dn.Task, error) {
	if err := s.accessVerificationToDeskForUser(ctx, userId, deskID); err != nil {
		return nil, err
	}

	repoTasks, err := s.repo.GetTasksWithParams(ctx, deskID, done)
	if err != nil {
		return nil, err
	}

	servTasks := make([]dn.Task, len(repoTasks))

	for i := 0; i < len(repoTasks); i++ {
		repoTask := repoTasks[i]

		servTask := dn.Task{
			Id:          repoTask.Id,
			UserId:      repoTask.UserId,
			DeskId:      repoTask.DeskId,
			Name:        repoTask.Name,
			Description: repoTask.Description,
			Done:        repoTask.Done,
			Time:        repoTask.Time,
			Created_at:  repoTask.Created_at,
		}

		servTasks[i] = servTask
	}

	return servTasks, nil
}

func (s *Service) CreateTask(ctx context.Context, task *dn.Task) error {
	if err := s.accessVerificationToDeskForUser(ctx, task.UserId, task.DeskId); err != nil {
		return err
	}

	if err := s.repo.CreateTask(ctx, task.ToRepoTask()); err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteTask(ctx context.Context, taskID, userID int) error {
	ownerID, err := s.repo.GetTaskOwner(ctx, taskID)
	if err != nil {
		return err
	}

	if ownerID != userID {
		return er.UserNotOwnerOfTask(userID, taskID)
	}

	if err := s.repo.DeleteTask(ctx, taskID); err != nil {
		return err
	}

	return nil
}

func (s *Service) ChangeTaskDescription(ctx context.Context, userID, taskID int, description string) error {
	ownerID, err := s.repo.GetTaskOwner(ctx, taskID)
	if err != nil {
		return err
	}

	if ownerID != userID {
		return er.UserNotOwnerOfTask(userID, taskID)
	}

	if err := s.repo.UpdateTaskDecription(ctx, taskID, description); err != nil {
		return err
	}

	return nil

}

func (s *Service) UpdateTaskTime(ctx context.Context, userID, taskID, userTime int) error {

	ownerID, err := s.repo.GetTaskOwner(ctx, taskID)
	if err != nil {
		return err
	}

	if ownerID != userID {
		return er.UserNotOwnerOfTask(userID, taskID)
	}

	resultTime := time.Now().Add(time.Hour * time.Duration(userTime))

	if err := s.repo.UpdateTaskTime(ctx, taskID, resultTime); err != nil {
		return err
	}

	return nil

}

func (s *Service) ComplyteTask(ctx context.Context, userID, taskID int) error {
	deskID, err := s.repo.GetDeskIDByTask(ctx, taskID)
	if err != nil {
		return err
	}

	if err := s.accessVerificationToDeskForUser(ctx, userID, deskID); err != nil {
		return err
	}

	if err := s.repo.UpdateTaskDone(ctx, taskID); err != nil {
		return err
	}

	return nil
}
