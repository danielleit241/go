package internal

type Service interface {
	DoBusinessLogic()
}

type Builder interface {
	reset()
	buildName(name string)
	buildLogger(logger Logger)
	buildNotifier(notifier Notifier)
	buildDataLayer(dataLayer DataLayer)
	buildUploader(uploader Uploader)
	result() Service
}

type serviceBuilder struct {
	service *complexService
}

func NewBuilder() *serviceBuilder {
	return &serviceBuilder{}
}

func (b *serviceBuilder) reset() {
	b.service = &complexService{}
}

func (b *serviceBuilder) buildName(name string) {
	b.service.name = name
}

func (b *serviceBuilder) buildLogger(logger Logger) {
	if logger == nil {
		logger = &StdLogger{}
	}
	b.service.logger = logger
}

func (b *serviceBuilder) buildNotifier(notifier Notifier) {
	if notifier == nil {
		notifier = &EmailNotifier{}
	}
	b.service.notifier = notifier
}

func (b *serviceBuilder) buildDataLayer(dataLayer DataLayer) {
	if dataLayer == nil {
		dataLayer = &MySQLDataLayer{}
	}
	b.service.dataLayer = dataLayer
}

func (b *serviceBuilder) buildUploader(uploader Uploader) {
	if uploader == nil {
		uploader = &FirebaseUploader{}
	}
	b.service.uploader = uploader
}

func (b *serviceBuilder) result() Service {
	return b.service
}
