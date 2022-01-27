package failure

import (
	"context"
)

type ServiceImpl struct {
	Client FailureServiceClient
}

/*
   Injects a failure.

   Parameters
   ----------
   failureUnit *FailureUnit


   failureType *FailureType


   instance int32


*/

func (s *ServiceImpl) Inject(failureUnit *FailureUnit, failureType *FailureType, instance int32) (*InjectResponse, error) {

	request := &InjectRequest{}
	ctx := context.Background()
	request.FailureUnit = *failureUnit
	request.FailureType = *failureType
	request.Instance = instance
	response, err := s.Client.Inject(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
