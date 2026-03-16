package internal

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

type complexService struct {
	name      string
	logger    Logger
	notifier  Notifier
	dataLayer DataLayer
	uploader  Uploader
}

func (s *complexService) setLogger(logger Logger)          { s.logger = logger }
func (s *complexService) setNotifier(notifier Notifier)    { s.notifier = notifier }
func (s *complexService) setDataLayer(dataLayer DataLayer) { s.dataLayer = dataLayer }
func (s *complexService) setUploader(uploader Uploader)    { s.uploader = uploader }

func (s *complexService) DoBusinessLogic() {
	s.logger.Log("Starting business logic")
	s.uploader.Upload("file.txt")
	s.dataLayer.Save()
	s.notifier.Send("Business logic completed")
}
