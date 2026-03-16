package main

type Logger interface {
	Log(...any)
}

type StdLogger struct{}
type FileLogger struct{}

func (l *StdLogger) Log(...any)  {}
func (l *FileLogger) Log(...any) {}

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct{}
type SMSNotifier struct{}

func (n *EmailNotifier) Send(message string) {}

func (n *SMSNotifier) Send(message string) {}

type DataLayer interface {
	Save()
}

type MySQLDataLayer struct{}
type MongoDBDataLayer struct{}

func (d *MySQLDataLayer) Save()   {}
func (d *MongoDBDataLayer) Save() {}

type Uploader interface {
	Upload(file string)
}

type S3Uploader struct{}
type FirebaseUploader struct{}

func (u *S3Uploader) Upload(file string)       {}
func (u *FirebaseUploader) Upload(file string) {}

type ComplexService struct {
	name      string
	logger    Logger
	notifier  Notifier
	dataLayer DataLayer
	uploader  Uploader
}

func (s *ComplexService) SetLogger(logger Logger)          { s.logger = logger }
func (s *ComplexService) SetNotifier(notifier Notifier)    { s.notifier = notifier }
func (s *ComplexService) SetDataLayer(dataLayer DataLayer) { s.dataLayer = dataLayer }
func (s *ComplexService) SetUploader(uploader Uploader)    { s.uploader = uploader }

func (s *ComplexService) DoBusinessLogic() {
	s.logger.Log("Starting business logic")
	s.uploader.Upload("file.txt")
	s.dataLayer.Save()
	s.notifier.Send("Business logic completed")
}

// God constructor is too complex to maintain
func NewService(name string, logger Logger, notifier Notifier, dataLayer DataLayer, uploader Uploader) *ComplexService {

	if logger == nil {
		logger = &StdLogger{}
	}
	if notifier == nil {
		notifier = &EmailNotifier{}
	}
	if dataLayer == nil {
		dataLayer = &MySQLDataLayer{}
	}
	if uploader == nil {
		uploader = &FirebaseUploader{}
	}

	return &ComplexService{
		name:      name,
		logger:    logger,
		notifier:  notifier,
		dataLayer: dataLayer,
		uploader:  uploader,
	}
}

//Or convenient constructor

func NewServiceWithDefault(name string) *ComplexService {
	return NewService(name, &StdLogger{}, &EmailNotifier{}, &MySQLDataLayer{}, &FirebaseUploader{})
}

func main() {
	s := NewServiceWithDefault("ComplexService")
	s.DoBusinessLogic()
}
