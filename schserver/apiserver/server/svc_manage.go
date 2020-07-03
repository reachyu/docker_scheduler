package server

import "schserver/apiserver/server/dockersvc"

type SvcManager struct {
	dockerSvc  dockersvc.DockerSvcInterface
}

func (r *SvcManager) CreateContainerInBackground(imageName string) string {
	return r.dockerSvc.CreateContainerInBackground(imageName)
}

func (r *SvcManager) StopByContainerId(containerId string) error {
	return r.dockerSvc.StopByContainerId(containerId)
}

func (r *SvcManager) DeleteByContainerId(containerId string) error {
	return r.dockerSvc.DeleteByContainerId(containerId)
}

func (r *SvcManager) ListContainers()  {
	r.dockerSvc.ListContainers()
}

func NewSvcManager() *SvcManager {
	return &SvcManager{
		dockerSvc:  dockersvc.NewDockerSvc(),
	}
}