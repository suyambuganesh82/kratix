package pipeline

type pipelineArgs struct {
	names map[string]string
}

func newPipelineArgs(promiseIdentifier, resourceRequestIdentifier, namespace string) pipelineArgs {
	names := map[string]string{
		"configure-pipeline-name": pipelineName("configure", promiseIdentifier),
		"promise-id":              promiseIdentifier,
		"resource-request-id":     resourceRequestIdentifier,
		"service-account":         promiseIdentifier + "-promise-pipeline",
		"role":                    promiseIdentifier + "-promise-pipeline",
		"role-binding":            promiseIdentifier + "-promise-pipeline",
		"config-map":              "scheduling-" + promiseIdentifier,
		"namespace":               namespace,
	}
	return pipelineArgs{
		names: names,
	}
}

func (p pipelineArgs) ConfigMapName() string {
	return p.names["config-map"]
}

func (p pipelineArgs) ServiceAccountName() string {
	return p.names["service-account"]
}

func (p pipelineArgs) RoleName() string {
	return p.names["role"]
}

func (p pipelineArgs) RoleBindingName() string {
	return p.names["role-binding"]
}

func (p pipelineArgs) Namespace() string {
	return p.names["namespace"]
}

func (p pipelineArgs) PromiseID() string {
	return p.names["promise-id"]
}

func (p pipelineArgs) ResourceRequestID() string {
	return p.names["resource-request-id"]
}

func (p pipelineArgs) ConfigurePipelineName() string {
	return p.names["configure-pipeline-name"]
}

func (p pipelineArgs) Labels() pipelineLabels {
	return newPipelineLabels().WithPromiseID(p.PromiseID())
}

func (p pipelineArgs) PipelinePodLabels() pipelineLabels {
	return p.Labels().
		WithResourceRequestID(p.ResourceRequestID())
}
