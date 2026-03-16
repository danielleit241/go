package internal

type ServiceDirector interface {
	BuildService(builder Builder) Service
}

type serviceBuilderDirector struct{}

func NewDirector() serviceBuilderDirector {
	return serviceBuilderDirector{}
}

func (sbd *serviceBuilderDirector) BuildService(builder Builder) Service {
	builder.reset()
	builder.buildName("ComplexService")
	builder.buildLogger(&StdLogger{})
	builder.buildNotifier(&EmailNotifier{})
	builder.buildDataLayer(&MySQLDataLayer{})
	builder.buildUploader(&FirebaseUploader{})
	return builder.result()
}
